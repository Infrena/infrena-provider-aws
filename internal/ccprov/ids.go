// Package ccprov is the generic provider: one implementation serving every catalog type through AWS Cloud Control API.
package ccprov

import (
	"fmt"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
)

// GlobalRegion is where Cloud Control is called for a type that has no region (IAM, Route 53, CloudFront).
const GlobalRegion = "us-east-1"

// FormatID is the one provider-ID form (P3): `<region>/<identifier>`, or `global/<identifier>`. Create, Discover and
// Import all use it, because infrata imports by matching `<type>.<provider id>` against what Discover returned.
func FormatID(t *catalog.Type, region, identifier string) string {
	if t.Global() {
		return catalog.GlobalScope + "/" + identifier
	}
	return region + "/" + identifier
}

// ParseID refuses an ID that cannot be this type's before any API call. It splits at the first slash only: an
// identifier may be an ARN.
func ParseID(t *catalog.Type, providerID string) (region, identifier string, err error) {
	shape := "<" + strings.Join(t.Identifier, "|") + ">"
	example := "us-east-1/" + shape
	if t.Global() {
		example = catalog.GlobalScope + "/" + shape
	}
	scope, id, ok := strings.Cut(providerID, "/")
	if !ok || scope == "" || id == "" {
		return "", "", fmt.Errorf("%q is not a %s ID: expected %s", providerID, t.Name, example)
	}
	if parts := strings.Count(id, "|") + 1; parts != len(t.Identifier) {
		return "", "", fmt.Errorf("%q is not a %s ID: expected %s, %d part(s) separated by |", providerID, t.Name, example, len(t.Identifier))
	}
	switch {
	case t.Global() && scope != catalog.GlobalScope:
		return "", "", fmt.Errorf("%q is not a %s ID: %s is not regional, so its ID is %s", providerID, t.Name, t.Name, example)
	case t.Global():
		return GlobalRegion, id, nil
	case scope == catalog.GlobalScope:
		return "", "", fmt.Errorf("%q is not a %s ID: %s is regional, so its ID starts with the region: %s", providerID, t.Name, t.Name, example)
	}
	return scope, id, nil
}
