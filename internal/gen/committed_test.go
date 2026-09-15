package gen

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

// TestTheCommittedCatalogMatchesTheBundleItClaims regenerates from the real bundle, when present, with the committed
// overlay and lock, and compares with the committed catalog. It skips without the bundle (run scripts/fetch-schemas).
// A different bundle hash means the schemas moved on: regenerate, do not edit this test.
func TestTheCommittedCatalogMatchesTheBundleItClaims(t *testing.T) {
	const bundlePath = "../../schemas/CloudformationSchema.zip"
	if _, err := os.Stat(bundlePath); errors.Is(err, fs.ErrNotExist) {
		t.Skip("no schema bundle; run scripts/fetch-schemas")
	}
	schemas, sha, err := cfn.ReadBundle(bundlePath)
	if err != nil {
		t.Fatal(err)
	}
	committed, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	if committed.Bundle != sha {
		t.Skipf("bundle %s differs from the committed catalog's %s: regenerate with go run ./cmd/gen-cloudcontrol", sha, committed.Bundle)
	}
	lock, err := LoadLock("../../gen/names.lock.json")
	if err != nil {
		t.Fatal(err)
	}
	o, err := LoadOverlay("../../gen/overlay.yaml")
	if err != nil {
		t.Fatal(err)
	}
	refs, err := LoadReferenceLock("../../gen/references.lock.json")
	if err != nil {
		t.Fatal(err)
	}
	before := len(lock.Names)
	cat, _, err := Generate(schemas, sha, lock, refs, o, ReferenceFlags{})
	if err != nil {
		t.Fatal(err)
	}
	if len(lock.Names) != before {
		t.Errorf("regeneration assigned %d new names: the committed lock is stale", len(lock.Names)-before)
	}
	committedRefs, err := os.ReadFile("../../gen/references.lock.json")
	if err != nil {
		t.Fatal(err)
	}
	regenerated := t.TempDir() + "/references.lock.json"
	if err := refs.Save(regenerated); err != nil {
		t.Fatal(err)
	}
	if again, _ := os.ReadFile(regenerated); string(again) != string(committedRefs) {
		t.Error("regeneration changed gen/references.lock.json: the committed reference lock is stale")
	}
	usable := 0
	for _, r := range refs.References {
		if r.Usable() {
			usable++
		}
	}
	carried := 0
	for _, typ := range committed.Types {
		for _, a := range typ.Attributes {
			if a.References != nil {
				carried++
			}
		}
	}
	if carried != usable {
		t.Errorf("committed catalog carries %d references, the lock allows %d", carried, usable)
	}
	if len(cat.Types) != len(committed.Types) {
		t.Errorf("regenerated %d types, committed catalog has %d", len(cat.Types), len(committed.Types))
	}
}
