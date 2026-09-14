// Package cfnschema is THROWAWAY SPIKE CODE. It asks one question: can an infrata resource
// definition be derived mechanically from an AWS CloudFormation resource type schema — the schema
// AWS Cloud Control API is driven by — with no handwritten knowledge of the resource?
package cfnschema

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

// Doc is the subset of a resource type schema this spike reads.
// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
type Doc struct {
	TypeName                        string                     `json:"typeName"`
	Description                     string                     `json:"description"`
	Properties                      map[string]json.RawMessage `json:"properties"`
	Definitions                     map[string]json.RawMessage `json:"definitions"`
	Required                        []string                   `json:"required"`
	PrimaryIdentifier               []string                   `json:"primaryIdentifier"`
	ReadOnlyProperties              []string                   `json:"readOnlyProperties"`
	CreateOnlyProperties            []string                   `json:"createOnlyProperties"`
	WriteOnlyProperties             []string                   `json:"writeOnlyProperties"`
	ConditionalCreateOnlyProperties []string                   `json:"conditionalCreateOnlyProperties"`
	Handlers                        map[string]json.RawMessage `json:"handlers"`
}

type propSchema struct {
	Type        json.RawMessage `json:"type"`
	Ref         string          `json:"$ref"`
	Description string          `json:"description"`
}

// Report records what the mapping could not express, which is the interesting output of the spike.
type Report struct {
	Type string
	// WriteOnly are top-level properties AWS never returns. infrata needs a plugin to carry these
	// forward from current state on Read, or every plan shows a change.
	WriteOnly []string
	// Conditional are properties that are sometimes updatable and sometimes force replacement; the
	// schema does not say when.
	Conditional []string
	// NestedFlags are JSON pointers deeper than /properties/X. infrata's flags are per top-level
	// attribute, so these cannot be represented.
	NestedFlags []string
	// Guessed are properties whose kind had to be inferred (no type, a union, or a $ref chain).
	Guessed []string
	// Identifier is how a provider ID would be formed.
	Identifier []string
	HasUpdate  bool
}

// TypeName maps AWS::EC2::VPC to aws.ec2.vpc.
func TypeName(cfn string) string {
	return "aws." + strings.ToLower(strings.ReplaceAll(strings.TrimPrefix(cfn, "AWS::"), "::", "."))
}

// Parse reads a schema document.
func Parse(raw []byte) (*Doc, error) {
	var d Doc
	if err := json.Unmarshal(raw, &d); err != nil {
		return nil, err
	}
	if d.TypeName == "" || len(d.Properties) == 0 {
		return nil, fmt.Errorf("not a resource type schema")
	}
	return &d, nil
}

// Definition maps a schema to an infrata ResourceDefinition, mechanically.
func Definition(d *Doc) (*schema.ResourceDefinition, *Report) {
	rep := &Report{Type: d.TypeName, HasUpdate: d.Handlers["update"] != nil}
	top := func(pointers []string, into map[string]bool, nested *[]string) {
		for _, p := range pointers {
			name, rest, _ := strings.Cut(strings.TrimPrefix(p, "/properties/"), "/")
			if rest != "" {
				*nested = append(*nested, p)
				continue
			}
			into[name] = true
		}
	}
	readOnly, createOnly, writeOnly, conditional := map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}
	top(d.ReadOnlyProperties, readOnly, &rep.NestedFlags)
	top(d.CreateOnlyProperties, createOnly, &rep.NestedFlags)
	top(d.WriteOnlyProperties, writeOnly, &rep.NestedFlags)
	top(d.ConditionalCreateOnlyProperties, conditional, &rep.NestedFlags)
	required := map[string]bool{}
	for _, r := range d.Required {
		required[r] = true
	}
	for _, p := range d.PrimaryIdentifier {
		rep.Identifier = append(rep.Identifier, strings.TrimPrefix(p, "/properties/"))
	}

	attrs := map[string]schema.Attribute{}
	for name, raw := range d.Properties {
		kind, guessed := kindOf(d, raw, 0)
		if guessed {
			rep.Guessed = append(rep.Guessed, name)
		}
		var ps propSchema
		_ = json.Unmarshal(raw, &ps)
		a := schema.Attribute{Kind: kind, Description: firstLine(ps.Description)}
		switch {
		case readOnly[name]:
			a.Computed = true // AWS sets it; configuration may not
		default:
			a.Required = required[name]
			// No update handler means every change replaces the resource.
			a.ForceNew = createOnly[name] || !rep.HasUpdate
		}
		if writeOnly[name] {
			rep.WriteOnly = append(rep.WriteOnly, name)
		}
		if conditional[name] {
			rep.Conditional = append(rep.Conditional, name)
		}
		attrs[name] = a
	}
	sort.Strings(rep.WriteOnly)
	sort.Strings(rep.Conditional)
	sort.Strings(rep.Guessed)
	sort.Strings(rep.NestedFlags)

	return &schema.ResourceDefinition{
		Type:         TypeName(d.TypeName),
		Description:  firstLine(d.Description),
		Attributes:   attrs,
		Capabilities: schema.Capabilities{Create: true, Read: true, Update: rep.HasUpdate, Delete: true, Import: true},
		ImportID:     schema.ImportSpec{Description: "the Cloud Control identifier: " + strings.Join(rep.Identifier, "|")},
	}, rep
}

// kindOf resolves a property's JSON-schema type to an infrata kind, following $ref into definitions.
func kindOf(d *Doc, raw json.RawMessage, depth int) (value.Kind, bool) {
	var ps propSchema
	if err := json.Unmarshal(raw, &ps); err != nil || depth > 8 {
		return value.KindString, true
	}
	if ps.Ref != "" {
		def, ok := d.Definitions[strings.TrimPrefix(ps.Ref, "#/definitions/")]
		if !ok {
			return value.KindString, true
		}
		return kindOf(d, def, depth+1)
	}
	var single string
	if err := json.Unmarshal(ps.Type, &single); err != nil {
		// a union such as ["string","object"], or no type at all
		return value.KindString, true
	}
	switch single {
	case "string":
		return value.KindString, false
	case "integer":
		return value.KindInt, false
	case "number":
		return value.KindFloat, false
	case "boolean":
		return value.KindBool, false
	case "array":
		return value.KindList, false
	case "object":
		return value.KindMap, false
	}
	return value.KindString, true
}

func firstLine(s string) string {
	s, _, _ = strings.Cut(strings.TrimSpace(s), "\n")
	return s
}
