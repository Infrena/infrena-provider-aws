package smithyinvoke

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/aws/smithy-go"
)

// Validate checks a generic value against a runtime schema before anything is sent, so a misspelt
// member is an error rather than a silently dropped field (the reflection approach's hazard).
func Validate(s *smithy.Schema, v any, path string) error {
	if v == nil {
		return nil
	}
	switch s.Type() {
	case smithy.ShapeTypeStructure, smithy.ShapeTypeUnion:
		obj, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("%s: want an object, got %T", path, v)
		}
		for k, item := range obj {
			m := s.Member(k)
			if m == nil {
				return fmt.Errorf("%s: %s has no member %q (members: %v)", path, s.ID().Name, k, memberNames(s))
			}
			if err := Validate(m, item, path+"."+k); err != nil {
				return err
			}
		}
	case smithy.ShapeTypeList, smithy.ShapeTypeSet:
		arr, ok := v.([]any)
		if !ok {
			return fmt.Errorf("%s: want a list, got %T", path, v)
		}
		for i, item := range arr {
			if err := Validate(s.ListMember(), item, path+"["+strconv.Itoa(i)+"]"); err != nil {
				return err
			}
		}
	case smithy.ShapeTypeMap:
		obj, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("%s: want a map, got %T", path, v)
		}
		for k, item := range obj {
			if err := Validate(s.MapValue(), item, path+"."+k); err != nil {
				return err
			}
		}
	case smithy.ShapeTypeString, smithy.ShapeTypeEnum:
		if _, ok := v.(string); !ok {
			return fmt.Errorf("%s: want a string, got %T", path, v)
		}
	case smithy.ShapeTypeBoolean:
		if _, ok := v.(bool); !ok {
			return fmt.Errorf("%s: want a boolean, got %T", path, v)
		}
	case smithy.ShapeTypeInteger, smithy.ShapeTypeLong, smithy.ShapeTypeShort, smithy.ShapeTypeByte, smithy.ShapeTypeIntEnum:
		if _, ok := toInt(v); !ok {
			return fmt.Errorf("%s: want an integer, got %T", path, v)
		}
	case smithy.ShapeTypeFloat, smithy.ShapeTypeDouble:
		if _, ok := toFloat(v); !ok {
			return fmt.Errorf("%s: want a number, got %T", path, v)
		}
	default:
		return fmt.Errorf("%s: shape type %d is not supported by this spike", path, s.Type())
	}
	return nil
}

// input is a generic value that serializes itself against a runtime schema.
type input struct {
	schema *smithy.Schema
	value  map[string]any
}

func (in input) Serialize(ss smithy.ShapeSerializer) {
	ss.WriteStruct(in.schema)
	writeMembers(ss, in.schema, in.value)
	ss.CloseStruct()
}

func writeMembers(ss smithy.ShapeSerializer, s *smithy.Schema, obj map[string]any) {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		writeValue(ss, s.Member(k), obj[k])
	}
}

func writeValue(ss smithy.ShapeSerializer, s *smithy.Schema, v any) {
	if v == nil {
		return
	}
	switch s.Type() {
	case smithy.ShapeTypeStructure, smithy.ShapeTypeUnion:
		ss.WriteStruct(s)
		writeMembers(ss, s, v.(map[string]any))
		ss.CloseStruct()
	case smithy.ShapeTypeList, smithy.ShapeTypeSet:
		ss.WriteList(s)
		for _, item := range v.([]any) {
			writeValue(ss, s.ListMember(), item)
		}
		ss.CloseList()
	case smithy.ShapeTypeMap:
		ss.WriteMap(s)
		for k, item := range v.(map[string]any) {
			ss.WriteKey(s.MapKey(), k)
			writeValue(ss, s.MapValue(), item)
		}
		ss.CloseMap()
	case smithy.ShapeTypeString, smithy.ShapeTypeEnum:
		ss.WriteString(s, v.(string))
	case smithy.ShapeTypeBoolean:
		ss.WriteBool(s, v.(bool))
	case smithy.ShapeTypeInteger, smithy.ShapeTypeIntEnum:
		n, _ := toInt(v)
		ss.WriteInt32(s, int32(n))
	case smithy.ShapeTypeLong:
		n, _ := toInt(v)
		ss.WriteInt64(s, n)
	case smithy.ShapeTypeShort:
		n, _ := toInt(v)
		ss.WriteInt16(s, int16(n))
	case smithy.ShapeTypeByte:
		n, _ := toInt(v)
		ss.WriteInt8(s, int8(n))
	case smithy.ShapeTypeFloat:
		f, _ := toFloat(v)
		ss.WriteFloat32(s, float32(f))
	case smithy.ShapeTypeDouble:
		f, _ := toFloat(v)
		ss.WriteFloat64(s, f)
	}
}

// output is a generic value that deserializes itself against a runtime schema.
type output struct {
	schema *smithy.Schema
	value  map[string]any
}

func (out *output) Deserialize(d smithy.ShapeDeserializer) error {
	v, err := readStruct(d, out.schema)
	if err != nil {
		return err
	}
	out.value = v
	return nil
}

func readStruct(d smithy.ShapeDeserializer, s *smithy.Schema) (map[string]any, error) {
	obj := map[string]any{}
	err := smithy.ReadStruct(d, s, func(m *smithy.Schema) error {
		v, err := readValue(d, m)
		if err != nil {
			return fmt.Errorf("%s.%s: %w", s.ID().Name, m.MemberName(), err)
		}
		obj[m.MemberName()] = v
		return nil
	})
	return obj, err
}

func readValue(d smithy.ShapeDeserializer, s *smithy.Schema) (any, error) {
	switch s.Type() {
	case smithy.ShapeTypeStructure, smithy.ShapeTypeUnion:
		return readStruct(d, s)
	case smithy.ShapeTypeList, smithy.ShapeTypeSet:
		arr := []any{}
		err := smithy.ReadList(d, s, func() error {
			v, err := readValue(d, s.ListMember())
			arr = append(arr, v)
			return err
		})
		return arr, err
	case smithy.ShapeTypeMap:
		obj := map[string]any{}
		err := smithy.ReadMap(d, s, func(k string) error {
			v, err := readValue(d, s.MapValue())
			obj[k] = v
			return err
		})
		return obj, err
	case smithy.ShapeTypeString, smithy.ShapeTypeEnum:
		var v string
		return v, d.ReadString(s, &v)
	case smithy.ShapeTypeBoolean:
		var v bool
		return v, d.ReadBool(s, &v)
	case smithy.ShapeTypeInteger, smithy.ShapeTypeIntEnum:
		var v int32
		return int64(v), d.ReadInt32(s, &v)
	case smithy.ShapeTypeLong:
		var v int64
		return v, d.ReadInt64(s, &v)
	case smithy.ShapeTypeDouble:
		var v float64
		return v, d.ReadFloat64(s, &v)
	default:
		return nil, fmt.Errorf("shape type %d is not supported by this spike", s.Type())
	}
}

func memberNames(s *smithy.Schema) []string {
	var out []string
	for k := range s.Members() {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func toInt(v any) (int64, bool) {
	switch n := v.(type) {
	case int:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	case float64:
		return int64(n), n == float64(int64(n))
	}
	return 0, false
}

func toFloat(v any) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int:
		return float64(n), true
	case int64:
		return float64(n), true
	}
	return 0, false
}
