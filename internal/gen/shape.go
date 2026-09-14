package gen

import (
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

const maxShapeDepth = 12

// BuildShape resolves a property's nested structure for reconciliation. Anything without a fixed set of named
// properties is opaque (plan decision P5).
func BuildShape(s *cfn.Schema, n *cfn.Node) *catalog.Shape {
	return buildShape(s, n, 0)
}

func buildShape(s *cfn.Schema, orig *cfn.Node, depth int) *catalog.Shape {
	n := s.Resolve(orig)
	if n == nil || depth > maxShapeDepth {
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	}
	if len(n.OneOf)+len(n.AnyOf)+len(n.AllOf) > 0 || len(n.Type) > 1 {
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	}
	typ := ""
	if len(n.Type) == 1 {
		typ = n.Type[0]
	} else if n.Items != nil {
		typ = "array"
	} else if len(n.Properties) > 0 || len(n.PatternProperties) > 0 {
		typ = "object"
	}
	switch typ {
	case "array":
		sh := &catalog.Shape{Kind: catalog.ShapeArray, Unordered: unordered(orig) || unordered(n)}
		if n.Items == nil {
			sh.Item = &catalog.Shape{Kind: catalog.ShapeOpaque}
		} else {
			sh.Item = buildShape(s, n.Items, depth+1)
		}
		return sh
	case "object":
		if len(n.Properties) == 0 || len(n.PatternProperties) > 0 {
			return &catalog.Shape{Kind: catalog.ShapeOpaque}
		}
		sh := &catalog.Shape{Kind: catalog.ShapeObject, Props: make(map[string]*catalog.Shape, len(n.Properties))}
		for name, child := range n.Properties {
			sh.Props[name] = buildShape(s, child, depth+1)
		}
		return sh
	case "":
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	default:
		return &catalog.Shape{Kind: catalog.ShapeScalar}
	}
}

// unordered reports insertionOrder: false. The schema default is true.
func unordered(n *cfn.Node) bool {
	return n != nil && n.InsertionOrder != nil && !*n.InsertionOrder
}
