package ccfake

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// client is the real SDK pointed at the fake: if the SDK decodes what the fake says, the wire format is right.
func client(s *Server, region string) *cloudcontrol.Client {
	return cloudcontrol.New(cloudcontrol.Options{
		Region: region, BaseEndpoint: aws.String(s.URL),
		Credentials: credentials.NewStaticCredentialsProvider("AKIDORACLE", "secret", ""),
	})
}

func vpcType() TypeConfig {
	return TypeConfig{
		TypeName: "AWS::EC2::VPC", Identifier: "VpcId", IDPrefix: "vpc-",
		ReadOnly:   map[string]string{"DefaultSecurityGroup": "sg-for-{id}"},
		Defaults:   map[string]any{"EnableDnsSupport": true, "InstanceTenancy": "default"},
		CreateOnly: []string{"CidrBlock"},
		WriteOnly:  []string{"Ipv4NetmaskLength"},
	}
}

func await(t *testing.T, c *cloudcontrol.Client, ev *types.ProgressEvent) *types.ProgressEvent {
	t.Helper()
	for i := 0; i < 20 && (ev.OperationStatus == types.OperationStatusPending || ev.OperationStatus == types.OperationStatusInProgress); i++ {
		out, err := c.GetResourceRequestStatus(context.Background(), &cloudcontrol.GetResourceRequestStatusInput{RequestToken: ev.RequestToken})
		if err != nil {
			t.Fatal(err)
		}
		ev = out.ProgressEvent
	}
	return ev
}

func TestTheSDKDrivesAFullLifecycle(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "eu-west-1")
	ctx := context.Background()

	created, err := c.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), ClientToken: aws.String("tok-1"),
		DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16","Ipv4NetmaskLength":16,"Tags":[{"Key":"team","Value":"a"}]}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.ProgressEvent.OperationStatus != types.OperationStatusInProgress || created.ProgressEvent.EventTime == nil {
		t.Fatalf("create event = %+v", created.ProgressEvent)
	}
	ev := await(t, c, created.ProgressEvent)
	id := aws.ToString(ev.Identifier)
	if ev.OperationStatus != types.OperationStatusSuccess || !strings.HasPrefix(id, "vpc-") {
		t.Fatalf("settled event = %+v", ev)
	}

	got, err := c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	if err != nil {
		t.Fatal(err)
	}
	var props map[string]any
	if err := json.Unmarshal([]byte(aws.ToString(got.ResourceDescription.Properties)), &props); err != nil {
		t.Fatal(err)
	}
	if props["VpcId"] != id || props["EnableDnsSupport"] != true || props["DefaultSecurityGroup"] != "sg-for-"+id {
		t.Errorf("properties = %v: want the identifier, a provider-chosen default and a read-only value", props)
	}
	if _, leaked := props["Ipv4NetmaskLength"]; leaked {
		t.Error("a write-only property was returned")
	}

	upd, err := c.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id), ClientToken: aws.String("tok-2"),
		PatchDocument: aws.String(`[{"op":"replace","path":"/EnableDnsSupport","value":false}]`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if ev := await(t, c, upd.ProgressEvent); ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("update = %+v", ev)
	}
	if res, _ := s.Resource("eu-west-1", "AWS::EC2::VPC", id); res["EnableDnsSupport"] != false {
		t.Errorf("after patch: %v", res)
	}

	_, err = c.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id),
		PatchDocument: aws.String(`[{"op":"replace","path":"/CidrBlock","value":"10.1.0.0/16"}]`),
	})
	var notUpdatable *types.NotUpdatableException
	if !errors.As(err, &notUpdatable) {
		t.Errorf("patching a create-only property = %v, want NotUpdatableException", err)
	}

	del, err := c.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	if err != nil {
		t.Fatal(err)
	}
	if ev := await(t, c, del.ProgressEvent); ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("delete = %+v", ev)
	}
	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	var nf *types.ResourceNotFoundException
	if !errors.As(err, &nf) {
		t.Fatalf("get after delete = %v", err)
	}
	if s.Calls("CreateResource") != 1 || len(s.Tokens("CreateResource")) != 1 {
		t.Errorf("calls/tokens not recorded")
	}
}

