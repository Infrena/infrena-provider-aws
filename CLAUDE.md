# CLAUDE.md

> Project notes (source of truth): Obsidian Vault/projects/labs/infra-tool.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this
repository.

## What this repository is

`infrata-provider-aws` is the **AWS provider for [infrata](https://github.com/infrata/infrata),
distributed as a plugin binary**: `infrata-plugin-aws`, module
`github.com/infrata/infrata-provider-aws`.

Infrata is a declarative infrastructure CLI. A provider plugin is a separate executable that infrata
launches as a child process and talks to over stdin/stdout in newline-delimited JSON. This repository
builds one, against real AWS APIs through the AWS SDK for Go v2.

It is the first plugin that touches a real cloud, so it is also where infrata's plugin contract meets
real credentials, many regions, pagination, throttling and eventual consistency for the first time.
When the contract is awkward here, that is a finding about infrata, not something to paper over.

## Current state

**Under construction on branch `first-slice`: Task 1 of 12 done** (schemas, binary, module pinned at infrata `v0.2.0`). The design and the first vertical slice (`aws.vpc` + `aws.subnet`) are in
`docs/plans/2026-09-13-first-slice-vpc-subnet.md`. Read it — its Decisions, James's answers, Findings
and Verification log — before writing code or re-planning. Every factual claim it relies on was
checked against infrata's source, the AWS SDK's source or AWS documentation, and the log records the
claims that came back wrong.

## Where the contract lives

The protocol and the interfaces are defined in the infrata repository, not here:

| What | Where | Read it for |
| --- | --- | --- |
| `PLAN.md` §31.1 | infrata repo | the plugin design, and what the host refuses to trust a plugin with |
| `PLAN.md` §31.2 | infrata repo | the `plugin.yaml` manifest, and why it is read at the git tag |
| `PLAN.md` §12.1 | infrata repo | provider instances, `defaults:`, and variables in `providers:` (amended by infrata `5895f8a`) |
| `pkg/provider` | infrata repo | `Plugin` and `Provider`, the two interfaces to implement |
| `pkg/schema`, `pkg/value`, `pkg/resource` | infrata repo | describing types, the value model, what CRUD receives and returns |
| `pkg/pluginsdk` | infrata repo | `Main(p)` — the whole of a plugin's `main()` |
| `pkg/plugintest` | infrata repo | the in-process harness the protocol tests use |
| `pkg/pluginmanifest`, `pkg/semver` | infrata repo | the parser `plugin.yaml` is checked with |
| `AGENT.md`, `docs/writing-a-provider.md` §14 | `infrata-provider-fake` repo | the authoring guide, and the AWS worked example this plugin started from |
| the whole repo | `infrata-provider-fake` | the reference implementation: copy its patterns (scripts, release gate, e2e harness, protocol tests) rather than reinventing them |

The fake plugin's guide was written before this repository existed and is already partly stale; the
plan's "Findings" section lists where. When the guide and infrata's code disagree, the code wins.

## Stack and commands

Go **1.27.0** — infrata's own `go.mod` floor, so this module must declare it too. With
`GOTOOLCHAIN=auto` (the default) Go fetches a new enough toolchain.

**Dependencies:** `github.com/infrata/infrata` plus the AWS SDK for Go v2 (`aws-sdk-go-v2`, its
`config`, `credentials`, `service/ec2`, `service/sts`, and `smithy-go`). The stdlib-only rule was the
fake plugin's, not this repository's. Add another dependency only with a reason written in the plan.

**How infrata is depended on (James, 2026-09-13).** infrata is private. `go.mod` REQUIRES a real infrata
version (a tag, or a pseudo-version of a commit until a tag contains what this plugin needs) and has **no
`replace`**. Two ways to build:

- **Local, against your checkout:** a gitignored `go.work` (`go work init . ../infrata`) substitutes the
  sibling checkout's WORKING TREE, committed or not — the fast loop while both repos change daily. Run
  `git -C ../infrata status` before trusting a result.
- **Pinned, as CI and releases build:** `GOWORK=off GOPRIVATE='github.com/infrata/*' go test -count=1 ./...`.
  This fetches the required version over git, so it needs credentials for `github.com/infrata/infrata`
  (for example `gh auth setup-git`). `go get` and `go mod tidy` ignore `go.work` and always need them.

CI's blocking job builds pinned; a second, non-blocking job builds against infrata's `main` through a
workspace, as early warning. `bump-infrata.yml` opens a PR when infrata tags a newer release — infrata
versions move often until the first official release, and keeping up is deliberate. All three use the
`INFRATA_CHECKOUT_TOKEN` secret.

```bash
go work init . ../infrata                    # once: build against the sibling checkout (go.work is gitignored)
go build ./cmd/infrata-plugin-aws           # build the plugin
go test -count=1 ./...                       # unit + fake-EC2 + pkg/plugintest protocol tests; no AWS account
go test -tags e2e -count=1 -v ./e2e/         # a real infrata binary against this binary and the fake EC2 endpoint
go test -tags live -count=1 -v ./live/       # REAL AWS: needs the INFRATA_AWS_LIVE_* variables; see live/README.md
go vet -tags e2e,live ./...
gofmt -l .
GOWORK=off GOPRIVATE='github.com/infrata/*' go test -count=1 ./...   # the pinned build CI blocks on
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
  create or delete; once sent, finish reading the answer (the plan uses `context.WithoutCancel` for
  the mutating call itself). Check `ctx` between pages of a paginated read.
- **`ClassifyError` is a pure function of the error.** The SDK asks any one configured instance to
  classify, not the one that failed. Anything unrecognised is `NotSafeToRetry`.

### AWS-specific

- **Never log a credential, a signed request, or a `Sensitive` value.** stderr reaches CI logs and
  the tail of it is quoted in crash errors. Do not enable the SDK's request/signing debug logging.
- **Never let the SDK silently resend a create.** `CreateVpc` and `CreateSubnet` take no client
  token, so a create is called with `RetryMaxAttempts = 1` and infrata's executor decides from the
  classification. Reads, deletes and tag changes keep the SDK's retryer: infrata never retries reads,
  and those mutations are idempotent here.
- **Classify EC2 server faults by HTTP status, not `ErrorFault()`.** EC2's query-protocol errors are
  `smithy.GenericAPIError` with no fault set, so `ErrorFault() == smithy.FaultServer` never matches.
- **Every regional type declares `region` as `Required` + `ForceNew`**; global types (IAM, Route 53)
  do not. The default comes from the instance's `defaults: {region: …}`, never from the plugin's
  configuration and never from `DiscoverRequest.Region` (always `""`).
- **Provider IDs are `<region>/<aws id>`** (`us-east-1/vpc-0abc123`) everywhere: `Create`,
  `Discover`, `Import`. `Import` refuses an ID whose AWS prefix is the wrong type before any API call.
- **An attribute AWS always reports must be `Required`, `Computed`, or carry a `Default`.** infrata's
  planner treats an attribute present on the resource but absent from configuration as "removed from
  configuration" (`internal/planner/diff.go`, `diffAttributes`), so an optional attribute AWS fills in
  plans a change forever. For the same reason, omit `tags` entirely when a resource has no tags, and
  never report `aws:`-prefixed tags (reserved by AWS; users cannot set or delete them).
- **A `NotFound` on read is not proof of absence** for a resource created seconds ago (EC2 is
  eventually consistent). `Read` retries a bounded number of times before reporting `(nil, nil)`.
- **Tests never reach real AWS** outside `-tags live`. Every test that constructs a configured
  instance isolates the SDK from the developer's machine (`AWS_CONFIG_FILE`,
  `AWS_SHARED_CREDENTIALS_FILE`, static keys, `AWS_EC2_METADATA_DISABLED=true`, and
  `AWS_ENDPOINT_URL_EC2`/`AWS_ENDPOINT_URL_STS` pointing at the fake) — see `testenv_test.go` in the plan.
- **Fake the cloud, not the code.** The fake is an `httptest` server speaking EC2's query protocol,
  so every test goes through the real SDK's serialisation, error decoding and retryer. The real SDK is
  the oracle for the fake's wire format.
- **A variable in `providers:` must resolve without an environment wherever `discover` is used.** infrata
  resolves `providers:` variables for every command given an environment (`5895f8a`, `76c3f28`), but
  `discover` takes none and refuses any value still unknown — in `defaults:` too, though it never uses
  them (`internal/cli/context.go`, `refuseUnresolvedInstances`). So examples write
  `defaults: {region: ${aws_region}}` with `aws_region` declared with a `default:` (overridden per
  environment), and keep `discover_regions` literal (James, 2026-09-13).

### Tests

- **A fixture must contradict its expected output.** A test whose fixture would satisfy the assertion
  by accident asserts nothing.
- **Sabotage every test**: break the code it covers so it still COMPILES and behaves wrongly, confirm
  the test fails, restore it (by re-applying the edit, not `git checkout --`, which can wipe
  uncommitted work). Record the sabotage in the commit message.

## Changing infrata

**Do not change the infrata repository from here.** A separate infrata session owns it. If something
there is wrong or missing, write it down with evidence (file, symbol, what was observed) in the plan's
Findings and in the project note's follow-ups, and tell James.

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
