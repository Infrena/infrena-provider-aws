package ccprov

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/value"
)

// systemTagPrefix marks tags AWS sets itself. Configuration cannot set them, so they are never reported.
const systemTagPrefix = "aws:"

// tagsToJSON sends a tag map as the [{Key, Value}] list every such schema uses, sorted so a request is stable.
func tagsToJSON(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error) {
	items, ok := v.Raw.(map[string]value.Value)
	if v.Kind != value.KindMap || !ok {
		return nil, fmt.Errorf("%s.%s must be a map of tag names to values, e.g. {team: platform}, got a %s", t.Name, a.Name, v.Kind)
	}
	out := make([]any, 0, len(items))
	for _, k := range sortedKeys(items) {
		text, isString := items[k].AsString()
		if !isString {
			raw, err := json.Marshal(plain(items[k]))
			if err != nil {
				return nil, err
			}
			text = string(raw)
		}
		out = append(out, map[string]any{"Key": k, "Value": text})
	}
	return out, nil
}

// tagsFromJSON reports AWS's tag list as a map.
func tagsFromJSON(t *catalog.Type, a *catalog.Attribute, datum any, ref *value.Value) (value.Value, bool, error) {
	if datum == nil {
		return value.Value{}, false, nil
	}
	items, ok := datum.([]any)
	if !ok {
		return value.Value{}, false, fmt.Errorf("AWS returned %s.%s as %T, not a list of {Key, Value}", t.Name, a.Name, datum)
	}
	refItems, _ := refMap(ref)
	out := make(map[string]value.Value, len(items))
	for _, item := range items {
		tag, _ := item.(map[string]any)
		key, isString := tag["Key"].(string)
		if !isString {
			return value.Value{}, false, fmt.Errorf("AWS returned a tag without a Key in %s.%s", t.Name, a.Name)
		}
		if strings.HasPrefix(key, systemTagPrefix) {
			continue
		}
		if r, has := refItems[key]; has && sameScalar(r, tag["Value"]) {
			out[key] = r
			continue
		}
		text, _ := tag["Value"].(string)
		out[key] = value.String(text, value.SourceProvider)
	}
	return value.Map(out, value.SourceProvider), true, nil
}
