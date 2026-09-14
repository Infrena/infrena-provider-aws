# Generic Cloud Control provider Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the handwritten `aws.vpc`/`aws.subnet` with one generic provider that manages every AWS resource type
AWS Cloud Control API supports, generated from the published CloudFormation resource schemas.

**Architecture:** A build-time generator (`cmd/gen-cloudcontrol`) reads the pinned schema bundle, a committed name lock
and a curated overlay, and writes an embedded catalog: infrata resource definitions plus the runtime metadata the
provider needs. One provider (`internal/ccprov`) serves every catalog type through Cloud Control, reconciling nested
values so plans converge. Tests run against an in-process Cloud Control fake (`internal/ccfake`); a live suite runs by
hand against a real account.

**Tech Stack:** Go 1.27.0; `github.com/infrata/infrata` **v0.3.0** (plugin protocol 2); AWS SDK for Go v2:
`aws-sdk-go-v2 v1.47.0`, `config v1.33.4`, `credentials v1.20.4`, `service/sts v1.50.0`, `service/cloudcontrol v1.38.0`,
`smithy-go v1.28.1`; `gopkg.in/yaml.v3 v3.0.1` (already in infrata's module graph) for the overlay.

**Spec:** `docs/specs/2026-09-14-generic-cloudcontrol-provider.md` (read its §1 decisions J1 to J10 and §3 before any task).
Evidence: `docs/investigations/2026-09-13-generic-aws-provider.md`.

## Global Constraints

- Module `github.com/infrata/infrata-provider-aws`; branch `generic-cloudcontrol`.
- `go 1.27.0`; `require github.com/infrata/infrata v0.3.0`, no `replace`; local work through a gitignored `go.work`.
- Plugin name `aws`; binary `infrata-plugin-aws`; every type prefixed `aws.`.
- `plugin.yaml`: `protocol: [2]`, `infrata: ">= 0.3.0"`.
- `Version` defaults to `"0.0.0-dev"`, stamped only by `-ldflags -X` at release.
- Type names follow J2/J3 and never change once in `gen/names.lock.json`.
- Attribute canonical name is AWS's property name; `Aliases` lists curated aliases first, then snake_case (J4, J7).
- Nothing writes to stdout. No credential, signed request or sensitive value is ever logged.
- `Create`/`Update` never return `(nil, nil)`; `Create` never returns an error once AWS created something.
- Do not reimplement host rules (sensitivity forcing, provenance, bookkeeping, undeclared-attribute refusal, alias
  canonicalisation).
- No test reaches real AWS outside `-tags live`.
- Every test command uses `-count=1`. Every test is sabotage-verified; the sabotage goes in the commit message.
- Stage explicit paths only. Never `git add -A`, `git add .` or `git commit -am`. Ask James before any push or tag.
- **Commit messages (James's rule):** plain English, short, human-sounding, no em-dashes, and no AI or Claude
  attribution of any kind.
- Do not modify `../infrata`. Defects there go to the vault note's follow-ups and to James.

---

## Plan-level decisions (not in the spec)

| # | Decision | Why |
| --- | --- | --- |
| P1 | The generic provider package is `internal/ccprov`, not `internal/cloudcontrol` | The SDK package is `cloudcontrol`; the same name would shadow it at every import site. |
| P2 | Shared runtime pieces move into `ccprov`: clients, error classification, provider IDs, patience | The provider needs them and `awsprov` must import `ccprov`; keeping them in `awsprov` would make an import cycle. `awsprov` keeps only the plugin, instance configuration and credentials. |
| P3 | Provider IDs are `<scope>/<identifier>`, where scope is the region, or `global` for global types | Identifiers can be ARNs with `/`, so the split is at the first `/` only; global types have no region. |
| P4 | New instance key **`discover_types`** (list of infrata type names). Default: the overlay's `discover_default` list (the core set, J10) | infrata's discovery, and every `import`, asks every type an instance offers. Unfiltered that is ~1,584 `ListResources` calls per region per run. Found while planning; not in the spec. |
| P5 | Reconciliation (spec §3.4) treats a nested value as **opaque** (copied exactly, no key translation, nothing dropped, no reordering) when its schema node has `patternProperties`, `additionalProperties` without `properties`, no `properties` at all, `oneOf`/`anyOf`/`allOf`, or a multi-type `type` | The bundle has 447 free-form maps, 323 opaque objects, 384 combinators and 74 type unions; translating or pruning their keys would corrupt user data. |
| P6 | The catalog stores each attribute's nested **shape** (object properties, array item, `insertionOrder`) resolved from `$ref` | Reconciliation needs nested names and ordering at runtime, offline. Its size is measured in Task 6. |
| P7 | Load-cost gate (Task 6): acceptable if the median extra time `infrata validate` takes with the full catalog is at most 500 ms over the fake plugin. Above that, stop and report to James with the numbers | The spec says measure first; this makes "acceptable" checkable. |
| P8 | The first real catalog is generated in Task 5 and committed; regenerating is a normal reviewed diff | Generated code is reviewed as a diff, never produced at build time. |

## Verification log (facts the tasks rely on, checked 2026-09-14)

| Claim | Checked against | Result |
| --- | --- | --- |
| infrata v0.3.0 is `45deb30`, protocol `Version = 2`, `Supported = {2, 1}` | `git show v0.3.0:pkg/pluginproto/proto.go` | ✔ |
| `schema.Attribute` has `Optional` and `Aliases`; `Validate` refuses `Optional` without `Computed` and names that fold together | `v0.3.0:pkg/schema/attribute.go`, `definition.go:94-103`, `alias.go` | ✔ |
| Plans and `import --generate` show `Display` (first alias) | `v0.3.0:internal/planner/render.go:128,197`, `internal/generator/generate.go:276` | ✔ |
| `provider.DiscoverRequest` is `{Types []string}` (no Region) | `v0.3.0:pkg/provider/provider.go:68` | ✔ |
| Reserved attributes are exactly `prevent_destroy`, `retain` (exact match) | `v0.3.0:internal/registry/registry.go:291`, `internal/pluginhost/adapter.go:367` | ✔ |
| Resource keys `type`, `depends_on`, `provider`, `skip`, `only`, `lifecycle` match exactly | `v0.3.0:internal/config/decode.go:455-521` | ✔ |
| `value.Equal` needs identical map keys and counts, lists by position | `v0.3.0:pkg/value/value.go:235-280` | ✔ |
| infrata release asset for linux/amd64 is `infrata_0.3.0_linux_amd64.tar.gz` | `gh release view v0.3.0 --repo infrata/infrata` | ✔ |
| Schema bundle URL and size; 1,729 `AWS::` schemas; 1,584 with create/read/delete handlers = Cloud Control's provisionable set | download + `cloudformation list-types` | ✔ |
| No property-name fold or snake_case collisions within a type | bundle analysis (Python regex equivalent to Task 1's `SnakeCase`) | ✔ (Task 5 re-proves it through `Validate` in Go) |
| Keyword clashes: `Type` 95, `Provider` 8, `Region` 5, `Lifecycle` 1 | bundle analysis | ✔ |
| Tag property shapes: array 919, object map 55, `$ref` 112, other 2 | bundle analysis | ✔ |
| 2,069 properties set `insertionOrder: false` (default is true) | bundle analysis; schema docs | ✔ |
| Free-form maps 447, opaque objects 323, `oneOf` 306, `anyOf` 66, `allOf` 12, type unions 74, `$ref` to `#/properties/` 2 | bundle analysis | ✔ |
| Cloud Control error codes and HTTP statuses (e.g. `ThrottlingException` 429, `ConcurrentOperationException` 409, `ResourceNotFoundException` 404, `HandlerInternalFailureException` 502, `ConcurrentModificationException` 500) | `cloudcontrol.json` model `smithy.api#error` / `httpError` traits | ✔ |
| awsJson1_0 errors: code from `X-Amzn-ErrorType` or body `__type`; message from body `message`/`Message`; timestamps epoch seconds | `smithy-go@v1.28.1 transport/http/protocol/awsjson/awsjson.go:190-240`, `internal/json/shape_deserializer.go:302` | ✔ |
| `X-Amz-Target: <service>.<Operation>`, `Content-Type: application/x-amz-json-1.0` | `awsjson.go:87,93` | ✔ |
| Cloud Control service name in the target header is `CloudApiService` | `cloudcontrol.json` service shape `com.amazonaws.cloudcontrol#CloudApiService` | ✔ |
| `AWS_ENDPOINT_URL_CLOUDCONTROL` overrides the endpoint | `service/cloudcontrol@v1.38.0/endpoints.go` | ✔ |
| Live Cloud Control create/update/delete of a VPC works; `GetResource` returns provider-chosen properties | spike `cmd/cloudcontrol-poc -live` | ✔ |
| Correction (Task 8): a known/registered error shape (e.g. `InvalidRequestException`) is decoded by the schema deserializer, which matches the member name `Message` exactly (no `JSONName` trait on it); the case-insensitive `message`/`Message` fallback only covers the generic `ProtocolErrorInfo` path for an *unregistered* code. `ccfake.writeError` sent lowercase `message`, so a registered error's `Message` field decoded empty; fixed to send `Message` | `service/cloudcontrol@v1.38.0/schemas/schemas.go` (`AddMember("Message", _ErrorMessage)`), `smithy-go@v1.28.1/transport/http/protocol/internal/json/shape_deserializer.go` (`memberFromToken`, exact byte match) | ✔ (found by `TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause` in Task 8) |
| P7 load gate (Task 6): median extra time `infrata validate` took with the full catalog over a no-op plugin was 487–489 ms | `scripts/measure-load`, run under the project's pre-rename name | ✔ — under the 500 ms gate as originally written |
| The project (module `github.com/infrata/infrata-provider-aws` → `github.com/infrena/infrena-provider-aws`, CLI `infrata` → `infrena`, binary `infrata-plugin-aws` → `infrena-plugin-aws`, env vars `INFRATA_*` → `INFRENA_*`, `plugin.yaml`'s `infrata:` floor → `infrena: ">= 0.4.0"`, CI secret `INFRATA_CHECKOUT_TOKEN` → `INFRENA_CHECKOUT_TOKEN`) was renamed from infrata to infrena, pinned to infrena `v0.4.0` | `go.mod`, `plugin.yaml`, repo-wide rename commits `0c68c9f`, `2825f7f` and others on `generic-cloudcontrol` | ✔ |
| Rerun of the P7 load gate after the rename: median extra time was about 437 ms, still comfortably under the gate. James decided ~500 ms extra is fine either way and only wants a flag if the extra goes over about 1 second | `scripts/measure-load`, run post-rename | ✔ — no flag needed; CLAUDE.md's load-cost rule states the 1 s concern threshold instead of a hard 500 ms cutoff |
| infrena host finding: `Update` receives the last **persisted** state as current, not a state `apply` freshly refreshed, so diffing against it directly misses or fabricates changes | found running the e2e suite (a VPC's tags, drifted outside infrena then corrected in configuration, never got patched); fixed here by having `Update` read fresh before building its patch | ✔ — commit `c7ff98e` in this repo. Reported to the infrena session; they've confirmed it and plan to fix it on the host side after the rename. The extra read here stays until the Verification log records that fix has landed |
| Task 17 whole-suite verification (2026-09-14, post-rename): `gofmt -l .` clean; `go vet ./...` and `go vet -tags e2e,live ./...` clean; `go test -count=1 ./...` and the pinned `GOWORK=off GOPRIVATE='github.com/infrena/*' go test -count=1 ./...` both pass every package; `go test -tags e2e -count=1 -v ./e2e/` against a read-only snapshot of infrena `v0.4.0` (commit `e2be8bf`, extracted with `git archive`) passes every test (`TestTheWorkflow` and its 9 subtests, plus 4 more top-level tests, 0 failures); `go test -tags live -count=1 ./live/` skips both tests (no `INFRENA_AWS_LIVE_*` set, no AWS credentials touched); `scripts/release-check v0.1.0` passes (tag, manifest and binary agree on version `0.1.0` and protocol `[2]`). `scripts/measure-load` was not rerun in this step (already measured post-rename, row above) | this session's command output | ✔ |
| The live suite against real AWS has never been run. It needs James's explicit approval each time, and a dedicated AWS profile — the `infrata` profile holds root keys that `live/live_test.go`'s `guard` refuses before any Cloud Control call | `live/README.md`, `live/live_test.go` | ✔ — still true; not run in this session either |

## File structure

```text
gen/overlay.yaml                          curated: global types, aliases, sensitive, requirements, discover_default
gen/names.lock.json                       committed: CFN type -> infrata type, never shrinks
scripts/fetch-schemas                     downloads the bundle into schemas/ (gitignored) and prints its SHA-256
internal/cfn/        schema.go bundle.go snake.go          parse CloudFormation resource schemas
internal/catalog/    catalog.go embed.go catalog.json.gz   catalog types, infrata definitions, embedded data
internal/gen/        lock.go overlay.go build.go shape.go generate.go   generator logic
cmd/gen-cloudcontrol/main.go              generator command
internal/ccprov/     ids.go errors.go patience.go clients.go await.go        Task 8: plumbing
                     provider.go values.go crud.go                         Task 9: create, read, delete, import
                     reconcile.go tags.go                                  Task 10: nested values, tags as a map
                     patch.go                                              Task 11: update
                     discover.go                                           Task 12: discover
internal/ccfake/     server.go                             Task 7: in-process Cloud Control (and STS) fake
internal/awstest/    awstest.go faketype.go core.go        test helpers: SDK isolation, fake types from the catalog
internal/awsprov/    plugin.go config.go credentials.go    plugin, instance configuration, credentials
scripts/measure-load                      Task 6: the P7 load-cost gate
e2e/ (Task 14), plugin.yaml scripts/release-check scripts/build-release .github/workflows/{ci,release,bump-infrata,bump-schemas}.yml (Task 15), live/ (Task 16)
REMOVED: internal/awsprov/{vpc,subnet,tags,ids,attrs,values,definitions,errors,clients,patience,provider}.go and their tests,
         internal/ec2fake/
```

Tasks: 1 `internal/cfn` · 2 `internal/catalog` · 3 name lock · 4 overlay and type building · 5 the real catalog ·
6 the plugin serves it, and the load gate · 7 `ccfake` · 8 ccprov plumbing · 9 create/read/delete/import ·
10 reconciliation and tags · 11 update · 12 discover · 13 through infrata's host · 14 e2e · 15 release and CI ·
16 live suite · 17 docs and close-out.

---
### Task 1: Parse CloudFormation resource schemas

**Files:**
- Create: `scripts/fetch-schemas`, `internal/cfn/schema.go`, `internal/cfn/bundle.go`, `internal/cfn/snake.go`
- Create: `internal/cfn/testdata/*.json` (11 real schemas copied from the bundle)
- Modify: `.gitignore` (add `/schemas/`)
- Test: `internal/cfn/schema_test.go`, `internal/cfn/snake_test.go`

**Interfaces:**
- Produces:
  - `type Schema struct { TypeName, Description string; Properties, Definitions map[string]*Node; Required, PrimaryIdentifier, ReadOnlyProperties, CreateOnlyProperties, WriteOnlyProperties []string; Handlers map[string]Handler; Tagging *Tagging }`
  - `type Node struct { Type TypeList; Ref string; Description string; Properties, PatternProperties map[string]*Node; AdditionalProperties json.RawMessage; Items *Node; InsertionOrder *bool; OneOf, AnyOf, AllOf []*Node }`
  - `type TypeList []string`; `type Handler struct { Permissions []string; TimeoutInMinutes int; HandlerSchema json.RawMessage }`; `type Tagging struct { Taggable *bool; TagProperty string }`
  - `func Parse(raw []byte) (*Schema, error)`
  - `func (s *Schema) Provisionable() bool`, `func (s *Schema) HasHandler(name string) bool`, `func (s *Schema) ListNeedsModel() bool`
  - `func (s *Schema) Resolve(n *Node) *Node` (follows `$ref`, loop-safe)
  - `func (s *Schema) TopLevel(pointers []string) map[string]bool`, `func (s *Schema) Nested(pointers []string) []string`
  - `func (s *Schema) Timeout(handler string) int` (minutes; 120 when unset)
  - `func ReadBundle(zipPath string) ([]*Schema, string, error)` (schemas and the zip's SHA-256 hex)
  - `func SnakeCase(name string) string`

- [ ] **Step 1: Fetch the bundle and copy the fixtures**

`scripts/fetch-schemas`:

```bash
#!/usr/bin/env bash
# fetch-schemas: download AWS's CloudFormation resource schema bundle for us-east-1, which the generator reads.
# The bundle is not committed (about 3 MB, changes weekly); its SHA-256 is recorded in the generated catalog.
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p schemas
curl -fsSL -o schemas/CloudformationSchema.zip https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip
sha256sum schemas/CloudformationSchema.zip
```

```bash
chmod +x scripts/fetch-schemas
printf '/schemas/\n' >> .gitignore
scripts/fetch-schemas
mkdir -p internal/cfn/testdata
unzip -o -j schemas/CloudformationSchema.zip \
  aws-ec2-vpc.json aws-ec2-subnet.json aws-ec2-securitygroup.json aws-s3-bucket.json aws-iam-role.json \
  aws-rds-dbinstance.json aws-acmpca-certificate.json aws-acmpca-certificateauthority.json \
  aws-aps-anomalydetector.json aws-arczonalshift-autoshiftobservernotificationstatus.json \
  aws-codepipeline-customactiontype.json -d internal/cfn/testdata
ls internal/cfn/testdata | wc -l    # 11
```

Why these: VPC and Subnet (the spike's types), SecurityGroup (unordered nested lists), S3 Bucket (a bucket name AWS
generates), IAM Role (global), RDS DBInstance (101 properties, write-only password), ACMPCA Certificate (composite
identifier), ACMPCA CertificateAuthority (a `Type` property), APS AnomalyDetector (list needs a parent model),
ARCZonalShift AutoshiftObserverNotificationStatus (a `Region` property), CodePipeline CustomActionType (a `Provider`
property).

- [ ] **Step 2: Write the failing tests**

`internal/cfn/snake_test.go`:

```go
package cfn

import "testing"

// TestSnakeCase pins the rule the bundle analysis used: an underscore before an upper-case letter that follows a
// lower-case letter or digit, or that follows an upper-case letter and precedes a lower-case one.
func TestSnakeCase(t *testing.T) {
	for in, want := range map[string]string{
		"CidrBlock":            "cidr_block",
		"VpcId":                "vpc_id",
		"EnableDnsHostnames":   "enable_dns_hostnames",
		"DBInstanceIdentifier": "db_instance_identifier",
		"Ipv6CidrBlocks":       "ipv6_cidr_blocks",
		"SSESpecification":     "sse_specification",
		"S3Bucket":             "s3_bucket",
		"ARN":                  "arn",
		"Type":                 "type",
		"already_snake":        "already_snake",
	} {
		if got := SnakeCase(in); got != want {
			t.Errorf("SnakeCase(%q) = %q, want %q", in, got, want)
		}
	}
}
```

`internal/cfn/schema_test.go`:

```go
package cfn

import (
	"os"
	"strings"
	"testing"
)

func load(t *testing.T, file string) *Schema {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}
	s, err := Parse(raw)
	if err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	return s
}

func TestParseReadsTheFieldsTheGeneratorUses(t *testing.T) {
	s := load(t, "aws-ec2-vpc.json")
	if s.TypeName != "AWS::EC2::VPC" || !s.Provisionable() || !s.HasHandler("update") {
		t.Fatalf("type=%q provisionable=%v update=%v", s.TypeName, s.Provisionable(), s.HasHandler("update"))
	}
	if strings.Join(s.PrimaryIdentifier, ",") != "/properties/VpcId" {
		t.Errorf("identifier = %v", s.PrimaryIdentifier)
	}
	if !s.TopLevel(s.CreateOnlyProperties)["CidrBlock"] || !s.TopLevel(s.ReadOnlyProperties)["VpcId"] {
		t.Error("top-level flags not read")
	}
	if len(s.Nested(s.ReadOnlyProperties)) == 0 {
		t.Error("the VPC schema has nested read-only pointers; Nested returned none")
	}
	if s.Tagging == nil || s.Tagging.TagProperty != "/properties/Tags" {
		t.Errorf("tagging = %+v", s.Tagging)
	}
	if s.Timeout("create") <= 0 {
		t.Errorf("create timeout = %d", s.Timeout("create"))
	}
}

func TestResolveFollowsReferencesToAnArrayOfObjects(t *testing.T) {
	s := load(t, "aws-ec2-securitygroup.json")
	ingress := s.Properties["SecurityGroupIngress"]
	if ingress.InsertionOrder == nil || *ingress.InsertionOrder {
		t.Fatalf("SecurityGroupIngress insertionOrder = %v, want false", ingress.InsertionOrder)
	}
	item := s.Resolve(ingress.Items)
	if item == nil || item.Properties["CidrIp"] == nil || item.Properties["IpProtocol"] == nil {
		t.Fatalf("Ingress item did not resolve to its definition: %+v", item)
	}
}

func TestListNeedsModelAndCompositeIdentifiers(t *testing.T) {
	if !load(t, "aws-aps-anomalydetector.json").ListNeedsModel() {
		t.Error("AnomalyDetector lists only under a workspace; ListNeedsModel = false")
	}
	if load(t, "aws-ec2-vpc.json").ListNeedsModel() {
		t.Error("VPC lists without a model; ListNeedsModel = true")
	}
	if n := len(load(t, "aws-acmpca-certificate.json").PrimaryIdentifier); n < 2 {
		t.Errorf("ACMPCA Certificate identifier parts = %d, want composite", n)
	}
}

func TestTypeAcceptsAStringOrAList(t *testing.T) {
	s, err := Parse([]byte(`{"typeName":"AWS::X::Y","properties":{"A":{"type":"string"},"B":{"type":["object","string"]}}}`))
	if err != nil {
		t.Fatal(err)
	}
	if got := s.Properties["A"].Type; len(got) != 1 || got[0] != "string" {
		t.Errorf("A type = %v", got)
	}
	if got := s.Properties["B"].Type; len(got) != 2 {
		t.Errorf("B type = %v", got)
	}
}

func TestResolveIsLoopSafe(t *testing.T) {
	s, err := Parse([]byte(`{"typeName":"AWS::X::Y","properties":{"A":{"$ref":"#/definitions/L"}},
		"definitions":{"L":{"$ref":"#/definitions/M"},"M":{"$ref":"#/definitions/L"}}}`))
	if err != nil {
		t.Fatal(err)
	}
	if got := s.Resolve(s.Properties["A"]); got != nil {
		t.Errorf("a reference loop resolved to %+v, want nil", got)
	}
}

func TestParseRefusesSomethingThatIsNotASchema(t *testing.T) {
	if _, err := Parse([]byte(`{"hello":"world"}`)); err == nil {
		t.Fatal("accepted a document with no typeName")
	}
}
```

- [ ] **Step 3: Run them to see them fail**

Run: `go test -count=1 ./internal/cfn/`
Expected: FAIL — `undefined: SnakeCase`, `undefined: Parse`.

- [ ] **Step 4: Implement**

`internal/cfn/snake.go`:

```go
package cfn

import (
	"strings"
	"unicode"
)

// SnakeCase converts a CloudFormation property name to the snake_case alias the generator adds (J4).
func SnakeCase(name string) string {
	rs := []rune(name)
	var b strings.Builder
	for i, r := range rs {
		if i > 0 && unicode.IsUpper(r) {
			prev := rs[i-1]
			nextLower := i+1 < len(rs) && unicode.IsLower(rs[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextLower) {
				b.WriteByte('_')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
```

`internal/cfn/schema.go`:

```go
// Package cfn reads AWS CloudFormation resource type schemas: the documents AWS Cloud Control API is driven by.
// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
package cfn

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Schema is the part of a resource type schema the generator and provider read.
type Schema struct {
	TypeName             string           `json:"typeName"`
	Description          string           `json:"description"`
	Properties           map[string]*Node `json:"properties"`
	Definitions          map[string]*Node `json:"definitions"`
	Required             []string         `json:"required"`
	PrimaryIdentifier    []string         `json:"primaryIdentifier"`
	ReadOnlyProperties   []string         `json:"readOnlyProperties"`
	CreateOnlyProperties []string         `json:"createOnlyProperties"`
	WriteOnlyProperties  []string         `json:"writeOnlyProperties"`
	Handlers             map[string]Handler `json:"handlers"`
	Tagging              *Tagging         `json:"tagging"`
}

// Node is one JSON-schema node inside a resource schema.
type Node struct {
	Type                 TypeList         `json:"type"`
	Ref                  string           `json:"$ref"`
	Description          string           `json:"description"`
	Properties           map[string]*Node `json:"properties"`
	PatternProperties    map[string]*Node `json:"patternProperties"`
	AdditionalProperties json.RawMessage  `json:"additionalProperties"`
	Items                *Node            `json:"items"`
	InsertionOrder       *bool            `json:"insertionOrder"`
	OneOf                []*Node          `json:"oneOf"`
	AnyOf                []*Node          `json:"anyOf"`
	AllOf                []*Node          `json:"allOf"`
}

// TypeList is a JSON-schema "type", which may be a string or a list of strings.
type TypeList []string

func (t *TypeList) UnmarshalJSON(b []byte) error {
	var one string
	if err := json.Unmarshal(b, &one); err == nil {
		*t = TypeList{one}
		return nil
	}
	var many []string
	if err := json.Unmarshal(b, &many); err != nil {
		return fmt.Errorf("type is neither a string nor a list of strings: %s", b)
	}
	*t = many
	return nil
}

// Handler is one provisioning handler.
type Handler struct {
	Permissions      []string        `json:"permissions"`
	TimeoutInMinutes int             `json:"timeoutInMinutes"`
	HandlerSchema    json.RawMessage `json:"handlerSchema"`
}

// Tagging is how a type supports tags.
type Tagging struct {
	Taggable    *bool  `json:"taggable"`
	TagProperty string `json:"tagProperty"`
}

// Parse reads one schema document.
func Parse(raw []byte) (*Schema, error) {
	var s Schema
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, err
	}
	if s.TypeName == "" {
		return nil, fmt.Errorf("not a resource type schema: no typeName")
	}
	return &s, nil
}

// HasHandler reports whether the schema declares a handler.
func (s *Schema) HasHandler(name string) bool {
	_, ok := s.Handlers[name]
	return ok
}

// Provisionable reports whether Cloud Control can manage the type: it needs create, read and delete handlers.
// Measured against list-types on 2026-09-14, this selects exactly Cloud Control's provisionable set.
func (s *Schema) Provisionable() bool {
	return s.HasHandler("create") && s.HasHandler("read") && s.HasHandler("delete")
}

// ListNeedsModel reports whether listing requires a parent resource model (the list handler's schema has required
// properties).
func (s *Schema) ListNeedsModel() bool {
	h, ok := s.Handlers["list"]
	if !ok || len(h.HandlerSchema) == 0 {
		return false
	}
	var hs struct {
		Required []string `json:"required"`
	}
	if err := json.Unmarshal(h.HandlerSchema, &hs); err != nil {
		return false
	}
	return len(hs.Required) > 0
}

// Timeout is a handler's timeout in minutes; the schema default is 120.
func (s *Schema) Timeout(handler string) int {
	if h, ok := s.Handlers[handler]; ok && h.TimeoutInMinutes > 0 {
		return h.TimeoutInMinutes
	}
	return 120
}

// TopLevel returns the property names of pointers of the form /properties/Name.
func (s *Schema) TopLevel(pointers []string) map[string]bool {
	out := map[string]bool{}
	for _, p := range pointers {
		if name, ok := strings.CutPrefix(p, "/properties/"); ok && !strings.Contains(name, "/") {
			out[name] = true
		}
	}
	return out
}

// Nested returns the pointers deeper than /properties/Name, which infrata's per-attribute flags cannot express.
func (s *Schema) Nested(pointers []string) []string {
	var out []string
	for _, p := range pointers {
		if name, ok := strings.CutPrefix(p, "/properties/"); ok && strings.Contains(name, "/") {
			out = append(out, p)
		}
	}
	return out
}

// Resolve follows $ref until it reaches a node without one. References to #/definitions/... and #/properties/... are
// supported. A loop, or a reference to something missing, resolves to nil.
func (s *Schema) Resolve(n *Node) *Node {
	seen := map[string]bool{}
	for n != nil && n.Ref != "" {
		if seen[n.Ref] {
			return nil
		}
		seen[n.Ref] = true
		n = s.pointer(n.Ref)
	}
	return n
}

func (s *Schema) pointer(ref string) *Node {
	path, ok := strings.CutPrefix(ref, "#/")
	if !ok {
		return nil
	}
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return nil
	}
	var n *Node
	switch parts[0] {
	case "definitions":
		n = s.Definitions[parts[1]]
	case "properties":
		n = s.Properties[parts[1]]
	default:
		return nil
	}
	for i := 2; n != nil && i+1 < len(parts); i += 2 {
		if parts[i] != "properties" {
			return nil
		}
		n = n.Properties[parts[i+1]]
	}
	return n
}
```

`internal/cfn/bundle.go`:

```go
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
```

- [ ] **Step 5: Run the tests**

Run: `go test -count=1 ./internal/cfn/ && go vet ./internal/cfn/ && gofmt -l internal/cfn`
Expected: PASS, no gofmt output.

- [ ] **Step 6: Sabotage, then commit**

Sabotages (each must compile and fail): drop the `unicode.IsDigit(prev)` condition (`S3Bucket` becomes `s3bucket`);
`Resolve` returns its argument without following `$ref` (security-group test); `ListNeedsModel` returns
`len(h.HandlerSchema) > 0` (a schema-only list handler would wrongly need a model; the VPC assertion fails if its list
handler carries a schema, so also check APS still passes); remove the `seen` guard (the loop test hangs, so run it with
`-timeout 10s`).

```bash
git add scripts/fetch-schemas .gitignore internal/cfn/schema.go internal/cfn/bundle.go internal/cfn/snake.go \
  internal/cfn/schema_test.go internal/cfn/snake_test.go internal/cfn/testdata
git commit -m "Read CloudFormation resource schemas

Adds a parser for the schema bundle AWS publishes, the snake_case rule used
for aliases, and eleven real schemas as fixtures. Checked by breaking the
digit rule, ref resolution, the list model check and the loop guard." -- \
  scripts/fetch-schemas .gitignore internal/cfn/schema.go internal/cfn/bundle.go internal/cfn/snake.go \
  internal/cfn/schema_test.go internal/cfn/snake_test.go internal/cfn/testdata
```

---
### Task 2: The catalog — what the generator writes and the plugin reads

**Files:**
- Create: `internal/catalog/catalog.go`
- Test: `internal/catalog/catalog_test.go`

**Interfaces:**
- Produces (every later task uses these exact names):
  - `type Catalog struct { Bundle string; Types []*Type }` (JSON `bundle`, `types`)
  - `type Type struct { Name, CFN, Description, RegionAttr, TagsAsMap string; Identifier, WriteOnly []string; HasUpdate, HasList, ListNeedsModel bool; Timeouts map[string]int; Attributes []*Attribute; Requirements []Requirement }`
  - `type Attribute struct { Name, Kind, Description string; Required, Computed, Optional, ForceNew, Sensitive bool; Aliases []string; Shape *Shape }`
  - `type Shape struct { Kind string; Props map[string]*Shape; Item *Shape; Unordered bool }` with kinds `ShapeObject`, `ShapeArray`, `ShapeScalar`, `ShapeOpaque`
  - `type Requirement struct { Name string; Types []string; Description string }`
  - `func Read(r io.Reader) (*Catalog, error)` (gzip JSON), `func (c *Catalog) Write(w io.Writer) error`
  - `func (c *Catalog) Lookup(name string) (*Type, bool)`, `func (c *Catalog) Definitions() []*schema.ResourceDefinition`
  - `func (t *Type) Definition() *schema.ResourceDefinition`, `func (t *Type) Attribute(name string) (*Attribute, bool)`, `func (t *Type) Global() bool`
  - `const GlobalScope = "global"`

- [ ] **Step 1: Write the failing test**

`internal/catalog/catalog_test.go`:

```go
package catalog

import (
	"bytes"
	"strings"
	"testing"
)

func sample() *Catalog {
	return &Catalog{Bundle: "abc123", Types: []*Type{
		{
			Name: "aws.vpc", CFN: "AWS::EC2::VPC", Description: "A VPC.", RegionAttr: "region",
			Identifier: []string{"VpcId"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
			Timeouts: map[string]int{"create": 120},
			Attributes: []*Attribute{
				{Name: "CidrBlock", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"cidr", "cidr_block"}},
				{Name: "VpcId", Kind: "string", Computed: true, Aliases: []string{"vpc_id"}},
				{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: &Shape{Kind: ShapeOpaque}},
			},
		},
		{
			Name: "aws.role", CFN: "AWS::IAM::Role", Identifier: []string{"RoleName"}, HasUpdate: true,
			Attributes: []*Attribute{
				{Name: "AssumeRolePolicyDocument", Kind: "map", Required: true, Shape: &Shape{Kind: ShapeOpaque}},
				{Name: "RoleName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"role_name"}},
			},
			Requirements: []Requirement{{Name: "vpc", Types: []string{"aws.vpc"}, Description: "only a test"}},
		},
	}}
}

func TestDefinitionsValidateAndCarryTheFlags(t *testing.T) {
	c := sample()
	defs := c.Definitions()
	if len(defs) != 2 {
		t.Fatalf("got %d definitions", len(defs))
	}
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Fatalf("%s: %v", d.Type, err)
		}
	}
	vpc, _ := c.Lookup("aws.vpc")
	d := vpc.Definition()
	cidr := d.Attributes["CidrBlock"]
	if !cidr.Optional || !cidr.Computed || !cidr.ForceNew || strings.Join(cidr.Aliases, ",") != "cidr,cidr_block" {
		t.Errorf("CidrBlock = %+v", cidr)
	}
	if d.Display("CidrBlock") != "cidr" {
		t.Errorf("Display = %q, want the first alias", d.Display("CidrBlock"))
	}
	region, ok := d.Attributes["region"]
	if !ok || !region.Required || !region.ForceNew {
		t.Errorf("regional type's region attribute = %+v, %v", region, ok)
	}
	if !d.Capabilities.Update || !d.Capabilities.Import {
		t.Errorf("capabilities = %+v", d.Capabilities)
	}
}

func TestAGlobalTypeHasNoRegionAndSaysHowToImport(t *testing.T) {
	role, ok := sample().Lookup("aws.role")
	if !ok || !role.Global() {
		t.Fatalf("aws.role global = %v", ok && role.Global())
	}
	d := role.Definition()
	if _, has := d.Attributes["region"]; has {
		t.Error("a global type declares region")
	}
	if !strings.Contains(d.ImportID.Description, GlobalScope+"/") {
		t.Errorf("import description = %q", d.ImportID.Description)
	}
	if len(d.Requirements) != 1 || d.Requirements[0].Types[0] != "aws.vpc" {
		t.Errorf("requirements = %+v", d.Requirements)
	}
}

func TestAnImmutableTypeCannotUpdate(t *testing.T) {
	c := sample()
	c.Types[0].HasUpdate = false
	vpc, _ := c.Lookup("aws.vpc")
	if vpc.Definition().Capabilities.Update {
		t.Error("an immutable type offers update")
	}
}

func TestCatalogRoundTripsThroughGzipJSON(t *testing.T) {
	var buf bytes.Buffer
	if err := sample().Write(&buf); err != nil {
		t.Fatal(err)
	}
	got, err := Read(&buf)
	if err != nil {
		t.Fatal(err)
	}
	vpc, ok := got.Lookup("aws.vpc")
	if !ok || got.Bundle != "abc123" || vpc.TagsAsMap != "Tags" {
		t.Fatalf("round trip lost data: %+v", got)
	}
	a, ok := vpc.Attribute("Tags")
	if !ok || a.Shape == nil || a.Shape.Kind != ShapeOpaque {
		t.Errorf("Tags shape after round trip = %+v", a)
	}
}

func TestReadRefusesDuplicateNames(t *testing.T) {
	c := sample()
	c.Types[1].Name = "aws.vpc"
	var buf bytes.Buffer
	if err := c.Write(&buf); err != nil {
		t.Fatal(err)
	}
	if _, err := Read(&buf); err == nil || !strings.Contains(err.Error(), "aws.vpc") {
		t.Fatalf("err = %v", err)
	}
}
```

- [ ] **Step 2: Run it to see it fail**

Run: `go test -count=1 ./internal/catalog/`
Expected: FAIL — `undefined: Catalog`.

- [ ] **Step 3: Implement**

`internal/catalog/catalog.go`:

```go
// Package catalog holds the generated description of every Cloud Control resource type the plugin serves: the
// infrata definitions it hands the host, and the runtime metadata the provider needs. It is produced by
// cmd/gen-cloudcontrol and embedded; nothing here is written by hand.
package catalog

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

// GlobalScope is the provider-ID scope of a type that has no region.
const GlobalScope = "global"

// Shape kinds.
const (
	ShapeObject = "object" // named properties: keys are translated and reconciled
	ShapeArray  = "array"  // items share one shape
	ShapeScalar = "scalar"
	ShapeOpaque = "opaque" // copied exactly: free-form maps, unions, anything without a fixed set of properties
)

// Catalog is every generated type.
type Catalog struct {
	Bundle          string   `json:"bundle"` // SHA-256 of the schema bundle it was generated from
	DiscoverDefault []string `json:"discover_default,omitempty"` // infrata type names discovered when an instance sets no discover_types (P4)
	Types           []*Type  `json:"types"`

	index map[string]*Type
}

// Type is one Cloud Control resource type.
type Type struct {
	Name           string         `json:"name"` // infrata type, e.g. aws.vpc
	CFN            string         `json:"cfn"`  // CloudFormation type, e.g. AWS::EC2::VPC
	Description    string         `json:"description,omitempty"`
	RegionAttr     string         `json:"region_attr,omitempty"` // "region", "aws_region", or "" for a global type
	Identifier     []string       `json:"identifier"`            // primary identifier property names, in order
	WriteOnly      []string       `json:"write_only,omitempty"`  // top-level properties AWS never returns
	HasUpdate      bool           `json:"has_update,omitempty"`
	HasList        bool           `json:"has_list,omitempty"`
	ListNeedsModel bool           `json:"list_needs_model,omitempty"`
	TagsAsMap      string         `json:"tags_as_map,omitempty"` // property exposed as a map instead of [{Key, Value}]
	Timeouts       map[string]int `json:"timeouts,omitempty"`    // handler timeouts in minutes
	Attributes     []*Attribute   `json:"attributes"`
	Requirements   []Requirement  `json:"requirements,omitempty"`
}

// Attribute is one top-level property.
type Attribute struct {
	Name        string   `json:"name"` // AWS's property name: the canonical, stored name
	Kind        string   `json:"kind"` // an infrata kind name: string, integer, float, boolean, list, map
	Required    bool     `json:"required,omitempty"`
	Computed    bool     `json:"computed,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	ForceNew    bool     `json:"force_new,omitempty"`
	Sensitive   bool     `json:"sensitive,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
	Description string   `json:"description,omitempty"`
	Shape       *Shape   `json:"shape,omitempty"` // for list and map kinds
}

// Shape describes a nested value for reconciliation.
type Shape struct {
	Kind      string            `json:"k"`
	Props     map[string]*Shape `json:"p,omitempty"`
	Item      *Shape            `json:"i,omitempty"`
	Unordered bool              `json:"u,omitempty"`
}

// Requirement is a pre-flight hint between types, from the overlay.
type Requirement struct {
	Name        string   `json:"name"`
	Types       []string `json:"types"` // infrata type names
	Description string   `json:"description"`
}

// Read decodes a gzip-compressed JSON catalog.
func Read(r io.Reader) (*Catalog, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var c Catalog
	if err := json.NewDecoder(zr).Decode(&c); err != nil {
		return nil, err
	}
	c.index = make(map[string]*Type, len(c.Types))
	for _, t := range c.Types {
		if _, dup := c.index[t.Name]; dup {
			return nil, fmt.Errorf("catalog declares %s twice", t.Name)
		}
		c.index[t.Name] = t
	}
	return &c, nil
}

// Write encodes the catalog as gzip-compressed JSON, types sorted by name so regeneration diffs cleanly.
func (c *Catalog) Write(w io.Writer) error {
	sort.Slice(c.Types, func(i, j int) bool { return c.Types[i].Name < c.Types[j].Name })
	zw := gzip.NewWriter(w)
	enc := json.NewEncoder(zw)
	if err := enc.Encode(c); err != nil {
		return err
	}
	return zw.Close()
}

// Lookup finds a type by its infrata name.
func (c *Catalog) Lookup(name string) (*Type, bool) {
	if c.index == nil {
		c.index = make(map[string]*Type, len(c.Types))
		for _, t := range c.Types {
			c.index[t.Name] = t
		}
	}
	t, ok := c.index[name]
	return t, ok
}

// Definitions converts every type to an infrata definition.
func (c *Catalog) Definitions() []*schema.ResourceDefinition {
	out := make([]*schema.ResourceDefinition, 0, len(c.Types))
	for _, t := range c.Types {
		out = append(out, t.Definition())
	}
	return out
}

// Global reports whether the type has no region.
func (t *Type) Global() bool { return t.RegionAttr == "" }

// Attribute finds an attribute by its canonical name.
func (t *Type) Attribute(name string) (*Attribute, bool) {
	for _, a := range t.Attributes {
		if a.Name == name {
			return a, true
		}
	}
	return nil, false
}

// Definition is the infrata schema for the type. The region attribute is the plugin's own (spec §3.1).
func (t *Type) Definition() *schema.ResourceDefinition {
	attrs := make(map[string]schema.Attribute, len(t.Attributes)+1)
	for _, a := range t.Attributes {
		kind, _ := value.ParseKind(a.Kind)
		attrs[a.Name] = schema.Attribute{
			Kind: kind, Required: a.Required, Computed: a.Computed, Optional: a.Optional,
			ForceNew: a.ForceNew, Sensitive: a.Sensitive, Aliases: a.Aliases, Description: a.Description,
		}
	}
	scope := "<region>"
	if t.Global() {
		scope = GlobalScope
	} else {
		attrs[t.RegionAttr] = schema.Attribute{
			Kind: value.KindString, Required: true, ForceNew: true,
			Description: "AWS region, usually set once with the provider's defaults: {region: ...}",
		}
	}
	var reqs []schema.Requirement
	for _, r := range t.Requirements {
		reqs = append(reqs, schema.Requirement{Name: r.Name, Types: r.Types, Description: r.Description})
	}
	return &schema.ResourceDefinition{
		Type:         t.Name,
		Description:  t.Description,
		Attributes:   attrs,
		Requirements: reqs,
		Capabilities: schema.Capabilities{Create: true, Read: true, Update: t.HasUpdate, Delete: true, Import: true},
		ImportID: schema.ImportSpec{Description: fmt.Sprintf("%s/<identifier>, the identifier being %s (%s)",
			scope, strings.Join(t.Identifier, "|"), t.CFN)},
	}
}
```

- [ ] **Step 4: Run the tests**

Run: `go test -count=1 ./internal/catalog/ && go vet ./internal/catalog/`
Expected: PASS.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: omit `Optional` when building `schema.Attribute` (`TestDefinitionsValidate…` fails on the CidrBlock flags);
add the region attribute for global types too (global test); ignore `HasUpdate` (immutable test); drop the duplicate
check in `Read` (duplicate test).

```bash
git add internal/catalog/catalog.go internal/catalog/catalog_test.go
git commit -m "Add the catalog types the generator writes

A catalog type turns into an infrata definition with the right flags,
aliases and region attribute. Checked by dropping Optional, adding region to
global types, ignoring the update handler and allowing duplicate names." -- \
  internal/catalog/catalog.go internal/catalog/catalog_test.go
```

---
### Task 3: Type names, and the lock that keeps them

**Files:**
- Create: `internal/gen/lock.go`
- Test: `internal/gen/lock_test.go`

**Interfaces:**
- Produces:
  - `type Lock struct { Names map[string]string }` (JSON `names`: CloudFormation type → infrata type)
  - `func LoadLock(path string) (*Lock, error)` (a missing file is an empty lock)
  - `func (l *Lock) Save(path string) error` (sorted, trailing newline)
  - `func (l *Lock) Assign(cfnTypes []string) (map[string]string, error)` (names for the given types; new ones are added to `l.Names`; entries for types AWS removed stay as tombstones)

- [ ] **Step 1: Write the failing test**

`internal/gen/lock_test.go`:

```go
package gen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestAFreshLockUsesShortNamesOnlyWhenUnique. J2: the resource segment lowercased when no other supported type
// shares it, otherwise service.segment.
func TestAFreshLockUsesShortNamesOnlyWhenUnique(t *testing.T) {
	l := &Lock{}
	got, err := l.Assign([]string{"AWS::EC2::VPC", "AWS::EC2::Instance", "AWS::Lightsail::Instance", "AWS::S3::Bucket"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"AWS::EC2::VPC":            "aws.vpc",
		"AWS::EC2::Instance":       "aws.ec2.instance",
		"AWS::Lightsail::Instance": "aws.lightsail.instance",
		"AWS::S3::Bucket":          "aws.bucket",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s -> %q, want %q", k, got[k], v)
		}
	}
	if len(l.Names) != 4 {
		t.Errorf("lock holds %d names, want the 4 assigned", len(l.Names))
	}
}

// TestALockedShortNameNeverMoves. J3: a new type whose segment clashes with a locked short name gets the qualified
// form, and the locked type keeps its name. The fixture's lock would otherwise be rewritten by the uniqueness rule.
func TestALockedShortNameNeverMoves(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::EC2::VPC": "aws.vpc"}}
	got, err := l.Assign([]string{"AWS::EC2::VPC", "AWS::NetworkFirewall::VPC"})
	if err != nil {
		t.Fatal(err)
	}
	if got["AWS::EC2::VPC"] != "aws.vpc" || got["AWS::NetworkFirewall::VPC"] != "aws.networkfirewall.vpc" {
		t.Fatalf("assigned %v", got)
	}
}

// TestARemovedTypeKeepsItsNameReserved. A tombstone means a name never comes back meaning something else.
func TestARemovedTypeKeepsItsNameReserved(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::Old::Widget": "aws.widget"}}
	got, err := l.Assign([]string{"AWS::New::Widget"})
	if err != nil {
		t.Fatal(err)
	}
	if got["AWS::New::Widget"] != "aws.new.widget" {
		t.Errorf("new widget = %q, want the qualified form: aws.widget belongs to a removed type", got["AWS::New::Widget"])
	}
	if _, kept := l.Names["AWS::Old::Widget"]; !kept {
		t.Error("the tombstone was dropped")
	}
	if _, returned := got["AWS::Old::Widget"]; returned {
		t.Error("a removed type was returned as current")
	}
}

func TestTwoTypesCannotShareAName(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::A::Thing": "aws.b.thing"}}
	if _, err := l.Assign([]string{"AWS::A::Thing", "AWS::B::Thing"}); err == nil || !strings.Contains(err.Error(), "aws.b.thing") {
		t.Fatalf("err = %v", err)
	}
}

func TestTheLockRoundTripsSorted(t *testing.T) {
	path := filepath.Join(t.TempDir(), "names.lock.json")
	l, err := LoadLock(path)
	if err != nil || len(l.Names) != 0 {
		t.Fatalf("missing lock = %+v, %v", l, err)
	}
	if _, err := l.Assign([]string{"AWS::S3::Bucket", "AWS::EC2::VPC"}); err != nil {
		t.Fatal(err)
	}
	if err := l.Save(path); err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(path)
	if strings.Index(string(raw), "AWS::EC2::VPC") > strings.Index(string(raw), "AWS::S3::Bucket") || !strings.HasSuffix(string(raw), "\n") {
		t.Errorf("lock not sorted or no trailing newline:\n%s", raw)
	}
	again, err := LoadLock(path)
	if err != nil || again.Names["AWS::S3::Bucket"] != "aws.bucket" {
		t.Fatalf("reloaded = %+v, %v", again, err)
	}
}
```

- [ ] **Step 2: Run it to see it fail**

Run: `go test -count=1 ./internal/gen/`
Expected: FAIL — `undefined: Lock`.

- [ ] **Step 3: Implement**

`internal/gen/lock.go`:

```go
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
```

- [ ] **Step 4: Run the tests**

Run: `go test -count=1 ./internal/gen/ && go vet ./internal/gen/`
Expected: PASS.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: ignore the lock (always recompute from uniqueness): `TestALockedShortNameNeverMoves` fails because
`aws.vpc` becomes `aws.ec2.vpc`. Ignore tombstones when building `owner`: the removed-type test fails. Drop the
`taken` error: the shared-name test fails.

```bash
git add internal/gen/lock.go internal/gen/lock_test.go
git commit -m "Assign type names and lock them

Short names like aws.vpc when the resource name is unique, service names
when not, and a lock so an assigned name never changes or gets reused.
Checked by ignoring the lock, ignoring removed types and allowing a shared name." -- \
  internal/gen/lock.go internal/gen/lock_test.go
```

---
### Task 4: From one schema to one catalog type — flags, spellings, shapes, overlay

**Files:**
- Create: `internal/gen/overlay.go`, `internal/gen/build.go`, `internal/gen/shape.go`
- Test: `internal/gen/build_test.go`, `internal/gen/testdata/overlay.yaml`
- Modify: `go.mod` (require `gopkg.in/yaml.v3 v3.0.1` directly; it is already in infrata's module graph)

**Interfaces:**
- Consumes: `cfn.Schema`, `cfn.SnakeCase`, `(*cfn.Schema).Resolve/TopLevel/Nested/Timeout/HasHandler/ListNeedsModel` (Task 1); `catalog.Type`, `catalog.Attribute`, `catalog.Shape`, `catalog.Requirement`, shape kind constants (Task 2).
- Produces:
  - `type Overlay struct { Global, DiscoverDefault []string; Aliases map[string]map[string][]string; Sensitive map[string][]string; Requirements map[string][]OverlayRequirement }`
  - `type OverlayRequirement struct { Name string; Types []string; Description string }`
  - `func LoadOverlay(path string) (*Overlay, error)`, `func (o *Overlay) IsGlobal(cfn string) bool`
  - `func BuildType(s *cfn.Schema, names map[string]string, o *Overlay) (*catalog.Type, []string, error)` (the type, warnings, error)
  - `func BuildShape(s *cfn.Schema, n *cfn.Node) *catalog.Shape`

- [ ] **Step 1: Write the test overlay and the failing tests**

`internal/gen/testdata/overlay.yaml`:

```yaml
# Test overlay: exercises every section the generator reads.
global:
  - "AWS::IAM::*"
discover_default:
  - AWS::EC2::VPC
aliases:
  AWS::EC2::VPC:
    CidrBlock: [cidr]
sensitive:
  AWS::RDS::DBInstance: [MasterUserPassword]
requirements:
  AWS::EC2::Subnet:
    - name: vpc
      types: [AWS::EC2::VPC]
      description: A subnet must be created inside a VPC
```

`internal/gen/build_test.go`:

```go
package gen

import (
	"os"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

var testNames = map[string]string{
	"AWS::EC2::VPC": "aws.vpc", "AWS::EC2::Subnet": "aws.subnet", "AWS::EC2::SecurityGroup": "aws.securitygroup",
	"AWS::S3::Bucket": "aws.bucket", "AWS::IAM::Role": "aws.role", "AWS::RDS::DBInstance": "aws.dbinstance",
	"AWS::ACMPCA::Certificate": "aws.acmpca.certificate", "AWS::ACMPCA::CertificateAuthority": "aws.certificateauthority",
	"AWS::APS::AnomalyDetector": "aws.anomalydetector",
	"AWS::ARCZonalShift::AutoshiftObserverNotificationStatus": "aws.autoshiftobservernotificationstatus",
	"AWS::CodePipeline::CustomActionType": "aws.customactiontype",
}

var fixtureFiles = map[string]string{
	"AWS::EC2::VPC": "aws-ec2-vpc.json", "AWS::EC2::Subnet": "aws-ec2-subnet.json",
	"AWS::EC2::SecurityGroup": "aws-ec2-securitygroup.json", "AWS::S3::Bucket": "aws-s3-bucket.json",
	"AWS::IAM::Role": "aws-iam-role.json", "AWS::RDS::DBInstance": "aws-rds-dbinstance.json",
	"AWS::ACMPCA::Certificate": "aws-acmpca-certificate.json",
	"AWS::ACMPCA::CertificateAuthority": "aws-acmpca-certificateauthority.json",
	"AWS::APS::AnomalyDetector": "aws-aps-anomalydetector.json",
	"AWS::ARCZonalShift::AutoshiftObserverNotificationStatus": "aws-arczonalshift-autoshiftobservernotificationstatus.json",
	"AWS::CodePipeline::CustomActionType": "aws-codepipeline-customactiontype.json",
}

func overlay(t *testing.T) *Overlay {
	t.Helper()
	o, err := LoadOverlay("testdata/overlay.yaml")
	if err != nil {
		t.Fatal(err)
	}
	return o
}

func build(t *testing.T, cfnType string) *catalog.Type {
	t.Helper()
	raw, err := os.ReadFile("../cfn/testdata/" + fixtureFiles[cfnType])
	if err != nil {
		t.Fatal(err)
	}
	s, err := cfn.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	typ, _, err := BuildType(s, testNames, overlay(t))
	if err != nil {
		t.Fatalf("%s: %v", cfnType, err)
	}
	return typ
}

func attr(t *testing.T, typ *catalog.Type, name string) *catalog.Attribute {
	t.Helper()
	a, ok := typ.Attribute(name)
	if !ok {
		t.Fatalf("%s has no attribute %s", typ.Name, name)
	}
	return a
}

// TestEveryFixtureBuildsADefinitionInfrataAccepts: schema.Validate checks flag combinations and case-folded name
// collisions, which is exactly where a generator bug would show.
func TestEveryFixtureBuildsADefinitionInfrataAccepts(t *testing.T) {
	for cfnType := range fixtureFiles {
		if err := build(t, cfnType).Definition().Validate(); err != nil {
			t.Errorf("%s: %v", cfnType, err)
		}
	}
}

func TestVPCFlagsSpellingsAndTags(t *testing.T) {
	vpc := build(t, "AWS::EC2::VPC")
	cidr := attr(t, vpc, "CidrBlock")
	if !cidr.Optional || !cidr.Computed || !cidr.ForceNew || cidr.Required {
		t.Errorf("CidrBlock = %+v, want Optional+Computed+ForceNew", cidr)
	}
	if strings.Join(cidr.Aliases, ",") != "cidr,cidr_block" {
		t.Errorf("CidrBlock aliases = %v, want the curated alias first, then snake_case", cidr.Aliases)
	}
	if id := attr(t, vpc, "VpcId"); !id.Computed || id.Optional || id.ForceNew {
		t.Errorf("VpcId = %+v, want Computed only", id)
	}
	if dns := attr(t, vpc, "EnableDnsHostnames"); dns.ForceNew || !dns.Optional || strings.Join(dns.Aliases, ",") != "enable_dns_hostnames" {
		t.Errorf("EnableDnsHostnames = %+v", dns)
	}
	if vpc.TagsAsMap != "Tags" || attr(t, vpc, "Tags").Kind != "map" {
		t.Errorf("tags as map = %q, kind %q", vpc.TagsAsMap, attr(t, vpc, "Tags").Kind)
	}
	if vpc.RegionAttr != "region" || strings.Join(vpc.Identifier, ",") != "VpcId" || !vpc.HasUpdate || !vpc.HasList {
		t.Errorf("vpc = %+v", vpc)
	}
}

func TestSubnetRequiresAVPCAndLetsAWSPickTheZone(t *testing.T) {
	subnet := build(t, "AWS::EC2::Subnet")
	if v := attr(t, subnet, "VpcId"); !v.Required || !v.ForceNew || v.Computed {
		t.Errorf("VpcId = %+v", v)
	}
	if az := attr(t, subnet, "AvailabilityZone"); !az.Optional || !az.Computed || !az.ForceNew {
		t.Errorf("AvailabilityZone = %+v, want Optional+Computed+ForceNew", az)
	}
	if len(subnet.Requirements) != 1 || subnet.Requirements[0].Types[0] != "aws.vpc" {
		t.Errorf("requirements = %+v, want the overlay's, in infrata names", subnet.Requirements)
	}
}

func TestSecurityGroupRulesAreUnorderedObjects(t *testing.T) {
	sh := attr(t, build(t, "AWS::EC2::SecurityGroup"), "SecurityGroupIngress").Shape
	if sh == nil || sh.Kind != catalog.ShapeArray || !sh.Unordered {
		t.Fatalf("ingress shape = %+v, want an unordered array", sh)
	}
	if sh.Item.Kind != catalog.ShapeObject || sh.Item.Props["CidrIp"] == nil || sh.Item.Props["CidrIp"].Kind != catalog.ShapeScalar {
		t.Errorf("ingress item = %+v", sh.Item)
	}
}

func TestGlobalSensitiveCompositeAndListModel(t *testing.T) {
	if role := build(t, "AWS::IAM::Role"); !role.Global() {
		t.Error("the overlay makes AWS::IAM::* global")
	}
	rds := build(t, "AWS::RDS::DBInstance")
	if !attr(t, rds, "MasterUserPassword").Sensitive {
		t.Error("the overlay marks MasterUserPassword sensitive")
	}
	if !strings.Contains(strings.Join(rds.WriteOnly, ","), "MasterUserPassword") {
		t.Errorf("write-only = %v", rds.WriteOnly)
	}
	if n := len(build(t, "AWS::ACMPCA::Certificate").Identifier); n < 2 {
		t.Errorf("composite identifier parts = %d", n)
	}
	if !build(t, "AWS::APS::AnomalyDetector").ListNeedsModel {
		t.Error("AnomalyDetector lists under a parent")
	}
}

// TestPropertiesNamedLikeInfrataKeywordsGetUsableNames: J9.
func TestPropertiesNamedLikeInfrataKeywordsGetUsableNames(t *testing.T) {
	ca := build(t, "AWS::ACMPCA::CertificateAuthority")
	typ := attr(t, ca, "Type")
	if len(typ.Aliases) == 0 || typ.Aliases[0] != "type_value" || ca.Definition().Display("Type") != "type_value" {
		t.Errorf("Type aliases = %v", typ.Aliases)
	}
	for _, a := range typ.Aliases {
		if strings.EqualFold(a, "type") {
			t.Errorf("alias %q would decode as infrata's resource key", a)
		}
	}
	if p := attr(t, build(t, "AWS::CodePipeline::CustomActionType"), "Provider"); p.Aliases[0] != "provider_value" {
		t.Errorf("Provider aliases = %v", p.Aliases)
	}
	zs := build(t, "AWS::ARCZonalShift::AutoshiftObserverNotificationStatus")
	if zs.RegionAttr != "aws_region" {
		t.Errorf("region attribute on a type with its own Region = %q, want aws_region", zs.RegionAttr)
	}
}

func TestOverlayRefusesUnknownKeys(t *testing.T) {
	path := t.TempDir() + "/bad.yaml"
	if err := os.WriteFile(path, []byte("globals:\n  - AWS::IAM::*\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadOverlay(path); err == nil {
		t.Fatal("a misspelt overlay key was accepted; it would silently apply nothing")
	}
}
```

- [ ] **Step 2: Run them to see them fail**

Run: `go test -count=1 -run 'Fixture|VPC|Subnet|Security|Global|Keyword|Overlay' ./internal/gen/`
Expected: FAIL — `undefined: LoadOverlay`.

- [ ] **Step 3: Implement**

```bash
go get gopkg.in/yaml.v3@v3.0.1
```

`internal/gen/overlay.go`:

```go
package gen

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Overlay is the hand-maintained part of generation: what CloudFormation schemas cannot say (spec §3.2).
type Overlay struct {
	Global          []string                       `yaml:"global"`
	DiscoverDefault []string                       `yaml:"discover_default"`
	Aliases         map[string]map[string][]string `yaml:"aliases"`
	Sensitive       map[string][]string            `yaml:"sensitive"`
	Requirements    map[string][]OverlayRequirement `yaml:"requirements"`
}

// OverlayRequirement is a pre-flight hint, written with CloudFormation type names.
type OverlayRequirement struct {
	Name        string   `yaml:"name"`
	Types       []string `yaml:"types"`
	Description string   `yaml:"description"`
}

// LoadOverlay reads the overlay, refusing unknown keys: a misspelt section would silently apply nothing.
func LoadOverlay(path string) (*Overlay, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	var o Overlay
	if err := dec.Decode(&o); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &o, nil
}

// IsGlobal matches a CloudFormation type against the global patterns: an exact name, or a prefix ending in "*".
func (o *Overlay) IsGlobal(cfn string) bool {
	for _, p := range o.Global {
		if prefix, ok := strings.CutSuffix(p, "*"); ok && strings.HasPrefix(cfn, prefix) {
			return true
		}
		if p == cfn {
			return true
		}
	}
	return false
}
```

`internal/gen/shape.go`:

```go
package gen

import (
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

const maxShapeDepth = 12

// BuildShape resolves a property's nested structure for reconciliation. Anything without a fixed set of named
// properties is opaque (plan decision P5).
func BuildShape(s *cfn.Schema, n *cfn.Node) *catalog.Shape {
	return buildShape(s, n, 0)
}

func buildShape(s *cfn.Schema, orig *cfn.Node, depth int) *catalog.Shape {
	n := s.Resolve(orig)
	if n == nil || depth > maxShapeDepth {
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	}
	if len(n.OneOf)+len(n.AnyOf)+len(n.AllOf) > 0 || len(n.Type) > 1 {
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	}
	typ := ""
	if len(n.Type) == 1 {
		typ = n.Type[0]
	} else if n.Items != nil {
		typ = "array"
	} else if len(n.Properties) > 0 || len(n.PatternProperties) > 0 {
		typ = "object"
	}
	switch typ {
	case "array":
		sh := &catalog.Shape{Kind: catalog.ShapeArray, Unordered: unordered(orig) || unordered(n)}
		if n.Items == nil {
			sh.Item = &catalog.Shape{Kind: catalog.ShapeOpaque}
		} else {
			sh.Item = buildShape(s, n.Items, depth+1)
		}
		return sh
	case "object":
		if len(n.Properties) == 0 || len(n.PatternProperties) > 0 {
			return &catalog.Shape{Kind: catalog.ShapeOpaque}
		}
		sh := &catalog.Shape{Kind: catalog.ShapeObject, Props: make(map[string]*catalog.Shape, len(n.Properties))}
		for name, child := range n.Properties {
			sh.Props[name] = buildShape(s, child, depth+1)
		}
		return sh
	case "":
		return &catalog.Shape{Kind: catalog.ShapeOpaque}
	default:
		return &catalog.Shape{Kind: catalog.ShapeScalar}
	}
}

// unordered reports insertionOrder: false. The schema default is true.
func unordered(n *cfn.Node) bool {
	return n != nil && n.InsertionOrder != nil && !*n.InsertionOrder
}
```

`internal/gen/build.go`:

```go
package gen

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

// infrataKeys are the resource keys infrata's configuration decoder claims, matched exactly
// (internal/config/decode.go at v0.3.0). A property whose lower-case spelling is one of them needs another name to be
// shown and written by (J9).
var infrataKeys = map[string]bool{"type": true, "provider": true, "lifecycle": true, "depends_on": true, "skip": true, "only": true}

var secretName = regexp.MustCompile(`(?i)password|secret|token|privatekey|credential`)

// BuildType maps one provisionable schema to a catalog type (spec §3.1).
func BuildType(s *cfn.Schema, names map[string]string, o *Overlay) (*catalog.Type, []string, error) {
	name, ok := names[s.TypeName]
	if !ok {
		return nil, nil, fmt.Errorf("%s has no assigned name", s.TypeName)
	}
	var warnings []string
	readOnly := s.TopLevel(s.ReadOnlyProperties)
	createOnly := s.TopLevel(s.CreateOnlyProperties)
	writeOnly := s.TopLevel(s.WriteOnlyProperties)
	required := map[string]bool{}
	for _, r := range s.Required {
		required[r] = true
	}
	sensitive := map[string]bool{}
	for _, p := range o.Sensitive[s.TypeName] {
		if s.Properties[p] == nil {
			return nil, nil, fmt.Errorf("overlay marks %s.%s sensitive, but the schema has no such property", s.TypeName, p)
		}
		sensitive[p] = true
	}
	for p := range o.Aliases[s.TypeName] {
		if s.Properties[p] == nil {
			return nil, nil, fmt.Errorf("overlay aliases %s.%s, but the schema has no such property", s.TypeName, p)
		}
	}

	t := &catalog.Type{
		Name: name, CFN: s.TypeName, Description: firstLine(s.Description),
		HasUpdate: s.HasHandler("update"), HasList: s.HasHandler("list"), ListNeedsModel: s.ListNeedsModel(),
		Timeouts: map[string]int{"create": s.Timeout("create"), "update": s.Timeout("update"), "delete": s.Timeout("delete")},
	}
	if !o.IsGlobal(s.TypeName) {
		t.RegionAttr = "region"
		for p := range s.Properties {
			if strings.EqualFold(p, "region") || cfn.SnakeCase(p) == "region" {
				t.RegionAttr = "aws_region"
			}
		}
	}
	for _, p := range s.PrimaryIdentifier {
		t.Identifier = append(t.Identifier, strings.TrimPrefix(p, "/properties/"))
	}
	for p := range writeOnly {
		t.WriteOnly = append(t.WriteOnly, p)
	}
	sort.Strings(t.WriteOnly)
	if n := len(s.Nested(s.CreateOnlyProperties)); n > 0 {
		warnings = append(warnings, fmt.Sprintf("%s: %d nested create-only pointers cannot be expressed", s.TypeName, n))
	}
	tagProp := tagsAsMap(s)
	t.TagsAsMap = tagProp

	props := make([]string, 0, len(s.Properties))
	for p := range s.Properties {
		props = append(props, p)
	}
	sort.Strings(props)
	for _, p := range props {
		node := s.Properties[p]
		kind, guessed := kindOf(s, node)
		if guessed {
			warnings = append(warnings, fmt.Sprintf("%s.%s: kind guessed as %s", s.TypeName, p, kind))
		}
		a := &catalog.Attribute{Name: p, Kind: kind, Description: firstLine(s.Resolve(node).descriptionOr(node.Description)), Sensitive: sensitive[p]}
		switch {
		case readOnly[p]:
			a.Computed = true
		case required[p]:
			a.Required = true
		default:
			a.Optional, a.Computed = true, true
		}
		a.ForceNew = !readOnly[p] && (createOnly[p] || !t.HasUpdate)
		if writeOnly[p] && !sensitive[p] && secretName.MatchString(p) {
			warnings = append(warnings, fmt.Sprintf("%s.%s: write-only and looks secret; consider the overlay's sensitive list", s.TypeName, p))
		}
		if p == tagProp {
			a.Kind = "map"
			a.Shape = &catalog.Shape{Kind: catalog.ShapeOpaque}
		} else if kind == "list" || kind == "map" {
			a.Shape = BuildShape(s, node)
		}
		a.Aliases = aliases(p, o.Aliases[s.TypeName][p], t.RegionAttr)
		t.Attributes = append(t.Attributes, a)
	}

	for _, r := range o.Requirements[s.TypeName] {
		req := catalog.Requirement{Name: r.Name, Description: r.Description}
		for _, cfnType := range r.Types {
			n, ok := names[cfnType]
			if !ok {
				return nil, nil, fmt.Errorf("overlay requirement on %s names %s, which is not a generated type", s.TypeName, cfnType)
			}
			req.Types = append(req.Types, n)
		}
		t.Requirements = append(t.Requirements, req)
	}

	if err := t.Definition().Validate(); err != nil {
		return nil, warnings, fmt.Errorf("%s: generated definition refused by infrata: %w", s.TypeName, err)
	}
	return t, warnings, nil
}

// aliases orders a property's spellings: curated first (Display shows the first), then a keyword-safe name when the
// property's own spelling is an infrata key, then snake_case. Duplicates under case folding, and anything that folds to
// the canonical name, an infrata key or the region attribute, are dropped.
func aliases(prop string, curated []string, regionAttr string) []string {
	snake := cfn.SnakeCase(prop)
	var candidates []string
	candidates = append(candidates, curated...)
	if infrataKeys[strings.ToLower(prop)] || infrataKeys[snake] {
		candidates = append(candidates, snake+"_value")
	}
	candidates = append(candidates, snake)
	seen := map[string]bool{strings.ToLower(prop): true}
	var out []string
	for _, c := range candidates {
		f := strings.ToLower(c)
		if seen[f] || infrataKeys[f] || (regionAttr != "" && f == strings.ToLower(regionAttr)) {
			continue
		}
		seen[f] = true
		out = append(out, c)
	}
	return out
}

// tagsAsMap returns the tag property's name when it is a list of {Key, Value} objects, which the provider exposes as
// a map (spec §3.4).
func tagsAsMap(s *cfn.Schema) string {
	if s.Tagging == nil || (s.Tagging.Taggable != nil && !*s.Tagging.Taggable) {
		return ""
	}
	name, ok := strings.CutPrefix(s.Tagging.TagProperty, "/properties/")
	if !ok || strings.Contains(name, "/") {
		return ""
	}
	prop := s.Resolve(s.Properties[name])
	if prop == nil || prop.Items == nil {
		return ""
	}
	item := s.Resolve(prop.Items)
	if item == nil || len(item.Properties) != 2 || item.Properties["Key"] == nil || item.Properties["Value"] == nil {
		return ""
	}
	return name
}

// kindOf maps a property's JSON-schema type to an infrata kind name, reporting whether it had to guess.
func kindOf(s *cfn.Schema, node *cfn.Node) (string, bool) {
	n := s.Resolve(node)
	if n == nil {
		return "string", true
	}
	if len(n.Type) == 1 {
		switch n.Type[0] {
		case "string":
			return "string", false
		case "integer":
			return "integer", false
		case "number":
			return "float", false
		case "boolean":
			return "boolean", false
		case "array":
			return "list", false
		case "object":
			return "map", false
		}
	}
	if len(n.Type) == 0 {
		if n.Items != nil {
			return "list", false
		}
		if len(n.Properties) > 0 || len(n.PatternProperties) > 0 {
			return "map", false
		}
	}
	return "string", true
}

func firstLine(s string) string {
	s, _, _ = strings.Cut(strings.TrimSpace(s), "\n")
	return s
}
```

Add to `internal/cfn/schema.go` (the description of a resolved definition is often empty while the referring property
carries one):

```go
// DescriptionOr returns the node's description, or fallback when the node is nil or has none.
func (n *Node) DescriptionOr(fallback string) string {
	if n == nil || n.Description == "" {
		return fallback
	}
	return n.Description
}
```

- [ ] **Step 4: Run the tests**

Run: `go test -count=1 ./internal/gen/ ./internal/cfn/ && go vet ./internal/gen/ ./internal/cfn/`
Expected: PASS. If `TestEveryFixtureBuildsADefinitionInfrataAccepts` fails on a fold collision, print the
definition's attribute names and aliases: the aliases rule, not the fixture, is wrong.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: mark every non-read-only property plainly optional (drop `a.Optional, a.Computed = true, true`): the VPC
and fixture-validation tests fail; put snake_case before curated aliases: the VPC aliases test fails; skip the
keyword rule: the keyword test fails (and `Validate` still passes, which is why the test exists); treat
`insertionOrder` as ordered by default: the security-group test fails; drop the `aws_region` switch: the keyword test's
region assertion fails and `Validate` refuses the fixture.

```bash
git add go.mod go.sum internal/cfn/schema.go internal/gen/overlay.go internal/gen/build.go internal/gen/shape.go \
  internal/gen/build_test.go internal/gen/testdata/overlay.yaml
git commit -m "Build catalog types from CloudFormation schemas

Maps schema flags to infrata attributes, orders aliases so the friendly name
shows, renames properties that clash with infrata keywords or the region
attribute, and records nested shapes. Checked by breaking each rule." -- \
  go.mod go.sum internal/cfn/schema.go internal/gen/overlay.go internal/gen/build.go internal/gen/shape.go \
  internal/gen/build_test.go internal/gen/testdata/overlay.yaml
```

---
### Task 5: Generate the real catalog

**Files:**
- Create: `internal/gen/generate.go`, `cmd/gen-cloudcontrol/main.go`, `gen/overlay.yaml`
- Create (generated, committed): `gen/names.lock.json`, `gen/warnings.txt`, `internal/catalog/catalog.json.gz`
- Create: `internal/catalog/embed.go`
- Test: `internal/gen/generate_test.go`, `internal/gen/committed_test.go`, `internal/catalog/embed_test.go`

**Interfaces:**
- Consumes: `cfn.ReadBundle`, `cfn.Schema` (Task 1); `catalog.Catalog` including `DiscoverDefault` (Task 2); `Lock`, `LoadLock`, `(*Lock).Assign`, `(*Lock).Save` (Task 3); `Overlay`, `LoadOverlay`, `BuildType` (Task 4).
- Produces:
  - `func Generate(schemas []*cfn.Schema, bundleSHA string, lock *Lock, o *Overlay) (*catalog.Catalog, []string, error)`
  - `func Embedded() (*catalog.Catalog, error)` in package `catalog` (loaded once)
  - `gen/overlay.yaml` and `gen/names.lock.json`, which later regenerations extend

- [ ] **Step 1: Write the real overlay**

`gen/overlay.yaml` (property names checked against the us-east-1 bundle, 2026-09-14; `AWS::IAM::Policy` is not
provisionable, so the core set uses `AWS::IAM::ManagedPolicy`; `AWS::ECS::Service` gets no `name` alias because it
already has a read-only `Name` property):

```yaml
# The hand-maintained part of generation (spec §3.2). Every name here is validated against the schema bundle; a
# misspelt type or property fails generation.

# Types that are not regional. They get no region attribute, their provider IDs start with "global/", and Cloud
# Control is called in us-east-1 for them.
global:
  - "AWS::IAM::*"
  - "AWS::Organizations::*"
  - "AWS::CloudFront::*"
  - "AWS::Route53::*"

# What `discover`, and `import` (which discovers first), look at when an instance sets no discover_types (P4).
discover_default:
  - AWS::EC2::VPC
  - AWS::EC2::Subnet
  - AWS::EC2::SecurityGroup
  - AWS::EC2::InternetGateway
  - AWS::EC2::RouteTable
  - AWS::EC2::Route
  - AWS::EC2::Instance
  - AWS::S3::Bucket
  - AWS::IAM::Role
  - AWS::IAM::ManagedPolicy
  - AWS::RDS::DBInstance
  - AWS::RDS::DBSubnetGroup
  - AWS::Lambda::Function
  - AWS::ECS::Cluster
  - AWS::ECS::Service

# Friendly names, shown in plans (J5, J10). Listed first, so infrata's Display picks them.
aliases:
  AWS::EC2::VPC:
    CidrBlock: [cidr]
  AWS::EC2::Subnet:
    CidrBlock: [cidr]
    AvailabilityZone: [az]
  AWS::EC2::SecurityGroup:
    GroupName: [name]
    GroupDescription: [description]
    SecurityGroupIngress: [ingress]
    SecurityGroupEgress: [egress]
  AWS::EC2::Route:
    DestinationCidrBlock: [destination_cidr]
  AWS::EC2::Instance:
    ImageId: [ami]
  AWS::S3::Bucket:
    BucketName: [name]
  AWS::IAM::Role:
    RoleName: [name]
    AssumeRolePolicyDocument: [assume_role_policy]
  AWS::IAM::ManagedPolicy:
    ManagedPolicyName: [name]
    PolicyDocument: [policy]
  AWS::RDS::DBInstance:
    DBInstanceIdentifier: [identifier]
    DBInstanceClass: [instance_class]
    MasterUsername: [username]
    MasterUserPassword: [password]
  AWS::RDS::DBSubnetGroup:
    DBSubnetGroupName: [name]
    DBSubnetGroupDescription: [description]
  AWS::Lambda::Function:
    FunctionName: [name]
  AWS::ECS::Cluster:
    ClusterName: [name]

# Secrets. Seeded from the write-only properties whose names look secret, excluding client tokens, ARNs and
# on/off flags. The generator warns about new candidates in gen/warnings.txt.
sensitive:
  AWS::Amplify::App: [AccessToken, OauthToken]
  AWS::BackupGateway::Hypervisor: [Password]
  AWS::CodeBuild::SourceCredential: [Token]
  AWS::Connect::User: [Password]
  AWS::DMS::Endpoint: [Password]
  AWS::DataSync::LocationFSxWindows: [Password]
  AWS::DataSync::LocationObjectStorage: [SecretKey]
  AWS::DataSync::LocationSMB: [Password]
  AWS::DirectoryService::SimpleAD: [Password]
  AWS::DocDBElastic::Cluster: [AdminUserPassword]
  AWS::ElastiCache::ReplicationGroup: [AuthToken]
  AWS::ElastiCache::User: [Passwords]
  AWS::IAM::SAMLProvider: [AddPrivateKey]
  AWS::IAM::ServerCertificate: [PrivateKey]
  AWS::Invoicing::ProcurementPortalPreference: [ProcurementPortalSharedSecret]
  AWS::Lightsail::Database: [MasterUserPassword]
  AWS::QuickSight::DataSource: [Credentials]
  AWS::QuickSight::OAuthClientApplication: [ClientSecret]
  AWS::RDS::DBCluster: [MasterUserPassword]
  AWS::RDS::DBInstance: [MasterUserPassword, TdeCredentialPassword]
  AWS::Redshift::Cluster: [MasterUserPassword]
  AWS::RedshiftServerless::Namespace: [AdminUserPassword]
  AWS::SecretsManager::Secret: [SecretString]
  AWS::SystemsManagerSAP::Application: [Credentials]
  AWS::Timestream::InfluxDBCluster: [Password]
  AWS::Timestream::InfluxDBInstance: [Password]
  AWS::Transfer::Certificate: [PrivateKey]

requirements:
  AWS::EC2::Subnet:
    - name: vpc
      types: [AWS::EC2::VPC]
      description: A subnet must be created inside a VPC
  AWS::EC2::RouteTable:
    - name: vpc
      types: [AWS::EC2::VPC]
      description: A route table belongs to a VPC
  AWS::EC2::Route:
    - name: route_table
      types: [AWS::EC2::RouteTable]
      description: A route belongs to a route table
  AWS::RDS::DBSubnetGroup:
    - name: subnets
      types: [AWS::EC2::Subnet]
      description: A DB subnet group needs subnets
```

- [ ] **Step 2: Write the failing tests**

`internal/gen/generate_test.go`:

```go
package gen

import (
	"os"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

func fixtureSchemas(t *testing.T) []*cfn.Schema {
	t.Helper()
	var out []*cfn.Schema
	for _, file := range fixtureFiles {
		raw, err := os.ReadFile("../cfn/testdata/" + file)
		if err != nil {
			t.Fatal(err)
		}
		s, err := cfn.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, s)
	}
	return out
}

func TestGenerateNamesBuildsAndKeepsTheBundleHash(t *testing.T) {
	cat, _, err := Generate(fixtureSchemas(t), "sha-of-bundle", &Lock{}, overlay(t))
	if err != nil {
		t.Fatal(err)
	}
	if cat.Bundle != "sha-of-bundle" || len(cat.Types) != len(fixtureFiles) {
		t.Fatalf("bundle=%q types=%d", cat.Bundle, len(cat.Types))
	}
	vpc, ok := cat.Lookup("aws.vpc")
	if !ok || vpc.CFN != "AWS::EC2::VPC" {
		t.Fatalf("aws.vpc = %+v", vpc)
	}
	if strings.Join(cat.DiscoverDefault, ",") != "aws.vpc" {
		t.Errorf("discover default = %v, want the overlay's list in infrata names", cat.DiscoverDefault)
	}
}

func TestGenerateSkipsTypesCloudControlCannotManage(t *testing.T) {
	schemas := fixtureSchemas(t)
	nonProvisionable, _ := cfn.Parse([]byte(`{"typeName":"AWS::AppMesh::Mesh","properties":{"MeshName":{"type":"string"}},"handlers":{"read":{}}}`))
	cat, _, err := Generate(append(schemas, nonProvisionable), "x", &Lock{}, overlay(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range cat.Types {
		if typ.CFN == "AWS::AppMesh::Mesh" {
			t.Fatal("a type without create/read/delete handlers was generated")
		}
	}
}

func TestGenerateRefusesAnOverlayNamingAnUnknownType(t *testing.T) {
	o := overlay(t)
	o.Aliases["AWS::EC2::VPCC"] = map[string][]string{"CidrBlock": {"cidr"}}
	if _, _, err := Generate(fixtureSchemas(t), "x", &Lock{}, o); err == nil || !strings.Contains(err.Error(), "AWS::EC2::VPCC") {
		t.Fatalf("err = %v", err)
	}
}

```

`internal/catalog/embed_test.go`:

```go
package catalog

import (
	"encoding/json"
	"testing"
)

func TestTheEmbeddedCatalogLoadsAndValidates(t *testing.T) {
	cat, err := Embedded()
	if err != nil {
		t.Fatal(err)
	}
	if len(cat.Types) < 1500 {
		t.Fatalf("embedded catalog has %d types; expected the ~1,584 Cloud Control supports", len(cat.Types))
	}
	defs := cat.Definitions()
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Fatalf("%s: %v", d.Type, err)
		}
	}
	for name, cfn := range map[string]string{
		"aws.vpc": "AWS::EC2::VPC", "aws.subnet": "AWS::EC2::Subnet", "aws.securitygroup": "AWS::EC2::SecurityGroup",
		"aws.role": "AWS::IAM::Role", "aws.ec2.instance": "AWS::EC2::Instance", "aws.s3.bucket": "AWS::S3::Bucket",
	} {
		if typ, ok := cat.Lookup(name); !ok || typ.CFN != cfn {
			t.Errorf("%s = %+v", name, typ)
		}
	}
	for _, name := range cat.DiscoverDefault {
		if _, ok := cat.Lookup(name); !ok {
			t.Errorf("discover default names %s, which the catalog lacks", name)
		}
	}
	raw, _ := json.Marshal(defs)
	t.Logf("%d definitions, %d bytes as JSON", len(defs), len(raw))
}
```

- [ ] **Step 3: Run them to see them fail**

Run: `go test -count=1 ./internal/gen/`
Expected: FAIL — `undefined: Generate`.

(`internal/catalog/embed_test.go` is written in Step 2 but committed with `embed.go` in Step 6: `go:embed` cannot compile until
the generator has written `catalog.json.gz`. Keep it aside until then: write it to a scratch file, not into the package.)

- [ ] **Step 4: Implement the generator**

`internal/gen/generate.go`:

```go
package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

// Generate builds the catalog from every provisionable AWS schema. The lock gains names for new types; saving it is
// the caller's job.
func Generate(schemas []*cfn.Schema, bundleSHA string, lock *Lock, o *Overlay) (*catalog.Catalog, []string, error) {
	var provisionable []*cfn.Schema
	var cfnTypes []string
	for _, s := range schemas {
		if strings.HasPrefix(s.TypeName, "AWS::") && s.Provisionable() {
			provisionable = append(provisionable, s)
			cfnTypes = append(cfnTypes, s.TypeName)
		}
	}
	if err := checkOverlay(o, cfnTypes); err != nil {
		return nil, nil, err
	}
	names, err := lock.Assign(cfnTypes)
	if err != nil {
		return nil, nil, err
	}
	cat := &catalog.Catalog{Bundle: bundleSHA}
	var warnings []string
	for _, s := range provisionable {
		t, w, err := BuildType(s, names, o)
		warnings = append(warnings, w...)
		if err != nil {
			return nil, warnings, err
		}
		cat.Types = append(cat.Types, t)
	}
	for _, cfnType := range o.DiscoverDefault {
		cat.DiscoverDefault = append(cat.DiscoverDefault, names[cfnType])
	}
	sort.Slice(cat.Types, func(i, j int) bool { return cat.Types[i].Name < cat.Types[j].Name })
	sort.Strings(warnings)
	return cat, warnings, nil
}

// checkOverlay refuses an overlay that names a type the bundle does not provide: it would silently apply nothing.
func checkOverlay(o *Overlay, cfnTypes []string) error {
	known := make(map[string]bool, len(cfnTypes))
	for _, t := range cfnTypes {
		known[t] = true
	}
	var unknown []string
	check := func(section, t string) {
		if !known[t] {
			unknown = append(unknown, section+": "+t)
		}
	}
	for t := range o.Aliases {
		check("aliases", t)
	}
	for t := range o.Sensitive {
		check("sensitive", t)
	}
	for t, reqs := range o.Requirements {
		check("requirements", t)
		for _, r := range reqs {
			for _, rt := range r.Types {
				check("requirements types", rt)
			}
		}
	}
	for _, t := range o.DiscoverDefault {
		check("discover_default", t)
	}
	for _, p := range o.Global {
		if !strings.HasSuffix(p, "*") {
			check("global", p)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return fmt.Errorf("overlay names types the bundle does not provide:\n  %s", strings.Join(unknown, "\n  "))
	}
	return nil
}
```

`cmd/gen-cloudcontrol/main.go`:

```go
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
```

- [ ] **Step 5: Generate**

```bash
scripts/fetch-schemas
go test -count=1 ./internal/gen/                    # the generator's unit tests pass
go run ./cmd/gen-cloudcontrol                        # prints: bundle <sha>: 1584 types, 1584 new names, N warnings
ls -la internal/catalog/catalog.json.gz gen/names.lock.json gen/warnings.txt
```

- [ ] **Step 6: Embed the catalog, and test what was generated**

Now that `catalog.json.gz` exists, add the embed and the two tests that read it: move `embed_test.go` from Step 2 into
`internal/catalog/`, and create `internal/gen/committed_test.go`.

`internal/catalog/embed.go`:

```go
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
```

`internal/gen/committed_test.go`:

```go
package gen

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
)

// TestTheCommittedCatalogMatchesTheBundleItClaims regenerates from the real bundle, when present, with the committed
// overlay and lock, and compares with the committed catalog. It skips without the bundle (run scripts/fetch-schemas).
// A different bundle hash means the schemas moved on: regenerate, do not edit this test.
func TestTheCommittedCatalogMatchesTheBundleItClaims(t *testing.T) {
	const bundlePath = "../../schemas/CloudformationSchema.zip"
	if _, err := os.Stat(bundlePath); errors.Is(err, fs.ErrNotExist) {
		t.Skip("no schema bundle; run scripts/fetch-schemas")
	}
	schemas, sha, err := cfn.ReadBundle(bundlePath)
	if err != nil {
		t.Fatal(err)
	}
	committed, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	if committed.Bundle != sha {
		t.Skipf("bundle %s differs from the committed catalog's %s: regenerate with go run ./cmd/gen-cloudcontrol", sha, committed.Bundle)
	}
	lock, err := LoadLock("../../gen/names.lock.json")
	if err != nil {
		t.Fatal(err)
	}
	o, err := LoadOverlay("../../gen/overlay.yaml")
	if err != nil {
		t.Fatal(err)
	}
	before := len(lock.Names)
	cat, _, err := Generate(schemas, sha, lock, o)
	if err != nil {
		t.Fatal(err)
	}
	if len(lock.Names) != before {
		t.Errorf("regeneration assigned %d new names: the committed lock is stale", len(lock.Names)-before)
	}
	if len(cat.Types) != len(committed.Types) {
		t.Errorf("regenerated %d types, committed catalog has %d", len(cat.Types), len(committed.Types))
	}
}
```

```bash
go test -count=1 ./internal/gen/ ./internal/catalog/ ./internal/cfn/
go vet ./... && gofmt -l .
```

Expected: PASS. The embed test logs the definitions' JSON size; record it in this plan's Verification log. If
generation fails with "generated definition refused by infrata", the error names the type and the rule: fix the rule
in `build.go` (or add an overlay entry) and regenerate. Never hand-edit the catalog.

Read `gen/warnings.txt` before committing: guessed kinds and secret-looking properties are listed there. Add any real
secret to the overlay's `sensitive:` and regenerate.

- [ ] **Step 7: Sabotage, then commit**

Sabotages: drop the `Provisionable()` filter (non-provisionable test); drop `checkOverlay` (unknown-type test); build
`DiscoverDefault` from CloudFormation names instead of infrata names (the discover-default assertions in both test
files fail).

```bash
git add internal/gen/generate.go internal/gen/generate_test.go internal/gen/committed_test.go cmd/gen-cloudcontrol/main.go gen/overlay.yaml \
  gen/names.lock.json gen/warnings.txt internal/catalog/embed.go internal/catalog/embed_test.go internal/catalog/catalog.json.gz
git commit -m "Generate the catalog for every Cloud Control type

Adds the generator command, the curated overlay for the core types, the name
lock and the generated catalog for all 1584 types, embedded in the plugin.
Checked by removing the provisionable filter, the overlay check and the
discover default name mapping." -- \
  internal/gen/generate.go internal/gen/generate_test.go internal/gen/committed_test.go cmd/gen-cloudcontrol/main.go gen/overlay.yaml \
  gen/names.lock.json gen/warnings.txt internal/catalog/embed.go internal/catalog/embed_test.go internal/catalog/catalog.json.gz
```

---
### Task 6: The plugin serves the catalog — and the load-cost gate

**Files:**
- Delete: `internal/awsprov/{vpc,subnet,tags,ids,attrs,values,definitions,errors,clients,patience,provider}.go` and
  `internal/awsprov/{vpc,subnet,tags,ids,definitions,errors,helpers}_test.go`; `internal/ec2fake/` (all three files)
- Create: `internal/awstest/awstest.go`, `internal/awsprov/unconfigured.go`, `scripts/measure-load`
- Modify: `go.mod` (infrata v0.3.0), `internal/awsprov/plugin.go`, `internal/awsprov/config.go`,
  `internal/awsprov/config_test.go`, `internal/awsprov/protocol_test.go`
- Delete: `internal/awsprov/testenv_test.go` (moved to `internal/awstest`)

**Interfaces:**
- Consumes: `catalog.Embedded`, `(*catalog.Catalog).Definitions/Lookup` (Tasks 2, 5).
- Produces:
  - `func awstest.Isolate(t testing.TB, endpoint string) string` (sets `AWS_ENDPOINT_URL_CLOUDCONTROL` and
    `AWS_ENDPOINT_URL_STS` to `endpoint`; returns the directory holding the empty AWS config files)
  - `instanceConfig.DiscoverTypes []string`; config key `discover_types`
  - `Plugin.New` validates `discover_types` against the catalog; `unconfigured` is a temporary provider replaced in Task 9

- [ ] **Step 1: Move to infrata v0.3.0 and remove the handwritten resources**

```bash
export GOPRIVATE='github.com/infrata/*'
GOWORK=off go get github.com/infrata/infrata@v0.3.0
git rm -q internal/awsprov/vpc.go internal/awsprov/subnet.go internal/awsprov/tags.go internal/awsprov/ids.go \
  internal/awsprov/attrs.go internal/awsprov/values.go internal/awsprov/definitions.go internal/awsprov/errors.go \
  internal/awsprov/clients.go internal/awsprov/patience.go internal/awsprov/provider.go \
  internal/awsprov/vpc_test.go internal/awsprov/subnet_test.go internal/awsprov/tags_test.go internal/awsprov/ids_test.go \
  internal/awsprov/definitions_test.go internal/awsprov/errors_test.go internal/awsprov/helpers_test.go \
  internal/awsprov/testenv_test.go internal/ec2fake/server.go internal/ec2fake/server_test.go internal/ec2fake/xml.go
```

The removed classification, provider-ID and patience code is rewritten for Cloud Control in Task 8. Task 5 of the
`first-slice` plan is where its reasoning lives.

- [ ] **Step 2: Write the shared test helper**

`internal/awstest/awstest.go`:

```go
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
```

- [ ] **Step 3: Write the failing tests**

In `internal/awsprov/config_test.go`: replace every `isolateAWS(t, …)` with `awstest.Isolate(t, …)` (import
`github.com/infrata/infrata-provider-aws/internal/awstest`), delete `TestAssumeRoleSignsEC2CallsWithTheAssumedCredentials`
(Task 9 restores it against the Cloud Control fake), remove the now-unused `context`, `ec2` and `ec2fake` imports, and add:

```go
func TestDiscoverTypesAreReadAndChecked(t *testing.T) {
	ic, err := parseConfig(map[string]value.Value{"discover_types": list("aws.vpc", "aws.subnet")})
	if err != nil || strings.Join(ic.DiscoverTypes, ",") != "aws.vpc,aws.subnet" {
		t.Fatalf("discover_types = %v, %v", ic.DiscoverTypes, err)
	}
	awstest.Isolate(t, "")
	_, err = NewPlugin().New(provider.Config{Instance: "main", Values: map[string]value.Value{"discover_types": list("aws.vpc", "aws.vpcc")}})
	if err == nil || !strings.Contains(err.Error(), "aws.vpcc") || !strings.Contains(err.Error(), "infrata explain") {
		t.Fatalf("err = %v", err)
	}
}
```

Also extend `TestAnUnknownKeyIsRefusedNamingWhatIsAccepted`'s wanted list with `"discover_types"`.

Replace `internal/awsprov/protocol_test.go` entirely:

```go
package awsprov

import (
	"context"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/plugintest"
)

func openHost(t *testing.T) *plugintest.Host {
	t.Helper()
	host, err := plugintest.Open(context.Background(), NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatalf("the host refused this plugin's schemas: %v", err)
	}
	t.Cleanup(func() { _ = host.Close() })
	return host
}

// TestTheWholeCatalogLoadsThroughTheHost: every generated definition crosses the wire and passes infrata's load checks
// (prefix, reserved names, validation, alias folding) at protocol 2.
func TestTheWholeCatalogLoadsThroughTheHost(t *testing.T) {
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	host := openHost(t)
	if got := host.Version(); got != Version {
		t.Errorf("handshake version = %q, want %q", got, Version)
	}
	defs := host.Definitions()
	if len(defs) != len(cat.Types) {
		t.Fatalf("host holds %d definitions, catalog has %d", len(defs), len(cat.Types))
	}
	for _, d := range defs {
		if d.Type != "aws.vpc" {
			continue
		}
		if got := d.Display("CidrBlock"); got != "cidr" {
			t.Errorf("aws.vpc CidrBlock displays as %q after the wire, want cidr", got)
		}
		return
	}
	t.Fatal("aws.vpc did not arrive")
}
```

- [ ] **Step 4: Run them to see them fail**

Run: `go test -count=1 ./internal/awsprov/`
Expected: FAIL — compile errors (`definitions` undefined in `plugin.go`, `DiscoverTypes` unknown).

- [ ] **Step 5: Implement**

`internal/awsprov/config.go`: add the key and field, and read it exactly like `discover_regions`:

```go
const keyDiscoverTypes = "discover_types"
```

- In `instanceConfig`, add `DiscoverTypes []string`.
- In `parseConfig`'s accepted-key switch, add `keyDiscoverTypes`, and update the error message to
  `"unknown configuration %s; the aws provider accepts only %s, %s, %s and %s (a resource's region belongs under defaults: {region: …}, not here)"`
  with arguments `strings.Join(unknown, ", "), keyAssumeRoleARN, keyDiscoverRegions, keyDiscoverTypes, keyProfile`.
- Replace the `discover_regions` block with a helper used for both keys:

```go
	if ic.DiscoverRegions, err = stringList(values, keyDiscoverRegions, "a region name such as us-east-1"); err != nil {
		return ic, err
	}
	if ic.DiscoverTypes, err = stringList(values, keyDiscoverTypes, "an infrata type name such as aws.vpc"); err != nil {
		return ic, err
	}
```

```go
// stringList reads an optional list of non-empty strings, dropping duplicates in written order.
func stringList(values map[string]value.Value, key, what string) ([]string, error) {
	v, ok := values[key]
	if !ok {
		return nil, nil
	}
	items, isList := v.Raw.([]value.Value)
	if v.Kind != value.KindList || !isList {
		return nil, fmt.Errorf("`%s` must be a list, got %s", key, v.Kind)
	}
	var out []string
	seen := map[string]bool{}
	for i, item := range items {
		s, isString := item.AsString()
		if !isString || s == "" {
			return nil, fmt.Errorf("`%s` item %d must be %s", key, i+1, what)
		}
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out, nil
}
```

`internal/awsprov/unconfigured.go` (removed in Task 9):

```go
package awsprov

import (
	"context"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/schema"
)

// unconfigured stands in for the Cloud Control provider until Task 9 of the plan builds it: a configured instance that
// says so rather than guessing.
type unconfigured struct{ cat *catalog.Catalog }

var _ provider.Provider = unconfigured{}

func (u unconfigured) Name() string                              { return PluginName }
func (u unconfigured) Definitions() []*schema.ResourceDefinition { return u.cat.Definitions() }
func (u unconfigured) Read(context.Context, *resource.ResourceState) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Create(context.Context, *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Update(context.Context, *resource.ResourceState, *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Delete(context.Context, *resource.ResourceState) error { return provider.ErrNotImplemented }
func (u unconfigured) Discover(context.Context, provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Import(context.Context, string, string) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) ClassifyError(error) provider.Retryability { return provider.NotSafeToRetry }
```

`internal/awsprov/plugin.go`: replace `Definitions` and `New`:

```go
// Definitions are the embedded catalog's. They need no configuration and make no network call.
func (pl *Plugin) Definitions() []*schema.ResourceDefinition {
	cat, err := catalog.Embedded()
	if err != nil {
		// Only a broken build gets here; the catalog tests fail first.
		fmt.Fprintln(os.Stderr, "the aws plugin's embedded catalog is unreadable:", err)
		return nil
	}
	return cat.Definitions()
}

// New builds one configured instance.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	ic, err := parseConfig(cfg.Values)
	if err != nil {
		return nil, err
	}
	cat, err := catalog.Embedded()
	if err != nil {
		return nil, err
	}
	var unknown []string
	for _, name := range ic.DiscoverTypes {
		if _, ok := cat.Lookup(name); !ok {
			unknown = append(unknown, strconv.Quote(name))
		}
	}
	if len(unknown) > 0 {
		return nil, fmt.Errorf("`discover_types` names %s, which the aws plugin does not serve; run `infrata explain <type>` to check a name",
			strings.Join(unknown, ", "))
	}
	if _, err := loadAWSConfig(context.Background(), cfg.Instance, ic); err != nil {
		return nil, err
	}
	return unconfigured{cat: cat}, nil
}
```

(imports: `context`, `fmt`, `os`, `strconv`, `strings`, `github.com/infrata/infrata-provider-aws/internal/catalog`,
`github.com/infrata/infrata/pkg/provider`, `github.com/infrata/infrata/pkg/schema`.)

- [ ] **Step 6: Run the suite**

```bash
go mod tidy
go test -count=1 ./... && go vet ./... && gofmt -l .
GOWORK=off GOPRIVATE='github.com/infrata/*' go test -count=1 ./...
```

Expected: PASS in both modes. `TestTheWholeCatalogLoadsThroughTheHost` is the proof that all ~1,584 definitions load.

- [ ] **Step 7: The load-cost gate (P7)**

`scripts/measure-load`:

```bash
#!/usr/bin/env bash
# measure-load: how much longer `infrata validate` takes with this plugin's full catalog than with the fake plugin.
# Plan decision P7: acceptable up to 500 ms extra (median of 11 runs). Exits non-zero above that.
set -euo pipefail
cd "$(dirname "$0")/.."
work="$(mktemp -d)"
trap 'rm -rf "$work"' EXIT
mkdir -p "$work/plugins" "$work/aws" "$work/fake"

gh release download v0.3.0 --repo infrata/infrata --pattern 'infrata_0.3.0_linux_amd64.tar.gz' --dir "$work"
tar -xzf "$work/infrata_0.3.0_linux_amd64.tar.gz" -C "$work"
infrata="$(find "$work" -type f -name infrata -perm -u+x | head -1)"
go build -o "$work/plugins/infrata-plugin-aws" ./cmd/infrata-plugin-aws
go -C ../infrata-provider-fake build -o "$work/plugins/infrata-plugin-fake" ./cmd/infrata-plugin-fake

cat > "$work/aws/infra.yml" <<'EOF'
project: measure
environments:
  dev: {}
providers:
  - plugin: aws
    defaults:
      region: us-east-1
resources:
  vpc:
    type: aws.vpc
    cidr: 10.0.0.0/16
EOF
cat > "$work/fake/infra.yml" <<'EOF'
project: measure
environments:
  dev: {}
resources:
  net:
    type: fake.network
    cidr: 10.0.0.0/16
EOF

export HOME="$work" AWS_CONFIG_FILE=/dev/null AWS_SHARED_CREDENTIALS_FILE=/dev/null \
  AWS_ACCESS_KEY_ID=measure AWS_SECRET_ACCESS_KEY=measure AWS_EC2_METADATA_DISABLED=true INFRATA_PLUGIN_PATH=
median() { sort -n | awk '{v[NR]=$1} END {print v[int((NR+1)/2)]}'; }
timed() {
  local dir="$1"
  for _ in $(seq 11); do
    start=$(date +%s%N)
    (cd "$dir" && "$infrata" validate dev --plugin-dir "$work/plugins" >/dev/null)
    end=$(date +%s%N)
    echo $(( (end - start) / 1000000 ))
  done | median
}
fake_ms="$(timed "$work/fake")"
aws_ms="$(timed "$work/aws")"
extra=$(( aws_ms - fake_ms ))
echo "infrata validate, median of 11: fake plugin ${fake_ms} ms, aws plugin ${aws_ms} ms, extra ${extra} ms (gate 500 ms)"
[ "$extra" -le 500 ]
```

```bash
chmod +x scripts/measure-load
scripts/measure-load
```

Expected: one line with the three numbers, and exit 0. `validate` succeeding with `cidr:` also proves alias
resolution end to end on a real infrata v0.3.0 binary.

**If the gate fails (exit 1): stop the plan here** and report the numbers to James, with the options from spec §5:
drop descriptions from the catalog (measured at about half the size), omit nested shapes from the definitions sent to
infrata (they are provider-side only already; check they are not in `Definition()`), or ask the infrata session about
lazy schema loading. Record the measured numbers in the Verification log either way.

- [ ] **Step 8: Sabotage, then commit**

Sabotages: `Definitions` returns only the first 10 catalog types (the whole-catalog test fails on the count); skip the
`discover_types` lookup (the discover-types test fails); drop `discover_types` from the accepted keys (the parse
assertion fails with "unknown configuration").

```bash
git add go.mod go.sum internal/awstest/awstest.go internal/awsprov/unconfigured.go internal/awsprov/plugin.go \
  internal/awsprov/config.go internal/awsprov/config_test.go internal/awsprov/protocol_test.go scripts/measure-load
git commit -m "Serve the generated catalog and measure what it costs to load

The plugin now offers every Cloud Control type from the embedded catalog on
infrata 0.3.0, and the handwritten vpc and subnet code is gone. Adds
discover_types and a script that times infrata validate against the fake
plugin." -- \
  go.mod go.sum internal/awstest/awstest.go internal/awsprov/unconfigured.go internal/awsprov/plugin.go \
  internal/awsprov/config.go internal/awsprov/config_test.go internal/awsprov/protocol_test.go scripts/measure-load \
  internal/awsprov/vpc.go internal/awsprov/subnet.go internal/awsprov/tags.go internal/awsprov/ids.go \
  internal/awsprov/attrs.go internal/awsprov/values.go internal/awsprov/definitions.go internal/awsprov/errors.go \
  internal/awsprov/clients.go internal/awsprov/patience.go internal/awsprov/provider.go \
  internal/awsprov/vpc_test.go internal/awsprov/subnet_test.go internal/awsprov/tags_test.go internal/awsprov/ids_test.go \
  internal/awsprov/definitions_test.go internal/awsprov/errors_test.go internal/awsprov/helpers_test.go \
  internal/awsprov/testenv_test.go internal/ec2fake/server.go internal/ec2fake/server_test.go internal/ec2fake/xml.go
```

---
### Task 7: `internal/ccfake` — an in-process Cloud Control (and STS) endpoint

**Files:**
- Create: `internal/ccfake/server.go`, `internal/ccfake/server_test.go`

**Interfaces:**
- Produces (used by every provider, protocol and e2e test):
  - `func New() *Server` (an `httptest.Server`; `Server.URL`, `Server.Close()`)
  - `type TypeConfig struct { TypeName, Identifier, IDPrefix string; ReadOnly map[string]string; Defaults map[string]any; CreateOnly, WriteOnly []string; OnRead func(props map[string]any) }`
    — `ReadOnly` values may contain `{id}`; `Defaults` are properties AWS picks when the desired state omits them;
    `OnRead` lets a test mutate what `GetResource` returns (e.g. add nested keys)
  - `func (s *Server) Register(tc TypeConfig)`
  - `type Fault struct { Action string; Nth, Status int; Code, Message string }`; `func (s *Server) Inject(f Fault)`
  - `func (s *Server) FailNext(operation, handlerCode, message string, keepResource bool)` — the next async request of
    `operation` (`CREATE`, `UPDATE`, `DELETE`) ends `FAILED`; for `CREATE` with `keepResource` the resource still exists
  - `func (s *Server) HideFromGet(identifier string, times int)`
  - `func (s *Server) Calls(action string) int`, `func (s *Server) Tokens(action string) []string`, `func (s *Server) AccessKeys() []string`
  - `func (s *Server) Resource(region, typeName, identifier string) (map[string]any, bool)`
  - `func (s *Server) Put(region, typeName, identifier string, props map[string]any)` (behind infrata's back)
  - `func (s *Server) Resources(region, typeName string) map[string]map[string]any`
  - `func (s *Server) LastPatch() []map[string]any`
  - Fields: `PollsToComplete int` (default 1), `PageSize int` (0 = one page)
  - `const AssumedAccessKey = "ASIAFAKEASSUMED"`
  - A create whose desired state already sets the identifier property (a role or bucket name) keeps it; a second
    create of the same identifier ends `FAILED` with `AlreadyExists`. Error responses carry `X-Amzn-Requestid: fake-request-err`.

Wire format (verified, see the Verification log): `POST /`, `X-Amz-Target: CloudApiService.<Operation>`,
`Content-Type: application/x-amz-json-1.0`, JSON bodies with the model's member names, timestamps as epoch seconds,
errors as HTTP status + `X-Amzn-ErrorType` header + `{"__type": code, "message": text}`. STS `AssumeRole` arrives
form-encoded (`Action=AssumeRole`) and answers XML. Resources are kept per region, taken from the SigV4 credential scope.

- [ ] **Step 1: Write the oracle tests**

`internal/ccfake/server_test.go`:

```go
package ccfake

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// client is the real SDK pointed at the fake: if the SDK decodes what the fake says, the wire format is right.
func client(s *Server, region string) *cloudcontrol.Client {
	return cloudcontrol.New(cloudcontrol.Options{
		Region: region, BaseEndpoint: aws.String(s.URL),
		Credentials: credentials.NewStaticCredentialsProvider("AKIDORACLE", "secret", ""),
	})
}

func vpcType() TypeConfig {
	return TypeConfig{
		TypeName: "AWS::EC2::VPC", Identifier: "VpcId", IDPrefix: "vpc-",
		ReadOnly:   map[string]string{"DefaultSecurityGroup": "sg-for-{id}"},
		Defaults:   map[string]any{"EnableDnsSupport": true, "InstanceTenancy": "default"},
		CreateOnly: []string{"CidrBlock"},
		WriteOnly:  []string{"Ipv4NetmaskLength"},
	}
}

func await(t *testing.T, c *cloudcontrol.Client, ev *types.ProgressEvent) *types.ProgressEvent {
	t.Helper()
	for i := 0; i < 20 && (ev.OperationStatus == types.OperationStatusPending || ev.OperationStatus == types.OperationStatusInProgress); i++ {
		out, err := c.GetResourceRequestStatus(context.Background(), &cloudcontrol.GetResourceRequestStatusInput{RequestToken: ev.RequestToken})
		if err != nil {
			t.Fatal(err)
		}
		ev = out.ProgressEvent
	}
	return ev
}

func TestTheSDKDrivesAFullLifecycle(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "eu-west-1")
	ctx := context.Background()

	created, err := c.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), ClientToken: aws.String("tok-1"),
		DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16","Ipv4NetmaskLength":16,"Tags":[{"Key":"team","Value":"a"}]}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.ProgressEvent.OperationStatus != types.OperationStatusInProgress || created.ProgressEvent.EventTime == nil {
		t.Fatalf("create event = %+v", created.ProgressEvent)
	}
	ev := await(t, c, created.ProgressEvent)
	id := aws.ToString(ev.Identifier)
	if ev.OperationStatus != types.OperationStatusSuccess || !strings.HasPrefix(id, "vpc-") {
		t.Fatalf("settled event = %+v", ev)
	}

	got, err := c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	if err != nil {
		t.Fatal(err)
	}
	var props map[string]any
	if err := json.Unmarshal([]byte(aws.ToString(got.ResourceDescription.Properties)), &props); err != nil {
		t.Fatal(err)
	}
	if props["VpcId"] != id || props["EnableDnsSupport"] != true || props["DefaultSecurityGroup"] != "sg-for-"+id {
		t.Errorf("properties = %v: want the identifier, a provider-chosen default and a read-only value", props)
	}
	if _, leaked := props["Ipv4NetmaskLength"]; leaked {
		t.Error("a write-only property was returned")
	}

	upd, err := c.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id), ClientToken: aws.String("tok-2"),
		PatchDocument: aws.String(`[{"op":"replace","path":"/EnableDnsSupport","value":false}]`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if ev := await(t, c, upd.ProgressEvent); ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("update = %+v", ev)
	}
	if res, _ := s.Resource("eu-west-1", "AWS::EC2::VPC", id); res["EnableDnsSupport"] != false {
		t.Errorf("after patch: %v", res)
	}

	_, err = c.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id),
		PatchDocument: aws.String(`[{"op":"replace","path":"/CidrBlock","value":"10.1.0.0/16"}]`),
	})
	var notUpdatable *types.NotUpdatableException
	if !errors.As(err, &notUpdatable) {
		t.Errorf("patching a create-only property = %v, want NotUpdatableException", err)
	}

	del, err := c.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	if err != nil {
		t.Fatal(err)
	}
	if ev := await(t, c, del.ProgressEvent); ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("delete = %+v", ev)
	}
	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: aws.String(id)})
	var nf *types.ResourceNotFoundException
	if !errors.As(err, &nf) {
		t.Fatalf("get after delete = %v", err)
	}
	if s.Calls("CreateResource") != 1 || len(s.Tokens("CreateResource")) != 1 {
		t.Errorf("calls/tokens not recorded")
	}
}

func TestAClientTokenMakesACreateIdempotent(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "us-east-1")
	in := &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), ClientToken: aws.String("same"), DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16"}`)}
	a, err := c.CreateResource(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	b, err := c.CreateResource(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	if aws.ToString(a.ProgressEvent.RequestToken) != aws.ToString(b.ProgressEvent.RequestToken) || len(s.Resources("us-east-1", "AWS::EC2::VPC")) != 1 {
		t.Fatal("the same client token made a second resource")
	}
}

// TestRegionsArePartitionedAndListPages: three in one region, one in another, page size one.
func TestRegionsArePartitionedAndListPages(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	s.PageSize = 1
	for _, id := range []string{"vpc-1", "vpc-2", "vpc-3"} {
		s.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
	}
	s.Put("eu-west-1", "AWS::EC2::VPC", "vpc-9", map[string]any{"VpcId": "vpc-9"})
	var ids []string
	pg := cloudcontrol.NewListResourcesPaginator(client(s, "us-east-1"), &cloudcontrol.ListResourcesInput{TypeName: aws.String("AWS::EC2::VPC")})
	for pg.HasMorePages() {
		page, err := pg.NextPage(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range page.ResourceDescriptions {
			ids = append(ids, aws.ToString(d.Identifier))
		}
	}
	if strings.Join(ids, ",") != "vpc-1,vpc-2,vpc-3" {
		t.Fatalf("listed %v", ids)
	}
}

func TestFaultsAndFailedRequests(t *testing.T) {
	s := New()
	defer s.Close()
	s.Register(vpcType())
	c := client(s, "us-east-1")
	ctx := context.Background()

	s.Inject(Fault{Action: "CreateResource", Nth: 1, Status: 429, Code: "ThrottlingException", Message: "slow down"})
	_, err := cloudcontrol.New(cloudcontrol.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(s.URL), RetryMaxAttempts: 1,
		Credentials: credentials.NewStaticCredentialsProvider("AKID", "secret", ""),
	}).CreateResource(ctx, &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`)})
	var throttled *types.ThrottlingException
	if !errors.As(err, &throttled) {
		t.Fatalf("injected throttle = %v", err)
	}

	s.FailNext("CREATE", "NotStabilized", "not stable yet", true)
	out, err := c.CreateResource(ctx, &cloudcontrol.CreateResourceInput{TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{"CidrBlock":"10.0.0.0/16"}`)})
	if err != nil {
		t.Fatal(err)
	}
	ev := await(t, c, out.ProgressEvent)
	if ev.OperationStatus != types.OperationStatusFailed || ev.ErrorCode != types.HandlerErrorCodeNotStabilized || aws.ToString(ev.Identifier) == "" {
		t.Fatalf("failed create = %+v, want FAILED NotStabilized with an identifier", ev)
	}
	if _, exists := s.Resource("us-east-1", "AWS::EC2::VPC", aws.ToString(ev.Identifier)); !exists {
		t.Error("keepResource: the resource should still exist")
	}

	s.HideFromGet(aws.ToString(ev.Identifier), 1)
	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::EC2::VPC"), Identifier: ev.Identifier})
	var nf *types.ResourceNotFoundException
	if !errors.As(err, &nf) {
		t.Errorf("hidden get = %v", err)
	}

	_, err = c.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String("AWS::Nope::Thing"), Identifier: aws.String("x")})
	var tnf *types.TypeNotFoundException
	if !errors.As(err, &tnf) {
		t.Errorf("unregistered type = %v", err)
	}
}

func TestAssumeRoleAnswersAndSignsAreRecorded(t *testing.T) {
	s := New()
	defer s.Close()
	out, err := sts.New(sts.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(s.URL),
		Credentials: credentials.NewStaticCredentialsProvider("AKIDSTATIC", "secret", ""),
	}).AssumeRole(context.Background(), &sts.AssumeRoleInput{RoleArn: aws.String("arn:aws:iam::123456789012:role/x"), RoleSessionName: aws.String("t")})
	if err != nil {
		t.Fatal(err)
	}
	if aws.ToString(out.Credentials.AccessKeyId) != AssumedAccessKey || !out.Credentials.Expiration.After(time.Now()) {
		t.Fatalf("credentials = %+v", out.Credentials)
	}
	if keys := s.AccessKeys(); len(keys) != 1 || keys[0] != "AKIDSTATIC" {
		t.Errorf("access keys = %v", keys)
	}
}
```

- [ ] **Step 2: Run them to see them fail**

```bash
export GOPRIVATE='github.com/infrata/*'
go get github.com/aws/aws-sdk-go-v2/service/cloudcontrol@v1.38.0
go test -count=1 ./internal/ccfake/
```

Expected: FAIL — `undefined: New`.

- [ ] **Step 3: Implement the fake**

`internal/ccfake/server.go`:

```go
// Package ccfake is an in-process stand-in for AWS Cloud Control API (and STS AssumeRole). Tests drive the real SDK
// against it, so the SDK is the judge of the wire format.
package ccfake

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// AssumedAccessKey is what AssumeRole hands out, so a test can tell assumed-role calls from static-key calls.
const AssumedAccessKey = "ASIAFAKEASSUMED"

// TypeConfig describes how the fake behaves for one resource type.
type TypeConfig struct {
	TypeName   string            // AWS::EC2::VPC
	Identifier string            // the property the fake assigns on create, e.g. VpcId
	IDPrefix   string            // e.g. "vpc-"
	ReadOnly   map[string]string // properties added on create; "{id}" is replaced by the identifier
	Defaults   map[string]any    // properties AWS picks when the desired state omits them
	CreateOnly []string          // patching these is refused with NotUpdatableException
	WriteOnly  []string          // accepted, never returned
	OnRead     func(props map[string]any)
}

// Fault makes the Nth call of Action fail with an HTTP error.
type Fault struct {
	Action  string
	Nth     int
	Status  int
	Code    string
	Message string
}

type failure struct {
	code, message string
	keep          bool
}

type request struct {
	token, operation, typeName, identifier, region string
	polls                                         int
	failCode, failMessage                         string
}

// Server is the fake.
type Server struct {
	*httptest.Server
	PollsToComplete int
	PageSize        int

	mu        sync.Mutex
	next      int
	types     map[string]TypeConfig
	resources map[string]map[string]any // key: region|type|identifier
	writeOnly map[string]map[string]any
	requests  map[string]*request
	byToken   map[string]string // client token -> request token
	faults    []Fault
	failNext  map[string]failure
	hidden    map[string]int
	calls     map[string]int
	tokens    map[string][]string
	keys      []string
	lastPatch []map[string]any
}

// New starts a fake.
func New() *Server {
	s := &Server{
		PollsToComplete: 1, types: map[string]TypeConfig{}, resources: map[string]map[string]any{},
		writeOnly: map[string]map[string]any{}, requests: map[string]*request{}, byToken: map[string]string{},
		failNext: map[string]failure{}, hidden: map[string]int{}, calls: map[string]int{}, tokens: map[string][]string{},
	}
	s.Server = httptest.NewServer(http.HandlerFunc(s.serve))
	return s
}

var scope = regexp.MustCompile(`Credential=([^/]+)/[^/]+/([^/]+)/`)

func key(region, typeName, id string) string { return region + "|" + typeName + "|" + id }

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	accessKey, region := "", ""
	if m := scope.FindStringSubmatch(r.Header.Get("Authorization")); m != nil {
		accessKey, region = m[1], m[2]
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys = append(s.keys, accessKey)

	target := r.Header.Get("X-Amz-Target")
	if target == "" {
		_ = r.ParseForm()
		s.calls[r.Form.Get("Action")]++
		s.assumeRole(w)
		return
	}
	action := strings.TrimPrefix(target, "CloudApiService.")
	s.calls[action]++
	for i, f := range s.faults {
		if f.Action == action && f.Nth == s.calls[action] {
			s.faults = append(s.faults[:i], s.faults[i+1:]...)
			writeError(w, f.Status, f.Code, f.Message)
			return
		}
	}
	body, _ := io.ReadAll(r.Body)
	var in map[string]any
	_ = json.Unmarshal(body, &in)
	str := func(k string) string { v, _ := in[k].(string); return v }
	if tok := str("ClientToken"); tok != "" {
		s.tokens[action] = append(s.tokens[action], tok)
	}

	switch action {
	case "CreateResource":
		s.create(w, region, str("TypeName"), str("DesiredState"), str("ClientToken"))
	case "GetResource":
		s.get(w, region, str("TypeName"), str("Identifier"))
	case "UpdateResource":
		s.update(w, region, str("TypeName"), str("Identifier"), str("PatchDocument"), str("ClientToken"))
	case "DeleteResource":
		s.delete(w, region, str("TypeName"), str("Identifier"), str("ClientToken"))
	case "ListResources":
		s.list(w, region, str("TypeName"), str("NextToken"))
	case "GetResourceRequestStatus":
		s.status(w, str("RequestToken"))
	default:
		writeError(w, 400, "UnsupportedActionException", "ccfake does not implement "+action)
	}
}

func (s *Server) typeConfig(w http.ResponseWriter, typeName string) (TypeConfig, bool) {
	tc, ok := s.types[typeName]
	if !ok {
		writeError(w, 404, "TypeNotFoundException", "The type "+typeName+" cannot be found")
	}
	return tc, ok
}

func (s *Server) newRequest(region, operation, typeName, identifier, clientToken string) (*request, bool) {
	if clientToken != "" {
		if tok, seen := s.byToken[operation+"|"+clientToken]; seen {
			return s.requests[tok], true
		}
	}
	s.next++
	req := &request{token: fmt.Sprintf("req-%d", s.next), operation: operation, typeName: typeName, identifier: identifier, region: region}
	if f, ok := s.failNext[operation]; ok {
		delete(s.failNext, operation)
		req.failCode, req.failMessage = f.code, f.message
	}
	s.requests[req.token] = req
	if clientToken != "" {
		s.byToken[operation+"|"+clientToken] = req.token
	}
	return req, false
}

func (s *Server) create(w http.ResponseWriter, region, typeName, desired, clientToken string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	keep := s.failNext["CREATE"].keep
	req, repeat := s.newRequest(region, "CREATE", typeName, "", clientToken)
	if !repeat && (req.failCode == "" || keep) {
		var props map[string]any
		if err := json.Unmarshal([]byte(desired), &props); err != nil {
			writeError(w, 400, "InvalidRequestException", "DesiredState is not a JSON object")
			return
		}
		// A name the caller chose (a bucket or role name) is the identifier, as in AWS; otherwise the fake makes one.
		id, named := props[tc.Identifier].(string)
		if !named || id == "" {
			s.next++
			id = fmt.Sprintf("%s%017x", tc.IDPrefix, s.next)
		}
		if _, taken := s.resources[key(region, typeName, id)]; taken {
			req.failCode, req.failMessage = "AlreadyExists", id+" already exists"
			writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
			return
		}
		props[tc.Identifier] = id
		for p, v := range tc.Defaults {
			if _, set := props[p]; !set {
				props[p] = v
			}
		}
		for p, tmpl := range tc.ReadOnly {
			props[p] = strings.ReplaceAll(tmpl, "{id}", id)
		}
		wo := map[string]any{}
		for _, p := range tc.WriteOnly {
			if v, set := props[p]; set {
				wo[p] = v
				delete(props, p)
			}
		}
		k := key(region, typeName, id)
		s.resources[k], s.writeOnly[k] = props, wo
		req.identifier = id
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) get(w http.ResponseWriter, region, typeName, id string) {
	if _, ok := s.typeConfig(w, typeName); !ok {
		return
	}
	props, ok := s.resources[key(region, typeName, id)]
	if s.hidden[id] > 0 {
		s.hidden[id]--
		ok = false
	}
	if !ok {
		writeError(w, 404, "ResourceNotFoundException", typeName+" Handler returned status FAILED: "+id+" does not exist (HandlerErrorCode: NotFound)")
		return
	}
	out := cloneMap(props)
	if tc := s.types[typeName]; tc.OnRead != nil {
		tc.OnRead(out)
	}
	raw, _ := json.Marshal(out)
	writeJSON(w, map[string]any{"TypeName": typeName, "ResourceDescription": map[string]any{"Identifier": id, "Properties": string(raw)}})
}

func (s *Server) update(w http.ResponseWriter, region, typeName, id, patch, clientToken string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	props, exists := s.resources[key(region, typeName, id)]
	if !exists {
		writeError(w, 404, "ResourceNotFoundException", id+" does not exist")
		return
	}
	var ops []map[string]any
	if err := json.Unmarshal([]byte(patch), &ops); err != nil {
		writeError(w, 400, "InvalidRequestException", "PatchDocument is not a JSON Patch array")
		return
	}
	for _, op := range ops {
		path, _ := op["path"].(string)
		name := strings.TrimPrefix(path, "/")
		if strings.Contains(name, "/") {
			writeError(w, 400, "InvalidRequestException", "ccfake supports top-level patch paths only: "+path)
			return
		}
		for _, co := range tc.CreateOnly {
			if co == name {
				writeError(w, 400, "NotUpdatableException", "Invalid patch update: createOnlyProperties ["+path+"] cannot be updated")
				return
			}
		}
	}
	req, repeat := s.newRequest(region, "UPDATE", typeName, id, clientToken)
	if !repeat && req.failCode == "" {
		s.lastPatch = ops
		for _, op := range ops {
			name := strings.TrimPrefix(op["path"].(string), "/")
			switch op["op"] {
			case "add", "replace":
				if isWriteOnly(tc, name) {
					s.writeOnly[key(region, typeName, id)][name] = op["value"]
				} else {
					props[name] = op["value"]
				}
			case "remove":
				delete(props, name)
			}
		}
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) delete(w http.ResponseWriter, region, typeName, id, clientToken string) {
	if _, ok := s.typeConfig(w, typeName); !ok {
		return
	}
	req, repeat := s.newRequest(region, "DELETE", typeName, id, clientToken)
	k := key(region, typeName, id)
	if _, exists := s.resources[k]; !exists && req.failCode == "" {
		req.failCode, req.failMessage = "NotFound", id+" does not exist"
	}
	if !repeat && req.failCode == "" {
		delete(s.resources, k)
		delete(s.writeOnly, k)
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) list(w http.ResponseWriter, region, typeName, token string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	var ids []string
	prefix := region + "|" + typeName + "|"
	for k := range s.resources {
		if id, ok := strings.CutPrefix(k, prefix); ok {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	start, _ := strconv.Atoi(token)
	end, next := len(ids), ""
	if s.PageSize > 0 && start+s.PageSize < len(ids) {
		end, next = start+s.PageSize, strconv.Itoa(start+s.PageSize)
	}
	var descs []map[string]any
	for _, id := range ids[min(start, len(ids)):end] {
		raw, _ := json.Marshal(map[string]any{tc.Identifier: id}) // like the real API: often only the identifier
		descs = append(descs, map[string]any{"Identifier": id, "Properties": string(raw)})
	}
	out := map[string]any{"TypeName": typeName, "ResourceDescriptions": descs}
	if next != "" {
		out["NextToken"] = next
	}
	writeJSON(w, out)
}

func (s *Server) status(w http.ResponseWriter, token string) {
	req, ok := s.requests[token]
	if !ok {
		writeError(w, 404, "RequestTokenNotFoundException", "no request "+token)
		return
	}
	req.polls++
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) event(req *request) map[string]any {
	ev := map[string]any{
		"TypeName": req.typeName, "RequestToken": req.token, "Operation": req.operation,
		"EventTime": float64(time.Now().UnixMilli()) / 1000, "OperationStatus": "IN_PROGRESS",
	}
	if req.identifier != "" {
		ev["Identifier"] = req.identifier
	}
	if req.polls >= s.PollsToComplete {
		ev["OperationStatus"] = "SUCCESS"
		if req.failCode != "" {
			ev["OperationStatus"], ev["ErrorCode"], ev["StatusMessage"] = "FAILED", req.failCode, req.failMessage
		}
	}
	return ev
}

func (s *Server) assumeRole(w http.ResponseWriter) {
	type creds struct {
		AccessKeyID     string `xml:"AccessKeyId"`
		SecretAccessKey string `xml:"SecretAccessKey"`
		SessionToken    string `xml:"SessionToken"`
		Expiration      string `xml:"Expiration"`
	}
	var out struct {
		XMLName xml.Name `xml:"AssumeRoleResponse"`
		Result  struct {
			Credentials creds `xml:"Credentials"`
		} `xml:"AssumeRoleResult"`
	}
	out.Result.Credentials = creds{AssumedAccessKey, "assumed-secret", "assumed-token", "2099-01-01T00:00:00Z"}
	w.Header().Set("Content-Type", "text/xml")
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(out)
}

func writeJSON(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/x-amz-json-1.0")
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/x-amz-json-1.0")
	w.Header().Set("X-Amzn-ErrorType", code)
	w.Header().Set("X-Amzn-Requestid", "fake-request-err") // the SDK reads it into ServiceRequestID
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"__type": code, "message": message})
}

func isWriteOnly(tc TypeConfig, name string) bool {
	for _, p := range tc.WriteOnly {
		if p == name {
			return true
		}
	}
	return false
}

func cloneMap(m map[string]any) map[string]any {
	raw, _ := json.Marshal(m)
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	return out
}

// ---- test controls ----

// Register makes the fake serve a type.
func (s *Server) Register(tc TypeConfig) { s.mu.Lock(); defer s.mu.Unlock(); s.types[tc.TypeName] = tc }

// Inject queues an HTTP fault.
func (s *Server) Inject(f Fault) { s.mu.Lock(); defer s.mu.Unlock(); s.faults = append(s.faults, f) }

// FailNext makes the next async request of operation end FAILED with a handler error code.
func (s *Server) FailNext(operation, handlerCode, message string, keepResource bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failNext[operation] = failure{code: handlerCode, message: message, keep: keepResource}
}

// HideFromGet makes the next `times` GetResource calls for identifier answer ResourceNotFoundException.
func (s *Server) HideFromGet(identifier string, times int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hidden[identifier] = times
}

// Calls counts requests for an action.
func (s *Server) Calls(action string) int { s.mu.Lock(); defer s.mu.Unlock(); return s.calls[action] }

// Tokens lists the client tokens an action was called with.
func (s *Server) Tokens(action string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.tokens[action]...)
}

// AccessKeys lists the access key ID each request was signed with.
func (s *Server) AccessKeys() []string { s.mu.Lock(); defer s.mu.Unlock(); return append([]string(nil), s.keys...) }

// Resource snapshots one resource's stored properties (write-only ones excluded).
func (s *Server) Resource(region, typeName, identifier string) (map[string]any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.resources[key(region, typeName, identifier)]
	if !ok {
		return nil, false
	}
	return cloneMap(p), true
}

// Put stores a resource behind infrata's back: pre-existing infrastructure, or drift.
func (s *Server) Put(region, typeName, identifier string, props map[string]any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := key(region, typeName, identifier)
	s.resources[k] = cloneMap(props)
	if s.writeOnly[k] == nil {
		s.writeOnly[k] = map[string]any{}
	}
}

// Resources snapshots every resource of a type in a region.
func (s *Server) Resources(region, typeName string) map[string]map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := map[string]map[string]any{}
	prefix := region + "|" + typeName + "|"
	for k, p := range s.resources {
		if id, ok := strings.CutPrefix(k, prefix); ok {
			out[id] = cloneMap(p)
		}
	}
	return out
}

// LastPatch is the most recent patch document applied.
func (s *Server) LastPatch() []map[string]any { s.mu.Lock(); defer s.mu.Unlock(); return s.lastPatch }
```

- [ ] **Step 4: Run the oracle tests**

Run: `go test -count=1 ./internal/ccfake/ && go vet ./internal/ccfake/`
Expected: PASS. If the SDK fails to decode a response, the fake's JSON member names or error shape are wrong: fix them
against `service/cloudcontrol@v1.38.0` (the SDK is authoritative) and note the correction in the Verification log.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: complete requests immediately (`PollsToComplete` ignored): the lifecycle test's `IN_PROGRESS` assertion
fails; ignore the region when listing (the partition test lists 4); drop client-token reuse (the idempotency test sees
two resources); return write-only properties from `GetResource` (the lifecycle test fails on `Ipv4NetmaskLength`).

```bash
git add go.mod go.sum internal/ccfake/server.go internal/ccfake/server_test.go
git commit -m "Add a fake Cloud Control endpoint

Serves create, get, update, delete, list and request status like the real
API, including async requests, client tokens, create-only refusals and
injected failures, and answers STS AssumeRole. The real SDK decodes it.
Checked by breaking polling, region separation, tokens and write-only
handling." -- go.mod go.sum internal/ccfake/server.go internal/ccfake/server_test.go
```

---
### Task 8: `internal/ccprov` foundations — IDs, errors, patience, clients, awaiting a request

**Files:**
- Create: `internal/ccprov/ids.go`, `internal/ccprov/errors.go`, `internal/ccprov/patience.go`,
  `internal/ccprov/clients.go`, `internal/ccprov/await.go`
- Test: `internal/ccprov/catalog_test.go` (the test catalog every ccprov test uses), `internal/ccprov/helpers_test.go`,
  `internal/ccprov/ids_test.go`, `internal/ccprov/errors_test.go`, `internal/ccprov/await_test.go`

**Interfaces:**
- Consumes: `catalog.Catalog`, `catalog.Type`, `catalog.Attribute`, `catalog.Shape`, `catalog.GlobalScope`,
  `(*catalog.Type).Global`, `Definition` (Task 2); `ccfake.New`, `Inject`, `Fault`, `PollsToComplete`, `Calls` (Task 7).
- Produces (used by Tasks 9 to 12):
  - `const GlobalRegion = "us-east-1"`
  - `func FormatID(t *catalog.Type, region, identifier string) string`
  - `func ParseID(t *catalog.Type, providerID string) (region, identifier string, err error)`
  - `type HandlerError struct { Operation, Code, Message, Identifier, Token string }`
  - `func classify(err error) provider.Retryability`
  - `func errorCode(err error) string`
  - `func failure(instance, action string, t *catalog.Type, where string, err error) error`
  - `func isNotFound(err error) bool`
  - `type patience struct { attempts int; base, max time.Duration; sleep func(context.Context, time.Duration) error }`,
    `var notFoundPatience`, `var once`, `func (pt patience) wait(ctx, try func() (bool, error)) (bool, error)`
  - `type clients struct`; `func newClients(cfg aws.Config) *clients`; `func (c *clients) get(region string) *cloudcontrol.Client`
  - `type statusAPI interface`; `type pacing struct { first, max time.Duration; sleep …; now func() time.Time }`;
    `var defaultPacing`; `func (pc pacing) await(ctx, api statusAPI, ev *types.ProgressEvent, timeout time.Duration) (*types.ProgressEvent, error)`
  - `func timeoutFor(t *catalog.Type, handler string) time.Duration`
  - test-only: `testCatalog() *catalog.Catalog`, `testConfig(endpoint string, attempts int) aws.Config`,
    `instantPatience`, `instantPacing`

P2 moves these pieces into `ccprov`. Their reasoning is carried over from `first-slice` Task 4 and Task 5 (classification
order, patience) and changed only where Cloud Control differs: every mutation carries a `ClientToken`, so the SDK's
retryer stays on, and a mutation is a request that must be awaited.

- [ ] **Step 1: The test catalog**

Hand-written, not the embedded catalog: each type exercises one thing later tasks need, and a regenerated bundle cannot
change what a unit test means. Spellings follow J4/J7 exactly as the generator would write them.

`internal/ccprov/catalog_test.go`:

```go
package ccprov

import (
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
)

func scalar() *catalog.Shape { return &catalog.Shape{Kind: catalog.ShapeScalar} }
func opaque() *catalog.Shape { return &catalog.Shape{Kind: catalog.ShapeOpaque} }
func object(props map[string]*catalog.Shape) *catalog.Shape {
	return &catalog.Shape{Kind: catalog.ShapeObject, Props: props}
}
func array(item *catalog.Shape, unordered bool) *catalog.Shape {
	return &catalog.Shape{Kind: catalog.ShapeArray, Item: item, Unordered: unordered}
}

// testCatalog is every shape of type the provider handles:
//
//	aws.vpc              regional, provider-chosen defaults, a write-only property, tags as a map
//	aws.securitygroup    an unordered list of objects (nested spelling, AWS-added keys, order)
//	aws.bucket           nested objects inside ordered lists inside objects
//	aws.role             global, an opaque policy document, an unordered list holding opaque values
//	aws.dbinstance       sensitive and write-only, a long create timeout, a computed object
//	aws.test.regioned    its own Region property (so the plugin's is aws_region); no update handler, no list handler
//	aws.test.child       a composite identifier; listing needs a parent model
func testCatalog() *catalog.Catalog {
	return &catalog.Catalog{
		Bundle:          "test",
		DiscoverDefault: []string{"aws.vpc", "aws.role"},
		Types: []*catalog.Type{
			{
				Name: "aws.vpc", CFN: "AWS::EC2::VPC", RegionAttr: "region", Identifier: []string{"VpcId"},
				WriteOnly: []string{"Ipv4NetmaskLength"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "CidrBlock", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"cidr", "cidr_block"}},
					{Name: "EnableDnsSupport", Kind: "boolean", Optional: true, Computed: true, Aliases: []string{"enable_dns_support"}},
					{Name: "InstanceTenancy", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"instance_tenancy"}},
					{Name: "Ipv4NetmaskLength", Kind: "integer", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"ipv4_netmask_length"}},
					{Name: "VpcId", Kind: "string", Computed: true, Aliases: []string{"vpc_id"}},
					{Name: "DefaultSecurityGroup", Kind: "string", Computed: true, Aliases: []string{"default_security_group"}},
					{Name: "CidrBlockAssociations", Kind: "list", Computed: true, Aliases: []string{"cidr_block_associations"}, Shape: array(scalar(), true)},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.securitygroup", CFN: "AWS::EC2::SecurityGroup", RegionAttr: "region", Identifier: []string{"GroupId"},
				HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "GroupDescription", Kind: "string", Required: true, ForceNew: true, Aliases: []string{"description", "group_description"}},
					{Name: "GroupName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "group_name"}},
					{Name: "VpcId", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"vpc_id"}},
					{Name: "GroupId", Kind: "string", Computed: true, Aliases: []string{"group_id"}},
					{Name: "SecurityGroupIngress", Kind: "list", Optional: true, Computed: true, Aliases: []string{"ingress", "security_group_ingress"},
						Shape: array(object(map[string]*catalog.Shape{
							"IpProtocol": scalar(), "FromPort": scalar(), "ToPort": scalar(), "CidrIp": scalar(), "Description": scalar(),
						}), true)},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.bucket", CFN: "AWS::S3::Bucket", RegionAttr: "region", Identifier: []string{"BucketName"},
				HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "BucketName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "bucket_name"}},
					{Name: "Arn", Kind: "string", Computed: true},
					{Name: "VersioningConfiguration", Kind: "map", Optional: true, Computed: true, Aliases: []string{"versioning_configuration"},
						Shape: object(map[string]*catalog.Shape{"Status": scalar()})},
					{Name: "LifecycleConfiguration", Kind: "map", Optional: true, Computed: true, Aliases: []string{"lifecycle_configuration"},
						Shape: object(map[string]*catalog.Shape{
							"Rules": array(object(map[string]*catalog.Shape{
								"Id": scalar(), "Status": scalar(), "ExpirationInDays": scalar(),
								"Transitions": array(object(map[string]*catalog.Shape{"StorageClass": scalar(), "TransitionInDays": scalar()}), false),
							}), false),
						})},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.role", CFN: "AWS::IAM::Role", Identifier: []string{"RoleName"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "RoleName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "role_name"}},
					{Name: "Arn", Kind: "string", Computed: true},
					{Name: "AssumeRolePolicyDocument", Kind: "map", Required: true, Aliases: []string{"assume_role_policy", "assume_role_policy_document"}, Shape: opaque()},
					{Name: "Policies", Kind: "list", Optional: true, Computed: true,
						Shape: array(object(map[string]*catalog.Shape{"PolicyName": scalar(), "PolicyDocument": opaque()}), true)},
					{Name: "MaxSessionDuration", Kind: "integer", Optional: true, Computed: true, Aliases: []string{"max_session_duration"}},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.dbinstance", CFN: "AWS::RDS::DBInstance", RegionAttr: "region", Identifier: []string{"DBInstanceIdentifier"},
				WriteOnly: []string{"MasterUserPassword"}, HasUpdate: true, HasList: true, Timeouts: map[string]int{"create": 2160},
				Attributes: []*catalog.Attribute{
					{Name: "DBInstanceIdentifier", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"identifier", "db_instance_identifier"}},
					{Name: "DBInstanceClass", Kind: "string", Optional: true, Computed: true, Aliases: []string{"instance_class", "db_instance_class"}},
					{Name: "MasterUserPassword", Kind: "string", Optional: true, Computed: true, Sensitive: true, Aliases: []string{"password", "master_user_password"}},
					{Name: "Endpoint", Kind: "map", Computed: true, Shape: object(map[string]*catalog.Shape{"Address": scalar(), "Port": scalar()})},
				},
			},
			{
				Name: "aws.test.regioned", CFN: "AWS::Test::Regioned", RegionAttr: "aws_region", Identifier: []string{"Id"},
				Attributes: []*catalog.Attribute{
					{Name: "Id", Kind: "string", Computed: true},
					{Name: "Name", Kind: "string", Required: true, ForceNew: true},
					{Name: "Region", Kind: "string", Optional: true, Computed: true, ForceNew: true},
				},
			},
			{
				Name: "aws.test.child", CFN: "AWS::Test::Child", RegionAttr: "region", Identifier: []string{"ParentId", "ChildId"},
				HasUpdate: true, HasList: true, ListNeedsModel: true,
				Attributes: []*catalog.Attribute{
					{Name: "ParentId", Kind: "string", Required: true, ForceNew: true, Aliases: []string{"parent_id"}},
					{Name: "ChildId", Kind: "string", Computed: true, Aliases: []string{"child_id"}},
					{Name: "Setting", Kind: "string", Optional: true, Computed: true},
				},
			},
		},
	}
}

