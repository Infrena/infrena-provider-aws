package cfn

import (
	"os"
	"strings"
	"testing"
)

func load(t *testing.T, file string) *Schema {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}
	s, err := Parse(raw)
	if err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	return s
}

func TestParseReadsTheFieldsTheGeneratorUses(t *testing.T) {
	s := load(t, "aws-ec2-vpc.json")
	if s.TypeName != "AWS::EC2::VPC" || !s.Provisionable() || !s.HasHandler("update") {
		t.Fatalf("type=%q provisionable=%v update=%v", s.TypeName, s.Provisionable(), s.HasHandler("update"))
	}
	if strings.Join(s.PrimaryIdentifier, ",") != "/properties/VpcId" {
		t.Errorf("identifier = %v", s.PrimaryIdentifier)
	}
	if !s.TopLevel(s.CreateOnlyProperties)["CidrBlock"] || !s.TopLevel(s.ReadOnlyProperties)["VpcId"] {
		t.Error("top-level flags not read")
	}
	if len(s.Nested(s.ReadOnlyProperties)) == 0 {
		t.Error("the VPC schema has nested read-only pointers; Nested returned none")
	}
	if s.Tagging == nil || s.Tagging.TagProperty != "/properties/Tags" {
		t.Errorf("tagging = %+v", s.Tagging)
	}
	if s.Timeout("create") <= 0 {
		t.Errorf("create timeout = %d", s.Timeout("create"))
	}
}

func TestResolveFollowsReferencesToAnArrayOfObjects(t *testing.T) {
	s := load(t, "aws-ec2-securitygroup.json")
	ingress := s.Properties["SecurityGroupIngress"]
	if ingress.InsertionOrder == nil || *ingress.InsertionOrder {
		t.Fatalf("SecurityGroupIngress insertionOrder = %v, want false", ingress.InsertionOrder)
	}
	item := s.Resolve(ingress.Items)
	if item == nil || item.Properties["CidrIp"] == nil || item.Properties["IpProtocol"] == nil {
		t.Fatalf("Ingress item did not resolve to its definition: %+v", item)
	}
}

func TestListNeedsModelAndCompositeIdentifiers(t *testing.T) {
	if !load(t, "aws-aps-anomalydetector.json").ListNeedsModel() {
		t.Error("AnomalyDetector lists only under a workspace; ListNeedsModel = false")
	}
	if load(t, "aws-ec2-vpc.json").ListNeedsModel() {
		t.Error("VPC lists without a model; ListNeedsModel = true")
	}
	if n := len(load(t, "aws-acmpca-certificate.json").PrimaryIdentifier); n < 2 {
		t.Errorf("ACMPCA Certificate identifier parts = %d, want composite", n)
	}
}

func TestListNeedsModelLooksInsideOneOf(t *testing.T) {
	// The ELBv2 Listener list handler states its requirement (LoadBalancerArn or ListenerArns) only inside a
	// top-level oneOf, with no top-level required. A version of ListNeedsModel that only checks the top level
	// misses this and reports the type as plainly listable, which is wrong: AWS refuses ListResources for it with
	// no parent resource model (found live 2026-09-16).
	if !load(t, "aws-elasticloadbalancingv2-listener.json").ListNeedsModel() {
		t.Error("Listener lists only under a load balancer (required nested in oneOf); ListNeedsModel = false")
	}
	if load(t, "aws-ec2-vpc.json").ListNeedsModel() {
		t.Error("VPC lists without a model and has no oneOf; ListNeedsModel = true")
	}
}

func TestTypeAcceptsAStringOrAList(t *testing.T) {
	s, err := Parse([]byte(`{"typeName":"AWS::X::Y","properties":{"A":{"type":"string"},"B":{"type":["object","string"]}}}`))
	if err != nil {
		t.Fatal(err)
	}
	if got := s.Properties["A"].Type; len(got) != 1 || got[0] != "string" {
		t.Errorf("A type = %v", got)
	}
	if got := s.Properties["B"].Type; len(got) != 2 {
		t.Errorf("B type = %v", got)
	}
}

func TestResolveIsLoopSafe(t *testing.T) {
	s, err := Parse([]byte(`{"typeName":"AWS::X::Y","properties":{"A":{"$ref":"#/definitions/L"}},
		"definitions":{"L":{"$ref":"#/definitions/M"},"M":{"$ref":"#/definitions/L"}}}`))
	if err != nil {
		t.Fatal(err)
	}
	if got := s.Resolve(s.Properties["A"]); got != nil {
		t.Errorf("a reference loop resolved to %+v, want nil", got)
	}
}

func TestParseRefusesSomethingThatIsNotASchema(t *testing.T) {
	if _, err := Parse([]byte(`{"hello":"world"}`)); err == nil {
		t.Fatal("accepted a document with no typeName")
	}
}

// TestWhollyNestedSplitsPropertiesByWhetherEveryLeafIsNamed covers the gap the first live Lambda run exposed
// (2026-09-17): Code's write-only pointers are all one level down, so TopLevel dropped every one of them, the
// catalog marked Code as ordinary, Cloud Control returned it empty, and the plan never converged.
func TestWhollyNestedSplitsPropertiesByWhetherEveryLeafIsNamed(t *testing.T) {
	// Lambda's Code names all seven of its leaves write-only, so the whole property is.
	fn := load(t, "aws-lambda-function.json")
	whole, partial := fn.WhollyNested(fn.WriteOnlyProperties)
	if !whole["Code"] {
		t.Errorf("Code: every leaf is write-only, so it must be wholly write-only; whole = %v", whole)
	}
	for _, p := range partial {
		if p == "Code" {
			t.Errorf("Code was called partial, which would leave the plan unable to converge")
		}
	}
	// The fixture must be able to contradict the assertion: if Code ever stops naming every leaf, this test
	// should stop claiming it does.
	if leaves := fn.leafNames(fn.Properties["Code"]); len(leaves) != 7 {
		t.Errorf("Code leaves = %v, want the 7 the 2026-09-14 bundle declares", leaves)
	}

	// A security group names ONE leaf inside SecurityGroupIngress. Flagging the whole property would carry
	// forward the leaves AWS does return and hide real drift in them, so it must stay ordinary.
	sg := load(t, "aws-ec2-securitygroup.json")
	whole, partial = sg.WhollyNested(sg.WriteOnlyProperties)
	if whole["SecurityGroupIngress"] {
		t.Error("SecurityGroupIngress names only one leaf write-only; flagging it whole would hide drift")
	}
	found := false
	for _, p := range partial {
		if p == "SecurityGroupIngress" {
			found = true
		}
	}
	if !found {
		t.Errorf("SecurityGroupIngress should be reported partial so it is visible, not silent; partial = %v", partial)
	}
}
