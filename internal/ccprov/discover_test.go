package ccprov

import (
	"context"
	"errors"
	"sort"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/provider"
)

func everything() provider.DiscoverRequest {
	var names []string
	for _, typ := range testCatalog().Types {
		names = append(names, typ.Name)
	}
	return provider.DiscoverRequest{Types: names}
}

func ids(found []provider.DiscoveredResource) []string {
	out := make([]string, len(found))
	for i, r := range found {
		out[i] = r.Type + " " + r.ProviderID
	}
	sort.Strings(out)
	return out
}

func TestDiscoveringEverythingMeansTheDefaultSetInEveryRegion(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	p.opts.DiscoverRegions = []string{"us-east-1", "eu-west-1"}
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "10.0.0.0/16"})
	fake.Put("eu-west-1", "AWS::EC2::VPC", "vpc-2", map[string]any{"VpcId": "vpc-2", "CidrBlock": "10.1.0.0/16"})
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})

	found, err := p.Discover(ctx, everything())
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"aws.role global/deploy", "aws.vpc eu-west-1/vpc-2", "aws.vpc us-east-1/vpc-1"}
	if got := ids(found); strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("found %v, want %v (the bucket is outside the default set)", got, want)
	}
	if n := fake.Calls("ListResources"); n != 3 {
		t.Errorf("ListResources calls = %d, want 3: VPCs in two regions, roles once", n)
	}
	for _, r := range found {
		if r.Type == "aws.vpc" && r.Attributes["region"].Raw == nil {
			t.Errorf("%s has no region attribute: import --generate would write a resource that does not plan", r.ProviderID)
		}
	}
}

func TestDiscoverTypesReplacesTheDefaultSet(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	p.opts.DiscoverTypes = []string{"aws.bucket"}
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1"})
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})
	found, err := p.Discover(ctx, everything())
	if err != nil || strings.Join(ids(found), ",") != "aws.bucket us-east-1/logs" {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

func TestARequestNamingSomeTypesIsHonouredAsGiven(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.bucket"}})
	if err != nil || len(found) != 1 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

func TestTypesDiscoveryCannotListAreSkippedAndNamed(t *testing.T) {
	p, fake, log := fakeProvider(t)
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.test.child", "aws.test.regioned", "aws.vpc"}}); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"aws.test.child", "parent", "aws.test.regioned"} {
		if !strings.Contains(log.String(), want) {
			t.Errorf("stderr lacks %q:\n%s", want, log)
		}
	}
	if n := fake.Calls("ListResources"); n != 1 {
		t.Errorf("ListResources calls = %d, want 1 (aws.vpc only)", n)
	}
}

func TestEveryPageIsFollowed(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.PageSize = 1
	for _, id := range []string{"vpc-1", "vpc-2", "vpc-3"} {
		fake.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
	}
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil || len(found) != 3 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

// TestDiscoveredValuesAreWhatImportReports: import --generate writes discovery's values, then imports; the two must
// agree or the next plan is not clean.
func TestDiscoveredValuesAreWhatImportReports(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::SecurityGroup", "sg-1", map[string]any{
		"GroupId": "sg-1", "GroupDescription": "web",
		"SecurityGroupIngress": []any{map[string]any{"IpProtocol": "tcp", "FromPort": 443, "ToPort": 443, "CidrIp": "0.0.0.0/0"}},
	})
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.securitygroup"}})
	if err != nil || len(found) != 1 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
	ingress := found[0].Attributes["SecurityGroupIngress"]
	if !ingress.Equal(list(rule("tcp", 443, 443, snake))) {
		t.Errorf("discovered ingress = %v, want snake_case keys", ingress)
	}
	imported, err := p.Import(ctx, "aws.securitygroup", found[0].ProviderID)
	if err != nil {
		t.Fatal(err)
	}
	for name, v := range found[0].Attributes {
		if !imported.Attributes[name].Equal(v) {
			t.Errorf("%s: discovered %v, imported %v", name, v, imported.Attributes[name])
		}
	}
}

func TestOneFailingTypeDoesNotHideTheRest(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	fake.Inject(ccfake.Fault{Action: "ListResources", Nth: 1, Status: 400, Code: "AccessDeniedException", Message: "no ec2:DescribeVpcs"})
	found, err := p.Discover(ctx, everything())
	if err != nil || strings.Join(ids(found), ",") != "aws.role global/deploy" {
		t.Fatalf("found %v, %v", ids(found), err)
	}
	if !strings.Contains(log.String(), "AccessDeniedException") || !strings.Contains(log.String(), "aws.vpc") {
		t.Errorf("stderr does not report the failed type:\n%s", log)
	}
}

func TestWhenEveryAttemptFailsDiscoverSaysSo(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	for n := 1; n <= 2; n++ {
		fake.Inject(ccfake.Fault{Action: "ListResources", Nth: n, Status: 400, Code: "AccessDeniedException", Message: "denied"})
	}
	if _, err := p.Discover(ctx, everything()); err == nil || !strings.Contains(err.Error(), "AccessDeniedException") {
		t.Fatalf("err = %v", err)
	}
}

func TestWithoutDiscoverRegionsOnlyGlobalTypesCanBeDiscovered(t *testing.T) {
	p, _, _ := fakeProvider(t)
	p.opts.DiscoverRegions = nil
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}}); err == nil || !strings.Contains(err.Error(), "discover_regions") {
		t.Errorf("err = %v", err)
	}
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.role"}}); err != nil {
		t.Errorf("a global type needs no region: %v", err)
	}
}

func TestCancellationStopsDiscovery(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := p.Discover(cctx, everything()); !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
	if n := fake.Calls("ListResources"); n != 0 {
		t.Errorf("ListResources calls = %d after cancellation", n)
	}
}