func mustType(t *testing.T, name string) *catalog.Type {
	t.Helper()
	typ, ok := testCatalog().Lookup(name)
	if !ok {
		t.Fatalf("test catalog has no %s", name)
	}
	return typ
}

// TestTheTestCatalogIsOneInfrataWouldLoad. A test catalog infrata would refuse would make every test built on it moot.
func TestTheTestCatalogIsOneInfrataWouldLoad(t *testing.T) {
	for _, d := range testCatalog().Definitions() {
		if err := d.Validate(); err != nil {
			t.Errorf("%s: %v", d.Type, err)
		}
	}
}
```

`internal/ccprov/helpers_test.go`:

```go
package ccprov

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// testConfig is an SDK configuration pointed at a fake, with the standard retryer making `attempts` attempts and no
// backoff: the SDK's retry decisions are real, only the waiting is removed.
func testConfig(endpoint string, attempts int) aws.Config {
	return aws.Config{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(endpoint),
		Credentials:  credentials.NewStaticCredentialsProvider("AKIDTEST", "secret", ""),
		Retryer: func() aws.Retryer {
			return retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = attempts
				o.Backoff = retry.BackoffDelayerFunc(func(int, error) (time.Duration, error) { return 0, nil })
			})
		},
	}
}

func noSleep(context.Context, time.Duration) error { return nil }

