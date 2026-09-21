package catalog

import (
	"bytes"
	_ "embed"
	"sync"
)

//go:embed catalog.json.gz
var embedded []byte

var (
	embeddedOnce sync.Once
	embeddedCat  *Catalog
	embeddedErr  error
)

// Embedded is the catalog generated into this build, decoded once.
func Embedded() (*Catalog, error) {
	embeddedOnce.Do(func() {
		embeddedCat, embeddedErr = Read(bytes.NewReader(embedded))
	})
	return embeddedCat, embeddedErr
}
