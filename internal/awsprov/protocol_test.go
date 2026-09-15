package awsprov

import (
	"context"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/plugintest"
	"github.com/infrena/infrena/pkg/schema"
)

func openHost(t *testing.T) *plugintest.Host {
	t.Helper()
	host, err := plugintest.Open(context.Background(), NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatalf("the host refused this plugin's schemas: %v", err)
	}
	t.Cleanup(func() { _ = host.Close() })
	return host
}

// TestTheWholeCatalogLoadsThroughTheHost: every generated definition crosses the wire and passes infrena's load checks
// (prefix, reserved names, validation, alias folding) at protocol 2.
func TestTheWholeCatalogLoadsThroughTheHost(t *testing.T) {
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	host := openHost(t)
	if got := host.Version(); got != Version {
		t.Errorf("handshake version = %q, want %q", got, Version)
	}
	defs := host.Definitions()
	if len(defs) != len(cat.Types) {
		t.Fatalf("host holds %d definitions, catalog has %d", len(defs), len(cat.Types))
	}
	for _, d := range defs {
		if d.Type != "aws.vpc" {
			continue
		}
		if got := d.Display("CidrBlock"); got != "cidr" {
			t.Errorf("aws.vpc CidrBlock displays as %q after the wire, want cidr", got)
		}
		return
	}
	t.Fatal("aws.vpc did not arrive")
}

// TestTheWholeCatalogPassesValidateAll: plugintest.Open does not run schema.ValidateAll in infrena v0.6.0, and the
// host does on load, so a dangling reference would fail a user's load and no test. This runs it on the definitions as
// they arrive over the wire, and checks every accepted and approved edge arrived as a References and nothing else did.
func TestTheWholeCatalogPassesValidateAll(t *testing.T) {
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	if err := schema.ValidateAll(cat.Definitions()); err != nil {
		t.Fatalf("the catalog's definitions fail infrena's ValidateAll: %v", err)
	}
	defs := openHost(t).Definitions()
	if err := schema.ValidateAll(defs); err != nil {
		t.Fatalf("the definitions after the wire fail infrena's ValidateAll: %v", err)
	}
	want := 0
	for _, typ := range cat.Types {
		for _, a := range typ.Attributes {
			if a.References != nil {
				want++
			}
		}
	}
	got := 0
	var subnetVpc *schema.Reference
	for _, d := range defs {
		for name, a := range d.Attributes {
			if a.Fields != nil {
				t.Errorf("%s.%s declares Fields; every map is open on purpose", d.Type, name)
			}
			if a.References == nil {
				continue
			}
			got++
			if d.Type == "aws.subnet" && name == "VpcId" {
				subnetVpc = a.References
			}
		}
	}
	if want == 0 || got != want {
		t.Errorf("%d attributes arrived with References, the catalog carries %d", got, want)
	}
	if subnetVpc == nil || *subnetVpc != (schema.Reference{Type: "aws.vpc", Attribute: "VpcId"}) {
		t.Errorf("aws.subnet VpcId references %+v, want aws.vpc's VpcId", subnetVpc)
	}
}