var (
	instantPatience = patience{attempts: notFoundPatience.attempts, base: time.Millisecond, max: time.Millisecond, sleep: noSleep}
	instantPacing   = pacing{first: time.Millisecond, max: time.Millisecond, sleep: noSleep, now: time.Now}
)
```

- [ ] **Step 2: Write the failing tests**

`internal/ccprov/ids_test.go`:

```go
package ccprov

import (
	"strings"
	"testing"
)

func TestProviderIDsRoundTrip(t *testing.T) {
	cases := []struct {
		typ                    string
		region, identifier, id string
	}{
		{"aws.vpc", "eu-west-1", "vpc-0abc", "eu-west-1/vpc-0abc"},
		{"aws.role", GlobalRegion, "deploy", "global/deploy"},
		// Identifiers may be ARNs, which hold slashes: the split is at the first slash only.
		{"aws.vpc", "us-east-1", "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-1", "us-east-1/arn:aws:ec2:us-east-1:123456789012:vpc/vpc-1"},
		{"aws.test.child", "us-east-1", "parent-1|child-2", "us-east-1/parent-1|child-2"},
	}
	for _, c := range cases {
		typ := mustType(t, c.typ)
		if got := FormatID(typ, c.region, c.identifier); got != c.id {
			t.Errorf("FormatID(%s, %s, %s) = %q, want %q", c.typ, c.region, c.identifier, got, c.id)
		}
		region, identifier, err := ParseID(typ, c.id)
		if err != nil || region != c.region || identifier != c.identifier {
			t.Errorf("ParseID(%s, %q) = %q, %q, %v", c.typ, c.id, region, identifier, err)
		}
	}
}

