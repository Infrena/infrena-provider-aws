// Package cfn reads AWS CloudFormation resource type schemas: the documents AWS Cloud Control API is driven by.
// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
package cfn

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Schema is the part of a resource type schema the generator and provider read.
type Schema struct {
	TypeName             string             `json:"typeName"`
	Description          string             `json:"description"`
	Properties           map[string]*Node   `json:"properties"`
	Definitions          map[string]*Node   `json:"definitions"`
	Required             []string           `json:"required"`
	PrimaryIdentifier    []string           `json:"primaryIdentifier"`
	ReadOnlyProperties   []string           `json:"readOnlyProperties"`
	CreateOnlyProperties []string           `json:"createOnlyProperties"`
	WriteOnlyProperties  []string           `json:"writeOnlyProperties"`
	Handlers             map[string]Handler `json:"handlers"`
	Tagging              *Tagging           `json:"tagging"`
}

// Node is one JSON-schema node inside a resource schema.
type Node struct {
	Type                 TypeList         `json:"type"`
	Ref                  string           `json:"$ref"`
	Description          string           `json:"description"`
	Properties           map[string]*Node `json:"properties"`
	PatternProperties    map[string]*Node `json:"patternProperties"`
	AdditionalProperties json.RawMessage  `json:"additionalProperties"`
	Items                *Node            `json:"items"`
	InsertionOrder       *bool            `json:"insertionOrder"`
	OneOf                []*Node          `json:"oneOf"`
	AnyOf                []*Node          `json:"anyOf"`
	AllOf                []*Node          `json:"allOf"`
}

// TypeList is a JSON-schema "type", which may be a string or a list of strings.
type TypeList []string

func (t *TypeList) UnmarshalJSON(b []byte) error {
	var one string
	if err := json.Unmarshal(b, &one); err == nil {
		*t = TypeList{one}
		return nil
	}
	var many []string
	if err := json.Unmarshal(b, &many); err != nil {
		return fmt.Errorf("type is neither a string nor a list of strings: %s", b)
	}
	*t = many
	return nil
}

// Handler is one provisioning handler.
type Handler struct {
	Permissions      []string        `json:"permissions"`
	TimeoutInMinutes int             `json:"timeoutInMinutes"`
	HandlerSchema    json.RawMessage `json:"handlerSchema"`
}

// Tagging is how a type supports tags.
type Tagging struct {
	Taggable    *bool  `json:"taggable"`
	TagProperty string `json:"tagProperty"`
}

// Parse reads one schema document.
func Parse(raw []byte) (*Schema, error) {
	var s Schema
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, err
	}
	if s.TypeName == "" {
		return nil, fmt.Errorf("not a resource type schema: no typeName")
	}
	return &s, nil
}

// HasHandler reports whether the schema declares a handler.
func (s *Schema) HasHandler(name string) bool {
	_, ok := s.Handlers[name]
	return ok
}

// Provisionable reports whether Cloud Control can manage the type: it needs create, read and delete handlers.
// Measured against list-types on 2026-09-14, this selects exactly Cloud Control's provisionable set.
func (s *Schema) Provisionable() bool {
	return s.HasHandler("create") && s.HasHandler("read") && s.HasHandler("delete")
}

// ListNeedsModel reports whether listing requires a parent resource model: the list handler's schema requires
// something, whether as a plain top-level `required`, or inside a `oneOf`, `anyOf` or `allOf` branch (checked
// recursively, since a branch can itself contain another branch). AWS sometimes states the requirement only inside
// one of those instead of at the top level — for example AWS::ElasticLoadBalancingV2::Listener, whose list handler
// needs LoadBalancerArn or ListenerArns via a top-level `oneOf` with no top-level `required`.
func (s *Schema) ListNeedsModel() bool {
	h, ok := s.Handlers["list"]
	if !ok || len(h.HandlerSchema) == 0 {
		return false
	}
	var hs listHandlerRequirement
	if err := json.Unmarshal(h.HandlerSchema, &hs); err != nil {
		return false
	}
	return hs.needsModel()
}

// listHandlerRequirement is the part of a list handler's handlerSchema ListNeedsModel inspects.
type listHandlerRequirement struct {
	Required []string                 `json:"required"`
	OneOf    []listHandlerRequirement `json:"oneOf"`
	AnyOf    []listHandlerRequirement `json:"anyOf"`
	AllOf    []listHandlerRequirement `json:"allOf"`
}

func (hs listHandlerRequirement) needsModel() bool {
	if len(hs.Required) > 0 {
		return true
	}
	for _, branches := range [][]listHandlerRequirement{hs.OneOf, hs.AnyOf, hs.AllOf} {
		for _, b := range branches {
			if b.needsModel() {
				return true
			}
		}
	}
	return false
}

