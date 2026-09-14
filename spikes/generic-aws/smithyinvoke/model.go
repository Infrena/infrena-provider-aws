// Package smithyinvoke is THROWAWAY SPIKE CODE for approaches A and C: invoke ANY AWS operation from a
// published Smithy model, with no generated per-operation Go type at all.
//
// The model (the JSON AST AWS publishes in github.com/aws/api-models-aws, and vendors into
// aws-sdk-go-v2/codegen/sdk-codegen/aws-models) is read at runtime and turned into smithy-go's own
// runtime schema type, *smithy.Schema. smithy-go's protocol codecs (here ec2query) then do the wire
// format — the same codecs AWS's newest generated clients use. Only the value walk is ours.
package smithyinvoke

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/traits"
)

type shapeRef struct {
	Target string `json:"target"`
}

type member struct {
	Target string                     `json:"target"`
	Traits map[string]json.RawMessage `json:"traits"`
}

type shape struct {
	Type    string                     `json:"type"`
	Version string                     `json:"version"`
	Members map[string]member          `json:"members"`
	Member  *member                    `json:"member"`
	Key     *member                    `json:"key"`
	Value   *member                    `json:"value"`
	Input   *shapeRef                  `json:"input"`
	Output  *shapeRef                  `json:"output"`
	Traits  map[string]json.RawMessage `json:"traits"`
}

// Model is a loaded Smithy JSON AST for one service.
type Model struct {
	shapes    map[string]shape
	namespace string
	service   string
	version   string
	protocols []string
	cache     map[string]*smithy.Schema
}

