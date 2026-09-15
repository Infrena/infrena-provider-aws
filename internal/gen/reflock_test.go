package gen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func resolved(source, property, target, attribute string, tier int) Derivation {
	return Derivation{Source: source, Property: property, Target: target, Attribute: attribute, Tier: tier}
}

var liveTypes = map[string][]string{
	"AWS::EC2::Subnet":      {"VpcId", "SubnetId"},
	"AWS::EC2::VPC":         {"VpcId"},
	"AWS::IAM::Role":        {"Arn"},
	"AWS::RDS::DBInstance":  {"MonitoringRoleArn", "DBSystemId"},
	"AWS::Lambda::Function": {"Role", "KmsKeyArn"},
	"AWS::KMS::Key":         {"Arn", "KeyId"},
	"AWS::X::System":        {"SystemId"},
}

func sampleDerivation() []Derivation {
	return []Derivation{
		resolved("AWS::EC2::Subnet", "VpcId", "AWS::EC2::VPC", "VpcId", 1),
		resolved("AWS::Lambda::Function", "KmsKeyArn", "AWS::KMS::Key", "Arn", 2),
		resolved("AWS::RDS::DBInstance", "DBSystemId", "AWS::X::System", "SystemId", 2),
		resolved("AWS::RDS::DBInstance", "MonitoringRoleArn", "AWS::IAM::Role", "Arn", 2),
	}
}

func status(t *testing.T, l *ReferenceLock, key string) string {
	t.Helper()
	for _, r := range l.References {
		if r.Source+"."+r.Property == key {
			return r.Status
		}
	}
	t.Fatalf("%s is not in the lock", key)
	return ""
}

func keys(refs []LockedReference) string {
	var out []string
	for _, r := range refs {
		out = append(out, r.Source+"."+r.Property)
	}
	return strings.Join(out, ",")
}

// TestAcceptingNewReferencesAssignsStatusFromTierAndOverlay: tier 1 accepted, tier 2 pending unless its target is
// approved, and a reject wins even over tier 1. Only accepted and approved edges are usable.
func TestAcceptingNewReferencesAssignsStatusFromTierAndOverlay(t *testing.T) {
	l := &ReferenceLock{}
	o := OverlayReferences{ApproveTargets: []string{"AWS::IAM::Role"}, Reject: []string{"AWS::EC2::Subnet.VpcId"}}
	usable, _, err := l.Reconcile(sampleDerivation(), liveTypes, o, ReferenceFlags{AcceptNew: true})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"AWS::EC2::Subnet.VpcId":                 StatusRejected,
		"AWS::Lambda::Function.KmsKeyArn":        StatusPending,
		"AWS::RDS::DBInstance.DBSystemId":        StatusPending,
		"AWS::RDS::DBInstance.MonitoringRoleArn": StatusApproved,
	}
	for k, v := range want {
		if got := status(t, l, k); got != v {
			t.Errorf("%s = %s, want %s", k, got, v)
		}
	}
	if keys(usable) != "AWS::RDS::DBInstance.MonitoringRoleArn" {
		t.Errorf("usable = %s, want only the approved edge", keys(usable))
	}
	l2 := &ReferenceLock{}
	usable, _, _ = l2.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{}, ReferenceFlags{AcceptNew: true})
	if keys(usable) != "AWS::EC2::Subnet.VpcId" || status(t, l2, "AWS::EC2::Subnet.VpcId") != StatusAccepted {
		t.Errorf("without overlay entries usable = %s, want the tier-1 edge accepted", keys(usable))
	}
}

