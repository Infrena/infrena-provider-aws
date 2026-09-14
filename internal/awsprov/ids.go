package awsprov

import (
	"fmt"
	"strings"
)

// formatID is the one provider-ID form used everywhere — Create, Discover, Import — because infrata
// imports by matching `<type>.<provider id>` against what Discover returned, and Import is given no
// region of its own.
func formatID(region, awsID string) string { return region + "/" + awsID }

var idPrefixes = map[string]string{typeVPC: "vpc-", typeSubnet: "subnet-"}

// parseID refuses a malformed ID or one naming a different kind of resource before any API call.
func parseID(resourceType, providerID string) (string, string, error) {
	prefix := idPrefixes[resourceType]
	region, awsID, ok := strings.Cut(providerID, "/")
	if !ok || region == "" || awsID == "" || strings.Contains(awsID, "/") {
		return "", "", fmt.Errorf("%q is not a %s ID: expected <region>/<id>, e.g. us-east-1/%s0abc123",
			providerID, resourceType, prefix)
	}
	if !strings.HasPrefix(awsID, prefix) {
		return "", "", fmt.Errorf("%q is not a %s: a %s ID starts with %q", providerID, resourceType, resourceType, prefix)
	}
	return region, awsID, nil
}
