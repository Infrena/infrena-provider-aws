package ccprov

import (
	"strings"
	"testing"
)

func TestProviderIDsRoundTrip(t *testing.T) {
	cases := []struct {
		typ                    string
		region, identifier, id string
	}{
		{"aws.vpc", "eu-west-1", "vpc-0abc", "eu-west-1/vpc-0abc"},
		{"aws.role", GlobalRegion, "deploy", "global/deploy"},
		// Identifiers may be ARNs, which hold slashes: the split is at the first slash only.
		{"aws.vpc", "us-east-1", "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-1", "us-east-1/arn:aws:ec2:us-east-1:123456789012:vpc/vpc-1"},
		{"aws.test.child", "us-east-1", "parent-1|child-2", "us-east-1/parent-1|child-2"},
	}
	for _, c := range cases {
		typ := mustType(t, c.typ)
		if got := FormatID(typ, c.region, c.identifier); got != c.id {
			t.Errorf("FormatID(%s, %s, %s) = %q, want %q", c.typ, c.region, c.identifier, got, c.id)
		}
		region, identifier, err := ParseID(typ, c.id)
		if err != nil || region != c.region || identifier != c.identifier {
			t.Errorf("ParseID(%s, %q) = %q, %q, %v", c.typ, c.id, region, identifier, err)
		}
	}
}

// TestMalformedIDsAreRefusedBeforeAnyCall. An import typo must say what the ID should look like.
func TestMalformedIDsAreRefusedBeforeAnyCall(t *testing.T) {
	cases := []struct{ typ, id, want string }{
		{"aws.vpc", "vpc-0abc", "us-east-1/<VpcId>"},
		{"aws.vpc", "/vpc-0abc", "us-east-1/<VpcId>"},
		{"aws.vpc", "us-east-1/", "us-east-1/<VpcId>"},
		{"aws.vpc", "global/vpc-0abc", "regional"},
		{"aws.role", "us-east-1/deploy", "global/<RoleName>"},
		{"aws.test.child", "us-east-1/child-2", "<ParentId|ChildId>"},
	}
	for _, c := range cases {
		if _, _, err := ParseID(mustType(t, c.typ), c.id); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("ParseID(%s, %q) = %v, want an error mentioning %q", c.typ, c.id, err, c.want)
		}
	}
}
