package ccprov

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// plain converts an infrata value to a JSON datum.
func plain(v value.Value) any {
	switch raw := v.Raw.(type) {
	case []value.Value:
		out := make([]any, len(raw))
		for i, item := range raw {
			out[i] = plain(item)
		}
		return out
	case map[string]value.Value:
		out := make(map[string]any, len(raw))
		for k, item := range raw {
			out[k] = plain(item)
		}
		return out
	default:
		return raw // string, int64, float64 or bool
	}
}

// infer converts a JSON datum decoded with UseNumber, taking the kind from the JSON. A null is absent.
func infer(datum any) (value.Value, bool) {
	const src = value.SourceProvider
	switch d := datum.(type) {
	case string:
		return value.String(d, src), true
	case bool:
		return value.Bool(d, src), true
	case json.Number:
		if i, err := d.Int64(); err == nil {
			return value.Int(i, src), true
		}
		f, err := d.Float64()
		return value.Float(f, src), err == nil
	case []any:
		items := make([]value.Value, 0, len(d))
		for _, item := range d {
			if v, ok := infer(item); ok {
				items = append(items, v)
			}
		}
		return value.List(items, src), true
	case map[string]any:
		items := make(map[string]value.Value, len(d))
		for k, item := range d {
			if v, ok := infer(item); ok {
				items[k] = v
			}
		}
		return value.Map(items, src), true
	}
	return value.Value{}, false
}

// coerce makes a returned top-level value the kind the schema declares. A string attribute takes anything, as JSON
// text: the generator types unions and untyped properties as strings.
func coerce(t *catalog.Type, a *catalog.Attribute, v value.Value, datum any) (value.Value, error) {
	want, _ := value.ParseKind(a.Kind)
	if v.Kind == want {
		return v, nil
	}
	const src = value.SourceProvider
	switch want {
	case value.KindString:
		raw, err := json.Marshal(datum)
		if err == nil {
			return value.String(string(raw), src), nil
		}
	case value.KindFloat:
		if i, ok := v.Raw.(int64); ok {
			return value.Float(float64(i), src), nil
		}
	case value.KindInt:
		if f, ok := v.Raw.(float64); ok && f == math.Trunc(f) {
			return value.Int(int64(f), src), nil
		}
	case value.KindBool:
		if s, ok := v.Raw.(string); ok {
			if b, err := strconv.ParseBool(s); err == nil {
				return value.Bool(b, src), nil
			}
		}
	}
	return value.Value{}, fmt.Errorf("AWS returned a %s for %s.%s, which the schema says is a %s", v.Kind, t.Name, a.Name, want)
}

// encodeAttr converts one top-level attribute to the JSON Cloud Control takes: nested keys under AWS's names, and tags
// as a list.
func encodeAttr(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error) {
	if a.Name == t.TagsAsMap {
		return tagsToJSON(t, a, v)
	}
	return encode(a.Shape, v, t.Name+"."+a.Name)
}

// decodeAttr converts one top-level property AWS returned, reconciled against reference when there is one. It reports
// false when there is nothing to record.
func decodeAttr(t *catalog.Type, a *catalog.Attribute, datum any, reference *value.Value) (value.Value, bool, error) {
	if a.Name == t.TagsAsMap {
		return tagsFromJSON(t, a, datum, reference)
	}
	v, ok := decode(a.Shape, datum, reference)
	if !ok {
		return value.Value{}, false, nil
	}
	v, err := coerce(t, a, v, datum)
	return v, err == nil, err
}

// desiredJSON is the Cloud Control desired state: every attribute configuration may set, under AWS's names. Never the
// plugin's region, never a read-only property (Update's desired attributes include observed ones).
func desiredJSON(t *catalog.Type, attrs map[string]value.Value) (map[string]any, error) {
	out := map[string]any{}
	for _, a := range t.Attributes {
		v, set := attrs[a.Name]
		if !set || (a.Computed && !a.Optional) {
			continue
		}
		enc, err := encodeAttr(t, a, v)
		if err != nil {
			return nil, err
		}
		out[a.Name] = enc
	}
	return out, nil
}

// decodeProperties parses GetResource's Properties, keeping numbers exact.
func decodeProperties(doc string) (map[string]any, error) {
	dec := json.NewDecoder(strings.NewReader(doc))
	dec.UseNumber()
	var props map[string]any
	if err := dec.Decode(&props); err != nil {
		return nil, err
	}
	return props, nil
}

// stateFrom builds the state infrata records from what AWS returned. reference is configuration's values (Create,
// Update), the previous state (Read), or nil (Discover, Import).
func stateFrom(t *catalog.Type, region, identifier string, props map[string]any, reference map[string]value.Value) (*resource.ResourceState, error) {
	attrs := map[string]value.Value{}
	for _, a := range t.Attributes {
		var ref *value.Value
		if r, ok := reference[a.Name]; ok && r.Known {
			ref = &r
		}
		datum, present := props[a.Name]
		if !present || datum == nil {
			switch {
			case ref != nil && slices.Contains(t.WriteOnly, a.Name):
				attrs[a.Name] = *ref // AWS never returns it
			case ref != nil && emptyCollection(*ref):
				attrs[a.Name] = *ref // AWS omits an empty list or map; an empty one was asked for
			}
			continue
		}
		v, ok, err := decodeAttr(t, a, datum, ref)
		if err != nil {
			return nil, err
		}
		if ok {
			attrs[a.Name] = v
		}
	}
	if !t.Global() {
		attrs[t.RegionAttr] = value.String(region, value.SourceProvider)
	}
	return &resource.ResourceState{Type: t.Name, ProviderID: FormatID(t, region, identifier), Attributes: attrs}, nil
}

func emptyCollection(v value.Value) bool {
	switch raw := v.Raw.(type) {
	case []value.Value:
		return len(raw) == 0
	case map[string]value.Value:
		return len(raw) == 0
	}
	return false
}
