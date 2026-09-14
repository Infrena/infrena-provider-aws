package awsprov

import (
	"context"
	"testing"
	"time"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// fakeProvider configures a real instance — LoadDefaultConfig and all — against a fresh fake, with
// the eventual-consistency backoff made instant. Its attempt count is kept, so tests can count calls.
func fakeProvider(t *testing.T, values map[string]value.Value) (*Provider, *ec2fake.Server) {
	t.Helper()
	fake := ec2fake.New()
	t.Cleanup(fake.Close)
	isolateAWS(t, fake.URL)
	prov, err := NewPlugin().New(provider.Config{Instance: "test", Values: values})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	p := prov.(*Provider)
	p.patience.sleep = func(context.Context, time.Duration) error { return nil }
	return p, fake
}

func desired(typ string, attrs map[string]value.Value) *resource.DesiredResource {
	return &resource.DesiredResource{Address: address.Address{Name: "r"}, Type: typ, Attrs: attrs}
}

func tagsValue(kv ...string) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i]] = value.String(kv[i+1], value.SourceExplicit)
	}
	return value.Map(items, value.SourceExplicit)
}

func vpcAttrs(extra map[string]value.Value) map[string]value.Value {
	attrs := map[string]value.Value{"region": s("us-east-1"), "cidr": s("10.0.0.0/16")}
	for k, v := range extra {
		attrs[k] = v
	}
	return attrs
}

func tagsOf(t *testing.T, st *resource.ResourceState) map[string]string {
	t.Helper()
	got, err := tagsFrom(st.Attributes)
	if err != nil {
		t.Fatal(err)
	}
	return got
}
