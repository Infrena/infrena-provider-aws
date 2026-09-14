package awsprov

import (
	"os"
	"path/filepath"
	"testing"
)

// isolateAWS stops the SDK reading anything from the developer's machine: their config and
// credentials files, their profile, EC2 instance metadata (which otherwise costs a second of
// timeouts), and any real endpoint. endpoint may be "" for tests that make no API call.
// Uses t.Setenv, so callers cannot run in parallel. Returns the directory holding the (empty)
// config and credentials files, for a test that wants to write a profile into them.
func isolateAWS(t *testing.T, endpoint string) string {
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
		"AWS_CONFIG_FILE":             cfgFile,
		"AWS_SHARED_CREDENTIALS_FILE": credFile,
		"AWS_PROFILE":                 "",
		"AWS_DEFAULT_PROFILE":         "",
		"AWS_REGION":                  "",
		"AWS_DEFAULT_REGION":          "",
		"AWS_MAX_ATTEMPTS":            "",
		"AWS_ACCESS_KEY_ID":           "AKIDTESTSTATIC",
		"AWS_SECRET_ACCESS_KEY":       "test-secret",
		"AWS_SESSION_TOKEN":           "",
		"AWS_EC2_METADATA_DISABLED":   "true",
		"AWS_ENDPOINT_URL":            "",
		"AWS_ENDPOINT_URL_EC2":        endpoint,
		"AWS_ENDPOINT_URL_STS":        endpoint,
	} {
		t.Setenv(k, v)
	}
	return dir
}