// TestMalformedIDsAreRefusedBeforeAnyCall. An import typo must say what the ID should look like.
func TestMalformedIDsAreRefusedBeforeAnyCall(t *testing.T) {
	cases := []struct{ typ, id, want string }{
		{"aws.vpc", "vpc-0abc", "us-east-1/<VpcId>"},
		{"aws.vpc", "/vpc-0abc", "us-east-1/<VpcId>"},
		{"aws.vpc", "us-east-1/", "us-east-1/<VpcId>"},
		{"aws.vpc", "global/vpc-0abc", "regional"},
		{"aws.role", "us-east-1/deploy", "global/<RoleName>"},
		{"aws.test.child", "us-east-1/child-2", "<ParentId|ChildId>"},
	}
	for _, c := range cases {
		if _, _, err := ParseID(mustType(t, c.typ), c.id); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("ParseID(%s, %q) = %v, want an error mentioning %q", c.typ, c.id, err, c.want)
		}
	}
}
```

`internal/ccprov/errors_test.go`:

```go
package ccprov

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
	"github.com/infrata/infrata/pkg/provider"
)

func create(t *testing.T, endpoint string, attempts int) error {
	t.Helper()
	_, err := cloudcontrol.NewFromConfig(testConfig(endpoint, attempts)).CreateResource(context.Background(), &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`), ClientToken: aws.String("t"),
	})
	return err
}

// TestClassificationOfRealSDKErrors. Every case goes through the real SDK against the fake, so the error chains are
// the ones production sees. Statuses and codes are the Cloud Control model's (Verification log).
func TestClassificationOfRealSDKErrors(t *testing.T) {
	cases := []struct {
		status int
		code   string
		want   provider.Retryability
	}{
		{429, "ThrottlingException", provider.SafeToRetry},
		{409, "ConcurrentOperationException", provider.SafeToRetry},
		{409, "ResourceConflictException", provider.SafeToRetry},
		{502, "HandlerInternalFailureException", provider.ConditionallyRetryable},
		{502, "ServiceInternalErrorException", provider.ConditionallyRetryable},
		{502, "NetworkFailureException", provider.ConditionallyRetryable},
		{500, "ConcurrentModificationException", provider.ConditionallyRetryable},
		{400, "NotStabilizedException", provider.ConditionallyRetryable},
		{400, "InvalidRequestException", provider.NotSafeToRetry},
		{409, "ClientTokenConflictException", provider.NotSafeToRetry},
		{400, "AlreadyExistsException", provider.NotSafeToRetry},
		{401, "InvalidCredentialsException", provider.NotSafeToRetry},
	}
	for _, c := range cases {
		t.Run(c.code, func(t *testing.T) {
			fake := ccfake.New()
			defer fake.Close()
			fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: 1, Status: c.status, Code: c.code, Message: "injected"})
			err := create(t, fake.URL, 1)
			if err == nil {
				t.Fatal("expected the injected failure")
			}
			if got := classify(err); got != c.want {
				t.Errorf("classify(%v) = %v, want %v", err, got, c.want)
			}
			if errorCode(err) != c.code {
				t.Errorf("errorCode = %q, want %q", errorCode(err), c.code)
			}
		})
	}
}

// TestAnErrorTheSDKGaveUpOnIsStillClassified. The SDK wraps the last attempt's error; classification sees through it.
func TestAnErrorTheSDKGaveUpOnIsStillClassified(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	for n := 1; n <= 3; n++ {
		fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: n, Status: 429, Code: "ThrottlingException"})
	}
	err := create(t, fake.URL, 3)
	if err == nil || fake.Calls("CreateResource") != 3 {
		t.Fatalf("err = %v after %d calls; want failure after 3", err, fake.Calls("CreateResource"))
	}
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

func TestARefusedConnectionIsSafeToRetry(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()
	if got := classify(create(t, "http://"+addr, 1)); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

// TestHandlerFailuresAreClassifiedByTheirCode. A FAILED request was accepted and may have acted.
func TestHandlerFailuresAreClassifiedByTheirCode(t *testing.T) {
	cases := map[string]provider.Retryability{
		"NotStabilized": provider.ConditionallyRetryable, "ServiceInternalError": provider.ConditionallyRetryable,
		"InternalFailure": provider.ConditionallyRetryable, "NetworkFailure": provider.ConditionallyRetryable,
		"ServiceTimeout": provider.ConditionallyRetryable, "Throttling": provider.ConditionallyRetryable,
		"ResourceConflict": provider.ConditionallyRetryable,
		"AlreadyExists": provider.NotSafeToRetry, "InvalidRequest": provider.NotSafeToRetry, "AccessDenied": provider.NotSafeToRetry,
		"NotUpdatable": provider.NotSafeToRetry, "ServiceLimitExceeded": provider.NotSafeToRetry, "Cancelled": provider.NotSafeToRetry,
	}
	for code, want := range cases {
		err := error(&HandlerError{Operation: "CREATE", Code: code, Message: "m", Token: "req-1"})
		if got := classify(err); got != want {
			t.Errorf("classify(%s) = %v, want %v", code, got, want)
		}
		if errorCode(err) != code {
			t.Errorf("errorCode(%s) = %q", code, errorCode(err))
		}
	}
}

func TestCancellationAndUnknownsAreNotSafe(t *testing.T) {
	for _, err := range []error{nil, context.Canceled, errors.New("something new")} {
		if got := classify(err); got != provider.NotSafeToRetry {
			t.Errorf("classify(%v) = %v", err, got)
		}
	}
	if got := classify(context.DeadlineExceeded); got != provider.ConditionallyRetryable {
		t.Errorf("classify(DeadlineExceeded) = %v", got)
	}
}

// TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause. Never anything from the request: it may hold a secret.
func TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	fake.Inject(ccfake.Fault{Action: "CreateResource", Nth: 1, Status: 400, Code: "InvalidRequestException", Message: "CidrBlock is malformed"})
	err := failure("prod", "CreateResource", mustType(t, "aws.vpc"), "us-east-1", create(t, fake.URL, 1))
	for _, want := range []string{`"prod"`, "CreateResource", "aws.vpc", "AWS::EC2::VPC", "us-east-1", "InvalidRequestException", "CidrBlock is malformed", "fake-request-err"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%q lacks %q", err, want)
		}
	}
	if errorCode(err) != "InvalidRequestException" || classify(err) != provider.NotSafeToRetry {
		t.Error("wrapping lost the cause")
	}

	denied := failure("prod", "create", mustType(t, "aws.role"), "global",
		&HandlerError{Operation: "CREATE", Code: "AccessDenied", Message: "not authorized to perform iam:CreateRole", Token: "req-9"})
	for _, want := range []string{"AccessDenied", "iam:CreateRole", "req-9", "IAM permissions"} {
		if !strings.Contains(denied.Error(), want) {
			t.Errorf("%q lacks %q", denied, want)
		}
	}
	exists := failure("prod", "create", mustType(t, "aws.role"), "global", &HandlerError{Code: "AlreadyExists", Message: "deploy exists", Token: "req-3"})
	if !strings.Contains(exists.Error(), "infrata import") {
		t.Errorf("%q does not suggest importing", exists)
	}
}
```

`internal/ccprov/await_test.go`:

```go
package ccprov

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
	"github.com/infrata/infrata/pkg/provider"
)

// scripted answers GetResourceRequestStatus from a list, repeating the last, and records whether each call's context
// had been cancelled.
type scripted struct {
	events  []types.ProgressEvent
	calls   int
	ctxErrs []error
}

func (s *scripted) GetResourceRequestStatus(ctx context.Context, _ *cloudcontrol.GetResourceRequestStatusInput, _ ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error) {
	s.ctxErrs = append(s.ctxErrs, ctx.Err())
	ev := s.events[min(s.calls, len(s.events)-1)]
	s.calls++
	return &cloudcontrol.GetResourceRequestStatusOutput{ProgressEvent: &ev}, nil
}

// clock is a fake time source whose sleeps advance it.
type clock struct {
	now   time.Time
	slept []time.Duration
}

func (c *clock) pacing() pacing {
	return pacing{
		first: time.Second, max: 4 * time.Second,
		now:   func() time.Time { return c.now },
		sleep: func(_ context.Context, d time.Duration) error { c.slept = append(c.slept, d); c.now = c.now.Add(d); return nil },
	}
}

func event(status types.OperationStatus) types.ProgressEvent {
	return types.ProgressEvent{OperationStatus: status, RequestToken: aws.String("req-1"), Operation: types.OperationCreate}
}

func TestAwaitPollsWithGrowingDelaysUntilSuccess(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress), event(types.OperationStatusInProgress), event(types.OperationStatusSuccess)}}
	start := event(types.OperationStatusPending)
	ev, err := c.pacing().await(context.Background(), api, &start, time.Hour)
	if err != nil || ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("await = %+v, %v", ev, err)
	}
	if got := c.slept; len(got) != 3 || got[0] != time.Second || got[1] != 2*time.Second || got[2] != 4*time.Second {
		t.Errorf("slept %v, want [1s 2s 4s]", got)
	}
}

func TestAwaitHonoursRetryAfter(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	start := event(types.OperationStatusInProgress)
	start.RetryAfter = aws.Time(c.now.Add(7 * time.Second))
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusSuccess)}}
	if _, err := c.pacing().await(context.Background(), api, &start, time.Hour); err != nil {
		t.Fatal(err)
	}
	if len(c.slept) != 1 || c.slept[0] != 7*time.Second {
		t.Errorf("slept %v, want [7s]", c.slept)
	}
}

// TestCancellationNeverAbandonsARequestAlreadySent. The host waits for the real answer: an abandoned create is an
// orphan.
func TestCancellationNeverAbandonsARequestAlreadySent(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress), event(types.OperationStatusSuccess)}}
	start := event(types.OperationStatusInProgress)
	if _, err := c.pacing().await(ctx, api, &start, time.Hour); err != nil {
		t.Fatal(err)
	}
	for i, e := range api.ctxErrs {
		if e != nil {
			t.Errorf("status call %d ran with a cancelled context: %v", i+1, e)
		}
	}
}

func TestAFailedRequestIsAHandlerError(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	failed := event(types.OperationStatusFailed)
	failed.ErrorCode, failed.StatusMessage, failed.Identifier = types.HandlerErrorCodeNotStabilized, aws.String("still attaching"), aws.String("vpc-9")
	api := &scripted{events: []types.ProgressEvent{failed}}
	start := event(types.OperationStatusInProgress)
	ev, err := c.pacing().await(context.Background(), api, &start, time.Hour)
	var he *HandlerError
	if !errors.As(err, &he) || he.Code != "NotStabilized" || he.Identifier != "vpc-9" || he.Token != "req-1" || he.Message != "still attaching" {
		t.Fatalf("err = %#v", err)
	}
	if ev == nil || aws.ToString(ev.Identifier) != "vpc-9" {
		t.Errorf("the final event was not returned with the error: %+v", ev)
	}
}

func TestAwaitGivesUpAtTheHandlerTimeoutAndSaysTheRequestMayFinish(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress)}}
	start := event(types.OperationStatusInProgress)
	_, err := c.pacing().await(context.Background(), api, &start, 10*time.Second)
	if err == nil || !strings.Contains(err.Error(), "req-1") || !strings.Contains(err.Error(), "refresh") {
		t.Fatalf("err = %v", err)
	}
	if classify(err) != provider.ConditionallyRetryable {
		t.Errorf("classify = %v, want ConditionallyRetryable", classify(err))
	}
	var total time.Duration
	for _, d := range c.slept {
		total += d
	}
	if total > 10*time.Second {
		t.Errorf("waited %s, beyond the 10s timeout", total)
	}
}

func TestTimeoutsComeFromTheSchemaOrCloudFormationsDefault(t *testing.T) {
	db := mustType(t, "aws.dbinstance")
	if got := timeoutFor(db, "create"); got != 2160*time.Minute {
		t.Errorf("create = %s", got)
	}
	if got := timeoutFor(db, "delete"); got != 120*time.Minute {
		t.Errorf("delete = %s, want the 120 minute default", got)
	}
}

// TestAwaitAgainstTheFake: the real SDK's GetResourceRequestStatus, and the per-region client cache.
func TestAwaitAgainstTheFake(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	fake.Register(ccfake.TypeConfig{TypeName: "AWS::EC2::VPC", Identifier: "VpcId", IDPrefix: "vpc-"})
	fake.PollsToComplete = 3
	cl := newClients(testConfig(fake.URL, 1))
	if cl.get("eu-west-1") != cl.get("eu-west-1") || cl.get("eu-west-1") == cl.get("us-east-1") {
		t.Fatal("clients are not cached per region")
	}
	if got := cl.get("eu-west-1").Options().Region; got != "eu-west-1" {
		t.Fatalf("client region = %q", got)
	}
	out, err := cl.get("eu-west-1").CreateResource(context.Background(), &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`), ClientToken: aws.String("t"),
	})
	if err != nil {
		t.Fatal(err)
	}
	ev, err := instantPacing.await(context.Background(), cl.get("eu-west-1"), out.ProgressEvent, time.Minute)
	if err != nil || ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("await = %+v, %v", ev, err)
	}
	if n := fake.Calls("GetResourceRequestStatus"); n != 3 {
		t.Errorf("status calls = %d, want 3", n)
	}
}
```

- [ ] **Step 3: Run them to see them fail**

Run: `go test -count=1 ./internal/ccprov/`
Expected: FAIL — `undefined: FormatID`, `undefined: classify`, `undefined: pacing` and so on.

- [ ] **Step 4: Implement**

`internal/ccprov/ids.go`:

```go
// Package ccprov is the generic provider: one implementation serving every catalog type through AWS Cloud Control API.
package ccprov

