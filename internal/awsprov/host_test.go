package awsprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/address"
	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

// hosted is the plugin reached through infrena's host, configured against a fresh fake.
func hosted(t *testing.T, values map[string]value.Value) (provider.Provider, *ccfake.Server, *catalog.Catalog) {
	t.Helper()
	fake := ccfake.New()
	t.Cleanup(fake.Close)
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	awstest.RegisterCore(t, fake, cat)
	awstest.Isolate(t, fake.URL)
	prov, err := openHost(t).Configure(provider.Config{Instance: "main", Values: values})
	if err != nil {
		t.Fatal(err)
	}
	return prov, fake, cat
}

func want(t *testing.T, desired map[string]value.Value, st *resource.ResourceState) {
	t.Helper()
	for name, v := range desired {
		if got, ok := st.Attributes[name]; !ok || !got.Equal(v) {
			t.Errorf("%s = %v, want %v: the next plan would not be clean", name, got, v)
		}
	}
}

func m(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}

func TestTheLifecycleThroughTheHost(t *testing.T) {
	prov, fake, _ := hosted(t, nil)
	ctx := context.Background()
	attrs := map[string]value.Value{"region": s("us-east-1"), "CidrBlock": s("10.0.0.0/16"), "Tags": m("team", s("platform"))}

	st, err := prov.Create(ctx, &resource.DesiredResource{Address: address.Address{Name: "vpc"}, Type: "aws.vpc", Attrs: attrs})
	if err != nil {
		t.Fatal(err)
	}
	want(t, attrs, st)
	if v := st.Attributes["EnableDnsSupport"]; v.Raw != true || v.Source != value.SourceProvider {
		t.Errorf("EnableDnsSupport = %+v, want AWS's true, recorded as the provider's", v)
	}

	got, err := prov.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	want(t, attrs, got)

	changed := map[string]value.Value{}
	for k, v := range got.Attributes {
		changed[k] = v
	}
	changed["EnableDnsHostnames"] = value.Bool(true, value.SourceExplicit)
	updated, err := prov.Update(ctx, got, &resource.DesiredResource{Address: address.Address{Name: "vpc"}, Type: "aws.vpc", Attrs: changed})
	if err != nil || updated.Attributes["EnableDnsHostnames"].Raw != true {
		t.Fatalf("Update = %v, %v", updated, err)
	}

	if err := prov.Delete(ctx, updated); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("us-east-1", "AWS::EC2::VPC")); n != 0 {
		t.Errorf("%d VPCs remain", n)
	}
}

// TestASecretIsSensitiveWhereverItAppears. The host forces the flag from the schema; this proves the overlay's
// sensitive list reached the schema, and that the write-only value survives the wire into state.
func TestASecretIsSensitiveWhereverItAppears(t *testing.T) {
	prov, _, cat := hosted(t, nil)
	db := awstest.TypeFor(t, cat, "AWS::RDS::DBInstance")
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "db"}, Type: db.Name,
		Attrs: map[string]value.Value{"region": s("us-east-1"), "DBInstanceIdentifier": s("app"), "DBInstanceClass": s("db.t3.micro"),
			"MasterUsername": s("app"), "MasterUserPassword": s("hunter2")}})
	if err != nil {
		t.Fatal(err)
	}
	pw := st.Attributes["MasterUserPassword"]
	if pw.Raw != "hunter2" || !pw.Sensitive {
		t.Errorf("MasterUserPassword = %+v, want the carried value, marked sensitive", pw)
	}
}

func TestNestedSpellingSurvivesTheWire(t *testing.T) {
	prov, _, _ := hosted(t, nil)
	ingress := value.List([]value.Value{m("ip_protocol", s("tcp"), "from_port", value.Int(443, value.SourceExplicit),
		"to_port", value.Int(443, value.SourceExplicit), "cidr_ip", s("0.0.0.0/0"))}, value.SourceExplicit)
	attrs := map[string]value.Value{"region": s("us-east-1"), "GroupDescription": s("web"), "SecurityGroupIngress": ingress}
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "web"}, Type: "aws.securitygroup", Attrs: attrs})
	if err != nil {
		t.Fatal(err)
	}
	want(t, attrs, st)
}

func TestDiscoverAndImportThroughTheHost(t *testing.T) {
	prov, fake, cat := hosted(t, map[string]value.Value{"discover_regions": list("us-east-1")})
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-legacy", map[string]any{"VpcId": "vpc-legacy", "CidrBlock": "172.16.0.0/16"})
	role := awstest.TypeFor(t, cat, "AWS::IAM::Role")
	fake.Put("us-east-1", "AWS::IAM::Role", "ops", map[string]any{"RoleName": "ops", "Arn": "arn:aws:iam::123456789012:role/ops"})

	var everything []string
	for _, d := range openHost(t).Definitions() {
		everything = append(everything, d.Type)
	}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, r := range found {
		got = append(got, r.Type+" "+r.ProviderID)
	}
	for _, w := range []string{"aws.vpc us-east-1/vpc-legacy", role.Name + " global/ops"} {
		if !strings.Contains(strings.Join(got, "\n"), w) {
			t.Errorf("discovered %v, want %s", got, w)
		}
	}
	st, err := prov.Import(context.Background(), "aws.vpc", "us-east-1/vpc-legacy")
	if err != nil || st.Attributes["CidrBlock"].Raw != "172.16.0.0/16" {
		t.Fatalf("Import = %v, %v", st, err)
	}
}

// TestAnUnknownNestedKeyIsAnErrorBeforeAnyCall. The plan cannot see nested keys; the apply must say which one.
func TestAnUnknownNestedKeyIsAnErrorBeforeAnyCall(t *testing.T) {
	prov, fake, _ := hosted(t, nil)
	bad := value.List([]value.Value{m("ip_protocol", s("tcp"), "port", value.Int(443, value.SourceExplicit))}, value.SourceExplicit)
	_, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "web"}, Type: "aws.securitygroup",
		Attrs: map[string]value.Value{"region": s("us-east-1"), "GroupDescription": s("web"), "SecurityGroupIngress": bad}})
	if err == nil || !strings.Contains(err.Error(), `"port"`) || !strings.Contains(err.Error(), "from_port") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("CreateResource") != 0 {
		t.Error("CreateResource was called")
	}
}