// TestANewReferenceFailsWithoutTheFlag: a schema refresh cannot add a relationship silently.
func TestANewReferenceFailsWithoutTheFlag(t *testing.T) {
	l := &ReferenceLock{}
	if _, _, err := l.Reconcile(sampleDerivation()[:1], liveTypes, OverlayReferences{}, ReferenceFlags{AcceptNew: true}); err != nil {
		t.Fatal(err)
	}
	_, _, err := l.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{}, ReferenceFlags{})
	if err == nil || !strings.Contains(err.Error(), "AWS::RDS::DBInstance.MonitoringRoleArn -> AWS::IAM::Role.Arn") || !strings.Contains(err.Error(), "-accept-new-references") {
		t.Fatalf("err = %v, want the new edges named and the flag that accepts them", err)
	}
	if len(l.References) != 1 {
		t.Errorf("a refused generation changed the lock: %d entries", len(l.References))
	}
	// The already-locked edge alone regenerates without the flag.
	if _, _, err := l.Reconcile(sampleDerivation()[:1], liveTypes, OverlayReferences{}, ReferenceFlags{}); err != nil {
		t.Errorf("an unchanged derivation needed the flag: %v", err)
	}
}

// TestALockedReferenceNeverMoves: when a refresh derives a different target, or none, the lock's edge stands and a
// warning says so. The fixture derivation contradicts the lock on purpose.
func TestALockedReferenceNeverMoves(t *testing.T) {
	l := &ReferenceLock{References: []LockedReference{
		{Source: "AWS::EC2::Subnet", Property: "VpcId", Target: "AWS::EC2::VPC", Attribute: "VpcId", Tier: 1, Status: StatusAccepted},
		{Source: "AWS::RDS::DBInstance", Property: "MonitoringRoleArn", Target: "AWS::IAM::Role", Attribute: "Arn", Tier: 2, Status: StatusPending},
	}}
	derived := []Derivation{
		resolved("AWS::EC2::Subnet", "VpcId", "AWS::X::System", "SystemId", 1),
		{Source: "AWS::RDS::DBInstance", Property: "MonitoringRoleArn", Unresolved: UnresolvedAmbiguous, Candidates: []string{"AWS::IAM::Role", "AWS::X::Role"}},
	}
	usable, warnings, err := l.Reconcile(derived, liveTypes, OverlayReferences{}, ReferenceFlags{})
	if err != nil {
		t.Fatal(err)
	}
	if keys(usable) != "AWS::EC2::Subnet.VpcId" || usable[0].Target != "AWS::EC2::VPC" {
		t.Errorf("usable = %+v, want the locked VPC edge", usable)
	}
	joined := strings.Join(warnings, "\n")
	if !strings.Contains(joined, "AWS::EC2::Subnet.VpcId") || !strings.Contains(joined, "AWS::RDS::DBInstance.MonitoringRoleArn") {
		t.Errorf("warnings = %q, want both disagreements named", joined)
	}
}

// TestALockedReferenceToAGoneTargetIsKeptAndWarned, like a names.lock.json tombstone.
func TestALockedReferenceToAGoneTargetIsKeptAndWarned(t *testing.T) {
	l := &ReferenceLock{References: []LockedReference{
		{Source: "AWS::EC2::Subnet", Property: "VpcId", Target: "AWS::EC2::Gone", Attribute: "GoneId", Tier: 1, Status: StatusAccepted},
		{Source: "AWS::KMS::Key", Property: "RoleArn", Target: "AWS::IAM::Role", Attribute: "Arn", Tier: 2, Status: StatusApproved},
	}}
	usable, warnings, err := l.Reconcile(nil, liveTypes, OverlayReferences{ApproveTargets: []string{"AWS::IAM::Role"}}, ReferenceFlags{})
	if err != nil {
		t.Fatal(err)
	}
	if len(usable) != 0 {
		t.Errorf("usable = %s, want nothing: one target and one source property are gone", keys(usable))
	}
	if len(l.References) != 2 {
		t.Errorf("lock holds %d entries, want both kept", len(l.References))
	}
	joined := strings.Join(warnings, "\n")
	if !strings.Contains(joined, "AWS::EC2::Gone") || !strings.Contains(joined, "AWS::KMS::Key.RoleArn") {
		t.Errorf("warnings = %q", joined)
	}
}

