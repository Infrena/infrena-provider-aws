package awsprov

import (
	"context"
	"fmt"
	"maps"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func subnetState(region string, sub types.Subnet) *resource.ResourceState {
	id := aws.ToString(sub.SubnetId)
	attrs := map[string]value.Value{
		"region":                  str(region),
		"vpc_id":                  str(aws.ToString(sub.VpcId)),
		"cidr":                    str(aws.ToString(sub.CidrBlock)),
		"availability_zone":       str(aws.ToString(sub.AvailabilityZone)),
		"map_public_ip_on_launch": boolean(aws.ToBool(sub.MapPublicIpOnLaunch)),
		"id":                      str(id),
		"arn":                     str(aws.ToString(sub.SubnetArn)),
		"owner_id":                str(aws.ToString(sub.OwnerId)),
	}
	putTags(attrs, fromAWSTags(sub.Tags))
	return &resource.ResourceState{Type: typeSubnet, ProviderID: formatID(region, id), Attributes: attrs}
}

func (p *Provider) createSubnet(ctx context.Context, attrs map[string]value.Value) (*resource.ResourceState, error) {
	var in struct{ region, vpcID, cidr, az string }
	for name, dst := range map[string]*string{"region": &in.region, "vpc_id": &in.vpcID, "cidr": &in.cidr, "availability_zone": &in.az} {
		v, err := stringAttr(attrs, name)
		if err != nil {
			return nil, err
		}
		*dst = v
	}
	public, err := boolAttr(attrs, "map_public_ip_on_launch")
	if err != nil {
		return nil, err
	}
	tags, err := tagsFrom(attrs)
	if err != nil {
		return nil, err
	}
	client := p.clients.ec2(in.region)

	var out *ec2.CreateSubnetOutput
	var lastErr error
	// A VPC created a moment ago may not be visible to CreateSubnet yet. A NotFound here is a refusal —
	// nothing was created — so waiting it out is safe in a way retrying an ambiguous failure is not.
	found, err := p.patience.wait(ctx, func() (bool, error) {
		if err := ctx.Err(); err != nil {
			return false, err
		}
		var err error
		out, err = client.CreateSubnet(context.WithoutCancel(ctx), &ec2.CreateSubnetInput{
			VpcId: aws.String(in.vpcID), CidrBlock: aws.String(in.cidr), AvailabilityZone: aws.String(in.az),
			TagSpecifications: tagSpecs(types.ResourceTypeSubnet, tags),
		}, createOnce)
		if hasCode(err, "InvalidVpcID.NotFound") {
			lastErr = err
			return false, nil
		}
		return err == nil, err
	})
	if err != nil {
		return nil, p.failed("CreateSubnet", in.region, "", err)
	}
	if !found {
		return nil, p.failed("CreateSubnet", in.region, "", fmt.Errorf("VPC %s does not exist in %s: %w", in.vpcID, in.region, lastErr))
	}
	if out.Subnet == nil || aws.ToString(out.Subnet.SubnetId) == "" {
		return nil, fmt.Errorf("aws instance %q: CreateSubnet in %s succeeded but returned no subnet ID, so infrata cannot record it: "+
			"look for subnet %s in VPC %s and import it", p.instance, in.region, in.cidr, in.vpcID)
	}
	sub := *out.Subnet
	sub.Tags = toAWSTags(tags)
	sub.MapPublicIpOnLaunch = aws.Bool(false) // what AWS created, until the call below says otherwise

	if public {
		// The subnet exists now. However this call ends, the subnet is reported: an error from a create is
		// dropped by the host, and the subnet would be untracked.
		if err := p.setMapPublicIP(ctx, in.region, aws.ToString(sub.SubnetId), true); err != nil {
			fmt.Fprintf(os.Stderr, "created %s but could not set map_public_ip_on_launch; the next plan will propose it: %v\n",
				formatID(in.region, aws.ToString(sub.SubnetId)), err)
		} else {
			sub.MapPublicIpOnLaunch = aws.Bool(true)
		}
	}
	return subnetState(in.region, sub), nil
}

// setMapPublicIP waits out a subnet ID that has not propagated, as a just-created one may not have.
func (p *Provider) setMapPublicIP(ctx context.Context, region, awsID string, on bool) error {
	client := p.clients.ec2(region)
	var lastErr error
	found, err := p.patience.wait(ctx, func() (bool, error) {
		_, err := client.ModifySubnetAttribute(ctx, &ec2.ModifySubnetAttributeInput{
			SubnetId: aws.String(awsID), MapPublicIpOnLaunch: &types.AttributeBooleanValue{Value: aws.Bool(on)},
		})
		if hasCode(err, "InvalidSubnetID.NotFound") {
			lastErr = err
			return false, nil
		}
		return err == nil, err
	})
	if err == nil && !found {
		err = lastErr
	}
	if err != nil {
		return p.failed("ModifySubnetAttribute", region, awsID, err)
	}
	return nil
}

func (p *Provider) describeSubnet(ctx context.Context, region, awsID string) (*types.Subnet, error) {
	out, err := p.clients.ec2(region).DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{SubnetIds: []string{awsID}})
	if hasCode(err, "InvalidSubnetID.NotFound") {
		return nil, nil
	}
	if err != nil {
		return nil, p.failed("DescribeSubnets", region, awsID, err)
	}
	if len(out.Subnets) == 0 {
		return nil, nil
	}
	return &out.Subnets[0], nil
}

func (p *Provider) readSubnet(ctx context.Context, providerID string, pt patience) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeSubnet, providerID)
	if err != nil {
		return nil, err
	}
	var sub *types.Subnet
	found, err := pt.wait(ctx, func() (bool, error) {
		var err error
		sub, err = p.describeSubnet(ctx, region, awsID)
		return sub != nil, err
	})
	if err != nil || !found {
		return nil, err
	}
	return subnetState(region, *sub), nil
}

// updateSubnet changes what can change in place: map_public_ip_on_launch and tags.
func (p *Provider) updateSubnet(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeSubnet, current.ProviderID)
	if err != nil {
		return nil, err
	}
	attrs := maps.Clone(current.Attributes)
	have, _ := boolAttr(current.Attributes, "map_public_ip_on_launch")
	want, err := boolAttr(desired.Attrs, "map_public_ip_on_launch")
	if err != nil {
		return nil, err
	}
	if have != want {
		if err := p.setMapPublicIP(ctx, region, awsID, want); err != nil {
			return nil, err
		}
		attrs["map_public_ip_on_launch"] = boolean(want)
	}
	if err := p.syncTags(ctx, region, awsID, current.Attributes, desired.Attrs); err != nil {
		return nil, err
	}
	tags, _ := tagsFrom(desired.Attrs)
	putTags(attrs, tags)
	return &resource.ResourceState{Type: typeSubnet, ProviderID: current.ProviderID, Attributes: attrs}, nil
}

func (p *Provider) deleteSubnet(ctx context.Context, current *resource.ResourceState) error {
	region, awsID, err := parseID(typeSubnet, current.ProviderID)
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	_, err = p.clients.ec2(region).DeleteSubnet(context.WithoutCancel(ctx), &ec2.DeleteSubnetInput{SubnetId: aws.String(awsID)})
	if hasCode(err, "InvalidSubnetID.NotFound") {
		return nil
	}
	if err != nil {
		return p.failed("DeleteSubnet", region, awsID, err)
	}
	return nil
}
