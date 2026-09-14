// Package cfn reads AWS CloudFormation resource type schemas: the documents AWS Cloud Control API is driven by.
// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
package cfn

import (
	"encoding/json"
	"fmt"
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

// ListNeedsModel reports whether listing requires a parent resource model (the list handler's schema has required
// properties).
func (s *Schema) ListNeedsModel() bool {
	h, ok := s.Handlers["list"]
	if !ok || len(h.HandlerSchema) == 0 {
		return false
	}
	var hs struct {
		Required []string `json:"required"`
	}
	if err := json.Unmarshal(h.HandlerSchema, &hs); err != nil {
		return false
	}
	return len(hs.Required) > 0
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

// Nested returns the pointers deeper than /properties/Name, which infrata's per-attribute flags cannot express.
func (s *Schema) Nested(pointers []string) []string {
	var out []string
	for _, p := range pointers {
		if name, ok := strings.CutPrefix(p, "/properties/"); ok && strings.Contains(name, "/") {
			out = append(out, p)
		}
	}
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
