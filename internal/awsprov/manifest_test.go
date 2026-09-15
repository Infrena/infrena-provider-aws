package awsprov

import (
	"os"
	"testing"

	"github.com/infrena/infrena/pkg/pluginmanifest"
	"github.com/infrena/infrena/pkg/pluginproto"
)

// readManifest parses the repository's plugin.yaml with infrena's own parser: the same code
// `infrena plugins install` runs against it (infrena PLAN.md §31.2), so a manifest that passes here is one
// install will accept rather than one a second parser merely agreed with.
func readManifest(t *testing.T) *pluginmanifest.Manifest {
	t.Helper()
	data, err := os.ReadFile("../../plugin.yaml")
	if err != nil {
		t.Fatal(err)
	}
	m, warnings, err := pluginmanifest.Parse(data)
	if err != nil {
		t.Fatalf("plugin.yaml is not a manifest infrena accepts: %v", err)
	}
	if len(warnings) != 0 {
		t.Errorf("plugin.yaml parses with warnings, which install would print: %v", warnings)
	}
	return m
}

// TestTheManifestDescribesThisPlugin. The manifest is authoritative for a release (infrena PLAN.md §31.2), so it
// must name the plugin this binary is and list exactly the protocol this SDK speaks. The version is not
// compared here: the code reports 0.0.0-dev until a release stamps it, and scripts/release-check
// is what asserts tag == manifest == binary.
//
// Exactly, not "includes": §31.2 (amended 2026-09-14) defines `protocol:` as the versions this
// release's binary speaks, and an SDK-built binary announces one. A manifest listing the host's
// whole Supported set, [2, 1], claims a protocol this binary cannot speak.
func TestTheManifestDescribesThisPlugin(t *testing.T) {
	m := readManifest(t)
	if m.Name != PluginName {
		t.Errorf("plugin.yaml names %q, but this plugin is %q", m.Name, PluginName)
	}
	if len(m.Protocol) != 1 || m.Protocol[0] != pluginproto.Version {
		t.Errorf("plugin.yaml's protocol is %v, but the SDK this plugin is built with speaks exactly [%d]; "+
			"change it in the same commit as go.mod's infrena require",
			m.Protocol, pluginproto.Version)
	}
}

// TestTheManifestFloorIsAtLeastTheRequiredRelease. go.mod's require is the oldest infrena release CI
// verifies this plugin against, so the manifest's floor must admit at least that release, and never
// admit a release published under the product's old name (which cannot pair with this plugin's
// module path, CLI, or plugin binary name). The floor is 0.6.2: the release that fixed the compiler
// bug where a whole-resource reference into an aliased attribute (`vpc: ${vpc}`, `vpc_id: ${vpc}`)
// wrongly failed with "declares no reference"; 0.6.0 and 0.6.1 admit only the canonical AWS name.
func TestTheManifestFloorIsAtLeastTheRequiredRelease(t *testing.T) {
	m := readManifest(t)
	if m.Infrena.IsZero() {
		t.Fatal("plugin.yaml has no infrena: floor, so it claims to work with releases under the old name too")
	}
	for version, want := range map[string]bool{"0.3.9": false, "0.4.9": false, "0.5.0": false, "0.5.9": false, "0.6.0": false, "0.6.1": false, "0.6.2": true} {
		if got := m.AllowsInfrena(version); got != want {
			t.Errorf("plugin.yaml's infrena: %q allows %s = %v, want %v", m.Infrena, version, got, want)
		}
	}
}
