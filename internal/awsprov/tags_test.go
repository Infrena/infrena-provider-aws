package awsprov

import (
	"testing"

	"github.com/infrata/infrata/pkg/value"
)

func TestTagsMustBeAMapOfStrings(t *testing.T) {
	for name, v := range map[string]value.Value{
		"not a map":         s("team=platform"),
		"non-string value":  value.Map(map[string]value.Value{"n": value.Int(1, value.SourceExplicit)}, value.SourceExplicit),
		"reserved aws: key": tagsValue("aws:x", "y"),
	} {
		if _, err := tagsFrom(map[string]value.Value{"tags": v}); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}