func TestAClientTokenMakesACreateIdempotent(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "us-east-1")
	in := &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), ClientToken: aws.String("same"), DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16"}`)}
	a, err := c.CreateResource(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	b, err := c.CreateResource(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	if aws.ToString(a.ProgressEvent.RequestToken) != aws.ToString(b.ProgressEvent.RequestToken) || len(s.Resources("us-east-1", "AWS::EC2::VPC")) != 1 {
		t.Fatal("the same client token made a second resource")
	}
}

// TestRegionsArePartitionedAndListPages: three in one region, one in another, page size one.
func TestRegionsArePartitionedAndListPages(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	s.PageSize = 1
	for _, id := range []string{"vpc-1", "vpc-2", "vpc-3"} {
		s.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
	}
	s.Put("eu-west-1", "AWS::EC2::VPC", "vpc-9", map[string]any{"VpcId": "vpc-9"})
	var ids []string
	pg := cloudcontrol.NewListResourcesPaginator(client(s, "us-east-1"), &cloudcontrol.ListResourcesInput{TypeName: aws.String("AWS::EC2::VPC")})
	for pg.HasMorePages() {
		page, err := pg.NextPage(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range page.ResourceDescriptions {
			ids = append(ids, aws.ToString(d.Identifier))
		}
	}
	if strings.Join(ids, ",") != "vpc-1,vpc-2,vpc-3" {
		t.Fatalf("listed %v", ids)
	}
}

func TestFaultsAndFailedRequests(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "us-east-1")
	ctx := context.Background()

	s.Inject(Fault{Action: "CreateResource", Nth: 1, Status: 429, Code: "ThrottlingException", Message: "slow down"})
	_, err := cloudcontrol.New(cloudcontrol.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(s.URL), RetryMaxAttempts: 1,
		Credentials: credentials.NewStaticCredentialsProvider("AKID", "secret", ""),
	}).CreateResource(ctx, &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`)})
	var throttled *types.ThrottlingException
	if !errors.As(err, &throttled) {
		t.Fatalf("injected throttle = %v", err)
	}

	s.FailNext("CREATE", "NotStabilized", "not stable yet", true)
	out, err := c.CreateResource(ctx, &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16"}`)})
	if err != nil {
		t.Fatal(err)
	}
	ev := await(t, c, out.ProgressEvent)
	if ev.OperationStatus != types.OperationStatusFailed || ev.ErrorCode != types.HandlerErrorCodeNotStabilized || aws.ToString(ev.Identifier) == "" {
		t.Fatalf("failed create = %+v, want FAILED NotStabilized with an identifier", ev)
	}
	if _, exists := s.Resource("us-east-1", "AWS::EC2::VPC", aws.ToString(ev.Identifier)); !exists {
		t.Error("keepResource: the resource should still exist")
	}

	s.HideFromGet(aws.ToString(ev.Identifier), 1)
	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: ev.Identifier})
	var nf *types.ResourceNotFoundException
	if !errors.As(err, &nf) {
		t.Errorf("hidden get = %v", err)
	}

	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::Nope::Thing"), Identifier: aws.String("x")})
	var tnf *types.TypeNotFoundException
	if !errors.As(err, &tnf) {
		t.Errorf("unregistered type = %v", err)
	}
}

func TestAssumeRoleAnswersAndSignsAreRecorded(t *testing.T) {
	s := New()
	defer s.Close()
	out, err := sts.New(sts.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(s.URL),
		Credentials: credentials.NewStaticCredentialsProvider("AKIDSTATIC", "secret", ""),
	}).AssumeRole(context.Background(), &sts.AssumeRoleInput{RoleArn: aws.String("arn:aws:iam::123456789012:role/x"), RoleSessionName: aws.String("t")})
	if err != nil {
		t.Fatal(err)
	}
	if aws.ToString(out.Credentials.AccessKeyId) != AssumedAccessKey || !out.Credentials.Expiration.After(time.Now()) {
		t.Fatalf("credentials = %+v", out.Credentials)
	}
	if keys := s.AccessKeys(); len(keys) != 1 || keys[0] != "AKIDSTATIC" {
		t.Errorf("access keys = %v", keys)
	}
}