import (
	"fmt"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
)

// GlobalRegion is where Cloud Control is called for a type that has no region (IAM, Route 53, CloudFront).
const GlobalRegion = "us-east-1"

// FormatID is the one provider-ID form (P3): `<region>/<identifier>`, or `global/<identifier>`. Create, Discover and
// Import all use it, because infrata imports by matching `<type>.<provider id>` against what Discover returned.
func FormatID(t *catalog.Type, region, identifier string) string {
	if t.Global() {
		return catalog.GlobalScope + "/" + identifier
	}
	return region + "/" + identifier
}

// ParseID refuses an ID that cannot be this type's before any API call. It splits at the first slash only: an
// identifier may be an ARN.
func ParseID(t *catalog.Type, providerID string) (region, identifier string, err error) {
	shape := "<" + strings.Join(t.Identifier, "|") + ">"
	example := "us-east-1/" + shape
	if t.Global() {
		example = catalog.GlobalScope + "/" + shape
	}
	scope, id, ok := strings.Cut(providerID, "/")
	if !ok || scope == "" || id == "" {
		return "", "", fmt.Errorf("%q is not a %s ID: expected %s", providerID, t.Name, example)
	}
	if parts := strings.Count(id, "|") + 1; parts != len(t.Identifier) {
		return "", "", fmt.Errorf("%q is not a %s ID: expected %s, %d part(s) separated by |", providerID, t.Name, example, len(t.Identifier))
	}
	switch {
	case t.Global() && scope != catalog.GlobalScope:
		return "", "", fmt.Errorf("%q is not a %s ID: %s is not regional, so its ID is %s", providerID, t.Name, t.Name, example)
	case t.Global():
		return GlobalRegion, id, nil
	case scope == catalog.GlobalScope:
		return "", "", fmt.Errorf("%q is not a %s ID: %s is regional, so its ID starts with the region: %s", providerID, t.Name, t.Name, example)
	}
	return scope, id, nil
}
```

`internal/ccprov/errors.go`:

```go
package ccprov

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
)

// HandlerError is a request Cloud Control accepted whose resource handler ended FAILED, or that was cancelled.
type HandlerError struct {
	Operation  string // CREATE, UPDATE or DELETE
	Code       string // a handler error code, e.g. NotStabilized
	Message    string
	Identifier string // set when AWS had assigned one
	Token      string // the request token, for `aws cloudcontrol get-resource-request-status`
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("%s request %s ended %s: %s", e.Operation, e.Token, e.Code, e.Message)
}

var (
	// Refused before acting.
	safeCodes = map[string]bool{"ConcurrentOperationException": true, "ResourceConflictException": true}
	// A server-side failure that may have acted.
	maybeCodes = map[string]bool{
		"HandlerInternalFailureException": true, "HandlerFailureException": true, "ServiceInternalErrorException": true,
		"NetworkFailureException": true, "NotStabilizedException": true, "ConcurrentModificationException": true,
	}
	// Handler outcomes worth another attempt, though the failed request may have changed something.
	maybeHandlerCodes = map[string]bool{
		"NotStabilized": true, "ServiceInternalError": true, "InternalFailure": true, "NetworkFailure": true,
		"ServiceTimeout": true, "Throttling": true, "ResourceConflict": true,
	}
)

// classify answers infrata's question, how dangerous is another attempt, from the error alone. It must be a pure
// function: the plugin SDK may ask any configured instance, not the one that failed.
//
// Order matters: a throttle is checked before the HTTP status, and a refused dial before the generic network error
// (a *url.Error is a net.Error).
func classify(err error) provider.Retryability {
	if err == nil || errors.Is(err, context.Canceled) {
		return provider.NotSafeToRetry
	}
	var he *HandlerError
	if errors.As(err, &he) {
		if maybeHandlerCodes[he.Code] {
			return provider.ConditionallyRetryable
		}
		return provider.NotSafeToRetry
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		code := apiErr.ErrorCode()
		if _, throttled := retry.DefaultThrottleErrorCodes[code]; throttled || safeCodes[code] {
			return provider.SafeToRetry
		}
		if maybeCodes[code] {
			return provider.ConditionallyRetryable
		}
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return provider.SafeToRetry // no connection, so nothing was sent
	}
	var respErr *smithyhttp.ResponseError
	if errors.As(err, &respErr) && respErr.HTTPStatusCode() >= 500 {
		return provider.ConditionallyRetryable
	}
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.As(err, &netErr) {
		return provider.ConditionallyRetryable // sent, answer lost
	}
	return provider.NotSafeToRetry
}

// errorCode is the AWS error code or handler error code an error carries, or "".
func errorCode(err error) string {
	var he *HandlerError
	if errors.As(err, &he) {
		return he.Code
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		return apiErr.ErrorCode()
	}
	return ""
}

func isNotFound(err error) bool {
	switch errorCode(err) {
	case "ResourceNotFoundException", "NotFound":
		return true
	}
	return false
}

// apiFailure is the message a user reads in a failed apply, with the cause kept underneath for classification.
type apiFailure struct {
	msg string
	err error
}

func (e *apiFailure) Error() string { return e.msg }
func (e *apiFailure) Unwrap() error { return e.err }

// failure says which instance, which call, which type, where, what AWS said and the ID AWS support asks for. Never
// anything from the request itself, which may hold a secret.
func failure(instance, action string, t *catalog.Type, where string, err error) error {
	msg := fmt.Sprintf("aws instance %q: Cloud Control %s of %s (%s) at %s failed: ", instance, action, t.Name, t.CFN, where)
	var he *HandlerError
	var apiErr smithy.APIError
	switch {
	case errors.As(err, &he):
		msg += he.Code + ": " + he.Message + " (request token " + he.Token + ")"
	case errors.As(err, &apiErr):
		msg += apiErr.ErrorCode() + ": " + apiErr.ErrorMessage()
	default:
		msg += err.Error()
	}
	var re *awshttp.ResponseError
	if errors.As(err, &re) && re.ServiceRequestID() != "" {
		msg += " (request ID " + re.ServiceRequestID() + ")"
	}
	switch errorCode(err) {
	case "AccessDenied", "AccessDeniedException", "UnauthorizedTaggingOperation":
		msg += fmt.Sprintf("\ncheck the IAM permissions of the credentials instance %q uses", instance)
	case "InvalidCredentials", "InvalidCredentialsException", "UnrecognizedClientException", "ExpiredTokenException":
		msg += fmt.Sprintf("\nthe credentials instance %q uses were refused: check `profile` or `assume_role_arn`", instance)
	case "TypeNotFoundException":
		msg += "\nCloud Control does not offer " + t.CFN + " in this region"
	case "AlreadyExists", "AlreadyExistsException":
		msg += "\nsomething with that name already exists: adopt it with `infrata import`, or choose another name"
	}
	return &apiFailure{msg: msg, err: err}
}
```

`internal/ccprov/patience.go`:

```go
package ccprov

import (
	"context"
	"time"
)

// patience is how long a read waits for a resource AWS may not have propagated yet.
type patience struct {
	attempts int
	base     time.Duration
	max      time.Duration
	sleep    func(context.Context, time.Duration) error
}

// notFoundPatience: 5 attempts, 250ms doubling, capped at 2s: at most 3.75s of waiting.
var notFoundPatience = patience{attempts: 5, base: 250 * time.Millisecond, max: 2 * time.Second, sleep: sleepCtx}

// once reads a single time: Import and Discover, where the identifier came from AWS a moment ago.
var once = patience{attempts: 1}

// wait calls try until it reports found, returns an error, or the attempts run out. An error stops at once: only
// absence is worth waiting out.
func (pt patience) wait(ctx context.Context, try func() (bool, error)) (bool, error) {
	delay := pt.base
	for attempt := 1; ; attempt++ {
		found, err := try()
		if err != nil || found {
			return found, err
		}
		if attempt >= pt.attempts {
			return false, nil
		}
		if err := pt.sleep(ctx, delay); err != nil {
			return false, err
		}
		delay = min(delay*2, pt.max)
	}
}

func sleepCtx(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
```

`internal/ccprov/clients.go`:

```go
package ccprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// clients holds one Cloud Control client per region, built on first use. The SDK's standard retryer stays on: every
// mutation carries a ClientToken, so a resend cannot make a second resource.
type clients struct {
	cfg      aws.Config
	mu       sync.Mutex
	byRegion map[string]*cloudcontrol.Client
}

func newClients(cfg aws.Config) *clients {
	return &clients{cfg: cfg, byRegion: map[string]*cloudcontrol.Client{}}
}

func (c *clients) get(region string) *cloudcontrol.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.byRegion[region]; ok {
		return cl
	}
	cl := cloudcontrol.NewFromConfig(c.cfg, func(o *cloudcontrol.Options) { o.Region = region })
	c.byRegion[region] = cl
	return cl
}
```

`internal/ccprov/await.go`:

```go
package ccprov

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
)

type statusAPI interface {
	GetResourceRequestStatus(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error)
}

// pacing is how often a request's status is checked.
type pacing struct {
	first, max time.Duration
	sleep      func(context.Context, time.Duration) error
	now        func() time.Time
}

var defaultPacing = pacing{first: time.Second, max: 10 * time.Second, sleep: sleepCtx, now: time.Now}

// defaultTimeoutMinutes is CloudFormation's handler timeout when a schema names none.
const defaultTimeoutMinutes = 120

func timeoutFor(t *catalog.Type, handler string) time.Duration {
	if m := t.Timeouts[handler]; m > 0 {
		return time.Duration(m) * time.Minute
	}
	return defaultTimeoutMinutes * time.Minute
}

// waitTimeout unwraps to context.DeadlineExceeded, so it classifies as ConditionallyRetryable.
type waitTimeout struct {
	token  string
	status types.OperationStatus
	after  time.Duration
}

func (e *waitTimeout) Error() string {
	return fmt.Sprintf("request %s was still %s after %s; it may yet finish, so run `infrata refresh` before trying again",
		e.token, e.status, e.after)
}
func (e *waitTimeout) Unwrap() error { return context.DeadlineExceeded }

// await follows a request until it settles. Once a request is sent, cancellation never abandons it: the host waits for
// the real answer, because an abandoned create is a resource nothing records. The handler's timeout bounds the wait.
// The last event is returned with any error, so a caller can still see an identifier AWS assigned.
func (pc pacing) await(ctx context.Context, api statusAPI, ev *types.ProgressEvent, timeout time.Duration) (*types.ProgressEvent, error) {
	ctx = context.WithoutCancel(ctx)
	token := aws.ToString(ev.RequestToken)
	deadline := pc.now().Add(timeout)
	delay := pc.first
	for {
		switch ev.OperationStatus {
		case types.OperationStatusSuccess:
			return ev, nil
		case types.OperationStatusFailed, types.OperationStatusCancelComplete:
			code := string(ev.ErrorCode)
			if code == "" {
				code = "Cancelled"
			}
			return ev, &HandlerError{Operation: string(ev.Operation), Code: code, Message: aws.ToString(ev.StatusMessage),
				Identifier: aws.ToString(ev.Identifier), Token: token}
		}
		now := pc.now()
		if !now.Before(deadline) {
			return ev, &waitTimeout{token: token, status: ev.OperationStatus, after: timeout}
		}
		wait := delay
		if ev.RetryAfter != nil {
			if d := ev.RetryAfter.Sub(now); d > wait {
				wait = d
			}
		}
		wait = min(wait, deadline.Sub(now))
		_ = pc.sleep(ctx, wait) // ctx is never cancelled here
		delay = min(delay*2, pc.max)
		out, err := api.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: aws.String(token)})
		if err != nil {
			return ev, fmt.Errorf("checking request %s: %w", token, err)
		}
		ev = out.ProgressEvent
	}
}
```

- [ ] **Step 5: Run the tests**

```bash
go test -count=1 ./internal/ccprov/ && go vet ./internal/ccprov/ && gofmt -l internal/ccprov/
```

Expected: PASS, no vet findings, no unformatted files.

- [ ] **Step 6: Sabotage, then commit**

Sabotages: check the HTTP status before the throttle code (`ThrottlingException` at 429 still passes, so also inject
`ThrottlingException` at 503 in a scratch case and see it become ConditionallyRetryable); drop `context.WithoutCancel`
in `await` (the cancellation test fails); ignore `RetryAfter` (the RetryAfter test fails); split IDs at the last slash
(the ARN case fails).

```bash
git add internal/ccprov/ids.go internal/ccprov/errors.go internal/ccprov/patience.go internal/ccprov/clients.go \
  internal/ccprov/await.go internal/ccprov/catalog_test.go internal/ccprov/helpers_test.go internal/ccprov/ids_test.go \
  internal/ccprov/errors_test.go internal/ccprov/await_test.go
git commit -m "Add the Cloud Control plumbing the provider builds on

Provider IDs, retry classification for Cloud Control errors and failed
requests, per region clients, and waiting on a request until it settles
without giving up when the caller cancels. Checked by breaking the throttle
order, cancellation, RetryAfter and the ID split." -- \
  internal/ccprov/ids.go internal/ccprov/errors.go internal/ccprov/patience.go internal/ccprov/clients.go \
  internal/ccprov/await.go internal/ccprov/catalog_test.go internal/ccprov/helpers_test.go internal/ccprov/ids_test.go \
  internal/ccprov/errors_test.go internal/ccprov/await_test.go
```

---
### Task 9: Create, Read, Delete and Import through Cloud Control

**Files:**
- Create: `internal/ccprov/provider.go`, `internal/ccprov/values.go`, `internal/ccprov/crud.go`,
  `internal/ccprov/pending.go` (removed again in Tasks 11 and 12), `internal/awstest/faketype.go`
- Test: `internal/ccprov/fake_test.go`, `internal/ccprov/values_test.go`, `internal/ccprov/crud_test.go`
- Modify: `internal/awsprov/plugin.go`, `internal/awsprov/config_test.go`
- Delete: `internal/awsprov/unconfigured.go`

**Interfaces:**
- Consumes: everything Task 8 produces; `catalog.Embedded` (Task 5); `ccfake` (Task 7); `awstest.Isolate` (Task 6);
  `loadAWSConfig`, `parseConfig`, `instanceConfig.DiscoverRegions/DiscoverTypes` (Task 6 and `first-slice`).
- Produces (used by Tasks 10 to 16):
  - `const PluginName = "aws"`; `type Options struct { DiscoverRegions, DiscoverTypes []string }`
  - `func New(instance string, cat *catalog.Catalog, cfg aws.Config, opts Options) *Provider`
  - `(*Provider).Name/Definitions/Create/Read/Delete/Import/ClassifyError` (Update and Discover in `pending.go` until Tasks 11 and 12)
  - `func (p *Provider) read(ctx, t *catalog.Type, region, id string, reference map[string]value.Value, pt patience) (*resource.ResourceState, error)`
  - `func (p *Provider) lookup(name string) (*catalog.Type, error)`; `func regionOf(t *catalog.Type, attrs map[string]value.Value) (string, error)`
  - `func plain(v value.Value) any`; `func infer(datum any) (value.Value, bool)`; `func coerce(t, a, v, datum) (value.Value, error)`
  - `func encodeAttr(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error)`
  - `func decodeAttr(t *catalog.Type, a *catalog.Attribute, datum any, reference *value.Value) (value.Value, bool, error)`
  - `func desiredJSON(t *catalog.Type, attrs map[string]value.Value) (map[string]any, error)`
  - `func stateFrom(t *catalog.Type, region, identifier string, props map[string]any, reference map[string]value.Value) (*resource.ResourceState, error)`
  - `func decodeProperties(doc string) (map[string]any, error)`
  - `func awstest.FakeType(t *catalog.Type, idPrefix string, defaults map[string]any) ccfake.TypeConfig`
  - test-only: `fakeProvider(t) (*Provider, *ccfake.Server, *bytes.Buffer)`, `sv`, `iv`, `desired`

Values cross at the top level only in this task: `encodeAttr` and `decodeAttr` copy nested values as they are. Task 10
replaces those two functions with nested reconciliation and the tags transform, and nothing else in this task changes.

- [ ] **Step 1: The fake, described from the catalog**

`internal/awstest/faketype.go`:

```go
package awstest

import (
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
)

// FakeType describes a catalog type to the Cloud Control fake the way its schema does, so the fake cannot drift from
// the catalog: create-only properties are the ones the catalog marks force-new on an updatable type, write-only ones
// are the catalog's, and every read-only string property gets a value derived from the identifier.
func FakeType(t *catalog.Type, idPrefix string, defaults map[string]any) ccfake.TypeConfig {
	tc := ccfake.TypeConfig{
		TypeName: t.CFN, Identifier: t.Identifier[0], IDPrefix: idPrefix,
		Defaults: defaults, WriteOnly: t.WriteOnly, ReadOnly: map[string]string{},
	}
	for _, a := range t.Attributes {
		if a.Computed && !a.Optional && a.Name != tc.Identifier && a.Kind == "string" {
			tc.ReadOnly[a.Name] = strings.ToLower(a.Name) + "-{id}"
		}
		if a.ForceNew && t.HasUpdate {
			tc.CreateOnly = append(tc.CreateOnly, a.Name)
		}
	}
	return tc
}
```

`internal/ccprov/fake_test.go`:

```go
package ccprov

import (
	"bytes"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/awstest"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

var fakePrefixes = map[string]string{
	"aws.vpc": "vpc-", "aws.securitygroup": "sg-", "aws.bucket": "bucket-", "aws.role": "role-",
	"aws.dbinstance": "db-", "aws.test.regioned": "reg-", "aws.test.child": "child-",
}

// fakeProvider is a real Provider over the test catalog, talking to a fresh fake through the real SDK, with every wait
// made instant. What it writes to stderr is captured.
func fakeProvider(t *testing.T) (*Provider, *ccfake.Server, *bytes.Buffer) {
	t.Helper()
	fake := ccfake.New()
	t.Cleanup(fake.Close)
	cat := testCatalog()
	defaults := map[string]map[string]any{
		"aws.vpc":  {"EnableDnsSupport": true, "InstanceTenancy": "default"},
		"aws.role": {"MaxSessionDuration": 3600},
	}
	for _, typ := range cat.Types {
		fake.Register(awstest.FakeType(typ, fakePrefixes[typ.Name], defaults[typ.Name]))
	}
	p := New("test", cat, testConfig(fake.URL, 1), Options{DiscoverRegions: []string{"us-east-1"}})
	p.patience, p.pacing = instantPatience, instantPacing
	var log bytes.Buffer
	p.log = &log
	return p, fake, &log
}

func sv(v string) value.Value { return value.String(v, value.SourceExplicit) }
func iv(n int64) value.Value  { return value.Int(n, value.SourceExplicit) }

func desired(typ string, attrs map[string]value.Value) *resource.DesiredResource {
	return &resource.DesiredResource{Address: address.Address{Name: "r"}, Type: typ, Attrs: attrs}
}
```

- [ ] **Step 2: Write the failing tests**

`internal/ccprov/values_test.go`:

```go
package ccprov

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/infrata/infrata/pkg/value"
)

func jsonDatum(t *testing.T, doc string) any {
	t.Helper()
	dec := json.NewDecoder(strings.NewReader(doc))
	dec.UseNumber()
	var out any
	if err := dec.Decode(&out); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestValuesCrossJSONAndComeBackEqual(t *testing.T) {
	v := value.Map(map[string]value.Value{
		"name": sv("web"), "count": iv(3), "ratio": value.Float(0.5, value.SourceExplicit), "on": value.Bool(true, value.SourceExplicit),
		"items": value.List([]value.Value{sv("a"), iv(2)}, value.SourceExplicit),
	}, value.SourceExplicit)
	raw, err := json.Marshal(plain(v))
	if err != nil {
		t.Fatal(err)
	}
	back, ok := infer(jsonDatum(t, string(raw)))
	if !ok || !back.Equal(v) {
		t.Fatalf("round trip = %v, want %v", back, v)
	}
}

func TestNullsAreAbsent(t *testing.T) {
	got, ok := infer(jsonDatum(t, `{"a": null, "b": [1, null], "c": "x"}`))
	items := got.Raw.(map[string]value.Value)
	if !ok || len(items) != 2 || len(items["b"].Raw.([]value.Value)) != 1 {
		t.Fatalf("infer = %v", got)
	}
	if _, ok := infer(nil); ok {
		t.Error("a top-level null was recorded")
	}
}

// TestReturnedValuesBecomeTheKindTheSchemaDeclares. A property typed as a union is a string attribute (spec §3.1):
// whatever AWS returns for it is recorded as JSON text, never refused.
func TestReturnedValuesBecomeTheKindTheSchemaDeclares(t *testing.T) {
	vpc, db := mustType(t, "aws.vpc"), mustType(t, "aws.dbinstance")
	cidr, _ := vpc.Attribute("CidrBlock")
	netmask, _ := vpc.Attribute("Ipv4NetmaskLength")
	dns, _ := vpc.Attribute("EnableDnsSupport")
	class, _ := db.Attribute("DBInstanceClass")

	if got, ok, err := decodeAttr(vpc, cidr, jsonDatum(t, `{"b": 1, "a": [true]}`), nil); err != nil || !ok || got.Raw != `{"a":[true],"b":1}` {
		t.Errorf("object into a string attribute = %v, %v", got, err)
	}
	if got, _, err := decodeAttr(vpc, netmask, jsonDatum(t, `16.0`), nil); err != nil || got.Raw != int64(16) {
		t.Errorf("16.0 into an integer attribute = %v, %v", got, err)
	}
	if got, _, err := decodeAttr(vpc, dns, jsonDatum(t, `"true"`), nil); err != nil || got.Raw != true {
		t.Errorf(`"true" into a boolean attribute = %v, %v`, got, err)
	}
	if _, _, err := decodeAttr(vpc, netmask, jsonDatum(t, `16.5`), nil); err == nil || !strings.Contains(err.Error(), "aws.vpc.Ipv4NetmaskLength") {
		t.Errorf("16.5 into an integer attribute: err = %v", err)
	}
	if got, _, err := decodeAttr(db, class, jsonDatum(t, `7`), nil); err != nil || got.Raw != "7" {
		t.Errorf("a number into a string attribute = %v, %v", got, err)
	}
}

func TestTheDesiredStateHoldsOnlyWhatConfigurationMaySet(t *testing.T) {
	body, err := desiredJSON(mustType(t, "aws.vpc"), map[string]value.Value{
		"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "VpcId": sv("vpc-from-state"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(body) != 1 || body["CidrBlock"] != "10.0.0.0/16" {
		t.Errorf("desired state = %v, want only CidrBlock: never the plugin's region, never a read-only property", body)
	}
}
```

`internal/ccprov/crud_test.go`:

```go
package ccprov

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

var ctx = context.Background()

func attr(t *testing.T, st *resource.ResourceState, name string) any {
	t.Helper()
	v, ok := st.Attributes[name]
	if !ok {
		t.Fatalf("state has no %s: %v", name, st.Attributes)
	}
	return v.Raw
}

func vpcAttrs(region string) map[string]value.Value {
	return map[string]value.Value{"region": sv(region), "CidrBlock": sv("10.0.0.0/16"), "Ipv4NetmaskLength": iv(16)}
}

func roleAttrs(name string) map[string]value.Value {
	return map[string]value.Value{
		"RoleName":                 sv(name),
		"AssumeRolePolicyDocument": value.Map(map[string]value.Value{"Version": sv("2012-10-17")}, value.SourceExplicit),
	}
}

func TestCreateReturnsWhatAWSReports(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	id, ok := strings.CutPrefix(st.ProviderID, "us-east-1/")
	if !ok || !strings.HasPrefix(id, "vpc-") {
		t.Fatalf("provider ID = %q", st.ProviderID)
	}
	for name, want := range map[string]any{
		"region": "us-east-1", "CidrBlock": "10.0.0.0/16", "VpcId": id,
		"EnableDnsSupport":     true,                        // chosen by AWS: Optional+Computed
		"DefaultSecurityGroup": "defaultsecuritygroup-" + id, // read-only
		"Ipv4NetmaskLength":    int64(16),                   // write-only, carried from what was asked
	} {
		if got := attr(t, st, name); got != want {
			t.Errorf("%s = %#v, want %#v", name, got, want)
		}
	}
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	if _, leaked := stored["region"]; leaked {
		t.Error("the plugin's region attribute was sent to AWS as a property")
	}
	if tokens := fake.Tokens("CreateResource"); len(tokens) != 1 || tokens[0] == "" {
		t.Errorf("client tokens = %v, want one", tokens)
	}
}

func TestRegionalTypesUseTheirRegionAndGlobalTypesUSEast1(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("eu-west-1"))); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("eu-west-1", "AWS::EC2::VPC")); n != 1 {
		t.Errorf("eu-west-1 holds %d VPCs, want 1", n)
	}
	role, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil {
		t.Fatal(err)
	}
	if role.ProviderID != "global/deploy" {
		t.Errorf("role provider ID = %q", role.ProviderID)
	}
	if _, ok := fake.Resource("us-east-1", "AWS::IAM::Role", "deploy"); !ok {
		t.Error("a global type was not created through us-east-1")
	}
	if _, has := role.Attributes["region"]; has {
		t.Error("a global type reported a region attribute")
	}
}

// TestATypeWithItsOwnRegionPropertyUsesAwsRegion (J9).
func TestATypeWithItsOwnRegionPropertyUsesAwsRegion(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.test.regioned", map[string]value.Value{
		"aws_region": sv("eu-west-1"), "Region": sv("ap-south-1"), "Name": sv("x"),
	}))
	if err != nil {
		t.Fatal(err)
	}
	if attr(t, st, "aws_region") != "eu-west-1" || attr(t, st, "Region") != "ap-south-1" {
		t.Errorf("state = %v", st.Attributes)
	}
	if n := len(fake.Resources("eu-west-1", "AWS::Test::Regioned")); n != 1 {
		t.Errorf("created in the wrong region: eu-west-1 holds %d", n)
	}
}

func TestARegionalTypeWithoutARegionIsRefusedBeforeAnyCall(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	_, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"CidrBlock": sv("10.0.0.0/16")}))
	if err == nil || !strings.Contains(err.Error(), "defaults: {region") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("CreateResource") != 0 {
		t.Error("CreateResource was called")
	}
}

// TestAFailedCreateThatLeftAResourceIsRecordedNotOrphaned. The host drops an errored create's result, so returning the
// error would leave a real VPC nothing records.
func TestAFailedCreateThatLeftAResourceIsRecordedNotOrphaned(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.FailNext("CREATE", "NotStabilized", "attachment did not settle", true)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil || st == nil {
		t.Fatalf("Create = %v, %v; want the state of what exists", st, err)
	}
	for _, want := range []string{`"test"`, "NotStabilized", "attachment did not settle", st.ProviderID} {
		if !strings.Contains(log.String(), want) {
			t.Errorf("stderr lacks %q:\n%s", want, log)
		}
	}
}

func TestAFailedCreateThatLeftNothingIsAnError(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.FailNext("CREATE", "ServiceLimitExceeded", "too many VPCs", false)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if st != nil || err == nil || !strings.Contains(err.Error(), "ServiceLimitExceeded") || !strings.Contains(err.Error(), "too many VPCs") {
		t.Fatalf("Create = %v, %v", st, err)
	}
	if p.ClassifyError(err) != provider.NotSafeToRetry {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestANameThatExistsSuggestsImporting(t *testing.T) {
	p, _, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy"))); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy"))); err == nil || !strings.Contains(err.Error(), "infrata import") {
		t.Fatalf("second create: err = %v", err)
	}
}

func TestCancellationAfterTheCreateIsSentStillRecordsTheResource(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.PollsToComplete = 2
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	p.pacing.sleep = func(context.Context, time.Duration) error { cancel(); return nil }
	st, err := p.Create(cctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil || st == nil {
		t.Fatalf("Create after cancellation = %v, %v", st, err)
	}
}

// TestCreateNeverReportsFailureOnceAWSCreated. AWS said SUCCESS but the resource is not readable yet: the configured
// values are recorded, and the next refresh reads the rest.
func TestCreateNeverReportsFailureOnceAWSCreated(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.HideFromGet("deploy", 100)
	st, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil || st == nil || st.ProviderID != "global/deploy" || attr(t, st, "RoleName") != "deploy" {
		t.Fatalf("Create = %+v, %v", st, err)
	}
	if !strings.Contains(log.String(), "could not be read back") {
		t.Errorf("stderr does not say so:\n%s", log)
	}
}

func TestReadWaitsOutPropagation(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "10.0.0.0/16"})
	fake.HideFromGet("vpc-1", 2)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "us-east-1/vpc-1"})
	if err != nil || st == nil {
		t.Fatalf("Read = %v, %v", st, err)
	}
	if n := fake.Calls("GetResource"); n != 3 {
		t.Errorf("GetResource calls = %d, want 3", n)
	}
}

func TestReadReportsGoneAfterPatience(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "us-east-1/vpc-missing"})
	if err != nil || st != nil {
		t.Fatalf("Read = %v, %v; want (nil, nil)", st, err)
	}
	if n := fake.Calls("GetResource"); n != notFoundPatience.attempts {
		t.Errorf("GetResource calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

func TestReadCarriesWriteOnlyValuesForward(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::RDS::DBInstance", "db-1", map[string]any{"DBInstanceIdentifier": "db-1"})
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.dbinstance", ProviderID: "us-east-1/db-1",
		Attributes: map[string]value.Value{"MasterUserPassword": sv("hunter2"), "DBInstanceClass": sv("db.t3.micro")}})
	if err != nil {
		t.Fatal(err)
	}
	if attr(t, st, "MasterUserPassword") != "hunter2" {
		t.Error("a write-only value AWS never returns was dropped")
	}
	if _, has := st.Attributes["DBInstanceClass"]; has {
		t.Error("a readable property AWS did not return was carried forward: drift would be hidden")
	}
}

// TestAnEmptyCollectionAWSOmitsStaysEmpty. `policies: []` in configuration, and AWS leaving the property out, would
// otherwise plan a change forever.
func TestAnEmptyCollectionAWSOmitsStaysEmpty(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	empty := value.List(nil, value.SourceExplicit)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "global/deploy",
		Attributes: map[string]value.Value{"Policies": empty}})
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := st.Attributes["Policies"]; !ok || !got.Equal(empty) {
		t.Errorf("Policies = %v, want []", got)
	}
	notEmpty := value.List([]value.Value{sv("x")}, value.SourceExplicit)
	st, _ = p.Read(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "global/deploy",
		Attributes: map[string]value.Value{"Policies": notEmpty}})
	if _, has := st.Attributes["Policies"]; has {
		t.Error("a non-empty list AWS no longer has was carried forward")
	}
}

func TestMalformedIDsNeverReachAWS(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	if _, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "vpc-1"}); err == nil {
		t.Error("an ID without a region was read")
	}
	if err := p.Delete(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "us-east-1/deploy"}); err == nil {
		t.Error("a global type was deleted by a regional ID")
	}
	if n := fake.Calls("GetResource") + fake.Calls("DeleteResource"); n != 0 {
		t.Errorf("%d calls were made", n)
	}
}

func TestDeleteRemovesAndTreatsGoneAsDone(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Delete(ctx, st); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("us-east-1", "AWS::EC2::VPC")); n != 0 {
		t.Fatalf("%d VPCs remain", n)
	}
	if err := p.Delete(ctx, st); err != nil {
		t.Errorf("deleting what is already gone = %v, want success", err)
	}
}

func TestADeleteThatFailsSaysWhy(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	fake.FailNext("DELETE", "ResourceConflict", "the VPC has dependencies", false)
	err = p.Delete(ctx, st)
	if err == nil || !strings.Contains(err.Error(), "the VPC has dependencies") {
		t.Fatalf("err = %v", err)
	}
	if p.ClassifyError(err) != provider.ConditionallyRetryable {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestImportReadsOnceAndSaysWhatIsMissing(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "172.16.0.0/16"})
	st, err := p.Import(ctx, "aws.vpc", "us-east-1/vpc-1")
	if err != nil || attr(t, st, "CidrBlock") != "172.16.0.0/16" || st.ProviderID != "us-east-1/vpc-1" {
		t.Fatalf("Import = %+v, %v", st, err)
	}
	before := fake.Calls("GetResource")
	if _, err := p.Import(ctx, "aws.vpc", "us-east-1/vpc-missing"); err == nil || !strings.Contains(err.Error(), "us-east-1/vpc-missing") {
		t.Errorf("importing what does not exist: err = %v", err)
	}
	if n := fake.Calls("GetResource") - before; n != 1 {
		t.Errorf("import made %d reads, want 1", n)
	}
}

func TestAnUnknownTypeIsRefused(t *testing.T) {
	p, _, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.nope", nil)); err == nil || !strings.Contains(err.Error(), "aws.nope") {
		t.Fatalf("err = %v", err)
	}
}

// TestNothingIsWrittenToStdout. Stdout is the plugin protocol's pipe.
func TestNothingIsWrittenToStdout(t *testing.T) {
	if p := New("x", testCatalog(), testConfig("http://127.0.0.1:1", 1), Options{}); p.log != os.Stderr {
		t.Error("the provider's log is not stderr")
	}
}
```

In `internal/awsprov/config_test.go`, add this test and the imports `context`,
`github.com/infrata/infrata-provider-aws/internal/catalog`, `github.com/infrata/infrata-provider-aws/internal/ccfake`
and `github.com/infrata/infrata/pkg/resource`:

```go
// TestAssumeRoleSignsCloudControlCallsWithTheAssumedCredentials. The static key would sign every call if the role were
// ignored, so the assertion cannot pass by accident.
func TestAssumeRoleSignsCloudControlCallsWithTheAssumedCredentials(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	vpc, _ := cat.Lookup("aws.vpc")
	fake.Register(awstest.FakeType(vpc, "vpc-", nil))
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "10.0.0.0/16"})
	awstest.Isolate(t, fake.URL)

	prov, err := NewPlugin().New(provider.Config{Instance: "deploy", Values: map[string]value.Value{
		"assume_role_arn": s("arn:aws:iam::123456789012:role/deploy"),
	}})
	if err != nil {
		t.Fatal(err)
	}
	st, err := prov.Read(context.Background(), &resource.ResourceState{Type: "aws.vpc", ProviderID: "us-east-1/vpc-1"})
	if err != nil || st == nil {
		t.Fatalf("Read = %v, %v", st, err)
	}
	if fake.Calls("AssumeRole") != 1 {
		t.Errorf("AssumeRole calls = %d, want 1", fake.Calls("AssumeRole"))
	}
	keys := fake.AccessKeys()
	if last := keys[len(keys)-1]; last != ccfake.AssumedAccessKey {
		t.Errorf("GetResource was signed with %q, want the assumed role's %q", last, ccfake.AssumedAccessKey)
	}
}
```

- [ ] **Step 3: Run them to see them fail**

Run: `go test -count=1 ./internal/ccprov/ ./internal/awsprov/`
Expected: FAIL — `undefined: New`, `undefined: plain`, `undefined: desiredJSON`, `undefined: awstest.FakeType`.

- [ ] **Step 4: Implement**

`internal/ccprov/provider.go`:

```go
package ccprov

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = "aws"

// Options is what an instance's configuration tells the provider beyond credentials.
type Options struct {
	DiscoverRegions []string // regions Discover scans for regional types
	DiscoverTypes   []string // what Discover lists when asked about everything; empty means the catalog's DiscoverDefault (P4)
}

// Provider is one configured instance: one account's credentials, every catalog type, any number of regions.
type Provider struct {
	instance string
	cat      *catalog.Catalog
	opts     Options
	clients  *clients
	patience patience
	pacing   pacing
	log      io.Writer // stderr: stdout is the plugin protocol
	token    func() string

	defsOnce sync.Once
	defs     []*schema.ResourceDefinition
}

var _ provider.Provider = (*Provider)(nil)

// New builds an instance. It makes no network call.
func New(instance string, cat *catalog.Catalog, cfg aws.Config, opts Options) *Provider {
	return &Provider{
		instance: instance, cat: cat, opts: opts, clients: newClients(cfg),
		patience: notFoundPatience, pacing: defaultPacing, log: os.Stderr, token: newToken,
	}
}

func (p *Provider) Name() string { return PluginName }

func (p *Provider) Definitions() []*schema.ResourceDefinition {
	p.defsOnce.Do(func() { p.defs = p.cat.Definitions() })
	return p.defs
}

// ClassifyError delegates to classify, a pure function of the error.
func (p *Provider) ClassifyError(err error) provider.Retryability { return classify(err) }

func (p *Provider) lookup(name string) (*catalog.Type, error) {
	t, ok := p.cat.Lookup(name)
	if !ok {
		return nil, fmt.Errorf("the aws plugin does not serve %q; run `infrata explain <type>` to check a name", name)
	}
	return t, nil
}

// regionOf is where a resource lives: its region attribute, or GlobalRegion for a global type.
func regionOf(t *catalog.Type, attrs map[string]value.Value) (string, error) {
	if t.Global() {
		return GlobalRegion, nil
	}
	region, ok := attrs[t.RegionAttr].AsString()
	if !ok || region == "" {
		return "", fmt.Errorf("%s needs %s: set it on the resource, or once for every resource with the provider's defaults: {%s: ...}",
			t.Name, t.RegionAttr, t.RegionAttr)
	}
	return region, nil
}

// assumedState records what configuration asked for, when AWS created a resource the plugin cannot read back yet.
func assumedState(t *catalog.Type, region, identifier string, attrs map[string]value.Value) *resource.ResourceState {
	out := map[string]value.Value{}
	for _, a := range t.Attributes {
		if v, ok := attrs[a.Name]; ok && !(a.Computed && !a.Optional) {
			out[a.Name] = v
		}
	}
	if !t.Global() {
		out[t.RegionAttr] = value.String(region, value.SourceProvider)
	}
	return &resource.ResourceState{Type: t.Name, ProviderID: FormatID(t, region, identifier), Attributes: out}
}

// newToken is a ClientToken: Cloud Control treats a repeat within 36 hours as the same request.
func newToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
```

`internal/ccprov/values.go`:

```go
package ccprov

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// plain converts an infrata value to a JSON datum.
func plain(v value.Value) any {
	switch raw := v.Raw.(type) {
	case []value.Value:
		out := make([]any, len(raw))
		for i, item := range raw {
			out[i] = plain(item)
		}
		return out
	case map[string]value.Value:
		out := make(map[string]any, len(raw))
		for k, item := range raw {
			out[k] = plain(item)
		}
		return out
	default:
		return raw // string, int64, float64 or bool
	}
}

// infer converts a JSON datum decoded with UseNumber, taking the kind from the JSON. A null is absent.
func infer(datum any) (value.Value, bool) {
	const src = value.SourceProvider
	switch d := datum.(type) {
	case string:
		return value.String(d, src), true
	case bool:
		return value.Bool(d, src), true
	case json.Number:
		if i, err := d.Int64(); err == nil {
			return value.Int(i, src), true
		}
		f, err := d.Float64()
		return value.Float(f, src), err == nil
	case []any:
		items := make([]value.Value, 0, len(d))
		for _, item := range d {
			if v, ok := infer(item); ok {
				items = append(items, v)
			}
		}
		return value.List(items, src), true
	case map[string]any:
		items := make(map[string]value.Value, len(d))
		for k, item := range d {
			if v, ok := infer(item); ok {
				items[k] = v
			}
		}
		return value.Map(items, src), true
	}
	return value.Value{}, false
}

