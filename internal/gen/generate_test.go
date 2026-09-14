package gen

import (
	"os"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

func fixtureSchemas(t *testing.T) []*cfn.Schema {
	t.Helper()
	var out []*cfn.Schema
	for _, file := range fixtureFiles {
		raw, err := os.ReadFile("../cfn/testdata/" + file)
		if err != nil {
			t.Fatal(err)
		}
		s, err := cfn.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, s)
	}
	return out
}

func TestGenerateNamesBuildsAndKeepsTheBundleHash(t *testing.T) {
	cat, _, err := Generate(fixtureSchemas(t), "sha-of-bundle", &Lock{}, overlay(t))
	if err != nil {
		t.Fatal(err)
	}
	if cat.Bundle != "sha-of-bundle" || len(cat.Types) != len(fixtureFiles) {
		t.Fatalf("bundle=%q types=%d", cat.Bundle, len(cat.Types))
	}
	vpc, ok := cat.Lookup("aws.vpc")
	if !ok || vpc.CFN != "AWS::EC2::VPC" {
		t.Fatalf("aws.vpc = %+v", vpc)
	}
	if strings.Join(cat.DiscoverDefault, ",") != "aws.vpc" {
		t.Errorf("discover default = %v, want the overlay's list in infrata names", cat.DiscoverDefault)
	}
}

func TestGenerateSkipsTypesCloudControlCannotManage(t *testing.T) {
	schemas := fixtureSchemas(t)
	nonProvisionable, _ := cfn.Parse([]byte(`{"typeName":"AWS::AppMesh::Mesh","properties":{"MeshName":{"type":"string"}},"handlers":{"read":{}}}`))
	cat, _, err := Generate(append(schemas, nonProvisionable), "x", &Lock{}, overlay(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range cat.Types {
		if typ.CFN == "AWS::AppMesh::Mesh" {
			t.Fatal("a type without create/read/delete handlers was generated")
		}
	}
}

func TestGenerateRefusesAnOverlayNamingAnUnknownType(t *testing.T) {
	o := overlay(t)
	o.Aliases["AWS::EC2::VPCC"] = map[string][]string{"CidrBlock": {"cidr"}}
	if _, _, err := Generate(fixtureSchemas(t), "x", &Lock{}, o); err == nil || !strings.Contains(err.Error(), "AWS::EC2::VPCC") {
		t.Fatalf("err = %v", err)
	}
}
