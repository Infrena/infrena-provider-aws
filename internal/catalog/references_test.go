package catalog

import (
	"strings"
	"testing"
)

// TestEveryEmbeddedReferenceNamesADeclaredTypeAndCanonicalAttribute: infrena refuses at load a reference whose type
// the plugin does not declare or whose attribute that type lacks, and it wants the canonical name, never an alias.
// internal/awsprov's TestTheWholeCatalogPassesValidateAll runs infrena's own check on the definitions as well.
func TestEveryEmbeddedReferenceNamesADeclaredTypeAndCanonicalAttribute(t *testing.T) {
	c, err := Embedded()
	if err != nil {
		t.Fatal(err)
	}
	carried := 0
	for _, typ := range c.Types {
		for _, a := range typ.Attributes {
			r := a.References
			if r == nil {
				continue
			}
			carried++
			target, ok := c.Lookup(r.Type)
			if !ok {
				t.Errorf("%s.%s refers to %s, which the catalog does not declare", typ.Name, a.Name, r.Type)
				continue
			}
			if _, ok := target.Attribute(r.Attribute); !ok {
				t.Errorf("%s.%s refers to %s.%s, which is not a canonical attribute of that type", typ.Name, a.Name, r.Type, r.Attribute)
			}
			if strings.Contains(a.Name, "/") {
				t.Errorf("%s.%s: a reference on a nested attribute", typ.Name, a.Name)
			}
		}
	}
	if carried == 0 {
		t.Fatal("the embedded catalog carries no references; the check above checked nothing")
	}
	vpc, _ := c.Lookup("aws.subnet")
	if a, ok := vpc.Attribute("VpcId"); !ok || a.References == nil || a.References.Type != "aws.vpc" || a.References.Attribute != "VpcId" {
		t.Errorf("aws.subnet VpcId = %+v, want a reference to aws.vpc's VpcId", a)
	}
}