// coerce makes a returned top-level value the kind the schema declares. A string attribute takes anything, as JSON
// text: the generator types unions and untyped properties as strings.
func coerce(t *catalog.Type, a *catalog.Attribute, v value.Value, datum any) (value.Value, error) {
	want, _ := value.ParseKind(a.Kind)
	if v.Kind == want {
		return v, nil
	}
	const src = value.SourceProvider
	switch want {
	case value.KindString:
		raw, err := json.Marshal(datum)
		if err == nil {
			return value.String(string(raw), src), nil
		}
	case value.KindFloat:
		if i, ok := v.Raw.(int64); ok {
			return value.Float(float64(i), src), nil
		}
	case value.KindInt:
		if f, ok := v.Raw.(float64); ok && f == math.Trunc(f) {
			return value.Int(int64(f), src), nil
		}
	case value.KindBool:
		if s, ok := v.Raw.(string); ok {
			if b, err := strconv.ParseBool(s); err == nil {
				return value.Bool(b, src), nil
			}
		}
	}
	return value.Value{}, fmt.Errorf("AWS returned a %s for %s.%s, which the schema says is a %s", v.Kind, t.Name, a.Name, want)
}

// encodeAttr converts one top-level attribute to the JSON Cloud Control takes.
func encodeAttr(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error) {
	return plain(v), nil
}

// decodeAttr converts one top-level property AWS returned. reference is what configuration asked for or the plugin last
// reported, when there is one. It reports false when there is nothing to record.
func decodeAttr(t *catalog.Type, a *catalog.Attribute, datum any, reference *value.Value) (value.Value, bool, error) {
	v, ok := infer(datum)
	if !ok {
		return value.Value{}, false, nil
	}
	v, err := coerce(t, a, v, datum)
	return v, err == nil, err
}

// desiredJSON is the Cloud Control desired state: every attribute configuration may set, under AWS's names. Never the
// plugin's region, never a read-only property (Update's desired attributes include observed ones).
func desiredJSON(t *catalog.Type, attrs map[string]value.Value) (map[string]any, error) {
	out := map[string]any{}
	for _, a := range t.Attributes {
		v, set := attrs[a.Name]
		if !set || (a.Computed && !a.Optional) {
			continue
		}
		enc, err := encodeAttr(t, a, v)
		if err != nil {
			return nil, err
		}
		out[a.Name] = enc
	}
	return out, nil
}

// decodeProperties parses GetResource's Properties, keeping numbers exact.
func decodeProperties(doc string) (map[string]any, error) {
	dec := json.NewDecoder(strings.NewReader(doc))
	dec.UseNumber()
	var props map[string]any
	if err := dec.Decode(&props); err != nil {
		return nil, err
	}
	return props, nil
}

// stateFrom builds the state infrata records from what AWS returned. reference is configuration's values (Create,
// Update), the previous state (Read), or nil (Discover, Import).
func stateFrom(t *catalog.Type, region, identifier string, props map[string]any, reference map[string]value.Value) (*resource.ResourceState, error) {
	attrs := map[string]value.Value{}
	for _, a := range t.Attributes {
		var ref *value.Value
		if r, ok := reference[a.Name]; ok && r.Known {
			ref = &r
		}
		datum, present := props[a.Name]
		if !present || datum == nil {
			switch {
			case ref != nil && slices.Contains(t.WriteOnly, a.Name):
				attrs[a.Name] = *ref // AWS never returns it
			case ref != nil && emptyCollection(*ref):
				attrs[a.Name] = *ref // AWS omits an empty list or map; an empty one was asked for
			}
			continue
		}
		v, ok, err := decodeAttr(t, a, datum, ref)
		if err != nil {
			return nil, err
		}
		if ok {
			attrs[a.Name] = v
		}
	}
	if !t.Global() {
		attrs[t.RegionAttr] = value.String(region, value.SourceProvider)
	}
	return &resource.ResourceState{Type: t.Name, ProviderID: FormatID(t, region, identifier), Attributes: attrs}, nil
}

func emptyCollection(v value.Value) bool {
	switch raw := v.Raw.(type) {
	case []value.Value:
		return len(raw) == 0
	case map[string]value.Value:
		return len(raw) == 0
	}
	return false
}
```

`internal/ccprov/crud.go`:

```go
package ccprov

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// Create sends the desired state, waits for the request, and reads the resource back. Once AWS may have created
// something it never returns an error for a resource that exists: the host drops an errored create's result, and the
// resource would be real and recorded nowhere.
func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	t, err := p.lookup(desired.Type)
	if err != nil {
		return nil, err
	}
	region, err := regionOf(t, desired.Attrs)
	if err != nil {
		return nil, err
	}
	body, err := desiredJSON(t, desired.Attrs)
	if err != nil {
		return nil, err
	}
	doc, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	cl := p.clients.get(region)
	out, err := cl.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName: aws.String(t.CFN), DesiredState: aws.String(string(doc)), ClientToken: aws.String(p.token()),
	})
	if err != nil {
		return nil, failure(p.instance, "CreateResource", t, region, err)
	}

	ctx = context.WithoutCancel(ctx) // from here AWS may have acted
	ev, waitErr := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "create"))
	id := aws.ToString(ev.Identifier)
	if id == "" {
		id = aws.ToString(out.ProgressEvent.Identifier)
	}
	if id == "" {
		if waitErr == nil {
			waitErr = fmt.Errorf("request %s succeeded without an identifier", aws.ToString(out.ProgressEvent.RequestToken))
		}
		return nil, failure(p.instance, "create", t, region, waitErr)
	}
	where := FormatID(t, region, id)

	st, readErr := p.read(ctx, t, region, id, desired.Attrs, p.patience)
	switch {
	case readErr == nil && st != nil:
		if waitErr != nil {
			fmt.Fprintf(p.log, "aws instance %q: creating %s %s: %v\nthe resource exists, so it is recorded rather than orphaned: check it before relying on it\n",
				p.instance, t.Name, where, waitErr)
		}
		return st, nil
	case waitErr != nil:
		return nil, failure(p.instance, "create", t, where, waitErr)
	}
	why := "it is not visible yet"
	if readErr != nil {
		why = readErr.Error()
	}
	fmt.Fprintf(p.log, "aws instance %q: created %s %s, but it could not be read back (%s); recording the configured values until the next refresh\n",
		p.instance, t.Name, where, why)
	return assumedState(t, region, id, desired.Attrs), nil
}

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	t, err := p.lookup(current.Type)
	if err != nil {
		return nil, err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return nil, err
	}
	return p.read(ctx, t, region, id, current.Attributes, p.patience)
}

// read is GetResource with patience for propagation: (nil, nil) when the resource is gone.
func (p *Provider) read(ctx context.Context, t *catalog.Type, region, id string, reference map[string]value.Value, pt patience) (*resource.ResourceState, error) {
	cl := p.clients.get(region)
	var desc *types.ResourceDescription
	found, err := pt.wait(ctx, func() (bool, error) {
		out, err := cl.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String(t.CFN), Identifier: aws.String(id)})
		if isNotFound(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		desc = out.ResourceDescription
		return true, nil
	})
	if err != nil {
		return nil, failure(p.instance, "GetResource", t, FormatID(t, region, id), err)
	}
	if !found {
		return nil, nil
	}
	props, err := decodeProperties(aws.ToString(desc.Properties))
	if err != nil {
		return nil, fmt.Errorf("aws instance %q: %s %s: AWS returned properties that are not a JSON object: %w",
			p.instance, t.Name, FormatID(t, region, id), err)
	}
	identifier := aws.ToString(desc.Identifier)
	if identifier == "" {
		identifier = id
	}
	return stateFrom(t, region, identifier, props, reference)
}

// Delete succeeds for a resource that is already gone, whenever that is discovered.
func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	t, err := p.lookup(current.Type)
	if err != nil {
		return err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return err
	}
	cl := p.clients.get(region)
	out, err := cl.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{
		TypeName: aws.String(t.CFN), Identifier: aws.String(id), ClientToken: aws.String(p.token()),
	})
	if isNotFound(err) {
		return nil
	}
	if err != nil {
		return failure(p.instance, "DeleteResource", t, current.ProviderID, err)
	}
	if _, err := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "delete")); err != nil && !isNotFound(err) {
		return failure(p.instance, "delete", t, current.ProviderID, err)
	}
	return nil
}

// Import reads once: the ID came from discovery a moment ago. Nested keys come back in snake_case, having no reference.
func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	t, err := p.lookup(resourceType)
	if err != nil {
		return nil, err
	}
	region, identifier, err := ParseID(t, id)
	if err != nil {
		return nil, err
	}
	st, err := p.read(ctx, t, region, identifier, nil, once)
	if err != nil {
		return nil, err
	}
	if st == nil {
		return nil, fmt.Errorf("aws instance %q: there is no %s %s (%s in %s)", p.instance, t.Name, id, t.CFN, region)
	}
	return st, nil
}
```

`internal/ccprov/pending.go`:

```go
package ccprov

import (
	"context"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
)

// Update and Discover are not built yet. Each says so rather than guessing.

func (p *Provider) Update(context.Context, *resource.ResourceState, *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Discover(context.Context, provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
```

`internal/awsprov/plugin.go`, whole file:

```go
// Package awsprov is infrata's AWS provider plugin: instance configuration and credentials around the generic Cloud
// Control provider.
package awsprov

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccprov"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/schema"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = ccprov.PluginName

// Version is reported in the handshake. "0.0.0-dev" in every build a release did not stamp, so a broken -ldflags path
// cannot pass the release gate by coincidence. scripts/build-release stamps it.
var Version = "0.0.0-dev"

// Plugin is the AWS provider before configuration.
type Plugin struct{}

// NewPlugin returns the AWS plugin.
func NewPlugin() *Plugin { return &Plugin{} }

var _ provider.Plugin = (*Plugin)(nil)

// Name is the plugin's name.
func (pl *Plugin) Name() string { return PluginName }

// Version reports this build's version.
func (pl *Plugin) Version() string { return Version }

// Definitions are the embedded catalog's. They need no configuration and make no network call.
func (pl *Plugin) Definitions() []*schema.ResourceDefinition {
	cat, err := catalog.Embedded()
	if err != nil {
		// Only a broken build gets here; the catalog tests fail first.
		fmt.Fprintln(os.Stderr, "the aws plugin's embedded catalog is unreadable:", err)
		return nil
	}
	return cat.Definitions()
}

// New builds one configured instance. An error here is rendered against the `providers:` entry, so it says what is
// wrong and what to set. It makes no network call: every compiling command constructs instances.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	ic, err := parseConfig(cfg.Values)
	if err != nil {
		return nil, err
	}
	cat, err := catalog.Embedded()
	if err != nil {
		return nil, err
	}
	var unknown []string
	for _, name := range ic.DiscoverTypes {
		if _, ok := cat.Lookup(name); !ok {
			unknown = append(unknown, strconv.Quote(name))
		}
	}
	if len(unknown) > 0 {
		return nil, fmt.Errorf("`discover_types` names %s, which the aws plugin does not serve; run `infrata explain <type>` to check a name",
			strings.Join(unknown, ", "))
	}
	awsCfg, err := loadAWSConfig(context.Background(), cfg.Instance, ic)
	if err != nil {
		return nil, err
	}
	return ccprov.New(cfg.Instance, cat, awsCfg, ccprov.Options{DiscoverRegions: ic.DiscoverRegions, DiscoverTypes: ic.DiscoverTypes}), nil
}
```

```bash
git rm -q internal/awsprov/unconfigured.go
```

- [ ] **Step 5: Run the suite**

```bash
go test -count=1 ./... && go vet ./... && gofmt -l .
```

Expected: PASS. `TestTheWholeCatalogLoadsThroughTheHost` (Task 6) still passes: nothing about definitions changed.

- [ ] **Step 6: Sabotage, then commit**

Sabotages: return the error when a FAILED create left a resource (the orphan test fails); drop
`context.WithoutCancel` in `Create` (the cancellation test fails: the read-back runs on a cancelled context); stop
carrying write-only values (the write-only test fails); send `region` in the desired state (the create test fails);
make `Delete` return the error for `NotFound` (the gone-is-done test fails).

```bash
git add internal/ccprov/provider.go internal/ccprov/values.go internal/ccprov/crud.go internal/ccprov/pending.go \
  internal/awstest/faketype.go internal/ccprov/fake_test.go internal/ccprov/values_test.go internal/ccprov/crud_test.go \
  internal/awsprov/plugin.go internal/awsprov/config_test.go
git commit -m "Create, read, delete and import any Cloud Control type

The plugin now builds a real provider. A create that AWS reports as failed
but that left a resource behind is recorded instead of orphaned, and a
cancel after sending still waits for the answer. Checked by breaking the
orphan case, cancellation, write-only carry, the region leak and delete of
something already gone." -- \
  internal/ccprov/provider.go internal/ccprov/values.go internal/ccprov/crud.go internal/ccprov/pending.go \
  internal/awstest/faketype.go internal/ccprov/fake_test.go internal/ccprov/values_test.go internal/ccprov/crud_test.go \
  internal/awsprov/plugin.go internal/awsprov/config_test.go internal/awsprov/unconfigured.go
```

---
### Task 10: Nested values reconcile, and tags are a map

**Files:**
- Create: `internal/ccprov/reconcile.go`, `internal/ccprov/tags.go`
- Test: `internal/ccprov/reconcile_test.go`, `internal/ccprov/tags_test.go`
- Modify: `internal/ccprov/values.go` (`encodeAttr` and `decodeAttr` only)

**Interfaces:**
- Consumes: `catalog.Shape` and its kinds (Task 2); `cfn.SnakeCase` (Task 1); `plain`, `infer`, `coerce`,
  `stateFrom`, `fakeProvider`, `sv`, `iv`, `desired`, `jsonDatum`, `attr` (Task 9); `awstest.FakeType` (Task 9).
- Produces (used by Tasks 11 to 16):
  - `func matchProp(props map[string]*catalog.Shape, key string) (string, bool)`
  - `func encode(shape *catalog.Shape, v value.Value, path string) (any, error)`
  - `func decode(shape *catalog.Shape, datum any, ref *value.Value) (value.Value, bool)`
  - `func sameScalar(ref value.Value, datum any) bool`
  - `func tagsToJSON(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error)`
  - `func tagsFromJSON(t *catalog.Type, a *catalog.Attribute, datum any, ref *value.Value) (value.Value, bool, error)`
  - `encodeAttr` and `decodeAttr` keep their Task 9 signatures

Why this exists (spec §3.4, J8): infrata's `value.Equal` needs nested maps to have exactly the same keys and lists to
match position by position, and it does not canonicalise inside them. Without this, a user who writes `from_port`,
an AWS handler that adds `Description: ""`, or a list AWS returns in another order would each plan an update forever.
Opaque shapes (P5) are copied exactly in both directions.

- [ ] **Step 1: Write the failing tests**

`internal/ccprov/reconcile_test.go`:

```go
package ccprov

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/awstest"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// obj builds a map value from key/value pairs.
func obj(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}

func list(items ...value.Value) value.Value { return value.List(items, value.SourceExplicit) }

func rule(protocol string, from, to int64, spelling map[string]string) value.Value {
	name := func(k string) string {
		if s, ok := spelling[k]; ok {
			return s
		}
		return k
	}
	return obj(name("IpProtocol"), sv(protocol), name("FromPort"), iv(from), name("ToPort"), iv(to), name("CidrIp"), sv("0.0.0.0/0"))
}

var snake = map[string]string{"IpProtocol": "ip_protocol", "FromPort": "from_port", "ToPort": "to_port", "CidrIp": "cidr_ip"}

func TestNestedKeysAreSentUnderAWSNamesWhateverTheSpelling(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	v := list(obj("ip_protocol", sv("tcp"), "FROMPORT", iv(443), "to_port", iv(443), "CidrIp", sv("0.0.0.0/0")))
	got, err := encodeAttr(sg, a, v)
	if err != nil {
		t.Fatal(err)
	}
	want := []any{map[string]any{"IpProtocol": "tcp", "FromPort": int64(443), "ToPort": int64(443), "CidrIp": "0.0.0.0/0"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("encoded = %#v\nwant      %#v", got, want)
	}
}

func TestAnUnknownNestedKeyIsRefusedNamingWhatIsAccepted(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	_, err := encodeAttr(sg, a, list(obj("ip_protocol", sv("tcp"), "port", iv(443))))
	for _, want := range []string{"aws.securitygroup.SecurityGroupIngress[0]", `"port"`, "FromPort (from_port)"} {
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("err = %v, want it to mention %s", err, want)
		}
	}
	_, err = encodeAttr(sg, a, list(obj("cidr_ip", sv("a"), "CidrIp", sv("b"))))
	if err == nil || !strings.Contains(err.Error(), "twice") {
		t.Errorf("two spellings of one key: err = %v", err)
	}
}

// TestOpaqueValuesAreSentAndReturnedExactly (P5). A policy document's keys are data, not schema names.
func TestOpaqueValuesAreSentAndReturnedExactly(t *testing.T) {
	role := mustType(t, "aws.role")
	a, _ := role.Attribute("AssumeRolePolicyDocument")
	doc := obj("Version", sv("2012-10-17"), "Statement", list(obj("Effect", sv("Allow"), "principal", obj("Service", sv("ec2.amazonaws.com")))))
	got, err := encodeAttr(role, a, doc)
	if err != nil {
		t.Fatal(err)
	}
	stmt := got.(map[string]any)["Statement"].([]any)[0].(map[string]any)
	if _, kept := stmt["principal"]; !kept {
		t.Errorf("an opaque key was translated: %v", stmt)
	}
	back, ok, err := decodeAttr(role, a, jsonDatum(t, `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","principal":{"Service":"ec2.amazonaws.com"}}]}`), nil)
	if err != nil || !ok || !back.Equal(doc) {
		t.Errorf("decoded without a reference = %v, want the keys exactly as AWS returned them", back)
	}
}

func TestReturnedValuesUseTheReferenceSpellingAndDropKeysAWSAdded(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0","Description":""}]`)
	got, ok, err := decodeAttr(sg, a, datum, &ref)
	if err != nil || !ok || !got.Equal(ref) {
		t.Fatalf("decoded = %v, want %v", got, ref)
	}
}

func TestUnorderedListsComeBackInTheReferenceOrder(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 80, 80, nil), rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[
		{"IpProtocol":"udp","FromPort":53,"ToPort":53,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":80,"ToPort":80,"CidrIp":"0.0.0.0/0"}]`)
	got, _, err := decodeAttr(sg, a, datum, &ref)
	if err != nil {
		t.Fatal(err)
	}
	items := got.Raw.([]value.Value)
	refItems := ref.Raw.([]value.Value)
	if len(items) != 3 || !items[0].Equal(refItems[0]) || !items[1].Equal(refItems[1]) {
		t.Fatalf("decoded = %v", got)
	}
	if extra := items[2].Raw.(map[string]value.Value); extra["from_port"].Raw != int64(53) {
		t.Errorf("the item AWS added = %v, want it appended in snake_case", items[2])
	}
}

// TestAChangedItemKeepsItsSpellingAndStillShowsTheChange. Drift inside a list is a real change and must stay visible.
func TestAChangedItemKeepsItsSpellingAndStillShowsTheChange(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 80, 80, nil), rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[
		{"IpProtocol":"tcp","FromPort":443,"ToPort":444,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":80,"ToPort":80,"CidrIp":"0.0.0.0/0"}]`)
	got, _, _ := decodeAttr(sg, a, datum, &ref)
	items, refItems := got.Raw.([]value.Value), ref.Raw.([]value.Value)
	if len(items) != 2 || !items[0].Equal(refItems[0]) {
		t.Fatalf("decoded = %v", got)
	}
	changed := items[1].Raw.(map[string]value.Value)
	if changed["to_port"].Raw != int64(444) || items[1].Equal(refItems[1]) {
		t.Errorf("the changed item = %v, want snake_case keys and to_port 444", items[1])
	}
}

func TestOrderedListsInsideObjectsPairByPosition(t *testing.T) {
	bucket := mustType(t, "aws.bucket")
	a, _ := bucket.Attribute("LifecycleConfiguration")
	ref := obj("rules", list(obj("id", sv("archive"), "status", sv("Enabled"),
		"transitions", list(obj("storage_class", sv("GLACIER"), "transition_in_days", iv(30)), obj("StorageClass", sv("DEEP_ARCHIVE"), "TransitionInDays", iv(180))))))
	datum := jsonDatum(t, `{"Rules":[{"Id":"archive","Status":"Enabled","Prefix":"",
		"Transitions":[{"StorageClass":"GLACIER","TransitionInDays":30},{"StorageClass":"DEEP_ARCHIVE","TransitionInDays":180}]}]}`)
	got, _, err := decodeAttr(bucket, a, datum, &ref)
	if err != nil || !got.Equal(ref) {
		t.Fatalf("decoded = %v, %v\nwant %v", got, err, ref)
	}
}

func TestWithNoReferenceNestedKeysAreSnakeCase(t *testing.T) {
	bucket := mustType(t, "aws.bucket")
	a, _ := bucket.Attribute("VersioningConfiguration")
	got, _, _ := decodeAttr(bucket, a, jsonDatum(t, `{"Status":"Enabled"}`), nil)
	if !got.Equal(obj("status", sv("Enabled"))) {
		t.Errorf("decoded = %v, want {status: Enabled}", got)
	}
}

// TestAScalarAWSReturnsInAnotherFormKeepsTheWrittenForm. `from_port: "443"` and AWS's 443 are one value.
func TestAScalarAWSReturnsInAnotherFormKeepsTheWrittenForm(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(obj("ip_protocol", sv("tcp"), "from_port", sv("443"), "to_port", sv("443"), "cidr_ip", sv("0.0.0.0/0")))
	got, _, _ := decodeAttr(sg, a, jsonDatum(t, `[{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0"}]`), &ref)
	if !got.Equal(ref) {
		t.Errorf("decoded = %v, want %v", got, ref)
	}
}

// TestNestedValuesConvergeThroughTheProvider: the user's spelling survives create and every later read, although AWS
// adds keys and reorders the list.
func TestNestedValuesConvergeThroughTheProvider(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	cfg := awstest.FakeType(mustType(t, "aws.securitygroup"), "sg-", nil)
	cfg.OnRead = func(props map[string]any) {
		rules, _ := props["SecurityGroupIngress"].([]any)
		for _, r := range rules {
			r.(map[string]any)["Description"] = ""
		}
		slices.Reverse(rules)
	}
	fake.Register(cfg)
	ingress := list(rule("tcp", 80, 80, snake), rule("tcp", 443, 443, map[string]string{"FromPort": "fromPort"}))

	st, err := p.Create(ctx, desired("aws.securitygroup", map[string]value.Value{
		"region": sv("us-east-1"), "GroupDescription": sv("web"), "SecurityGroupIngress": ingress,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if got := st.Attributes["SecurityGroupIngress"]; !got.Equal(ingress) {
		t.Fatalf("after create = %v\nwant %v: a plan would never converge", got, ingress)
	}
	again, err := p.Read(ctx, &resource.ResourceState{Type: st.Type, ProviderID: st.ProviderID, Attributes: st.Attributes})
	if err != nil || !again.Attributes["SecurityGroupIngress"].Equal(ingress) {
		t.Fatalf("after read = %v, %v", again.Attributes["SecurityGroupIngress"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::SecurityGroup", id)
	first := stored["SecurityGroupIngress"].([]any)[0].(map[string]any)
	if _, ok := first["FromPort"]; !ok {
		t.Errorf("AWS was sent %v, want AWS's names", first)
	}
}
```

`internal/ccprov/tags_test.go`:

```go
package ccprov

import (
	"reflect"
	"strings"
	"testing"

	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func TestTagsAreAMapInConfigurationAndAListInAWS(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	got, err := encodeAttr(vpc, a, obj("team", sv("platform"), "app", sv("web")))
	if err != nil {
		t.Fatal(err)
	}
	want := []any{map[string]any{"Key": "app", "Value": "web"}, map[string]any{"Key": "team", "Value": "platform"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("encoded = %#v", got)
	}
	if _, err := encodeAttr(vpc, a, list(sv("team"))); err == nil || !strings.Contains(err.Error(), "map of tag names") {
		t.Errorf("tags as a list: err = %v", err)
	}
}

// TestSystemTagsAreNeverReported. `aws:` tags are AWS's; configuration cannot set them, so reporting them is a diff
// nobody can resolve.
func TestSystemTagsAreNeverReported(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	got, _, err := decodeAttr(vpc, a, jsonDatum(t, `[{"Key":"team","Value":"platform"},{"Key":"aws:cloudformation:stack-name","Value":"x"}]`), nil)
	if err != nil || !got.Equal(obj("team", sv("platform"))) {
		t.Errorf("decoded = %v, %v", got, err)
	}
}

func TestATagValueWrittenAsANumberKeepsItsForm(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	ref := obj("cost_centre", iv(1234))
	sent, _ := encodeAttr(vpc, a, ref)
	if sent.([]any)[0].(map[string]any)["Value"] != "1234" {
		t.Errorf("sent %v, want the value as text", sent)
	}
	got, _, _ := decodeAttr(vpc, a, jsonDatum(t, `[{"Key":"cost_centre","Value":"1234"}]`), &ref)
	if !got.Equal(ref) {
		t.Errorf("decoded = %v, want %v", got, ref)
	}
}

func TestTagsRoundTripThroughTheProviderAndDriftShows(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	tags := obj("team", sv("platform"))
	st, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "Tags": tags}))
	if err != nil || !st.Attributes["Tags"].Equal(tags) {
		t.Fatalf("after create Tags = %v, %v", st.Attributes["Tags"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	stored["Tags"] = []any{map[string]any{"Key": "team", "Value": "someone-else"}}
	fake.Put("us-east-1", "AWS::EC2::VPC", id, stored)
	got, err := p.Read(ctx, st)
	if err != nil || got.Attributes["Tags"].Equal(tags) {
		t.Fatalf("after drift Tags = %v, %v; want the changed value", got.Attributes["Tags"], err)
	}
}

func TestAnEmptyTagMapStaysEmpty(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	none := value.Map(map[string]value.Value{}, value.SourceExplicit)
	st, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "Tags": none}))
	if err != nil || !st.Attributes["Tags"].Equal(none) {
		t.Fatalf("after create Tags = %v, %v", st.Attributes["Tags"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	delete(stored, "Tags")
	fake.Put("us-east-1", "AWS::EC2::VPC", id, stored)
	got, _ := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: st.ProviderID, Attributes: st.Attributes})
	if !got.Attributes["Tags"].Equal(none) {
		t.Errorf("AWS omitting an empty tag list: Tags = %v, want {}", got.Attributes["Tags"])
	}
}
```

- [ ] **Step 2: Run them to see them fail**

Run: `go test -count=1 ./internal/ccprov/`
Expected: FAIL — the nested-key test sends `ip_protocol` unchanged, the tags tests see a map sent as a map, and
`TestNestedValuesConvergeThroughTheProvider` reports "a plan would never converge".

- [ ] **Step 3: Implement**

`internal/ccprov/reconcile.go`:

```go
package ccprov

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/cfn"
	"github.com/infrata/infrata/pkg/value"
)

// Nested values are reconciled here (spec §3.4, J8). Outgoing, every key is sent under AWS's name. Incoming, a value is
// rewritten to its reference: the reference's spelling, without keys AWS added, and in its order where the schema says
// order does not matter. Opaque shapes (P5) are copied exactly both ways.

// matchProp finds the schema property a written key names: exactly, ignoring case, or by its snake_case form. The
// generator proved no two properties of one object fold together (Verification log).
func matchProp(props map[string]*catalog.Shape, key string) (string, bool) {
	if _, ok := props[key]; ok {
		return key, true
	}
	lower := strings.ToLower(key)
	for name := range props {
		if strings.ToLower(name) == lower || cfn.SnakeCase(name) == lower {
			return name, true
		}
	}
	return "", false
}

func spellings(props map[string]*catalog.Shape) string {
	names := make([]string, 0, len(props))
	for name := range props {
		if snake := cfn.SnakeCase(name); snake != strings.ToLower(name) {
			name += " (" + snake + ")"
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func copiedExactly(shape *catalog.Shape) bool {
	return shape == nil || shape.Kind == catalog.ShapeOpaque || shape.Kind == catalog.ShapeScalar
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// encode converts what a user wrote to the JSON Cloud Control takes. path names the value in errors.
func encode(shape *catalog.Shape, v value.Value, path string) (any, error) {
	if copiedExactly(shape) {
		return plain(v), nil
	}
	switch shape.Kind {
	case catalog.ShapeObject:
		items, ok := v.Raw.(map[string]value.Value)
		if v.Kind != value.KindMap || !ok {
			return nil, fmt.Errorf("%s must be a map with the keys %s, got a %s", path, spellings(shape.Props), v.Kind)
		}
		out := make(map[string]any, len(items))
		written := make(map[string]string, len(items))
		for _, k := range sortedKeys(items) {
			name, found := matchProp(shape.Props, k)
			if !found {
				return nil, fmt.Errorf("%s has no key %q; it accepts %s", path, k, spellings(shape.Props))
			}
			if first, dup := written[name]; dup {
				return nil, fmt.Errorf("%s sets %s twice, as %q and %q", path, name, first, k)
			}
			written[name] = k
			enc, err := encode(shape.Props[name], items[k], path+"."+name)
			if err != nil {
				return nil, err
			}
			out[name] = enc
		}
		return out, nil
	case catalog.ShapeArray:
		items, ok := v.Raw.([]value.Value)
		if v.Kind != value.KindList || !ok {
			return nil, fmt.Errorf("%s must be a list, got a %s", path, v.Kind)
		}
		out := make([]any, len(items))
		for i, item := range items {
			enc, err := encode(shape.Item, item, fmt.Sprintf("%s[%d]", path, i))
			if err != nil {
				return nil, err
			}
			out[i] = enc
		}
		return out, nil
	}
	return plain(v), nil
}

// decode converts what AWS returned, shaped by ref when there is one. It reports false for a null.
func decode(shape *catalog.Shape, datum any, ref *value.Value) (value.Value, bool) {
	if datum == nil {
		return value.Value{}, false
	}
	if ref != nil && !ref.Known {
		ref = nil
	}
	if ref != nil && sameScalar(*ref, datum) {
		return *ref, true
	}
	if copiedExactly(shape) {
		return infer(datum)
	}
	switch shape.Kind {
	case catalog.ShapeObject:
		if m, ok := datum.(map[string]any); ok {
			return decodeObject(shape, m, ref), true
		}
	case catalog.ShapeArray:
		if items, ok := datum.([]any); ok {
			return decodeArray(shape, items, ref), true
		}
	}
	return infer(datum)
}

func decodeObject(shape *catalog.Shape, m map[string]any, ref *value.Value) value.Value {
	out := map[string]value.Value{}
	if refItems, ok := refMap(ref); ok {
		// The reference's keys, in its spelling. What AWS added is dropped; what AWS no longer has is absent.
		for rk, rv := range refItems {
			if name, found := matchProp(shape.Props, rk); found {
				if dv, ok := decode(shape.Props[name], m[name], &rv); ok {
					out[rk] = dv
				}
			}
		}
		return value.Map(out, value.SourceProvider)
	}
	for k, d := range m {
		if dv, ok := decode(shape.Props[k], d, nil); ok {
			out[cfn.SnakeCase(k)] = dv
		}
	}
	return value.Map(out, value.SourceProvider)
}

func decodeArray(shape *catalog.Shape, items []any, ref *value.Value) value.Value {
	refItems, hasRef := refList(ref)
	out := make([]value.Value, 0, len(items))
	if !hasRef || !shape.Unordered {
		for i, d := range items {
			var r *value.Value
			if i < len(refItems) {
				r = &refItems[i]
			}
			if dv, ok := decode(shape.Item, d, r); ok {
				out = append(out, dv)
			}
		}
		return value.List(out, value.SourceProvider)
	}

	// Unordered: each reference item takes the returned item equal to it, so the list comes back in the reference's
	// order.
	slots := make([]*value.Value, len(refItems))
	used := make([]bool, len(items))
	for i := range refItems {
		for j, d := range items {
			if used[j] {
				continue
			}
			if dv, ok := decode(shape.Item, d, &refItems[i]); ok && dv.Equal(refItems[i]) {
				slots[i], used[j] = &dv, true
				break
			}
		}
	}
	// A returned item nothing matched fills the next unmatched reference slot, keeping that item's spelling so the
	// change shows as a change; anything beyond is appended.
	var extra []value.Value
	next := 0
	for j, d := range items {
		if used[j] {
			continue
		}
		for next < len(slots) && slots[next] != nil {
			next++
		}
		var r *value.Value
		if next < len(slots) {
			r = &refItems[next]
		}
		dv, ok := decode(shape.Item, d, r)
		if !ok {
			continue
		}
		if next < len(slots) {
			slots[next] = &dv
			next++
		} else {
			extra = append(extra, dv)
		}
	}
	for _, s := range slots {
		if s != nil {
			out = append(out, *s)
		}
	}
	return value.List(append(out, extra...), value.SourceProvider)
}

func refMap(ref *value.Value) (map[string]value.Value, bool) {
	if ref == nil {
		return nil, false
	}
	m, ok := ref.Raw.(map[string]value.Value)
	return m, ok && ref.Kind == value.KindMap
}

func refList(ref *value.Value) ([]value.Value, bool) {
	if ref == nil {
		return nil, false
	}
	l, ok := ref.Raw.([]value.Value)
	return l, ok && ref.Kind == value.KindList
}

// sameScalar reports whether a returned scalar is the reference scalar written another way: 443 for "443".
func sameScalar(ref value.Value, datum any) bool {
	var text string
	switch d := datum.(type) {
	case string:
		text = d
	case json.Number:
		text = d.String()
	case bool:
		text = strconv.FormatBool(d)
	default:
		return false
	}
	switch r := ref.Raw.(type) {
	case string:
		return r == text
	case int64:
		return strconv.FormatInt(r, 10) == text
	case float64:
		f, err := strconv.ParseFloat(text, 64)
		return err == nil && f == r
	case bool:
		return strconv.FormatBool(r) == text
	}
	return false
}
```

`internal/ccprov/tags.go`:

```go
package ccprov

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/value"
)

// systemTagPrefix marks tags AWS sets itself. Configuration cannot set them, so they are never reported.
const systemTagPrefix = "aws:"

// tagsToJSON sends a tag map as the [{Key, Value}] list every such schema uses, sorted so a request is stable.
func tagsToJSON(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error) {
	items, ok := v.Raw.(map[string]value.Value)
	if v.Kind != value.KindMap || !ok {
		return nil, fmt.Errorf("%s.%s must be a map of tag names to values, e.g. {team: platform}, got a %s", t.Name, a.Name, v.Kind)
	}
	out := make([]any, 0, len(items))
	for _, k := range sortedKeys(items) {
		text, isString := items[k].AsString()
		if !isString {
			raw, err := json.Marshal(plain(items[k]))
			if err != nil {
				return nil, err
			}
			text = string(raw)
		}
		out = append(out, map[string]any{"Key": k, "Value": text})
	}
	return out, nil
}

// tagsFromJSON reports AWS's tag list as a map.
func tagsFromJSON(t *catalog.Type, a *catalog.Attribute, datum any, ref *value.Value) (value.Value, bool, error) {
	if datum == nil {
		return value.Value{}, false, nil
	}
	items, ok := datum.([]any)
	if !ok {
		return value.Value{}, false, fmt.Errorf("AWS returned %s.%s as %T, not a list of {Key, Value}", t.Name, a.Name, datum)
	}
	refItems, _ := refMap(ref)
	out := make(map[string]value.Value, len(items))
	for _, item := range items {
		tag, _ := item.(map[string]any)
		key, isString := tag["Key"].(string)
		if !isString {
			return value.Value{}, false, fmt.Errorf("AWS returned a tag without a Key in %s.%s", t.Name, a.Name)
		}
		if strings.HasPrefix(key, systemTagPrefix) {
			continue
		}
		if r, has := refItems[key]; has && sameScalar(r, tag["Value"]) {
			out[key] = r
			continue
		}
		text, _ := tag["Value"].(string)
		out[key] = value.String(text, value.SourceProvider)
	}
	return value.Map(out, value.SourceProvider), true, nil
}
```

In `internal/ccprov/values.go`, replace `encodeAttr` and `decodeAttr` with:

```go
// encodeAttr converts one top-level attribute to the JSON Cloud Control takes: nested keys under AWS's names, and tags
// as a list.
func encodeAttr(t *catalog.Type, a *catalog.Attribute, v value.Value) (any, error) {
	if a.Name == t.TagsAsMap {
		return tagsToJSON(t, a, v)
	}
	return encode(a.Shape, v, t.Name+"."+a.Name)
}

// decodeAttr converts one top-level property AWS returned, reconciled against reference when there is one. It reports
// false when there is nothing to record.
func decodeAttr(t *catalog.Type, a *catalog.Attribute, datum any, reference *value.Value) (value.Value, bool, error) {
	if a.Name == t.TagsAsMap {
		return tagsFromJSON(t, a, datum, reference)
	}
	v, ok := decode(a.Shape, datum, reference)
	if !ok {
		return value.Value{}, false, nil
	}
	v, err := coerce(t, a, v, datum)
	return v, err == nil, err
}
```

- [ ] **Step 4: Run the suite**

```bash
go test -count=1 ./internal/ccprov/ && go vet ./... && gofmt -l .
```

Expected: PASS, including every Task 9 test (the top-level behaviour is unchanged).

- [ ] **Step 5: Sabotage, then commit**

Sabotages: remove the snake_case comparison from `matchProp` (the nested-key test fails); keep every returned key
instead of the reference's (the AWS-added-keys test fails); skip the unordered matching and pair by position (the
order test and the provider convergence test fail); drop the `aws:` filter (the system-tags test fails); echo the
reference for a changed item (the change-shows test fails).

```bash
git add internal/ccprov/reconcile.go internal/ccprov/tags.go internal/ccprov/reconcile_test.go internal/ccprov/tags_test.go \
  internal/ccprov/values.go
git commit -m "Reconcile nested values so plans converge

Nested keys can be written in any spelling and go to AWS under its own
names. What comes back keeps the spelling that was written, drops keys AWS
added, and keeps the written order for lists where order does not matter.
Tags are a plain map. Checked by breaking each of those rules." -- \
  internal/ccprov/reconcile.go internal/ccprov/tags.go internal/ccprov/reconcile_test.go internal/ccprov/tags_test.go \
  internal/ccprov/values.go
```

---
### Task 11: Update through a JSON Patch

**Files:**
- Create: `internal/ccprov/patch.go`
- Test: `internal/ccprov/update_test.go`
- Modify: `internal/ccprov/pending.go` (remove `Update`)

**Interfaces:**
- Consumes: `encodeAttr` (Task 10); `read`, `lookup`, `ParseID`, `failure`, `errorCode`, `timeoutFor`, `pacing.await` (Tasks 8, 9);
  `ccfake.LastPatch`, `FailNext`, `Calls`, `Tokens` (Task 7); `fakeProvider`, `sv`, `iv`, `desired`, `attr`, `obj`, `list`, `rule`, `snake` (Tasks 9, 10).
- Produces:
  - `func patchOps(t *catalog.Type, current, desired map[string]value.Value) ([]map[string]any, error)`
  - `func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error)`

The rules (spec §3.3, and the Update fix recorded in the spec): one operation per top-level property whose desired
value differs from current, compared under AWS's names so a change of spelling alone sends nothing; `replace` when
current holds the property, `add` when it does not; never `remove`, because every settable property is
Optional+Computed and a property dropped from configuration keeps AWS's value (PLAN §14.1); never a read-only
property, nor anything `Computed && !Optional`, which Update's desired attributes carry because infrata fills in
observed values.

- [ ] **Step 1: Write the failing tests**

`internal/ccprov/update_test.go`:

```go
package ccprov

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// withChanges is what infrata hands Update: the state's attributes, observed values included, overlaid with configuration.
func withChanges(st *resource.ResourceState, changes map[string]value.Value) *resource.DesiredResource {
	attrs := map[string]value.Value{}
	for k, v := range st.Attributes {
		attrs[k] = v
	}
	for k, v := range changes {
		attrs[k] = v
	}
	return desired(st.Type, attrs)
}

func createVPC(t *testing.T, p *Provider, extra map[string]value.Value) *resource.ResourceState {
	t.Helper()
	attrs := map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16")}
	for k, v := range extra {
		attrs[k] = v
	}
	st, err := p.Create(ctx, desired("aws.vpc", attrs))
	if err != nil {
		t.Fatal(err)
	}
	return st
}

func TestAChangedPropertyIsReplacedAndReadBack(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err != nil {
		t.Fatal(err)
	}
	want := []map[string]any{{"op": "replace", "path": "/EnableDnsSupport", "value": false}}
	if !reflect.DeepEqual(fake.LastPatch(), want) {
		t.Errorf("patch = %v, want %v: observed values such as VpcId must never be patched", fake.LastPatch(), want)
	}
	if attr(t, got, "EnableDnsSupport") != false {
		t.Errorf("state after update = %v", got.Attributes)
	}
	if tokens := fake.Tokens("UpdateResource"); len(tokens) != 1 || tokens[0] == "" {
		t.Errorf("client tokens = %v", tokens)
	}
}

func TestAPropertyStateDoesNotHoldIsAdded(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil {
		t.Fatal(err)
	}
	policies := list(obj("PolicyName", sv("read"), "PolicyDocument", obj("Version", sv("2012-10-17"))))
	if _, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"Policies": policies})); err != nil {
		t.Fatal(err)
	}
	if patch := fake.LastPatch(); len(patch) != 1 || patch[0]["op"] != "add" || patch[0]["path"] != "/Policies" {
		t.Errorf("patch = %v, want one add of /Policies", patch)
	}
}

// TestASpellingChangeAloneSendsNothingAndConverges. The next plan compares the new spelling against state, so the
// state returned must carry it.
func TestASpellingChangeAloneSendsNothingAndConverges(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.securitygroup", map[string]value.Value{
		"region": sv("us-east-1"), "GroupDescription": sv("web"), "SecurityGroupIngress": list(rule("tcp", 443, 443, nil)),
	}))
	if err != nil {
		t.Fatal(err)
	}
	respelled := list(rule("tcp", 443, 443, snake))
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"SecurityGroupIngress": respelled}))
	if err != nil {
		t.Fatal(err)
	}
	if n := fake.Calls("UpdateResource"); n != 0 {
		t.Errorf("UpdateResource calls = %d, want 0", n)
	}
	if !got.Attributes["SecurityGroupIngress"].Equal(respelled) {
		t.Errorf("state = %v, want the new spelling", got.Attributes["SecurityGroupIngress"])
	}
}

// TestADroppedPropertyIsNeverRemoved (PLAN §14.1): infrata sends no desired value for it, and AWS's is kept.
func TestADroppedPropertyIsNeverRemoved(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)})
	changes := withChanges(st, map[string]value.Value{"Tags": obj("team", sv("platform"))})
	delete(changes.Attrs, "EnableDnsSupport")
	if _, err := p.Update(ctx, st, changes); err != nil {
		t.Fatal(err)
	}
	for _, op := range fake.LastPatch() {
		if op["op"] == "remove" || op["path"] == "/EnableDnsSupport" {
			t.Errorf("patch %v touches a property configuration dropped", fake.LastPatch())
		}
	}
}

