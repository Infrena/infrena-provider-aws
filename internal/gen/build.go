package gen

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

// infrenaKeys are the resource keys infrena's configuration decoder claims, matched exactly
// (internal/config/decode.go at v0.3.0). A property whose lower-case spelling is one of them needs another name to be
// shown and written by (J9).
var infrenaKeys = map[string]bool{"type": true, "provider": true, "lifecycle": true, "depends_on": true, "skip": true, "only": true}

var secretName = regexp.MustCompile(`(?i)password|secret|token|privatekey|credential`)

// BuildType maps one provisionable schema to a catalog type.
func BuildType(s *cfn.Schema, names map[string]string, o *Overlay) (*catalog.Type, []string, error) {
	name, ok := names[s.TypeName]
	if !ok {
		return nil, nil, fmt.Errorf("%s has no assigned name", s.TypeName)
	}
	var warnings []string
	readOnly := s.TopLevel(s.ReadOnlyProperties)
	createOnly := s.TopLevel(s.CreateOnlyProperties)
	writeOnly := s.TopLevel(s.WriteOnlyProperties)
	required := map[string]bool{}
	for _, r := range s.Required {
		required[r] = true
	}
	sensitive := map[string]bool{}
	for _, p := range o.Sensitive[s.TypeName] {
		if s.Properties[p] == nil {
			return nil, nil, fmt.Errorf("overlay marks %s.%s sensitive, but the schema has no such property", s.TypeName, p)
		}
		sensitive[p] = true
	}
	for p := range o.Aliases[s.TypeName] {
		if s.Properties[p] == nil {
			return nil, nil, fmt.Errorf("overlay aliases %s.%s, but the schema has no such property", s.TypeName, p)
		}
	}

	t := &catalog.Type{
		Name: name, CFN: s.TypeName, Description: firstLine(s.Description),
		HasUpdate: s.HasHandler("update"), HasList: s.HasHandler("list"), ListNeedsModel: s.ListNeedsModel(),
		Timeouts: map[string]int{"create": s.Timeout("create"), "update": s.Timeout("update"), "delete": s.Timeout("delete")},
	}
	if !o.IsGlobal(s.TypeName) {
		t.RegionAttr = "region"
		for p := range s.Properties {
			if strings.EqualFold(p, "region") || cfn.SnakeCase(p) == "region" {
				t.RegionAttr = "aws_region"
			}
		}
	}
	for _, p := range s.PrimaryIdentifier {
		t.Identifier = append(t.Identifier, strings.TrimPrefix(p, "/properties/"))
	}
	// A property whose every leaf is write-only is write-only as a whole: AWS never returns any of it, so its
	// configured value must be carried forward or the plan never converges (see Schema.WhollyNested).
	whollyWriteOnly, partlyWriteOnly := s.WhollyNested(s.WriteOnlyProperties)
	for p := range whollyWriteOnly {
		writeOnly[p] = true
	}
	for p := range writeOnly {
		t.WriteOnly = append(t.WriteOnly, p)
	}
	sort.Strings(t.WriteOnly)
	if n := len(s.Nested(s.CreateOnlyProperties)); n > 0 {
		warnings = append(warnings, fmt.Sprintf("%s: %d nested create-only pointers cannot be expressed", s.TypeName, n))
	}
	// Partly-write-only properties get the same visibility create-only ones already had. Flagging one whole would
	// carry forward the leaves AWS does return and hide drift in them, so they stay ordinary and are named here.
	for _, p := range partlyWriteOnly {
		warnings = append(warnings, fmt.Sprintf("%s.%s: some but not all leaves are write-only; left ordinary", s.TypeName, p))
	}
	tagProp := tagsAsMap(s)
	t.TagsAsMap = tagProp

	props := make([]string, 0, len(s.Properties))
	for p := range s.Properties {
		props = append(props, p)
	}
	sort.Strings(props)
	for _, p := range props {
		node := s.Properties[p]
		kind, guessed := kindOf(s, node)
		if guessed {
			warnings = append(warnings, fmt.Sprintf("%s.%s: kind guessed as %s", s.TypeName, p, kind))
		}
		a := &catalog.Attribute{Name: p, Kind: kind, Description: firstLine(s.Resolve(node).DescriptionOr(node.Description)), Sensitive: sensitive[p]}
		switch {
		case readOnly[p]:
			a.Computed = true
		case required[p]:
			a.Required = true
		default:
			a.Optional, a.Computed = true, true
		}
		a.ForceNew = !readOnly[p] && (createOnly[p] || !t.HasUpdate)
		if writeOnly[p] && !sensitive[p] && secretName.MatchString(p) {
			warnings = append(warnings, fmt.Sprintf("%s.%s: write-only and looks secret; consider the overlay's sensitive list", s.TypeName, p))
		}
		if p == tagProp {
			a.Kind = "map"
			a.Shape = &catalog.Shape{Kind: catalog.ShapeOpaque}
			// Tags are user data, not a value the cloud chooses, so the tag map is Optional and NOT Computed.
			// infrena's configuration generator omits an optional+computed attribute unless it is also ForceNew
			// (it reads that combination as "a setting with a cloud default", like EnableDnsSupport), so while
			// this was computed, `import --generate` wrote no tags at all for 905 of the 958 taggable types —
			// they reached state correctly and then vanished from the generated configuration. The damage was
			// worse than a missing line: adding one tag by hand to an imported resource diffs one tag against
			// the three in state, so the plan proposes deleting the two the user never saw.
			// Reported from the infrena session, which could not fix it there: pkg/schema has no flag that
			// separates user data from a cloud default, so the catalog is the only place that knows.
			//
			// Safe because `aws:` system tags never reach state (internal/ccprov/tags.go's systemTagPrefix, and
			// TestSystemTagsAreNeverReported), so no plan proposes stripping AWS's own tags. It does mean a
			// hand-written resource that omits tags now proposes removing the ones AWS holds, which is the one
			// place this repository's "a dropped property keeps AWS's value" rule no longer applies — deliberate,
			// and the declarative answer. Nothing in this plugin reads the flag: ccprov's three Computed checks
			// all test `Computed && !Optional`, false either way once Computed is false.
			//
			// NEITHER flag, not Optional alone. infrena refuses Optional without Computed outright
			// (pkg/schema/definition.go: "every attribute that is not Required is already optional. Optional
			// exists to pair with Computed"), so a first attempt at Optional+!Computed was rejected by the host's
			// own validator and generated nothing. An attribute carrying no flag at all is the legal way to say
			// "settable, not required, not cloud-chosen", and it is what makes the host write the value:
			// its generator omits on `Computed && !(Optional && ForceNew)`, and its diff skips an unconfigured
			// attribute on `!defined || attr.Computed` (internal/planner/diff.go). With Computed false,
			// configuration is authoritative for tags in both places, which is the point.
			a.Optional, a.Computed = false, false
		} else if kind == "list" || kind == "map" {
			a.Shape = BuildShape(s, node)
		}
		a.Aliases = aliases(p, o.Aliases[s.TypeName][p], t.RegionAttr)
		t.Attributes = append(t.Attributes, a)
	}

	for _, r := range o.Requirements[s.TypeName] {
		req := catalog.Requirement{Name: r.Name, Description: r.Description}
		for _, cfnType := range r.Types {
			n, ok := names[cfnType]
			if !ok {
				return nil, nil, fmt.Errorf("overlay requirement on %s names %s, which is not a generated type", s.TypeName, cfnType)
			}
			req.Types = append(req.Types, n)
		}
		t.Requirements = append(t.Requirements, req)
	}

	if err := t.Definition().Validate(); err != nil {
		return nil, warnings, fmt.Errorf("%s: generated definition refused by infrena: %w", s.TypeName, err)
	}
	return t, warnings, nil
}

