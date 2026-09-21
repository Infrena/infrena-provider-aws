package gen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

// Reference edge statuses in gen/references.lock.json.
const (
	StatusAccepted = "accepted" // tier 1: an exact type-name match, accepted without review
	StatusApproved = "approved" // tier 2, its target listed in the overlay's references.approve_targets
	StatusPending  = "pending"  // tier 2, not yet approved: kept out of the catalog
	StatusRejected = "rejected" // named in the overlay's references.reject, or its target in reject_targets, whatever its tier
)

// AcceptNewReferencesFlag is the gen-cloudcontrol flag that lets new edges into the lock.
const AcceptNewReferencesFlag = "-accept-new-references"

// AcceptReclassifiedReferencesFlag is the gen-cloudcontrol flag that lets the overlay's cross_service_targets move a
// locked edge between tier 1 and tier 2.
const AcceptReclassifiedReferencesFlag = "-accept-reclassified-references"

// ReferenceFlags says which lock changes one generation may make. Without its flag, each fails generation.
type ReferenceFlags struct {
	AcceptNew          bool // derived edges the lock does not hold yet
	AcceptReclassified bool // locked edges whose tier the cross-service rule now puts elsewhere
}

// ReferenceLock records every reference edge generation has ever derived, so a schema refresh cannot add, move or drop
// a relationship silently. Like the names lock, it only grows.
type ReferenceLock struct {
	References []LockedReference `json:"references"`
	// Changes lists every locked edge whose status or tier the last Reconcile changed, for the generator to print.
	Changes []string `json:"-"`
}

// LockedReference is one edge: Source.Property holds Target's Attribute.
type LockedReference struct {
	Source    string `json:"source"`
	Property  string `json:"property"`
	Target    string `json:"target"`
	Attribute string `json:"attribute"`
	Tier      int    `json:"tier"`
	Status    string `json:"status"`
}

// Usable reports whether the edge reaches the catalog.
func (r LockedReference) Usable() bool {
	return r.Status == StatusAccepted || r.Status == StatusApproved
}

func (r LockedReference) key() string { return r.Source + "." + r.Property }

func (r LockedReference) String() string {
	return fmt.Sprintf("%s.%s -> %s.%s (tier %d)", r.Source, r.Property, r.Target, r.Attribute, r.Tier)
}

// LoadReferenceLock reads a reference lock; a missing file is an empty lock.
func LoadReferenceLock(path string) (*ReferenceLock, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return &ReferenceLock{}, nil
	}
	if err != nil {
		return nil, err
	}
	var l ReferenceLock
	if err := json.Unmarshal(raw, &l); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &l, nil
}

