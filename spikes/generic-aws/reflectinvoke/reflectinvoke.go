// Package reflectinvoke is THROWAWAY SPIKE CODE for approach B: call any operation on a generated AWS SDK
// client by NAME, with a generic map as input and a generic map as output, using Go reflection over the
// generated types. It shows what reflection can and cannot do; it is not a recommendation.
package reflectinvoke

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

// Invoke calls client.<operation>(ctx, input) where input is built from a generic map.
//
// The map is converted through encoding/json into the generated *<Op>Input struct. That works because the
// generated structs have exported fields named exactly as the Smithy members, with no JSON tags, and
// encoding/json matches field names case-insensitively. Enums are string types, timestamps *time.Time,
// numbers *int32 etc., so JSON's coercions mostly line up.
func Invoke(ctx context.Context, client any, operation string, input map[string]any) (map[string]any, error) {
	method := reflect.ValueOf(client).MethodByName(operation)
	if !method.IsValid() {
		return nil, fmt.Errorf("%T has no operation %q", client, operation)
	}
	mt := method.Type()
	// Generated operations are func(ctx, *XInput, ...func(*Options)) (*XOutput, error).
	if mt.NumIn() != 3 || mt.In(1).Kind() != reflect.Pointer || mt.NumOut() != 2 {
		return nil, fmt.Errorf("%s does not have the generated operation shape: %s", operation, mt)
	}
	in := reflect.New(mt.In(1).Elem())
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	// DisallowUnknownFields would be the strict choice; it is left out to show the silent-drop hazard.
	if err := json.Unmarshal(raw, in.Interface()); err != nil {
		return nil, fmt.Errorf("input does not fit %s: %w", mt.In(1).Elem(), err)
	}
	outs := method.Call([]reflect.Value{reflect.ValueOf(ctx), in})
	if errV := outs[1]; !errV.IsNil() {
		return nil, errV.Interface().(error)
	}
	// The output carries ResultMetadata (middleware.Metadata) and unexported noSmithyDocumentSerde
	// fields; JSON skips unexported ones.
	outRaw, err := json.Marshal(outs[0].Interface())
	if err != nil {
		return nil, fmt.Errorf("output does not marshal: %w", err)
	}
	var out map[string]any
	if err := json.Unmarshal(outRaw, &out); err != nil {
		return nil, err
	}
	delete(out, "ResultMetadata")
	return out, nil
}
