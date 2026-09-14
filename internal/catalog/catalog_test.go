package catalog

import (
	"bytes"
	"strings"
	"testing"
)

func sample() *Catalog {
	return &Catalog{Bundle: "abc123", Types: []*Type{
		{
			Name: "aws.vpc", CFN: "AWS::EC2::VPC", Description: "A VPC.", RegionAttr: "region",
			Identifier: []string{"VpcId"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
			Timeouts: map[string]int{"create": 120},
			Attributes: []*Attribute{
				{Name: "CidrBlock", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"cidr", "cidr_block"}},
				{Name: "VpcId", Kind: "string", Computed: true, Aliases: []string{"vpc_id"}},
				{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: &Shape{Kind: ShapeOpaque}},
			},
		},
		{
			Name: "aws.role", CFN: "AWS::IAM::Role", Identifier: []string{"RoleName"}, HasUpdate: true,
			Attributes: []*Attribute{
				{Name: "AssumeRolePolicyDocument", Kind: "map", Required: true, Shape: &Shape{Kind: ShapeOpaque}},
				{Name: "RoleName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"role_name"}},
			},
			Requirements: []Requirement{{Name: "vpc", Types: []string{"aws.vpc"}, Description: "only a test"}},
		},
	}}
}

func TestDefinitionsValidateAndCarryTheFlags(t *testing.T) {
	c := sample()
	defs := c.Definitions()
	if len(defs) != 2 {
		t.Fatalf("got %d definitions", len(defs))
	}
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Fatalf("%s: %v", d.Type, err)
		}
	}
	vpc, _ := c.Lookup("aws.vpc")
	d := vpc.Definition()
	cidr := d.Attributes["CidrBlock"]
	if !cidr.Optional || !cidr.Computed || !cidr.ForceNew || strings.Join(cidr.Aliases, ",") != "cidr,cidr_block" {
		t.Errorf("CidrBlock = %+v", cidr)
	}
	if d.Display("CidrBlock") != "cidr" {
		t.Errorf("Display = %q, want the first alias", d.Display("CidrBlock"))
	}
	region, ok := d.Attributes["region"]
	if !ok || !region.Required || !region.ForceNew {
		t.Errorf("regional type's region attribute = %+v, %v", region, ok)
	}
	if !d.Capabilities.Update || !d.Capabilities.Import {
		t.Errorf("capabilities = %+v", d.Capabilities)
	}
}

func TestAGlobalTypeHasNoRegionAndSaysHowToImport(t *testing.T) {
	role, ok := sample().Lookup("aws.role")
	if !ok || !role.Global() {
		t.Fatalf("aws.role global = %v", ok && role.Global())
	}
	d := role.Definition()
	if _, has := d.Attributes["region"]; has {
		t.Error("a global type declares region")
	}
	if !strings.Contains(d.ImportID.Description, GlobalScope+"/") {
		t.Errorf("import description = %q", d.ImportID.Description)
	}
	if len(d.Requirements) != 1 || d.Requirements[0].Types[0] != "aws.vpc" {
		t.Errorf("requirements = %+v", d.Requirements)
	}
}

func TestAnImmutableTypeCannotUpdate(t *testing.T) {
	c := sample()
	c.Types[0].HasUpdate = false
	vpc, _ := c.Lookup("aws.vpc")
	if vpc.Definition().Capabilities.Update {
		t.Error("an immutable type offers update")
	}
}

func TestCatalogRoundTripsThroughGzipJSON(t *testing.T) {
	var buf bytes.Buffer
	if err := sample().Write(&buf); err != nil {
		t.Fatal(err)
	}
	got, err := Read(&buf)
	if err != nil {
		t.Fatal(err)
	}
	vpc, ok := got.Lookup("aws.vpc")
	if !ok || got.Bundle != "abc123" || vpc.TagsAsMap != "Tags" {
		t.Fatalf("round trip lost data: %+v", got)
	}
	a, ok := vpc.Attribute("Tags")
	if !ok || a.Shape == nil || a.Shape.Kind != ShapeOpaque {
		t.Errorf("Tags shape after round trip = %+v", a)
	}
}

func TestReadRefusesDuplicateNames(t *testing.T) {
	c := sample()
	c.Types[1].Name = "aws.vpc"
	var buf bytes.Buffer
	if err := c.Write(&buf); err != nil {
		t.Fatal(err)
	}
	if _, err := Read(&buf); err == nil || !strings.Contains(err.Error(), "aws.vpc") {
		t.Fatalf("err = %v", err)
	}
}
