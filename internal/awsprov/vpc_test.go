package awsprov

import (
	"context"
	"maps"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func TestAVPCIsCreatedWithItsTagsAndReadBack(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("team", "platform")})))
	if err != nil {
		t.Fatal(err)
	}
	vpcs := fake.VPCs()
	if len(vpcs) != 1 || vpcs[0].Tags["team"] != "platform" {
		t.Fatalf("fake holds %+v", vpcs)
	}
	if st.ProviderID != "us-east-1/"+vpcs[0].ID || st.Type != typeVPC {
		t.Errorf("state = %s %s", st.Type, st.ProviderID)
	}
	if id, _ := st.Attributes["id"].AsString(); id != vpcs[0].ID {
		t.Errorf("id = %q", id)
	}
	if owner, _ := st.Attributes["owner_id"].AsString(); owner != ec2fake.Owner {
		t.Errorf("owner_id = %q", owner)
	}
	if st.Address.Name != "" || st.Provider != "" {
		t.Error("bookkeeping is the host's: the plugin must leave Address and Provider unset")
	}
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	if !maps.Equal(tagsOf(t, got), map[string]string{"team": "platform"}) {
		t.Errorf("read tags = %v", tagsOf(t, got))
	}
	if cidr, _ := got.Attributes["cidr"].AsString(); cidr != "10.0.0.0/16" {
		t.Errorf("cidr = %q", cidr)
	}
}

// TestACreateIsSentOnceEvenWhenTheAnswerIsLost. AWS_MAX_ATTEMPTS=5 makes the SDK want to retry: the
// assertion is that it does not, and that the one VPC that exists is exactly one.
func TestACreateIsSentOnceEvenWhenTheAnswerIsLost(t *testing.T) {
	for name, fault := range map[string]ec2fake.Fault{
		"connection dropped after acting": {Drop: true},
		"server fault":                    {Status: 500, Code: "InternalError", Message: "oops"},
	} {
		t.Run(name, func(t *testing.T) {
			fake := ec2fake.New()
			defer fake.Close()
			isolateAWS(t, fake.URL)
			t.Setenv("AWS_MAX_ATTEMPTS", "5")
			prov, err := NewPlugin().New(provider.Config{Instance: "test"})
			if err != nil {
				t.Fatal(err)
			}
			fault.Action, fault.Nth = "CreateVpc", 1
			fake.Inject(fault)
			_, err = prov.Create(context.Background(), desired(typeVPC, vpcAttrs(nil)))
			if err == nil {
				t.Fatal("expected the create to fail")
			}
			if n := fake.Calls("CreateVpc"); n != 1 {
				t.Errorf("CreateVpc sent %d times; a resent create makes an untracked VPC", n)
			}
			if got := prov.ClassifyError(err); got != provider.ConditionallyRetryable {
				t.Errorf("classification = %v, want ConditionallyRetryable so infrata does not retry the create", got)
			}
		})
	}
}

func TestAReadWaitsForAVPCEC2HasNotPropagated(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.HideFromDescribe(fake.VPCs()[0].ID, 3)
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read of a just-created VPC = %v, %v; reporting it gone would plan a second VPC", got, err)
	}
	if n := fake.Calls("DescribeVpcs"); n != 4 {
		t.Errorf("DescribeVpcs calls = %d, want 4 (three hidden, then found)", n)
	}
}

func TestAReadReportsAVPCGoneOnlyAfterItsPatienceRunsOut(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.Remove(fake.VPCs()[0].ID)
	got, err := p.Read(ctx, st)
	if err != nil || got != nil {
		t.Fatalf("Read of a deleted VPC = %v, %v; want (nil, nil)", got, err)
	}
	if n := fake.Calls("DescribeVpcs"); n != notFoundPatience.attempts {
		t.Errorf("DescribeVpcs calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

// TestUpdateRemovesTagsConfigurationDropped. Merging instead of replacing leaves `a` in place and
// every later plan proposes removing it again.
func TestUpdateRemovesTagsConfigurationDropped(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("a", "1", "b", "2")})))
	if err != nil {
		t.Fatal(err)
	}
	st, err = p.Update(ctx, st, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("b", "3", "c", "4")})))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"b": "3", "c": "4"}
	if got := fake.VPCs()[0].Tags; !maps.Equal(got, want) {
		t.Errorf("AWS tags = %v, want %v", got, want)
	}
	if got := tagsOf(t, st); !maps.Equal(got, want) {
		t.Errorf("returned tags = %v, want %v", got, want)
	}
	if _, ok := st.Attributes["id"]; !ok {
		t.Error("Update dropped a computed attribute: desired never carries them")
	}
	st, err = p.Update(ctx, st, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := st.Attributes["tags"]; ok || len(fake.VPCs()[0].Tags) != 0 {
		t.Errorf("after removing tags: state tags present=%v, AWS tags=%v", ok, fake.VPCs()[0].Tags)
	}
}

// TestAWSReservedTagsAreNeitherReportedNorAccepted. Reporting them plans their removal forever;
// AWS refuses to remove them.
func TestAWSReservedTagsAreNeitherReportedNorAccepted(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.SetTag(fake.VPCs()[0].ID, "aws:cloudformation:stack-name", "legacy")
	got, err := p.Read(ctx, st)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := got.Attributes["tags"]; ok {
		t.Errorf("reserved tags were reported: %v", got.Attributes["tags"])
	}
	_, err = p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("aws:owner", "me")})))
	if err == nil || !strings.Contains(err.Error(), "aws:owner") {
		t.Fatalf("create with a reserved tag = %v", err)
	}
	if n := fake.Calls("CreateVpc"); n != 1 {
		t.Errorf("CreateVpc calls = %d: a reserved tag must be refused before any call", n)
	}
}

func TestDeletingAVPCThatIsAlreadyGoneSucceeds(t *testing.T) {
	p, _ := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := p.Delete(ctx, st); err != nil {
			t.Fatalf("delete %d: %v", i+1, err)
		}
	}
}

func TestACancelledCreateSendsNothing(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil))); err == nil {
		t.Fatal("a cancelled create succeeded")
	}
	if n := fake.Calls("CreateVpc"); n != 0 {
		t.Errorf("CreateVpc calls = %d, want 0", n)
	}
}

func TestReadRefusesAnIDOfAnotherType(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	_, err := p.Read(context.Background(), &resource.ResourceState{Type: typeVPC, ProviderID: "us-east-1/subnet-1"})
	if err == nil || fake.Calls("DescribeVpcs") != 0 {
		t.Fatalf("Read = %v after %d calls", err, fake.Calls("DescribeVpcs"))
	}
}