func TestAWriteOnlyValueIsPatchedOnlyWhenItChanges(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.dbinstance", map[string]value.Value{
		"region": sv("us-east-1"), "DBInstanceIdentifier": sv("db-1"), "MasterUserPassword": sv("first"), "DBInstanceClass": sv("db.t3.micro"),
	}))
	if err != nil {
		t.Fatal(err)
	}
	got, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"MasterUserPassword": sv("second")}))
	if err != nil {
		t.Fatal(err)
	}
	if patch := fake.LastPatch(); len(patch) != 1 || patch[0]["path"] != "/MasterUserPassword" {
		t.Errorf("patch paths = %v", patch)
	}
	if attr(t, got, "MasterUserPassword") != "second" {
		t.Error("the new write-only value was not carried into state")
	}
	before := fake.Calls("UpdateResource")
	if _, err := p.Update(ctx, got, withChanges(got, map[string]value.Value{"DBInstanceClass": sv("db.t3.micro")})); err != nil {
		t.Fatal(err)
	}
	if fake.Calls("UpdateResource") != before {
		t.Error("an unchanged write-only value caused an update")
	}
}

func TestAChangeAWSWillNotMakeInPlaceSaysSo(t *testing.T) {
	p, _, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"CidrBlock": sv("10.1.0.0/16")}))
	if err == nil || !strings.Contains(err.Error(), "NotUpdatable") || !strings.Contains(err.Error(), "/CidrBlock") {
		t.Fatalf("err = %v", err)
	}
	if p.ClassifyError(err) != provider.NotSafeToRetry {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestAFailedUpdateRequestIsAnError(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	fake.FailNext("UPDATE", "InvalidRequest", "tenancy cannot be dedicated here", false)
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err == nil || !strings.Contains(err.Error(), "tenancy cannot be dedicated here") {
		t.Fatalf("err = %v", err)
	}
}

func TestATypeWithoutAnUpdateHandlerIsNeverPatched(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := &resource.ResourceState{Type: "aws.test.regioned", ProviderID: "us-east-1/reg-1", Attributes: map[string]value.Value{"Name": sv("a")}}
	_, err := p.Update(ctx, st, withChanges(st, map[string]value.Value{"Name": sv("b")}))
	if err == nil || !strings.Contains(err.Error(), "replacement") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("UpdateResource") != 0 {
		t.Error("UpdateResource was called")
	}
}

func TestAnUpdateAlreadySentSurvivesCancellation(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st := createVPC(t, p, nil)
	fake.PollsToComplete = 2
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	p.pacing.sleep = func(context.Context, time.Duration) error { cancel(); return nil }
	got, err := p.Update(cctx, st, withChanges(st, map[string]value.Value{"EnableDnsSupport": value.Bool(false, value.SourceExplicit)}))
	if err != nil || got == nil {
		t.Fatalf("Update after cancellation = %v, %v", got, err)
	}
}
```

- [ ] **Step 2: Run them to see them fail**

Run: `go test -count=1 -run 'Update|Patch|Replaced|Added|Spelling|Dropped|WriteOnlyValueIsPatched|InPlace|UpdateHandler|AlreadySent' ./internal/ccprov/`
Expected: FAIL — every Update returns `not implemented`.

- [ ] **Step 3: Implement**

`internal/ccprov/patch.go`:

```go
package ccprov

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// patchOps is the RFC 6902 patch from current to desired, one operation per top-level property whose value differs
// under AWS's names. Sorted by path, so a request is stable.
func patchOps(t *catalog.Type, current, desired map[string]value.Value) ([]map[string]any, error) {
	attrs := append([]*catalog.Attribute(nil), t.Attributes...)
	sort.Slice(attrs, func(i, j int) bool { return attrs[i].Name < attrs[j].Name })
	var ops []map[string]any
	for _, a := range attrs {
		v, set := desired[a.Name]
		if !set || (a.Computed && !a.Optional) {
			continue
		}
		want, err := encodeAttr(t, a, v)
		if err != nil {
			return nil, err
		}
		op := "add"
		if cur, has := current[a.Name]; has && cur.Known {
			if have, err := encodeAttr(t, a, cur); err == nil {
				if sameJSON(want, have) {
					continue
				}
				op = "replace"
			}
		}
		ops = append(ops, map[string]any{"op": op, "path": "/" + a.Name, "value": want})
	}
	return ops, nil
}

func sameJSON(a, b any) bool {
	ja, errA := json.Marshal(a)
	jb, errB := json.Marshal(b)
	return errA == nil && errB == nil && bytes.Equal(ja, jb)
}

// Update patches what changed, waits for the request, and reads the resource back with configuration as the reference.
// When nothing differs under AWS's names (a spelling change) it only reads back, so state takes the new spelling.
func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	t, err := p.lookup(current.Type)
	if err != nil {
		return nil, err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return nil, err
	}
	if !t.HasUpdate {
		return nil, fmt.Errorf("%s (%s) has no update handler, so every change needs a replacement; infrata was sent an update, which is a defect in the catalog's force-new flags",
			t.Name, t.CFN)
	}
	ops, err := patchOps(t, current.Attributes, desired.Attrs)
	if err != nil {
		return nil, err
	}
	if len(ops) > 0 {
		doc, err := json.Marshal(ops)
		if err != nil {
			return nil, err
		}
		cl := p.clients.get(region)
		out, err := cl.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
			TypeName: aws.String(t.CFN), Identifier: aws.String(id), PatchDocument: aws.String(string(doc)), ClientToken: aws.String(p.token()),
		})
		if err != nil {
			return nil, updateFailure(p.instance, "UpdateResource", t, current.ProviderID, ops, err)
		}
		ctx = context.WithoutCancel(ctx) // from here AWS may have acted
		if _, err := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "update")); err != nil {
			return nil, updateFailure(p.instance, "update", t, current.ProviderID, ops, err)
		}
	}
	st, err := p.read(ctx, t, region, id, desired.Attrs, p.patience)
	if err != nil {
		return nil, err
	}
	if st == nil {
		return nil, fmt.Errorf("aws instance %q: %s %s no longer exists after the update", p.instance, t.Name, current.ProviderID)
	}
	return st, nil
}

// updateFailure names the patched paths when AWS refuses to change them in place, which is the schema being wrong
// about what is create-only. Only paths: the values may be secret.
func updateFailure(instance, action string, t *catalog.Type, where string, ops []map[string]any, err error) error {
	wrapped := failure(instance, action, t, where, err)
	switch errorCode(err) {
	case "NotUpdatable", "NotUpdatableException":
		paths := make([]string, len(ops))
		for i, op := range ops {
			paths[i] = op["path"].(string)
		}
		return fmt.Errorf("%w\nAWS will not change %s in place, although the schema does not mark it create-only; replace the resource, and report the schema",
			wrapped, strings.Join(paths, ", "))
	}
	return wrapped
}
```

`internal/ccprov/pending.go`, whole file:

```go
package ccprov

import (
	"context"

	"github.com/infrata/infrata/pkg/provider"
)

// Discover is not built yet. It says so rather than guessing.
func (p *Provider) Discover(context.Context, provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
```

- [ ] **Step 4: Run the suite**

```bash
go test -count=1 ./internal/ccprov/ && go vet ./... && gofmt -l .
```

Expected: PASS.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: drop the `Computed && !Optional` skip (the replace test sees `/VpcId` and `/DefaultSecurityGroup`);
compare with `value.Equal` instead of encoded JSON (the spelling test sends an update); always use `replace` (the add
test fails); remove `context.WithoutCancel` (the cancellation test fails).

```bash
git add internal/ccprov/patch.go internal/ccprov/update_test.go internal/ccprov/pending.go
git commit -m "Update resources with a JSON patch of what changed

Only properties whose value differs are sent, compared under AWS's names so
a spelling change alone sends nothing. Observed values are never patched
and a property removed from the config keeps its AWS value. Checked by
breaking the observed value skip, the comparison, add versus replace and
cancellation." -- internal/ccprov/patch.go internal/ccprov/update_test.go internal/ccprov/pending.go
```

---
### Task 12: Discover

**Files:**
- Create: `internal/ccprov/discover.go`
- Test: `internal/ccprov/discover_test.go`
- Delete: `internal/ccprov/pending.go`

**Interfaces:**
- Consumes: `read`, `once`, `failure`, `FormatID`, `GlobalRegion`, `Options.DiscoverRegions/DiscoverTypes`,
  `catalog.DiscoverDefault`, `HasList`, `ListNeedsModel` (Tasks 2, 8, 9); `ccfake.Put`, `PageSize`, `Inject`, `Calls` (Task 7).
- Produces:
  - `func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error)`
  - `func (p *Provider) discoverTypes(requested []string) []string`

How infrata calls it (verified in `v0.3.0:internal/discovery/walk.go`): with no types named, `Walk` asks each instance
about every type it offers, which for this plugin is the whole catalog; `import` always walks with none. An error from
`Discover` drops all of that instance's results. So:

- A request covering the whole catalog means the instance's `discover_types`, else the catalog's `DiscoverDefault` (P4).
  A request naming fewer types was narrowed by the user and is honoured as given.
- Types with no list handler, or whose list needs a parent model, are skipped and named once on stderr (spec §4).
- One type or region failing is reported on stderr and the rest still returned. Only when every attempt fails is it
  an error.
- `ListResources` often returns only identifiers, so each is read with `GetResource`, once, with no reference: nested keys
  come back in snake_case, which is also what `Import` reports, so `import --generate` plans clean.

- [ ] **Step 1: Write the failing tests**

`internal/ccprov/discover_test.go`:

```go
package ccprov

import (
	"context"
	"errors"
	"sort"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/ccfake"
	"github.com/infrata/infrata/pkg/provider"
)

func everything() provider.DiscoverRequest {
	var names []string
	for _, typ := range testCatalog().Types {
		names = append(names, typ.Name)
	}
	return provider.DiscoverRequest{Types: names}
}

func ids(found []provider.DiscoveredResource) []string {
	out := make([]string, len(found))
	for i, r := range found {
		out[i] = r.Type + " " + r.ProviderID
	}
	sort.Strings(out)
	return out
}

func TestDiscoveringEverythingMeansTheDefaultSetInEveryRegion(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	p.opts.DiscoverRegions = []string{"us-east-1", "eu-west-1"}
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "10.0.0.0/16"})
	fake.Put("eu-west-1", "AWS::EC2::VPC", "vpc-2", map[string]any{"VpcId": "vpc-2", "CidrBlock": "10.1.0.0/16"})
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})

	found, err := p.Discover(ctx, everything())
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"aws.role global/deploy", "aws.vpc eu-west-1/vpc-2", "aws.vpc us-east-1/vpc-1"}
	if got := ids(found); strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("found %v, want %v (the bucket is outside the default set)", got, want)
	}
	if n := fake.Calls("ListResources"); n != 3 {
		t.Errorf("ListResources calls = %d, want 3: VPCs in two regions, roles once", n)
	}
	for _, r := range found {
		if r.Type == "aws.vpc" && r.Attributes["region"].Raw == nil {
			t.Errorf("%s has no region attribute: import --generate would write a resource that does not plan", r.ProviderID)
		}
	}
}

func TestDiscoverTypesReplacesTheDefaultSet(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	p.opts.DiscoverTypes = []string{"aws.bucket"}
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1"})
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})
	found, err := p.Discover(ctx, everything())
	if err != nil || strings.Join(ids(found), ",") != "aws.bucket us-east-1/logs" {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

func TestARequestNamingSomeTypesIsHonouredAsGiven(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::S3::Bucket", "logs", map[string]any{"BucketName": "logs"})
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.bucket"}})
	if err != nil || len(found) != 1 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

func TestTypesDiscoveryCannotListAreSkippedAndNamed(t *testing.T) {
	p, fake, log := fakeProvider(t)
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.test.child", "aws.test.regioned", "aws.vpc"}}); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"aws.test.child", "parent", "aws.test.regioned"} {
		if !strings.Contains(log.String(), want) {
			t.Errorf("stderr lacks %q:\n%s", want, log)
		}
	}
	if n := fake.Calls("ListResources"); n != 1 {
		t.Errorf("ListResources calls = %d, want 1 (aws.vpc only)", n)
	}
}

func TestEveryPageIsFollowed(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.PageSize = 1
	for _, id := range []string{"vpc-1", "vpc-2", "vpc-3"} {
		fake.Put("us-east-1", "AWS::EC2::VPC", id, map[string]any{"VpcId": id})
	}
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}})
	if err != nil || len(found) != 3 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
}

// TestDiscoveredValuesAreWhatImportReports: import --generate writes discovery's values, then imports; the two must
// agree or the next plan is not clean.
func TestDiscoveredValuesAreWhatImportReports(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::SecurityGroup", "sg-1", map[string]any{
		"GroupId": "sg-1", "GroupDescription": "web",
		"SecurityGroupIngress": []any{map[string]any{"IpProtocol": "tcp", "FromPort": 443, "ToPort": 443, "CidrIp": "0.0.0.0/0"}},
	})
	found, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.securitygroup"}})
	if err != nil || len(found) != 1 {
		t.Fatalf("found %v, %v", ids(found), err)
	}
	ingress := found[0].Attributes["SecurityGroupIngress"]
	if !ingress.Equal(list(rule("tcp", 443, 443, snake))) {
		t.Errorf("discovered ingress = %v, want snake_case keys", ingress)
	}
	imported, err := p.Import(ctx, "aws.securitygroup", found[0].ProviderID)
	if err != nil {
		t.Fatal(err)
	}
	for name, v := range found[0].Attributes {
		if !imported.Attributes[name].Equal(v) {
			t.Errorf("%s: discovered %v, imported %v", name, v, imported.Attributes[name])
		}
	}
}

func TestOneFailingTypeDoesNotHideTheRest(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	fake.Inject(ccfake.Fault{Action: "ListResources", Nth: 1, Status: 400, Code: "AccessDeniedException", Message: "no ec2:DescribeVpcs"})
	found, err := p.Discover(ctx, everything())
	if err != nil || strings.Join(ids(found), ",") != "aws.role global/deploy" {
		t.Fatalf("found %v, %v", ids(found), err)
	}
	if !strings.Contains(log.String(), "AccessDeniedException") || !strings.Contains(log.String(), "aws.vpc") {
		t.Errorf("stderr does not report the failed type:\n%s", log)
	}
}

func TestWhenEveryAttemptFailsDiscoverSaysSo(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	for n := 1; n <= 2; n++ {
		fake.Inject(ccfake.Fault{Action: "ListResources", Nth: n, Status: 400, Code: "AccessDeniedException", Message: "denied"})
	}
	if _, err := p.Discover(ctx, everything()); err == nil || !strings.Contains(err.Error(), "AccessDeniedException") {
		t.Fatalf("err = %v", err)
	}
}

func TestWithoutDiscoverRegionsOnlyGlobalTypesCanBeDiscovered(t *testing.T) {
	p, _, _ := fakeProvider(t)
	p.opts.DiscoverRegions = nil
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.vpc"}}); err == nil || !strings.Contains(err.Error(), "discover_regions") {
		t.Errorf("err = %v", err)
	}
	if _, err := p.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.role"}}); err != nil {
		t.Errorf("a global type needs no region: %v", err)
	}
}

func TestCancellationStopsDiscovery(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := p.Discover(cctx, everything()); !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
	if n := fake.Calls("ListResources"); n != 0 {
		t.Errorf("ListResources calls = %d after cancellation", n)
	}
}
```

- [ ] **Step 2: Run them to see them fail**

Run: `go test -count=1 -run 'Discover|Discovered|Discovery|Page|Failing|Fails' ./internal/ccprov/`
Expected: FAIL — `not implemented`.

- [ ] **Step 3: Implement**

`internal/ccprov/discover.go`:

```go
package ccprov

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
)

// Discover lists what exists for the requested types in the instance's regions, reading each resource once.
func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	var types []*catalog.Type
	var needsParent, unlistable []string
	for _, name := range p.discoverTypes(req.Types) {
		t, ok := p.cat.Lookup(name)
		switch {
		case !ok:
			continue
		case !t.HasList:
			unlistable = append(unlistable, name)
		case t.ListNeedsModel:
			needsParent = append(needsParent, name)
		default:
			types = append(types, t)
		}
	}
	if len(unlistable) > 0 {
		fmt.Fprintf(p.log, "aws instance %q: not discovering %s: AWS offers no way to list them\n", p.instance, strings.Join(unlistable, ", "))
	}
	if len(needsParent) > 0 {
		fmt.Fprintf(p.log, "aws instance %q: not discovering %s: listing them needs a parent resource, which discovery does not have\n",
			p.instance, strings.Join(needsParent, ", "))
	}
	if len(p.opts.DiscoverRegions) == 0 && slices.ContainsFunc(types, func(t *catalog.Type) bool { return !t.Global() }) {
		return nil, fmt.Errorf("aws instance %q has no `discover_regions`, so discovery does not know where to look: set it, e.g. discover_regions: [us-east-1]",
			p.instance)
	}

	var out []provider.DiscoveredResource
	var failed []error
	attempts := 0
	for _, t := range types {
		regions := p.opts.DiscoverRegions
		if t.Global() {
			regions = []string{GlobalRegion}
		}
		for _, region := range regions {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			attempts++
			found, err := p.discoverIn(ctx, t, region)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil, err
				}
				fmt.Fprintln(p.log, err)
				failed = append(failed, err)
				continue
			}
			out = append(out, found...)
		}
	}
	if attempts > 0 && len(failed) == attempts {
		return nil, errors.Join(failed...)
	}
	return out, nil
}

func (p *Provider) discoverIn(ctx context.Context, t *catalog.Type, region string) ([]provider.DiscoveredResource, error) {
	pages := cloudcontrol.NewListResourcesPaginator(p.clients.get(region), &cloudcontrol.ListResourcesInput{TypeName: aws.String(t.CFN)})
	var out []provider.DiscoveredResource
	for pages.HasMorePages() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, failure(p.instance, "ListResources", t, region, err)
		}
		for _, d := range page.ResourceDescriptions {
			st, err := p.read(ctx, t, region, aws.ToString(d.Identifier), nil, once)
			if err != nil {
				return nil, err
			}
			if st != nil { // gone between the list and the read
				out = append(out, provider.DiscoveredResource{Type: t.Name, ProviderID: st.ProviderID, Attributes: st.Attributes})
			}
		}
	}
	return out, nil
}

// discoverTypes is what a request covers (P4). A request for every type the catalog holds is infrata asking about
// everything, which here is ~1,584 ListResources calls per region: it means the instance's discover_types, or the
// catalog's default set. A request naming fewer was narrowed by the user.
func (p *Provider) discoverTypes(requested []string) []string {
	if len(requested) < len(p.cat.Types) {
		return requested
	}
	if len(p.opts.DiscoverTypes) > 0 {
		return p.opts.DiscoverTypes
	}
	return p.cat.DiscoverDefault
}
```

```bash
git rm -q internal/ccprov/pending.go
```

- [ ] **Step 4: Run the suite**

```bash
go test -count=1 ./... && go vet ./... && gofmt -l .
```

Expected: PASS. `Provider` now implements every method itself: the `var _ provider.Provider = (*Provider)(nil)` line
in `provider.go` fails to compile if anything is missing.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: treat every request as narrowed (the everything test lists the bucket); return the first error instead of
continuing (the one-failing-type test fails); decode discovered resources with AWS's names instead of snake_case by
skipping `cfn.SnakeCase` in `decodeObject`'s no-reference branch (the discovered-equals-snake-case test fails); drop
the page loop (the pages test finds 1).

```bash
git add internal/ccprov/discover.go internal/ccprov/discover_test.go
git commit -m "Discover resources across the configured regions

Asking about everything means the instance's discover_types or a default
set of common types, not all 1584. Types that cannot be listed are named
and skipped, and one failing type no longer hides the rest. Checked by
breaking the default set, error handling, what import sees and paging." -- \
  internal/ccprov/discover.go internal/ccprov/discover_test.go internal/ccprov/pending.go
```

---
### Task 13: The real catalog through infrata's host — `pkg/plugintest`

**Files:**
- Create: `internal/awstest/core.go`
- Test: `internal/awsprov/host_test.go`

**Interfaces:**
- Consumes: `plugintest.Open`, `(*Host).Configure`, `Definitions` (infrata v0.3.0); `catalog.Embedded` (Task 5);
  `awstest.Isolate` (Task 6); `awstest.FakeType` (Task 9); `ccfake` (Task 7); the whole provider (Tasks 9 to 12).
- Produces (used by Tasks 14 and 16):
  - `func awstest.TypeFor(t testing.TB, cat *catalog.Catalog, cfn string) *catalog.Type`
  - `func awstest.RegisterCore(t testing.TB, fake *ccfake.Server, cat *catalog.Catalog)` — VPC, Subnet, SecurityGroup,
    IAM Role and RDS DBInstance, with AWS's defaults for the properties the e2e suite leaves unset

Every earlier provider test used the hand-written test catalog and called the provider directly. This task runs the
generated catalog through the host adapter a real infrata uses, so every rule it enforces applies: undeclared
attributes refused, sensitivity forced from the schema, provenance overwritten. Type names are looked up by
CloudFormation name, except the three Task 5 already pins (`aws.vpc`, `aws.subnet`, `aws.securitygroup`).

- [ ] **Step 1: The core fake types**

`internal/awstest/core.go`:

```go
package awstest

import (
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
)

// TypeFor finds a catalog type by CloudFormation name, so a test does not depend on the name the generator assigned.
func TypeFor(t testing.TB, cat *catalog.Catalog, cfn string) *catalog.Type {
	t.Helper()
	for _, typ := range cat.Types {
		if typ.CFN == cfn {
			return typ
		}
	}
	t.Fatalf("the catalog has no %s", cfn)
	return nil
}

// RegisterCore describes the types the protocol, e2e and live suites use to the fake, with the defaults AWS chooses
// for properties those suites leave unset: the Optional+Computed attributes whose values must never plan a change.
func RegisterCore(t testing.TB, fake *ccfake.Server, cat *catalog.Catalog) {
	t.Helper()
	egress := []any{map[string]any{"IpProtocol": "-1", "CidrIp": "0.0.0.0/0"}}
	for cfn, c := range map[string]struct {
		prefix   string
		defaults map[string]any
	}{
		"AWS::EC2::VPC":           {"vpc-", map[string]any{"EnableDnsSupport": true, "EnableDnsHostnames": false, "InstanceTenancy": "default"}},
		"AWS::EC2::Subnet":        {"subnet-", map[string]any{"MapPublicIpOnLaunch": false, "AvailabilityZoneId": "use1-az1"}},
		"AWS::EC2::SecurityGroup": {"sg-", map[string]any{"SecurityGroupEgress": egress}},
		"AWS::IAM::Role":          {"role-", map[string]any{"MaxSessionDuration": 3600, "Path": "/"}},
		"AWS::RDS::DBInstance":    {"db-", map[string]any{"Engine": "postgres", "StorageEncrypted": false}},
	} {
		fake.Register(FakeType(TypeFor(t, cat, cfn), c.prefix, c.defaults))
	}
}
```

- [ ] **Step 2: Write the tests**

`internal/awsprov/host_test.go`:

```go
package awsprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/awstest"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// hosted is the plugin reached through infrata's host, configured against a fresh fake.
func hosted(t *testing.T, values map[string]value.Value) (provider.Provider, *ccfake.Server, *catalog.Catalog) {
	t.Helper()
	fake := ccfake.New()
	t.Cleanup(fake.Close)
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	awstest.RegisterCore(t, fake, cat)
	awstest.Isolate(t, fake.URL)
	prov, err := openHost(t).Configure(provider.Config{Instance: "main", Values: values})
	if err != nil {
		t.Fatal(err)
	}
	return prov, fake, cat
}

func want(t *testing.T, desired map[string]value.Value, st *resource.ResourceState) {
	t.Helper()
	for name, v := range desired {
		if got, ok := st.Attributes[name]; !ok || !got.Equal(v) {
			t.Errorf("%s = %v, want %v: the next plan would not be clean", name, got, v)
		}
	}
}

func m(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}

func TestTheLifecycleThroughTheHost(t *testing.T) {
	prov, fake, _ := hosted(t, nil)
	ctx := context.Background()
	attrs := map[string]value.Value{"region": s("us-east-1"), "CidrBlock": s("10.0.0.0/16"), "Tags": m("team", s("platform"))}

	st, err := prov.Create(ctx, &resource.DesiredResource{Address: address.Address{Name: "vpc"}, Type: "aws.vpc", Attrs: attrs})
	if err != nil {
		t.Fatal(err)
	}
	want(t, attrs, st)
	if v := st.Attributes["EnableDnsSupport"]; v.Raw != true || v.Source != value.SourceProvider {
		t.Errorf("EnableDnsSupport = %+v, want AWS's true, recorded as the provider's", v)
	}

	got, err := prov.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	want(t, attrs, got)

	changed := map[string]value.Value{}
	for k, v := range got.Attributes {
		changed[k] = v
	}
	changed["EnableDnsHostnames"] = value.Bool(true, value.SourceExplicit)
	updated, err := prov.Update(ctx, got, &resource.DesiredResource{Address: address.Address{Name: "vpc"}, Type: "aws.vpc", Attrs: changed})
	if err != nil || updated.Attributes["EnableDnsHostnames"].Raw != true {
		t.Fatalf("Update = %v, %v", updated, err)
	}

	if err := prov.Delete(ctx, updated); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("us-east-1", "AWS::EC2::VPC")); n != 0 {
		t.Errorf("%d VPCs remain", n)
	}
}

// TestASecretIsSensitiveWhereverItAppears. The host forces the flag from the schema; this proves the overlay's
// sensitive list reached the schema, and that the write-only value survives the wire into state.
func TestASecretIsSensitiveWhereverItAppears(t *testing.T) {
	prov, _, cat := hosted(t, nil)
	db := awstest.TypeFor(t, cat, "AWS::RDS::DBInstance")
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "db"}, Type: db.Name,
		Attrs: map[string]value.Value{"region": s("us-east-1"), "DBInstanceIdentifier": s("app"), "DBInstanceClass": s("db.t3.micro"),
			"MasterUsername": s("app"), "MasterUserPassword": s("hunter2")}})
	if err != nil {
		t.Fatal(err)
	}
	pw := st.Attributes["MasterUserPassword"]
	if pw.Raw != "hunter2" || !pw.Sensitive {
		t.Errorf("MasterUserPassword = %+v, want the carried value, marked sensitive", pw)
	}
}

func TestNestedSpellingSurvivesTheWire(t *testing.T) {
	prov, _, _ := hosted(t, nil)
	ingress := value.List([]value.Value{m("ip_protocol", s("tcp"), "from_port", value.Int(443, value.SourceExplicit),
		"to_port", value.Int(443, value.SourceExplicit), "cidr_ip", s("0.0.0.0/0"))}, value.SourceExplicit)
	attrs := map[string]value.Value{"region": s("us-east-1"), "GroupDescription": s("web"), "SecurityGroupIngress": ingress}
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "web"}, Type: "aws.securitygroup", Attrs: attrs})
	if err != nil {
		t.Fatal(err)
	}
	want(t, attrs, st)
}

func TestDiscoverAndImportThroughTheHost(t *testing.T) {
	prov, fake, cat := hosted(t, map[string]value.Value{"discover_regions": list("us-east-1")})
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-legacy", map[string]any{"VpcId": "vpc-legacy", "CidrBlock": "172.16.0.0/16"})
	role := awstest.TypeFor(t, cat, "AWS::IAM::Role")
	fake.Put("us-east-1", "AWS::IAM::Role", "ops", map[string]any{"RoleName": "ops", "Arn": "arn:aws:iam::123456789012:role/ops"})

	var everything []string
	for _, d := range openHost(t).Definitions() {
		everything = append(everything, d.Type)
	}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, r := range found {
		got = append(got, r.Type+" "+r.ProviderID)
	}
	for _, w := range []string{"aws.vpc us-east-1/vpc-legacy", role.Name + " global/ops"} {
		if !strings.Contains(strings.Join(got, "\n"), w) {
			t.Errorf("discovered %v, want %s", got, w)
		}
	}
	st, err := prov.Import(context.Background(), "aws.vpc", "us-east-1/vpc-legacy")
	if err != nil || st.Attributes["CidrBlock"].Raw != "172.16.0.0/16" {
		t.Fatalf("Import = %v, %v", st, err)
	}
}

// TestAnUnknownNestedKeyIsAnErrorBeforeAnyCall. The plan cannot see nested keys; the apply must say which one.
func TestAnUnknownNestedKeyIsAnErrorBeforeAnyCall(t *testing.T) {
	prov, fake, _ := hosted(t, nil)
	bad := value.List([]value.Value{m("ip_protocol", s("tcp"), "port", value.Int(443, value.SourceExplicit))}, value.SourceExplicit)
	_, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: "web"}, Type: "aws.securitygroup",
		Attrs: map[string]value.Value{"region": s("us-east-1"), "GroupDescription": s("web"), "SecurityGroupIngress": bad}})
	if err == nil || !strings.Contains(err.Error(), `"port"`) || !strings.Contains(err.Error(), "from_port") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("CreateResource") != 0 {
		t.Error("CreateResource was called")
	}
}
```

- [ ] **Step 3: Run**

```bash
go test -count=1 -run 'Host|Secret|Wire|BeforeAnyCall' ./internal/awsprov/ && go vet ./... && gofmt -l .
```

Expected: PASS. If the host refuses an attribute ("which its own schema does not declare"), the provider returned
something outside the catalog: fix the provider, never the test. If a real schema's shape differs from what a test
assumes (for example `MasterUsername` not existing), check `internal/cfn/testdata` and correct the test's attributes,
noting the correction in the Verification log.

- [ ] **Step 4: Sabotage, then commit**

Sabotages: remove `MasterUserPassword` from `gen/overlay.yaml`'s sensitive list and regenerate (the secret test fails;
restore and regenerate again); return an attribute named `region` for a global type from `stateFrom` (the host refuses
the discovered role); skip `matchProp` in `encode` and pass the key through (the wire test fails at AWS names, the
unknown-key test sends the create).

```bash
git add internal/awstest/core.go internal/awsprov/host_test.go
git commit -m "Test the generated catalog through infrata's own host

Create, read, update, delete, discover and import all go through the same
adapter a real infrata uses, against the fake, so the host's checks on
attributes, secrets and provenance apply. Checked by dropping a secret from
the overlay, leaking a region on a global type and skipping key
translation." -- internal/awstest/core.go internal/awsprov/host_test.go
```

---
### Task 14: A real infrata against the real binary and the fake (`-tags e2e`)

**Files:**
- Create: `e2e/e2e_test.go`, `e2e/testdata/basic/infra.yml`

**Interfaces:**
- Consumes: `ccfake`, `awstest.RegisterCore`, `awstest.TypeFor` (Tasks 7, 13); `catalog.Embedded` (Task 5); the
  binary `cmd/infrata-plugin-aws`; infrata v0.3.0's CLI (`plan --output`, `apply --auto-approve`, `discover`,
  `import <env> <type>.<id> --generate`, `explain`, `destroy`).
- Produces: `e2e/testdata/basic/infra.yml`, which the README quotes byte for byte (Task 17).

The harness is `infrata-provider-fake/e2e/e2e_test.go`'s: `TestMain` builds `infrata` from `$INFRATA_SRC` (default
`../../infrata`, skipping loudly with `E2E SKIPPED` when absent) and this plugin into a temporary plugin directory.
The difference: the cloud is an in-process `ccfake`, so each project owns one, and every command's environment points
the plugin at it. The plugin inherits infrata's environment.

Facts this relies on, checked in infrata v0.3.0: references canonicalise aliases (`${vpc.vpc_id}` becomes
`${vpc.VpcId}`, `internal/compiler/bind.go`); import selectors match `<type>.<provider id>` exactly against discovery,
never split at a dot (`internal/cli/import.go`, `narrowToSelectors`); `import --generate` writes Required and
Optional+ForceNew attributes only (`internal/generator/generate.go`); `explain` lists every spelling as
"also …" (`internal/cli/explain.go`); the summary lines are `Plan: %d to create, …`, `Apply complete: %d applied, %d failed, %d skipped.`
and `%d resource%s imported into %q.`.

- [ ] **Step 1: The fixture**

`e2e/testdata/basic/infra.yml`:

```yaml
project: demo

environments:
  dev: {}

variables:
  aws_region:
    type: string
    default: us-east-1

providers:
  - plugin: aws
    discover_regions: [us-east-1]
    discover_types: [aws.vpc, aws.subnet, aws.securitygroup]
    defaults:
      region: ${aws_region}

resources:
  vpc:
    type: aws.vpc
    cidr: 10.0.0.0/16
    tags:
      team: platform

  private_a:
    type: aws.subnet
    vpc_id: ${vpc.vpc_id}
    cidr: 10.0.1.0/24
    az: us-east-1a

  web:
    type: aws.securitygroup
    description: web servers
    vpc_id: ${vpc.vpc_id}
    ingress:
      - ip_protocol: tcp
        from_port: 443
        to_port: 443
        cidr_ip: 0.0.0.0/0
      - ip_protocol: tcp
        from_port: 80
        to_port: 80
        cidr_ip: 0.0.0.0/0
```

Every spelling here is an alias or snake_case form, never AWS's name: the suite is also the proof that J4, J5 and J8
work end to end.

- [ ] **Step 2: Write the suite**

`e2e/e2e_test.go`:

```go
//go:build e2e

// Package e2e runs a real infrata binary against a real infrata-plugin-aws binary and an in-process fake Cloud Control.
//
// Not part of `go test ./...`: it builds infrata from source. Run it with `go test -tags e2e -count=1 -v ./e2e/`.
// INFRATA_SRC points at the checkout; the default is the sibling ../infrata.
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

	"github.com/infrata/infrata-provider-aws/internal/awstest"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccfake"
)

