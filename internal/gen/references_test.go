package gen

import (
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

// inline builds a provisionable schema from a type name, its properties, the read-only ones and a primary identifier.
func inline(t *testing.T, typeName string, props []string, readOnly []string, primary ...string) *cfn.Schema {
	t.Helper()
	s := &cfn.Schema{TypeName: typeName, Properties: map[string]*cfn.Node{},
		Handlers: map[string]cfn.Handler{"create": {}, "read": {}, "delete": {}}}
	for _, p := range props {
		s.Properties[p] = &cfn.Node{Type: cfn.TypeList{"string"}}
	}
	for _, p := range readOnly {
		s.Properties[p] = &cfn.Node{Type: cfn.TypeList{"string"}}
		s.ReadOnlyProperties = append(s.ReadOnlyProperties, "/properties/"+p)
	}
	for _, p := range primary {
		s.PrimaryIdentifier = append(s.PrimaryIdentifier, "/properties/"+p)
	}
	return s
}

func derivedFor(t *testing.T, derived []Derivation, source, property string) Derivation {
	t.Helper()
	for _, d := range derived {
		if d.Source == source && d.Property == property {
			return d
		}
	}
	t.Fatalf("%s.%s is not a candidate; derived %d", source, property, len(derived))
	return Derivation{}
}

func notCandidate(t *testing.T, derived []Derivation, source, property string) {
	t.Helper()
	for _, d := range derived {
		if d.Source == source && d.Property == property {
			t.Errorf("%s.%s is a candidate (%+v); it should not be", source, property, d)
		}
	}
}

// TestFixturesDeriveTheRelationshipsTheirNamesSay covers tier 1, tier 2 and the negatives on real schemas.
func TestFixturesDeriveTheRelationshipsTheirNamesSay(t *testing.T) {
	derived := DeriveReferences(fixtureSchemas(t))
	cases := []struct {
		source, property, target, attribute string
		tier                                int
	}{
		{"AWS::EC2::Subnet", "VpcId", "AWS::EC2::VPC", "VpcId", 1},
		{"AWS::EC2::SecurityGroup", "VpcId", "AWS::EC2::VPC", "VpcId", 1},
		{"AWS::ACMPCA::Certificate", "CertificateAuthorityArn", "AWS::ACMPCA::CertificateAuthority", "Arn", 1},
		{"AWS::RDS::DBInstance", "MonitoringRoleArn", "AWS::IAM::Role", "Arn", 2},
	}
	for _, c := range cases {
		d := derivedFor(t, derived, c.source, c.property)
		if d.Target != c.target || d.Attribute != c.attribute || d.Tier != c.tier {
			t.Errorf("%s.%s -> %s.%s tier %d (%s), want %s.%s tier %d", c.source, c.property, d.Target, d.Attribute, d.Tier, d.Unresolved, c.target, c.attribute, c.tier)
		}
	}
	// No IPAM pool type among the fixtures, and no type segment that IpamPool ends with.
	if d := derivedFor(t, derived, "AWS::EC2::Subnet", "Ipv4IpamPoolId"); d.Tier != 0 || d.Unresolved != UnresolvedNoType {
		t.Errorf("Ipv4IpamPoolId = %+v, want unresolved: no such type", d)
	}
	// Read-only properties are the resource's own identifiers, not references.
	notCandidate(t, derived, "AWS::EC2::VPC", "VpcId")
	notCandidate(t, derived, "AWS::EC2::Subnet", "SubnetId")
	notCandidate(t, derived, "AWS::IAM::Role", "Arn")
}

// TestSameServiceBreaksATieAndNothingElseDoes: EC2's Route picks EC2's Instance; a type in a third service is left
// unresolved with its candidates named rather than guessed.
func TestSameServiceBreaksATieAndNothingElseDoes(t *testing.T) {
	schemas := []*cfn.Schema{
		inline(t, "AWS::EC2::Route", []string{"InstanceId"}, nil),
		inline(t, "AWS::EC2::Instance", nil, []string{"InstanceId"}, "InstanceId"),
		inline(t, "AWS::Connect::Instance", nil, []string{"Id", "Arn"}, "Arn"),
		inline(t, "AWS::Lightsail::Instance", nil, []string{"InstanceName", "InstanceId"}, "InstanceName"),
		inline(t, "AWS::SSM::Association", []string{"InstanceId"}, nil),
	}
	derived := DeriveReferences(schemas)
	if d := derivedFor(t, derived, "AWS::EC2::Route", "InstanceId"); d.Target != "AWS::EC2::Instance" || d.Attribute != "InstanceId" || d.Tier != 1 {
		t.Errorf("Route.InstanceId = %+v, want EC2's Instance", d)
	}
	d := derivedFor(t, derived, "AWS::SSM::Association", "InstanceId")
	if d.Tier != 0 || d.Unresolved != UnresolvedAmbiguous || strings.Join(d.Candidates, ",") != "AWS::Connect::Instance,AWS::EC2::Instance,AWS::Lightsail::Instance" {
		t.Errorf("Association.InstanceId = %+v, want ambiguous between the three Instance types", d)
	}
}

// TestPrimaryIdentifierPluralsAndCase: ManagedPolicy has no Arn property, only its identifier PolicyArn; Arns maps to
// Arn; the segment match ignores case (IpamPool vs IPAMPool).
func TestPrimaryIdentifierPluralsAndCase(t *testing.T) {
	schemas := []*cfn.Schema{
		inline(t, "AWS::IAM::Role", []string{"ManagedPolicyArns"}, []string{"Arn", "RoleId"}, "RoleName"),
		inline(t, "AWS::IAM::ManagedPolicy", nil, []string{"PolicyArn"}, "PolicyArn"),
		inline(t, "AWS::EC2::IPAMPool", nil, []string{"IpamPoolId", "Arn"}, "IpamPoolId"),
		inline(t, "AWS::EC2::VPC", []string{"Ipv4IpamPoolId"}, []string{"VpcId"}, "VpcId"),
		inline(t, "AWS::EC2::VPCEndpoint", []string{"VpcId", "SubnetIds"}, nil, "Id"),
		inline(t, "AWS::EC2::Subnet", nil, []string{"SubnetId"}, "SubnetId"),
	}
	derived := DeriveReferences(schemas)
	if d := derivedFor(t, derived, "AWS::IAM::Role", "ManagedPolicyArns"); d.Target != "AWS::IAM::ManagedPolicy" || d.Attribute != "PolicyArn" || d.Tier != 1 {
		t.Errorf("ManagedPolicyArns = %+v", d)
	}
	if d := derivedFor(t, derived, "AWS::EC2::VPCEndpoint", "VpcId"); d.Target != "AWS::EC2::VPC" || d.Attribute != "VpcId" {
		t.Errorf("VPCEndpoint.VpcId = %+v", d)
	}
	if d := derivedFor(t, derived, "AWS::EC2::VPCEndpoint", "SubnetIds"); d.Target != "AWS::EC2::Subnet" || d.Attribute != "SubnetId" || d.Tier != 1 {
		t.Errorf("SubnetIds = %+v", d)
	}
	// Ipv4IpamPool ends with IpamPool: a suffix match, so tier 2.
	if d := derivedFor(t, derived, "AWS::EC2::VPC", "Ipv4IpamPoolId"); d.Target != "AWS::EC2::IPAMPool" || d.Attribute != "IpamPoolId" || d.Tier != 2 {
		t.Errorf("Ipv4IpamPoolId = %+v", d)
	}
}

// TestSuffixFallbackTakesTheLongestSegment: MonitoringIpamPoolId ends with both IpamPool and Pool; IpamPool wins. And an
// exact match is never demoted to the fallback.
func TestSuffixFallbackTakesTheLongestSegment(t *testing.T) {
	schemas := []*cfn.Schema{
		inline(t, "AWS::EC2::IPAMPool", nil, []string{"IpamPoolId"}, "IpamPoolId"),
		inline(t, "AWS::EC2::Pool", nil, []string{"PoolId"}, "PoolId"),
		inline(t, "AWS::EC2::Thing", []string{"MonitoringIpamPoolId", "PoolId"}, nil),
	}
	derived := DeriveReferences(schemas)
	if d := derivedFor(t, derived, "AWS::EC2::Thing", "MonitoringIpamPoolId"); d.Target != "AWS::EC2::IPAMPool" || d.Tier != 2 {
		t.Errorf("MonitoringIpamPoolId = %+v, want the longer IpamPool segment", d)
	}
	if d := derivedFor(t, derived, "AWS::EC2::Thing", "PoolId"); d.Target != "AWS::EC2::Pool" || d.Tier != 1 {
		t.Errorf("PoolId = %+v, want an exact tier-1 match", d)
	}
}

// TestUnresolvedReasons: a type without the attribute kind, a target Cloud Control cannot provision, and property
// names that are only a suffix.
func TestUnresolvedReasons(t *testing.T) {
	notProvisionable := inline(t, "AWS::B::Gadget", nil, []string{"GadgetId"}, "GadgetId")
	delete(notProvisionable.Handlers, "delete")
	schemas := []*cfn.Schema{
		inline(t, "AWS::A::Widget", nil, []string{"Name"}, "Name"),
		notProvisionable,
		inline(t, "AWS::C::User", []string{"WidgetArn", "GadgetId", "Id", "Arns"}, nil),
	}
	derived := DeriveReferences(schemas)
	if d := derivedFor(t, derived, "AWS::C::User", "WidgetArn"); d.Tier != 0 || d.Unresolved != UnresolvedNoAttribute {
		t.Errorf("WidgetArn = %+v, want unresolved: target lacks the attribute", d)
	}
	if d := derivedFor(t, derived, "AWS::C::User", "GadgetId"); d.Tier != 0 || d.Unresolved != UnresolvedNoType {
		t.Errorf("GadgetId = %+v, want unresolved: a non-provisionable type is no candidate", d)
	}
	notCandidate(t, derived, "AWS::C::User", "Id")
	notCandidate(t, derived, "AWS::C::User", "Arns")
	for _, d := range derived {
		if d.Source == "AWS::B::Gadget" {
			t.Errorf("a non-provisionable type was a source: %+v", d)
		}
	}
}

// TestDerivationIsSorted: the lock and the review are written from it, so order must not depend on map iteration.
func TestDerivationIsSorted(t *testing.T) {
	derived := DeriveReferences(fixtureSchemas(t))
	for i := 1; i < len(derived); i++ {
		a, b := derived[i-1], derived[i]
		if a.Source > b.Source || (a.Source == b.Source && a.Property >= b.Property) {
			t.Fatalf("not sorted at %d: %s.%s then %s.%s", i, a.Source, a.Property, b.Source, b.Property)
		}
	}
}
