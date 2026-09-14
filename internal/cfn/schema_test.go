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
