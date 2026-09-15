package gen

import (
	"sort"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

// Why a candidate property did not resolve.
const (
	UnresolvedAmbiguous   = "ambiguous"                  // more than one type remains after the same-service tiebreak
	UnresolvedNoType      = "no such type"               // no provisionable type's last segment matches
	UnresolvedNoAttribute = "target lacks the attribute" // one type matched, but it has no Id or Arn to hold
)

// Derivation is what the name heuristic says about one candidate property: a writable top-level property named
// <Base><Suffix>, Suffix one of Id, Ids, Arn, Arns (resource references design §4). Tier is 1 for an exact segment
// match, 2 for the suffix fallback, and 0 when unresolved, in which case Unresolved says why.
type Derivation struct {
	Source     string
	Property   string
	Target     string
	Attribute  string
	Tier       int
	Unresolved string
	Candidates []string // the types left when ambiguous, sorted
}

// referenceSuffixes, longest first so Ids is not read as I-d-s.
var referenceSuffixes = []struct{ suffix, kind string }{{"Ids", "Id"}, {"Arns", "Arn"}, {"Id", "Id"}, {"Arn", "Arn"}}

// DeriveReferences runs the heuristic over every provisionable schema, sorted by source type then property. Only
// provisionable types are sources or targets.
func DeriveReferences(schemas []*cfn.Schema) []Derivation {
	bySegment := map[string][]*cfn.Schema{}
	var sources []*cfn.Schema
	for _, s := range schemas {
		if !s.Provisionable() {
			continue
		}
		if _, seg, err := split(s.TypeName); err == nil {
			bySegment[seg] = append(bySegment[seg], s)
			sources = append(sources, s)
		}
	}
	var out []Derivation
	for _, s := range sources {
		readOnly := s.TopLevel(s.ReadOnlyProperties)
		for p := range s.Properties {
			if readOnly[p] {
				continue
			}
			base, kind, ok := cutReferenceSuffix(p)
			if !ok {
				continue
			}
			out = append(out, derive(s, p, base, kind, bySegment))
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Source != out[j].Source {
			return out[i].Source < out[j].Source
		}
		return out[i].Property < out[j].Property
	})
	return out
}

func cutReferenceSuffix(p string) (base, kind string, ok bool) {
	for _, s := range referenceSuffixes {
		if b, found := strings.CutSuffix(p, s.suffix); found {
			return b, s.kind, b != ""
		}
	}
	return "", "", false
}

func derive(s *cfn.Schema, property, base, kind string, bySegment map[string][]*cfn.Schema) Derivation {
	d := Derivation{Source: s.TypeName, Property: property}
	lower := strings.ToLower(base)
	candidates, tier := bySegment[lower], 1
	if len(candidates) == 0 {
		longest := ""
		for seg := range bySegment {
			if len(seg) > len(longest) && len(seg) < len(lower) && strings.HasSuffix(lower, seg) {
				longest = seg
			}
		}
		candidates, tier = bySegment[longest], 2
	}
	if len(candidates) == 0 {
		d.Unresolved = UnresolvedNoType
		return d
	}
	if len(candidates) > 1 {
		service, _, _ := split(s.TypeName)
		var same []*cfn.Schema
		for _, c := range candidates {
			if cs, _, _ := split(c.TypeName); cs == service {
				same = append(same, c)
			}
		}
		if len(same) == 1 {
			candidates = same
		}
	}
	if len(candidates) > 1 {
		d.Unresolved = UnresolvedAmbiguous
		for _, c := range candidates {
			d.Candidates = append(d.Candidates, c.TypeName)
		}
		sort.Strings(d.Candidates)
		return d
	}
	target := candidates[0]
	attr := targetAttribute(target, kind)
	if attr == "" {
		d.Unresolved = UnresolvedNoAttribute
		return d
	}
	d.Target, d.Attribute, d.Tier = target.TypeName, attr, tier
	return d
}

// targetAttribute picks which of the target's properties a reference of this kind holds: a property named exactly Id or
// Arn, else <Segment>Id or <Segment>Arn (VPC's VpcId), else a primary identifier ending in the kind (ManagedPolicy's
// PolicyArn). Empty when none applies.
func targetAttribute(target *cfn.Schema, kind string) string {
	if target.Properties[kind] != nil {
		return kind
	}
	_, seg, _ := split(target.TypeName)
	for p := range target.Properties {
		if strings.EqualFold(p, seg+kind) {
			return p
		}
	}
	var fromIdentifier []string
	for _, pointer := range target.PrimaryIdentifier {
		p, ok := strings.CutPrefix(pointer, "/properties/")
		if ok && strings.HasSuffix(p, kind) && target.Properties[p] != nil {
			fromIdentifier = append(fromIdentifier, p)
		}
	}
	if len(fromIdentifier) == 1 {
		return fromIdentifier[0]
	}
	return ""
}
