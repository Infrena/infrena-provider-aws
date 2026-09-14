package awsprov

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/provider"
)

// sdkClient is a real EC2 client with the SDK's standard retryer, making `attempts` attempts with no
// wait between them: the retry decisions are the SDK's own, only the backoff is removed so a throttled
// test does not sit through seconds of real delay.
func sdkClient(endpoint string, attempts int) *ec2.Client {
	return ec2.New(ec2.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentialsProvider("AKIDTEST", "secret", ""),
		Retryer: retry.NewStandard(func(o *retry.StandardOptions) {
			o.MaxAttempts = attempts
			o.Backoff = retry.BackoffDelayerFunc(func(int, error) (time.Duration, error) { return 0, nil })
		}),
	})
}

// TestClassificationOfRealSDKErrors. Every case goes through the real SDK against the fake, so the
// error chains are the ones production sees, not ones built by hand to match the classifier.
func TestClassificationOfRealSDKErrors(t *testing.T) {
	cases := []struct {
		name  string
		fault ec2fake.Fault
		want  provider.Retryability
	}{
		// A throttle carrying a 503: the code must win over the status.
		{"throttle", ec2fake.Fault{Status: 503, Code: "RequestLimitExceeded", Message: "slow down"}, provider.SafeToRetry},
		{"server fault", ec2fake.Fault{Status: 500, Code: "InternalError", Message: "oops"}, provider.ConditionallyRetryable},
		{"unavailable", ec2fake.Fault{Status: 503, Code: "Unavailable", Message: "later"}, provider.ConditionallyRetryable},
		{"validation", ec2fake.Fault{Status: 400, Code: "InvalidParameterValue", Message: "bad cidr"}, provider.NotSafeToRetry},
		{"access", ec2fake.Fault{Status: 403, Code: "UnauthorizedOperation", Message: "no"}, provider.NotSafeToRetry},
		{"dropped after acting", ec2fake.Fault{Drop: true}, provider.ConditionallyRetryable},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fake := ec2fake.New()
			defer fake.Close()
			c.fault.Action, c.fault.Nth = "CreateVpc", 1
			fake.Inject(c.fault)
			_, err := sdkClient(fake.URL, 1).CreateVpc(context.Background(), &ec2.CreateVpcInput{CidrBlock: aws.String("10.0.0.0/16")})
			if err == nil {
				t.Fatal("expected the injected failure")
			}
			if got := classify(err); got != c.want {
				t.Errorf("classify(%v) = %v, want %v", err, got, c.want)
			}
		})
	}
}

// TestAnErrorTheSDKGaveUpOnIsStillClassified. MaxAttemptsError wraps the last error; classification
// must see through it.
func TestAnErrorTheSDKGaveUpOnIsStillClassified(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	for n := 1; n <= 3; n++ {
		fake.Inject(ec2fake.Fault{Action: "DescribeVpcs", Nth: n, Status: 503, Code: "RequestLimitExceeded"})
	}
	_, err := sdkClient(fake.URL, 3).DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	if err == nil || fake.Calls("DescribeVpcs") != 3 {
		t.Fatalf("err = %v after %d calls; want failure after 3", err, fake.Calls("DescribeVpcs"))
	}
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

// TestARefusedConnectionIsSafeToRetry. Nothing was sent, so nothing can have happened.
func TestARefusedConnectionIsSafeToRetry(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()
	_, err = sdkClient("http://"+addr, 1).CreateVpc(context.Background(), &ec2.CreateVpcInput{CidrBlock: aws.String("10.0.0.0/16")})
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify(%v) = %v, want SafeToRetry", err, got)
	}
}

func TestCancellationAndUnknownsAreNotSafe(t *testing.T) {
	for _, err := range []error{context.Canceled, errors.New("something new")} {
		if got := classify(err); got != provider.NotSafeToRetry {
			t.Errorf("classify(%v) = %v", err, got)
		}
	}
	if got := classify(context.DeadlineExceeded); got != provider.ConditionallyRetryable {
		t.Errorf("classify(DeadlineExceeded) = %v", got)
	}
}

// TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause.
func TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	fake.Inject(ec2fake.Fault{Action: "DeleteVpc", Nth: 1, Status: 400, Code: "DependencyViolation", Message: "has dependencies"})
	_, err := sdkClient(fake.URL, 1).DeleteVpc(context.Background(), &ec2.DeleteVpcInput{VpcId: aws.String("vpc-1")})
	p := &Provider{instance: "prod"}
	wrapped := p.failed("DeleteVpc", "us-east-1", "vpc-1", err)
	for _, want := range []string{`"prod"`, "DeleteVpc", "us-east-1/vpc-1", "DependencyViolation", "has dependencies", "fake-request-err"} {
		if !strings.Contains(wrapped.Error(), want) {
			t.Errorf("%q lacks %q", wrapped, want)
		}
	}
	if !hasCode(wrapped, "DependencyViolation") || classify(wrapped) != provider.NotSafeToRetry {
		t.Error("wrapping lost the cause")
	}
}
