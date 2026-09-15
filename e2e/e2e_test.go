//go:build e2e

// Package e2e runs a real infrena binary against a real infrena-plugin-aws binary and an in-process fake Cloud Control.
//
// Not part of `go test ./...`: it builds infrena from source. Run it with `go test -tags e2e -count=1 -v ./e2e/`.
// INFRENA_SRC points at the checkout; the default is the sibling ../infrena.
package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/pluginmanifest"
)

var (
	infrenaBin string
	pluginDir  string
	skipReason string
)

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "aws-e2e-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := func() int {
		defer os.RemoveAll(tmp)
		src := os.Getenv("INFRENA_SRC")
		if src == "" {
			src = filepath.Join("..", "..", "infrena")
		}
		if _, err := os.Stat(filepath.Join(src, "cmd", "infrena")); err != nil {
			skipReason = fmt.Sprintf("no infrena checkout at %s (set INFRENA_SRC): %v", src, err)
			fmt.Fprintln(os.Stderr, "E2E SKIPPED: "+skipReason)
			return m.Run()
		}
		infrenaBin = filepath.Join(tmp, "infrena")
		pluginDir = filepath.Join(tmp, "plugins")
		for _, b := range []struct{ dir, out, pkg string }{
			{src, infrenaBin, "./cmd/infrena"},
			{"..", filepath.Join(pluginDir, "infrena-plugin-aws"), "./cmd/infrena-plugin-aws"},
		} {
			cmd := exec.Command("go", "build", "-o", b.out, b.pkg)
			cmd.Dir = b.dir
			if out, err := cmd.CombinedOutput(); err != nil {
				fmt.Fprintf(os.Stderr, "building %s: %v\n%s", b.pkg, err, out)
				return 1
			}
		}
		return m.Run()
	}()
	os.Exit(code)
}

// env is one project and the account it talks to.
type env struct {
	dir  string
	fake *ccfake.Server
}

func project(t *testing.T, body string) *env {
	t.Helper()
	if skipReason != "" {
		t.Skip(skipReason)
	}
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	e := &env{dir: t.TempDir(), fake: ccfake.New()}
	t.Cleanup(e.fake.Close)
	awstest.RegisterCore(t, e.fake, cat)
	// AWS reports ingress rules in its own order, with a Description nobody set.
	sg := awstest.FakeType(awstest.TypeFor(t, cat, "AWS::EC2::SecurityGroup"), "sg-", map[string]any{
		"SecurityGroupEgress": []any{map[string]any{"IpProtocol": "-1", "CidrIp": "0.0.0.0/0"}},
	})
	sg.OnRead = func(props map[string]any) {
		rules, _ := props["SecurityGroupIngress"].([]any)
		for _, r := range rules {
			r.(map[string]any)["Description"] = ""
		}
		slices.Reverse(rules)
	}
	e.fake.Register(sg)
	writeFile(t, filepath.Join(e.dir, "infra.yml"), body)
	return e
}

func fixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name, "infra.yml"))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// infrena runs the CLI in the project. Nothing from the developer's machine reaches the plugin: no plugins but ours,
// no AWS files, static test keys, no instance metadata, and Cloud Control and STS pointed at this project's fake.
func (e *env) infrena(t *testing.T, args ...string) (string, int) {
	t.Helper()
	home := t.TempDir()
	cmd := exec.Command(infrenaBin, append(args, "--plugin-dir", pluginDir)...)
	cmd.Dir = e.dir
	cmd.Env = append(os.Environ(),
		"INFRENA_PLUGIN_PATH=", "HOME="+home,
		"AWS_CONFIG_FILE="+filepath.Join(home, "none"), "AWS_SHARED_CREDENTIALS_FILE="+filepath.Join(home, "none"),
		"AWS_PROFILE=", "AWS_REGION=", "AWS_DEFAULT_REGION=", "AWS_MAX_ATTEMPTS=", "AWS_SESSION_TOKEN=",
		"AWS_ACCESS_KEY_ID=AKIDE2E", "AWS_SECRET_ACCESS_KEY=e2e-secret", "AWS_EC2_METADATA_DISABLED=true",
		"AWS_ENDPOINT_URL=", "AWS_ENDPOINT_URL_CLOUDCONTROL="+e.fake.URL, "AWS_ENDPOINT_URL_STS="+e.fake.URL,
	)
	out, err := cmd.CombinedOutput()
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("running infrena %v: %v", args, err)
	}
	return string(out), code
}

