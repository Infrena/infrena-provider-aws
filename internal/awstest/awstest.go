// Package awstest isolates the AWS SDK from the developer's machine in tests. It is imported only by _test files.
package awstest

import (
	"os"
	"path/filepath"
	"testing"
)

// Isolate stops the SDK reading the developer's config and credentials files, profile, instance metadata and real
// endpoints, and points Cloud Control and STS at endpoint (which may be ""). It uses t.Setenv, so callers cannot run in
// parallel. It returns the directory holding the empty config and credentials files.
func Isolate(t testing.TB, endpoint string) string {
	t.Helper()
	dir := t.TempDir()
	cfgFile := filepath.Join(dir, "config")
	credFile := filepath.Join(dir, "credentials")
	for _, f := range []string{cfgFile, credFile} {
		if err := os.WriteFile(f, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	for k, v := range map[string]string{
		"AWS_CONFIG_FILE":               cfgFile,
		"AWS_SHARED_CREDENTIALS_FILE":   credFile,
		"AWS_PROFILE":                   "",
		"AWS_DEFAULT_PROFILE":           "",
		"AWS_REGION":                    "",
		"AWS_DEFAULT_REGION":            "",
		"AWS_MAX_ATTEMPTS":              "",
		"AWS_ACCESS_KEY_ID":             "AKIDTESTSTATIC",
		"AWS_SECRET_ACCESS_KEY":         "test-secret",
		"AWS_SESSION_TOKEN":             "",
		"AWS_EC2_METADATA_DISABLED":     "true",
		"AWS_ENDPOINT_URL":              "",
		"AWS_ENDPOINT_URL_CLOUDCONTROL": endpoint,
		"AWS_ENDPOINT_URL_STS":          endpoint,
	} {
		t.Setenv(k, v)
	}
	return dir
}