// Load reads a model file.
func Load(path string) (*Model, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc struct {
		Shapes map[string]shape `json:"shapes"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}
	m := &Model{shapes: doc.Shapes, cache: map[string]*smithy.Schema{}}
	for id, s := range doc.Shapes {
		if s.Type != "service" {
			continue
		}
		m.service, m.version = id, s.Version
		m.namespace, _, _ = strings.Cut(id, "#")
		for t := range s.Traits {
			if strings.HasPrefix(t, "aws.protocols#") {
				m.protocols = append(m.protocols, t)
			}
		}
	}
	if m.service == "" {
		return nil, fmt.Errorf("%s has no service shape", path)
	}
	return m, nil
}

// Operations lists every operation the model declares — discovered, not written down.
func (m *Model) Operations() []string {
	var out []string
	for id, s := range m.shapes {
		if s.Type == "operation" {
			_, name, _ := strings.Cut(id, "#")
			out = append(out, name)
		}
	}
	return out
}

// Version is the service's API version.
func (m *Model) Version() string { return m.version }

// Protocols are the aws.protocols traits on the service.
func (m *Model) Protocols() []string { return m.protocols }

// Operation builds runtime schemas for one operation, its input and its output.
func (m *Model) Operation(name string) (*smithy.OperationSchema, error) {
	id := m.namespace + "#" + name
	op, ok := m.shapes[id]
	if !ok || op.Type != "operation" {
		return nil, fmt.Errorf("no operation %q in %s", name, m.service)
	}
	var in, out *smithy.Schema
	var err error
	if op.Input != nil {
		if in, err = m.schema(op.Input.Target); err != nil {
			return nil, err
		}
	}
	if op.Output != nil {
		if out, err = m.schema(op.Output.Target); err != nil {
			return nil, err
		}
	}
	return smithy.NewOperationSchema(smithy.NewSchema(stoid(id), smithy.ShapeTypeOperation, 0), in, out), nil
}

// Service is the service schema a protocol is constructed from.
func (m *Model) Service() *smithy.ServiceSchema {
	return smithy.NewServiceSchema(smithy.NewSchema(stoid(m.service), smithy.ShapeTypeService, 0), m.version)
}

var shapeTypes = map[string]smithy.ShapeType{
	"blob": smithy.ShapeTypeBlob, "boolean": smithy.ShapeTypeBoolean, "string": smithy.ShapeTypeString,
	"timestamp": smithy.ShapeTypeTimestamp, "byte": smithy.ShapeTypeByte, "short": smithy.ShapeTypeShort,
	"integer": smithy.ShapeTypeInteger, "long": smithy.ShapeTypeLong, "float": smithy.ShapeTypeFloat,
	"double": smithy.ShapeTypeDouble, "document": smithy.ShapeTypeDocument, "bigDecimal": smithy.ShapeTypeBigDecimal,
	"bigInteger": smithy.ShapeTypeBigInteger, "enum": smithy.ShapeTypeEnum, "intEnum": smithy.ShapeTypeIntEnum,
	"list": smithy.ShapeTypeList, "set": smithy.ShapeTypeSet, "map": smithy.ShapeTypeMap,
	"structure": smithy.ShapeTypeStructure, "union": smithy.ShapeTypeUnion,
}

// prelude shapes are not in a service model; their type is their name.
var prelude = map[string]smithy.ShapeType{
	"String": smithy.ShapeTypeString, "Boolean": smithy.ShapeTypeBoolean, "PrimitiveBoolean": smithy.ShapeTypeBoolean,
	"Integer": smithy.ShapeTypeInteger, "PrimitiveInteger": smithy.ShapeTypeInteger, "Long": smithy.ShapeTypeLong,
	"PrimitiveLong": smithy.ShapeTypeLong, "Double": smithy.ShapeTypeDouble, "Float": smithy.ShapeTypeFloat,
	"Timestamp": smithy.ShapeTypeTimestamp, "Blob": smithy.ShapeTypeBlob, "Document": smithy.ShapeTypeDocument,
	"Unit": smithy.ShapeTypeStructure,
}

// schema builds (and caches) the runtime schema for a shape ID. The schema is cached BEFORE its members
// are added, so recursive shapes terminate; AddMember shares the target's member map, so members added
// later are still visible through it.
func (m *Model) schema(id string) (*smithy.Schema, error) {
	if s, ok := m.cache[id]; ok {
		return s, nil
	}
	if ns, name, _ := strings.Cut(id, "#"); ns == "smithy.api" {
		typ, ok := prelude[name]
		if !ok {
			return nil, fmt.Errorf("unsupported prelude shape %s", id)
		}
		s := smithy.NewSchema(stoid(id), typ, 0)
		m.cache[id] = s
		return s, nil
	}
	sh, ok := m.shapes[id]
	if !ok {
		return nil, fmt.Errorf("model has no shape %s", id)
	}
	typ, ok := shapeTypes[sh.Type]
	if !ok {
		return nil, fmt.Errorf("shape %s has unsupported type %s", id, sh.Type)
	}
	s := smithy.NewSchema(stoid(id), typ, len(sh.Members), convertTraits(sh.Traits)...)
	m.cache[id] = s

	add := func(name string, mem *member) error {
		target, err := m.schema(mem.Target)
		if err != nil {
			return err
		}
		s.AddMember(name, target, convertTraits(mem.Traits)...)
		return nil
	}
	for name, mem := range sh.Members {
		mem := mem
		if err := add(name, &mem); err != nil {
			return nil, err
		}
	}
	for name, mem := range map[string]*member{"member": sh.Member, "key": sh.Key, "value": sh.Value} {
		if mem != nil {
			if err := add(name, mem); err != nil {
				return nil, err
			}
		}
	}
	return s, nil
}

// convertTraits maps the protocol-relevant traits from the model onto smithy-go's trait types. Traits the
// codecs do not read (documentation, required, ...) are dropped.
func convertTraits(in map[string]json.RawMessage) []smithy.Trait {
	var out []smithy.Trait
	str := func(raw json.RawMessage) string {
		var v string
		_ = json.Unmarshal(raw, &v)
		return v
	}
	for id, raw := range in {
		switch id {
		case "smithy.api#xmlName":
			out = append(out, &traits.XMLName{Name: str(raw)})
		case "aws.protocols#ec2QueryName":
			out = append(out, &traits.EC2QueryName{Name: str(raw)})
		case "smithy.api#xmlFlattened":
			out = append(out, &traits.XMLFlattened{})
		case "smithy.api#jsonName":
			out = append(out, &traits.JSONName{Name: str(raw)})
		}
	}
	return out
}

func stoid(s string) smithy.ShapeID {
	ns, n, _ := strings.Cut(s, "#")
	n, mem, _ := strings.Cut(n, "$")
	return smithy.ShapeID{Namespace: ns, Name: n, Member: mem}
}
