package gen

import (
	"os"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/cfn"
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
	cat, _, err := Generate(fixtureSchemas(t), "sha-of-bundle", &Lock{}, &ReferenceLock{}, overlay(t), true)
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
		t.Errorf("discover default = %v, want the overlay's list in infrena names", cat.DiscoverDefault)
	}
}

// TestGenerateCarriesOnlyAcceptedAndApprovedReferences: Subnet.VpcId is tier 1, DBInstance.MonitoringRoleArn is a
// pending tier-2 edge until IAM::Role is approved, and the overlay's Subnet requirement is backed by an edge.
func TestGenerateCarriesOnlyAcceptedAndApprovedReferences(t *testing.T) {
	refs := &ReferenceLock{}
	cat, warnings, err := Generate(fixtureSchemas(t), "x", &Lock{}, refs, overlay(t), true)
	if err != nil {
		t.Fatal(err)
	}
	subnet, _ := cat.Lookup("aws.subnet")
	if r := attr(t, subnet, "VpcId").References; r == nil || r.Type != "aws.vpc" || r.Attribute != "VpcId" {
		t.Errorf("Subnet.VpcId references = %+v, want aws.vpc's VpcId in infrena names", r)
	}
	if r := attr(t, subnet, "CidrBlock").References; r != nil {
		t.Errorf("CidrBlock references = %+v, want nil", r)
	}
	rds, _ := cat.Lookup("aws.dbinstance")
	if r := attr(t, rds, "MonitoringRoleArn").References; r != nil {
		t.Errorf("a pending tier-2 edge reached the catalog: %+v", r)
	}
	if status(t, refs, "AWS::RDS::DBInstance.MonitoringRoleArn") != StatusPending {
		t.Errorf("lock = %+v", refs.References)
	}
	for _, w := range warnings {
		if strings.HasPrefix(w, "requirement") {
			t.Errorf("unexpected requirement warning: %s", w)
		}
	}

	o := overlay(t)
	o.References.ApproveTargets = []string{"AWS::IAM::Role"}
	cat, _, err = Generate(fixtureSchemas(t), "x", &Lock{}, refs, o, false)
	if err != nil {
		t.Fatal(err)
	}
	rds, _ = cat.Lookup("aws.dbinstance")
	if r := attr(t, rds, "MonitoringRoleArn").References; r == nil || r.Type != "aws.role" || r.Attribute != "Arn" {
		t.Errorf("approved MonitoringRoleArn references = %+v", r)
	}

	if _, _, err := Generate(fixtureSchemas(t), "x", &Lock{}, &ReferenceLock{}, overlay(t), false); err == nil || !strings.Contains(err.Error(), AcceptNewReferencesFlag) {
		t.Errorf("an empty reference lock generated without the flag: err = %v", err)
	}
	o = overlay(t)
	o.References.ApproveTargets = []string{"AWS::IAM::Rolee"}
	if _, _, err := Generate(fixtureSchemas(t), "x", &Lock{}, &ReferenceLock{}, o, true); err == nil || !strings.Contains(err.Error(), "AWS::IAM::Rolee") {
		t.Errorf("a misspelt approve target was accepted: err = %v", err)
	}
}

func TestGenerateSkipsTypesCloudControlCannotManage(t *testing.T) {
	schemas := fixtureSchemas(t)
	nonProvisionable, _ := cfn.Parse([]byte(`{"typeName":"AWS::AppMesh::Mesh","properties":{"MeshName":{"type":"string"}},"handlers":{"read":{}}}`))
	cat, _, err := Generate(append(schemas, nonProvisionable), "x", &Lock{}, &ReferenceLock{}, overlay(t), true)
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
	if _, _, err := Generate(fixtureSchemas(t), "x", &Lock{}, &ReferenceLock{}, o, true); err == nil || !strings.Contains(err.Error(), "AWS::EC2::VPCC") {
		t.Fatalf("err = %v", err)
	}
}