// Timeout is a handler's timeout in minutes; the schema default is 120.
func (s *Schema) Timeout(handler string) int {
	if h, ok := s.Handlers[handler]; ok && h.TimeoutInMinutes > 0 {
		return h.TimeoutInMinutes
	}
	return 120
}

// TopLevel returns the property names of pointers of the form /properties/Name.
func (s *Schema) TopLevel(pointers []string) map[string]bool {
	out := map[string]bool{}
	for _, p := range pointers {
		if name, ok := strings.CutPrefix(p, "/properties/"); ok && !strings.Contains(name, "/") {
			out[name] = true
		}
	}
	return out
}

// Nested returns the pointers deeper than /properties/Name, which infrena's per-attribute flags cannot express.
func (s *Schema) Nested(pointers []string) []string {
	var out []string
	for _, p := range pointers {
		if name, ok := strings.CutPrefix(p, "/properties/"); ok && strings.Contains(name, "/") {
			out = append(out, p)
		}
	}
	return out
}

// WhollyNested splits nested pointers by their top-level property into those that cover EVERY leaf of that property,
// and those that cover only some. infrena's flags are per-attribute, so a property is only safe to flag as a whole
// when nothing inside it is left out: AWS::Lambda::Function's Code names all seven of its leaves write-only, so the
// attribute is write-only and its configured value must be carried forward. A property with only some leaves named
// must NOT be flagged, because carrying the whole thing forward would hide real drift in the leaves AWS does return.
//
// Found by the first live Lambda run (2026-09-17): Code's pointers are all of the form /properties/Code/ZipFile, so
// TopLevel dropped every one of them, the catalog marked Code as ordinary, Cloud Control returned Code as an empty
// object, and the plan never converged. Across the 2026-09-14 bundle this splits 100 nested write-only properties
// into 13 wholly-nested and 87 partial.
func (s *Schema) WhollyNested(pointers []string) (whole map[string]bool, partial []string) {
	whole = map[string]bool{}
	named := map[string]map[string]bool{}
	for _, p := range s.Nested(pointers) {
		rest := strings.TrimPrefix(p, "/properties/")
		top, leaf, _ := strings.Cut(rest, "/")
		if named[top] == nil {
			named[top] = map[string]bool{}
		}
		// Only the first segment below the property is a leaf this can reason about; anything deeper is covered
		// by its own parent being named, or is partial, which is the safe answer either way.
		first, _, _ := strings.Cut(leaf, "/")
		named[top][first] = true
	}
	tops := make([]string, 0, len(named))
	for top := range named {
		tops = append(tops, top)
	}
	sort.Strings(tops)
	for _, top := range tops {
		leaves := s.leafNames(s.Properties[top])
		if len(leaves) == 0 {
			partial = append(partial, top)
			continue
		}
		covered := true
		for _, leaf := range leaves {
			if !named[top][leaf] {
				covered = false
				break
			}
		}
		if covered {
			whole[top] = true
		} else {
			partial = append(partial, top)
		}
	}
	return whole, partial
}

// leafNames returns the immediate property names of a node, following $ref and looking through an array's items.
// It returns nil for anything without a fixed set of named properties, which WhollyNested treats as partial.
func (s *Schema) leafNames(n *Node) []string {
	r := s.Resolve(n)
	if r == nil {
		return nil
	}
	if len(r.Properties) == 0 && r.Items != nil {
		r = s.Resolve(r.Items)
		if r == nil {
			return nil
		}
	}
	if len(r.Properties) == 0 || len(r.PatternProperties) > 0 {
		return nil
	}
	out := make([]string, 0, len(r.Properties))
	for name := range r.Properties {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Resolve follows $ref until it reaches a node without one. References to #/definitions/... and #/properties/... are
// supported. A loop, or a reference to something missing, resolves to nil.
func (s *Schema) Resolve(n *Node) *Node {
	seen := map[string]bool{}
	for n != nil && n.Ref != "" {
		if seen[n.Ref] {
			return nil
		}
		seen[n.Ref] = true
		n = s.pointer(n.Ref)
	}
	return n
}

// DescriptionOr returns the node's description, or fallback when the node is nil or has none.
func (n *Node) DescriptionOr(fallback string) string {
	if n == nil || n.Description == "" {
		return fallback
	}
	return n.Description
}

func (s *Schema) pointer(ref string) *Node {
	path, ok := strings.CutPrefix(ref, "#/")
	if !ok {
		return nil
	}
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return nil
	}
	var n *Node
	switch parts[0] {
	case "definitions":
		n = s.Definitions[parts[1]]
	case "properties":
		n = s.Properties[parts[1]]
	default:
		return nil
	}
	for i := 2; n != nil && i+1 < len(parts); i += 2 {
		if parts[i] != "properties" {
			return nil
		}
		n = n.Properties[parts[i+1]]
	}
	return n
}
