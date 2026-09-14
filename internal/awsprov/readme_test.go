package awsprov

import (
	"os"
	"strings"
	"testing"
)

// TestReadmeQuotesTheTestedExample. The README's infra.yml is the first thing anyone copies; the e2e suite runs that
// exact file, so the README must quote it byte for byte.
func TestReadmeQuotesTheTestedExample(t *testing.T) {
	readme, err := os.ReadFile("../../README.md")
	if err != nil {
		t.Fatal(err)
	}
	fixture, err := os.ReadFile("../../e2e/testdata/basic/infra.yml")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(readme), string(fixture)) {
		t.Error("README.md does not quote e2e/testdata/basic/infra.yml verbatim")
	}
	for _, want := range []string{
		"Cloud Control", "infrena explain", "discover_regions", "discover_types", "assume_role_arn", "profile", "--provider",
		"defaults:", "${var.aws_region}", "us-east-1/vpc-", "global/", "aws_region", "type_value", "cidr_block", "tags:",
		"gen/overlay.yaml", "gen/names.lock.json", "scripts/fetch-schemas", "go run ./cmd/gen-cloudcontrol",
		"-tags e2e", "-tags live", "go work init", "GOWORK=off", "plugin.yaml", "scripts/release-check", "0.0.0-dev",
	} {
		if !strings.Contains(string(readme), want) {
			t.Errorf("README.md never mentions %q", want)
		}
	}
}
