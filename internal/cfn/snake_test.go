package cfn

import "testing"

// TestSnakeCase pins the rule the bundle analysis used: an underscore before an upper-case letter that follows a
// lower-case letter or digit, or that follows an upper-case letter and precedes a lower-case one.
func TestSnakeCase(t *testing.T) {
	for in, want := range map[string]string{
		"CidrBlock":            "cidr_block",
		"VpcId":                "vpc_id",
		"EnableDnsHostnames":   "enable_dns_hostnames",
		"DBInstanceIdentifier": "db_instance_identifier",
		"Ipv6CidrBlocks":       "ipv6_cidr_blocks",
		"SSESpecification":     "sse_specification",
		"S3Bucket":             "s3_bucket",
		"ARN":                  "arn",
		"Type":                 "type",
		"already_snake":        "already_snake",
	} {
		if got := SnakeCase(in); got != want {
			t.Errorf("SnakeCase(%q) = %q, want %q", in, got, want)
		}
	}
}
