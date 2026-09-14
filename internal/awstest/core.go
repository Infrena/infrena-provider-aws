package awstest

import (
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
)

// TypeFor finds a catalog type by CloudFormation name, so a test does not depend on the name the generator assigned.
func TypeFor(t testing.TB, cat *catalog.Catalog, cfn string) *catalog.Type {
	t.Helper()
	for _, typ := range cat.Types {
		if typ.CFN == cfn {
			return typ
		}
	}
	t.Fatalf("the catalog has no %s", cfn)
	return nil
}

// RegisterCore describes the types the protocol, e2e and live suites use to the fake, with the defaults AWS chooses
// for properties those suites leave unset: the Optional+Computed attributes whose values must never plan a change.
func RegisterCore(t testing.TB, fake *ccfake.Server, cat *catalog.Catalog) {
	t.Helper()
	egress := []any{map[string]any{"IpProtocol": "-1", "CidrIp": "0.0.0.0/0"}}
	for cfn, c := range map[string]struct {
		prefix   string
		defaults map[string]any
	}{
		"AWS::EC2::VPC":           {"vpc-", map[string]any{"EnableDnsSupport": true, "EnableDnsHostnames": false, "InstanceTenancy": "default"}},
		"AWS::EC2::Subnet":        {"subnet-", map[string]any{"MapPublicIpOnLaunch": false, "AvailabilityZoneId": "use1-az1"}},
		"AWS::EC2::SecurityGroup": {"sg-", map[string]any{"SecurityGroupEgress": egress}},
		"AWS::IAM::Role":          {"role-", map[string]any{"MaxSessionDuration": 3600, "Path": "/"}},
		"AWS::RDS::DBInstance":    {"db-", map[string]any{"Engine": "postgres", "StorageEncrypted": false}},
	} {
		fake.Register(FakeType(TypeFor(t, cat, cfn), c.prefix, c.defaults))
	}
}