// TestARejectMustNameAnEdge: a misspelt reject would silently refuse nothing.
func TestARejectMustNameAnEdge(t *testing.T) {
	l := &ReferenceLock{}
	_, _, err := l.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{Reject: []string{"AWS::EC2::Subnet.VpcID"}}, ReferenceFlags{AcceptNew: true})
	if err == nil || !strings.Contains(err.Error(), "AWS::EC2::Subnet.VpcID") {
		t.Fatalf("err = %v", err)
	}
}

func locked(t *testing.T, l *ReferenceLock, key string) LockedReference {
	t.Helper()
	for _, r := range l.References {
		if r.key() == key {
			return r
		}
	}
	t.Fatalf("%s is not in the lock", key)
	return LockedReference{}
}

// crossService: two exact matches that cross services, one to an allowlisted target (KMS::Key) and one not
// (ApiGateway::Resource, the FlowLog fabrication), beside a same-service exact match.
func crossService() []Derivation {
	return []Derivation{
		resolved("AWS::EC2::FlowLog", "ResourceId", "AWS::ApiGateway::Resource", "ResourceId", 1),
		resolved("AWS::EC2::Subnet", "VpcId", "AWS::EC2::VPC", "VpcId", 1),
		resolved("AWS::Lambda::Function", "KmsKeyId", "AWS::KMS::Key", "KeyId", 1),
	}
}

var crossServiceLive = map[string][]string{
	"AWS::EC2::FlowLog":         {"ResourceId"},
	"AWS::ApiGateway::Resource": {"ResourceId"},
	"AWS::EC2::Subnet":          {"VpcId"},
	"AWS::EC2::VPC":             {"VpcId"},
	"AWS::Lambda::Function":     {"KmsKeyId"},
	"AWS::KMS::Key":             {"KeyId"},
}

// TestACrossServiceExactMatchIsTierOneOnlyForAnAllowlistedTarget: a generic name like ResourceId matches exactly across
// services and is wrong, so an exact match to another service's type is tier 2 unless the target is listed under
// cross_service_targets. Same-service exact matches stay tier 1 with an empty allowlist.
func TestACrossServiceExactMatchIsTierOneOnlyForAnAllowlistedTarget(t *testing.T) {
	l := &ReferenceLock{}
	o := OverlayReferences{CrossServiceTargets: []string{"AWS::KMS::Key"}}
	usable, _, err := l.Reconcile(crossService(), crossServiceLive, o, ReferenceFlags{AcceptNew: true})
	if err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]LockedReference{
		"AWS::EC2::FlowLog.ResourceId":   {Tier: 2, Status: StatusPending},
		"AWS::EC2::Subnet.VpcId":         {Tier: 1, Status: StatusAccepted},
		"AWS::Lambda::Function.KmsKeyId": {Tier: 1, Status: StatusAccepted},
	} {
		if got := locked(t, l, key); got.Tier != want.Tier || got.Status != want.Status {
			t.Errorf("%s = tier %d %s, want tier %d %s", key, got.Tier, got.Status, want.Tier, want.Status)
		}
	}
	if keys(usable) != "AWS::EC2::Subnet.VpcId,AWS::Lambda::Function.KmsKeyId" {
		t.Errorf("usable = %s", keys(usable))
	}
	// Once tier 2, approve_targets approves it like any suffix match.
	o.ApproveTargets = []string{"AWS::ApiGateway::Resource"}
	if _, _, err := l.Reconcile(crossService(), crossServiceLive, o, ReferenceFlags{}); err != nil {
		t.Fatal(err)
	}
	if got := locked(t, l, "AWS::EC2::FlowLog.ResourceId"); got.Status != StatusApproved {
		t.Errorf("an approved demoted edge is %s", got.Status)
	}
}

