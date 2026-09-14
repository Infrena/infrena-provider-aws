// Package catalog holds the generated description of every Cloud Control resource type the plugin serves: the
// infrata definitions it hands the host, and the runtime metadata the provider needs. It is produced by
// cmd/gen-cloudcontrol and embedded; nothing here is written by hand.
package catalog

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

// GlobalScope is the provider-ID scope of a type that has no region.
const GlobalScope = "global"

// Shape kinds.
const (
	ShapeObject = "object" // named properties: keys are translated and reconciled
	ShapeArray  = "array"  // items share one shape
	ShapeScalar = "scalar"
	ShapeOpaque = "opaque" // copied exactly: free-form maps, unions, anything without a fixed set of properties
)

// Catalog is every generated type.
type Catalog struct {
	Bundle          string   `json:"bundle"`                     // SHA-256 of the schema bundle it was generated from
	DiscoverDefault []string `json:"discover_default,omitempty"` // infrata type names discovered when an instance sets no discover_types (P4)
	Types           []*Type  `json:"types"`

	index map[string]*Type
}

// Type is one Cloud Control resource type.
type Type struct {
	Name           string         `json:"name"` // infrata type, e.g. aws.vpc
	CFN            string         `json:"cfn"`  // CloudFormation type, e.g. AWS::EC2::VPC
	Description    string         `json:"description,omitempty"`
	RegionAttr     string         `json:"region_attr,omitempty"` // "region", "aws_region", or "" for a global type
	Identifier     []string       `json:"identifier"`            // primary identifier property names, in order
	WriteOnly      []string       `json:"write_only,omitempty"`  // top-level properties AWS never returns
	HasUpdate      bool           `json:"has_update,omitempty"`
	HasList        bool           `json:"has_list,omitempty"`
	ListNeedsModel bool           `json:"list_needs_model,omitempty"`
	TagsAsMap      string         `json:"tags_as_map,omitempty"` // property exposed as a map instead of [{Key, Value}]
	Timeouts       map[string]int `json:"timeouts,omitempty"`    // handler timeouts in minutes
	Attributes     []*Attribute   `json:"attributes"`
	Requirements   []Requirement  `json:"requirements,omitempty"`
}

// Attribute is one top-level property.
type Attribute struct {
	Name        string   `json:"name"` // AWS's property name: the canonical, stored name
	Kind        string   `json:"kind"` // an infrata kind name: string, integer, float, boolean, list, map
	Required    bool     `json:"required,omitempty"`
	Computed    bool     `json:"computed,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	ForceNew    bool     `json:"force_new,omitempty"`
	Sensitive   bool     `json:"sensitive,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
	Description string   `json:"description,omitempty"`
	Shape       *Shape   `json:"shape,omitempty"` // for list and map kinds
}

// Shape describes a nested value for reconciliation.
type Shape struct {
	Kind      string            `json:"k"`
	Props     map[string]*Shape `json:"p,omitempty"`
	Item      *Shape            `json:"i,omitempty"`
	Unordered bool              `json:"u,omitempty"`
}

// Requirement is a pre-flight hint between types, from the overlay.
type Requirement struct {
	Name        string   `json:"name"`
	Types       []string `json:"types"` // infrata type names
	Description string   `json:"description"`
}

// Read decodes a gzip-compressed JSON catalog.
func Read(r io.Reader) (*Catalog, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var c Catalog
	if err := json.NewDecoder(zr).Decode(&c); err != nil {
		return nil, err
	}
	c.index = make(map[string]*Type, len(c.Types))
	for _, t := range c.Types {
		if _, dup := c.index[t.Name]; dup {
			return nil, fmt.Errorf("catalog declares %s twice", t.Name)
		}
		c.index[t.Name] = t
	}
	return &c, nil
}

// Write encodes the catalog as gzip-compressed JSON, types sorted by name so regeneration diffs cleanly.
func (c *Catalog) Write(w io.Writer) error {
	sort.Slice(c.Types, func(i, j int) bool { return c.Types[i].Name < c.Types[j].Name })
	zw := gzip.NewWriter(w)
	enc := json.NewEncoder(zw)
	if err := enc.Encode(c); err != nil {
		return err
	}
	return zw.Close()
}

// Lookup finds a type by its infrata name.
func (c *Catalog) Lookup(name string) (*Type, bool) {
	if c.index == nil {
		c.index = make(map[string]*Type, len(c.Types))
		for _, t := range c.Types {
			c.index[t.Name] = t
		}
	}
	t, ok := c.index[name]
	return t, ok
}

// Definitions converts every type to an infrata definition.
func (c *Catalog) Definitions() []*schema.ResourceDefinition {
	out := make([]*schema.ResourceDefinition, 0, len(c.Types))
	for _, t := range c.Types {
		out = append(out, t.Definition())
	}
	return out
}

// Global reports whether the type has no region.
func (t *Type) Global() bool { return t.RegionAttr == "" }

// Attribute finds an attribute by its canonical name.
func (t *Type) Attribute(name string) (*Attribute, bool) {
	for _, a := range t.Attributes {
		if a.Name == name {
			return a, true
		}
	}
	return nil, false
}

// Definition is the infrata schema for the type. The region attribute is the plugin's own (spec §3.1).
func (t *Type) Definition() *schema.ResourceDefinition {
	attrs := make(map[string]schema.Attribute, len(t.Attributes)+1)
	for _, a := range t.Attributes {
		kind, _ := value.ParseKind(a.Kind)
		attrs[a.Name] = schema.Attribute{
			Kind: kind, Required: a.Required, Computed: a.Computed, Optional: a.Optional,
			ForceNew: a.ForceNew, Sensitive: a.Sensitive, Aliases: a.Aliases, Description: a.Description,
		}
	}
	scope := "<region>"
	if t.Global() {
		scope = GlobalScope
	} else {
		attrs[t.RegionAttr] = schema.Attribute{
			Kind: value.KindString, Required: true, ForceNew: true,
			Description: "AWS region, usually set once with the provider's defaults: {region: ...}",
		}
	}
	var reqs []schema.Requirement
	for _, r := range t.Requirements {
		reqs = append(reqs, schema.Requirement{Name: r.Name, Types: r.Types, Description: r.Description})
	}
	return &schema.ResourceDefinition{
		Type:         t.Name,
		Description:  t.Description,
		Attributes:   attrs,
		Requirements: reqs,
		Capabilities: schema.Capabilities{Create: true, Read: true, Update: t.HasUpdate, Delete: true, Import: true},
		ImportID: schema.ImportSpec{Description: fmt.Sprintf("%s/<identifier>, the identifier being %s (%s)",
			scope, strings.Join(t.Identifier, "|"), t.CFN)},
	}
}
