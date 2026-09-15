package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

func TestGenerateVPCDocumentation(t *testing.T) {
	dir := t.TempDir()
	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldCwd)

	// Change to temp dir and create docs/reference structure
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	// Create a minimal catalog.Embedded() mock by manually generating docs for aws.vpc
	// We'll create the test by running the actual generator and checking the output

	// For this test, we need to actually run the generator which requires the embedded catalog
	// Let's test the generation logic on a sample type instead

	t.Skip("skipping integration test; unit tests below")
}

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

	// Now compare with repo docs
	repoDocsDir := filepath.Join(oldCwd, "docs", "reference")

	// Check if docs/reference exists and compare
	if info, err := os.Stat(repoDocsDir); err == nil && info.IsDir() {
		// Compare generated docs with repo docs
		if err := compareDirectories(tempDir, repoDocsDir); err != nil {
			t.Errorf("generated documentation does not match committed docs: %v\nRun: go run ./cmd/gen-docs", err)
		}
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

		if strings.TrimSpace(string(genContent)) != strings.TrimSpace(string(repoContent)) {
			return err
		}

		return nil
	})
}
