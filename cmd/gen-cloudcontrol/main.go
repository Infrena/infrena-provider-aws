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

	"github.com/infrena/infrena-provider-aws/internal/cfn"
	"github.com/infrena/infrena-provider-aws/internal/gen"
)

func main() {
	bundle := flag.String("bundle", "schemas/CloudformationSchema.zip", "CloudFormation schema bundle")
	overlayPath := flag.String("overlay", "gen/overlay.yaml", "curated overlay")
	lockPath := flag.String("lock", "gen/names.lock.json", "type name lock")
	refsPath := flag.String("references-lock", "gen/references.lock.json", "reference edge lock")
	acceptNew := flag.Bool(strings.TrimPrefix(gen.AcceptNewReferencesFlag, "-"), false, "add derived reference edges the lock does not hold yet")
	out := flag.String("out", "internal/catalog/catalog.json.gz", "catalog to write")
	warningsPath := flag.String("warnings", "gen/warnings.txt", "warnings to write")
	flag.Parse()
	if err := run(*bundle, *overlayPath, *lockPath, *refsPath, *out, *warningsPath, *acceptNew); err != nil {
		fmt.Fprintln(os.Stderr, "gen-cloudcontrol:", err)
		os.Exit(1)
	}
}

func run(bundle, overlayPath, lockPath, refsPath, out, warningsPath string, acceptNew bool) error {
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
	refs, err := gen.LoadReferenceLock(refsPath)
	if err != nil {
		return err
	}
	before, refsBefore := len(lock.Names), len(refs.References)
	cat, warnings, err := gen.Generate(schemas, sha, lock, refs, o, acceptNew)
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
	if err := refs.Save(refsPath); err != nil {
		return err
	}
	if err := os.WriteFile(warningsPath, []byte(strings.Join(warnings, "\n")+"\n"), 0o644); err != nil {
		return err
	}
	statuses := map[string]int{}
	for _, r := range refs.References {
		statuses[r.Status]++
	}
	fmt.Fprintf(os.Stderr, "bundle %s: %d types, %d new names, %d new references (%d accepted, %d approved, %d pending, %d rejected), %d warnings\n",
		sha[:12], len(cat.Types), len(lock.Names)-before, len(refs.References)-refsBefore,
		statuses[gen.StatusAccepted], statuses[gen.StatusApproved], statuses[gen.StatusPending], statuses[gen.StatusRejected], len(warnings))
	return nil
}
