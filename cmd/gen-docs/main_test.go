package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

func TestFormatFlags(t *testing.T) {
	tests := []struct {
		name      string
		required  bool
		computed  bool
		optional  bool
		forceNew  bool
		sensitive bool
		want      string
	}{
		{
			name:     "optional computed",
			optional: true,
			computed: true,
			want:     "optional, computed, provider-chosen",
		},
		{
			name:     "required force new",
			required: true,
			forceNew: true,
			want:     "required, replaces on change",
		},
		{
			name:      "sensitive optional",
			optional:  true,
			sensitive: true,
			want:      "optional, sensitive",
		},
		{
			name: "no flags",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := &catalog.Attribute{
				Required:  tt.required,
				Computed:  tt.computed,
				Optional:  tt.optional,
				ForceNew:  tt.forceNew,
				Sensitive: tt.sensitive,
			}
			got := formatFlags(attr, &catalog.Type{})
			if got != tt.want {
				t.Errorf("formatFlags() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestServiceFromCFN(t *testing.T) {
	tests := []struct {
		cfn  string
		want string
	}{
		{"AWS::EC2::VPC", "ec2"},
		{"AWS::IAM::Role", "iam"},
		{"AWS::S3::Bucket", "s3"},
		{"AWS::DynamoDB::Table", "dynamodb"},
	}

	for _, tt := range tests {
		t.Run(tt.cfn, func(t *testing.T) {
			got := serviceFromCFN(tt.cfn)
			if got != tt.want {
				t.Errorf("serviceFromCFN(%q) = %q, want %q", tt.cfn, got, tt.want)
			}
		})
	}
}

func TestTypeNameToDashCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"aws.vpc", "vpc"},
		{"aws.subnet", "subnet"},
		{"aws.iam.role", "iam-role"},
		{"aws.ec2.security.group", "ec2-security-group"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := typeNameToDashCase(tt.name)
			if got != tt.want {
				t.Errorf("typeNameToDashCase(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestDocumentationIsUpToDate(t *testing.T) {
	// Load the real embedded catalog
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatalf("failed to load embedded catalog: %v", err)
	}

	// Generate into temp dir
	tempDir := t.TempDir()
	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldCwd)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}

	// Write type pages
	for _, typ := range cat.Types {
		if err := writeTypePage(typ); err != nil {
			t.Fatalf("writeTypePage(%s): %v", typ.Name, err)
		}
	}

	// Write index
	if err := writeIndex(cat.Types); err != nil {
		t.Fatalf("writeIndex: %v", err)
	}

	// The committed pages are at the repository root, two levels up. A test's working directory is its own
	// package directory, so joining "docs/reference" onto it named cmd/gen-docs/docs/reference, which has never
	// existed: os.Stat failed, the guard below skipped the comparison in silence, and this test passed however
	// stale the pages were. Not finding them is now a failure rather than a reason to skip.
	repoDocsDir := filepath.Join(oldCwd, "..", "..", "docs", "reference")
	info, err := os.Stat(repoDocsDir)
	if err != nil || !info.IsDir() {
		t.Fatalf("committed reference pages not found at %s (%v); this test compares against them", repoDocsDir, err)
	}
	if err := compareDirectories(tempDir, repoDocsDir); err != nil {
		t.Errorf("generated documentation does not match committed docs: %v\nRun: go run ./cmd/gen-docs", err)
	}
}

func compareDirectories(genDir, repoDir string) error {
	genDocsDir := filepath.Join(genDir, "docs", "reference")

	// Walk the generated directory
	return filepath.Walk(genDocsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(genDocsDir, path)
		if err != nil {
			return err
		}

		repoPath := filepath.Join(repoDir, relPath)
		genContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		repoContent, err := os.ReadFile(repoPath)
		if err != nil {
			if os.IsNotExist(err) {
				return err // File not in repo
			}
			return err
		}

		// This returned err, which is nil here — the walk callback checks it at the top, so a difference was
		// found, discarded, and reported as success. The test was green through every stale-docs commit it was
		// meant to catch, including two on 2026-09-17 that left 23 pages describing the catalog wrongly.
		if strings.TrimSpace(string(genContent)) != strings.TrimSpace(string(repoContent)) {
			return fmt.Errorf("%s differs from the generator's output", relPath)
		}

		return nil
	})
}
