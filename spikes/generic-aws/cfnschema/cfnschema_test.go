package cfnschema

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/infrata/infrata/pkg/plugintest"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/schema"
)

func load(t *testing.T, file string) (*schema.ResourceDefinition, *Report) {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	def, rep := Definition(doc)
	return def, rep
}

// TestTheRealVPCSchemaMapsToTheFlagsWeHandWrote. Task 1 hand-wrote aws.vpc: cidr ForceNew, id Computed.
// The real AWS::EC2::VPC schema, fetched with DescribeType on 2026-09-13, must say the same by itself.
func TestTheRealVPCSchemaMapsToTheFlagsWeHandWrote(t *testing.T) {
	def, rep := load(t, "aws-ec2-vpc.json")
	if def.Type != "aws.ec2.vpc" {
		t.Errorf("type = %q", def.Type)
	}
	if a := def.Attributes["CidrBlock"]; !a.ForceNew || a.Computed {
		t.Errorf("CidrBlock = %+v, want ForceNew and not Computed", a)
	}
	if a := def.Attributes["VpcId"]; !a.Computed {
		t.Errorf("VpcId = %+v, want Computed", a)
	}
	if a := def.Attributes["EnableDnsHostnames"]; a.ForceNew || a.Computed {
		t.Errorf("EnableDnsHostnames = %+v, want updatable in place", a)
	}
	if a := def.Attributes["Tags"]; a.Kind.String() != "list" {
		t.Errorf("Tags kind = %s: CloudFormation tags are a list of {Key, Value}, not a map", a.Kind)
	}
	if strings.Join(rep.Identifier, ",") != "VpcId" {
		t.Errorf("identifier = %v", rep.Identifier)
	}
	if err := def.Validate(); err != nil {
		t.Fatal(err)
	}
	t.Logf("write-only (carry forward on Read): %v", rep.WriteOnly)
	t.Logf("conditionally create-only (replace sometimes): %v", rep.Conditional)
	t.Logf("nested flags infrata cannot express: %d", len(rep.NestedFlags))
	t.Logf("kinds guessed: %v", rep.Guessed)
}

func TestTheRealSubnetSchemaNeedsAVpcIdAndReplacesOnAZ(t *testing.T) {
	def, _ := load(t, "aws-ec2-subnet.json")
	if a := def.Attributes["VpcId"]; !a.Required || !a.ForceNew {
		t.Errorf("VpcId = %+v, want Required and ForceNew", a)
	}
	if a := def.Attributes["AvailabilityZone"]; a.Required || !a.ForceNew {
		t.Errorf("AvailabilityZone = %+v: the schema makes it optional (AWS picks one) and create-only", a)
	}
	if a := def.Attributes["MapPublicIpOnLaunch"]; a.ForceNew {
		t.Errorf("MapPublicIpOnLaunch = %+v, want updatable", a)
	}
}

type spikePlugin struct{ defs []*schema.ResourceDefinition }

func (p spikePlugin) Name() string                              { return "aws" }
func (p spikePlugin) Definitions() []*schema.ResourceDefinition { return p.defs }
func (p spikePlugin) New(provider.Config) (provider.Provider, error) {
	return nil, errors.New("spike: no instances")
}

// TestInfrataAcceptsGeneratedDefinitions. The real test of "no handwritten knowledge": infrata's own
// plugin host must accept what the mapping produced, through the protocol.
func TestInfrataAcceptsGeneratedDefinitions(t *testing.T) {
	vpc, _ := load(t, "aws-ec2-vpc.json")
	subnet, _ := load(t, "aws-ec2-subnet.json")
	host, err := plugintest.Open(context.Background(), spikePlugin{defs: []*schema.ResourceDefinition{vpc, subnet}}, t.TempDir())
	if err != nil {
		t.Fatalf("infrata's host refused the generated definitions: %v", err)
	}
	defer host.Close()
	if n := len(host.Definitions()); n != 2 {
		t.Fatalf("host holds %d definitions", n)
	}
}