// aliases orders a property's spellings: curated first (Display shows the first), then a keyword-safe name when the
// property's own spelling is an infrena key, then snake_case. Duplicates under case folding, and anything that folds to
// the canonical name, an infrena key or the region attribute, are dropped.
func aliases(prop string, curated []string, regionAttr string) []string {
	snake := cfn.SnakeCase(prop)
	var candidates []string
	candidates = append(candidates, curated...)
	if infrenaKeys[strings.ToLower(prop)] || infrenaKeys[snake] {
		candidates = append(candidates, snake+"_value")
	}
	candidates = append(candidates, snake)
	seen := map[string]bool{strings.ToLower(prop): true}
	var out []string
	for _, c := range candidates {
		f := strings.ToLower(c)
		if seen[f] || infrenaKeys[f] || (regionAttr != "" && f == strings.ToLower(regionAttr)) {
			continue
		}
		seen[f] = true
		out = append(out, c)
	}
	return out
}

// tagsAsMap returns the tag property's name when it is a list of {Key, Value} objects, which the provider exposes as
// a map.
func tagsAsMap(s *cfn.Schema) string {
	if s.Tagging == nil || (s.Tagging.Taggable != nil && !*s.Tagging.Taggable) {
		return ""
	}
	name, ok := strings.CutPrefix(s.Tagging.TagProperty, "/properties/")
	if !ok || strings.Contains(name, "/") {
		return ""
	}
	prop := s.Resolve(s.Properties[name])
	if prop == nil || prop.Items == nil {
		return ""
	}
	item := s.Resolve(prop.Items)
	if item == nil || len(item.Properties) != 2 || item.Properties["Key"] == nil || item.Properties["Value"] == nil {
		return ""
	}
	return name
}

// kindOf maps a property's JSON-schema type to an infrena kind name, reporting whether it had to guess.
func kindOf(s *cfn.Schema, node *cfn.Node) (string, bool) {
	n := s.Resolve(node)
	if n == nil {
		return "string", true
	}
	if len(n.Type) == 1 {
		switch n.Type[0] {
		case "string":
			return "string", false
		case "integer":
			return "integer", false
		case "number":
			return "float", false
		case "boolean":
			return "boolean", false
		case "array":
			return "list", false
		case "object":
			return "map", false
		}
	}
	if len(n.Type) == 0 {
		if n.Items != nil {
			return "list", false
		}
		if len(n.Properties) > 0 || len(n.PatternProperties) > 0 {
			return "map", false
		}
	}
	return "string", true
}

func firstLine(s string) string {
	s, _, _ = strings.Cut(strings.TrimSpace(s), "\n")
	return s
}
