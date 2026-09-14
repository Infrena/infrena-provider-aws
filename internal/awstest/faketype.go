package awstest

import (
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
)

// FakeType describes a catalog type to the Cloud Control fake the way its schema does, so the fake cannot drift from
// the catalog: create-only properties are the ones the catalog marks force-new on an updatable type, write-only ones
// are the catalog's, and every read-only string property gets a value derived from the identifier.
func FakeType(t *catalog.Type, idPrefix string, defaults map[string]any) ccfake.TypeConfig {
	tc := ccfake.TypeConfig{
		TypeName: t.CFN, Identifier: t.Identifier[0], IDPrefix: idPrefix,
		Defaults: defaults, WriteOnly: t.WriteOnly, ReadOnly: map[string]string{},
	}
	for _, a := range t.Attributes {
		if a.Computed && !a.Optional && a.Name != tc.Identifier && a.Kind == "string" {
			tc.ReadOnly[a.Name] = strings.ToLower(a.Name) + "-{id}"
		}
		if a.ForceNew && t.HasUpdate {
			tc.CreateOnly = append(tc.CreateOnly, a.Name)
		}
	}
	return tc
}
