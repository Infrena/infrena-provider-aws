package ccprov

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// withChanges is what infrata hands Update: the state's attributes, observed values included, overlaid with configuration.
func withChanges(st *resource.ResourceState, changes map[string]value.Value) *resource.DesiredResource {
	attrs := map[string]value.Value{}
	for k, v := range st.Attributes {
		attrs[k] = v
	}
	for k, v := range changes {
		attrs[k] = v
	}
	return desired(st.Type, attrs)
}

func createVPC(t *testing.T, p *Provider, extra map[string]value.Value) *resource.ResourceState {
	t.Helper()
	attrs := map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16")}
	for k, v := range extra {
		attrs[k] = v
	}
	st, err := p.Create(ctx, desired("aws.vpc", attrs))
	if err != nil {
		t.Fatal(err)
	}
	return st
}

func TestAChangedPropertyIsReplacedAndReadBack(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err != nil {
		t.Fatal(err)
	}
	want := []map[string]any{{"op": "replace", "path": "/EnableDnsSupport", "value": false}}
	if !reflect.DeepEqual(fake.LastPatch(), want) {
		t.Errorf("patch = %v, want %v: observed values such as VpcId must never be patched", fake.LastPatch(), want)
	}
	if attr(t, got, "EnableDnsSupport") != false {
		t.Errorf("state after update = %v", got.Attributes)
	}
	if tokens := fake.Tokens("UpdateResource"); len(tokens) != 1 || tokens[0] == "" {
		t.Errorf("client tokens = %v", tokens)
	}
}

func TestAPropertyStateDoesNotHoldIsAdded(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil {
		t.Fatal(err)
	}
	policies := list(obj("PolicyName", sv("read"), "PolicyDocument", obj("Version", sv("2012-10-17"))))
	if _, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"Policies": policies})); err != nil {
		t.Fatal(err)
	}
	if patch := fake.LastPatch(); len(patch) != 1 || patch[0]["op"] != "add" || patch[0]["path"] != "/Policies" {
		t.Errorf("patch = %v, want one add of /Policies", patch)
	}
}

// TestASpellingChangeAloneSendsNothingAndConverges. The next plan compares the new spelling against state, so the
// state returned must carry it.
func TestASpellingChangeAloneSendsNothingAndConverges(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.securitygroup", map[string]value.Value{
		"region": sv("us-east-1"), "GroupDescription": sv("web"), "SecurityGroupIngress": list(rule("tcp", 443, 443, nil)),
	}))
	if err != nil {
		t.Fatal(err)
	}
	respelled := list(rule("tcp", 443, 443, snake))
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"SecurityGroupIngress": respelled}))
	if err != nil {
		t.Fatal(err)
	}
	if n := fake.Calls("UpdateResource"); n != 0 {
		t.Errorf("UpdateResource calls = %d, want 0", n)
	}
	if !got.Attributes["SecurityGroupIngress"].Equal(respelled) {
		t.Errorf("state = %v, want the new spelling", got.Attributes["SecurityGroupIngress"])
	}
}

// TestADroppedPropertyIsNeverRemoved (PLAN §14.1): infrata sends no desired value for it, and AWS's is kept.
func TestADroppedPropertyIsNeverRemoved(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)})
	changes := withChanges(st, map[string]value.Value{"Tags": obj("team", sv("platform"))})
	delete(changes.Attrs, "EnableDnsSupport")
	if _, err := p.Update(ctx, st, changes); err != nil {
		t.Fatal(err)
	}
	for _, op := range fake.LastPatch() {
		if op["op"] == "remove" || op["path"] == "/EnableDnsSupport" {
			t.Errorf("patch %v touches a property configuration dropped", fake.LastPatch())
		}
	}
}

func TestAWriteOnlyValueIsPatchedOnlyWhenItChanges(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.dbinstance", map[string]value.Value{
		"region": sv("us-east-1"), "DBInstanceIdentifier": sv("db-1"), "MasterUserPassword": sv("first"), "DBInstanceClass": sv("db.t3.micro"),
	}))
	if err != nil {
		t.Fatal(err)
	}
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"MasterUserPassword": sv("second")}))
	if err != nil {
		t.Fatal(err)
	}
	if patch := fake.LastPatch(); len(patch) != 1 || patch[0]["path"] != "/MasterUserPassword" {
		t.Errorf("patch paths = %v", patch)
	}
	if attr(t, got, "MasterUserPassword") != "second" {
		t.Error("the new write-only value was not carried into state")
	}
	before := fake.Calls("UpdateResource")
	if _, err := p.Update(ctx, got, withChanges(got, map[string]value.Value{"DBInstanceClass": sv("db.t3.micro")})); err != nil {
		t.Fatal(err)
	}
	if fake.Calls("UpdateResource") != before {
		t.Error("an unchanged write-only value caused an update")
	}
}

func TestAChangeAWSWillNotMakeInPlaceSaysSo(t *testing.T) {
	p, _, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"CidrBlock": sv("10.1.0.0/16")}))
	if err == nil || !strings.Contains(err.Error(), "NotUpdatable") || !strings.Contains(err.Error(), "/CidrBlock") {
		t.Fatalf("err = %v", err)
	}
	if p.ClassifyError(err) != provider.NotSafeToRetry {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestAFailedUpdateRequestIsAnError(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	fake.FailNext("UPDATE", "InvalidRequest", "tenancy cannot be dedicated here", false)
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err == nil || !strings.Contains(err.Error(), "tenancy cannot be dedicated here") {
		t.Fatalf("err = %v", err)
	}
}

func TestATypeWithoutAnUpdateHandlerIsNeverPatched(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := &resource.ResourceState{Type: "aws.test.regioned", ProviderID: "us-east-1/reg-1", Attributes: map[string]value.Value{"Name": sv("a")}}
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"Name": sv("b")}))
	if err == nil || !strings.Contains(err.Error(), "replacement") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("UpdateResource") != 0 {
		t.Error("UpdateResource was called")
	}
}

func TestAnUpdateAlreadySentSurvivesCancellation(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	fake.PollsToComplete = 2
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	p.pacing.sleep = func(context.Context, time.Duration) error { cancel(); return nil }
	got, err := p.Update(cctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err != nil || got == nil {
		t.Fatalf("Update after cancellation = %v, %v", got, err)
	}
}
