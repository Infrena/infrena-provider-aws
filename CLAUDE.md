# CLAUDE.md

> Project notes (source of truth): Obsidian Vault/projects/labs/infra-tool.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this
repository.

## What this repository is

`infrena-provider-aws` is the **AWS provider for [infrena](https://github.com/infrena/infrena),
distributed as a plugin binary**: `infrena-plugin-aws`, module
`github.com/infrena/infrena-provider-aws`.

Infrena is a declarative infrastructure CLI. A provider plugin is a separate executable that infrena
launches as a child process and talks to over stdin/stdout in newline-delimited JSON. This repository
builds one, against real AWS APIs through the AWS SDK for Go v2, as one generic provider serving every
AWS resource type Cloud Control API supports.

It is the first plugin that touches a real cloud, so it is also where infrena's plugin contract meets
real credentials, many regions, pagination, throttling and eventual consistency for the first time.
When the contract is awkward here, that is a finding about infrena, not something to paper over.

## Current state

**Built on branch `generic-cloudcontrol`, per `docs/plans/2026-09-14-generic-cloudcontrol-provider.md`
and `docs/specs/2026-09-14-generic-cloudcontrol-provider.md`** (read the spec's §1 decisions J1–J10 and
§3 before changing generator or provider code, and the plan's Verification log before trusting a claim
about what's been checked). A build-time generator (`cmd/gen-cloudcontrol`) turns AWS's published
CloudFormation schemas plus `gen/overlay.yaml` and `gen/names.lock.json` into a committed, embedded
catalog; one provider (`internal/ccprov`) serves every catalog type through Cloud Control; tests run
against an in-process fake (`internal/ccfake`). This replaces the handwritten `aws.vpc`/`aws.subnet`
slice entirely — `docs/plans/2026-09-13-first-slice-vpc-subnet.md` is history. The project (module,
CLI, binary, env vars) was renamed from infrata to infrena partway through this work; every name below
is the current one. The e2e suite passes against a real infrena binary; the live suite against real AWS
first passed on 2026-09-14 (VPC, subnet, security group, IAM role). It still needs James's approval each
time; see `live/README.md` before attempting it.

## Where the contract lives

The protocol and the interfaces are defined in the infrena repository, not here:

| What | Where | Read it for |
| --- | --- | --- |
| `PLAN.md` §31.1 | infrena repo | the plugin design, and what the host refuses to trust a plugin with |
| `PLAN.md` §31.2 | infrena repo | the `plugin.yaml` manifest, and why it is read at the git tag |
| `PLAN.md` §12.1 | infrena repo | provider instances, `defaults:`, and variables in `providers:` |
| `pkg/provider` | infrena repo | `Plugin` and `Provider`, the two interfaces to implement |
| `pkg/schema`, `pkg/value`, `pkg/resource` | infrena repo | describing types, the value model, what CRUD receives and returns |
| `pkg/pluginsdk` | infrena repo | `Main(p)` — the whole of a plugin's `main()` |
| `pkg/plugintest` | infrena repo | the in-process harness the protocol tests use |
| `pkg/pluginmanifest`, `pkg/semver` | infrena repo | the parser `plugin.yaml` is checked with |
| `AGENT.md`, `docs/writing-a-provider.md` §14 | `infrena-provider-fake` repo | the authoring guide, and the earlier AWS worked example this plugin started from |
| the whole repo | `infrena-provider-fake` | scripts, release gate, e2e harness and protocol test patterns worth copying rather than reinventing |

The fake plugin's guide predates this generic design; where they disagree, this repository's actual
code wins.

## Stack and commands

Go **1.27.0** — infrena's own `go.mod` floor, so this module must declare it too. With
`GOTOOLCHAIN=auto` (the default) Go fetches a new enough toolchain.

**Dependencies:** `github.com/infrena/infrena` plus the AWS SDK for Go v2 (`aws-sdk-go-v2`, its
`config`, `credentials`, `service/sts`, `service/cloudcontrol`, `service/ec2`, and `smithy-go`), and
`gopkg.in/yaml.v3` for the generator's overlay. Cloud Control replaced `service/ec2` for CRUD; it came
back in 0.4.0 for one reason only, that Cloud Control's schemas carry no `IsDefault` on a VPC and no
`DefaultForAz` on a subnet, so `DescribeVpcs` and `DescribeSubnets` are the only authoritative way to
mark those as owned by AWS (`internal/ccprov/defaults.go`).

**What an AWS SDK service costs, measured when `service/ec2` was added (2026-09-15):** no new indirect
modules, +503 KiB of binary (+3.3%), cold build 8.7 s to 28.9 s, warm rebuild and plugin load time
unchanged. The cold build is the number that will creep if more services arrive; measure it again
before adding one, and add a dependency only with a reason written down.

**How infrena is depended on.** infrena is private. `go.mod` REQUIRES a real infrena version (currently
`v0.4.0`, the first release under the infrena name) and has **no `replace`**. Two ways to build:

- **Local, against your checkout:** a gitignored `go.work` (`go work init . ../infrena`) substitutes the
  sibling checkout's WORKING TREE, committed or not — the fast loop while both repos change daily. Run
  `git -C ../infrena status` before trusting a result.
- **Pinned, as CI and releases build:** `GOWORK=off GOPRIVATE='github.com/infrena/*' go test -count=1 ./...`.
  This fetches the required version over git, so it needs credentials for `github.com/infrena/infrena`
  (for example `gh auth setup-git`). `go get` and `go mod tidy` ignore `go.work` and always need them.

CI's blocking job builds pinned; a second, non-blocking job builds against infrena's `main` through a
workspace, as early warning. `bump-infrena.yml` opens a PR when infrena tags a newer release. All three
use the `INFRENA_CHECKOUT_TOKEN` secret.

```bash
go work init . ../infrena                    # once: build against the sibling checkout (go.work is gitignored)
scripts/fetch-schemas                        # downloads AWS's CloudFormation schema bundle into schemas/ (gitignored)
go run ./cmd/gen-cloudcontrol                 # regenerates internal/catalog/catalog.json.gz; commit the diff
go build ./cmd/infrena-plugin-aws            # build the plugin
go test -count=1 ./...                       # unit + ccfake + pkg/plugintest protocol tests; no AWS account
go test -tags e2e -count=1 -v ./e2e/         # a real infrena binary against this binary; no EC2 fake needed anymore
go test -tags live -count=1 -v ./live/       # REAL AWS: needs the INFRENA_AWS_LIVE_* variables; see live/README.md
go vet -tags e2e,live ./...
gofmt -l .
GOWORK=off GOPRIVATE='github.com/infrena/*' go test -count=1 ./...   # the pinned build CI blocks on
scripts/measure-load                          # the load-cost check; see the rule below
scripts/release-check vX.Y.Z                  # tag/manifest/binary agreement, before any real release
```

`-count=1` is mandatory: Go caches test results, and a cached pass hides a fixture edit.

## Rules for code in this repository

These are the ones that are easy to get wrong and expensive to get wrong.

### From the plugin contract (same as the fake plugin)

- **stdout is the protocol. Never print to it.** One stray `fmt.Println` corrupts the stream; the
  symptom is an unrelated parse error much later. Log to stderr. The SDK redirects `os.Stdout`, but a
  direct write to fd 1 still escapes. The AWS SDK's own logger must be pointed at stderr too if it is
  ever enabled.
- **A schema is data, and holds no functions.** `Default` is a datum of the attribute's `Kind`.
- **Never return `(nil, nil)` from `Create` or `Update`.** And never return an error from `Create`
  once AWS has created something: the host drops the result of a failed create
  (`internal/pluginhost/adapter.go`, `Create`), so the resource would exist untracked. Report the
  truthful state instead and let the next plan converge it.
- **Do not reimplement what the host enforces**: sensitivity from the schema, provenance, bookkeeping
  carry-forward, undeclared-attribute rejection. Return only `Type`, `ProviderID` and `Attributes`.
- **Cancellation is not a licence to abandon a mutation in flight.** Check `ctx` before sending a
  create, update or delete; once sent, finish reading the answer (`context.WithoutCancel` covers the
  mutating call and its await). Check `ctx` between pages of a paginated read or list.
- **`ClassifyError` is a pure function of the error.** The SDK asks any one configured instance to
  classify, not the one that failed. Anything unrecognised is `NotSafeToRetry`.

### AWS / Cloud Control specific

- **The catalog is generated. Never hand-edit `internal/catalog/catalog.json.gz`.** Change
  `gen/overlay.yaml` or the generator in `internal/gen`/`internal/cfn`, then regenerate and commit the
  diff like any other reviewed change. `gen/names.lock.json` only grows: once a CloudFormation type is
  assigned an infrena name, that name never moves to a different type, even if AWS removes the type.
- **Never log a credential, a signed request, or a `Sensitive` value.** stderr reaches CI logs and
  the tail of it is quoted in crash errors. Do not enable the SDK's request/signing debug logging.
- **Every mutation carries a fresh `ClientToken`.** Unlike the old EC2 calls, `CreateResource`,
  `UpdateResource` and `DeleteResource` all take one, so the SDK's retryer stays on safely; infrena's
  executor still decides retry from `ClassifyError`. Once a request is sent, it is awaited under
  `context.WithoutCancel`, bounded by the handler's `timeoutInMinutes` (default 120).
- **Classify Cloud Control errors by HTTP status/error code, not `ErrorFault()`.** Same reasoning as the
  old EC2 rule: Cloud Control's awsJson1_0 errors do not reliably set fault.
- **`Create` returns the state of a resource that exists even when AWS reports the request `FAILED`**,
  as long as an identifier exists and a read-back finds it; the failure itself goes to stderr, never to
  an error return (that would orphan the resource — the host drops a failed create's result).
- **Every regional type declares its region attribute as `Required` + `ForceNew`**; global types (IAM,
  Organizations, CloudFront, Route 53) do not, and are always called against Cloud Control in
  `us-east-1`. The default comes from the instance's `defaults: {region: …}`, never from the plugin's
  own configuration. A type that has its own AWS `Region` property exposes the plugin's region as
  `aws_region` instead, to avoid the name clash.
- **Provider IDs are `<region>/<identifier>` or `global/<identifier>`**, split at the first `/` only —
  identifiers can themselves contain `/` (ARNs), and a composite identifier joins its parts with `|`
  before that. `Import` refuses an ID whose scope doesn't match the type (regional vs. global) before
  any API call.
- **Nested values are reconciled (`internal/ccprov/reconcile.go`), not just top-level ones**: outgoing
  keys are translated to AWS's names, incoming keys back to the reference spelling, AWS-added keys are
  dropped, and unordered lists (`insertionOrder: false`) are reordered to match the reference. A shape
  with `patternProperties`, `additionalProperties` and no `properties`, `oneOf`/`anyOf`/`allOf`, or a
  multi-type `type` is opaque and copied exactly — translating or pruning its keys would corrupt user
  data. A change to reconciliation is a change to whether plans converge; the e2e suite is the check.
- **Update reads the resource fresh before building its patch.** infrena's host passes `Update` the last
  *persisted* state as current, not a freshly refreshed one (commit `c7ff98e`, found running the e2e
  suite: drifted tags that were corrected in configuration never got patched because the diff ran
  against stale state). The infrena team has confirmed this and plans to fix it on their side after the
  rename; once the Verification log says that's landed, drop the extra read here — don't do it
  preemptively.
- **JSON text in a string attribute is kept as written when AWS returns the same document as an object**
  (`sameJSONText`, commit `973d4a5`). Schemas type object-or-string properties such as IAM policy documents as
  strings; without this, spacing or key order plans a change forever. Found by the first live run.
- **Update's patch only adds or replaces, never removes**, and never touches a `readOnly` property:
  every settable property is `Optional`+`Computed`, so a property dropped from configuration produces
  no diff and keeps AWS's current value, per PLAN §14.1. For the same reason, omit `tags` entirely when
  a resource has no tags, and never report `aws:`-prefixed tags (reserved by AWS).
- **A `NotFound` on read is not proof of absence** for a resource created seconds ago (Cloud Control is
  eventually consistent). `Read` retries a bounded number of times (currently up to ~4 seconds total)
  before reporting `(nil, nil)`.
- **`discover` with no `discover_types` set means the catalog's curated default set, never every type
  the catalog knows.** Unfiltered discovery is roughly 1,500 `ListResources` calls per region per run;
  see `gen/overlay.yaml`'s `discover_default`. Types that only list under a parent resource are skipped
  and named once on stderr, not silently dropped.
- **The load-cost check is informational, not a hard gate.** `scripts/measure-load` measured the full
  catalog costing `infrena validate` 487–489 ms extra over a no-op plugin before the rename, and about
  437 ms extra after it — both comfortably inside what James considers fine. About 500 ms extra is
  acceptable; only flag it to James if a run shows the extra going over roughly 1 second, since that
  would suggest something changed for the worse, not the catalog's normal size.
- **Tests never reach real AWS** outside `-tags live`. Every test that constructs a configured instance
  isolates the SDK from the developer's machine (`AWS_CONFIG_FILE`, `AWS_SHARED_CREDENTIALS_FILE`,
  static keys, `AWS_EC2_METADATA_DISABLED=true`, and `AWS_ENDPOINT_URL_CLOUDCONTROL`/
  `AWS_ENDPOINT_URL_STS` pointing at `internal/ccfake`) via `internal/awstest`.
- **Fake the cloud, not the code.** `internal/ccfake` is an `httptest` server speaking Cloud Control's
  awsJson1_0 protocol, so every test goes through the real SDK's serialisation, error decoding and
  retryer. The real SDK is the oracle for the fake's wire format — for example, a registered error type
  like `InvalidRequestException` decodes its message from a field named exactly `Message`, not
  `message`; the fake sends `Message` to match (commit `72ae74b`).
- **A variable in `providers:` must resolve without an environment wherever `discover` is used.** infrena
  resolves `providers:` variables for every command given an environment, but `discover` takes none and
  refuses any value still unknown — in `defaults:` too, though it never uses them
  (`internal/cli/context.go`, `refuseUnresolvedInstances`). So examples write
  `defaults: {region: ${var.aws_region}}` (infrena 0.5.0 grammar: variables are `${var.x}`) with `aws_region` declared with a `default:` (overridden per
  environment), and keep `discover_regions` literal.

### Tests

- **A fixture must contradict its expected output.** A test whose fixture would satisfy the assertion
  by accident asserts nothing.
- **Sabotage every test**: break the code it covers so it still COMPILES and behaves wrongly, confirm
  the test fails, restore it (by re-applying the edit, not `git checkout --`, which can wipe
  uncommitted work). Record the sabotage in the commit message.

## Changing infrena

**Do not change the infrena repository from here.** A separate infrena session owns it. If something
there is wrong or missing, write it down with evidence (file, symbol, what was observed) in the plan's
Verification log or Findings and in the project note's follow-ups, and tell James. The
"Update sees persisted state, not a fresh read" behaviour above is exactly this kind of finding: this
repository works around it, and the workaround comes out once infrena's side changes.

The live suite against real AWS needs James's explicit approval each time. It runs as the `infrena-live`
IAM user (profile `infrena-live`), never root keys, which `live/live_test.go`'s `guard` refuses before
making any Cloud Control call. Cloud Control's IAM actions are named `cloudformation:CreateResource` and
so on, not `cloudcontrol:*`.

## Commit discipline

Stage explicit paths. `git add -A`, `git add .` and `git commit -am` are forbidden: they sweep up
scratch files, editor droppings and unrelated work, and a commit that contains something its message
does not mention is a commit nobody can review.

```bash
git add path/one path/two
git commit -m "..." -- path/one path/two
```

Write commit messages that explain why, not what. **Ask James before any `git push`** and before
creating a tag: a tag starts the release workflow.
