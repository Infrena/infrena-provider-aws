package ccprov

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/provider"
)

func create(t *testing.T, endpoint string, attempts int) error {
	t.Helper()
	_, err := cloudcontrol.NewFromConfig(testConfig(endpoint, attempts)).CreateResource(context.Background(), &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`), ClientToken: aws.String("t"),
	})
	return err
}

// TestClassificationOfRealSDKErrors. Every case goes through the real SDK against the fake, so the error chains are
// the ones production sees. Statuses and codes are the Cloud Control model's (Verification log).
func TestClassificationOfRealSDKErrors(t *testing.T) {
	cases := []struct {
		status int
		code   string
		want   provider.Retryability
	}{
		{429, "ThrottlingException", provider.SafeToRetry},
		{409, "ConcurrentOperationException", provider.SafeToRetry},
		{409, "ResourceConflictException", provider.SafeToRetry},
		{502, "HandlerInternalFailureException", provider.ConditionallyRetryable},
		{502, "ServiceInternalErrorException", provider.ConditionallyRetryable},
		{502, "NetworkFailureException", provider.ConditionallyRetryable},
		{500, "ConcurrentModificationException", provider.ConditionallyRetryable},
		{400, "NotStabilizedException", provider.ConditionallyRetryable},
		{400, "InvalidRequestException", provider.NotSafeToRetry},
		{409, "ClientTokenConflictException", provider.NotSafeToRetry},
		{400, "AlreadyExistsException", provider.NotSafeToRetry},
		{401, "InvalidCredentialsException", provider.NotSafeToRetry},
	}
	for _, c := range cases {
		t.Run(c.code, func(t *testing.T) {
			fake := ccfake.New()
			defer fake.Close()
			fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: 1, Status: c.status, Code: c.code, Message: "injected"})
			err := create(t, fake.URL, 1)
			if err == nil {
				t.Fatal("expected the injected failure")
			}
			if got := classify(err); got != c.want {
				t.Errorf("classify(%v) = %v, want %v", err, got, c.want)
			}
			if errorCode(err) != c.code {
				t.Errorf("errorCode = %q, want %q", errorCode(err), c.code)
			}
		})
	}
}

// TestAnErrorTheSDKGaveUpOnIsStillClassified. The SDK wraps the last attempt's error; classification sees through it.
func TestAnErrorTheSDKGaveUpOnIsStillClassified(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	for n := 1; n <= 3; n++ {
		fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: n, Status: 429, Code: "ThrottlingException"})
	}
	err := create(t, fake.URL, 3)
	if err == nil || fake.Calls("CreateResource") != 3 {
		t.Fatalf("err = %v after %d calls; want failure after 3", err, fake.Calls("CreateResource"))
	}
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

func TestARefusedConnectionIsSafeToRetry(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()
	if got := classify(create(t, "http://"+addr, 1)); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

// TestHandlerFailuresAreClassifiedByTheirCode. A FAILED request was accepted and may have acted.
func TestHandlerFailuresAreClassifiedByTheirCode(t *testing.T) {
	cases := map[string]provider.Retryability{
		"NotStabilized": provider.ConditionallyRetryable, "ServiceInternalError": provider.ConditionallyRetryable,
		"InternalFailure": provider.ConditionallyRetryable, "NetworkFailure": provider.ConditionallyRetryable,
		"ServiceTimeout": provider.ConditionallyRetryable, "Throttling": provider.ConditionallyRetryable,
		"ResourceConflict": provider.ConditionallyRetryable,
		"AlreadyExists":    provider.NotSafeToRetry, "InvalidRequest": provider.NotSafeToRetry, "AccessDenied": provider.NotSafeToRetry,
		"NotUpdatable": provider.NotSafeToRetry, "ServiceLimitExceeded": provider.NotSafeToRetry, "Cancelled": provider.NotSafeToRetry,
	}
	for code, want := range cases {
		err := error(&HandlerError{Operation: "CREATE", Code: code, Message: "m", Token: "req-1"})
		if got := classify(err); got != want {
			t.Errorf("classify(%s) = %v, want %v", code, got, want)
		}
		if errorCode(err) != code {
			t.Errorf("errorCode(%s) = %q", code, errorCode(err))
		}
	}
}

func TestCancellationAndUnknownsAreNotSafe(t *testing.T) {
	for _, err := range []error{nil, context.Canceled, errors.New("something new")} {
		if got := classify(err); got != provider.NotSafeToRetry {
			t.Errorf("classify(%v) = %v", err, got)
		}
	}
	if got := classify(context.DeadlineExceeded); got != provider.ConditionallyRetryable {
		t.Errorf("classify(DeadlineExceeded) = %v", got)
	}
}

// TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause. Never anything from the request: it may hold a secret.
func TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: 1, Status: 400, Code: "InvalidRequestException", Message: "CidrBlock is malformed"})
	err := failure("prod", "CreateResource", mustType(t, "aws.vpc"), "us-east-1", create(t, fake.URL, 1))
	for _, want := range []string{`"prod"`, "CreateResource", "aws.vpc", "AWS::EC2::VPC", "us-east-1", "InvalidRequestException", "CidrBlock is malformed", "fake-request-err"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%q lacks %q", err, want)
		}
	}
	if errorCode(err) != "InvalidRequestException" || classify(err) != provider.NotSafeToRetry {
		t.Error("wrapping lost the cause")
	}

	denied := failure("prod", "create", mustType(t, "aws.role"), "global",
		&HandlerError{Operation: "CREATE", Code: "AccessDenied", Message: "not authorized to perform iam:CreateRole", Token: "req-9"})
	for _, want := range []string{"AccessDenied", "iam:CreateRole", "req-9", "IAM permissions"} {
		if !strings.Contains(denied.Error(), want) {
			t.Errorf("%q lacks %q", denied, want)
		}
	}
	exists := failure("prod", "create", mustType(t, "aws.role"), "global", &HandlerError{Code: "AlreadyExists", Message: "deploy exists", Token: "req-3"})
	if !strings.Contains(exists.Error(), "infrena import") {
		t.Errorf("%q does not suggest importing", exists)
	}
}
