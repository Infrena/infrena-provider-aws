package ccprov

import (
	"bytes"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/address"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

var fakePrefixes = map[string]string{
	"aws.vpc": "vpc-", "aws.subnet": "subnet-", "aws.securitygroup": "sg-", "aws.bucket": "bucket-", "aws.role": "role-",
	"aws.dbinstance": "db-", "aws.test.regioned": "reg-", "aws.test.child": "child-",
}

// fakeProvider is a real Provider over the test catalog, talking to a fresh fake through the real SDK, with every wait
// made instant. What it writes to stderr is captured.
func fakeProvider(t *testing.T) (*Provider, *ccfake.Server, *bytes.Buffer) {
	t.Helper()
	fake := ccfake.New()
	t.Cleanup(fake.Close)
	cat := testCatalog()
	defaults := map[string]map[string]any{
		"aws.vpc":  {"EnableDnsSupport": true, "InstanceTenancy": "default"},
		"aws.role": {"MaxSessionDuration": 3600},
	}
	for _, typ := range cat.Types {
		fake.Register(awstest.FakeType(typ, fakePrefixes[typ.Name], defaults[typ.Name]))
	}
	p := New("test", cat, testConfig(fake.URL, 1), Options{DiscoverRegions: []string{"us-east-1"}})
	p.patience, p.pacing = instantPatience, instantPacing
	var log bytes.Buffer
	p.log = &log
	return p, fake, &log
}

func sv(v string) value.Value { return value.String(v, value.SourceExplicit) }
func iv(n int64) value.Value  { return value.Int(n, value.SourceExplicit) }

func desired(typ string, attrs map[string]value.Value) *resource.DesiredResource {
	return &resource.DesiredResource{Address: address.Address{Name: "r"}, Type: typ, Attrs: attrs}
}