// TestAReclassificationFailsWithoutItsFlag: the lock records each edge's tier, so an allowlist change that moves a
// locked edge between tiers is a lock change, refused until someone runs with the flag. Both directions, and every
// moved edge is recorded.
func TestAReclassificationFailsWithoutItsFlag(t *testing.T) {
	// A lock written before the allowlist existed: every exact match tier 1.
	l := &ReferenceLock{References: []LockedReference{
		{Source: "AWS::EC2::FlowLog", Property: "ResourceId", Target: "AWS::ApiGateway::Resource", Attribute: "ResourceId", Tier: 1, Status: StatusAccepted},
		{Source: "AWS::EC2::Subnet", Property: "VpcId", Target: "AWS::EC2::VPC", Attribute: "VpcId", Tier: 1, Status: StatusAccepted},
		{Source: "AWS::Lambda::Function", Property: "KmsKeyId", Target: "AWS::KMS::Key", Attribute: "KeyId", Tier: 2, Status: StatusPending},
	}}
	o := OverlayReferences{CrossServiceTargets: []string{"AWS::KMS::Key"}}
	_, _, err := l.Reconcile(crossService(), crossServiceLive, o, ReferenceFlags{AcceptNew: true})
	if err == nil || !strings.Contains(err.Error(), "AWS::EC2::FlowLog.ResourceId") || !strings.Contains(err.Error(), "AWS::Lambda::Function.KmsKeyId") ||
		!strings.Contains(err.Error(), AcceptReclassifiedReferencesFlag) {
		t.Fatalf("err = %v, want both reclassified edges and the flag named", err)
	}
	if strings.Contains(err.Error(), "AWS::EC2::Subnet.VpcId") {
		t.Errorf("err names an edge that did not move: %v", err)
	}
	if locked(t, l, "AWS::EC2::FlowLog.ResourceId").Tier != 1 || len(l.Changes) != 0 {
		t.Fatal("a refused generation changed the lock")
	}
	usable, _, err := l.Reconcile(crossService(), crossServiceLive, o, ReferenceFlags{AcceptReclassified: true})
	if err != nil {
		t.Fatal(err)
	}
	if got := locked(t, l, "AWS::EC2::FlowLog.ResourceId"); got.Tier != 2 || got.Status != StatusPending {
		t.Errorf("FlowLog.ResourceId = tier %d %s, want tier 2 pending", got.Tier, got.Status)
	}
	if got := locked(t, l, "AWS::Lambda::Function.KmsKeyId"); got.Tier != 1 || got.Status != StatusAccepted {
		t.Errorf("Lambda KmsKeyId = tier %d %s, want tier 1 accepted", got.Tier, got.Status)
	}
	if keys(usable) != "AWS::EC2::Subnet.VpcId,AWS::Lambda::Function.KmsKeyId" {
		t.Errorf("usable = %s", keys(usable))
	}
	changes := strings.Join(l.Changes, "\n")
	if len(l.Changes) != 2 || !strings.Contains(changes, "AWS::EC2::FlowLog.ResourceId") || !strings.Contains(changes, "accepted -> pending") ||
		!strings.Contains(changes, "pending -> accepted") {
		t.Errorf("changes = %q, want both moves recorded with their statuses", l.Changes)
	}
	// Settled: the same run again needs no flag and records nothing.
	if _, _, err := l.Reconcile(crossService(), crossServiceLive, o, ReferenceFlags{}); err != nil || len(l.Changes) != 0 {
		t.Errorf("a settled lock: err = %v, changes = %q", err, l.Changes)
	}
}

