package gen

import (
	"os"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

var testNames = map[string]string{
	"AWS::EC2::VPC": "aws.vpc", "AWS::EC2::Subnet": "aws.subnet", "AWS::EC2::SecurityGroup": "aws.securitygroup",
	"AWS::S3::Bucket": "aws.bucket", "AWS::IAM::Role": "aws.role", "AWS::RDS::DBInstance": "aws.dbinstance",
	"AWS::ACMPCA::Certificate": "aws.acmpca.certificate", "AWS::ACMPCA::CertificateAuthority": "aws.certificateauthority",
	"AWS::APS::AnomalyDetector":                               "aws.anomalydetector",
	"AWS::ARCZonalShift::AutoshiftObserverNotificationStatus": "aws.autoshiftobservernotificationstatus",
	"AWS::CodePipeline::CustomActionType":                     "aws.customactiontype",
}

var fixtureFiles = map[string]string{
	"AWS::EC2::VPC": "aws-ec2-vpc.json", "AWS::EC2::Subnet": "aws-ec2-subnet.json",
	"AWS::EC2::SecurityGroup": "aws-ec2-securitygroup.json", "AWS::S3::Bucket": "aws-s3-bucket.json",
	"AWS::IAM::Role": "aws-iam-role.json", "AWS::RDS::DBInstance": "aws-rds-dbinstance.json",
	"AWS::ACMPCA::Certificate":                                "aws-acmpca-certificate.json",
	"AWS::ACMPCA::CertificateAuthority":                       "aws-acmpca-certificateauthority.json",
	"AWS::APS::AnomalyDetector":                               "aws-aps-anomalydetector.json",
	"AWS::ARCZonalShift::AutoshiftObserverNotificationStatus": "aws-arczonalshift-autoshiftobservernotificationstatus.json",
	"AWS::CodePipeline::CustomActionType":                     "aws-codepipeline-customactiontype.json",
}

func overlay(t *testing.T) *Overlay {
	t.Helper()
	o, err := LoadOverlay("testdata/overlay.yaml")
	if err != nil {
		t.Fatal(err)
	}
	return o
}

func build(t *testing.T, cfnType string) *catalog.Type {
	t.Helper()
	raw, err := os.ReadFile("../cfn/testdata/" + fixtureFiles[cfnType])
	if err != nil {
		t.Fatal(err)
	}
	s, err := cfn.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	typ, _, err := BuildType(s, testNames, overlay(t))
	if err != nil {
		t.Fatalf("%s: %v", cfnType, err)
	}
	return typ
}

func attr(t *testing.T, typ *catalog.Type, name string) *catalog.Attribute {
	t.Helper()
	a, ok := typ.Attribute(name)
	if !ok {
		t.Fatalf("%s has no attribute %s", typ.Name, name)
	}
	return a
}

// TestEveryFixtureBuildsADefinitionInfrenaAccepts: schema.Validate checks flag combinations and case-folded name
// collisions, which is exactly where a generator bug would show.
func TestEveryFixtureBuildsADefinitionInfrenaAccepts(t *testing.T) {
	for cfnType := range fixtureFiles {
		if err := build(t, cfnType).Definition().Validate(); err != nil {
			t.Errorf("%s: %v", cfnType, err)
		}
	}
}

func TestVPCFlagsSpellingsAndTags(t *testing.T) {
	vpc := build(t, "AWS::EC2::VPC")
	cidr := attr(t, vpc, "CidrBlock")
	if !cidr.Optional || !cidr.Computed || !cidr.ForceNew || cidr.Required {
		t.Errorf("CidrBlock = %+v, want Optional+Computed+ForceNew", cidr)
	}
	if strings.Join(cidr.Aliases, ",") != "cidr,cidr_block" {
		t.Errorf("CidrBlock aliases = %v, want the curated alias first, then snake_case", cidr.Aliases)
	}
	if id := attr(t, vpc, "VpcId"); !id.Computed || id.Optional || id.ForceNew {
		t.Errorf("VpcId = %+v, want Computed only", id)
	}
	if dns := attr(t, vpc, "EnableDnsHostnames"); dns.ForceNew || !dns.Optional || strings.Join(dns.Aliases, ",") != "enable_dns_hostnames" {
		t.Errorf("EnableDnsHostnames = %+v", dns)
	}
	if vpc.TagsAsMap != "Tags" || attr(t, vpc, "Tags").Kind != "map" {
		t.Errorf("tags as map = %q, kind %q", vpc.TagsAsMap, attr(t, vpc, "Tags").Kind)
	}
	if vpc.RegionAttr != "region" || strings.Join(vpc.Identifier, ",") != "VpcId" || !vpc.HasUpdate || !vpc.HasList {
		t.Errorf("vpc = %+v", vpc)
	}
}

func TestSubnetRequiresAVPCAndLetsAWSPickTheZone(t *testing.T) {
	subnet := build(t, "AWS::EC2::Subnet")
	if v := attr(t, subnet, "VpcId"); !v.Required || !v.ForceNew || v.Computed {
		t.Errorf("VpcId = %+v", v)
	}
	if az := attr(t, subnet, "AvailabilityZone"); !az.Optional || !az.Computed || !az.ForceNew {
		t.Errorf("AvailabilityZone = %+v, want Optional+Computed+ForceNew", az)
	}
	if len(subnet.Requirements) != 1 || subnet.Requirements[0].Types[0] != "aws.vpc" {
		t.Errorf("requirements = %+v, want the overlay's, in infrena names", subnet.Requirements)
	}
}

func TestSecurityGroupRulesAreUnorderedObjects(t *testing.T) {
	sh := attr(t, build(t, "AWS::EC2::SecurityGroup"), "SecurityGroupIngress").Shape
	if sh == nil || sh.Kind != catalog.ShapeArray || !sh.Unordered {
		t.Fatalf("ingress shape = %+v, want an unordered array", sh)
	}
	if sh.Item.Kind != catalog.ShapeObject || sh.Item.Props["CidrIp"] == nil || sh.Item.Props["CidrIp"].Kind != catalog.ShapeScalar {
		t.Errorf("ingress item = %+v", sh.Item)
	}
}

func TestGlobalSensitiveCompositeAndListModel(t *testing.T) {
	if role := build(t, "AWS::IAM::Role"); !role.Global() {
		t.Error("the overlay makes AWS::IAM::* global")
	}
	rds := build(t, "AWS::RDS::DBInstance")
	if !attr(t, rds, "MasterUserPassword").Sensitive {
		t.Error("the overlay marks MasterUserPassword sensitive")
	}
	if !strings.Contains(strings.Join(rds.WriteOnly, ","), "MasterUserPassword") {
		t.Errorf("write-only = %v", rds.WriteOnly)
	}
	if n := len(build(t, "AWS::ACMPCA::Certificate").Identifier); n < 2 {
		t.Errorf("composite identifier parts = %d", n)
	}
	if !build(t, "AWS::APS::AnomalyDetector").ListNeedsModel {
		t.Error("AnomalyDetector lists under a parent")
	}
}

// TestPropertiesNamedLikeInfrenaKeywordsGetUsableNames: J9.
func TestPropertiesNamedLikeInfrenaKeywordsGetUsableNames(t *testing.T) {
	ca := build(t, "AWS::ACMPCA::CertificateAuthority")
	typ := attr(t, ca, "Type")
	if len(typ.Aliases) == 0 || typ.Aliases[0] != "type_value" || ca.Definition().Display("Type") != "type_value" {
		t.Errorf("Type aliases = %v", typ.Aliases)
	}
	for _, a := range typ.Aliases {
		if strings.EqualFold(a, "type") {
			t.Errorf("alias %q would decode as infrena's resource key", a)
		}
	}
	if p := attr(t, build(t, "AWS::CodePipeline::CustomActionType"), "Provider"); p.Aliases[0] != "provider_value" {
		t.Errorf("Provider aliases = %v", p.Aliases)
	}
	zs := build(t, "AWS::ARCZonalShift::AutoshiftObserverNotificationStatus")
	if zs.RegionAttr != "aws_region" {
		t.Errorf("region attribute on a type with its own Region = %q, want aws_region", zs.RegionAttr)
	}
}

func TestOverlayRefusesUnknownKeys(t *testing.T) {
	path := t.TempDir() + "/bad.yaml"
	if err := os.WriteFile(path, []byte("globals:\n  - AWS::IAM::*\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadOverlay(path); err == nil {
		t.Fatal("a misspelt overlay key was accepted; it would silently apply nothing")
	}
}
