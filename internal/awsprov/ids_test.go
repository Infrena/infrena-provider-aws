package awsprov

import (
	"strings"
	"testing"
)

func TestProviderIDsRoundTrip(t *testing.T) {
	region, id, err := parseID(typeVPC, formatID("eu-west-1", "vpc-0abc123"))
	if err != nil || region != "eu-west-1" || id != "vpc-0abc123" {
		t.Fatalf("parseID = %q, %q, %v", region, id, err)
	}
}

// TestAnIDOfTheWrongTypeIsRefused. Importing a VPC as a subnet would record a VPC as a subnet in
// state, and the next plan would propose replacing real infrastructure to settle it.
func TestAnIDOfTheWrongTypeIsRefused(t *testing.T) {
	for _, c := range []struct{ typ, id, want string }{
		{typeSubnet, "us-east-1/vpc-0abc123", "subnet-"},
		{typeVPC, "us-east-1/subnet-0abc123", "vpc-"},
		{typeVPC, "vpc-0abc123", "<region>/"},
		{typeVPC, "us-east-1/", "<region>/"},
		{typeVPC, "/vpc-0abc123", "<region>/"},
		{typeVPC, "us-east-1/vpc-1/extra", "<region>/"},
	} {
		_, _, err := parseID(c.typ, c.id)
		if err == nil || !strings.Contains(err.Error(), c.want) || !strings.Contains(err.Error(), c.id) {
			t.Errorf("parseID(%s, %q) = %v, want an error naming the ID and %q", c.typ, c.id, err, c.want)
		}
	}
}