// Save writes the lock sorted by source type and property, with a trailing newline.
func (l *ReferenceLock) Save(path string) error {
	l.sort()
	if l.References == nil {
		l.References = []LockedReference{}
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(l); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

func (l *ReferenceLock) sort() {
	sort.Slice(l.References, func(i, j int) bool {
		a, b := l.References[i], l.References[j]
		if a.Source != b.Source {
			return a.Source < b.Source
		}
		return a.Property < b.Property
	})
}

// Reconcile brings the lock up to date with a derivation and the overlay, and returns the edges that reach the catalog.
//
// live maps every current provisionable type to its top-level property names. A locked edge always stands over what
// the heuristic now derives; a disagreement is a warning. A locked edge whose source or target no longer exists is
// kept, warned about, and not returned. A derived edge the lock does not hold fails generation unless flags.AcceptNew is
// set, and a locked edge the cross-service rule moves to another tier fails it unless flags.AcceptReclassified is set.
// Statuses follow the overlay on every run, since the overlay is itself reviewed; every locked edge whose status or tier
// changed is recorded in l.Changes. On error the lock is unchanged.
func (l *ReferenceLock) Reconcile(derived []Derivation, live map[string][]string, o OverlayReferences, flags ReferenceFlags) ([]LockedReference, []string, error) {
	locked := map[string]int{}
	for i, r := range l.References {
		if _, dup := locked[r.key()]; dup {
			return nil, nil, fmt.Errorf("reference lock holds %s twice", r.key())
		}
		locked[r.key()] = i
	}
	byKey := map[string]Derivation{}
	for _, d := range derived {
		byKey[d.Source+"."+d.Property] = d
	}
	rejected := map[string]bool{}
	var unknownRejects []string
	for _, k := range o.Reject {
		rejected[k] = true
		_, isLocked := locked[k]
		if d, ok := byKey[k]; !isLocked && (!ok || d.Tier == 0) {
			unknownRejects = append(unknownRejects, k)
		}
	}
	if len(unknownRejects) > 0 {
		sort.Strings(unknownRejects)
		return nil, nil, fmt.Errorf("overlay references.reject names edges that are neither derived nor locked:\n  %s", strings.Join(unknownRejects, "\n  "))
	}
	approved := map[string]bool{}
	for _, t := range o.ApproveTargets {
		approved[t] = true
	}
	rejectedTarget := map[string]bool{}
	for _, t := range o.RejectTargets {
		rejectedTarget[t] = true
	}
	statusOf := func(r LockedReference) string {
		sourceService, _, _ := split(r.Source)
		targetService, _, _ := split(r.Target)
		switch {
		case rejected[r.key()], rejectedTarget[r.Target] && sourceService != targetService:
			return StatusRejected
		case r.Tier == 1:
			return StatusAccepted
		case approved[r.Target]:
			return StatusApproved
		default:
			return StatusPending
		}
	}

	crossService := map[string]bool{}
	for _, t := range o.CrossServiceTargets {
		crossService[t] = true
	}

	var added []LockedReference
	for _, d := range derived {
		if d.Tier == 0 {
			continue
		}
		if _, ok := locked[d.Source+"."+d.Property]; !ok {
			added = append(added, LockedReference{Source: d.Source, Property: d.Property, Target: d.Target, Attribute: d.Attribute,
				Tier: tierOf(d.Source, d.Target, d.Tier, crossService)})
		}
	}
	if len(added) > 0 && !flags.AcceptNew {
		lines := make([]string, len(added))
		for i, r := range added {
			lines[i] = r.String()
		}
		return nil, nil, fmt.Errorf("%d reference edges are not in gen/references.lock.json:\n  %s\nreview them, then run go run ./cmd/gen-cloudcontrol %s to add them "+
			"(tier 1 as accepted, tier 2 as pending unless its target is in the overlay's references.approve_targets)",
			len(added), strings.Join(lines, "\n  "), AcceptNewReferencesFlag)
	}

	// A locked edge's heuristic tier is the derivation's while the two still agree on the target; otherwise the lock's.
	retier := map[int]int{}
	var moved []string
	for i, r := range l.References {
		heuristic := r.Tier
		if d, ok := byKey[r.key()]; ok && d.Tier > 0 && d.Target == r.Target && d.Attribute == r.Attribute {
			heuristic = d.Tier
		}
		if want := tierOf(r.Source, r.Target, heuristic, crossService); want != r.Tier {
			retier[i] = want
			moved = append(moved, fmt.Sprintf("%s.%s -> %s.%s: tier %d -> %d", r.Source, r.Property, r.Target, r.Attribute, r.Tier, want))
		}
	}
	if len(moved) > 0 && !flags.AcceptReclassified {
		sort.Strings(moved)
		return nil, nil, fmt.Errorf("%d locked reference edges change tier under the overlay's references.cross_service_targets:\n  %s\n"+
			"review them, then run go run ./cmd/gen-cloudcontrol %s to record the move", len(moved), strings.Join(moved, "\n  "), AcceptReclassifiedReferencesFlag)
	}

	next := append(append([]LockedReference(nil), l.References...), added...)
	var changes []string
	for i := range l.References {
		before := l.References[i]
		r := &next[i]
		if tier, ok := retier[i]; ok {
			r.Tier = tier
		}
		r.Status = statusOf(*r)
		if r.Status != before.Status || r.Tier != before.Tier {
			change := fmt.Sprintf("%s.%s -> %s.%s: %s -> %s", r.Source, r.Property, r.Target, r.Attribute, before.Status, r.Status)
			if r.Tier != before.Tier {
				change += fmt.Sprintf(" (tier %d -> %d)", before.Tier, r.Tier)
			}
			changes = append(changes, change)
		}
	}
	var usable []LockedReference
	var warnings []string
	has := func(typ, prop string) bool {
		for _, p := range live[typ] {
			if p == prop {
				return true
			}
		}
		return false
	}
	for i := range next {
		r := &next[i]
		r.Status = statusOf(*r)
		alive := true
		switch {
		case !has(r.Source, r.Property):
			warnings = append(warnings, fmt.Sprintf("reference %s: the source property no longer exists; kept in the lock, not in the catalog", r))
			alive = false
		case !has(r.Target, r.Attribute):
			warnings = append(warnings, fmt.Sprintf("reference %s: the target no longer exists; kept in the lock, not in the catalog", r))
			alive = false
		}
		if d, ok := byKey[r.key()]; alive && (!ok || d.Target != r.Target || d.Attribute != r.Attribute) {
			now := "nothing"
			if ok && d.Tier > 0 {
				now = d.Target + "." + d.Attribute
			} else if ok {
				now = d.Unresolved
				if len(d.Candidates) > 0 {
					now += " (" + strings.Join(d.Candidates, ", ") + ")"
				}
			}
			warnings = append(warnings, fmt.Sprintf("reference %s: the heuristic now derives %s; the lock keeps its edge", r, now))
		}
		if alive && r.Usable() {
			usable = append(usable, *r)
		}
	}
	l.References = next
	l.sort()
	sort.Strings(changes)
	l.Changes = changes
	sort.Slice(usable, func(i, j int) bool { return usable[i].key() < usable[j].key() })
	sort.Strings(warnings)
	return usable, warnings, nil
}

// tierOf applies the cross-service rule to a heuristic tier: an exact match to another service's type is tier 2 unless
// its target is in crossService, because a generic property name matches some other service's type exactly and wrongly
// (ResourceId, NotificationArns). Tier 2 stays tier 2.
func tierOf(source, target string, heuristic int, crossService map[string]bool) int {
	sourceService, _, _ := split(source)
	targetService, _, _ := split(target)
	if heuristic == 1 && sourceService != targetService && !crossService[target] {
		return 2
	}
	return heuristic
}

// CheckRequirements cross-checks the overlay's hand-maintained requirements against the reference edges: each
// requirement should have a usable edge from its type to one of its target types. all is every locked edge, used only
// to say when a pending or rejected one exists.
func CheckRequirements(reqs map[string][]OverlayRequirement, usable, all []LockedReference) []string {
	edge := func(refs []LockedReference, source string, targets []string) *LockedReference {
		for i, r := range refs {
			if r.Source != source {
				continue
			}
			for _, t := range targets {
				if r.Target == t {
					return &refs[i]
				}
			}
		}
		return nil
	}
	var warnings []string
	for source, rs := range reqs {
		for _, r := range rs {
			if edge(usable, source, r.Types) != nil {
				continue
			}
			msg := fmt.Sprintf("requirement %s %q (%s) has no reference edge in the catalog", source, r.Name, strings.Join(r.Types, ", "))
			if e := edge(all, source, r.Types); e != nil {
				msg += fmt.Sprintf("; %s is %s", e, e.Status)
			}
			warnings = append(warnings, msg)
		}
	}
	sort.Strings(warnings)
	return warnings
}
