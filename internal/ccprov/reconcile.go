package ccprov

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/cfn"
	"github.com/infrena/infrena/pkg/value"
)

// Nested values are reconciled here. Outgoing, every key is sent under AWS's name. Incoming, a value is
// rewritten to its reference: the reference's spelling, without keys AWS added, and in its order where the schema says
// order does not matter. Opaque shapes (P5) are copied exactly both ways.

// matchProp finds the schema property a written key names: exactly, ignoring case, or by its snake_case form. The
// generator proved no two properties of one object fold together (Verification log).
func matchProp(props map[string]*catalog.Shape, key string) (string, bool) {
	if _, ok := props[key]; ok {
		return key, true
	}
	lower := strings.ToLower(key)
	for name := range props {
		if strings.ToLower(name) == lower || cfn.SnakeCase(name) == lower {
			return name, true
		}
	}
	return "", false
}

func spellings(props map[string]*catalog.Shape) string {
	names := make([]string, 0, len(props))
	for name := range props {
		if snake := cfn.SnakeCase(name); snake != strings.ToLower(name) {
			name += " (" + snake + ")"
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func copiedExactly(shape *catalog.Shape) bool {
	return shape == nil || shape.Kind == catalog.ShapeOpaque || shape.Kind == catalog.ShapeScalar
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// encode converts what a user wrote to the JSON Cloud Control takes. path names the value in errors.
func encode(shape *catalog.Shape, v value.Value, path string) (any, error) {
	if copiedExactly(shape) {
		return plain(v), nil
	}
	switch shape.Kind {
	case catalog.ShapeObject:
		items, ok := v.Raw.(map[string]value.Value)
		if v.Kind != value.KindMap || !ok {
			return nil, fmt.Errorf("%s must be a map with the keys %s, got a %s", path, spellings(shape.Props), v.Kind)
		}
		out := make(map[string]any, len(items))
		written := make(map[string]string, len(items))
		for _, k := range sortedKeys(items) {
			name, found := matchProp(shape.Props, k)
			if !found {
				return nil, fmt.Errorf("%s has no key %q; it accepts %s", path, k, spellings(shape.Props))
			}
			if first, dup := written[name]; dup {
				return nil, fmt.Errorf("%s sets %s twice, as %q and %q", path, name, first, k)
			}
			written[name] = k
			enc, err := encode(shape.Props[name], items[k], path+"."+name)
			if err != nil {
				return nil, err
			}
			out[name] = enc
		}
		return out, nil
	case catalog.ShapeArray:
		items, ok := v.Raw.([]value.Value)
		if v.Kind != value.KindList || !ok {
			return nil, fmt.Errorf("%s must be a list, got a %s", path, v.Kind)
		}
		out := make([]any, len(items))
		for i, item := range items {
			enc, err := encode(shape.Item, item, fmt.Sprintf("%s[%d]", path, i))
			if err != nil {
				return nil, err
			}
			out[i] = enc
		}
		return out, nil
	}
	return plain(v), nil
}

// decode converts what AWS returned, shaped by ref when there is one. It reports false for a null.
func decode(shape *catalog.Shape, datum any, ref *value.Value) (value.Value, bool) {
	if datum == nil {
		return value.Value{}, false
	}
	if ref != nil && !ref.Known {
		ref = nil
	}
	if ref != nil && (sameScalar(*ref, datum) || sameJSONDocument(*ref, datum)) {
		return *ref, true
	}
	if copiedExactly(shape) {
		return infer(datum)
	}
	switch shape.Kind {
	case catalog.ShapeObject:
		if m, ok := datum.(map[string]any); ok {
			return decodeObject(shape, m, ref), true
		}
	case catalog.ShapeArray:
		if items, ok := datum.([]any); ok {
			return decodeArray(shape, items, ref), true
		}
	}
	return infer(datum)
}

func decodeObject(shape *catalog.Shape, m map[string]any, ref *value.Value) value.Value {
	out := map[string]value.Value{}
	if refItems, ok := refMap(ref); ok {
		// The reference's keys, in its spelling. What AWS added is dropped; what AWS no longer has is absent.
		for rk, rv := range refItems {
			if name, found := matchProp(shape.Props, rk); found {
				if dv, ok := decode(shape.Props[name], m[name], &rv); ok {
					out[rk] = dv
				}
			}
		}
		return value.Map(out, value.SourceProvider)
	}
	for k, d := range m {
		if dv, ok := decode(shape.Props[k], d, nil); ok {
			out[cfn.SnakeCase(k)] = dv
		}
	}
	return value.Map(out, value.SourceProvider)
}

func decodeArray(shape *catalog.Shape, items []any, ref *value.Value) value.Value {
	refItems, hasRef := refList(ref)
	out := make([]value.Value, 0, len(items))
	if !hasRef || !shape.Unordered {
		for i, d := range items {
			var r *value.Value
			if i < len(refItems) {
				r = &refItems[i]
			}
			if dv, ok := decode(shape.Item, d, r); ok {
				out = append(out, dv)
			}
		}
		return value.List(out, value.SourceProvider)
	}

	// Unordered: each reference item takes the returned item equal to it, so the list comes back in the reference's
	// order.
	slots := make([]*value.Value, len(refItems))
	used := make([]bool, len(items))
	for i := range refItems {
		for j, d := range items {
			if used[j] {
				continue
			}
			if dv, ok := decode(shape.Item, d, &refItems[i]); ok && dv.Equal(refItems[i]) {
				slots[i], used[j] = &dv, true
				break
			}
		}
	}
	// A returned item nothing matched fills the next unmatched reference slot, keeping that item's spelling so the
	// change shows as a change; anything beyond is appended.
	var extra []value.Value
	next := 0
	for j, d := range items {
		if used[j] {
			continue
		}
		for next < len(slots) && slots[next] != nil {
			next++
		}
		var r *value.Value
		if next < len(slots) {
			r = &refItems[next]
		}
		dv, ok := decode(shape.Item, d, r)
		if !ok {
			continue
		}
		if next < len(slots) {
			slots[next] = &dv
			next++
		} else {
			extra = append(extra, dv)
		}
	}
	for _, s := range slots {
		if s != nil {
			out = append(out, *s)
		}
	}
	return value.List(append(out, extra...), value.SourceProvider)
}

func refMap(ref *value.Value) (map[string]value.Value, bool) {
	if ref == nil {
		return nil, false
	}
	m, ok := ref.Raw.(map[string]value.Value)
	return m, ok && ref.Kind == value.KindMap
}

func refList(ref *value.Value) ([]value.Value, bool) {
	if ref == nil {
		return nil, false
	}
	l, ok := ref.Raw.([]value.Value)
	return l, ok && ref.Kind == value.KindList
}

// sameScalar reports whether a returned scalar is the reference scalar written another way: 443 for "443".
func sameScalar(ref value.Value, datum any) bool {
	var text string
	switch d := datum.(type) {
	case string:
		text = d
	case json.Number:
		text = d.String()
	case bool:
		text = strconv.FormatBool(d)
	default:
		return false
	}
	switch r := ref.Raw.(type) {
	case string:
		return r == text
	case int64:
		return strconv.FormatInt(r, 10) == text
	case float64:
		f, err := strconv.ParseFloat(text, 64)
		return err == nil && f == r
	case bool:
		return strconv.FormatBool(r) == text
	}
	return false
}
