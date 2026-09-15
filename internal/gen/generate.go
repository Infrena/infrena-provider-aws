package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/cfn"
)

// Generate builds the catalog from every provisionable AWS schema. The lock gains names for new types and refs gains
// derived reference edges and tier moves, those only as flags allows; saving both is the caller's job.
func Generate(schemas []*cfn.Schema, bundleSHA string, lock *Lock, refs *ReferenceLock, o *Overlay, flags ReferenceFlags) (*catalog.Catalog, []string, error) {
	var provisionable []*cfn.Schema
	var cfnTypes []string
	live := map[string][]string{}
	for _, s := range schemas {
		if strings.HasPrefix(s.TypeName, "AWS::") && s.Provisionable() {
			provisionable = append(provisionable, s)
			cfnTypes = append(cfnTypes, s.TypeName)
			for p := range s.Properties {
				live[s.TypeName] = append(live[s.TypeName], p)
			}
		}
	}
	if err := checkOverlay(o, cfnTypes); err != nil {
		return nil, nil, err
	}
	names, err := lock.Assign(cfnTypes)
	if err != nil {
		return nil, nil, err
	}
	usable, warnings, err := refs.Reconcile(DeriveReferences(provisionable), live, o.References, flags)
	if err != nil {
		return nil, nil, err
	}
	warnings = append(warnings, CheckRequirements(o.Requirements, usable, refs.References)...)
	cat := &catalog.Catalog{Bundle: bundleSHA}
	byCFN := map[string]*catalog.Type{}
	for _, s := range provisionable {
		t, w, err := BuildType(s, names, o)
		warnings = append(warnings, w...)
		if err != nil {
			return nil, warnings, err
		}
		cat.Types = append(cat.Types, t)
		byCFN[t.CFN] = t
	}
	for _, r := range usable {
		a, ok := byCFN[r.Source].Attribute(r.Property)
		target, tok := byCFN[r.Target]
		if !ok || !tok {
			return nil, warnings, fmt.Errorf("reference %s: source or target missing from the catalog", r)
		}
		if _, has := target.Attribute(r.Attribute); !has {
			return nil, warnings, fmt.Errorf("reference %s: %s has no attribute %s", r, target.Name, r.Attribute)
		}
		a.References = &catalog.Reference{Type: target.Name, Attribute: r.Attribute}
	}
	for _, cfnType := range o.DiscoverDefault {
		cat.DiscoverDefault = append(cat.DiscoverDefault, names[cfnType])
	}
	sort.Slice(cat.Types, func(i, j int) bool { return cat.Types[i].Name < cat.Types[j].Name })
	sort.Strings(warnings)
	return cat, warnings, nil
}

// checkOverlay refuses an overlay that names a type the bundle does not provide: it would silently apply nothing.
func checkOverlay(o *Overlay, cfnTypes []string) error {
	known := make(map[string]bool, len(cfnTypes))
	for _, t := range cfnTypes {
		known[t] = true
	}
	var unknown []string
	check := func(section, t string) {
		if !known[t] {
			unknown = append(unknown, section+": "+t)
		}
	}
	for t := range o.Aliases {
		check("aliases", t)
	}
	for t := range o.Sensitive {
		check("sensitive", t)
	}
	for t, reqs := range o.Requirements {
		check("requirements", t)
		for _, r := range reqs {
			for _, rt := range r.Types {
				check("requirements types", rt)
			}
		}
	}
	for _, t := range o.DiscoverDefault {
		check("discover_default", t)
	}
	for _, t := range o.References.CrossServiceTargets {
		check("references cross_service_targets", t)
	}
	for _, t := range o.References.ApproveTargets {
		check("references approve_targets", t)
	}
	for _, p := range o.Global {
		if !strings.HasSuffix(p, "*") {
			check("global", p)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return fmt.Errorf("overlay names types the bundle does not provide:\n  %s", strings.Join(unknown, "\n  "))
	}
	return nil
}
