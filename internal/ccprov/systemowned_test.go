package ccprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/provider"
)

// stubFacts is a region whose defaults are already settled: known ones, or, with the known flags false, a region EC2
// could not be asked about.
type stubFacts struct {
	vpc          string
	vpcKnown     bool
	subnets      map[string]bool
	subnetsKnown bool
}

func (f stubFacts) defaultVPC(context.Context) (string, bool) { return f.vpc, f.vpcKnown }

func (f stubFacts) defaultSubnets(context.Context) (map[string]bool, bool) {
	return f.subnets, f.subnetsKnown
}

func tagList(kv ...string) []any {
	var out []any
	for i := 0; i+1 < len(kv); i += 2 {
		out = append(out, map[string]any{"Key": kv[i], "Value": kv[i+1]})
	}
	return out
}

// TestTheFlagIsSetOnlyOnEvidence. Every row that is NOT flagged is a resource a user may have built by hand, and
// flagging one would drop it from `import` silently. The CIDR AWS gives a default VPC and a name that reads like a
// default are in here precisely because they are not evidence.
func TestTheFlagIsSetOnlyOnEvidence(t *testing.T) {
	known := stubFacts{
		vpc: "vpc-default", vpcKnown: true,
		subnets: map[string]bool{"subnet-default": true}, subnetsKnown: true,
	}
	blind := stubFacts{} // EC2 could not be asked

	for _, c := range []struct {
		name   string
		typ    string
		props  map[string]any
		facts  regionFacts
		reason string // "" means the flag must stay unset
	}{
		{"the default VPC, as EC2 reports it", "aws.vpc",
			map[string]any{"VpcId": "vpc-default", "CidrBlock": "172.31.0.0/16"}, known,
			"the default VPC for this region, reported by EC2"},
		{"a VPC on the CIDR AWS gives default VPCs", "aws.vpc",
			map[string]any{"VpcId": "vpc-mine", "CidrBlock": "172.31.0.0/16"}, known, ""},
		{"a VPC named like a default", "aws.vpc",
			map[string]any{"VpcId": "vpc-mine", "Tags": tagList("Name", "default")}, known, ""},
		{"the default VPC when EC2 could not be asked", "aws.vpc",
			map[string]any{"VpcId": "vpc-default"}, blind, ""},

		{"a default subnet, as EC2 reports it", "aws.subnet",
			map[string]any{"SubnetId": "subnet-default"}, known,
			"the default subnet for its availability zone, reported by EC2"},
		{"a subnet in the default VPC that is not a default subnet", "aws.subnet",
			map[string]any{"SubnetId": "subnet-mine", "VpcId": "vpc-default"}, known, ""},
		{"a default subnet when EC2 could not be asked", "aws.subnet",
			map[string]any{"SubnetId": "subnet-default"}, blind, ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			owned, reason := systemOwned(ctx, mustType(t, c.typ), c.props, c.facts)
			if owned != (c.reason != "") || reason != c.reason {
				t.Errorf("systemOwned = %t, %q; want %t, %q", owned, reason, c.reason != "", c.reason)
			}
		})
	}
}

// byID indexes a discovery result.
func byID(found []provider.DiscoveredResource) map[string]provider.DiscoveredResource {
	out := map[string]provider.DiscoveredResource{}
	for _, r := range found {
		out[r.ProviderID] = r
	}
	return out
}

func wantOwned(t *testing.T, found []provider.DiscoveredResource, id, reason string) {
	t.Helper()
	r, ok := byID(found)[id]
	switch {
	case !ok:
		t.Fatalf("%s was not discovered at all; found %v", id, ids(found))
	case !r.SystemOwned:
		t.Errorf("%s is not declared cloud-owned, so `import` would adopt it by default", id)
	case r.SystemOwnedReason != reason:
		t.Errorf("%s reads %q, want %q", id, r.SystemOwnedReason, reason)
	}
}

func wantNotOwned(t *testing.T, found []provider.DiscoveredResource, id string) {
	t.Helper()
	r, ok := byID(found)[id]
	switch {
	case !ok:
		t.Fatalf("%s was not discovered at all; found %v", id, ids(found))
	case r.SystemOwned:
		t.Errorf("%s is declared cloud-owned (%q), so `import` would skip a resource a user made", id, r.SystemOwnedReason)
	case r.SystemOwnedReason != "":
		t.Errorf("%s carries the reason %q with the flag unset", id, r.SystemOwnedReason)
	}
}