// TestARejectedTargetRejectsEveryEdgeToIt: a fabricated target type is refused for every edge reaching it from another
// service, whatever its tier and even over an approval of the same target, while an edge from the target's own service
// (ApiGateway::Method.ResourceId, which really holds an API Gateway resource id) keeps its normal status.
func TestARejectedTargetRejectsEveryEdgeToIt(t *testing.T) {
	derived := append(crossService(), resolved("AWS::ApiGateway::Method", "ResourceId", "AWS::ApiGateway::Resource", "ResourceId", 1))
	live := map[string][]string{"AWS::ApiGateway::Method": {"ResourceId"}}
	for k, v := range crossServiceLive {
		live[k] = v
	}
	l := &ReferenceLock{}
	o := OverlayReferences{
		CrossServiceTargets: []string{"AWS::KMS::Key"},
		ApproveTargets:      []string{"AWS::ApiGateway::Resource"},
		RejectTargets:       []string{"AWS::ApiGateway::Resource"},
	}
	usable, _, err := l.Reconcile(derived, live, o, ReferenceFlags{AcceptNew: true})
	if err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]string{
		"AWS::ApiGateway::Method.ResourceId": StatusAccepted, // tier 1, same service: not a fabrication
		"AWS::EC2::FlowLog.ResourceId":       StatusRejected, // tier 2, approved target
		"AWS::Lambda::Function.KmsKeyId":     StatusAccepted,
	} {
		if got := status(t, l, key); got != want {
			t.Errorf("%s = %s, want %s", key, got, want)
		}
	}
	if keys(usable) != "AWS::ApiGateway::Method.ResourceId,AWS::EC2::Subnet.VpcId,AWS::Lambda::Function.KmsKeyId" {
		t.Errorf("usable = %s", keys(usable))
	}
}

func TestTheReferenceLockRoundTripsSorted(t *testing.T) {
	path := filepath.Join(t.TempDir(), "references.lock.json")
	l, err := LoadReferenceLock(path)
	if err != nil || len(l.References) != 0 {
		t.Fatalf("missing lock = %+v, %v", l, err)
	}
	d := sampleDerivation()
	if _, _, err := l.Reconcile([]Derivation{d[3], d[0], d[2], d[1]}, liveTypes, OverlayReferences{}, ReferenceFlags{AcceptNew: true}); err != nil {
		t.Fatal(err)
	}
	if err := l.Save(path); err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(path)
	s := string(raw)
	if !(strings.Index(s, "AWS::EC2::Subnet") < strings.Index(s, "AWS::Lambda::Function") &&
		strings.Index(s, `"DBSystemId"`) < strings.Index(s, `"MonitoringRoleArn"`)) || !strings.HasSuffix(s, "\n") {
		t.Errorf("lock not sorted or no trailing newline:\n%s", raw)
	}
	again, err := LoadReferenceLock(path)
	if err != nil || len(again.References) != 4 || status(t, again, "AWS::Lambda::Function.KmsKeyArn") != StatusPending {
		t.Fatalf("reloaded = %+v, %v", again, err)
	}
}

// TestRequirementsWithoutAReferenceAreWarned: the overlay's requirements are human-approved relationship facts, so each
// one should have a usable edge from its type to one of its targets.
func TestRequirementsWithoutAReferenceAreWarned(t *testing.T) {
	reqs := map[string][]OverlayRequirement{
		"AWS::EC2::Subnet":     {{Name: "vpc", Types: []string{"AWS::EC2::VPC"}}},
		"AWS::RDS::DBInstance": {{Name: "role", Types: []string{"AWS::IAM::Role"}}},
		"AWS::KMS::Key":        {{Name: "vpc", Types: []string{"AWS::EC2::VPC"}}},
	}
	usable := []LockedReference{{Source: "AWS::EC2::Subnet", Property: "VpcId", Target: "AWS::EC2::VPC", Attribute: "VpcId", Tier: 1, Status: StatusAccepted}}
	all := append(append([]LockedReference(nil), usable...),
		LockedReference{Source: "AWS::RDS::DBInstance", Property: "MonitoringRoleArn", Target: "AWS::IAM::Role", Attribute: "Arn", Tier: 2, Status: StatusPending})
	warnings := CheckRequirements(reqs, usable, all)
	if len(warnings) != 2 {
		t.Fatalf("warnings = %q, want the DBInstance and Key requirements", warnings)
	}
	if !strings.Contains(warnings[0], "AWS::KMS::Key") || !strings.Contains(warnings[1], "AWS::RDS::DBInstance") || !strings.Contains(warnings[1], "pending") {
		t.Errorf("warnings = %q, want them sorted and the pending edge mentioned", warnings)
	}
}
