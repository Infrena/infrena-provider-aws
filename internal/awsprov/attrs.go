package awsprov

import (
	"fmt"

	"github.com/infrata/infrata/pkg/value"
)

// stringAttr reads a required string. infrata has already checked Required and Kind at compile time;
// this is the plugin refusing to guess if a caller skipped that.
func stringAttr(attrs map[string]value.Value, name string) (string, error) {
	v, ok := attrs[name]
	if !ok {
		return "", fmt.Errorf("attribute %q is missing", name)
	}
	text, isString := v.AsString()
	if !isString || text == "" {
		return "", fmt.Errorf("attribute %q must be a non-empty string, got %s", name, v.Kind)
	}
	return text, nil
}

// boolAttr reads a bool, absent meaning false (the schema's default).
func boolAttr(attrs map[string]value.Value, name string) (bool, error) {
	v, ok := attrs[name]
	if !ok {
		return false, nil
	}
	b, isBool := v.AsBool()
	if !isBool {
		return false, fmt.Errorf("attribute %q must be a boolean, got %s", name, v.Kind)
	}
	return b, nil
}
