package ccprov

import (
	"context"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

// owner decides whether AWS itself owns one discovered resource, from the properties Cloud Control returned and what
// the region can be asked about it.
//
// It is a CLAIM, never a guess. A signal that merely suggests AWS made something — a VPC whose CIDR is 172.31.0.0/16,
// a name that looks generated — is not evidence, and a resource a user built by hand must never be flagged because it
// resembles one AWS built. Where the only authoritative answer lives in another API, it is asked (see defaults.go);
// where it cannot be had, the flag stays unset.
type owner func(ctx context.Context, props map[string]any, facts regionFacts) (bool, string)

// owners is the evidence available per CloudFormation type.
var owners = map[string]owner{
	"AWS::EC2::VPC":    defaultVPC,
	"AWS::EC2::Subnet": defaultSubnet,
}

// systemOwned is what Discover reports to infrena, which prints the reason and leaves the resource out of `import`
// unless a selector names it. The reason is read by a person deciding whether to adopt the resource, so it names the
// evidence rather than restating the conclusion.
func systemOwned(ctx context.Context, t *catalog.Type, props map[string]any, facts regionFacts) (bool, string) {
	if own, ok := owners[t.CFN]; ok {
		return own(ctx, props, facts)
	}
	return false, ""
}

// defaultVPC: EC2 says which VPC is the region's default. Cloud Control's AWS::EC2::VPC schema has no IsDefault.
func defaultVPC(ctx context.Context, props map[string]any, facts regionFacts) (bool, string) {
	id, known := facts.defaultVPC(ctx)
	if !known || id == "" {
		return false, ""
	}
	if vpcID, _ := props["VpcId"].(string); vpcID == id {
		return true, "the default VPC for this region, reported by EC2"
	}
	return false, ""
}

// defaultSubnet: EC2 says which subnets are default for their availability zone. Cloud Control's AWS::EC2::Subnet
// schema has no DefaultForAz.
func defaultSubnet(ctx context.Context, props map[string]any, facts regionFacts) (bool, string) {
	ids, known := facts.defaultSubnets(ctx)
	if !known {
		return false, ""
	}
	if subnetID, _ := props["SubnetId"].(string); ids[subnetID] {
		return true, "the default subnet for its availability zone, reported by EC2"
	}
	return false, ""
}