var (
	infrataBin string
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
		src := os.Getenv("INFRATA_SRC")
		if src == "" {
			src = filepath.Join("..", "..", "infrata")
		}
		if _, err := os.Stat(filepath.Join(src, "cmd", "infrata")); err != nil {
			skipReason = fmt.Sprintf("no infrata checkout at %s (set INFRATA_SRC): %v", src, err)
			fmt.Fprintln(os.Stderr, "E2E SKIPPED: "+skipReason)
			return m.Run()
		}
		infrataBin = filepath.Join(tmp, "infrata")
		pluginDir = filepath.Join(tmp, "plugins")
		for _, b := range []struct{ dir, out, pkg string }{
			{src, infrataBin, "./cmd/infrata"},
			{"..", filepath.Join(pluginDir, "infrata-plugin-aws"), "./cmd/infrata-plugin-aws"},
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

// infrata runs the CLI in the project. Nothing from the developer's machine reaches the plugin: no plugins but ours,
// no AWS files, static test keys, no instance metadata, and Cloud Control and STS pointed at this project's fake.
func (e *env) infrata(t *testing.T, args ...string) (string, int) {
	t.Helper()
	home := t.TempDir()
	cmd := exec.Command(infrataBin, append(args, "--plugin-dir", pluginDir)...)
	cmd.Dir = e.dir
	cmd.Env = append(os.Environ(),
		"INFRATA_PLUGIN_PATH=", "HOME="+home,
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
		t.Fatalf("running infrata %v: %v", args, err)
	}
	return string(out), code
}

func (e *env) expect(t *testing.T, wantCode int, want []string, args ...string) string {
	t.Helper()
	out, code := e.infrata(t, args...)
	if code != wantCode {
		t.Fatalf("infrata %s: exit %d, want %d\n%s", strings.Join(args, " "), code, wantCode, out)
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("infrata %s: output lacks %q\n%s", strings.Join(args, " "), w, out)
		}
	}
	return out
}

// planOps runs `plan dev --output` and returns address → kind for every proposed operation, "noop" excluded.
func (e *env) planOps(t *testing.T) map[string]string {
	t.Helper()
	outPath := filepath.Join(t.TempDir(), "plan.json")
	e.infrata(t, "plan", "dev", "--output", outPath)
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

// drift changes a resource behind infrata's back.
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
			t.Fatalf("the subnet's VpcId = %v, want %s: ${vpc.vpc_id} did not resolve", subnet["VpcId"], vpcID)
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
	t.Run("a tag changed outside infrata is an update", func(t *testing.T) {
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
	t.Run("a create-only value changed outside infrata forces a replacement", func(t *testing.T) {
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) { p["CidrBlock"] = "10.50.0.0/16" })
		if kind := e.planOps(t)["vpc"]; kind != "replace" {
			t.Fatalf("vpc plans as %q, want replace", kind)
		}
		e.drift(t, "AWS::EC2::VPC", managed(t), func(p map[string]any) { p["CidrBlock"] = "10.0.0.0/16" })
	})
	t.Run("discover and import adopt a VPC infrata did not create", func(t *testing.T) {
		e.fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-legacy", map[string]any{
			"VpcId": "vpc-legacy", "CidrBlock": "172.16.0.0/16", "EnableDnsSupport": true, "EnableDnsHostnames": false, "InstanceTenancy": "default",
		})
		e.expect(t, 0, []string{"aws.vpc", "us-east-1/vpc-legacy"}, "discover")
		e.expect(t, 0, []string{"1 resource imported"}, "import", "dev", "aws.vpc.us-east-1/vpc-legacy", "--generate")
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after import --generate proposes %v", ops)
		}
	})
	t.Run("destroy removes everything infrata manages", func(t *testing.T) {
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

func TestAMisspelledKeyIsRefusedAgainstTheProvidersEntry(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "discover_regions:", "discover_region:", 1))
	e.expect(t, 1, []string{`unknown configuration "discover_region"`}, "plan", "dev")
}

func TestAnUnknownDiscoverTypeIsRefused(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "aws.securitygroup]", "aws.securitygroupp]", 1))
	e.expect(t, 1, []string{"aws.securitygroupp", "infrata explain"}, "plan", "dev")
}
```

- [ ] **Step 3: Run**

Run: `go test -tags e2e -count=1 -v ./e2e/`
Expected: PASS against the `../infrata` checkout (run `git -C ../infrata log -1 --format=%h` and record it). If a
quoted output line differs (an exit code, `1 resource imported`, the `explain` wording), read the actual output and
assert on it: the wording is infrata's. If a plan is not clean, that is a real finding: read `plan dev` to see which
attribute differs, and fix the provider or the overlay, never the assertion. If the cause is in infrata, stop and
report it to James with the evidence.

- [ ] **Step 4: Sabotage, then commit**

Sabotages: send tags as a map instead of a list (the tag subtest fails at AWS); skip the unordered reordering (the
first re-plan is not clean, because the fake reverses ingress); make `discoverTypes` return `nil` for a whole-catalog
request (`discover` no longer lists `vpc-legacy`); return AWS's names instead of snake_case with no reference (the
post-import plan is not clean).

```bash
git add e2e/e2e_test.go e2e/testdata/basic/infra.yml
git commit -m "Run a real infrata against the plugin and a fake Cloud Control

Plans, applies, re-plans clean, ignores values AWS chose, repairs drift,
patches in place, replaces on a create-only change, discovers, imports and
destroys, all written with friendly names and snake_case keys. Checked by
breaking tags, list order, discover_types and the import spelling." -- \
  e2e/e2e_test.go e2e/testdata/basic/infra.yml
```

---
### Task 15: Manifest, release gate, CI, and the two bump workflows

**Files:**
- Create: `plugin.yaml`, `internal/awsprov/manifest_test.go`, `scripts/release-check`, `scripts/build-release`,
  `scripts/scripts_test.go`
- Create: `.github/workflows/ci.yml`, `.github/workflows/release.yml`, `.github/workflows/bump-infrata.yml`,
  `.github/workflows/bump-schemas.yml`
- Modify: `e2e/e2e_test.go` (add the manifest test)

**Interfaces:**
- Consumes: `awsprov.Version`, `PluginName`; `pkg/pluginmanifest.Parse`, `(*Manifest).SpeaksProtocol/AllowsInfrata`,
  `pkg/pluginproto.Version` (infrata v0.3.0); `scripts/fetch-schemas`, `cmd/gen-cloudcontrol` (Tasks 1, 5).
- Produces: the release convention `infrata-plugin-aws_<version>_<goos>_<goarch>.tar.gz` (`.zip` on Windows).

Copied from `infrata-provider-fake` at `584802c` (its release after the protocol-2 amendment), changing only what names
the plugin and what differs in how infrata is depended on: the fake uses a `replace` and `scripts/ci-use-infrata-tag`;
this repository has no `replace` and builds pinned with `GOWORK=off` (CLAUDE.md), so that script and its three tests
are not copied.

- [ ] **Step 1: The manifest test first**

`internal/awsprov/manifest_test.go` — `infrata-provider-fake/internal/fake/manifest_test.go` with `package awsprov`
and nothing else changed (it reads `../../plugin.yaml` and compares against `PluginName` and exactly
`[pluginproto.Version]`).

Run: `go test -count=1 -run Manifest ./internal/awsprov/`
Expected: FAIL — `open ../../plugin.yaml: no such file or directory`.

- [ ] **Step 2: `plugin.yaml`**

```yaml
# plugin.yaml: what this plugin is, and what it works with. infrata PLAN.md §31.2.
# Read at a release TAG, never at the default branch, which describes unreleased code.
manifest: 1
name: aws
version: 0.1.0
# The protocol THIS RELEASE'S binary speaks: exactly the pluginproto.Version of the infrata go.mod requires. It changes
# in the same commit as that require (internal/awsprov/manifest_test.go and scripts/release-check refuse a mismatch).
protocol: [2]
platforms: [linux/amd64, linux/arm64, linux/arm, linux/386, darwin/amd64, darwin/arm64, windows/amd64, windows/arm64]
description: The AWS provider for infrata, serving every resource type AWS Cloud Control API supports.
# The oldest infrata release CI verifies this plugin against: go.mod's require. v0.3.0 is the first with
# Optional+Computed attributes and aliases, which every generated type uses.
infrata: ">= 0.3.0"
source: https://github.com/infrata/infrata-provider-aws
```

Run: `go test -count=1 -run Manifest ./internal/awsprov/`
Expected: PASS.

Append `TestTheInfrataUnderTestSpeaksTheManifestsProtocol` from `infrata-provider-fake/e2e/e2e_test.go` to
`e2e/e2e_test.go` unchanged, adding the import `github.com/infrata/infrata/pkg/pluginmanifest`. Run:
`go test -tags e2e -count=1 -run Manifest ./e2e/` — PASS.

- [ ] **Step 3: Scripts**

Copy `infrata-provider-fake/scripts/release-check` and `scripts/build-release` and substitute:

| In the fake's file | Here |
| --- | --- |
| `infrata-plugin-fake` | `infrata-plugin-aws` |
| `./cmd/infrata-plugin-fake` | `./cmd/infrata-plugin-aws` |
| `github.com/infrata/infrata-provider-fake/internal/fake.Version` | `github.com/infrata/infrata-provider-aws/internal/awsprov.Version` |
| `internal/fake.Version` (in the error message) | `internal/awsprov.Version` |
| `FAKE_VERSION_SYMBOL` | `PLUGIN_VERSION_SYMBOL` |
| `FAKE_MANIFEST` | `PLUGIN_MANIFEST` |

Copy `infrata-provider-fake/scripts/scripts_test.go` with the same substitutions, and delete from it
`moduleCopy`, `readRepoFile`, `TestGoSumCarriesWhatABuildWithoutTheReplaceNeeds`, `TestCheckRefusesAGoSumThatTidyStripped`
and `TestVersionRefusesARequireThatIsNotARelease` (they test `ci-use-infrata-tag`), then remove imports left unused.

```bash
chmod +x scripts/release-check scripts/build-release
go test -count=1 ./scripts/
scripts/release-check v0.1.0
```

Expected: PASS, including the doctored `protocol: [1]`, `[2, 1]` and `[99]` refusals and the unstamped-symbol refusal;
`release-check` prints `tag, plugin.yaml and binary all say 0.1.0, and speak protocol [2]`. The binary embeds the
catalog, so a cross-compile is slower than the fake's; `TestBuildReleaseNamesArchivesByTheInstallConvention` already
builds only two platforms.

- [ ] **Step 4: Workflows**

Every workflow fetches infrata as a private module: `GOPRIVATE` plus a git `insteadOf` carrying
`INFRATA_CHECKOUT_TOKEN`. The token goes through `env`, never a command line that is echoed.

`.github/workflows/ci.yml`:

```yaml
name: CI

# Two jobs, answering two questions.
#   pinned  Does this plugin work with the infrata go.mod requires? Builds with GOWORK=off against that release and runs
#           the e2e suite against a host built from the same tag. Blocking, and what a release runs.
#   main    Has infrata main broken us? Builds through a workspace against infrata's main. Advisory.

on:
  push:
    branches: [main]
  pull_request:
  workflow_call:
  workflow_dispatch:

permissions:
  contents: read

env:
  GOTOOLCHAIN: local
  GOPRIVATE: github.com/infrata/*

jobs:
  pinned:
    name: against the required infrata release
    runs-on: ubuntu-latest
    env:
      GOWORK: "off"
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - name: Let go fetch the private infrata module
        env:
          INFRATA_TOKEN: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
        run: git config --global url."https://x-access-token:${INFRATA_TOKEN}@github.com/infrata/".insteadOf "https://github.com/infrata/"
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: Read the infrata release go.mod requires
        id: infrata
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          v="$(go mod edit -json | jq -r '.Require[] | select(.Path == "github.com/infrata/infrata") | .Version')"
          [[ "$v" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "::error::go.mod requires infrata $v, which is not a release tag"; exit 1; }
          echo "version=$v" >> "$GITHUB_OUTPUT"
      # The e2e host, at ../infrata where the suite looks for it, built from the same tag the plugin compiles against.
      - uses: actions/checkout@v7
        with:
          repository: infrata/infrata
          ref: ${{ steps.infrata.outputs.version }}
          path: infrata
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - name: gofmt
        working-directory: infrata-provider-aws
        run: test -z "$(gofmt -l .)"
      - name: vet
        working-directory: infrata-provider-aws
        run: go vet ./... && go vet -tags e2e,live ./...
      - name: test
        working-directory: infrata-provider-aws
        run: go test -count=1 ./...
      - name: e2e against infrata ${{ steps.infrata.outputs.version }}
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          go test -tags e2e -count=1 -v ./e2e/ 2>&1 | tee "$RUNNER_TEMP/e2e.log"
          if grep -q 'E2E SKIPPED' "$RUNNER_TEMP/e2e.log"; then
            echo "::error::the e2e suite skipped instead of running against ../infrata"
            exit 1
          fi

  main:
    name: against infrata main (advisory)
    runs-on: ubuntu-latest
    continue-on-error: true
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - uses: actions/checkout@v7
        with:
          repository: infrata/infrata
          path: infrata
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: test and e2e against infrata main
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          go work init . ../infrata
          go test -count=1 ./...
          go test -tags e2e -count=1 -v ./e2e/ 2>&1 | tee "$RUNNER_TEMP/e2e.log"
          if grep -q 'E2E SKIPPED' "$RUNNER_TEMP/e2e.log"; then exit 1; fi
```

`.github/workflows/release.yml` — the fake plugin's file with these changes: the job `env` gains `GOPRIVATE:
github.com/infrata/*` and `GOWORK: "off"`; the `ci-use-infrata-tag use` step is replaced by `ci.yml`'s "Let go fetch
the private infrata module" step; everything else (the `ci` job with `secrets: inherit`, `release-check`,
`build-release`, `SHA256SUMS`, `gh release create --verify-tag`) unchanged.

`.github/workflows/bump-infrata.yml`. It runs the suites against a new infrata tag and opens the PR either way, saying
whether they passed; a bump that changes the plugin protocol fails the manifest test, and the PR body says what to do:

```yaml
name: Bump infrata

# infrata tags often until its first official release (James, 2026-09-13). This notices each tag and proposes it,
# having already run the suites against it.

on:
  schedule:
    - cron: "17 6 * * *"
  workflow_dispatch:

permissions:
  contents: write
  pull-requests: write

env:
  GOTOOLCHAIN: local
  GOPRIVATE: github.com/infrata/*
  GOWORK: "off"

jobs:
  bump:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - name: Let go fetch the private infrata module
        env:
          INFRATA_TOKEN: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
        run: git config --global url."https://x-access-token:${INFRATA_TOKEN}@github.com/infrata/".insteadOf "https://github.com/infrata/"
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: Upgrade to infrata's newest release
        id: bump
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          before="$(go list -m -f '{{.Version}}' github.com/infrata/infrata)"
          go get github.com/infrata/infrata@upgrade
          go mod tidy
          after="$(go list -m -f '{{.Version}}' github.com/infrata/infrata)"
          echo "before=$before" >> "$GITHUB_OUTPUT"
          echo "after=$after" >> "$GITHUB_OUTPUT"
          echo "changed=$([[ "$before" != "$after" ]] && echo true || echo false)" >> "$GITHUB_OUTPUT"
      - uses: actions/checkout@v7
        if: steps.bump.outputs.changed == 'true'
        with:
          repository: infrata/infrata
          path: infrata
          ref: ${{ steps.bump.outputs.after }}
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - name: Test against the new version
        id: test
        if: steps.bump.outputs.changed == 'true'
        continue-on-error: true
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          go vet ./... && go test -count=1 ./... && go test -tags e2e -count=1 ./e2e/
      - name: Open the PR
        if: steps.bump.outputs.changed == 'true'
        working-directory: infrata-provider-aws
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          set -euo pipefail
          branch="bump-infrata-${{ steps.bump.outputs.after }}"
          git config user.name "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git switch -c "$branch"
          git add go.mod go.sum
          git commit -m "Move to infrata ${{ steps.bump.outputs.after }}" -- go.mod go.sum
          git push origin "$branch"
          if [[ "${{ steps.test.outcome }}" == success ]]; then
            body="Upgraded from ${{ steps.bump.outputs.before }}. vet, the unit suite and the e2e suite passed against it."
          else
            body="Upgraded from ${{ steps.bump.outputs.before }}. The suites FAILED against it: see this workflow run. If the manifest test failed, the plugin protocol changed and plugin.yaml's protocol: must change in this PR."
          fi
          gh pr create --title "infrata ${{ steps.bump.outputs.after }}" --body "$body"
```

`.github/workflows/bump-schemas.yml`:

```yaml
name: Bump AWS schemas

# AWS publishes CloudFormation resource schemas continually. This regenerates the catalog from the newest bundle and
# proposes the diff. The name lock only grows (J3), so a regeneration never renames a type; a new warning in
# gen/warnings.txt is a reason to read the PR, not to merge it blind.

on:
  schedule:
    - cron: "41 5 * * 1"
  workflow_dispatch:

permissions:
  contents: write
  pull-requests: write

env:
  GOTOOLCHAIN: local
  GOPRIVATE: github.com/infrata/*
  GOWORK: "off"

jobs:
  regenerate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - name: Let go fetch the private infrata module
        env:
          INFRATA_TOKEN: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
        run: git config --global url."https://x-access-token:${INFRATA_TOKEN}@github.com/infrata/".insteadOf "https://github.com/infrata/"
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: Fetch the bundle and regenerate
        id: gen
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          scripts/fetch-schemas
          go run ./cmd/gen-cloudcontrol
          if git diff --quiet -- internal/catalog/catalog.json.gz gen/names.lock.json gen/warnings.txt; then
            echo "changed=false" >> "$GITHUB_OUTPUT"
          else
            echo "changed=true" >> "$GITHUB_OUTPUT"
          fi
      - name: Test the new catalog
        if: steps.gen.outputs.changed == 'true'
        working-directory: infrata-provider-aws
        run: go vet ./... && go test -count=1 ./...
      - name: Open the PR
        if: steps.gen.outputs.changed == 'true'
        working-directory: infrata-provider-aws
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          set -euo pipefail
          branch="bump-schemas-$(date +%Y-%m-%d)"
          git config user.name "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git switch -c "$branch"
          git add internal/catalog/catalog.json.gz gen/names.lock.json gen/warnings.txt
          git commit -m "Regenerate the catalog from this week's AWS schemas" -- internal/catalog/catalog.json.gz gen/names.lock.json gen/warnings.txt
          git push origin "$branch"
          gh pr create --title "AWS schemas $(date +%Y-%m-%d)" \
            --body "Regenerated from the newest us-east-1 schema bundle. The unit suite passed. Read the gen/warnings.txt diff for new secret-looking properties before merging."
```

Check locally what can be checked: `go vet -tags e2e,live ./...` (after Task 16), each workflow's shell by hand (the
`jq` line against `go.mod`, the `git diff --quiet` line after a regeneration). The workflows are only proven by the
first push, which is James's call, and the repository setting that lets Actions open pull requests is his to change.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: `name: aws2` in `plugin.yaml` (the manifest test); `protocol: [2, 1]` (the manifest test and
`release-check`); the ldflags symbol pointing at `internal/fake.Version` (the unstamped-binary test); `version: 0.1.1`
with tag `v0.1.0` (`release-check` refuses).

```bash
git add plugin.yaml internal/awsprov/manifest_test.go scripts/release-check scripts/build-release scripts/scripts_test.go \
  e2e/e2e_test.go .github/workflows/ci.yml .github/workflows/release.yml .github/workflows/bump-infrata.yml \
  .github/workflows/bump-schemas.yml
git commit -m "Add the manifest, release gate and CI workflows

Same three way version gate as the fake plugin, CI against the pinned
infrata release with an advisory run against main, and weekly jobs that
propose infrata and AWS schema bumps. Checked with a wrong name, a wrong
protocol, an unstamped binary and a mismatched version." -- \
  plugin.yaml internal/awsprov/manifest_test.go scripts/release-check scripts/build-release scripts/scripts_test.go \
  e2e/e2e_test.go .github/workflows/ci.yml .github/workflows/release.yml .github/workflows/bump-infrata.yml \
  .github/workflows/bump-schemas.yml
```

---
### Task 16: The live suite — real AWS, opt-in

**Files:**
- Create: `live/live_test.go`, `live/README.md`

**Interfaces:**
- Consumes: the plugin through `pkg/plugintest` (so the host's rules apply, as in production); `catalog.Embedded`,
  `awstest.TypeFor` (Tasks 5, 13).
- Produces: nothing other suites use.

Run by hand only, never in CI. It creates a VPC, a subnet and a security group (free) and an IAM role (free), all
tagged or named with the run. **Running it needs James's approval each time, and the `infrata` profile's root keys
replaced by a least-privilege IAM identity first** (see `live/README.md`). Writing and compiling it needs neither.

- [ ] **Step 1: Write the suite**

`live/live_test.go`:

```go
//go:build live

// Package live runs the plugin against a real AWS account. It is the only suite that sees real Cloud Control handlers,
// eventual consistency and IAM, and the only one that costs anything if it leaks, so it refuses to run unless told
// which account it may use and the credentials really are that account.
//
//	INFRATA_AWS_LIVE_PROFILE=infrata-live INFRATA_AWS_LIVE_ACCOUNT=111111111111 go test -tags live -count=1 -v -timeout 30m ./live/
package live

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/infrata/infrata-provider-aws/internal/awsprov"
	"github.com/infrata/infrata-provider-aws/internal/awstest"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/plugintest"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

const runTag = "infrata-live-run"

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }
func n(v int64) value.Value  { return value.Int(v, value.SourceExplicit) }
func m(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}
func l(items ...value.Value) value.Value { return value.List(items, value.SourceExplicit) }

// guard skips without the variables, and refuses when the profile is not the named account.
func guard(t *testing.T) (profile, region string) {
	t.Helper()
	profile, account := os.Getenv("INFRATA_AWS_LIVE_PROFILE"), os.Getenv("INFRATA_AWS_LIVE_ACCOUNT")
	if profile == "" || account == "" {
		t.Skip("set INFRATA_AWS_LIVE_PROFILE and INFRATA_AWS_LIVE_ACCOUNT to run against real AWS")
	}
	region = os.Getenv("INFRATA_AWS_LIVE_REGION")
	if region == "" {
		region = "us-east-1"
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		t.Fatal(err)
	}
	id, err := sts.NewFromConfig(cfg).GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := aws.ToString(id.Account); got != account {
		t.Fatalf("profile %q is account %s, not INFRATA_AWS_LIVE_ACCOUNT=%s: refusing to create anything", profile, got, account)
	}
	if strings.HasSuffix(aws.ToString(id.Arn), ":root") {
		t.Fatalf("profile %q holds root credentials: use a least-privilege IAM identity (live/README.md)", profile)
	}
	return profile, region
}

func configure(t *testing.T, profile, region string) (provider.Provider, *catalog.Catalog) {
	t.Helper()
	host, err := plugintest.Open(context.Background(), awsprov.NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = host.Close() })
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	role := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	prov, err := host.Configure(provider.Config{Instance: "live", Values: map[string]value.Value{
		"profile":          s(profile),
		"discover_regions": l(s(region)),
		"discover_types":   l(s("aws.vpc"), s("aws.subnet"), s("aws.securitygroup"), s(role)),
	}})
	if err != nil {
		t.Fatal(err)
	}
	return prov, cat
}

// converged fails unless every configured value reads back equal: otherwise the next plan is not clean.
func converged(t *testing.T, attrs map[string]value.Value, st *resource.ResourceState) {
	t.Helper()
	for name, v := range attrs {
		if got := st.Attributes[name]; !got.Equal(v) {
			t.Errorf("%s %s: %s reads back as %v, configured %v", st.Type, st.ProviderID, name, got, v)
		}
	}
}

func create(t *testing.T, prov provider.Provider, typ string, attrs map[string]value.Value) *resource.ResourceState {
	t.Helper()
	start := time.Now()
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: typ}, Type: typ, Attrs: attrs})
	if err != nil {
		t.Fatalf("create %s: %v", typ, err)
	}
	t.Logf("created %s %s in %s", typ, st.ProviderID, time.Since(start).Round(time.Millisecond))
	t.Cleanup(func() { _ = prov.Delete(context.Background(), st) })
	converged(t, attrs, st)
	read, err := prov.Read(context.Background(), st)
	if err != nil || read == nil {
		t.Fatalf("read %s straight after create = %v, %v", st.ProviderID, read, err)
	}
	converged(t, attrs, read)
	return read
}

func update(t *testing.T, prov provider.Provider, st *resource.ResourceState, changes map[string]value.Value) *resource.ResourceState {
	t.Helper()
	attrs := map[string]value.Value{}
	for k, v := range st.Attributes {
		attrs[k] = v
	}
	for k, v := range changes {
		attrs[k] = v
	}
	got, err := prov.Update(context.Background(), st, &resource.DesiredResource{Address: address.Address{Name: st.Type}, Type: st.Type, Attrs: attrs})
	if err != nil {
		t.Fatalf("update %s: %v", st.ProviderID, err)
	}
	converged(t, changes, got)
	return got
}

func TestTheLifecycleAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.99.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnet := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.99.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	ingress := l(m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")))
	sg := create(t, prov, "aws.securitygroup", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"GroupDescription": s("infrata live " + run), "SecurityGroupIngress": ingress, "Tags": tags})
	roleType := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	policy := m("Version", s("2012-10-17"), "Statement", l(m("Effect", s("Allow"), "Principal", m("Service", s("ec2.amazonaws.com")), "Action", s("sts:AssumeRole"))))
	role := create(t, prov, roleType, map[string]value.Value{"RoleName": s("infrata-live-" + run), "AssumeRolePolicyDocument": policy, "Tags": tags})

	vpc = update(t, prov, vpc, map[string]value.Value{"EnableDnsHostnames": value.Bool(true, value.SourceExplicit)})
	sg = update(t, prov, sg, map[string]value.Value{"SecurityGroupIngress": l(
		m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")),
		m("ip_protocol", s("tcp"), "from_port", n(80), "to_port", n(80), "cidr_ip", s("0.0.0.0/0")))})
	role = update(t, prov, role, map[string]value.Value{"MaxSessionDuration": n(7200)})

	var everything []string
	for _, typ := range cat.Types {
		everything = append(everything, typ.Name)
	}
	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range found {
		seen[r.ProviderID] = true
	}
	for _, st := range []*resource.ResourceState{vpc, subnet, sg, role} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}
	if imported, err := prov.Import(ctx, "aws.vpc", vpc.ProviderID); err != nil || !imported.Attributes["CidrBlock"].Equal(s("10.99.0.0/16")) {
		t.Errorf("import %s = %v, %v", vpc.ProviderID, imported, err)
	}

	for _, st := range []*resource.ResourceState{role, sg, subnet, vpc} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
}

// TestSweepLeftovers deletes what a crashed run left: anything tagged by this suite more than an hour ago.
func TestSweepLeftovers(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	cutoff := time.Now().Add(-time.Hour).Unix()
	order := []string{awstest.TypeFor(t, cat, "AWS::IAM::Role").Name, "aws.securitygroup", "aws.subnet", "aws.vpc"}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: order})
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range order {
		for _, r := range found {
			if r.Type != typ {
				continue
			}
			tags, _ := r.Attributes["Tags"].Raw.(map[string]value.Value)
			started, err := strconv.ParseInt(fmt.Sprint(tags[runTag].Raw), 10, 64)
			if err != nil || started >= cutoff {
				continue
			}
			t.Logf("deleting %s %s", r.Type, r.ProviderID)
			if err := prov.Delete(context.Background(), &resource.ResourceState{Type: r.Type, ProviderID: r.ProviderID, Attributes: r.Attributes}); err != nil {
				t.Error(err)
			}
		}
	}
}
```

`live/README.md` must say:
- what the suite needs: a dedicated account; a profile for an IAM user or role in it — never root keys, which the
  suite refuses — with a policy allowing `cloudcontrol:*`, `sts:GetCallerIdentity`, and the handler permissions for the
  four types (`ec2:CreateVpc`, `ec2:DeleteVpc`, `ec2:ModifyVpcAttribute`, `ec2:CreateSubnet`, `ec2:DeleteSubnet`,
  `ec2:CreateSecurityGroup`, `ec2:DeleteSecurityGroup`, `ec2:AuthorizeSecurityGroupIngress`,
  `ec2:RevokeSecurityGroupIngress`, `ec2:AuthorizeSecurityGroupEgress`, `ec2:RevokeSecurityGroupEgress`, `ec2:Describe*`,
  `ec2:CreateTags`, `ec2:DeleteTags`, `iam:CreateRole`, `iam:DeleteRole`, `iam:GetRole`, `iam:UpdateRole`,
  `iam:ListRoles`, `iam:TagRole`, `iam:UntagRole`, `iam:ListRolePolicies`, `iam:ListAttachedRolePolicies`), noting
  that a missing permission shows up as `AccessDenied` with the action named, and to add that action;
- the variables (`INFRATA_AWS_LIVE_PROFILE`, `INFRATA_AWS_LIVE_ACCOUNT`, optional `INFRATA_AWS_LIVE_REGION`), with no
  real account ID in the file;
- what it creates (VPC `10.99.0.0/16`, a subnet, a security group, a role named `infrata-live-<unix time>`, tagged
  `infrata-live-run`), that it takes minutes, that it refuses a mismatched account or root credentials, how to run
  `-run TestSweepLeftovers`, and that it never runs in CI.

- [ ] **Step 2: Check it compiles and skips**

Run: `go vet -tags live ./live/ && go test -tags live -count=1 -v ./live/`
Expected: both tests SKIP with "set INFRATA_AWS_LIVE_PROFILE…".

- [ ] **Step 3: The guard's sabotage, then commit**

Sabotage without touching AWS: set `INFRATA_AWS_LIVE_PROFILE=infrata INFRATA_AWS_LIVE_ACCOUNT=000000000000` and run
`-run TestSweepLeftovers`. `GetCallerIdentity` is read-only; the test must fail with "refusing to create anything"
before any Cloud Control call. With the current root keys it would instead hit the root refusal if the account
matched, which is also correct. Record which message appeared.

```bash
git add live/live_test.go live/README.md
git commit -m "Add an opt in live suite against real AWS

Creates, updates, discovers, imports and deletes a VPC, subnet, security
group and IAM role through the real host, and checks every value reads back
as configured. It refuses the wrong account and root credentials. Checked
the guard against a wrong account ID." -- live/live_test.go live/README.md
```

- [ ] **Step 4: The first real run (only with James's go-ahead)**

Not part of finishing the plan. When James approves and the least-privilege identity exists: run the suite, then
record in the Verification log how long each create took, whether any configured value did not read back equal (a
scalar AWS normalises goes in the overlay notes, spec §5), and any handler error code that differed from the fake's.

---
### Task 17: README, CLAUDE.md, and closing out

**Files:**
- Create: `README.md`, `internal/awsprov/readme_test.go`
- Modify: `CLAUDE.md`, `docs/specs/2026-09-14-generic-cloudcontrol-provider.md` (status line), this plan (Verification log)
- Vault: `projects/labs/infra-tool.md`, `projects/labs/daily/<date>.md`

- [ ] **Step 1: The README test first**

`internal/awsprov/readme_test.go`:

```go
package awsprov

import (
	"os"
	"strings"
	"testing"
)

// TestReadmeQuotesTheTestedExample. The README's infra.yml is the first thing anyone copies; the e2e suite runs that
// exact file, so the README must quote it byte for byte.
func TestReadmeQuotesTheTestedExample(t *testing.T) {
	readme, err := os.ReadFile("../../README.md")
	if err != nil {
		t.Fatal(err)
	}
	fixture, err := os.ReadFile("../../e2e/testdata/basic/infra.yml")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(readme), string(fixture)) {
		t.Error("README.md does not quote e2e/testdata/basic/infra.yml verbatim")
	}
	for _, want := range []string{
		"Cloud Control", "infrata explain", "discover_regions", "discover_types", "assume_role_arn", "profile", "--provider",
		"defaults:", "${aws_region}", "us-east-1/vpc-", "global/", "aws_region", "type_value", "cidr_block", "tags:",
		"gen/overlay.yaml", "gen/names.lock.json", "scripts/fetch-schemas", "go run ./cmd/gen-cloudcontrol",
		"-tags e2e", "-tags live", "go work init", "GOWORK=off", "plugin.yaml", "scripts/release-check", "0.0.0-dev",
	} {
		if !strings.Contains(string(readme), want) {
			t.Errorf("README.md never mentions %q", want)
		}
	}
}
```

Run: `go test -count=1 -run Readme ./internal/awsprov/` — FAIL (`open ../../README.md`).

- [ ] **Step 2: `README.md`**

It must cover, with real commands and output copied from the e2e run:
- what the plugin is in two sentences: every resource type AWS Cloud Control API supports, generated from AWS's
  published schemas; `infrata explain <type>` is the reference for any one of them;
- how to build it and where infrata finds it (`--plugin-dir`, `.infra/plugins/`, `~/.local/share/infrata/plugins/`,
  `$PATH`);
- type names: `aws.<resource>`, `aws.<service>.<resource>` when the short name is taken (J2), and that a name never
  changes once released (J3);
- attribute names: AWS's name in any case, its snake_case form (`cidr_block`), and a friendly alias where one is curated
  (`cidr`); plans show the alias or snake_case; keyword clashes show as `type_value`, `provider_value`,
  `lifecycle_value`; a type with its own `Region` property takes the plugin's region as `aws_region`; nested keys
  accept the same spellings; `tags:` is a map;
- values AWS chooses: an attribute left unset keeps AWS's value and never plans a change; removing an attribute from
  configuration keeps AWS's current value rather than resetting it;
- credentials: the default chain, `profile`, `assume_role_arn`; one instance per account; `import --provider <instance>`;
- regions: `defaults: {region: ${aws_region}}` with a `default:` and per-environment override, the per-resource
  override, the warning that changing the default region replaces every resource that inherits it; global types (IAM,
  Organizations, CloudFront, Route 53) have no region and IDs `global/<identifier>`;
- discovery: `discover_regions` (literal), `discover_types` and the default set when it is unset, why discovery does not
  scan every type, and that types needing a parent resource are not discovered;
- import IDs: `<region>/<identifier>` (`us-east-1/vpc-0abc`), `global/<identifier>`, composite identifiers joined with
  `|`, and `explain`'s "Import ID" section;
- the `e2e/testdata/basic/infra.yml` project quoted byte for byte;
- behaviour in user terms: creates and deletes wait for AWS to finish (a VPC takes about 12 seconds for real); a failed
  create that left a resource is recorded, not orphaned; a read may take up to ~4 seconds to report a resource gone;
- regenerating the catalog: `scripts/fetch-schemas`, `go run ./cmd/gen-cloudcontrol`, what `gen/overlay.yaml` and
  `gen/names.lock.json` are, and that the weekly workflow does this;
- building: `go work init . ../infrata` for local work, `GOWORK=off` plus credentials for the pinned build; the three
  suites (`go test ./...`, `-tags e2e`, `-tags live`); releases (`plugin.yaml`, `scripts/release-check`, versions are
  `0.0.0-dev` unless stamped); a pointer to the fake plugin's `AGENT.md` for plugin authors.

Run: `go test -count=1 -run Readme ./internal/awsprov/` — PASS. Sabotage: change one character in the README's quoted
fixture; the test fails; restore.

- [ ] **Step 3: `CLAUDE.md`**

Rewrite the sections that describe the handwritten slice; keep the rest (the contract table, the dependency model,
the test rules, "Changing infrata", commit discipline). Specifically:
- **Current state**: branch `generic-cloudcontrol`; built per `docs/plans/2026-09-14-generic-cloudcontrol-provider.md`
  from `docs/specs/2026-09-14-generic-cloudcontrol-provider.md`; the generator, catalog, provider, fake, e2e, release
  plumbing and live suite exist; the live suite has not run against real AWS until the Verification log says it has.
  The `first-slice` plan is history.
- **Stack**: dependencies become `aws-sdk-go-v2`, `config`, `credentials`, `service/sts`, `service/cloudcontrol`,
  `smithy-go`, `gopkg.in/yaml.v3`; `service/ec2` is gone. Add `scripts/fetch-schemas` and
  `go run ./cmd/gen-cloudcontrol` to the command block, and `go test -tags e2e` now needs no EC2 fake.
- **AWS-specific rules**, replacing the EC2 ones:
  - The catalog is generated. Never edit `internal/catalog/catalog.json.gz`; change `gen/overlay.yaml` or the generator
    and regenerate. `gen/names.lock.json` only grows: a name, once assigned, never moves (J3).
  - Every mutation carries a fresh `ClientToken`, so the SDK's retryer stays on; once a request is sent it is awaited
    under `context.WithoutCancel`, bounded by the handler timeout.
  - `Create` returns the state of a resource that exists even when AWS reported the request `FAILED`; the failure goes
    to stderr.
  - Nested values are reconciled (`internal/ccprov/reconcile.go`): outgoing keys to AWS's names, incoming to the
    reference spelling, without AWS-added keys, in the reference order for unordered lists; opaque shapes are copied
    exactly. A change there is a change to whether plans converge: the e2e suite is the check.
  - Update patches add or replace only, never remove, never a `Computed && !Optional` attribute.
  - Provider IDs are `<region>/<identifier>` or `global/<identifier>`, split at the first `/`.
  - `discover` with no types means `discover_types` or the catalog's default set, never the whole catalog.
  - Tests point `AWS_ENDPOINT_URL_CLOUDCONTROL` and `AWS_ENDPOINT_URL_STS` at `internal/ccfake` through
    `awstest.Isolate`; the real SDK is the fake's oracle.
  - Keep: never log credentials or sensitive values; `aws:` tags are never reported; the `discover` variable rule; no
    test reaches real AWS outside `-tags live`, and a live run needs James's approval and no root keys.

- [ ] **Step 4: Whole-suite verification**

```bash
git -C ../infrata status --short && git -C ../infrata log -1 --format=%h
gofmt -l . ; go vet ./... ; go vet -tags e2e,live ./...
go test -count=1 ./...
GOWORK=off GOPRIVATE='github.com/infrata/*' go test -count=1 ./...
go test -tags e2e -count=1 -v ./e2e/
go test -tags live -count=1 ./live/      # skips without the live variables
scripts/measure-load                     # the P7 gate again, with the finished provider
scripts/release-check v0.1.0
```

All green; keep the output for the close-out note and record the infrata commit the e2e suite ran against.

- [ ] **Step 5: Documentation**

- Spec: the status line becomes "Implemented on `generic-cloudcontrol` per the plan; live suite pending" (or its real
  state).
- This plan: every Verification log row still unconfirmed is confirmed or left marked with a reason; add the
  `measure-load` numbers from Task 6 and Step 4.
- Vault `projects/labs/infra-tool.md`: the AWS provider section's status, decisions made during execution, and
  follow-ups as tagged checkboxes (`- [ ] … #follow-up`): the first live run; replacing the root keys; any infrata
  finding. Today's `projects/labs/daily/<date>.md` gets one bullet linking the note.

- [ ] **Step 6: Commit, then ask**

```bash
git add README.md internal/awsprov/readme_test.go CLAUDE.md docs/specs/2026-09-14-generic-cloudcontrol-provider.md \
  docs/plans/2026-09-14-generic-cloudcontrol-provider.md
git commit -m "Document the generic provider

A README whose example is the one the e2e suite runs, CLAUDE.md rules for
the Cloud Control design, and the spec and plan brought up to date." -- \
  README.md internal/awsprov/readme_test.go CLAUDE.md docs/specs/2026-09-14-generic-cloudcontrol-provider.md \
  docs/plans/2026-09-14-generic-cloudcontrol-provider.md
```

Ask James before pushing `generic-cloudcontrol`, before merging it, and before any tag (a tag starts the release
workflow).

### Verification log additions: first live runs (2026-09-14)

| Claim | Checked against | Result |
| --- | --- | --- |
| Cloud Control's IAM actions are `cloudcontrol:*` (Task 16 README) | first live run, `AccessDeniedException` naming `cloudformation:CreateResource` | ✘ corrected: the actions are `cloudformation:CreateResource`, `GetResource`, `UpdateResource`, `DeleteResource`, `ListResources`, `GetResourceRequestStatus`, `ListResourceRequests`, `CancelResourceRequest` (`a1d6a32`) |
| Every configured value reads back equal from real AWS | live run 3 | ✘ `AWS::IAM::Role.AssumeRolePolicyDocument` (object-or-string, so a string attribute) came back as an object; fixed by keeping equivalent JSON text as written (`973d4a5`) |
| Create, update, discover, import and delete converge against real AWS for VPC, subnet, security group and IAM role | live run 4, after `973d4a5` | ✔ passed in 298 s; creates VPC 15.7 s, subnet 3.4 s, security group 10.0 s, role 24.7 s; one subnet delete took 2 m 52 s on AWS's side |
