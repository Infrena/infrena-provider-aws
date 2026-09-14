// Package gen turns CloudFormation resource schemas into the plugin's catalog.
package gen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

// Lock records every infrata type name ever assigned. It only grows: a name, once given, keeps its owner forever
// (J3), including after AWS removes the type.
type Lock struct {
	Names map[string]string `json:"names"`
}

// LoadLock reads a lock file; a missing file is an empty lock.
func LoadLock(path string) (*Lock, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return &Lock{Names: map[string]string{}}, nil
	}
	if err != nil {
		return nil, err
	}
	var l Lock
	if err := json.Unmarshal(raw, &l); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if l.Names == nil {
		l.Names = map[string]string{}
	}
	return &l, nil
}

// Save writes the lock with keys sorted (encoding/json sorts map keys) and a trailing newline.
func (l *Lock) Save(path string) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(l); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

func split(cfn string) (service, segment string, err error) {
	parts := strings.Split(cfn, "::")
	if len(parts) != 3 || parts[0] != "AWS" {
		return "", "", fmt.Errorf("%q is not an AWS::Service::Resource type name", cfn)
	}
	return strings.ToLower(parts[1]), strings.ToLower(parts[2]), nil
}

// Assign returns the infrata name for each current type, adding new types to the lock.
func (l *Lock) Assign(cfnTypes []string) (map[string]string, error) {
	if l.Names == nil {
		l.Names = map[string]string{}
	}
	sorted := append([]string(nil), cfnTypes...)
	sort.Strings(sorted)

	segmentCount := map[string]int{}
	for _, cfn := range sorted {
		_, seg, err := split(cfn)
		if err != nil {
			return nil, err
		}
		segmentCount[seg]++
	}
	owner := map[string]string{} // infrata name -> CFN type, across the whole lock, tombstones included
	for cfn, name := range l.Names {
		if other, dup := owner[name]; dup {
			return nil, fmt.Errorf("lock gives %s to both %s and %s", name, other, cfn)
		}
		owner[name] = cfn
	}

	out := make(map[string]string, len(sorted))
	for _, cfn := range sorted {
		if name, locked := l.Names[cfn]; locked {
			out[cfn] = name
			continue
		}
		service, seg, _ := split(cfn)
		name := "aws." + seg
		if _, taken := owner[name]; taken || segmentCount[seg] > 1 {
			name = "aws." + service + "." + seg
		}
		if other, taken := owner[name]; taken {
			return nil, fmt.Errorf("%s would be named %s, which already belongs to %s", cfn, name, other)
		}
		owner[name] = cfn
		l.Names[cfn] = name
		out[cfn] = name
	}
	return out, nil
}
