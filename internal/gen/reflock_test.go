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
	usable, _, err := l.Reconcile(sampleDerivation(), liveTypes, o, true)
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
	usable, _, _ = l2.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{}, true)
	if keys(usable) != "AWS::EC2::Subnet.VpcId" || status(t, l2, "AWS::EC2::Subnet.VpcId") != StatusAccepted {
		t.Errorf("without overlay entries usable = %s, want the tier-1 edge accepted", keys(usable))
	}
}

// TestANewReferenceFailsWithoutTheFlag: a schema refresh cannot add a relationship silently.
func TestANewReferenceFailsWithoutTheFlag(t *testing.T) {
	l := &ReferenceLock{}
	if _, _, err := l.Reconcile(sampleDerivation()[:1], liveTypes, OverlayReferences{}, true); err != nil {
		t.Fatal(err)
	}
	_, _, err := l.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{}, false)
	if err == nil || !strings.Contains(err.Error(), "AWS::RDS::DBInstance.MonitoringRoleArn -> AWS::IAM::Role.Arn") || !strings.Contains(err.Error(), "-accept-new-references") {
		t.Fatalf("err = %v, want the new edges named and the flag that accepts them", err)
	}
	if len(l.References) != 1 {
		t.Errorf("a refused generation changed the lock: %d entries", len(l.References))
	}
	// The already-locked edge alone regenerates without the flag.
	if _, _, err := l.Reconcile(sampleDerivation()[:1], liveTypes, OverlayReferences{}, false); err != nil {
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
	usable, warnings, err := l.Reconcile(derived, liveTypes, OverlayReferences{}, false)
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
	usable, warnings, err := l.Reconcile(nil, liveTypes, OverlayReferences{ApproveTargets: []string{"AWS::IAM::Role"}}, false)
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
	_, _, err := l.Reconcile(sampleDerivation(), liveTypes, OverlayReferences{Reject: []string{"AWS::EC2::Subnet.VpcID"}}, true)
	if err == nil || !strings.Contains(err.Error(), "AWS::EC2::Subnet.VpcID") {
		t.Fatalf("err = %v", err)
	}
}

func TestTheReferenceLockRoundTripsSorted(t *testing.T) {
	path := filepath.Join(t.TempDir(), "references.lock.json")
	l, err := LoadReferenceLock(path)
	if err != nil || len(l.References) != 0 {
		t.Fatalf("missing lock = %+v, %v", l, err)
	}
	d := sampleDerivation()
	if _, _, err := l.Reconcile([]Derivation{d[3], d[0], d[2], d[1]}, liveTypes, OverlayReferences{}, true); err != nil {
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
