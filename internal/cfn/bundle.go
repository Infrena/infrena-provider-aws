package cfn

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// ReadBundle reads every schema in the CloudFormation schema zip and returns them sorted by type name, with the
// zip's SHA-256.
func ReadBundle(zipPath string) ([]*Schema, string, error) {
	raw, err := os.ReadFile(zipPath)
	if err != nil {
		return nil, "", err
	}
	sum := sha256.Sum256(raw)
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, "", err
	}
	defer zr.Close()
	var out []*Schema
	for _, f := range zr.File {
		if !strings.HasSuffix(f.Name, ".json") {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, "", err
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, "", err
		}
		s, err := Parse(data)
		if err != nil {
			return nil, "", fmt.Errorf("%s: %w", f.Name, err)
		}
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].TypeName < out[j].TypeName })
	return out, hex.EncodeToString(sum[:]), nil
}
