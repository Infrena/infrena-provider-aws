package awsprov

import (
	"strings"
	"testing"
)

// TestEveryTypeIsValidAndPrefixed. The host refuses the whole plugin for either failure, naming the
// plugin rather than the definition, so catch it here first.
func TestEveryTypeIsValidAndPrefixed(t *testing.T) {
	defs := definitions()
	if len(defs) != 2 {
		t.Fatalf("got %d definitions, want 2", len(defs))
	}
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Errorf("%s: %v", d.Type, err)
		}
		if !strings.HasPrefix(d.Type, PluginName+".") {
			t.Errorf("%s is not prefixed %s.", d.Type, PluginName)
		}
	}
}

// TestRegionalAttributesAreRequiredAndForceNew. A resource cannot move between regions or VPCs, and
// every attribute AWS always reports must be Required, Computed or defaulted, or infrata plans
// "removed from configuration" forever (internal/planner/diff.go, diffAttributes).
func TestRegionalAttributesAreRequiredAndForceNew(t *testing.T) {
	want := map[string][]string{
		typeVPC:    {"cidr", "region"},
		typeSubnet: {"availability_zone", "cidr", "region", "vpc_id"},
	}
	for _, d := range definitions() {
		got := d.ForceNewAttributes()
		if strings.Join(got, ",") != strings.Join(want[d.Type], ",") {
			t.Errorf("%s ForceNew = %v, want %v", d.Type, got, want[d.Type])
		}
		for _, name := range want[d.Type] {
			if a, _ := d.Attribute(name); !a.Required {
				t.Errorf("%s.%s is not Required", d.Type, name)
			}
		}
		for name, a := range d.Attributes {
			if !a.Required && !a.Computed && a.Default == nil && name != "tags" {
				t.Errorf("%s.%s is optional with no default: AWS reports it, so every plan would show it removed", d.Type, name)
			}
		}
	}
}

func TestASubnetRequiresAVPC(t *testing.T) {
	for _, d := range definitions() {
		if d.Type != typeSubnet {
			continue
		}
		if len(d.Requirements) != 1 || d.Requirements[0].Types[0] != typeVPC || d.Requirements[0].Optional {
			t.Fatalf("aws.subnet requirements = %+v, want one mandatory aws.vpc", d.Requirements)
		}
	}
}
