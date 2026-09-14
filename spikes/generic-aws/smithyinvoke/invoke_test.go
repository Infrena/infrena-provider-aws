package smithyinvoke

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/smithy-go"
	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
)

// The EC2 model is 8 MB, so it is not committed: scripts/fetch-models puts it in testdata/.
const modelPath = "testdata/ec2.json"

func loadEC2(t *testing.T) *Model {
	t.Helper()
	if _, err := os.Stat(modelPath); err != nil {
		t.Skipf("no %s (run spikes/generic-aws/scripts/fetch-models): %v", modelPath, err)
	}
	m, err := Load(modelPath)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func target(fake *ec2fake.Server) Target {
	return Target{Endpoint: fake.URL, Region: "us-east-1", SigningName: "ec2",
		Credentials: credentials.NewStaticCredentialsProvider("AKIDSMITHY", "secret", "")}
}

// TestGenericInvocationOfCreateDescribeDelete. CreateVpc, DescribeVpcs and DeleteVpc are strings here; no
// CreateVpcInput exists anywhere in this package or its imports.
func TestGenericInvocationOfCreateDescribeDelete(t *testing.T) {
	m := loadEC2(t)
	t.Logf("model %s declares %d operations; protocols %v", m.Version(), len(m.Operations()), m.Protocols())
	fake := ec2fake.New()
	defer fake.Close()
	ctx := context.Background()

	created, err := Invoke(ctx, m, target(fake), "CreateVpc", map[string]any{
		"CidrBlock": "10.0.0.0/16",
		"TagSpecifications": []any{map[string]any{
			"ResourceType": "vpc",
			"Tags":         []any{map[string]any{"Key": "team", "Value": "platform"}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("CreateVpc -> %v", created)
	vpc, _ := created["Vpc"].(map[string]any)
	id, _ := vpc["VpcId"].(string)
	if id == "" || vpc["CidrBlock"] != "10.0.0.0/16" {
		t.Fatalf("CreateVpc output = %v", created)
	}
	if got := fake.VPCs(); len(got) != 1 || got[0].Tags["team"] != "platform" {
		t.Fatalf("the fake received %+v — the nested TagSpecifications did not serialize", got)
	}

	described, err := Invoke(ctx, m, target(fake), "DescribeVpcs", map[string]any{"VpcIds": []any{id}})
	if err != nil {
		t.Fatal(err)
	}
	vpcs, _ := described["Vpcs"].([]any)
	if len(vpcs) != 1 {
		t.Fatalf("DescribeVpcs = %v", described)
	}
	tags, _ := vpcs[0].(map[string]any)["Tags"].([]any)
	if len(tags) != 1 {
		t.Errorf("tags did not come back: %v", vpcs[0])
	}

	if _, err := Invoke(ctx, m, target(fake), "DeleteVpc", map[string]any{"VpcId": id}); err != nil {
		t.Fatal(err)
	}
	_, err = Invoke(ctx, m, target(fake), "DescribeVpcs", map[string]any{"VpcIds": []any{id}})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "InvalidVpcID.NotFound" {
		t.Fatalf("describe after delete = %v (%T), want InvalidVpcID.NotFound", err, err)
	}
}

// TestAnOperationTheSpikeNeverMentioned. Subnets and a nested attribute structure, to show the mechanism
// is not VPC-shaped.
func TestAnOperationTheSpikeNeverMentioned(t *testing.T) {
	m := loadEC2(t)
	fake := ec2fake.New()
	defer fake.Close()
	ctx := context.Background()
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)

	out, err := Invoke(ctx, m, target(fake), "CreateSubnet", map[string]any{
		"VpcId": vpc, "CidrBlock": "10.0.1.0/24", "AvailabilityZone": "us-east-1a",
	})
	if err != nil {
		t.Fatal(err)
	}
	subnet := out["Subnet"].(map[string]any)
	if _, err := Invoke(ctx, m, target(fake), "ModifySubnetAttribute", map[string]any{
		"SubnetId":            subnet["SubnetId"],
		"MapPublicIpOnLaunch": map[string]any{"Value": true},
	}); err != nil {
		t.Fatal(err)
	}
	if !fake.Subnets()[0].MapPublicIP {
		t.Fatal("the nested AttributeBooleanValue did not serialize")
	}
}

func TestAMisspeltMemberIsRefusedBeforeAnyRequest(t *testing.T) {
	m := loadEC2(t)
	fake := ec2fake.New()
	defer fake.Close()
	_, err := Invoke(context.Background(), m, target(fake), "CreateVpc", map[string]any{"CidrBlok": "10.0.0.0/16"})
	if err == nil || !strings.Contains(err.Error(), `"CidrBlok"`) || !strings.Contains(err.Error(), "CidrBlock") {
		t.Fatalf("err = %v", err)
	}
	if n := fake.Calls("CreateVpc"); n != 0 {
		t.Fatalf("a request was sent (%d)", n)
	}
}
