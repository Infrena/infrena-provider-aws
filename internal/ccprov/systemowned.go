package ccprov

import (
	"context"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

// cloudFormationTag is the tag CloudFormation puts on every resource a stack owns. AWS reserves the aws: prefix, so
// nothing but AWS can have written it.
const cloudFormationTag = "aws:cloudformation:stack-id"

// serviceRolePath is the path IAM gives a service-linked role, which only the service that owns it may create or
// delete.
const serviceRolePath = "/aws-service-role/"

// launcherTags are the tags AWS puts on an instance some other controller launched and will replace on its own
// schedule: an Auto Scaling group, an EC2 Fleet, a Spot Fleet request. All carry the aws: prefix AWS reserves, so
// nothing but AWS can have written them — the same standard cloudFormationTag is held to.
//
// A launch template is deliberately NOT here. aws:ec2launchtemplate:id says an instance was launched FROM a
// template, which a user does by hand all the time; it names no controller that will replace it.
var launcherTags = []string{
	"aws:autoscaling:groupName",
	"aws:ec2:fleet-id",
	"aws:ec2spot:fleet-request-id",
}

// owner decides whether AWS itself owns one discovered resource, from the properties Cloud Control returned and what
// the region can be asked about it.
//
// It is a CLAIM, never a guess. A signal that merely suggests AWS made something — a VPC whose CIDR is 172.31.0.0/16,
// a name that looks generated — is not evidence, and a resource a user built by hand must never be flagged because it
// resembles one AWS built. Where the only authoritative answer lives in another API, it is asked (see defaults.go);
// where it cannot be had, the flag stays unset.
type owner func(ctx context.Context, t *catalog.Type, props map[string]any, facts regionFacts) (bool, string)

// owners is the evidence available per CloudFormation type. A type that is not here can still be flagged by its tags.
var owners = map[string]owner{
	"AWS::EC2::VPC":           defaultVPC,
	"AWS::EC2::Subnet":        defaultSubnet,
	"AWS::EC2::SecurityGroup": defaultSecurityGroup,
	"AWS::EC2::Instance":      fleetLaunched,
	"AWS::IAM::Role":          serviceLinkedRole,
}

// systemOwned is what Discover reports to infrena, which prints the reason and leaves the resource out of `import`
// unless a selector names it. The reason is read by a person deciding whether to adopt the resource, so it names the
// evidence rather than restating the conclusion.
//
// The type's own evidence is preferred to the CloudFormation tag: both are true of a resource a stack made, and the
// specific claim is the more useful one to read.
func systemOwned(ctx context.Context, t *catalog.Type, props map[string]any, facts regionFacts) (bool, string) {
	if own, ok := owners[t.CFN]; ok {
		if owned, why := own(ctx, t, props, facts); owned {
			return true, why
		}
	}
	return stackOwned(t, props)
}

// defaultVPC: EC2 says which VPC is the region's default. Cloud Control's AWS::EC2::VPC schema has no IsDefault.
func defaultVPC(ctx context.Context, _ *catalog.Type, props map[string]any, facts regionFacts) (bool, string) {
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
func defaultSubnet(ctx context.Context, _ *catalog.Type, props map[string]any, facts regionFacts) (bool, string) {
	ids, known := facts.defaultSubnets(ctx)
	if !known {
		return false, ""
	}
	if subnetID, _ := props["SubnetId"].(string); ids[subnetID] {
		return true, "the default subnet for its availability zone, reported by EC2"
	}
	return false, ""
}

// defaultSecurityGroup: every VPC gets one, it cannot be deleted, and no other security group may be named `default`.
func defaultSecurityGroup(_ context.Context, _ *catalog.Type, props map[string]any, _ regionFacts) (bool, string) {
	if name, _ := props["GroupName"].(string); name == "default" {
		return true, "GroupName is default"
	}
	return false, ""
}

// fleetLaunched: an Auto Scaling group, an EC2 Fleet or a Spot Fleet request owns this instance's lifecycle and
// will replace it whenever it decides to — on a spot fleet draining short jobs, that is minutes.
//
// Adopting one produces a state record that is stale almost immediately, and then a plan that proposes CREATING an
// instance ID which no longer exists anywhere, on every run, for ever. It reads as an infrena bug and is not one: the
// resource really is gone, and infrena really did record it. The fix is to never take it on, because managing a
// lifecycle someone else controls is not a thing infrena can do — only a thing it can appear to do until the
// controller next acts.
//
// The tags are read from the raw properties for the reason stackOwned gives: decoding drops every aws:-prefixed tag
// (tags.go), so this is the only place they can be seen.
func fleetLaunched(_ context.Context, t *catalog.Type, props map[string]any, _ regionFacts) (bool, string) {
	for _, tag := range launcherTags {
		if hasTag(tagsOf(t, props), tag) {
			return true, "tagged " + tag
		}
	}
	return false, ""
}

// serviceLinkedRole: IAM reserves this path for roles a service creates for itself, and refuses it to anyone else.
func serviceLinkedRole(_ context.Context, _ *catalog.Type, props map[string]any, _ regionFacts) (bool, string) {
	if path, _ := props["Path"].(string); strings.HasPrefix(path, serviceRolePath) {
		return true, "Path starts " + serviceRolePath
	}
	return false, ""
}

// stackOwned reads the tags AWS returned, not the ones the plugin reports: decoding drops every aws:-prefixed tag
// (tags.go), and this is one of them, so the raw properties are the only place it can be seen.
func stackOwned(t *catalog.Type, props map[string]any) (bool, string) {
	if hasTag(tagsOf(t, props), cloudFormationTag) {
		return true, "tagged " + cloudFormationTag
	}
	return false, ""
}

// tagsOf returns the raw tags Cloud Control sent, from whichever property this type keeps them under.
func tagsOf(t *catalog.Type, props map[string]any) any {
	name := t.TagsAsMap
	if name == "" {
		name = "Tags"
	}
	return props[name]
}

// hasTag looks for a tag key in either shape AWS returns: a list of {Key, Value} pairs, or a map.
func hasTag(datum any, key string) bool {
	switch tags := datum.(type) {
	case []any:
		for _, item := range tags {
			tag, _ := item.(map[string]any)
			if k, _ := tag["Key"].(string); k == key {
				return true
			}
		}
	case map[string]any:
		_, ok := tags[key]
		return ok
	}
	return false
}
