package awsprov

import (
	"os"
	"testing"

	"github.com/infrena/infrena/pkg/pluginmanifest"
	"github.com/infrena/infrena/pkg/pluginproto"
)

// readManifest parses the repository's plugin.yaml with infrena's own parser: the same code
// `infrena plugins install` runs against it, so a manifest that passes here is one
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

// TestTheManifestDescribesThisPlugin. The manifest is authoritative for a release, so it
// must name the plugin this binary is and list exactly the protocol this SDK speaks. The version is not
// compared here: the code reports 0.0.0-dev until a release stamps it, and scripts/release-check
// is what asserts tag == manifest == binary.
//
// Exactly, not "includes": infrena defines `protocol:` as the versions this
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
// module path, CLI, or plugin binary name).
//
// THE FLOOR IS 0.13.0 SINCE 2026-09-18. The protocol set it to 0.12.0 first; deleting every release
// before 0.13.0 as stale, tags included, set it here, because a floor may only name a release that
// exists. This binary speaks
// plugin protocol 5; every host from 0.7.1 to 0.11.1 speaks at most 4 and refuses it at the
// handshake, so admitting them would be claiming a pairing that cannot connect. `Supported` being a
// set protects an old plugin on a new host, and nothing protects the reverse.
//
// The previous floor's reasoning is kept in plugin.yaml because it still governs the NEXT one: a
// floor follows what the binary genuinely cannot run without, not what go.mod happens to require.
// It was 0.7.1, the release that fixed Update and Delete to receive the refreshed observation
// rather than the last persisted state (infrena 1399f20), which this plugin's Update relies on
// directly. That requirement has not gone away — it is simply implied by anything at or above
// 0.12.0.
func TestTheManifestFloorIsAtLeastTheRequiredRelease(t *testing.T) {
	m := readManifest(t)
	if m.Infrena.IsZero() {
		t.Fatal("plugin.yaml has no infrena: floor, so it claims to work with releases under the old name too")
	}
	for version, want := range map[string]bool{
		// Under the old name, or speaking protocol 3.
		"0.3.9": false, "0.4.9": false, "0.5.9": false, "0.6.2": false,
		// Protocol 4: fine for the previous release of this plugin, too old for this one.
		"0.7.0": false, "0.7.1": false, "0.11.1": false,
		// Protocol 5, but deleted as stale and below the floor.
		"0.12.0": false, "0.12.9": false,
		// The oldest release that still exists, and everything after it.
		"0.14.0": true, "0.14.9": true, "1.0.0": true,
	} {
		if got := m.AllowsInfrena(version); got != want {
			t.Errorf("plugin.yaml's infrena: %q allows %s = %v, want %v", m.Infrena, version, got, want)
		}
	}
}