func TestDiscoveryDeclaresTheDefaultVPCFromWhatEC2Says(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	// Both VPCs have the CIDR AWS gives a default one: EC2's answer is the only thing that separates them.
	for _, id := range []string{"vpc-default", "vpc-mine"} {
		fake.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id, "CidrBlock": "172.31.0.0/16"})
	}
	fake.SetDefaultVPC("us-east-1", "vpc-default")

	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil {
		t.Fatal(err)
	}
	wantOwned(t, found, "us-east-1/vpc-default", "the default VPC for this region, reported by EC2")
	wantNotOwned(t, found, "us-east-1/vpc-mine")
	if got := fake.Filters("DescribeVpcs"); len(got) != 1 || got[0] != "is-default=true" {
		t.Errorf("DescribeVpcs filters = %v, want one call filtered to is-default=true", got)
	}
}

func TestDiscoveryDeclaresDefaultSubnetsFromWhatEC2Says(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	for _, id := range []string{"subnet-default", "subnet-mine"} {
		fake.Put("us-east-1", "AWS::EC2::Subnet", id, map[string]any{"SubnetId": id, "VpcId": "vpc-default", "AvailabilityZone": "us-east-1a"})
	}
	fake.SetDefaultSubnet("us-east-1", "subnet-default")

	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.subnet"}})
	if err != nil {
		t.Fatal(err)
	}
	wantOwned(t, found, "us-east-1/subnet-default", "the default subnet for its availability zone, reported by EC2")
	wantNotOwned(t, found, "us-east-1/subnet-mine")
	if got := fake.Filters("DescribeSubnets"); len(got) != 1 || got[0] != "default-for-az=true" {
		t.Errorf("DescribeSubnets filters = %v, want one call filtered to default-for-az=true", got)
	}
}

// TestEC2IsAskedOncePerRegionPerRun. Discovery reads every resource; asking EC2 per resource would turn a run over a
// large account into thousands of extra calls.
func TestEC2IsAskedOncePerRegionPerRun(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	p.opts.DiscoverRegions = []string{"us-east-1", "eu-west-1"}
	for _, region := range p.opts.DiscoverRegions {
		for _, id := range []string{region + "-vpc-1", region + "-vpc-2", region + "-vpc-3"} {
			fake.Put(region, "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
		}
		fake.SetDefaultVPC(region, region+"-vpc-1")
	}
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 6 {
		t.Fatalf("discovered %v", ids(found))
	}
	wantOwned(t, found, "eu-west-1/eu-west-1-vpc-1", "the default VPC for this region, reported by EC2")
	if n := fake.Calls("DescribeVpcs"); n != 2 {
		t.Errorf("DescribeVpcs calls = %d, want 2: once per region, not once per resource", n)
	}
	if n := fake.Calls("DescribeSubnets"); n != 0 {
		t.Errorf("DescribeSubnets calls = %d: no subnet was being discovered", n)
	}
}

// TestAFailingEC2CallDoesNotFailDiscovery. The flag is advisory; the discovery it hangs off is not. Losing the flag
// must leave the resources, unmarked, and say so where a user can see it.
func TestAFailingEC2CallDoesNotFailDiscovery(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-default", map[string]any{"VpcId": "vpc-default"})
	fake.SetDefaultVPC("us-east-1", "vpc-default")
	fake.Inject(ccfake.Fault{Action: "DescribeVpcs", Nth: 1, Status: 403, Code: "UnauthorizedOperation",
		Message: "You are not authorized to perform this operation"})

	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil {
		t.Fatalf("a failed EC2 call failed the whole discovery: %v", err)
	}
	if len(found) != 1 {
		t.Fatalf("discovered %v, want the VPC", ids(found))
	}
	wantNotOwned(t, found, "us-east-1/vpc-default")
	for _, want := range []string{"UnauthorizedOperation", "us-east-1", "none are marked as cloud-owned", "ec2:DescribeVpcs"} {
		if !strings.Contains(log.String(), want) {
			t.Errorf("stderr never says %q:\n%s", want, log)
		}
	}
}

// TestEC2IsNotAskedAgainAfterItFails: the failure is cached like an answer, so one broken permission cannot cost a
// call per resource.
func TestEC2IsNotAskedAgainAfterItFails(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	for _, id := range []string{"vpc-1", "vpc-2", "vpc-3"} {
		fake.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
	}
	fake.Inject(ccfake.Fault{Action: "DescribeVpcs", Nth: 1, Status: 403, Code: "UnauthorizedOperation", Message: "denied"})

	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil || len(found) != 3 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
	if n := fake.Calls("DescribeVpcs"); n != 1 {
		t.Errorf("DescribeVpcs calls = %d, want 1: a failure is remembered, not retried per resource", n)
	}
}
