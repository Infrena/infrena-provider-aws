package awsprov

import (
	"context"
	"testing"

	"github.com/infrata/infrata/pkg/plugintest"
)

// openHost connects this plugin to infrata's own host over an in-memory pipe: schema validation,
// the prefix rule and reserved names all apply, exactly as for a subprocess.
func openHost(t *testing.T) *plugintest.Host {
	t.Helper()
	host, err := plugintest.Open(context.Background(), NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatalf("the host refused this plugin's schemas: %v", err)
	}
	t.Cleanup(func() { _ = host.Close() })
	return host
}

// TestSchemasLoadThroughTheHost. The subnet's map_public_ip_on_launch default is the one datum
// that must survive JSON as a bool.
func TestSchemasLoadThroughTheHost(t *testing.T) {
	host := openHost(t)
	if got := host.Version(); got != Version {
		t.Errorf("handshake version = %q, want %q", got, Version)
	}
	var subnet bool
	for _, d := range host.Definitions() {
		if d.Type != typeSubnet {
			continue
		}
		subnet = true
		a, _ := d.Attribute("map_public_ip_on_launch")
		if b, ok := a.Default.(bool); !ok || b {
			t.Errorf("map_public_ip_on_launch default after the wire = %#v (%T), want false", a.Default, a.Default)
		}
	}
	if !subnet {
		t.Fatal("aws.subnet did not arrive")
	}
}
