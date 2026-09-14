//go:build live

// Package live runs the plugin against a real AWS account. It is the only suite that sees real Cloud Control handlers,
// eventual consistency and IAM, and the only one that costs anything if it leaks, so it refuses to run unless told
// which account it may use and the credentials really are that account.
//
//	INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 go test -tags live -count=1 -v -timeout 30m ./live/
package live

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/infrena/infrena-provider-aws/internal/awsprov"
	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/address"
	"github.com/infrena/infrena/pkg/plugintest"
	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

const runTag = "infrena-live-run"

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }
func n(v int64) value.Value  { return value.Int(v, value.SourceExplicit) }
func m(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}
func l(items ...value.Value) value.Value { return value.List(items, value.SourceExplicit) }

// guard skips without the variables, and refuses when the profile is not the named account.
func guard(t *testing.T) (profile, region string) {
	t.Helper()
	profile, account := os.Getenv("INFRENA_AWS_LIVE_PROFILE"), os.Getenv("INFRENA_AWS_LIVE_ACCOUNT")
	if profile == "" || account == "" {
		t.Skip("set INFRENA_AWS_LIVE_PROFILE and INFRENA_AWS_LIVE_ACCOUNT to run against real AWS")
	}
	region = os.Getenv("INFRENA_AWS_LIVE_REGION")
	if region == "" {
		region = "us-east-1"
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		t.Fatal(err)
	}
	id, err := sts.NewFromConfig(cfg).GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := aws.ToString(id.Account); got != account {
		t.Fatalf("profile %q is account %s, not INFRENA_AWS_LIVE_ACCOUNT=%s: refusing to create anything", profile, got, account)
	}
	if strings.HasSuffix(aws.ToString(id.Arn), ":root") {
		t.Fatalf("profile %q holds root credentials: use a least-privilege IAM identity (live/README.md)", profile)
	}
	return profile, region
}

func configure(t *testing.T, profile, region string) (provider.Provider, *catalog.Catalog) {
	t.Helper()
	host, err := plugintest.Open(context.Background(), awsprov.NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = host.Close() })
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	role := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	prov, err := host.Configure(provider.Config{Instance: "live", Values: map[string]value.Value{
		"profile":          s(profile),
		"discover_regions": l(s(region)),
		"discover_types":   l(s("aws.vpc"), s("aws.subnet"), s("aws.securitygroup"), s(role)),
	}})
	if err != nil {
		t.Fatal(err)
	}
	return prov, cat
}

// converged fails unless every configured value reads back equal: otherwise the next plan is not clean.
func converged(t *testing.T, attrs map[string]value.Value, st *resource.ResourceState) {
	t.Helper()
	for name, v := range attrs {
		if got := st.Attributes[name]; !got.Equal(v) {
			t.Errorf("%s %s: %s reads back as %v, configured %v", st.Type, st.ProviderID, name, got, v)
		}
	}
}

func create(t *testing.T, prov provider.Provider, typ string, attrs map[string]value.Value) *resource.ResourceState {
	t.Helper()
	start := time.Now()
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: typ}, Type: typ, Attrs: attrs})
	if err != nil {
		t.Fatalf("create %s: %v", typ, err)
	}
	t.Logf("created %s %s in %s", typ, st.ProviderID, time.Since(start).Round(time.Millisecond))
	t.Cleanup(func() { _ = prov.Delete(context.Background(), st) })
	converged(t, attrs, st)
	read, err := prov.Read(context.Background(), st)
	if err != nil || read == nil {
		t.Fatalf("read %s straight after create = %v, %v", st.ProviderID, read, err)
	}
	converged(t, attrs, read)
	return read
}

func update(t *testing.T, prov provider.Provider, st *resource.ResourceState, changes map[string]value.Value) *resource.ResourceState {
	t.Helper()
	attrs := map[string]value.Value{}
	for k, v := range st.Attributes {
		attrs[k] = v
	}
	for k, v := range changes {
		attrs[k] = v
	}
	got, err := prov.Update(context.Background(), st, &resource.DesiredResource{Address: address.Address{Name: st.Type}, Type: st.Type, Attrs: attrs})
	if err != nil {
		t.Fatalf("update %s: %v", st.ProviderID, err)
	}
	converged(t, changes, got)
	return got
}

func TestTheLifecycleAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.99.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnet := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.99.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	ingress := l(m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")))
	sg := create(t, prov, "aws.securitygroup", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"GroupDescription": s("infrena live " + run), "SecurityGroupIngress": ingress, "Tags": tags})
	roleType := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	policy := m("Version", s("2012-10-17"), "Statement", l(m("Effect", s("Allow"), "Principal", m("Service", s("ec2.amazonaws.com")), "Action", s("sts:AssumeRole"))))
	role := create(t, prov, roleType, map[string]value.Value{"RoleName": s("infrena-live-" + run), "AssumeRolePolicyDocument": policy, "Tags": tags})

	vpc = update(t, prov, vpc, map[string]value.Value{"EnableDnsHostnames": value.Bool(true, value.SourceExplicit)})
	sg = update(t, prov, sg, map[string]value.Value{"SecurityGroupIngress": l(
		m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")),
		m("ip_protocol", s("tcp"), "from_port", n(80), "to_port", n(80), "cidr_ip", s("0.0.0.0/0")))})
	role = update(t, prov, role, map[string]value.Value{"MaxSessionDuration": n(7200)})

	var everything []string
	for _, typ := range cat.Types {
		everything = append(everything, typ.Name)
	}
	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range found {
		seen[r.ProviderID] = true
	}
	for _, st := range []*resource.ResourceState{vpc, subnet, sg, role} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}
	if imported, err := prov.Import(ctx, "aws.vpc", vpc.ProviderID); err != nil || !imported.Attributes["CidrBlock"].Equal(s("10.99.0.0/16")) {
		t.Errorf("import %s = %v, %v", vpc.ProviderID, imported, err)
	}

	for _, st := range []*resource.ResourceState{role, sg, subnet, vpc} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
}

// TestSweepLeftovers deletes what a crashed run left: anything tagged by this suite more than an hour ago.
func TestSweepLeftovers(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	cutoff := time.Now().Add(-time.Hour).Unix()
	order := []string{awstest.TypeFor(t, cat, "AWS::IAM::Role").Name, "aws.securitygroup", "aws.subnet", "aws.vpc"}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: order})
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range order {
		for _, r := range found {
			if r.Type != typ {
				continue
			}
			tags, _ := r.Attributes["Tags"].Raw.(map[string]value.Value)
			started, err := strconv.ParseInt(fmt.Sprint(tags[runTag].Raw), 10, 64)
			if err != nil || started >= cutoff {
				continue
			}
			t.Logf("deleting %s %s", r.Type, r.ProviderID)
			if err := prov.Delete(context.Background(), &resource.ResourceState{Type: r.Type, ProviderID: r.ProviderID, Attributes: r.Attributes}); err != nil {
				t.Error(err)
			}
		}
	}
}
