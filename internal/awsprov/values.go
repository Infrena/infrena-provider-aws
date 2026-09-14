package awsprov

import "github.com/infrata/infrata/pkg/value"

// str and boolean build values as a provider reports them. Sensitivity and provenance are NOT set
// here beyond the source: the host forces both from the schema, and a plugin that also did it would
// have tests that pass when the host is broken.
func str(s string) value.Value   { return value.String(s, value.SourceProvider) }
func boolean(b bool) value.Value { return value.Bool(b, value.SourceProvider) }
