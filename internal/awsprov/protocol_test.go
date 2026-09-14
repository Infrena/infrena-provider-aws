package awsprov

import (
	"context"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/plugintest"
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
