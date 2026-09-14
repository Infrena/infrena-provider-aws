package awsprov

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/value"
)

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }

func list(items ...string) value.Value {
	vs := make([]value.Value, 0, len(items))
	for _, it := range items {
		vs = append(vs, s(it))
	}
	return value.List(vs, value.SourceExplicit)
}

// TestAnUnknownKeyIsRefusedNamingWhatIsAccepted. A misspelled `profil:` silently ignored falls back
// to the default credential chain, which may be another account.
func TestAnUnknownKeyIsRefusedNamingWhatIsAccepted(t *testing.T) {
	_, err := parseConfig(map[string]value.Value{"profil": s("prod"), "region": s("us-east-1")})
	if err == nil {
		t.Fatal("unknown keys were accepted")
	}
	for _, want := range []string{`"profil"`, `"region"`, "assume_role_arn", "discover_regions", "profile", "defaults: {region"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %s", err, want)
		}
	}
}

func TestConfigValuesAreRead(t *testing.T) {
	ic, err := parseConfig(map[string]value.Value{
		"profile":          s("prod"),
		"assume_role_arn":  s("arn:aws:iam::123456789012:role/deploy"),
		"discover_regions": list("us-east-1", "eu-west-1", "us-east-1"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if ic.Profile != "prod" || ic.AssumeRoleARN != "arn:aws:iam::123456789012:role/deploy" {
		t.Errorf("config = %+v", ic)
	}
	if strings.Join(ic.DiscoverRegions, ",") != "us-east-1,eu-west-1" {
		t.Errorf("discover_regions = %v, want duplicates removed in written order", ic.DiscoverRegions)
	}
}

func TestMalformedValuesAreRefused(t *testing.T) {
	for name, values := range map[string]map[string]value.Value{
		"profile not a string":           {"profile": list("a")},
		"empty profile":                  {"profile": s("")},
		"discover_regions not a list":    {"discover_regions": s("us-east-1")},
		"discover_regions holds a blank": {"discover_regions": list("us-east-1", "")},
		"assume_role_arn not an ARN":     {"assume_role_arn": s("deploy")},
	} {
		if _, err := parseConfig(values); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

// TestANamedProfileThatDoesNotExistFailsInNew. A file read, no network: cheap enough for every
// validate, and the one credential mistake worth catching before an apply starts.
func TestANamedProfileThatDoesNotExistFailsInNew(t *testing.T) {
	isolateAWS(t, "")
	_, err := NewPlugin().New(provider.Config{Instance: "prod", Values: map[string]value.Value{"profile": s("prodd")}})
	if err == nil {
		t.Fatal("a profile missing from the config files was accepted")
	}
	for _, want := range []string{`"prod"`, `"prodd"`, "aws configure list-profiles"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %s", err, want)
		}
	}
}

// TestAssumeRoleSignsEC2CallsWithTheAssumedCredentials. The fixture's static key would sign every
// call if the role were ignored, so the assertion cannot pass by accident.
func TestAssumeRoleSignsEC2CallsWithTheAssumedCredentials(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	isolateAWS(t, fake.URL)

	prov, err := NewPlugin().New(provider.Config{Instance: "deploy", Values: map[string]value.Value{
		"assume_role_arn": s("arn:aws:iam::123456789012:role/deploy"),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := prov.(*Provider).clients.ec2("us-east-1").DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{}); err != nil {
		t.Fatal(err)
	}
	if fake.Calls("AssumeRole") != 1 {
		t.Errorf("AssumeRole calls = %d, want 1", fake.Calls("AssumeRole"))
	}
	keys := fake.AccessKeys()
	if last := keys[len(keys)-1]; last != ec2fake.AssumedAccessKey {
		t.Errorf("DescribeVpcs was signed with %q, want the assumed role's %q", last, ec2fake.AssumedAccessKey)
	}
}

func TestAProfileThatExistsConfigures(t *testing.T) {
	dir := isolateAWS(t, "")
	if err := os.WriteFile(filepath.Join(dir, "config"), []byte("[profile staging]\nregion = eu-west-1\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := NewPlugin().New(provider.Config{Instance: "staging", Values: map[string]value.Value{"profile": s("staging")}}); err != nil {
		t.Fatalf("New: %v", err)
	}
}
