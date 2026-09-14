// Command gen-cloudcontrol generates the plugin's catalog from AWS's CloudFormation schema bundle. Run it by hand
// after scripts/fetch-schemas; commit what it writes.
//
//	go run ./cmd/gen-cloudcontrol
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/cfn"
	"github.com/infrata/infrata-provider-aws/internal/gen"
)

func main() {
	bundle := flag.String("bundle", "schemas/CloudformationSchema.zip", "CloudFormation schema bundle")
	overlayPath := flag.String("overlay", "gen/overlay.yaml", "curated overlay")
	lockPath := flag.String("lock", "gen/names.lock.json", "type name lock")
	out := flag.String("out", "internal/catalog/catalog.json.gz", "catalog to write")
	warningsPath := flag.String("warnings", "gen/warnings.txt", "warnings to write")
	flag.Parse()
	if err := run(*bundle, *overlayPath, *lockPath, *out, *warningsPath); err != nil {
		fmt.Fprintln(os.Stderr, "gen-cloudcontrol:", err)
		os.Exit(1)
	}
}

func run(bundle, overlayPath, lockPath, out, warningsPath string) error {
	schemas, sha, err := cfn.ReadBundle(bundle)
	if err != nil {
		return err
	}
	o, err := gen.LoadOverlay(overlayPath)
	if err != nil {
		return err
	}
	lock, err := gen.LoadLock(lockPath)
	if err != nil {
		return err
	}
	before := len(lock.Names)
	cat, warnings, err := gen.Generate(schemas, sha, lock, o)
	if err != nil {
		return err
	}
	f, err := os.Create(out)
	if err != nil {
		return err
	}
	if err := cat.Write(f); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := lock.Save(lockPath); err != nil {
		return err
	}
	if err := os.WriteFile(warningsPath, []byte(strings.Join(warnings, "\n")+"\n"), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "bundle %s: %d types, %d new names, %d warnings\n", sha[:12], len(cat.Types), len(lock.Names)-before, len(warnings))
	return nil
}
