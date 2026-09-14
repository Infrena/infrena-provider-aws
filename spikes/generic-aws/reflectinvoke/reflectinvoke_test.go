package reflectinvoke

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
)

func TestReflectionInvokesEC2ByName(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	client := ec2.New(ec2.Options{Region: "us-east-1", BaseEndpoint: aws.String(fake.URL),
		Credentials: credentials.NewStaticCredentialsProvider("AKID", "secret", "")})
	ctx := context.Background()

	created, err := Invoke(ctx, client, "CreateVpc", map[string]any{
		"CidrBlock": "10.0.0.0/16",
		"TagSpecifications": []any{map[string]any{
			"ResourceType": "vpc",
			"Tags":         []any{map[string]any{"Key": "team", "Value": "platform"}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	vpc := created["Vpc"].(map[string]any)
	id := vpc["VpcId"].(string)
	t.Logf("CreateVpc output: %v", created)

	described, err := Invoke(ctx, client, "DescribeVpcs", map[string]any{"VpcIds": []any{id}})
	if err != nil {
		t.Fatal(err)
	}
	if n := len(described["Vpcs"].([]any)); n != 1 {
		t.Fatalf("described %d", n)
	}
	if _, err := Invoke(ctx, client, "DeleteVpc", map[string]any{"VpcId": id}); err != nil {
		t.Fatal(err)
	}
	if len(fake.VPCs()) != 0 {
		t.Fatal("not deleted")
	}

	// The hazard: a misspelt member is silently dropped, and the call still goes out.
	_, err = Invoke(ctx, client, "CreateVpc", map[string]any{"CidrBlok": "10.1.0.0/16"})
	t.Logf("misspelt member: err=%v, VPCs now=%d (the typo was dropped without complaint)", err, len(fake.VPCs()))
}
