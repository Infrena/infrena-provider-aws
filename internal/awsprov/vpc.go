package awsprov

import (
	"context"
	"fmt"
	"maps"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func vpcState(region string, v types.Vpc) *resource.ResourceState {
	id := aws.ToString(v.VpcId)
	attrs := map[string]value.Value{
		"region":     str(region),
		"cidr":       str(aws.ToString(v.CidrBlock)),
		"id":         str(id),
		"owner_id":   str(aws.ToString(v.OwnerId)),
		"is_default": boolean(aws.ToBool(v.IsDefault)),
	}
	putTags(attrs, fromAWSTags(v.Tags))
	return &resource.ResourceState{Type: typeVPC, ProviderID: formatID(region, id), Attributes: attrs}
}

func (p *Provider) createVPC(ctx context.Context, attrs map[string]value.Value) (*resource.ResourceState, error) {
	region, err := stringAttr(attrs, "region")
	if err != nil {
		return nil, err
	}
	cidr, err := stringAttr(attrs, "cidr")
	if err != nil {
		return nil, err
	}
	tags, err := tagsFrom(attrs)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err // nothing sent, nothing lost
	}
	// Once sent, the answer is read whatever happens to ctx: a cancelled request loses the answer, not
	// the VPC.
	out, err := p.clients.ec2(region).CreateVpc(context.WithoutCancel(ctx), &ec2.CreateVpcInput{
		CidrBlock:         aws.String(cidr),
		TagSpecifications: tagSpecs(types.ResourceTypeVpc, tags),
	}, createOnce)
	if err != nil {
		return nil, p.failed("CreateVpc", region, "", err)
	}
	if out.Vpc == nil || aws.ToString(out.Vpc.VpcId) == "" {
		return nil, fmt.Errorf("aws instance %q: CreateVpc in %s succeeded but returned no VPC ID, so infrata cannot record it: "+
			"look for a VPC with CIDR %s in %s and import it", p.instance, region, cidr, region)
	}
	vpc := *out.Vpc
	// From what was sent, not a describe: tags in a create call are applied atomically, and a describe
	// made now may not see the VPC yet.
	vpc.Tags = toAWSTags(tags)
	return vpcState(region, vpc), nil
}

func (p *Provider) describeVPC(ctx context.Context, region, awsID string) (*types.Vpc, error) {
	out, err := p.clients.ec2(region).DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{awsID}})
	if hasCode(err, "InvalidVpcID.NotFound") {
		return nil, nil
	}
	if err != nil {
		return nil, p.failed("DescribeVpcs", region, awsID, err)
	}
	if len(out.Vpcs) == 0 {
		return nil, nil
	}
	return &out.Vpcs[0], nil
}

func (p *Provider) readVPC(ctx context.Context, providerID string, pt patience) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeVPC, providerID)
	if err != nil {
		return nil, err
	}
	var vpc *types.Vpc
	found, err := pt.wait(ctx, func() (bool, error) {
		var err error
		vpc, err = p.describeVPC(ctx, region, awsID)
		return vpc != nil, err
	})
	if err != nil || !found {
		return nil, err
	}
	return vpcState(region, *vpc), nil
}

// updateVPC can only be a tag change: region and cidr are ForceNew, so a plan replaces instead.
func (p *Provider) updateVPC(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeVPC, current.ProviderID)
	if err != nil {
		return nil, err
	}
	if err := p.syncTags(ctx, region, awsID, current.Attributes, desired.Attrs); err != nil {
		return nil, err
	}
	want, _ := tagsFrom(desired.Attrs) // already validated by syncTags
	attrs := maps.Clone(current.Attributes)
	putTags(attrs, want)
	return &resource.ResourceState{Type: typeVPC, ProviderID: current.ProviderID, Attributes: attrs}, nil
}

func (p *Provider) deleteVPC(ctx context.Context, current *resource.ResourceState) error {
	region, awsID, err := parseID(typeVPC, current.ProviderID)
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	_, err = p.clients.ec2(region).DeleteVpc(context.WithoutCancel(ctx), &ec2.DeleteVpcInput{VpcId: aws.String(awsID)})
	if hasCode(err, "InvalidVpcID.NotFound") {
		return nil // already gone is the outcome a delete asks for
	}
	if err != nil {
		return p.failed("DeleteVpc", region, awsID, err)
	}
	return nil
}