func (e *env) expect(t *testing.T, wantCode int, want []string, args ...string) string {
	t.Helper()
	out, code := e.infrena(t, args...)
	if code != wantCode {
		t.Fatalf("infrena %s: exit %d, want %d\n%s", strings.Join(args, " "), code, wantCode, out)
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("infrena %s: output lacks %q\n%s", strings.Join(args, " "), w, out)
		}
	}
	return out
}

// planOps runs `plan dev --output` and returns address → kind for every proposed operation, "noop" excluded.
func (e *env) planOps(t *testing.T) map[string]string {
	t.Helper()
	outPath := filepath.Join(t.TempDir(), "plan.json")
	e.infrena(t, "plan", "dev", "--output", outPath)
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("plan wrote no --output file: %v", err)
	}
	var doc struct {
		Operations []struct{ Address, Kind string } `json:"operations"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("plan --output is not the expected JSON: %v\n%s", err, data)
	}
	ops := map[string]string{}
	for _, op := range doc.Operations {
		if op.Kind != "noop" {
			ops[op.Address] = op.Kind
		}
	}
	return ops
}

func (e *env) editConfig(t *testing.T, from, to string) {
	t.Helper()
	path := filepath.Join(e.dir, "infra.yml")
	body, _ := os.ReadFile(path)
	if !strings.Contains(string(body), from) {
		t.Fatalf("infra.yml does not contain %q", from)
	}
	writeFile(t, path, strings.Replace(string(body), from, to, 1))
}

// only returns the one resource of a type the fake holds in us-east-1.
func (e *env) only(t *testing.T, cfn string) (string, map[string]any) {
	t.Helper()
	all := e.fake.Resources("us-east-1", cfn)
	if len(all) != 1 {
		t.Fatalf("the fake holds %d %s, want 1", len(all), cfn)
	}
	for id, props := range all {
		return id, props
	}
	return "", nil
}

// drift changes a resource behind infrena's back.
func (e *env) drift(t *testing.T, cfn, id string, change func(props map[string]any)) {
	t.Helper()
	props, ok := e.fake.Resource("us-east-1", cfn, id)
	if !ok {
		t.Fatalf("no %s %s", cfn, id)
	}
	change(props)
	e.fake.Put("us-east-1", cfn, id, props)
}

// TestTheWorkflow is AGENT.md's checklist against a VPC, a subnet and a security group.
func TestTheWorkflow(t *testing.T) {
	e := project(t, fixture(t, "basic"))
	var vpcID string
	managed := func(t *testing.T) string {
		if vpcID == "" {
			t.Fatal("no managed VPC yet: the apply subtest did not run or failed")
		}
		return vpcID
	}

	t.Run("explain shows every spelling", func(t *testing.T) {
		e.expect(t, 0, []string{"aws.vpc", "also cidr", "(replaces on change", "Optional, chosen by the provider if unset", "<region>/<identifier>"},
			"explain", "aws.vpc")
	})
	t.Run("plan proposes three creates in the friendly names", func(t *testing.T) {
		out := e.expect(t, 2, []string{"Plan: 3 to create", "cidr", "ingress"}, "plan", "dev")
		if strings.Contains(out, "CidrBlock") {
			t.Errorf("the plan shows AWS's name instead of the alias:\n%s", out)
		}
	})
	t.Run("apply creates them, and a re-plan is clean", func(t *testing.T) {
		e.expect(t, 2, []string{"Apply complete: 3 applied, 0 failed"}, "apply", "dev", "--auto-approve")
		vpcID, _ = e.only(t, "AWS::EC2::VPC")
		if _, subnet := e.only(t, "AWS::EC2::Subnet"); subnet["VpcId"] != vpcID {
			t.Fatalf("the subnet's VpcId = %v, want %s: ${vpc} did not project to the VPC's VpcId", subnet["VpcId"], vpcID)
		}
		if _, sg := e.only(t, "AWS::EC2::SecurityGroup"); sg["SecurityGroupIngress"].([]any)[0].(map[string]any)["FromPort"] == nil {
			t.Fatalf("ingress reached AWS as %v, want AWS's names", sg["SecurityGroupIngress"])
		}
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("re-plan after apply proposes %v: a value read back differs from what was configured", ops)
		}
	})
	t.Run("a value AWS chose that changes is not a diff", func(t *testing.T) {
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) { p["EnableDnsSupport"] = false })
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("an unset provider-chosen attribute plans %v, want nothing (PLAN §14.1)", ops)
		}
	})
	t.Run("a tag changed outside infrena is an update", func(t *testing.T) {
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) {
			p["Tags"] = []any{map[string]any{"Key": "team", "Value": "someone-else"}}
		})
		if kind := e.planOps(t)["vpc"]; kind != "update" {
			t.Fatalf("vpc plans as %q, want update", kind)
		}
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
		if _, vpc := e.only(t, "AWS::EC2::VPC"); vpc["Tags"].([]any)[0].(map[string]any)["Value"] != "platform" {
			t.Fatalf("tags after apply = %v", vpc["Tags"])
		}
	})
	t.Run("a setting changed in configuration is patched in place", func(t *testing.T) {
		e.editConfig(t, "    cidr: 10.0.0.0/16\n", "    cidr: 10.0.0.0/16\n    enable_dns_hostnames: true\n")
		if kind := e.planOps(t)["vpc"]; kind != "update" {
			t.Fatalf("vpc plans as %q, want update", kind)
		}
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
		if _, vpc := e.only(t, "AWS::EC2::VPC"); vpc["EnableDnsHostnames"] != true {
			t.Fatalf("EnableDnsHostnames = %v", vpc["EnableDnsHostnames"])
		}
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after the update proposes %v", ops)
		}
	})
	t.Run("a create-only value changed outside infrena forces a replacement", func(t *testing.T) {
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) { p["CidrBlock"] = "10.50.0.0/16" })
		if kind := e.planOps(t)["vpc"]; kind != "replace" {
			t.Fatalf("vpc plans as %q, want replace", kind)
		}
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) { p["CidrBlock"] = "10.0.0.0/16" })
	})
	t.Run("discover and import adopt a VPC infrena did not create", func(t *testing.T) {
		e.fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-legacy", map[string]any{
			"VpcId": "vpc-legacy", "CidrBlock": "172.16.0.0/16", "EnableDnsSupport": true, "EnableDnsHostnames": false, "InstanceTenancy": "default",
		})
		e.expect(t, 0, []string{"aws.vpc", "us-east-1/vpc-legacy"}, "discover")
		e.expect(t, 0, []string{"1 resource imported"}, "import", "dev", "aws.vpc.us-east-1/vpc-legacy", "--generate")
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after import --generate proposes %v", ops)
		}
	})
	t.Run("destroy removes everything infrena manages", func(t *testing.T) {
		e.expect(t, 2, []string{"0 failed"}, "destroy", "dev", "--auto-approve")
		for _, cfn := range []string{"AWS::EC2::VPC", "AWS::EC2::Subnet", "AWS::EC2::SecurityGroup"} {
			if n := len(e.fake.Resources("us-east-1", cfn)); n != 0 {
				t.Errorf("after destroy the fake holds %d %s", n, cfn)
			}
		}
	})
}

// TestAnUnknownNestedKeyFailsTheApplyNamingIt. The plan cannot check nested keys; the apply must, before any call.
func TestAnUnknownNestedKeyFailsTheApplyNamingIt(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "        from_port: 443\n", "        port: 443\n", 1))
	e.expect(t, 1, []string{`"port"`, "from_port"}, "apply", "dev", "--auto-approve")
	if n := len(e.fake.Resources("us-east-1", "AWS::EC2::SecurityGroup")); n != 0 {
		t.Errorf("the fake holds %d security groups: the create was sent", n)
	}
}

// TestAWholeResourceIntoTheCanonicalAttributeNameStillWorks: VpcId: ${vpc}, AWS's own property name, worked even
// on infrena 0.6.0 and 0.6.1, which had a compiler bug refusing the same reference written through an alias
// (vpc: ${vpc}, vpc_id: ${vpc}, both used elsewhere in this fixture). This confirms the canonical spelling is
// still accepted alongside the alias spellings the fixture now exercises.
func TestAWholeResourceIntoTheCanonicalAttributeNameStillWorks(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "    vpc: ${vpc}\n", "    VpcId: ${vpc}\n", 1))
	e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
	vpcID, _ := e.only(t, "AWS::EC2::VPC")
	if _, subnet := e.only(t, "AWS::EC2::Subnet"); subnet["VpcId"] != vpcID {
		t.Fatalf("the subnet's VpcId = %v, want %s: VpcId: ${vpc} did not project to the VPC's VpcId", subnet["VpcId"], vpcID)
	}
}

// TestAWholeResourceIntoAnUndeclaredAttributeNamesTheFix: CidrBlock declares no reference, so ${vpc} there is a
// compile error telling the user to name the attribute, and nothing is created.
func TestAWholeResourceIntoAnUndeclaredAttributeNamesTheFix(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "    cidr: 10.0.1.0/24\n", "    CidrBlock: ${vpc}\n", 1))
	e.expect(t, 1, []string{"declares no reference", "Name the attribute you mean"}, "apply", "dev", "--auto-approve")
	if n := len(e.fake.Resources("us-east-1", "AWS::EC2::VPC")); n != 0 {
		t.Errorf("the fake holds %d VPCs: a configuration that does not compile was applied", n)
	}
}

// TestAListOfWholeResourcesProjectsEachItem: DBSubnetGroup.SubnetIds declares aws.subnet's SubnetId, and infrena
// projects each ${subnet} in the list.
func TestAListOfWholeResourcesProjectsEachItem(t *testing.T) {
	e := project(t, fixture(t, "basic")+`
  db_subnets:
    type: aws.rds.dbsubnetgroup
    description: databases
    SubnetIds:
      - ${private_a}
`)
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	e.fake.Register(awstest.FakeType(awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup"), "dbsubnet-", nil))
	e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
	subnetID, _ := e.only(t, "AWS::EC2::Subnet")
	if _, group := e.only(t, "AWS::RDS::DBSubnetGroup"); fmt.Sprint(group["SubnetIds"]) != fmt.Sprint([]any{subnetID}) {
		t.Fatalf("SubnetIds = %v, want [%s]: ${private_a} in a list did not project", group["SubnetIds"], subnetID)
	}
}

func TestAMisspelledKeyIsRefusedAgainstTheProvidersEntry(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "discover_regions:", "discover_region:", 1))
	e.expect(t, 1, []string{`unknown configuration "discover_region"`}, "plan", "dev")
}

func TestAnUnknownDiscoverTypeIsRefused(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "aws.securitygroup]", "aws.securitygroupp]", 1))
	e.expect(t, 1, []string{"aws.securitygroupp", "infrena explain"}, "plan", "dev")
}

// TestTheInfrenaUnderTestSpeaksTheManifestsProtocol applies infrena PLAN.md §31.2's compatibility rules to the
// infrena this suite built, reading what that build says it speaks from `infrena version --output`.
func TestTheInfrenaUnderTestSpeaksTheManifestsProtocol(t *testing.T) {
	if skipReason != "" {
		t.Skip(skipReason)
	}
	out := filepath.Join(t.TempDir(), "version.json")
	if b, err := exec.Command(infrenaBin, "version", "--output", out).CombinedOutput(); err != nil {
		t.Fatalf("infrena version --output: %v\n%s", err, b)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	var info struct {
		Version string `json:"version"`
		Formats []struct {
			Name     string `json:"name"`
			Versions []int  `json:"versions"`
		} `json:"formats"`
	}
	if err := json.Unmarshal(data, &info); err != nil {
		t.Fatalf("infrena version --output is not the expected JSON: %v\n%s", err, data)
	}
	var protocols []int
	for _, f := range info.Formats {
		if f.Name == "plugin protocol" {
			protocols = f.Versions
		}
	}
	if len(protocols) == 0 {
		t.Fatalf("infrena version --output lists no plugin protocol:\n%s", data)
	}

	raw, err := os.ReadFile(filepath.Join("..", "plugin.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	m, _, err := pluginmanifest.Parse(raw)
	if err != nil {
		t.Fatalf("plugin.yaml: %v", err)
	}
	if !m.SpeaksProtocol(protocols) {
		t.Errorf("plugin.yaml speaks protocol %v; infrena %s speaks %v", m.Protocol, info.Version, protocols)
	}
	if !m.AllowsInfrena(info.Version) {
		t.Errorf("plugin.yaml's infrena constraint %q does not allow infrena %s", m.Infrena, info.Version)
	}
}
