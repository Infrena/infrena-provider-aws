package awsprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func subnetAttrs(vpcID string, extra map[string]value.Value) map[string]value.Value {
	attrs := map[string]value.Value{
		"region": s("us-east-1"), "vpc_id": s(vpcID), "cidr": s("10.0.1.0/24"),
		"availability_zone": s("us-east-1a"), "map_public_ip_on_launch": value.Bool(false, value.SourceDefault),
	}
	for k, v := range extra {
		attrs[k] = v
	}
	return attrs
}

func TestASubnetIsCreatedInItsVPCAndReadBack(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	st, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(vpc, map[string]value.Value{"tags": tagsValue("tier", "private")})))
	if err != nil {
		t.Fatal(err)
	}
	subs := fake.Subnets()
	if len(subs) != 1 || subs[0].VPCID != vpc || subs[0].AZ != "us-east-1a" || subs[0].Tags["tier"] != "private" {
		t.Fatalf("fake holds %+v", subs)
	}
	if st.ProviderID != "us-east-1/"+subs[0].ID {
		t.Errorf("provider ID = %q", st.ProviderID)
	}
	if arn, _ := st.Attributes["arn"].AsString(); !strings.HasSuffix(arn, ":subnet/"+subs[0].ID) {
		t.Errorf("arn = %q", arn)
	}
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	for _, name := range []string{"region", "vpc_id", "cidr", "availability_zone", "map_public_ip_on_launch", "tags", "id", "arn", "owner_id"} {
		if !got.Attributes[name].Equal(st.Attributes[name]) {
			t.Errorf("%s: read %v, created %v — a difference here is a change on every plan", name, got.Attributes[name], st.Attributes[name])
		}
	}
}

// TestASubnetWaitsForAVPCCreatedAMomentAgo. CreateSubnet refusing an unpropagated VPC means nothing
// was created, so retrying it inside the plugin is safe — unlike retrying after an ambiguous failure.
func TestASubnetWaitsForAVPCCreatedAMomentAgo(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	fake.HideFromDescribe(vpc, 2) // CreateSubnet consults the same visibility
	if _, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, nil))); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if n := fake.Calls("CreateSubnet"); n != 3 {
		t.Errorf("CreateSubnet calls = %d, want 3 (two refusals, one success)", n)
	}
	if len(fake.Subnets()) != 1 {
		t.Errorf("subnets = %d, want exactly 1", len(fake.Subnets()))
	}
}

func TestASubnetInAVPCThatReallyDoesNotExistFailsNamingIt(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	_, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs("vpc-0000000000000dead", nil)))
	if err == nil || !strings.Contains(err.Error(), "vpc-0000000000000dead") || !strings.Contains(err.Error(), "us-east-1") {
		t.Fatalf("err = %v; want it to name the VPC and the region", err)
	}
	if n := fake.Calls("CreateSubnet"); n != notFoundPatience.attempts {
		t.Errorf("CreateSubnet calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

func TestMapPublicIPIsSetOnCreateAndChangedInPlace(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	on := map[string]value.Value{"map_public_ip_on_launch": value.Bool(true, value.SourceExplicit)}
	st, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(vpc, on)))
	if err != nil {
		t.Fatal(err)
	}
	if !fake.Subnets()[0].MapPublicIP {
		t.Fatal("map_public_ip_on_launch was not applied on create")
	}
	st, err = p.Update(ctx, st, desired(typeSubnet, subnetAttrs(vpc, nil)))
	if err != nil {
		t.Fatal(err)
	}
	if fake.Subnets()[0].MapPublicIP {
		t.Error("update did not turn map_public_ip_on_launch off")
	}
	if b, _ := st.Attributes["map_public_ip_on_launch"].AsBool(); b {
		t.Error("returned state still says true")
	}
}

// TestAFailedAttributeAfterCreateStillReportsTheSubnet. infrata drops the result of an errored create,
// so an error here would leave a real subnet untracked.
func TestAFailedAttributeAfterCreateStillReportsTheSubnet(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	for n := 1; n <= 3; n++ { // outlast the SDK's own retries
		fake.Inject(ec2fake.Fault{Action: "ModifySubnetAttribute", Nth: n, Status: 400, Code: "UnauthorizedOperation", Message: "no"})
	}
	on := map[string]value.Value{"map_public_ip_on_launch": value.Bool(true, value.SourceExplicit)}
	st, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, on)))
	if err != nil || st == nil {
		t.Fatalf("Create = %v, %v; a subnet that exists must be reported", st, err)
	}
	if b, _ := st.Attributes["map_public_ip_on_launch"].AsBool(); b {
		t.Error("state claims an attribute AWS refused; it must be the truth, so the next plan converges")
	}
}

func TestDeletingASubnetThatIsAlreadyGoneSucceeds(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	st, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, nil)))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := p.Delete(context.Background(), st); err != nil {
			t.Fatalf("delete %d: %v", i+1, err)
		}
	}
}

// TestAVPCWithASubnetRefusesDeletionWithAHint. infrata orders destroys by reference, so this only
// happens when something outside infrata sits in the VPC.
func TestAVPCWithASubnetRefusesDeletionWithAHint(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpcState, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	id, _ := vpcState.Attributes["id"].AsString()
	if _, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(id, nil))); err != nil {
		t.Fatal(err)
	}
	err = p.Delete(ctx, &resource.ResourceState{Type: typeVPC, ProviderID: vpcState.ProviderID})
	if err == nil || !strings.Contains(err.Error(), "DependencyViolation") || !strings.Contains(err.Error(), "depends on it") {
		t.Fatalf("err = %v", err)
	}
	if len(fake.VPCs()) != 1 {
		t.Error("the VPC was deleted")
	}
}
