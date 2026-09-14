package ec2fake

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
)

// client is a real SDK client pointed at the fake. If the SDK can decode what the fake says, the
// fake's wire format is right: the SDK is the oracle, not this package's own reading of it.
func client(t *testing.T, s *Server, region string) *ec2.Client {
	t.Helper()
	return ec2.New(ec2.Options{
		Region:       region,
		BaseEndpoint: aws.String(s.URL),
		Credentials:  credentials.NewStaticCredentialsProvider("AKIDORACLE", "secret", ""),
	})
}

func TestTheSDKDecodesAVPCRoundTrip(t *testing.T) {
	s := New()
	defer s.Close()
	ctx := context.Background()
	c := client(t, s, "us-east-1")

	out, err := c.CreateVpc(ctx, &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
		TagSpecifications: []types.TagSpecification{{
			ResourceType: types.ResourceTypeVpc,
			Tags:         []types.Tag{{Key: aws.String("team"), Value: aws.String("platform")}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	id := aws.ToString(out.Vpc.VpcId)
	if id == "" || aws.ToString(out.Vpc.CidrBlock) != "10.0.0.0/16" || len(out.Vpc.Tags) != 1 {
		t.Fatalf("CreateVpc decoded as %+v", out.Vpc)
	}
	got, err := c.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{id}})
	if err != nil || len(got.Vpcs) != 1 || aws.ToString(got.Vpcs[0].OwnerId) != Owner {
		t.Fatalf("DescribeVpcs = %+v, %v", got, err)
	}
	if _, err := c.DeleteTags(ctx, &ec2.DeleteTagsInput{Resources: []string{id}, Tags: []types.Tag{{Key: aws.String("team")}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.DeleteVpc(ctx, &ec2.DeleteVpcInput{VpcId: aws.String(id)}); err != nil {
		t.Fatal(err)
	}
	_, err = c.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{id}})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "InvalidVpcID.NotFound" {
		t.Fatalf("describing a deleted VPC = %v, want InvalidVpcID.NotFound", err)
	}
	if keys := s.AccessKeys(); len(keys) == 0 || keys[0] != "AKIDORACLE" {
		t.Errorf("access keys seen = %v", keys)
	}
}

func TestTheSDKDecodesASubnetAndItsAttribute(t *testing.T) {
	s := New()
	defer s.Close()
	ctx := context.Background()
	c := client(t, s, "eu-west-1")
	vpc := s.AddVPC("eu-west-1", "10.1.0.0/16", nil)

	out, err := c.CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId: aws.String(vpc), CidrBlock: aws.String("10.1.1.0/24"), AvailabilityZone: aws.String("eu-west-1a"),
	})
	if err != nil {
		t.Fatal(err)
	}
	id := aws.ToString(out.Subnet.SubnetId)
	if _, err := c.ModifySubnetAttribute(ctx, &ec2.ModifySubnetAttributeInput{
		SubnetId: aws.String(id), MapPublicIpOnLaunch: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	}); err != nil {
		t.Fatal(err)
	}
	got, err := c.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{SubnetIds: []string{id}})
	if err != nil || len(got.Subnets) != 1 {
		t.Fatalf("DescribeSubnets = %+v, %v", got, err)
	}
	sub := got.Subnets[0]
	if !aws.ToBool(sub.MapPublicIpOnLaunch) || aws.ToString(sub.AvailabilityZone) != "eu-west-1a" ||
		aws.ToString(sub.SubnetArn) != "arn:aws:ec2:eu-west-1:"+Owner+":subnet/"+id || aws.ToString(sub.VpcId) != vpc {
		t.Fatalf("subnet decoded as %+v", sub)
	}
	// A subnet in a VPC the region does not hold is refused the way EC2 refuses it.
	_, err = client(t, s, "us-east-1").CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId: aws.String(vpc), CidrBlock: aws.String("10.1.2.0/24"), AvailabilityZone: aws.String("us-east-1a"),
	})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "InvalidVpcID.NotFound" {
		t.Fatalf("cross-region CreateSubnet = %v, want InvalidVpcID.NotFound", err)
	}
}

// TestRegionsArePartitionedAndPagesFollowNextToken. The fixture holds three VPCs in one region and
// one in another, with a page size of one, so neither a region leak nor a paginator that stops after
// the first page can pass.
func TestRegionsArePartitionedAndPagesFollowNextToken(t *testing.T) {
	s := New()
	defer s.Close()
	s.PageSize = 1
	for _, cidr := range []string{"10.0.0.0/16", "10.1.0.0/16", "10.2.0.0/16"} {
		s.AddVPC("us-east-1", cidr, nil)
	}
	s.AddVPC("eu-west-1", "10.9.0.0/16", nil)

	var n int
	pg := ec2.NewDescribeVpcsPaginator(client(t, s, "us-east-1"), &ec2.DescribeVpcsInput{})
	for pg.HasMorePages() {
		page, err := pg.NextPage(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		n += len(page.Vpcs)
	}
	if n != 3 {
		t.Fatalf("paginated us-east-1 VPCs = %d, want 3", n)
	}
}

func TestAnInjectedFaultDecodesWithItsStatus(t *testing.T) {
	s := New()
	defer s.Close()
	s.Inject(Fault{Action: "DescribeVpcs", Nth: 1, Status: 400, Code: "UnauthorizedOperation", Message: "no"})
	_, err := client(t, s, "us-east-1").DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "UnauthorizedOperation" {
		t.Fatalf("err = %v, want UnauthorizedOperation", err)
	}
	if s.Calls("DescribeVpcs") != 1 {
		t.Errorf("calls = %d, want 1 (a 400 is not retried by the SDK)", s.Calls("DescribeVpcs"))
	}
}

func TestHiddenResourcesAreNotFoundForAWhile(t *testing.T) {
	s := New()
	defer s.Close()
	id := s.AddVPC("us-east-1", "10.0.0.0/16", nil)
	s.HideFromDescribe(id, 2)
	c := client(t, s, "us-east-1")
	for i := 1; i <= 3; i++ {
		_, err := c.DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{VpcIds: []string{id}})
		if hidden := err != nil; hidden != (i <= 2) {
			t.Fatalf("describe %d: err = %v", i, err)
		}
	}
}
