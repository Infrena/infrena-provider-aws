# infrata-plugin-aws, first vertical slice: `aws.vpc` + `aws.subnet` — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build `infrata-plugin-aws` serving `aws.vpc` and `aws.subnet` end to end — credentials, per-region
clients, CRUD, paginated discovery, import, error classification, eventual consistency — tested without an AWS
account, released through the same version gate as the fake plugin, with an opt-in suite against real AWS.

**Architecture:** One Go module. `cmd/infrata-plugin-aws` is a one-line `pluginsdk.Main`. `internal/awsprov`
holds the plugin: instance configuration, one EC2 client per region, one file per resource type, error
classification and provider-ID handling. `internal/ec2fake` is an `httptest` server speaking EC2's query protocol
(and STS `AssumeRole`), so every test drives the real SDK. Three suites: the plain suite (unit, fake-EC2 and
`pkg/plugintest` protocol tests), `-tags e2e` (a real infrata binary against this binary and the fake), and
`-tags live` (real AWS, opt-in, never in CI).

**Tech Stack:** Go 1.27.0; `github.com/infrata/infrata` required by version (local `go.work` over `../infrata`); AWS SDK for Go v2 —
`aws-sdk-go-v2 v1.47.0`, `config v1.33.4`, `credentials v1.20.4`, `service/ec2 v1.332.0`, `service/sts v1.50.0`,
`smithy-go v1.28.1` (the versions current on 2026-09-13, the ones every SDK claim below was checked against).

**Spec:** this document. The design is recorded here (Decisions, Verification log) rather than in a separate
spec, following `infrata-provider-fake/docs/plans/2026-09-13-port-fake-provider.md`. Contract: infrata `PLAN.md`
§31.1, §31.2, §12.1 (as amended by `5895f8a`); `infrata-provider-fake/AGENT.md` and
`docs/writing-a-provider.md` §14 as the starting position, corrected by the Findings below.

## Global Constraints

- Module `github.com/infrata/infrata-provider-aws`. Never under `github.com/infrata/infrata/`.
- `go 1.27.0`; `require github.com/infrata/infrata <real version>` and **no `replace`**; a gitignored `go.work` (`use . ../infrata`) for local work (D19).
- Plugin name `aws`; binary `infrata-plugin-aws`; every type prefixed `aws.`.
- `Version` defaults to `"0.0.0-dev"`, stamped only by `-ldflags -X` at release.
- Nothing writes to stdout. No credential, signed request or sensitive value is ever logged.
- `Create`/`Update` never return `(nil, nil)`; `Create` never returns an error once AWS created something.
- Do not reimplement host rules (sensitivity, provenance, bookkeeping, undeclared-attribute refusal).
- Regional types declare `region` `Required` + `ForceNew`. Provider IDs are `<region>/<aws id>`.
- A variable in `providers:` must resolve without an environment where `discover` is used: declare it with a `default:`; `discover_regions` stays literal (D20).
- No test reaches real AWS outside `-tags live`.
- Every test command uses `-count=1`. Every test is sabotage-verified; the sabotage goes in the commit message.
- Stage explicit paths only. Never `git add -A` / `git add .` / `git commit -am`. Ask James before any push or tag.
- **Prerequisite for Task 1:** git credentials that can fetch `github.com/infrata/infrata` (e.g. `gh auth setup-git`): `go get` and `go mod tidy` ignore `go.work`. On 2026-09-13 this machine had none.
- Do not modify `../infrata`. infrata defects go to Findings and the vault note's follow-ups.

---

## Decisions

| # | Decision | Why | Cost |
| --- | --- | --- | --- |
| D1 | **First slice is `aws.vpc` + `aws.subnet`** (James's suggestion, kept) | Covers every mechanism the plugin needs once: a requirement edge and a real reference (`vpc_id: ${vpc.id}`), a computed ID feeding another resource, ForceNew attributes, an in-place update that is not only tags (`map_public_ip_on_launch`), paginated discovery, regional IDs, and cross-resource eventual consistency (`CreateSubnet` right after `CreateVpc`). Both are free in AWS, so the live suite costs nothing. | Two types rather than one. Rejected: **`aws.s3_bucket`** — user-chosen name so no computed ID, a different protocol (REST-XML) with a region model of its own (`GetBucketLocation`), no requirement edge; worth its own slice later. **`aws.security_group`** — rules are nested lists whose diffing is the hard part; better as slice two, on top of a working VPC. **VPC alone** — loses the edge and the cross-resource consistency case, which are the parts most likely to find infrata problems. |
| D2 | **Credentials: `config.LoadDefaultConfig` once per instance in `New`, overrides as named keys** `profile` (→ `config.WithSharedConfigProfile`) and `assume_role_arn` (→ `aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(sts client, arn))`). Accepted keys: `profile`, `assume_role_arn`, `discover_regions`. Anything else is refused, naming what is accepted. | The standard chain means a user who can run the AWS CLI configures nothing twice. A misspelled key silently ignored would fall back to default credentials — another account. | No `external_id`, `role_session_name`, or static keys in config yet (YAGNI; static keys in YAML would be a secret in a file). Added when a user needs them. |
| D3 | **Credentials are resolved lazily**, not with `Retrieve` in `New`. `New` still fails fast on a named profile that does not exist (`config.SharedConfigProfileNotExistError`, a file read with no network). | Every compiling command, `validate` included, constructs instances; an SSO or assume-role `Retrieve` would make `validate` hit the network. | A missing credential surfaces at the first API call, not as a configuration error. The error names the instance and profile. |
| D4 | **STS client region falls back to `us-east-1`** when the loaded config has no region | Regions come from resources (`defaults: {region}`), so a user may have no default region at all, and an STS client with an empty region cannot resolve an endpoint. | Assume-role always goes to one region's STS. Acceptable for a slice; revisit if a partition other than `aws` is needed. |
| D5 | **One EC2 client per region**, built lazily from the instance's `aws.Config` and cached | The SDK documents per-call option overrides as concurrency-safe, but a client per region caches endpoint resolution and makes "which region" visible at one call site. | A mutex-guarded map. |
| D6 | **The SDK does not retry creates** (`CreateVpc`, `CreateSubnet` are called with `func(o *ec2.Options){ o.RetryMaxAttempts = 1 }`); reads, deletes, tag changes and `ModifySubnetAttribute` keep the SDK's standard retryer | Neither create takes a client token (checked: no `ClientToken` field on either input), and the SDK's standard retryer resends on connection errors and 5xx. A resent create after AWS acted makes a second, untracked VPC. Reads are never retried by infrata, so the SDK is the only retrier there. Deletes and tag/attribute changes are idempotent here (`NotFound` on delete is success). | infrata's executor sees a throttled create once per attempt rather than after SDK backoff. That is the point: one retry loop decides, with the classification. |
| D7 | **Classification** (`ClassifyError`, pure function): throttle code in `retry.DefaultThrottleErrorCodes` → `SafeToRetry`; a **dial** failure (`*net.OpError` with `Op == "dial"`, request never sent) → `SafeToRetry`; HTTP status ≥ 500 via `*smithyhttp.ResponseError` → `ConditionallyRetryable`; `context.DeadlineExceeded`, any other `net.Error`, `io.ErrUnexpectedEOF` → `ConditionallyRetryable`; `context.Canceled` and everything else → `NotSafeToRetry`. Checked in that order. | A throttle can carry a 5xx status (the SDK's own retryer lists 503 as retryable), so the throttle-code check must precede the 5xx check, or a throttled create would be classified as possibly-acted. EC2 errors carry no smithy fault (see Findings F3), so status code is the only server-fault signal. `*url.Error` implements `net.Error`, so the dial check must precede the generic one. | A dial-level failure after a DNS answer is still pre-send; that reasoning is Go's `net` semantics, not an SDK guarantee, and is pinned by a test against a closed listener. |
| D8 | **Error messages** carry the instance, operation, region/ID, AWS code and message, and request ID (`*awshttp.ResponseError.ServiceRequestID()`), and wrap the original error (`Unwrap`) so classification still sees it | A user reads it in a failed apply summary in CI with nothing else. | One small wrapper type. |
| D9 | **Eventual consistency.** `Create` builds its result from the create response (tags go in `TagSpecifications`, so no follow-up call). `Read` retries `InvalidVpcID.NotFound`/`InvalidSubnetID.NotFound` or an empty result with backoff (5 attempts, 250ms doubling, capped at 2s: ≤ 3.75s) before returning `(nil, nil)`. `CreateSubnet` retries `InvalidVpcID.NotFound` with the same patience (the API rejected the call, so nothing was created). `Import` reads once. | AWS documents that a just-created resource's ID "might not have propagated" and that a `NotFound` "does not mean the instance does not exist". A `Read` reporting gone plans a duplicate. The plugin is never sent `CreatedAt`, so it cannot tell a fresh resource from an old one; the bounded retry applies to every read. | A really-deleted resource takes up to 3.75s longer to be reported gone. The backoff is a field on `Provider` so tests use a zero-sleep schedule. |
| D10 | **A mutation that follows a create never turns the create into an error.** If `ModifySubnetAttribute` fails after `CreateSubnet` succeeded, `Create` returns the subnet's truthful state (`map_public_ip_on_launch: false`) and logs one line to stderr; the next plan proposes the update | infrata's adapter returns `nil` for an errored create, so the created subnet would be untracked. Truthful state converges; an orphan does not. | An apply can report success with one attribute not yet applied. Visible in the next plan. |
| D11 | **Once a mutating request is sent, cancellation does not abandon it**: `ctx.Err()` is checked before; the call itself runs on `context.WithoutCancel(ctx)` | §6 of the authoring guide: the host waits for the real answer; a cancelled HTTP request mid-create loses the answer, not the VPC. | A cancelled apply waits for up to one in-flight call. |
| D12 | **Schemas** (see Task 1): `aws.vpc` — `region`, `cidr` (Required, ForceNew), `tags` (map), computed `id`, `owner_id`, `is_default`. `aws.subnet` — `region`, `vpc_id`, `cidr`, `availability_zone` (all Required, ForceNew), `map_public_ip_on_launch` (bool, `Default: false`), `tags`, computed `id`, `arn`, `owner_id`. Subnet requires `aws.vpc`. | infrata's planner reports an attribute present on the resource but absent from configuration as "removed from configuration", so anything AWS always returns must be Required, Computed or defaulted: an optional `availability_zone` would plan a replacement forever. No `arn` on VPC: `types.Vpc` has no ARN field (Finding F4). `vpc_id` is the real dependency; the requirement is only a pre-flight hint (per-instance, not per-region, not from state — infrata `6e968a8`). | `availability_zone` can't be left to AWS; the user names it. VPC DNS attributes (`DescribeVpcAttribute` per attribute) are out of the slice. |
| D13 | **Tags**: `aws:`-prefixed keys are never reported and refused if configured; `tags` is omitted when a resource has no user tags; `Update` sends `CreateTags` for new/changed keys and `DeleteTags` for removed keys | AWS reserves `aws:` ("you can't edit or delete"), so reporting them makes every plan propose removing them. An empty map versus an absent attribute would plan "removed from configuration". Merging instead of removing is the fake plugin's old `Update` bug. | — |
| D14 | **Discover** scans `discover_regions` × requested types with the SDK paginators, checking `ctx` between pages; includes default VPCs and everything infrata did not create. No `discover_regions` → an error naming the key | `DiscoverRequest.Region` is always `""`. A silent empty survey is invisible; infrata's `Walk` reports a per-instance error and continues with other instances. | A resource in an unscanned region cannot be imported (infrata's import selects from discovery). Documented in README. |
| D15 | **Testing without AWS: one fake, at the HTTP level** — `internal/ec2fake`, an `httptest.Server` implementing the nine EC2 actions and STS `AssumeRole` this slice calls, partitioned by the SigV4 credential-scope region, with fault injection (status/code on the Nth call, drop the connection after applying), hide-for-N-describes, call counts, and the access key IDs it saw. Clients reach it through `AWS_ENDPOINT_URL_EC2`/`AWS_ENDPOINT_URL_STS`. | The same mechanism works in unit tests and for the e2e binary (the plugin inherits the env), so there is one fake, not an interface double plus a server. Going through the real SDK proves serialisation, error decoding, retry settings and the endpoint override. The SDK's own deserialiser is the oracle for the fake's XML. | Hand-written XML for nine actions. Rejected: a narrow client interface with an in-memory double (cannot serve the e2e binary; proves nothing about the SDK); LocalStack/moto (Docker or Python in the default suite, and a second implementation whose EC2 fidelity is not ours to fix). |
| D16 | **Real-AWS suite: yes, behind `//go:build live`, in `live/`, run by hand only.** Requires `INFRATA_AWS_LIVE_PROFILE` and `INFRATA_AWS_LIVE_ACCOUNT`; refuses to run unless STS `GetCallerIdentity` returns that account. Everything it creates is tagged `infrata-live-run=<run id>` and deleted in `t.Cleanup`; `TestSweepLeftovers` deletes tagged resources older than an hour. Not in CI. | The only suite that can catch real eventual-consistency and IAM behaviour. The account guard stops a contributor's default profile from being used by accident. VPCs and subnets cost nothing. | Needs a dedicated account (Q2). |
| D17 | **CI**: `.github/workflows/ci.yml` on push and pull request (gofmt, vet including `-tags e2e,live`, the plain suite, the e2e suite), plus `release.yml` copied from the fake plugin with the same three-way version gate | The fake plugin only tests at release; this plugin changes more often and a tag should not be the first time e2e runs in CI. Vetting the tagged files keeps `live/` compiling though it never runs. | One more workflow using `INFRATA_CHECKOUT_TOKEN`. |
| D18 | **Package name `internal/awsprov`** | `internal/aws` would shadow the SDK's `aws` package at every import site; `internal/provider` would shadow infrata's `pkg/provider`. | — |
| D19 | **infrata is required by version, not replaced** (James, 2026-09-13, answering Q3). `go.mod` requires the newest infrata tag that contains what this plugin relies on, or a pseudo-version of a commit until such a tag exists; no `replace`. Locally a gitignored `go.work` uses `../infrata`. CI: a **blocking `pinned` job** (`GOWORK=off`, `GOPRIVATE=github.com/infrata/*`, git credentials from `INFRATA_CHECKOUT_TOKEN`; e2e builds the infrata CLI from a checkout at the same version) and a **non-blocking `infrata-main` job** (workspace over infrata `main`). Releases build pinned. **`bump-infrata.yml`** runs daily: `go get github.com/infrata/infrata@upgrade`, tests, and opens a PR. | A `replace` makes every build compile infrata's working tree, so a green suite proves nothing about a released infrata, and the release ships an SDK nobody can name. `go.work` keeps the fast local loop without committing it. `@upgrade` never moves from a newer pseudo-version back to an older tag. infrata versions will move often until its first official release, and James wants to keep up, so the bump is automated rather than remembered. | Contributors need credentials for the private module even to `go mod tidy`. The only tag, `v0.1.0`, predates every infrata change this plugin relies on (F10), so Task 1 pins a pseudo-version until infrata tags again. A bump PR opened with `GITHUB_TOKEN` gets CI runs in an approval-required state, so the bump workflow runs the suite itself before opening it. Releases are no longer tested against infrata `main` (a change from the fake plugin): the non-blocking job covers that signal. |
| D20 | **Examples use `defaults: {region: ${aws_region}}`, with `aws_region` declared with a `default:`** and overridden per environment; `discover_regions` stays literal (James, 2026-09-13: "probably", answering Q1) | infrata resolves `providers:` variables for every command given an environment (`5895f8a`, `76c3f28`). `discover` takes none and refuses any value still unknown, `defaults:` included, though discovery never applies defaults (F9). A variable with a `default:` resolves without an environment, so `discover` keeps working. | A region set only per environment breaks `discover` until F9 is answered. The README's example and the e2e `basic` fixture change together (Task 9, Task 12). |

## James's answers (2026-09-13)

- **Q1 — `providers:` variables in examples: "Probably."** Adopted as D20. Task 9's `basic` fixture and the README use
  `defaults: {region: ${aws_region}}` with a `default:`; `TestProviderVariablesReachEnvironmentCommands` pins the
  environment-only case, including `discover`'s refusal (F9).
- **Q2 — live account: "Not yet."** The live suite (Task 11) is still built, compiles in CI (`go vet -tags live`),
  and skips without its variables. It has never run; its unconfirmed Verification log rows stay unconfirmed.
- **Q3 — require an infrata version in CI: "Yes, although the version will likely increase frequently and we need
  to keep up until we make our first official release."** Adopted as D19, with `bump-infrata.yml` for keeping up.
  Blocked on F10 for a tag; a pseudo-version works meanwhile.

## Findings (to report, not to fix here)

| # | Where | Finding | Evidence |
| --- | --- | --- | --- |
| F1 | infrata (resolved) | The "providers: literal" limit was lifted for every command with an environment; `discover` keeps it for per-environment values. The authoring docs still describe the old rule. | infrata `5895f8a`, `76c3f28`; `internal/cli/context.go` `registerStateInstances` → `compiler.VariableScope`, `refuseUnresolvedInstances`; PLAN §12.1 "AMENDED 2026-09-13". The fake-provider session is updating its docs. |
| F2 | fake repo docs (**fixed** in `b221ac8`, `b6b3dc5`) | `AGENT.md` §10 and `writing-a-provider.md` §8/§14 said state-only commands take literals only, requirements ignore instance, and an import selector silently picks one account. All three changed today; the fake-provider session updated the docs. F3–F6 were sent to that session, which is checking and fixing them. | infrata `5895f8a`, `76c3f28`, `6e968a8` (`internal/compiler/validate.go` `checkRequirements` keys by `held{instance, type}`), `a1efc85` (`import --provider`, ambiguous selectors refused, `internal/cli/import.go:311`). |
| F3 | fake repo docs, §14 error table (**fixed** upstream in `e62d157`) | "A server fault … `apiErr.ErrorFault() == smithy.FaultServer`" never matches EC2. EC2's query-protocol deserialiser returns `&smithy.GenericAPIError{Code, Message}` with `Fault` unset. Use `*smithyhttp.ResponseError.HTTPStatusCode() >= 500`. | `service/ec2@v1.332.0/deserializers.go` `awsEc2query_deserializeOpErrorCreateVpc`; `smithy-go@v1.28.1/errors.go` `GenericAPIError`; `transport/http/response.go:22` `HTTPStatusCode`. |
| F4 | fake repo docs, §14 VPC sketch (**fixed** upstream in `e62d157`) | The sketch declares a computed `arn` on `aws.vpc`; `types.Vpc` has no ARN field (`types.Subnet` has `SubnetArn`). It would have to be synthesised. | `service/ec2@v1.332.0/types/types.go`, `type Vpc struct`. |
| F5 | fake repo docs (**fixed** upstream in `e62d157`) | The guide documents "configuration sets it, state lacks it → change" but not the converse, which is what forces AWS-reported optional attributes to be Required, Computed or defaulted. | infrata `internal/planner/diff.go` `diffAttributes`, the `!inConfig` branch: "removed from configuration". |
| F6 | fake repo docs, §14 retries (**fixed** upstream in `e62d157`) | `o.RetryMaxAttempts = 1` per call is correct but has a subtlety worth one sentence: the SDK ignores a per-call value equal to the client's (`finalizeOperationRetryMaxAttempts`), which is harmless only because equal means already 1. | `service/ec2@v1.332.0/api_client.go:602-608`. |
| F7 | infrata, minor | `selectForImport`'s doc comment says a selector "splits at the LAST dot"; the code looks the whole selector up in a map. | `internal/cli/import.go`, `selectForImport`. |
| F9 | infrata, for the infrata session | `discover` refuses an unresolved value in an instance's `defaults:`, though discovery never applies defaults. `defaults: {region: ${aws_region}}` with a per-environment-only value therefore breaks `discover` for no benefit. Either intended (say so in §12.1) or refuse configuration only. | `internal/cli/context.go` `refuseUnresolvedInstances` loops `inst.Config` and `inst.Defaults`; `discover` reaches it through `discoveryRegistry(opts, "")` → `registerStateInstances`. |
| F10 | infrata, for James | The only infrata tag, `v0.1.0` (`cfe996f`), contains none of `5895f8a`, `76c3f28`, `6e968a8`, `a1efc85`. A plugin pinned to it would lose per-environment `providers:` variables on refresh/destroy/import, per-instance requirements and `import --provider`. D19 needs a newer tag. | `git merge-base --is-ancestor <commit> v0.1.0` false for all four; `git ls-remote --tags origin` lists only `v0.1.0`. |
| F8 | infrata, known | `DiscoverRequest.Region` / `pluginproto` `region` still exist and are never set. Already a follow-up; this plugin's model needs neither (evidence for removal). | `pkg/provider/provider.go:64-67`; `internal/discovery/walk.go` builds `{Types: ask}`. |

## Verification log

Every claim the design rests on, what it was checked against on 2026-09-13, and the result. infrata was at
`a1efc85` with a clean working tree at the last check. SDK source is the module cache for the versions in Tech Stack.

| Claim | Checked against | Result |
| --- | --- | --- |
| Plugin interfaces: `Plugin{Name, Definitions, New(Config)}`, `Provider{Read, Create, Update, Delete, Discover, Import, ClassifyError}`; `Config{Instance, Values, ProjectDir}` | infrata `pkg/provider/provider.go` | ✔ |
| `plugintest.Open(ctx, p, dir)`, `Host.Configure/Definitions/Version/Close` | infrata `pkg/plugintest/plugintest.go` | ✔ |
| `pluginsdk.Main(p provider.Plugin)`; `pluginproto.Version = 1` | `pkg/pluginsdk/serve.go:33`, `pkg/pluginproto/proto.go:29` | ✔ |
| Schema kinds: string/int/float/bool/list/map; `Validate` refuses Required+Computed, Computed+Default, Required+Default | `pkg/value/kind.go`, `pkg/schema/definition.go` | ✔ |
| A `providers:` entry's keys other than `plugin`/`name`/`default`/`defaults` reach `New` as `Config.Values` (there is no literal `config:` key) | `internal/config/providers.go`, `default:` branch | ✔ — the prompt's "`config: discover_regions`" means a top-level key in the entry |
| `defaults:` is applied at compile time (`applyInstanceDefaults` → `applyDefaults` → `checkRequired`), never sent to the plugin; wrong-kind defaults skipped | `internal/compiler/schema.go` | ✔ |
| State-only commands take literal `providers:` values only | `internal/cli/context.go` at `2e88f81` vs `5895f8a` | ✘ **changed** at `5895f8a`/`76c3f28` (F1) |
| Requirements satisfied by type anywhere in configuration | `internal/compiler/validate.go` | ✘ **changed** at `6e968a8`: per instance; still not per region, not from state (intended) |
| Import selector silently resolves to one instance on an ID collision | `internal/cli/import.go` | ✘ **changed** at `a1efc85`: refused, `--provider` narrows |
| Import selects only from what `Discover` returned | `internal/cli/import.go` `selectForImport` → `discovery.Walk` | ✔ |
| `Walk` asks every instance, errors are per instance, results sorted by type then ID | `internal/discovery/walk.go` | ✔ |
| An errored `Create` result is dropped by the adapter | `internal/pluginhost/adapter.go` `Create`: `if err != nil { return nil, err }` | ✔ — basis of D10 |
| Executor retries create/delete only on `SafeToRetry`, update also on `ConditionallyRetryable`; reads never go through the loop; 3 attempts, 500ms base, 10s cap, full jitter | `internal/executor/retry.go` `retryable`; `internal/cli/apply.go` `defaultRetryPolicy`; `VerbRead` has no caller | ✔ (`retry.go`'s own `defaultMax` is 30s; the CLI passes 10s) |
| SDK classifies on the plugin side, asking any one configured instance | `pkg/pluginsdk/serve.go` `plugin2Retryability` | ✔ |
| Planner: attribute on resource but not in config (non-computed) → "removed from configuration" | `internal/planner/diff.go` `diffAttributes` | ✔ — drives D12/D13 |
| `config.LoadDefaultConfig`, `WithSharedConfigProfile`, `WithRetryMaxAttempts`, `WithRetryer` exist | `config@v1.33.4/config.go:203`, `load_options.go` | ✔ |
| A missing named profile yields `SharedConfigProfileNotExistError` (value receiver) | `config/shared_config.go:1539,1550` | ✔ |
| `stscreds.NewAssumeRoleProvider(AssumeRoleAPIClient, roleARN, …)`; `aws.NewCredentialsCache` | `credentials@v1.20.4/stscreds/assume_role_provider.go:258`; `aws/credential_cache.go:69` | ✔ |
| EC2 reads `AWS_ENDPOINT_URL_EC2` (and `AWS_ENDPOINT_URL`); STS reads `AWS_ENDPOINT_URL_STS` | `service/ec2/endpoints.go:209-219`; `service/sts/endpoints.go:211` | ✔ |
| Env names `AWS_CONFIG_FILE`, `AWS_SHARED_CREDENTIALS_FILE`, `AWS_EC2_METADATA_DISABLED`, `AWS_MAX_ATTEMPTS`, `AWS_PROFILE`, `AWS_ACCESS_KEY_ID` | `config/env_config.go` | ✔ |
| `ec2.Options` has `BaseEndpoint`, `Region`, `RetryMaxAttempts`, `Retryer`, `Credentials` | `service/ec2/options.go` | ✔ |
| Per-call `RetryMaxAttempts` wraps the retryer with `retry.AddWithMaxAttempts`, unless equal to the client's | `service/ec2/api_client.go:602-608` | ✔ with subtlety (F6) |
| `retry.MaxAttemptsError` has `Unwrap`; standard retryer 3 attempts, retries 500/502/503/504, throttle codes, connection errors | `aws-sdk-go-v2@v1.47.0/aws/retry/errors.go`, `standard.go`, `retryable_error.go` | ✔ |
| `retry.DefaultThrottleErrorCodes` includes `Throttling`, `RequestLimitExceeded`, `EC2ThrottledException`, `TooManyRequestsException` | `aws/retry/standard.go:69-84` | ✔ |
| `aws.NopRetryer` exists (max attempts 1) | `aws/retryer.go:99` | ✔ (not used; per-call attempts chosen) |
| `smithy.APIError{ErrorCode, ErrorMessage, ErrorFault}`; EC2 errors set a server fault | `smithy-go/errors.go`; ec2 deserializers | ✘ **EC2 sets no fault** (F3) |
| `awshttp.ResponseError.ServiceRequestID()`; it embeds `*smithyhttp.ResponseError` with `HTTPStatusCode()`; EC2 client installs `ResponseErrorWrapper` | `aws/transport/http/response_error.go`; `smithy-go/transport/http/response.go:22`; `service/ec2/api_client.go:848` | ✔ |
| `CreateVpc` and `CreateSubnet` accept a client token | `service/ec2/api_op_CreateVpc.go`, `api_op_CreateSubnet.go` | ✘ **neither has `ClientToken`** (confirms D6) |
| Both creates accept `TagSpecifications`; `types.ResourceTypeVpc`, `ResourceTypeSubnet` | api_op files; `types/enums.go:10573,10587` | ✔ |
| Paginators `NewDescribeVpcsPaginator`, `NewDescribeSubnetsPaginator`; interfaces `DescribeVpcsAPIClient`, `DescribeSubnetsAPIClient`; waiters exist | `service/ec2/api_op_DescribeVpcs.go`, `api_op_DescribeSubnets.go` | ✔ |
| `InvalidVpcID.NotFound` is the not-found code on `DescribeVpcs` | `api_op_DescribeVpcs.go:515` (the SDK's own waiter matches it) | ✔; `InvalidSubnetID.NotFound` by analogy — confirmed only by the live suite (Task 11) |
| `types.Vpc` fields `VpcId, CidrBlock, OwnerId, IsDefault, State, Tags`; `types.Subnet` adds `SubnetArn, AvailabilityZone, MapPublicIpOnLaunch, VpcId` | `service/ec2/types/types.go` | ✔; no VPC ARN (F4) |
| EC2 is the query protocol: form-encoded POST; error body `Errors>Error>Code`, `Errors>Error>Message`, `RequestID` | `service/ec2/serializers.go` (`Content-Type: application/x-www-form-urlencoded`); `aws/protocol/ec2query/error_utils.go` | ✔ |
| Request field names: `CidrBlock`, `VpcId`, `AvailabilityZone`, `TagSpecification.N.ResourceType`, `TagSpecification.N.Tag.M.Key/Value`, `VpcId.N`, `SubnetId.N`, `NextToken`, `ResourceId.N`, `Tag.N.Key/Value`, `MapPublicIpOnLaunch.Value` | `service/ec2/serializers.go` | ✔ |
| Response elements: `vpc`, `vpcSet>item`, `subnet`, `subnetSet>item`, `nextToken`, `tagSet>item>key/value`; VPC `vpcId cidrBlock ownerId isDefault state`; subnet `subnetId vpcId cidrBlock availabilityZone mapPublicIpOnLaunch subnetArn ownerId state` | `service/ec2/deserializers.go` | ✔ (the root element name is not checked by the deserialiser; the fake uses the real ones anyway) |
| STS `AssumeRole` response wraps `AssumeRoleResult>Credentials>AccessKeyId/SecretAccessKey/SessionToken/Expiration` | `service/sts/deserializers.go:77,1938-2003` | ✔ element names; full shape confirmed by Task 3's SDK oracle test |
| Per-operation option overrides are concurrency-safe; unit testing via caller-defined interfaces | SDK developer guide via context7 (`using.md`, `unit-testing.md`) | ✔ |
| EC2 API is eventually consistent; `NotFound` shortly after create "does not mean the instance does not exist"; retry describe with exponential backoff | docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html | ✔ |
| `aws:` tag prefix reserved, can't be edited or deleted; 50 tags, key 128 / value 256 chars | docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html | ✔ |
| SDK modules' `go` directives (1.24) are below this module's 1.27.0 | module `go.mod`s | ✔ |
| A `go.work` using `../infrata` substitutes for a required infrata version with no network and no `go.sum` | scratch module requiring `v0.1.0`, `GOPROXY=off go build` in workspace mode | ✔ (and `-mod=mod` is refused in workspace mode) |
| `go mod tidy` honours `go.work` | same scratch module, `GOPROXY=off go mod tidy` | ✘ **it ignores the workspace and fetches** — needs credentials |
| This machine can fetch the private module by version | `GOPRIVATE=github.com/infrata/* go list -m -versions` | ✘ **no git credentials** (`could not read Username for 'https://github.com'`) |
| `@upgrade` never moves from a newer pseudo-version to an older tag | Go toolchain source `cmd/go/internal/modload/query.go:45,66,315` (1.24.13 checkout; 1.27 not re-read) | ✔ |
| A PR opened with `GITHUB_TOKEN` does not trigger CI | docs.github.com, "GITHUB_TOKEN" concept page | ✘ **partly**: `opened`/`synchronize`/`reopened` create runs in an approval-required state |
| `discover` refuses an unresolved `defaults:` value | `internal/cli/context.go` `refuseUnresolvedInstances`, `discoveryRegistry` | ✔ (F9) |
| A declared variable takes `type:` and `default:` | infrata `examples/shop/modules/app-stack/module.yml`; PLAN §12.1 "a declared `default:`" | ✔ for module inputs; confirmed for top-level `variables:` by Task 9 |
| infrata tags and what they contain | `git tag`, `git merge-base --is-ancestor`, `git ls-remote --tags origin` | only `v0.1.0`, lacking all four needed commits (F10) |

## File structure

```text
go.mod, go.sum                        module, infrata required by version (no replace), SDK requires
.gitignore                            /infrata-plugin-aws, /bin/, /dist/, /go.work, /go.work.sum
cmd/infrata-plugin-aws/main.go        pluginsdk.Main(awsprov.NewPlugin())
internal/awsprov/
  plugin.go        Plugin: Name, Version, Definitions, New (config → aws.Config → Provider)
  config.go        instanceConfig, parseConfig: accepted keys, refusal of unknown keys
  credentials.go   loadAWSConfig: LoadDefaultConfig, profile, assume role
  clients.go       per-region EC2 client cache; createOnce option
  definitions.go   aws.vpc, aws.subnet schemas
  ids.go           formatID, parseID (region/id, type prefix)
  errors.go        classify, apiFailure wrapper, hasCode
  patience.go      bounded backoff for NotFound (eventual consistency)
  tags.go          tags ↔ value.Map, aws: filter, diff, TagSpecifications
  values.go        small constructors for provider-sourced values
  provider.go      Provider: dispatch by type, ClassifyError, Discover, Import
  vpc.go           aws.vpc operations
  subnet.go        aws.subnet operations
  *_test.go        per file; testenv_test.go isolates the SDK; protocol_test.go via pkg/plugintest
internal/ec2fake/
  server.go        Server, state, routing, faults, region from SigV4 scope
  xml.go           response/error encoding
  server_test.go   the real SDK as the oracle for the wire format
e2e/e2e_test.go, e2e/testdata/{basic,variables}/infra.yml   -tags e2e
live/live_test.go, live/README.md                           -tags live
plugin.yaml, scripts/release-check, scripts/build-release, scripts/scripts_test.go
internal/awsprov/manifest_test.go, internal/awsprov/readme_test.go
.github/workflows/ci.yml, .github/workflows/release.yml, .github/workflows/bump-infrata.yml
README.md
```

---
### Task 1: Module, schemas, and a binary the host accepts

**Files:**
- Create: `go.mod`, `go.sum`, `.gitignore`, `cmd/infrata-plugin-aws/main.go`
- Create: `internal/awsprov/plugin.go`, `internal/awsprov/definitions.go`, `internal/awsprov/values.go`
- Test: `internal/awsprov/definitions_test.go`, `internal/awsprov/protocol_test.go`

**Interfaces:**
- Produces: `const PluginName = "aws"`; `var Version = "0.0.0-dev"`; `type Plugin struct{}`; `func NewPlugin() *Plugin`;
  `func definitions() []*schema.ResourceDefinition`; `const typeVPC = "aws.vpc"`, `typeSubnet = "aws.subnet"`;
  `func str(s string) value.Value`, `func boolean(b bool) value.Value` (provider-sourced constructors).
  In this task `Plugin.New` returns an error `"not built yet"`; Task 2 replaces it.

- [ ] **Step 1: Create the module**

Prerequisite: git credentials for `github.com/infrata/infrata` (`gh auth setup-git`, or an SSH `insteadOf`).
Check with `GOWORK=off GOPRIVATE='github.com/infrata/*' go list -m -versions github.com/infrata/infrata`.

```bash
cd /home/james/projects/infrata-provider-aws
export GOPRIVATE='github.com/infrata/*'
go mod init github.com/infrata/infrata-provider-aws
go mod edit -go=1.27.0
# The newest infrata TAG containing a1efc85 (import --provider). Until one exists (F10), pin that commit:
# `go get` turns it into a pseudo-version. Check first: git -C ../infrata tag --contains a1efc85
GOWORK=off go get github.com/infrata/infrata@a1efc85
go get github.com/aws/aws-sdk-go-v2@v1.47.0 github.com/aws/aws-sdk-go-v2/config@v1.33.4 \
  github.com/aws/aws-sdk-go-v2/credentials@v1.20.4 github.com/aws/aws-sdk-go-v2/service/ec2@v1.332.0 \
  github.com/aws/aws-sdk-go-v2/service/sts@v1.50.0 github.com/aws/smithy-go@v1.28.1
go work init . ../infrata                 # local builds use the sibling checkout; go.work is never committed
git -C ../infrata status --short          # must be empty, or record what the build is compiling against
```

`go.mod` gets no `replace`. Above the infrata `require`, add:

```
// infrata is private: fetch it with GOPRIVATE=github.com/infrata/* and git credentials. Local work
// builds against ../infrata through a gitignored go.work; CI builds this exact version (GOWORK=off).
```

`.gitignore`:

```
/infrata-plugin-aws
/bin/
/dist/
/go.work
/go.work.sum
```

- [ ] **Step 2: Write the failing schema test**

`internal/awsprov/definitions_test.go`:

```go
package awsprov

import (
	"strings"
	"testing"
)

// TestEveryTypeIsValidAndPrefixed. The host refuses the whole plugin for either failure, naming the
// plugin rather than the definition, so catch it here first.
func TestEveryTypeIsValidAndPrefixed(t *testing.T) {
	defs := definitions()
	if len(defs) != 2 {
		t.Fatalf("got %d definitions, want 2", len(defs))
	}
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Errorf("%s: %v", d.Type, err)
		}
		if !strings.HasPrefix(d.Type, PluginName+".") {
			t.Errorf("%s is not prefixed %s.", d.Type, PluginName)
		}
	}
}

// TestRegionalAttributesAreRequiredAndForceNew. A resource cannot move between regions or VPCs, and
// every attribute AWS always reports must be Required, Computed or defaulted, or infrata plans
// "removed from configuration" forever (internal/planner/diff.go, diffAttributes).
func TestRegionalAttributesAreRequiredAndForceNew(t *testing.T) {
	want := map[string][]string{
		typeVPC:    {"cidr", "region"},
		typeSubnet: {"availability_zone", "cidr", "region", "vpc_id"},
	}
	for _, d := range definitions() {
		got := d.ForceNewAttributes()
		if strings.Join(got, ",") != strings.Join(want[d.Type], ",") {
			t.Errorf("%s ForceNew = %v, want %v", d.Type, got, want[d.Type])
		}
		for _, name := range want[d.Type] {
			if a, _ := d.Attribute(name); !a.Required {
				t.Errorf("%s.%s is not Required", d.Type, name)
			}
		}
		for name, a := range d.Attributes {
			if !a.Required && !a.Computed && a.Default == nil && name != "tags" {
				t.Errorf("%s.%s is optional with no default: AWS reports it, so every plan would show it removed", d.Type, name)
			}
		}
	}
}

func TestASubnetRequiresAVPC(t *testing.T) {
	for _, d := range definitions() {
		if d.Type != typeSubnet {
			continue
		}
		if len(d.Requirements) != 1 || d.Requirements[0].Types[0] != typeVPC || d.Requirements[0].Optional {
			t.Fatalf("aws.subnet requirements = %+v, want one mandatory aws.vpc", d.Requirements)
		}
	}
}
```

- [ ] **Step 3: Run it to see it fail**

Run: `go test -count=1 ./internal/awsprov/`
Expected: FAIL — `undefined: definitions`.

- [ ] **Step 4: Write the schemas, the value helpers and the plugin skeleton**

`internal/awsprov/values.go`:

```go
package awsprov

import "github.com/infrata/infrata/pkg/value"

// str and boolean build values as a provider reports them. Sensitivity and provenance are NOT set
// here beyond the source: the host forces both from the schema, and a plugin that also did it would
// have tests that pass when the host is broken.
func str(s string) value.Value    { return value.String(s, value.SourceProvider) }
func boolean(b bool) value.Value  { return value.Bool(b, value.SourceProvider) }
```

`internal/awsprov/definitions.go`:

```go
package awsprov

import (
	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

const (
	typeVPC    = "aws.vpc"
	typeSubnet = "aws.subnet"
)

// regionAttr is declared by every regional type. ForceNew because an AWS resource cannot move
// between regions; the default comes from the instance's `defaults: {region: …}`, applied by
// infrata at compile time and never sent to this plugin.
var regionAttr = schema.Attribute{Kind: value.KindString, Required: true, ForceNew: true,
	Description: "AWS region, e.g. us-east-1. Usually set once with the provider's defaults: {region: …}"}

var tagsAttr = schema.Attribute{Kind: value.KindMap,
	Description: "Tags. Keys starting aws: are reserved by AWS and are neither reported nor accepted"}

func definitions() []*schema.ResourceDefinition {
	return []*schema.ResourceDefinition{
		{
			Type:        typeVPC,
			Description: "An Amazon VPC.",
			Attributes: map[string]schema.Attribute{
				"region": regionAttr,
				// The primary CIDR block cannot be changed or disassociated after creation.
				"cidr":       {Kind: value.KindString, Required: true, ForceNew: true, Description: "Primary IPv4 CIDR block"},
				"tags":       tagsAttr,
				"id":         {Kind: value.KindString, Computed: true, Description: "VPC ID, e.g. vpc-0abc123"},
				"owner_id":   {Kind: value.KindString, Computed: true, Description: "Owning AWS account ID"},
				"is_default": {Kind: value.KindBool, Computed: true, Description: "Whether this is the region's default VPC"},
			},
			Capabilities: schema.Capabilities{Create: true, Read: true, Update: true, Delete: true, Import: true},
			ImportID:     schema.ImportSpec{Description: "<region>/<vpc id>, e.g. us-east-1/vpc-0abc123"},
		},
		{
			Type:        typeSubnet,
			Description: "A subnet inside an Amazon VPC.",
			Attributes: map[string]schema.Attribute{
				"region": regionAttr,
				"vpc_id": {Kind: value.KindString, Required: true, ForceNew: true,
					Description: "VPC to create the subnet in, usually ${<vpc>.id}"},
				"cidr": {Kind: value.KindString, Required: true, ForceNew: true, Description: "IPv4 CIDR block"},
				// Required, not optional-and-chosen-by-AWS: AWS always reports it, so an omitted
				// availability zone would plan a replacement on every run.
				"availability_zone": {Kind: value.KindString, Required: true, ForceNew: true,
					Description: "Availability zone, e.g. us-east-1a"},
				"map_public_ip_on_launch": {Kind: value.KindBool, Default: false,
					Description: "Give instances launched here a public IPv4 address"},
				"tags":     tagsAttr,
				"id":       {Kind: value.KindString, Computed: true, Description: "Subnet ID, e.g. subnet-0abc123"},
				"arn":      {Kind: value.KindString, Computed: true, Description: "Subnet ARN"},
				"owner_id": {Kind: value.KindString, Computed: true, Description: "Owning AWS account ID"},
			},
			// A pre-flight hint only: satisfied by any aws.vpc in the same provider instance,
			// whatever its region, and never by state. The real dependency is vpc_id.
			Requirements: []schema.Requirement{{
				Name: "vpc", Types: []string{typeVPC},
				Description: "A subnet must be created inside a VPC",
			}},
			Capabilities: schema.Capabilities{Create: true, Read: true, Update: true, Delete: true, Import: true},
			ImportID:     schema.ImportSpec{Description: "<region>/<subnet id>, e.g. us-east-1/subnet-0abc123"},
		},
	}
}
```

`internal/awsprov/plugin.go`:

```go
// Package awsprov is infrata's AWS provider plugin.
package awsprov

import (
	"errors"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/schema"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = "aws"

// Version is reported in the handshake. "0.0.0-dev" in every build a release did not stamp, so a
// broken -ldflags path cannot pass the release gate by coincidence. scripts/build-release stamps it.
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

// Definitions need no configuration and make no network call: every command loads them.
func (pl *Plugin) Definitions() []*schema.ResourceDefinition { return definitions() }

// New builds one configured instance. Replaced in Task 2.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	return nil, errors.New("the aws plugin cannot configure instances yet")
}
```

`cmd/infrata-plugin-aws/main.go`:

```go
// Command infrata-plugin-aws is infrata's AWS provider. Infrata runs it; you do not.
package main

import (
	"github.com/infrata/infrata-provider-aws/internal/awsprov"
	"github.com/infrata/infrata/pkg/pluginsdk"
)

func main() { pluginsdk.Main(awsprov.NewPlugin()) }
```

- [ ] **Step 5: Add the protocol test for schema loading**

`internal/awsprov/protocol_test.go`:

```go
package awsprov

import (
	"context"
	"testing"

	"github.com/infrata/infrata/pkg/plugintest"
)

// openHost connects this plugin to infrata's own host over an in-memory pipe: schema validation,
// the prefix rule and reserved names all apply, exactly as for a subprocess.
func openHost(t *testing.T) *plugintest.Host {
	t.Helper()
	host, err := plugintest.Open(context.Background(), NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatalf("the host refused this plugin's schemas: %v", err)
	}
	t.Cleanup(func() { _ = host.Close() })
	return host
}

// TestSchemasLoadThroughTheHost. The subnet's map_public_ip_on_launch default is the one datum
// that must survive JSON as a bool.
func TestSchemasLoadThroughTheHost(t *testing.T) {
	host := openHost(t)
	if got := host.Version(); got != Version {
		t.Errorf("handshake version = %q, want %q", got, Version)
	}
	var subnet bool
	for _, d := range host.Definitions() {
		if d.Type != typeSubnet {
			continue
		}
		subnet = true
		a, _ := d.Attribute("map_public_ip_on_launch")
		if b, ok := a.Default.(bool); !ok || b {
			t.Errorf("map_public_ip_on_launch default after the wire = %#v (%T), want false", a.Default, a.Default)
		}
	}
	if !subnet {
		t.Fatal("aws.subnet did not arrive")
	}
}
```

- [ ] **Step 6: Run the suite and the binary**

```bash
go test -count=1 ./...
go build -o bin/infrata-plugin-aws ./cmd/infrata-plugin-aws
./bin/infrata-plugin-aws; echo "exit $?"                       # expect the "run by infrata" message, exit 2
INFRATA_PLUGIN_COOKIE=x ./bin/infrata-plugin-aws </dev/null     # expect {"protocol":1,"name":"aws","version":"0.0.0-dev"}
go vet ./... && gofmt -l .
```

Expected: PASS; the two binary checks print as commented.

- [ ] **Step 7: Sabotage, then commit**

Sabotages (each must compile and fail a test): drop `ForceNew` from `subnet.vpc_id`
(`TestRegionalAttributesAreRequiredAndForceNew`); make `availability_zone` optional
(same test, the "optional with no default" branch); set the default to `true` (`TestSchemasLoadThroughTheHost`);
rename the type to `awsx.vpc` (both tests).

```bash
GOWORK=off go test -count=1 ./...   # the pinned build passes too, not only the workspace one
git add go.mod go.sum .gitignore cmd/infrata-plugin-aws/main.go internal/awsprov/plugin.go \
  internal/awsprov/definitions.go internal/awsprov/values.go internal/awsprov/definitions_test.go \
  internal/awsprov/protocol_test.go
git commit -m "aws: schemas the host accepts, and a binary that says what it is" -- <same paths>
```

---

### Task 2: Instance configuration and credentials

**Files:**
- Create: `internal/awsprov/config.go`, `internal/awsprov/credentials.go`, `internal/awsprov/testenv_test.go`
- Modify: `internal/awsprov/plugin.go` (`New`)
- Create (minimal, grown in Task 5): `internal/awsprov/provider.go`, `internal/awsprov/clients.go`
- Test: `internal/awsprov/config_test.go`

**Interfaces:**
- Consumes: Task 1's `Plugin`.
- Produces:
  - `type instanceConfig struct { Profile, AssumeRoleARN string; DiscoverRegions []string }`
  - `func parseConfig(values map[string]value.Value) (instanceConfig, error)`
  - `func loadAWSConfig(ctx context.Context, instance string, ic instanceConfig) (aws.Config, error)`
  - `type clients struct` with `func newClients(cfg aws.Config) *clients` and `func (c *clients) ec2(region string) *ec2.Client`
  - `type Provider struct { instance string; config instanceConfig; clients *clients; patience patience }`
    and `func newProvider(instance string, ic instanceConfig, cfg aws.Config) *Provider`
  - test helper `func isolateAWS(t *testing.T, endpoint string)` (in `testenv_test.go`)

This task does not yet need the fake server for most cases; the assume-role test is added in Task 3 once the
fake serves STS.

- [ ] **Step 1: Write the failing tests**

`internal/awsprov/testenv_test.go`:

```go
package awsprov

import (
	"os"
	"path/filepath"
	"testing"
)

// isolateAWS stops the SDK reading anything from the developer's machine: their config and
// credentials files, their profile, EC2 instance metadata (which otherwise costs a second of
// timeouts), and any real endpoint. endpoint may be "" for tests that make no API call.
// Uses t.Setenv, so callers cannot run in parallel.
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
```

`internal/awsprov/config_test.go`:

```go
package awsprov

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/value"
)

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }

func list(items ...string) value.Value {
	vs := make([]value.Value, 0, len(items))
	for _, it := range items {
		vs = append(vs, s(it))
	}
	return value.List(vs, value.SourceExplicit)
}

// TestAnUnknownKeyIsRefusedNamingWhatIsAccepted. A misspelled `profil:` silently ignored falls back
// to the default credential chain, which may be another account.
func TestAnUnknownKeyIsRefusedNamingWhatIsAccepted(t *testing.T) {
	_, err := parseConfig(map[string]value.Value{"profil": s("prod"), "region": s("us-east-1")})
	if err == nil {
		t.Fatal("unknown keys were accepted")
	}
	for _, want := range []string{`"profil"`, `"region"`, "assume_role_arn", "discover_regions", "profile", "defaults: {region"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %s", err, want)
		}
	}
}

func TestConfigValuesAreRead(t *testing.T) {
	ic, err := parseConfig(map[string]value.Value{
		"profile":          s("prod"),
		"assume_role_arn":  s("arn:aws:iam::123456789012:role/deploy"),
		"discover_regions": list("us-east-1", "eu-west-1", "us-east-1"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if ic.Profile != "prod" || ic.AssumeRoleARN != "arn:aws:iam::123456789012:role/deploy" {
		t.Errorf("config = %+v", ic)
	}
	if strings.Join(ic.DiscoverRegions, ",") != "us-east-1,eu-west-1" {
		t.Errorf("discover_regions = %v, want duplicates removed in written order", ic.DiscoverRegions)
	}
}

func TestMalformedValuesAreRefused(t *testing.T) {
	for name, values := range map[string]map[string]value.Value{
		"profile not a string":           {"profile": list("a")},
		"empty profile":                  {"profile": s("")},
		"discover_regions not a list":    {"discover_regions": s("us-east-1")},
		"discover_regions holds a blank": {"discover_regions": list("us-east-1", "")},
		"assume_role_arn not an ARN":     {"assume_role_arn": s("deploy")},
	} {
		if _, err := parseConfig(values); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

// TestANamedProfileThatDoesNotExistFailsInNew. A file read, no network: cheap enough for every
// validate, and the one credential mistake worth catching before an apply starts.
func TestANamedProfileThatDoesNotExistFailsInNew(t *testing.T) {
	isolateAWS(t, "")
	_, err := NewPlugin().New(provider.Config{Instance: "prod", Values: map[string]value.Value{"profile": s("prodd")}})
	if err == nil {
		t.Fatal("a profile missing from the config files was accepted")
	}
	for _, want := range []string{`"prod"`, `"prodd"`, "aws configure list-profiles"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %s", err, want)
		}
	}
}

func TestAProfileThatExistsConfigures(t *testing.T) {
	dir := isolateAWS(t, "")
	if err := os.WriteFile(filepath.Join(dir, "config"), []byte("[profile staging]\nregion = eu-west-1\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := NewPlugin().New(provider.Config{Instance: "staging", Values: map[string]value.Value{"profile": s("staging")}}); err != nil {
		t.Fatalf("New: %v", err)
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 -run 'Key|Config|Malformed|Profile' ./internal/awsprov/`
Expected: FAIL — `undefined: parseConfig`.

- [ ] **Step 3: Implement**

`internal/awsprov/config.go`:

```go
package awsprov

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/infrata/infrata/pkg/value"
)

const (
	keyAssumeRoleARN   = "assume_role_arn"
	keyDiscoverRegions = "discover_regions"
	keyProfile         = "profile"
)

// instanceConfig is one `providers:` entry's own configuration. Credentials and the regions
// discovery scans; never a region for CRUD, which comes from each resource.
type instanceConfig struct {
	Profile         string
	AssumeRoleARN   string
	DiscoverRegions []string
}

// parseConfig fails closed: a key this plugin does not read is refused, naming what it accepts.
func parseConfig(values map[string]value.Value) (instanceConfig, error) {
	var ic instanceConfig
	var unknown []string
	for k := range values {
		switch k {
		case keyAssumeRoleARN, keyDiscoverRegions, keyProfile:
		default:
			unknown = append(unknown, strconv.Quote(k))
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return ic, fmt.Errorf("unknown configuration %s; the aws provider accepts only %s, %s and %s "+
			"(a resource's region belongs under defaults: {region: …}, not here)",
			strings.Join(unknown, ", "), keyAssumeRoleARN, keyDiscoverRegions, keyProfile)
	}

	var err error
	if ic.Profile, err = optionalString(values, keyProfile); err != nil {
		return ic, err
	}
	if ic.AssumeRoleARN, err = optionalString(values, keyAssumeRoleARN); err != nil {
		return ic, err
	}
	if ic.AssumeRoleARN != "" && !strings.HasPrefix(ic.AssumeRoleARN, "arn:") {
		return ic, fmt.Errorf("`%s` must be a role ARN such as arn:aws:iam::123456789012:role/deploy, got %q",
			keyAssumeRoleARN, ic.AssumeRoleARN)
	}
	if v, ok := values[keyDiscoverRegions]; ok {
		items, isList := v.Raw.([]value.Value)
		if v.Kind != value.KindList || !isList {
			return ic, fmt.Errorf("`%s` must be a list of regions, e.g. [us-east-1, eu-west-1], got %s", keyDiscoverRegions, v.Kind)
		}
		seen := map[string]bool{}
		for i, item := range items {
			region, isString := item.AsString()
			if !isString || region == "" {
				return ic, fmt.Errorf("`%s` item %d must be a region name such as us-east-1", keyDiscoverRegions, i+1)
			}
			if !seen[region] {
				seen[region] = true
				ic.DiscoverRegions = append(ic.DiscoverRegions, region)
			}
		}
	}
	return ic, nil
}

func optionalString(values map[string]value.Value, key string) (string, error) {
	v, ok := values[key]
	if !ok {
		return "", nil
	}
	text, isString := v.AsString()
	if !isString {
		return "", fmt.Errorf("`%s` must be a string, got %s", key, v.Kind)
	}
	if text == "" {
		return "", fmt.Errorf("`%s` is empty: give a value, or omit the key to use the default credential chain", key)
	}
	return text, nil
}
```

`internal/awsprov/credentials.go`:

```go
package awsprov

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// stsFallbackRegion is used only when nothing configured a region: resources carry their own, so
// a user may have no default region, and an STS client without one cannot resolve an endpoint.
const stsFallbackRegion = "us-east-1"

// loadAWSConfig loads credentials the way the AWS CLI does, then applies this instance's overrides.
//
// It resolves nothing over the network: credentials are retrieved on first use. Every compiling
// command constructs instances, `validate` included, and an SSO or assume-role call there would put
// a network round trip in front of a syntax check.
func loadAWSConfig(ctx context.Context, instance string, ic instanceConfig) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error
	if ic.Profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(ic.Profile))
	}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		var missing config.SharedConfigProfileNotExistError
		if errors.As(err, &missing) {
			return aws.Config{}, fmt.Errorf("provider instance %q names profile %q, which is not in your AWS config files: "+
				"run \"aws configure list-profiles\" to see what exists, or correct `profile`", instance, ic.Profile)
		}
		return aws.Config{}, fmt.Errorf("provider instance %q: loading AWS configuration: %w", instance, err)
	}
	if ic.AssumeRoleARN != "" {
		client := sts.NewFromConfig(cfg, func(o *sts.Options) {
			if o.Region == "" {
				o.Region = stsFallbackRegion
			}
		})
		cfg.Credentials = aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(client, ic.AssumeRoleARN))
	}
	return cfg, nil
}
```

`internal/awsprov/clients.go`:

```go
package awsprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// clients holds one EC2 client per region, built on first use from the instance's configuration.
type clients struct {
	cfg      aws.Config
	mu       sync.Mutex
	byRegion map[string]*ec2.Client
}

func newClients(cfg aws.Config) *clients {
	return &clients{cfg: cfg, byRegion: map[string]*ec2.Client{}}
}

func (c *clients) ec2(region string) *ec2.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.byRegion[region]; ok {
		return cl
	}
	cl := ec2.NewFromConfig(c.cfg, func(o *ec2.Options) { o.Region = region })
	c.byRegion[region] = cl
	return cl
}

// createOnce stops the SDK resending a create. CreateVpc and CreateSubnet take no client token, so
// a resend after AWS acted makes a second resource nothing records; infrata's executor decides
// instead, from ClassifyError. The SDK ignores a per-call value equal to the client's own, which is
// harmless: equal means the client already makes one attempt.
func createOnce(o *ec2.Options) { o.RetryMaxAttempts = 1 }
```

`internal/awsprov/provider.go` (grown in Tasks 4–7; this is the minimal version that satisfies
`provider.Provider`):

```go
package awsprov

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/schema"
)

// Provider is one configured instance: one AWS account's credentials, any number of regions.
type Provider struct {
	instance string
	config   instanceConfig
	clients  *clients
	patience patience
}

var _ provider.Provider = (*Provider)(nil)

func newProvider(instance string, ic instanceConfig, cfg aws.Config) *Provider {
	return &Provider{instance: instance, config: ic, clients: newClients(cfg), patience: notFoundPatience}
}

func (p *Provider) Name() string                              { return PluginName }
func (p *Provider) Definitions() []*schema.ResourceDefinition { return definitions() }

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	return provider.ErrNotImplemented
}
func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (p *Provider) ClassifyError(err error) provider.Retryability { return provider.NotSafeToRetry }
```

Add a placeholder-free `patience` type now so the struct compiles; Task 5 gives it behaviour and tests.
`internal/awsprov/patience.go`:

```go
package awsprov

import (
	"context"
	"time"
)

// patience is how long a Read waits for a resource EC2 may not have propagated yet.
type patience struct {
	attempts int
	base     time.Duration
	max      time.Duration
	sleep    func(context.Context, time.Duration) error
}

// notFoundPatience: 5 attempts, 250ms doubling, capped at 2s — at most 3.75s of waiting.
var notFoundPatience = patience{attempts: 5, base: 250 * time.Millisecond, max: 2 * time.Second, sleep: sleepCtx}

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

Replace `Plugin.New` in `plugin.go` (and drop the `errors` import):

```go
// New builds one configured instance. An error here is rendered against the `providers:` entry, so
// it must say what is wrong and what to set.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	ic, err := parseConfig(cfg.Values)
	if err != nil {
		return nil, err
	}
	awsCfg, err := loadAWSConfig(context.Background(), cfg.Instance, ic)
	if err != nil {
		return nil, err
	}
	return newProvider(cfg.Instance, ic, awsCfg), nil
}
```

- [ ] **Step 4: Run to see them pass**

Run: `go test -count=1 ./... && go vet ./... && gofmt -l .`
Expected: PASS, no output from gofmt.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: accept every key (`TestAnUnknownKeyIsRefused…`); drop the dedupe (`TestConfigValuesAreRead`); swallow
`SharedConfigProfileNotExistError` and continue with defaults (`TestANamedProfileThatDoesNotExistFailsInNew`);
skip the `arn:` check (`TestMalformedValuesAreRefused`).

```bash
git add internal/awsprov/config.go internal/awsprov/credentials.go internal/awsprov/clients.go \
  internal/awsprov/provider.go internal/awsprov/patience.go internal/awsprov/plugin.go \
  internal/awsprov/testenv_test.go internal/awsprov/config_test.go
git commit -m "aws: instance configuration fails closed, and a missing profile fails before an apply" -- <same paths>
```

---
### Task 3: `internal/ec2fake` — a fake EC2 (and STS) endpoint, with the real SDK as its oracle

**Files:**
- Create: `internal/ec2fake/server.go`, `internal/ec2fake/xml.go`
- Test: `internal/ec2fake/server_test.go`, and add the assume-role case to `internal/awsprov/config_test.go`

**Interfaces:**
- Produces (used by every later test):
  - `func New() *Server` (starts an `httptest.Server`; `Server.URL`; `Server.Close()`)
  - `type Fault struct { Action string; Nth int; Status int; Code, Message string; Drop bool }` —
    `Nth` counts calls of `Action` from 1; `Drop` applies the action then closes the connection with no response
  - `func (s *Server) Inject(f Fault)`
  - `func (s *Server) HideFromDescribe(id string, times int)` — the next `times` describes omit `id` (by-ID
    describes answer `Invalid…ID.NotFound`)
  - `func (s *Server) Calls(action string) int`, `func (s *Server) AccessKeys() []string`
  - `func (s *Server) AddVPC(region, cidr string, tags map[string]string) string` (unmanaged infrastructure)
  - `func (s *Server) SetTag(id, key, value string)`, `func (s *Server) SetCIDR(id, cidr string)`, `func (s *Server) Remove(id string)` (drift)
  - `func (s *Server) VPCs() []VPC`, `func (s *Server) Subnets() []Subnet` (snapshots); `type VPC struct{ ID, Region, CIDR, Owner string; Default bool; Tags map[string]string }`,
    `type Subnet struct{ ID, Region, VPCID, CIDR, AZ, Owner string; MapPublicIP bool; Tags map[string]string }`
  - `PageSize int` field (0 = everything in one page)
  - `const Owner = "123456789012"`, `const AssumedAccessKey = "ASIAFAKEASSUMED"`

The fake keeps resources per region, taking the region from the SigV4 credential scope in `Authorization`
(`Credential=<key>/<date>/<region>/<service>/aws4_request`), which is also where it reads the access key ID.

- [ ] **Step 1: Write the oracle tests**

`internal/ec2fake/server_test.go`:

```go
package ec2fake

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
)

// client is a real SDK client pointed at the fake. If the SDK can decode what the fake says, the
// fake's wire format is right: the SDK is the oracle, not this package's own reading of it.
func client(t *testing.T, s *Server, region string) *ec2.Client {
	t.Helper()
	return ec2.New(ec2.Options{
		Region:       region,
		BaseEndpoint: aws.String(s.URL),
		Credentials:  credentials.NewStaticCredentialsProvider("AKIDORACLE", "secret", ""),
	})
}

func TestTheSDKDecodesAVPCRoundTrip(t *testing.T) {
	s := New()
	defer s.Close()
	ctx := context.Background()
	c := client(t, s, "us-east-1")

	out, err := c.CreateVpc(ctx, &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
		TagSpecifications: []types.TagSpecification{{
			ResourceType: types.ResourceTypeVpc,
			Tags:         []types.Tag{{Key: aws.String("team"), Value: aws.String("platform")}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	id := aws.ToString(out.Vpc.VpcId)
	if id == "" || aws.ToString(out.Vpc.CidrBlock) != "10.0.0.0/16" || len(out.Vpc.Tags) != 1 {
		t.Fatalf("CreateVpc decoded as %+v", out.Vpc)
	}
	got, err := c.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{id}})
	if err != nil || len(got.Vpcs) != 1 || aws.ToString(got.Vpcs[0].OwnerId) != Owner {
		t.Fatalf("DescribeVpcs = %+v, %v", got, err)
	}
	if _, err := c.DeleteTags(ctx, &ec2.DeleteTagsInput{Resources: []string{id}, Tags: []types.Tag{{Key: aws.String("team")}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.DeleteVpc(ctx, &ec2.DeleteVpcInput{VpcId: aws.String(id)}); err != nil {
		t.Fatal(err)
	}
	_, err = c.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{id}})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "InvalidVpcID.NotFound" {
		t.Fatalf("describing a deleted VPC = %v, want InvalidVpcID.NotFound", err)
	}
	if keys := s.AccessKeys(); len(keys) == 0 || keys[0] != "AKIDORACLE" {
		t.Errorf("access keys seen = %v", keys)
	}
}

func TestTheSDKDecodesASubnetAndItsAttribute(t *testing.T) {
	s := New()
	defer s.Close()
	ctx := context.Background()
	c := client(t, s, "eu-west-1")
	vpc := s.AddVPC("eu-west-1", "10.1.0.0/16", nil)

	out, err := c.CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId: aws.String(vpc), CidrBlock: aws.String("10.1.1.0/24"), AvailabilityZone: aws.String("eu-west-1a"),
	})
	if err != nil {
		t.Fatal(err)
	}
	id := aws.ToString(out.Subnet.SubnetId)
	if _, err := c.ModifySubnetAttribute(ctx, &ec2.ModifySubnetAttributeInput{
		SubnetId: aws.String(id), MapPublicIpOnLaunch: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	}); err != nil {
		t.Fatal(err)
	}
	got, err := c.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{SubnetIds: []string{id}})
	if err != nil || len(got.Subnets) != 1 {
		t.Fatalf("DescribeSubnets = %+v, %v", got, err)
	}
	sub := got.Subnets[0]
	if !aws.ToBool(sub.MapPublicIpOnLaunch) || aws.ToString(sub.AvailabilityZone) != "eu-west-1a" ||
		aws.ToString(sub.SubnetArn) != "arn:aws:ec2:eu-west-1:"+Owner+":subnet/"+id || aws.ToString(sub.VpcId) != vpc {
		t.Fatalf("subnet decoded as %+v", sub)
	}
	// A subnet in a VPC the region does not hold is refused the way EC2 refuses it.
	_, err = client(t, s, "us-east-1").CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId: aws.String(vpc), CidrBlock: aws.String("10.1.2.0/24"), AvailabilityZone: aws.String("us-east-1a"),
	})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "InvalidVpcID.NotFound" {
		t.Fatalf("cross-region CreateSubnet = %v, want InvalidVpcID.NotFound", err)
	}
}

// TestRegionsArePartitionedAndPagesFollowNextToken. The fixture holds three VPCs in one region and
// one in another, with a page size of one, so neither a region leak nor a paginator that stops after
// the first page can pass.
func TestRegionsArePartitionedAndPagesFollowNextToken(t *testing.T) {
	s := New()
	defer s.Close()
	s.PageSize = 1
	for _, cidr := range []string{"10.0.0.0/16", "10.1.0.0/16", "10.2.0.0/16"} {
		s.AddVPC("us-east-1", cidr, nil)
	}
	s.AddVPC("eu-west-1", "10.9.0.0/16", nil)

	var n int
	pg := ec2.NewDescribeVpcsPaginator(client(t, s, "us-east-1"), &ec2.DescribeVpcsInput{})
	for pg.HasMorePages() {
		page, err := pg.NextPage(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		n += len(page.Vpcs)
	}
	if n != 3 {
		t.Fatalf("paginated us-east-1 VPCs = %d, want 3", n)
	}
}

func TestAnInjectedFaultDecodesWithItsStatus(t *testing.T) {
	s := New()
	defer s.Close()
	s.Inject(Fault{Action: "DescribeVpcs", Nth: 1, Status: 400, Code: "UnauthorizedOperation", Message: "no"})
	_, err := client(t, s, "us-east-1").DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) || apiErr.ErrorCode() != "UnauthorizedOperation" {
		t.Fatalf("err = %v, want UnauthorizedOperation", err)
	}
	if s.Calls("DescribeVpcs") != 1 {
		t.Errorf("calls = %d, want 1 (a 400 is not retried by the SDK)", s.Calls("DescribeVpcs"))
	}
}

func TestHiddenResourcesAreNotFoundForAWhile(t *testing.T) {
	s := New()
	defer s.Close()
	id := s.AddVPC("us-east-1", "10.0.0.0/16", nil)
	s.HideFromDescribe(id, 2)
	c := client(t, s, "us-east-1")
	for i := 1; i <= 3; i++ {
		_, err := c.DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{VpcIds: []string{id}})
		if hidden := err != nil; hidden != (i <= 2) {
			t.Fatalf("describe %d: err = %v", i, err)
		}
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 ./internal/ec2fake/`
Expected: FAIL — `undefined: New`.

- [ ] **Step 3: Implement the fake**

`internal/ec2fake/xml.go`:

```go
package ec2fake

import (
	"encoding/xml"
	"net/http"
)

const ec2NS = "http://ec2.amazonaws.com/doc/2016-11-15/"

type xmlTag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type xmlVPC struct {
	VpcID     string   `xml:"vpcId"`
	CidrBlock string   `xml:"cidrBlock"`
	OwnerID   string   `xml:"ownerId"`
	State     string   `xml:"state"`
	IsDefault bool     `xml:"isDefault"`
	Tags      []xmlTag `xml:"tagSet>item"`
}

type xmlSubnet struct {
	SubnetID            string   `xml:"subnetId"`
	VpcID               string   `xml:"vpcId"`
	CidrBlock           string   `xml:"cidrBlock"`
	AvailabilityZone    string   `xml:"availabilityZone"`
	MapPublicIPOnLaunch bool     `xml:"mapPublicIpOnLaunch"`
	SubnetArn           string   `xml:"subnetArn"`
	OwnerID             string   `xml:"ownerId"`
	State               string   `xml:"state"`
	Tags                []xmlTag `xml:"tagSet>item"`
}

type createVpcResponse struct {
	XMLName   xml.Name `xml:"CreateVpcResponse"`
	Xmlns     string   `xml:"xmlns,attr"`
	RequestID string   `xml:"requestId"`
	Vpc       xmlVPC   `xml:"vpc"`
}

type describeVpcsResponse struct {
	XMLName   xml.Name `xml:"DescribeVpcsResponse"`
	Xmlns     string   `xml:"xmlns,attr"`
	RequestID string   `xml:"requestId"`
	Vpcs      []xmlVPC `xml:"vpcSet>item"`
	NextToken string   `xml:"nextToken,omitempty"`
}

type createSubnetResponse struct {
	XMLName   xml.Name  `xml:"CreateSubnetResponse"`
	Xmlns     string    `xml:"xmlns,attr"`
	RequestID string    `xml:"requestId"`
	Subnet    xmlSubnet `xml:"subnet"`
}

type describeSubnetsResponse struct {
	XMLName   xml.Name    `xml:"DescribeSubnetsResponse"`
	Xmlns     string      `xml:"xmlns,attr"`
	RequestID string      `xml:"requestId"`
	Subnets   []xmlSubnet `xml:"subnetSet>item"`
	NextToken string      `xml:"nextToken,omitempty"`
}

// returnResponse is every EC2 action here that answers only success.
type returnResponse struct {
	XMLName   xml.Name
	Xmlns     string `xml:"xmlns,attr"`
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type assumeRoleResponse struct {
	XMLName xml.Name `xml:"AssumeRoleResponse"`
	Result  struct {
		Credentials struct {
			AccessKeyID     string `xml:"AccessKeyId"`
			SecretAccessKey string `xml:"SecretAccessKey"`
			SessionToken    string `xml:"SessionToken"`
			Expiration      string `xml:"Expiration"`
		} `xml:"Credentials"`
	} `xml:"AssumeRoleResult"`
}

// errorResponse is EC2's query-protocol error body: Errors>Error>Code, Errors>Error>Message, RequestID
// (aws-sdk-go-v2 aws/protocol/ec2query/error_utils.go).
type errorResponse struct {
	XMLName   xml.Name `xml:"Response"`
	Code      string   `xml:"Errors>Error>Code"`
	Message   string   `xml:"Errors>Error>Message"`
	RequestID string   `xml:"RequestID"`
}

func writeXML(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "text/xml;charset=UTF-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeXML(w, status, errorResponse{Code: code, Message: message, RequestID: "fake-request-err"})
}
```

`internal/ec2fake/server.go`:

```go
// Package ec2fake is an in-process stand-in for the parts of EC2 and STS this plugin calls. It speaks
// the query protocol, so tests drive the real AWS SDK against it: serialisation, error decoding,
// retries and endpoint overrides are all the SDK's own.
package ec2fake

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	Owner            = "123456789012"
	AssumedAccessKey = "ASIAFAKEASSUMED"
)

type VPC struct {
	ID, Region, CIDR, Owner string
	Default                 bool
	Tags                    map[string]string
}

type Subnet struct {
	ID, Region, VPCID, CIDR, AZ, Owner string
	MapPublicIP                        bool
	Tags                               map[string]string
}

type Fault struct {
	Action  string
	Nth     int
	Status  int
	Code    string
	Message string
	Drop    bool
}

type Server struct {
	*httptest.Server
	PageSize int

	mu      sync.Mutex
	next    int
	vpcs    map[string]*VPC
	subnets map[string]*Subnet
	faults  []Fault
	hidden  map[string]int
	calls   map[string]int
	keys    []string
}

func New() *Server {
	s := &Server{vpcs: map[string]*VPC{}, subnets: map[string]*Subnet{}, hidden: map[string]int{}, calls: map[string]int{}}
	s.Server = httptest.NewServer(http.HandlerFunc(s.serve))
	return s
}

var scope = regexp.MustCompile(`Credential=([^/]+)/[^/]+/([^/]+)/`)

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		writeError(w, 400, "MalformedQueryString", err.Error())
		return
	}
	action := r.Form.Get("Action")
	key, region := "", ""
	if m := scope.FindStringSubmatch(r.Header.Get("Authorization")); m != nil {
		key, region = m[1], m[2]
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.calls[action]++
	s.keys = append(s.keys, key)
	fault, faulted := s.faultFor(action)
	if faulted && !fault.Drop {
		writeError(w, fault.Status, fault.Code, fault.Message)
		return
	}
	status, body := s.handle(action, region, r)
	if faulted && fault.Drop {
		// The action took effect; the answer never arrives. What a connection lost after the
		// request was processed looks like.
		if hj, ok := w.(http.Hijacker); ok {
			if conn, _, err := hj.Hijack(); err == nil {
				_ = conn.Close()
				return
			}
		}
	}
	if e, isErr := body.(errorResponse); isErr {
		writeError(w, status, e.Code, e.Message)
		return
	}
	writeXML(w, status, body)
}

func (s *Server) faultFor(action string) (Fault, bool) {
	for i, f := range s.faults {
		if f.Action == action && f.Nth == s.calls[action] {
			s.faults = append(s.faults[:i], s.faults[i+1:]...)
			return f, true
		}
	}
	return Fault{}, false
}

func notFound(code, id string) (int, any) {
	return 400, errorResponse{Code: code, Message: fmt.Sprintf("The ID '%s' does not exist", id)}
}

func (s *Server) handle(action, region string, r *http.Request) (int, any) {
	f := r.Form
	switch action {
	case "AssumeRole":
		var out assumeRoleResponse
		out.Result.Credentials.AccessKeyID = AssumedAccessKey
		out.Result.Credentials.SecretAccessKey = "assumed-secret"
		out.Result.Credentials.SessionToken = "assumed-token"
		out.Result.Credentials.Expiration = "2099-01-01T00:00:00Z"
		return 200, out

	case "CreateVpc":
		id := s.newID("vpc")
		v := &VPC{ID: id, Region: region, CIDR: f.Get("CidrBlock"), Owner: Owner, Tags: tagSpec(f)}
		s.vpcs[id] = v
		return 200, createVpcResponse{Xmlns: ec2NS, RequestID: "fake-request", Vpc: vpcXML(v)}

	case "DescribeVpcs":
		ids := indexed(f, "VpcId")
		var out []xmlVPC
		for _, id := range ids {
			v, ok := s.vpcs[id]
			if !ok || v.Region != region || s.hide(id) {
				return notFound("InvalidVpcID.NotFound", id)
			}
			out = append(out, vpcXML(v))
		}
		if len(ids) == 0 {
			for _, id := range sortedKeys(s.vpcs) {
				if v := s.vpcs[id]; v.Region == region && !s.hide(id) {
					out = append(out, vpcXML(v))
				}
			}
		}
		page, next := s.paginate(len(out), f.Get("NextToken"))
		return 200, describeVpcsResponse{Xmlns: ec2NS, RequestID: "fake-request", Vpcs: out[page[0]:page[1]], NextToken: next}

	case "DeleteVpc":
		id := f.Get("VpcId")
		v, ok := s.vpcs[id]
		if !ok || v.Region != region {
			return notFound("InvalidVpcID.NotFound", id)
		}
		for _, sub := range s.subnets {
			if sub.VPCID == id {
				return 400, errorResponse{Code: "DependencyViolation", Message: "The vpc '" + id + "' has dependencies and cannot be deleted."}
			}
		}
		delete(s.vpcs, id)
		return 200, returnResponse{XMLName: xmlName("DeleteVpcResponse"), Xmlns: ec2NS, RequestID: "fake-request", Return: true}

	case "CreateSubnet":
		vpcID := f.Get("VpcId")
		if v, ok := s.vpcs[vpcID]; !ok || v.Region != region || s.hide(vpcID) {
			return notFound("InvalidVpcID.NotFound", vpcID)
		}
		id := s.newID("subnet")
		sub := &Subnet{ID: id, Region: region, VPCID: vpcID, CIDR: f.Get("CidrBlock"), AZ: f.Get("AvailabilityZone"), Owner: Owner, Tags: tagSpec(f)}
		s.subnets[id] = sub
		return 200, createSubnetResponse{Xmlns: ec2NS, RequestID: "fake-request", Subnet: subnetXML(sub)}

	case "DescribeSubnets":
		ids := indexed(f, "SubnetId")
		var out []xmlSubnet
		for _, id := range ids {
			sub, ok := s.subnets[id]
			if !ok || sub.Region != region || s.hide(id) {
				return notFound("InvalidSubnetID.NotFound", id)
			}
			out = append(out, subnetXML(sub))
		}
		if len(ids) == 0 {
			for _, id := range sortedKeys(s.subnets) {
				if sub := s.subnets[id]; sub.Region == region && !s.hide(id) {
					out = append(out, subnetXML(sub))
				}
			}
		}
		page, next := s.paginate(len(out), f.Get("NextToken"))
		return 200, describeSubnetsResponse{Xmlns: ec2NS, RequestID: "fake-request", Subnets: out[page[0]:page[1]], NextToken: next}

	case "DeleteSubnet":
		id := f.Get("SubnetId")
		if sub, ok := s.subnets[id]; !ok || sub.Region != region {
			return notFound("InvalidSubnetID.NotFound", id)
		}
		delete(s.subnets, id)
		return 200, returnResponse{XMLName: xmlName("DeleteSubnetResponse"), Xmlns: ec2NS, RequestID: "fake-request", Return: true}

	case "ModifySubnetAttribute":
		id := f.Get("SubnetId")
		sub, ok := s.subnets[id]
		if !ok || sub.Region != region || s.hide(id) {
			return notFound("InvalidSubnetID.NotFound", id)
		}
		if v := f.Get("MapPublicIpOnLaunch.Value"); v != "" {
			sub.MapPublicIP = v == "true"
		}
		return 200, returnResponse{XMLName: xmlName("ModifySubnetAttributeResponse"), Xmlns: ec2NS, RequestID: "fake-request", Return: true}

	case "CreateTags", "DeleteTags":
		for _, id := range indexed(f, "ResourceId") {
			tags := s.tagsOf(id, region)
			if tags == nil {
				return notFound("InvalidID", id)
			}
			for i := 1; f.Has(fmt.Sprintf("Tag.%d.Key", i)); i++ {
				k := f.Get(fmt.Sprintf("Tag.%d.Key", i))
				if action == "CreateTags" {
					tags[k] = f.Get(fmt.Sprintf("Tag.%d.Value", i))
				} else {
					delete(tags, k)
				}
			}
		}
		return 200, returnResponse{XMLName: xmlName(action + "Response"), Xmlns: ec2NS, RequestID: "fake-request", Return: true}
	}
	return 400, errorResponse{Code: "InvalidAction", Message: "ec2fake does not implement " + action}
}

func (s *Server) newID(prefix string) string {
	s.next++
	return fmt.Sprintf("%s-%017x", prefix, s.next)
}

func (s *Server) hide(id string) bool {
	if s.hidden[id] > 0 {
		s.hidden[id]--
		return true
	}
	return false
}

// paginate returns the [start, end) window for this page and the token for the next.
func (s *Server) paginate(n int, token string) ([2]int, string) {
	start, _ := strconv.Atoi(token)
	if s.PageSize <= 0 || start+s.PageSize >= n {
		return [2]int{min(start, n), n}, ""
	}
	return [2]int{start, start + s.PageSize}, strconv.Itoa(start + s.PageSize)
}

func (s *Server) tagsOf(id, region string) map[string]string {
	if v, ok := s.vpcs[id]; ok && v.Region == region {
		if v.Tags == nil {
			v.Tags = map[string]string{}
		}
		return v.Tags
	}
	if sub, ok := s.subnets[id]; ok && sub.Region == region {
		if sub.Tags == nil {
			sub.Tags = map[string]string{}
		}
		return sub.Tags
	}
	return nil
}

// indexed reads a flattened query list: Name.1, Name.2, …
func indexed(f map[string][]string, name string) []string {
	var out []string
	for i := 1; ; i++ {
		v, ok := f[fmt.Sprintf("%s.%d", name, i)]
		if !ok {
			return out
		}
		out = append(out, v[0])
	}
}

// tagSpec reads TagSpecification.1.Tag.N.Key/Value.
func tagSpec(f map[string][]string) map[string]string {
	tags := map[string]string{}
	for i := 1; ; i++ {
		k, ok := f[fmt.Sprintf("TagSpecification.1.Tag.%d.Key", i)]
		if !ok {
			return tags
		}
		v := f[fmt.Sprintf("TagSpecification.1.Tag.%d.Value", i)]
		tags[k[0]] = strings.Join(v, "")
	}
}

func vpcXML(v *VPC) xmlVPC {
	return xmlVPC{VpcID: v.ID, CidrBlock: v.CIDR, OwnerID: v.Owner, State: "available", IsDefault: v.Default, Tags: tagsXML(v.Tags)}
}

func subnetXML(sub *Subnet) xmlSubnet {
	return xmlSubnet{
		SubnetID: sub.ID, VpcID: sub.VPCID, CidrBlock: sub.CIDR, AvailabilityZone: sub.AZ,
		MapPublicIPOnLaunch: sub.MapPublicIP, OwnerID: sub.Owner, State: "available",
		SubnetArn: fmt.Sprintf("arn:aws:ec2:%s:%s:subnet/%s", sub.Region, sub.Owner, sub.ID), Tags: tagsXML(sub.Tags),
	}
}

func tagsXML(tags map[string]string) []xmlTag {
	var out []xmlTag
	for _, k := range sortedKeys(tags) {
		out = append(out, xmlTag{Key: k, Value: tags[k]})
	}
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ---- test controls ----

func (s *Server) Inject(f Fault) { s.mu.Lock(); defer s.mu.Unlock(); s.faults = append(s.faults, f) }

func (s *Server) HideFromDescribe(id string, times int) { s.mu.Lock(); defer s.mu.Unlock(); s.hidden[id] = times }

func (s *Server) Calls(action string) int { s.mu.Lock(); defer s.mu.Unlock(); return s.calls[action] }

func (s *Server) AccessKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.keys...)
}

func (s *Server) AddVPC(region, cidr string, tags map[string]string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.newID("vpc")
	s.vpcs[id] = &VPC{ID: id, Region: region, CIDR: cidr, Owner: Owner, Tags: tags}
	return id
}

func (s *Server) SetTag(id, key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, region := range []string{s.regionOf(id)} {
		s.tagsOf(id, region)[key] = val
	}
}

func (s *Server) SetCIDR(id, cidr string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.vpcs[id]; ok {
		v.CIDR = cidr
	}
	if sub, ok := s.subnets[id]; ok {
		sub.CIDR = cidr
	}
}

func (s *Server) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.vpcs, id)
	delete(s.subnets, id)
}

func (s *Server) regionOf(id string) string {
	if v, ok := s.vpcs[id]; ok {
		return v.Region
	}
	if sub, ok := s.subnets[id]; ok {
		return sub.Region
	}
	return ""
}

func (s *Server) VPCs() []VPC {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []VPC
	for _, id := range sortedKeys(s.vpcs) {
		out = append(out, *s.vpcs[id])
	}
	return out
}

func (s *Server) Subnets() []Subnet {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []Subnet
	for _, id := range sortedKeys(s.subnets) {
		out = append(out, *s.subnets[id])
	}
	return out
}
```

Add to `xml.go`:

```go
func xmlName(local string) xml.Name { return xml.Name{Local: local} }
```

- [ ] **Step 4: Run the oracle tests**

Run: `go test -count=1 ./internal/ec2fake/`
Expected: PASS. If the SDK fails to decode a response, the element names in `xml.go` are wrong: correct them
against `service/ec2@v1.332.0/deserializers.go` (search the `awsEc2query_deserializeDocument…` function for the
type) and record the correction in the Verification log. The SDK is authoritative, not this plan.

- [ ] **Step 5: Add the assume-role test to `internal/awsprov/config_test.go`**

It needs a provider operation that makes an EC2 call; `Discover` exists only after Task 7, so call the client
directly through the provider's client cache:

```go
// TestAssumeRoleSignsEC2CallsWithTheAssumedCredentials. The fixture's static key would sign every
// call if the role were ignored, so the assertion cannot pass by accident.
func TestAssumeRoleSignsEC2CallsWithTheAssumedCredentials(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	isolateAWS(t, fake.URL)

	prov, err := NewPlugin().New(provider.Config{Instance: "deploy", Values: map[string]value.Value{
		"assume_role_arn": s("arn:aws:iam::123456789012:role/deploy"),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := prov.(*Provider).clients.ec2("us-east-1").DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{}); err != nil {
		t.Fatal(err)
	}
	if fake.Calls("AssumeRole") != 1 {
		t.Errorf("AssumeRole calls = %d, want 1", fake.Calls("AssumeRole"))
	}
	keys := fake.AccessKeys()
	if last := keys[len(keys)-1]; last != ec2fake.AssumedAccessKey {
		t.Errorf("DescribeVpcs was signed with %q, want the assumed role's %q", last, ec2fake.AssumedAccessKey)
	}
}
```

(imports: `context`, `github.com/aws/aws-sdk-go-v2/service/ec2`, `github.com/infrata/infrata-provider-aws/internal/ec2fake`.)

Run: `go test -count=1 ./...` — Expected: PASS.

- [ ] **Step 6: Sabotage, then commit**

Sabotages: fake ignores the region on `DescribeVpcs` (`TestRegionsArePartitioned…` fails — the fixture has a
VPC in another region); fake never sets `nextToken` with `PageSize` (same test: 1, not 3); rename `vpcSet` to
`vpcs` (`TestTheSDKDecodesAVPCRoundTrip`); in `loadAWSConfig`, skip the assume-role wrap
(`TestAssumeRoleSigns…`).

```bash
git add internal/ec2fake/server.go internal/ec2fake/xml.go internal/ec2fake/server_test.go internal/awsprov/config_test.go
git commit -m "ec2fake: a fake EC2 endpoint the real SDK decodes, so tests fake the cloud and not the code" -- <same paths>
```

---

### Task 4: Provider IDs and error classification

**Files:**
- Create: `internal/awsprov/ids.go`, `internal/awsprov/errors.go`
- Modify: `internal/awsprov/provider.go` (`ClassifyError` → `classify`)
- Test: `internal/awsprov/ids_test.go`, `internal/awsprov/errors_test.go`

**Interfaces:**
- Produces:
  - `func formatID(region, awsID string) string`
  - `func parseID(resourceType, providerID string) (region, awsID string, err error)`
  - `func classify(err error) provider.Retryability`
  - `func (p *Provider) failed(op, region, awsID string, err error) error` — the D8 message wrapper (type `apiFailure`, with `Unwrap`)
  - `func hasCode(err error, codes ...string) bool`

- [ ] **Step 1: Write the failing tests**

`internal/awsprov/ids_test.go`:

```go
package awsprov

import (
	"strings"
	"testing"
)

func TestProviderIDsRoundTrip(t *testing.T) {
	region, id, err := parseID(typeVPC, formatID("eu-west-1", "vpc-0abc123"))
	if err != nil || region != "eu-west-1" || id != "vpc-0abc123" {
		t.Fatalf("parseID = %q, %q, %v", region, id, err)
	}
}

// TestAnIDOfTheWrongTypeIsRefused. Importing a VPC as a subnet would record a VPC as a subnet in
// state, and the next plan would propose replacing real infrastructure to settle it.
func TestAnIDOfTheWrongTypeIsRefused(t *testing.T) {
	for _, c := range []struct{ typ, id, want string }{
		{typeSubnet, "us-east-1/vpc-0abc123", "subnet-"},
		{typeVPC, "us-east-1/subnet-0abc123", "vpc-"},
		{typeVPC, "vpc-0abc123", "<region>/"},
		{typeVPC, "us-east-1/", "<region>/"},
		{typeVPC, "/vpc-0abc123", "<region>/"},
		{typeVPC, "us-east-1/vpc-1/extra", "<region>/"},
	} {
		_, _, err := parseID(c.typ, c.id)
		if err == nil || !strings.Contains(err.Error(), c.want) || !strings.Contains(err.Error(), c.id) {
			t.Errorf("parseID(%s, %q) = %v, want an error naming the ID and %q", c.typ, c.id, err, c.want)
		}
	}
}
```

`internal/awsprov/errors_test.go` — every case goes through the real SDK against the fake, or against a closed
listener, so the error chains are the ones production sees:

```go
package awsprov

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/provider"
)

func sdkClient(endpoint string, attempts int) *ec2.Client {
	return ec2.New(ec2.Options{
		Region: "us-east-1", BaseEndpoint: aws.String(endpoint), RetryMaxAttempts: attempts,
		Credentials: credentials.NewStaticCredentialsProvider("AKIDTEST", "secret", ""),
	})
}

func TestClassificationOfRealSDKErrors(t *testing.T) {
	cases := []struct {
		name  string
		fault ec2fake.Fault
		want  provider.Retryability
	}{
		// A throttle carrying a 503: the code must win over the status.
		{"throttle", ec2fake.Fault{Status: 503, Code: "RequestLimitExceeded", Message: "slow down"}, provider.SafeToRetry},
		{"server fault", ec2fake.Fault{Status: 500, Code: "InternalError", Message: "oops"}, provider.ConditionallyRetryable},
		{"unavailable", ec2fake.Fault{Status: 503, Code: "Unavailable", Message: "later"}, provider.ConditionallyRetryable},
		{"validation", ec2fake.Fault{Status: 400, Code: "InvalidParameterValue", Message: "bad cidr"}, provider.NotSafeToRetry},
		{"access", ec2fake.Fault{Status: 403, Code: "UnauthorizedOperation", Message: "no"}, provider.NotSafeToRetry},
		{"dropped after acting", ec2fake.Fault{Drop: true}, provider.ConditionallyRetryable},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fake := ec2fake.New()
			defer fake.Close()
			c.fault.Action, c.fault.Nth = "CreateVpc", 1
			fake.Inject(c.fault)
			_, err := sdkClient(fake.URL, 1).CreateVpc(context.Background(), &ec2.CreateVpcInput{CidrBlock: aws.String("10.0.0.0/16")})
			if err == nil {
				t.Fatal("expected the injected failure")
			}
			if got := classify(err); got != c.want {
				t.Errorf("classify(%v) = %v, want %v", err, got, c.want)
			}
		})
	}
}

// TestAnErrorTheSDKGaveUpOnIsStillClassified. MaxAttemptsError wraps the last error; classification
// must see through it.
func TestAnErrorTheSDKGaveUpOnIsStillClassified(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	for n := 1; n <= 3; n++ {
		fake.Inject(ec2fake.Fault{Action: "DescribeVpcs", Nth: n, Status: 503, Code: "RequestLimitExceeded"})
	}
	_, err := sdkClient(fake.URL, 3).DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	if err == nil || fake.Calls("DescribeVpcs") != 3 {
		t.Fatalf("err = %v after %d calls; want failure after 3", err, fake.Calls("DescribeVpcs"))
	}
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify = %v, want SafeToRetry", got)
	}
}

// TestARefusedConnectionIsSafeToRetry. Nothing was sent, so nothing can have happened.
func TestARefusedConnectionIsSafeToRetry(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()
	_, err = sdkClient("http://"+addr, 1).CreateVpc(context.Background(), &ec2.CreateVpcInput{CidrBlock: aws.String("10.0.0.0/16")})
	if got := classify(err); got != provider.SafeToRetry {
		t.Errorf("classify(%v) = %v, want SafeToRetry", err, got)
	}
}

func TestCancellationAndUnknownsAreNotSafe(t *testing.T) {
	for _, err := range []error{context.Canceled, errors.New("something new")} {
		if got := classify(err); got != provider.NotSafeToRetry {
			t.Errorf("classify(%v) = %v", err, got)
		}
	}
	if got := classify(context.DeadlineExceeded); got != provider.ConditionallyRetryable {
		t.Errorf("classify(DeadlineExceeded) = %v", got)
	}
}

// TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause.
func TestAFailureMessageSaysWhatAndWhereAndKeepsItsCause(t *testing.T) {
	fake := ec2fake.New()
	defer fake.Close()
	fake.Inject(ec2fake.Fault{Action: "DeleteVpc", Nth: 1, Status: 400, Code: "DependencyViolation", Message: "has dependencies"})
	_, err := sdkClient(fake.URL, 1).DeleteVpc(context.Background(), &ec2.DeleteVpcInput{VpcId: aws.String("vpc-1")})
	p := &Provider{instance: "prod"}
	wrapped := p.failed("DeleteVpc", "us-east-1", "vpc-1", err)
	for _, want := range []string{`"prod"`, "DeleteVpc", "us-east-1/vpc-1", "DependencyViolation", "has dependencies", "fake-request-err"} {
		if !strings.Contains(wrapped.Error(), want) {
			t.Errorf("%q lacks %q", wrapped, want)
		}
	}
	if !hasCode(wrapped, "DependencyViolation") || classify(wrapped) != provider.NotSafeToRetry {
		t.Error("wrapping lost the cause")
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 -run 'ID|Classif|SDKGave|Refused|Cancellation|FailureMessage' ./internal/awsprov/`
Expected: FAIL — `undefined: parseID`, `undefined: classify`.

- [ ] **Step 3: Implement**

`internal/awsprov/ids.go`:

```go
package awsprov

import (
	"fmt"
	"strings"
)

// formatID is the one provider-ID form used everywhere — Create, Discover, Import — because infrata
// imports by matching `<type>.<provider id>` against what Discover returned, and Import is given no
// region of its own.
func formatID(region, awsID string) string { return region + "/" + awsID }

var idPrefixes = map[string]string{typeVPC: "vpc-", typeSubnet: "subnet-"}

// parseID refuses a malformed ID or one naming a different kind of resource before any API call.
func parseID(resourceType, providerID string) (string, string, error) {
	prefix := idPrefixes[resourceType]
	region, awsID, ok := strings.Cut(providerID, "/")
	if !ok || region == "" || awsID == "" || strings.Contains(awsID, "/") {
		return "", "", fmt.Errorf("%q is not a %s ID: expected <region>/<id>, e.g. us-east-1/%s0abc123",
			providerID, resourceType, prefix)
	}
	if !strings.HasPrefix(awsID, prefix) {
		return "", "", fmt.Errorf("%q is not a %s: a %s ID starts with %q", providerID, resourceType, resourceType, prefix)
	}
	return region, awsID, nil
}
```

`internal/awsprov/errors.go`:

```go
package awsprov

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
	"github.com/infrata/infrata/pkg/provider"
)

// classify answers infrata's question — how dangerous is another attempt — from the error alone. It
// must be a pure function: the plugin SDK asks any one configured instance, not the one that failed.
//
// Order matters. A throttle may carry a 5xx status, so the throttle code is checked before the status. A
// *url.Error is a net.Error, so the dial case is checked before the generic one. EC2's query-protocol
// errors carry no smithy fault, so the HTTP status is the only server-fault signal.
func classify(err error) provider.Retryability {
	if err == nil || errors.Is(err, context.Canceled) {
		return provider.NotSafeToRetry
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		if _, throttled := retry.DefaultThrottleErrorCodes[apiErr.ErrorCode()]; throttled {
			return provider.SafeToRetry // refused before acting
		}
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return provider.SafeToRetry // no connection, so nothing was sent
	}
	var respErr *smithyhttp.ResponseError
	if errors.As(err, &respErr) && respErr.HTTPStatusCode() >= 500 {
		return provider.ConditionallyRetryable // may have acted: no client token tells us otherwise
	}
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.As(err, &netErr) {
		return provider.ConditionallyRetryable // sent, answer lost
	}
	return provider.NotSafeToRetry
}

// hasCode reports whether err carries one of the given AWS error codes.
func hasCode(err error, codes ...string) bool {
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		return false
	}
	for _, c := range codes {
		if apiErr.ErrorCode() == c {
			return true
		}
	}
	return false
}

// apiFailure is the message a user reads in a failed apply, with the SDK error kept underneath for
// classification.
type apiFailure struct {
	msg string
	err error
}

func (e *apiFailure) Error() string { return e.msg }
func (e *apiFailure) Unwrap() error { return e.err }

// failed says which instance, which call, where, what AWS said, and the request ID AWS support will
// ask for. Never anything from the request itself, which may hold a secret.
func (p *Provider) failed(op, region, awsID string, err error) error {
	where := region
	if awsID != "" {
		where = formatID(region, awsID)
	}
	msg := fmt.Sprintf("aws instance %q: EC2 %s for %s failed: ", p.instance, op, where)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		msg += apiErr.ErrorCode() + ": " + apiErr.ErrorMessage()
	} else {
		msg += err.Error()
	}
	var re *awshttp.ResponseError
	if errors.As(err, &re) && re.ServiceRequestID() != "" {
		msg += " (request ID " + re.ServiceRequestID() + ")"
	}
	switch {
	case hasCode(err, "UnauthorizedOperation", "AuthFailure"):
		msg += fmt.Sprintf("\ncheck the IAM permissions of the credentials instance %q uses", p.instance)
	case hasCode(err, "DependencyViolation"):
		msg += "\nsomething infrata does not manage still depends on it (subnets, network interfaces, gateways): remove that first"
	}
	return &apiFailure{msg: msg, err: err}
}
```

In `provider.go`, replace `ClassifyError`:

```go
// ClassifyError delegates to classify, a pure function of the error.
func (p *Provider) ClassifyError(err error) provider.Retryability { return classify(err) }
```

- [ ] **Step 4: Run**

Run: `go test -count=1 ./... && go vet ./...`
Expected: PASS. If "dropped after acting" classifies as `NotSafeToRetry`, print the error chain
(`fmt.Printf("%#v", err)`) and add the concrete type it carries to the `ConditionallyRetryable` branch; record
what it was in the Verification log.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: move the throttle check after the 5xx check ("throttle" case fails); drop the dial case (refused
connection → `ConditionallyRetryable`); classify every `smithy.APIError` as `SafeToRetry` ("validation" fails);
remove `Unwrap` from `apiFailure` (message test's `hasCode` fails).

```bash
git add internal/awsprov/ids.go internal/awsprov/errors.go internal/awsprov/provider.go \
  internal/awsprov/ids_test.go internal/awsprov/errors_test.go
git commit -m "aws: classify SDK errors by what AWS may have done, and say where they happened" -- <same paths>
```

---
### Task 5: `aws.vpc` — create, read, update, delete

**Files:**
- Create: `internal/awsprov/vpc.go`, `internal/awsprov/tags.go`, `internal/awsprov/attrs.go`
- Modify: `internal/awsprov/patience.go` (add `wait`), `internal/awsprov/provider.go` (dispatch)
- Test: `internal/awsprov/vpc_test.go`, `internal/awsprov/tags_test.go`, `internal/awsprov/helpers_test.go`

**Interfaces:**
- Consumes: `clients.ec2`, `createOnce`, `parseID`/`formatID`, `p.failed`, `hasCode`, `patience`, `str`/`boolean`.
- Produces:
  - `func (pt patience) wait(ctx context.Context, try func() (found bool, err error)) (bool, error)`
  - `var once = patience{attempts: 1}` (no waiting; used by Import)
  - `func stringAttr(attrs map[string]value.Value, name string) (string, error)`, `func boolAttr(attrs map[string]value.Value, name string) (bool, error)`
  - tags: `tagsFrom(attrs) (map[string]string, error)`, `fromAWSTags([]types.Tag) map[string]string`,
    `putTags(attrs map[string]value.Value, tags map[string]string)`, `toAWSTags(map[string]string) []types.Tag`,
    `tagSpecs(types.ResourceType, map[string]string) []types.TagSpecification`,
    `func (p *Provider) syncTags(ctx context.Context, region, awsID string, current, desired map[string]value.Value) error`
  - vpc: `createVPC`, `readVPC(ctx, providerID string, pt patience)`, `updateVPC`, `deleteVPC`, `vpcState(region string, v types.Vpc) *resource.ResourceState`
  - test helpers: `fakeProvider(t, values) (*Provider, *ec2fake.Server)`, `desired(typ string, attrs map[string]value.Value) *resource.DesiredResource`, `tagsValue(kv ...string) value.Value`

- [ ] **Step 1: Write the test helpers and failing tests**

`internal/awsprov/helpers_test.go`:

```go
package awsprov

import (
	"context"
	"testing"
	"time"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// fakeProvider configures a real instance — LoadDefaultConfig and all — against a fresh fake, with
// the eventual-consistency backoff made instant. Its attempt count is kept, so tests can count calls.
func fakeProvider(t *testing.T, values map[string]value.Value) (*Provider, *ec2fake.Server) {
	t.Helper()
	fake := ec2fake.New()
	t.Cleanup(fake.Close)
	isolateAWS(t, fake.URL)
	prov, err := NewPlugin().New(provider.Config{Instance: "test", Values: values})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	p := prov.(*Provider)
	p.patience.sleep = func(context.Context, time.Duration) error { return nil }
	return p, fake
}

func desired(typ string, attrs map[string]value.Value) *resource.DesiredResource {
	return &resource.DesiredResource{Address: address.Address{Name: "r"}, Type: typ, Attrs: attrs}
}

func tagsValue(kv ...string) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i]] = value.String(kv[i+1], value.SourceExplicit)
	}
	return value.Map(items, value.SourceExplicit)
}

func vpcAttrs(extra map[string]value.Value) map[string]value.Value {
	attrs := map[string]value.Value{"region": s("us-east-1"), "cidr": s("10.0.0.0/16")}
	for k, v := range extra {
		attrs[k] = v
	}
	return attrs
}

func tagsOf(t *testing.T, st *resource.ResourceState) map[string]string {
	t.Helper()
	got, err := tagsFrom(st.Attributes)
	if err != nil {
		t.Fatal(err)
	}
	return got
}
```

`internal/awsprov/vpc_test.go`:

```go
package awsprov

import (
	"context"
	"maps"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func TestAVPCIsCreatedWithItsTagsAndReadBack(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("team", "platform")})))
	if err != nil {
		t.Fatal(err)
	}
	vpcs := fake.VPCs()
	if len(vpcs) != 1 || vpcs[0].Tags["team"] != "platform" {
		t.Fatalf("fake holds %+v", vpcs)
	}
	if st.ProviderID != "us-east-1/"+vpcs[0].ID || st.Type != typeVPC {
		t.Errorf("state = %s %s", st.Type, st.ProviderID)
	}
	if id, _ := st.Attributes["id"].AsString(); id != vpcs[0].ID {
		t.Errorf("id = %q", id)
	}
	if owner, _ := st.Attributes["owner_id"].AsString(); owner != ec2fake.Owner {
		t.Errorf("owner_id = %q", owner)
	}
	if st.Address.Name != "" || st.Provider != "" {
		t.Error("bookkeeping is the host's: the plugin must leave Address and Provider unset")
	}
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	if !maps.Equal(tagsOf(t, got), map[string]string{"team": "platform"}) {
		t.Errorf("read tags = %v", tagsOf(t, got))
	}
	if cidr, _ := got.Attributes["cidr"].AsString(); cidr != "10.0.0.0/16" {
		t.Errorf("cidr = %q", cidr)
	}
}

// TestACreateIsSentOnceEvenWhenTheAnswerIsLost. AWS_MAX_ATTEMPTS=5 makes the SDK want to retry: the
// assertion is that it does not, and that the one VPC that exists is exactly one.
func TestACreateIsSentOnceEvenWhenTheAnswerIsLost(t *testing.T) {
	for name, fault := range map[string]ec2fake.Fault{
		"connection dropped after acting": {Drop: true},
		"server fault":                    {Status: 500, Code: "InternalError", Message: "oops"},
	} {
		t.Run(name, func(t *testing.T) {
			fake := ec2fake.New()
			defer fake.Close()
			isolateAWS(t, fake.URL)
			t.Setenv("AWS_MAX_ATTEMPTS", "5")
			prov, err := NewPlugin().New(provider.Config{Instance: "test"})
			if err != nil {
				t.Fatal(err)
			}
			fault.Action, fault.Nth = "CreateVpc", 1
			fake.Inject(fault)
			_, err = prov.Create(context.Background(), desired(typeVPC, vpcAttrs(nil)))
			if err == nil {
				t.Fatal("expected the create to fail")
			}
			if n := fake.Calls("CreateVpc"); n != 1 {
				t.Errorf("CreateVpc sent %d times; a resent create makes an untracked VPC", n)
			}
			if got := prov.ClassifyError(err); got != provider.ConditionallyRetryable {
				t.Errorf("classification = %v, want ConditionallyRetryable so infrata does not retry the create", got)
			}
		})
	}
}

func TestAReadWaitsForAVPCEC2HasNotPropagated(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.HideFromDescribe(fake.VPCs()[0].ID, 3)
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read of a just-created VPC = %v, %v; reporting it gone would plan a second VPC", got, err)
	}
	if n := fake.Calls("DescribeVpcs"); n != 4 {
		t.Errorf("DescribeVpcs calls = %d, want 4 (three hidden, then found)", n)
	}
}

func TestAReadReportsAVPCGoneOnlyAfterItsPatienceRunsOut(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.Remove(fake.VPCs()[0].ID)
	got, err := p.Read(ctx, st)
	if err != nil || got != nil {
		t.Fatalf("Read of a deleted VPC = %v, %v; want (nil, nil)", got, err)
	}
	if n := fake.Calls("DescribeVpcs"); n != notFoundPatience.attempts {
		t.Errorf("DescribeVpcs calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

// TestUpdateRemovesTagsConfigurationDropped. Merging instead of replacing leaves `a` in place and
// every later plan proposes removing it again.
func TestUpdateRemovesTagsConfigurationDropped(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("a", "1", "b", "2")})))
	if err != nil {
		t.Fatal(err)
	}
	st, err = p.Update(ctx, st, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("b", "3", "c", "4")})))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"b": "3", "c": "4"}
	if got := fake.VPCs()[0].Tags; !maps.Equal(got, want) {
		t.Errorf("AWS tags = %v, want %v", got, want)
	}
	if got := tagsOf(t, st); !maps.Equal(got, want) {
		t.Errorf("returned tags = %v, want %v", got, want)
	}
	if _, ok := st.Attributes["id"]; !ok {
		t.Error("Update dropped a computed attribute: desired never carries them")
	}
	st, err = p.Update(ctx, st, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := st.Attributes["tags"]; ok || len(fake.VPCs()[0].Tags) != 0 {
		t.Errorf("after removing tags: state tags present=%v, AWS tags=%v", ok, fake.VPCs()[0].Tags)
	}
}

// TestAWSReservedTagsAreNeitherReportedNorAccepted. Reporting them plans their removal forever;
// AWS refuses to remove them.
func TestAWSReservedTagsAreNeitherReportedNorAccepted(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	fake.SetTag(fake.VPCs()[0].ID, "aws:cloudformation:stack-name", "legacy")
	got, err := p.Read(ctx, st)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := got.Attributes["tags"]; ok {
		t.Errorf("reserved tags were reported: %v", got.Attributes["tags"])
	}
	_, err = p.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("aws:owner", "me")})))
	if err == nil || !strings.Contains(err.Error(), "aws:owner") {
		t.Fatalf("create with a reserved tag = %v", err)
	}
	if n := fake.Calls("CreateVpc"); n != 1 {
		t.Errorf("CreateVpc calls = %d: a reserved tag must be refused before any call", n)
	}
}

func TestDeletingAVPCThatIsAlreadyGoneSucceeds(t *testing.T) {
	p, _ := fakeProvider(t, nil)
	ctx := context.Background()
	st, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := p.Delete(ctx, st); err != nil {
			t.Fatalf("delete %d: %v", i+1, err)
		}
	}
}

func TestACancelledCreateSendsNothing(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil))); err == nil {
		t.Fatal("a cancelled create succeeded")
	}
	if n := fake.Calls("CreateVpc"); n != 0 {
		t.Errorf("CreateVpc calls = %d, want 0", n)
	}
}

func TestReadRefusesAnIDOfAnotherType(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	_, err := p.Read(context.Background(), &resource.ResourceState{Type: typeVPC, ProviderID: "us-east-1/subnet-1"})
	if err == nil || fake.Calls("DescribeVpcs") != 0 {
		t.Fatalf("Read = %v after %d calls", err, fake.Calls("DescribeVpcs"))
	}
}
```

`internal/awsprov/tags_test.go`:

```go
package awsprov

import (
	"testing"

	"github.com/infrata/infrata/pkg/value"
)

func TestTagsMustBeAMapOfStrings(t *testing.T) {
	for name, v := range map[string]value.Value{
		"not a map":          s("team=platform"),
		"non-string value":   value.Map(map[string]value.Value{"n": value.Int(1, value.SourceExplicit)}, value.SourceExplicit),
		"reserved aws: key":  tagsValue("aws:x", "y"),
	} {
		if _, err := tagsFrom(map[string]value.Value{"tags": v}); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 -run 'VPC|Create|Read|Update|Reserved|Delet|Tags' ./internal/awsprov/`
Expected: FAIL — `undefined: tagsFrom` and friends.

- [ ] **Step 3: Implement**

`internal/awsprov/attrs.go`:

```go
package awsprov

import (
	"fmt"

	"github.com/infrata/infrata/pkg/value"
)

// stringAttr reads a required string. infrata has already checked Required and Kind at compile time;
// this is the plugin refusing to guess if a caller skipped that.
func stringAttr(attrs map[string]value.Value, name string) (string, error) {
	v, ok := attrs[name]
	if !ok {
		return "", fmt.Errorf("attribute %q is missing", name)
	}
	text, isString := v.AsString()
	if !isString || text == "" {
		return "", fmt.Errorf("attribute %q must be a non-empty string, got %s", name, v.Kind)
	}
	return text, nil
}

// boolAttr reads a bool, absent meaning false (the schema's default).
func boolAttr(attrs map[string]value.Value, name string) (bool, error) {
	v, ok := attrs[name]
	if !ok {
		return false, nil
	}
	b, isBool := v.AsBool()
	if !isBool {
		return false, fmt.Errorf("attribute %q must be a boolean, got %s", name, v.Kind)
	}
	return b, nil
}
```

`internal/awsprov/tags.go`:

```go
package awsprov

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/value"
)

// reservedPrefix is AWS's own: tags under it cannot be edited or deleted, so they are neither
// reported (every plan would propose removing them) nor accepted.
const reservedPrefix = "aws:"

func tagsFrom(attrs map[string]value.Value) (map[string]string, error) {
	v, ok := attrs["tags"]
	if !ok {
		return map[string]string{}, nil
	}
	items, isMap := v.Raw.(map[string]value.Value)
	if v.Kind != value.KindMap || !isMap {
		return nil, fmt.Errorf("tags must be a map of strings, got %s", v.Kind)
	}
	out := make(map[string]string, len(items))
	for k, item := range items {
		if strings.HasPrefix(k, reservedPrefix) {
			return nil, fmt.Errorf("tag %q: keys starting %q are reserved by AWS and cannot be set", k, reservedPrefix)
		}
		text, isString := item.AsString()
		if !isString {
			return nil, fmt.Errorf("tag %q must be a string, got %s", k, item.Kind)
		}
		out[k] = text
	}
	return out, nil
}

func fromAWSTags(tags []types.Tag) map[string]string {
	out := map[string]string{}
	for _, t := range tags {
		if k := aws.ToString(t.Key); !strings.HasPrefix(k, reservedPrefix) {
			out[k] = aws.ToString(t.Value)
		}
	}
	return out
}

// putTags sets `tags`, or removes it when there are none: an empty map where configuration has no
// `tags:` would plan "removed from configuration" on every run.
func putTags(attrs map[string]value.Value, tags map[string]string) {
	if len(tags) == 0 {
		delete(attrs, "tags")
		return
	}
	items := make(map[string]value.Value, len(tags))
	for k, v := range tags {
		items[k] = str(v)
	}
	attrs["tags"] = value.Map(items, value.SourceProvider)
}

func toAWSTags(tags map[string]string) []types.Tag {
	keys := make([]string, 0, len(tags))
	for k := range tags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]types.Tag, 0, len(keys))
	for _, k := range keys {
		out = append(out, types.Tag{Key: aws.String(k), Value: aws.String(tags[k])})
	}
	return out
}

// tagSpecs tags a resource in its create call, which AWS applies atomically: a resource is created
// with its tags or not at all. No follow-up CreateTags, so no window where it exists untagged.
func tagSpecs(rt types.ResourceType, tags map[string]string) []types.TagSpecification {
	if len(tags) == 0 {
		return nil
	}
	return []types.TagSpecification{{ResourceType: rt, Tags: toAWSTags(tags)}}
}

// syncTags makes AWS's tags equal desired's: sets what is new or changed, deletes what is gone.
func (p *Provider) syncTags(ctx context.Context, region, awsID string, current, desired map[string]value.Value) error {
	have, err := tagsFrom(current)
	if err != nil {
		return err
	}
	want, err := tagsFrom(desired)
	if err != nil {
		return err
	}
	set := map[string]string{}
	for k, v := range want {
		if old, ok := have[k]; !ok || old != v {
			set[k] = v
		}
	}
	var remove []types.Tag
	for k := range have {
		if _, ok := want[k]; !ok {
			remove = append(remove, types.Tag{Key: aws.String(k)})
		}
	}
	sort.Slice(remove, func(i, j int) bool { return *remove[i].Key < *remove[j].Key })
	client := p.clients.ec2(region)
	if len(set) > 0 {
		if _, err := client.CreateTags(ctx, &ec2.CreateTagsInput{Resources: []string{awsID}, Tags: toAWSTags(set)}); err != nil {
			return p.failed("CreateTags", region, awsID, err)
		}
	}
	if len(remove) > 0 {
		if _, err := client.DeleteTags(ctx, &ec2.DeleteTagsInput{Resources: []string{awsID}, Tags: remove}); err != nil {
			return p.failed("DeleteTags", region, awsID, err)
		}
	}
	return nil
}
```

Append to `internal/awsprov/patience.go`:

```go
// once reads a single time: Import, where the ID came from discovery and so is already visible.
var once = patience{attempts: 1}

// wait calls try until it reports found, returns an error, or the attempts run out — the last
// answering (false, nil). An error from try stops at once: only absence is worth waiting out.
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
```

`internal/awsprov/vpc.go`:

```go
package awsprov

import (
	"context"
	"fmt"
	"maps"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func vpcState(region string, v types.Vpc) *resource.ResourceState {
	id := aws.ToString(v.VpcId)
	attrs := map[string]value.Value{
		"region":     str(region),
		"cidr":       str(aws.ToString(v.CidrBlock)),
		"id":         str(id),
		"owner_id":   str(aws.ToString(v.OwnerId)),
		"is_default": boolean(aws.ToBool(v.IsDefault)),
	}
	putTags(attrs, fromAWSTags(v.Tags))
	return &resource.ResourceState{Type: typeVPC, ProviderID: formatID(region, id), Attributes: attrs}
}

func (p *Provider) createVPC(ctx context.Context, attrs map[string]value.Value) (*resource.ResourceState, error) {
	region, err := stringAttr(attrs, "region")
	if err != nil {
		return nil, err
	}
	cidr, err := stringAttr(attrs, "cidr")
	if err != nil {
		return nil, err
	}
	tags, err := tagsFrom(attrs)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err // nothing sent, nothing lost
	}
	// Once sent, the answer is read whatever happens to ctx: a cancelled request loses the answer,
	// not the VPC.
	out, err := p.clients.ec2(region).CreateVpc(context.WithoutCancel(ctx), &ec2.CreateVpcInput{
		CidrBlock:         aws.String(cidr),
		TagSpecifications: tagSpecs(types.ResourceTypeVpc, tags),
	}, createOnce)
	if err != nil {
		return nil, p.failed("CreateVpc", region, "", err)
	}
	if out.Vpc == nil || aws.ToString(out.Vpc.VpcId) == "" {
		return nil, fmt.Errorf("aws instance %q: CreateVpc in %s succeeded but returned no VPC ID, so infrata cannot record it: "+
			"look for a VPC with CIDR %s in %s and import it", p.instance, region, cidr, region)
	}
	vpc := *out.Vpc
	// From what was sent, not a describe: tags in a create call are applied atomically, and a
	// describe made now may not see the VPC yet.
	vpc.Tags = toAWSTags(tags)
	return vpcState(region, vpc), nil
}

func (p *Provider) describeVPC(ctx context.Context, region, awsID string) (*types.Vpc, error) {
	out, err := p.clients.ec2(region).DescribeVpcs(ctx, &ec2.DescribeVpcsInput{VpcIds: []string{awsID}})
	if hasCode(err, "InvalidVpcID.NotFound") {
		return nil, nil
	}
	if err != nil {
		return nil, p.failed("DescribeVpcs", region, awsID, err)
	}
	if len(out.Vpcs) == 0 {
		return nil, nil
	}
	return &out.Vpcs[0], nil
}

func (p *Provider) readVPC(ctx context.Context, providerID string, pt patience) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeVPC, providerID)
	if err != nil {
		return nil, err
	}
	var vpc *types.Vpc
	found, err := pt.wait(ctx, func() (bool, error) {
		var err error
		vpc, err = p.describeVPC(ctx, region, awsID)
		return vpc != nil, err
	})
	if err != nil || !found {
		return nil, err
	}
	return vpcState(region, *vpc), nil
}

// updateVPC can only be a tag change: region and cidr are ForceNew, so a plan replaces instead.
func (p *Provider) updateVPC(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeVPC, current.ProviderID)
	if err != nil {
		return nil, err
	}
	if err := p.syncTags(ctx, region, awsID, current.Attributes, desired.Attrs); err != nil {
		return nil, err
	}
	want, _ := tagsFrom(desired.Attrs) // already validated by syncTags
	attrs := maps.Clone(current.Attributes)
	putTags(attrs, want)
	return &resource.ResourceState{Type: typeVPC, ProviderID: current.ProviderID, Attributes: attrs}, nil
}

func (p *Provider) deleteVPC(ctx context.Context, current *resource.ResourceState) error {
	region, awsID, err := parseID(typeVPC, current.ProviderID)
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	_, err = p.clients.ec2(region).DeleteVpc(context.WithoutCancel(ctx), &ec2.DeleteVpcInput{VpcId: aws.String(awsID)})
	if hasCode(err, "InvalidVpcID.NotFound") {
		return nil // already gone is the outcome a delete asks for
	}
	if err != nil {
		return p.failed("DeleteVpc", region, awsID, err)
	}
	return nil
}
```

In `provider.go`, replace the four CRUD stubs with type dispatch (Task 6 adds the subnet cases; `Discover` and
`Import` stay stubs until Task 7):

```go
func unknownType(t string) error { return fmt.Errorf("the aws plugin does not serve %q", t) }

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	switch current.Type {
	case typeVPC:
		return p.readVPC(ctx, current.ProviderID, p.patience)
	}
	return nil, unknownType(current.Type)
}

func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	switch desired.Type {
	case typeVPC:
		return p.createVPC(ctx, desired.Attrs)
	}
	return nil, unknownType(desired.Type)
}

func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	switch current.Type {
	case typeVPC:
		return p.updateVPC(ctx, current, desired)
	}
	return nil, unknownType(current.Type)
}

func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	switch current.Type {
	case typeVPC:
		return p.deleteVPC(ctx, current)
	}
	return unknownType(current.Type)
}
```

(add `"fmt"` to `provider.go`'s imports.)

- [ ] **Step 4: Run**

Run: `go test -count=1 ./... && go vet ./... && gofmt -l .`
Expected: PASS.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: drop `createOnce` from `CreateVpc` (`TestACreateIsSentOnce…` — 5 calls under `AWS_MAX_ATTEMPTS=5`);
`patience.wait` returns after the first miss (`TestAReadWaitsFor…`); `syncTags` skips `DeleteTags`
(`TestUpdateRemovesTags…`); `fromAWSTags` keeps `aws:` keys (`TestAWSReservedTags…`); `deleteVPC` returns
the NotFound error (`TestDeletingAVPC…`); remove the `ctx.Err()` check (`TestACancelledCreate…`).

```bash
git add internal/awsprov/vpc.go internal/awsprov/tags.go internal/awsprov/attrs.go internal/awsprov/patience.go \
  internal/awsprov/provider.go internal/awsprov/vpc_test.go internal/awsprov/tags_test.go internal/awsprov/helpers_test.go
git commit -m "aws.vpc: a create sent once, a read that waits out propagation, tags that converge" -- <same paths>
```

---
### Task 6: `aws.subnet` — the edge, a non-tag update, and cross-resource consistency

**Files:**
- Create: `internal/awsprov/subnet.go`
- Modify: `internal/awsprov/provider.go` (subnet cases in the four dispatch switches)
- Test: `internal/awsprov/subnet_test.go`

**Interfaces:**
- Consumes: everything Task 5 produced.
- Produces: `createSubnet`, `readSubnet(ctx, providerID string, pt patience)`, `updateSubnet`, `deleteSubnet`,
  `subnetState(region string, sub types.Subnet) *resource.ResourceState`.

- [ ] **Step 1: Write the failing tests**

`internal/awsprov/subnet_test.go`:

```go
package awsprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrata/infrata-provider-aws/internal/ec2fake"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func subnetAttrs(vpcID string, extra map[string]value.Value) map[string]value.Value {
	attrs := map[string]value.Value{
		"region": s("us-east-1"), "vpc_id": s(vpcID), "cidr": s("10.0.1.0/24"),
		"availability_zone": s("us-east-1a"), "map_public_ip_on_launch": value.Bool(false, value.SourceDefault),
	}
	for k, v := range extra {
		attrs[k] = v
	}
	return attrs
}

func TestASubnetIsCreatedInItsVPCAndReadBack(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	st, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(vpc, map[string]value.Value{"tags": tagsValue("tier", "private")})))
	if err != nil {
		t.Fatal(err)
	}
	subs := fake.Subnets()
	if len(subs) != 1 || subs[0].VPCID != vpc || subs[0].AZ != "us-east-1a" || subs[0].Tags["tier"] != "private" {
		t.Fatalf("fake holds %+v", subs)
	}
	if st.ProviderID != "us-east-1/"+subs[0].ID {
		t.Errorf("provider ID = %q", st.ProviderID)
	}
	if arn, _ := st.Attributes["arn"].AsString(); !strings.HasSuffix(arn, ":subnet/"+subs[0].ID) {
		t.Errorf("arn = %q", arn)
	}
	got, err := p.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read = %v, %v", got, err)
	}
	for _, name := range []string{"region", "vpc_id", "cidr", "availability_zone", "map_public_ip_on_launch", "tags", "id", "arn", "owner_id"} {
		if !got.Attributes[name].Equal(st.Attributes[name]) {
			t.Errorf("%s: read %v, created %v — a difference here is a change on every plan", name, got.Attributes[name], st.Attributes[name])
		}
	}
}

// TestASubnetWaitsForAVPCCreatedAMomentAgo. CreateSubnet refusing an unpropagated VPC means nothing
// was created, so retrying it inside the plugin is safe — unlike retrying after an ambiguous failure.
func TestASubnetWaitsForAVPCCreatedAMomentAgo(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	fake.HideFromDescribe(vpc, 2) // CreateSubnet consults the same visibility
	if _, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, nil))); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if n := fake.Calls("CreateSubnet"); n != 3 {
		t.Errorf("CreateSubnet calls = %d, want 3 (two refusals, one success)", n)
	}
	if len(fake.Subnets()) != 1 {
		t.Errorf("subnets = %d, want exactly 1", len(fake.Subnets()))
	}
}

func TestASubnetInAVPCThatReallyDoesNotExistFailsNamingIt(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	_, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs("vpc-0000000000000dead", nil)))
	if err == nil || !strings.Contains(err.Error(), "vpc-0000000000000dead") || !strings.Contains(err.Error(), "us-east-1") {
		t.Fatalf("err = %v; want it to name the VPC and the region", err)
	}
	if n := fake.Calls("CreateSubnet"); n != notFoundPatience.attempts {
		t.Errorf("CreateSubnet calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

func TestMapPublicIPIsSetOnCreateAndChangedInPlace(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	on := map[string]value.Value{"map_public_ip_on_launch": value.Bool(true, value.SourceExplicit)}
	st, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(vpc, on)))
	if err != nil {
		t.Fatal(err)
	}
	if !fake.Subnets()[0].MapPublicIP {
		t.Fatal("map_public_ip_on_launch was not applied on create")
	}
	st, err = p.Update(ctx, st, desired(typeSubnet, subnetAttrs(vpc, nil)))
	if err != nil {
		t.Fatal(err)
	}
	if fake.Subnets()[0].MapPublicIP {
		t.Error("update did not turn map_public_ip_on_launch off")
	}
	if b, _ := st.Attributes["map_public_ip_on_launch"].AsBool(); b {
		t.Error("returned state still says true")
	}
}

// TestAFailedAttributeAfterCreateStillReportsTheSubnet. infrata drops the result of an errored
// create, so an error here would leave a real subnet untracked.
func TestAFailedAttributeAfterCreateStillReportsTheSubnet(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	for n := 1; n <= 3; n++ { // outlast the SDK's own retries
		fake.Inject(ec2fake.Fault{Action: "ModifySubnetAttribute", Nth: n, Status: 400, Code: "UnauthorizedOperation", Message: "no"})
	}
	on := map[string]value.Value{"map_public_ip_on_launch": value.Bool(true, value.SourceExplicit)}
	st, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, on)))
	if err != nil || st == nil {
		t.Fatalf("Create = %v, %v; a subnet that exists must be reported", st, err)
	}
	if b, _ := st.Attributes["map_public_ip_on_launch"].AsBool(); b {
		t.Error("state claims an attribute AWS refused; it must be the truth, so the next plan converges")
	}
}

func TestDeletingASubnetThatIsAlreadyGoneSucceeds(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	vpc := fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	st, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, nil)))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := p.Delete(context.Background(), st); err != nil {
			t.Fatalf("delete %d: %v", i+1, err)
		}
	}
}

// TestAVPCWithASubnetRefusesDeletionWithAHint. infrata orders destroys by reference, so this only
// happens when something outside infrata sits in the VPC.
func TestAVPCWithASubnetRefusesDeletionWithAHint(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpcState, err := p.Create(ctx, desired(typeVPC, vpcAttrs(nil)))
	if err != nil {
		t.Fatal(err)
	}
	id, _ := vpcState.Attributes["id"].AsString()
	if _, err := p.Create(ctx, desired(typeSubnet, subnetAttrs(id, nil))); err != nil {
		t.Fatal(err)
	}
	err = p.Delete(ctx, &resource.ResourceState{Type: typeVPC, ProviderID: vpcState.ProviderID})
	if err == nil || !strings.Contains(err.Error(), "DependencyViolation") || !strings.Contains(err.Error(), "depends on it") {
		t.Fatalf("err = %v", err)
	}
	if len(fake.VPCs()) != 1 {
		t.Error("the VPC was deleted")
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 -run 'Subnet|MapPublic|FailedAttribute|VPCWithASubnet' ./internal/awsprov/`
Expected: FAIL — `the aws plugin does not serve "aws.subnet"`.

- [ ] **Step 3: Implement**

`internal/awsprov/subnet.go`:

```go
package awsprov

import (
	"context"
	"fmt"
	"maps"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

func subnetState(region string, sub types.Subnet) *resource.ResourceState {
	id := aws.ToString(sub.SubnetId)
	attrs := map[string]value.Value{
		"region":                  str(region),
		"vpc_id":                  str(aws.ToString(sub.VpcId)),
		"cidr":                    str(aws.ToString(sub.CidrBlock)),
		"availability_zone":       str(aws.ToString(sub.AvailabilityZone)),
		"map_public_ip_on_launch": boolean(aws.ToBool(sub.MapPublicIpOnLaunch)),
		"id":                      str(id),
		"arn":                     str(aws.ToString(sub.SubnetArn)),
		"owner_id":                str(aws.ToString(sub.OwnerId)),
	}
	putTags(attrs, fromAWSTags(sub.Tags))
	return &resource.ResourceState{Type: typeSubnet, ProviderID: formatID(region, id), Attributes: attrs}
}

func (p *Provider) createSubnet(ctx context.Context, attrs map[string]value.Value) (*resource.ResourceState, error) {
	var in struct{ region, vpcID, cidr, az string }
	for name, dst := range map[string]*string{"region": &in.region, "vpc_id": &in.vpcID, "cidr": &in.cidr, "availability_zone": &in.az} {
		v, err := stringAttr(attrs, name)
		if err != nil {
			return nil, err
		}
		*dst = v
	}
	public, err := boolAttr(attrs, "map_public_ip_on_launch")
	if err != nil {
		return nil, err
	}
	tags, err := tagsFrom(attrs)
	if err != nil {
		return nil, err
	}
	client := p.clients.ec2(in.region)

	var out *ec2.CreateSubnetOutput
	var lastErr error
	// A VPC created a moment ago may not be visible to CreateSubnet yet. A NotFound here is a
	// refusal — nothing was created — so waiting it out is safe in a way retrying an ambiguous
	// failure is not.
	found, err := p.patience.wait(ctx, func() (bool, error) {
		if err := ctx.Err(); err != nil {
			return false, err
		}
		var err error
		out, err = client.CreateSubnet(context.WithoutCancel(ctx), &ec2.CreateSubnetInput{
			VpcId: aws.String(in.vpcID), CidrBlock: aws.String(in.cidr), AvailabilityZone: aws.String(in.az),
			TagSpecifications: tagSpecs(types.ResourceTypeSubnet, tags),
		}, createOnce)
		if hasCode(err, "InvalidVpcID.NotFound") {
			lastErr = err
			return false, nil
		}
		return err == nil, err
	})
	if err != nil {
		return nil, p.failed("CreateSubnet", in.region, "", err)
	}
	if !found {
		return nil, p.failed("CreateSubnet", in.region, "", fmt.Errorf("VPC %s does not exist in %s: %w", in.vpcID, in.region, lastErr))
	}
	if out.Subnet == nil || aws.ToString(out.Subnet.SubnetId) == "" {
		return nil, fmt.Errorf("aws instance %q: CreateSubnet in %s succeeded but returned no subnet ID, so infrata cannot record it: "+
			"look for subnet %s in VPC %s and import it", p.instance, in.region, in.cidr, in.vpcID)
	}
	sub := *out.Subnet
	sub.Tags = toAWSTags(tags)
	sub.MapPublicIpOnLaunch = aws.Bool(false) // what AWS created, until the call below says otherwise

	if public {
		// The subnet exists now. However this call ends, the subnet is reported: an error from a
		// create is dropped by the host, and the subnet would be untracked.
		if err := p.setMapPublicIP(ctx, in.region, aws.ToString(sub.SubnetId), true); err != nil {
			fmt.Fprintf(os.Stderr, "created %s but could not set map_public_ip_on_launch; the next plan will propose it: %v\n",
				formatID(in.region, aws.ToString(sub.SubnetId)), err)
		} else {
			sub.MapPublicIpOnLaunch = aws.Bool(true)
		}
	}
	return subnetState(in.region, sub), nil
}

// setMapPublicIP waits out a subnet ID that has not propagated, as a just-created one may not have.
func (p *Provider) setMapPublicIP(ctx context.Context, region, awsID string, on bool) error {
	client := p.clients.ec2(region)
	var lastErr error
	found, err := p.patience.wait(ctx, func() (bool, error) {
		_, err := client.ModifySubnetAttribute(ctx, &ec2.ModifySubnetAttributeInput{
			SubnetId: aws.String(awsID), MapPublicIpOnLaunch: &types.AttributeBooleanValue{Value: aws.Bool(on)},
		})
		if hasCode(err, "InvalidSubnetID.NotFound") {
			lastErr = err
			return false, nil
		}
		return err == nil, err
	})
	if err == nil && !found {
		err = lastErr
	}
	if err != nil {
		return p.failed("ModifySubnetAttribute", region, awsID, err)
	}
	return nil
}

func (p *Provider) describeSubnet(ctx context.Context, region, awsID string) (*types.Subnet, error) {
	out, err := p.clients.ec2(region).DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{SubnetIds: []string{awsID}})
	if hasCode(err, "InvalidSubnetID.NotFound") {
		return nil, nil
	}
	if err != nil {
		return nil, p.failed("DescribeSubnets", region, awsID, err)
	}
	if len(out.Subnets) == 0 {
		return nil, nil
	}
	return &out.Subnets[0], nil
}

func (p *Provider) readSubnet(ctx context.Context, providerID string, pt patience) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeSubnet, providerID)
	if err != nil {
		return nil, err
	}
	var sub *types.Subnet
	found, err := pt.wait(ctx, func() (bool, error) {
		var err error
		sub, err = p.describeSubnet(ctx, region, awsID)
		return sub != nil, err
	})
	if err != nil || !found {
		return nil, err
	}
	return subnetState(region, *sub), nil
}

// updateSubnet changes what can change in place: map_public_ip_on_launch and tags.
func (p *Provider) updateSubnet(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	region, awsID, err := parseID(typeSubnet, current.ProviderID)
	if err != nil {
		return nil, err
	}
	attrs := maps.Clone(current.Attributes)
	have, _ := boolAttr(current.Attributes, "map_public_ip_on_launch")
	want, err := boolAttr(desired.Attrs, "map_public_ip_on_launch")
	if err != nil {
		return nil, err
	}
	if have != want {
		if err := p.setMapPublicIP(ctx, region, awsID, want); err != nil {
			return nil, err
		}
		attrs["map_public_ip_on_launch"] = boolean(want)
	}
	if err := p.syncTags(ctx, region, awsID, current.Attributes, desired.Attrs); err != nil {
		return nil, err
	}
	tags, _ := tagsFrom(desired.Attrs)
	putTags(attrs, tags)
	return &resource.ResourceState{Type: typeSubnet, ProviderID: current.ProviderID, Attributes: attrs}, nil
}

func (p *Provider) deleteSubnet(ctx context.Context, current *resource.ResourceState) error {
	region, awsID, err := parseID(typeSubnet, current.ProviderID)
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	_, err = p.clients.ec2(region).DeleteSubnet(context.WithoutCancel(ctx), &ec2.DeleteSubnetInput{SubnetId: aws.String(awsID)})
	if hasCode(err, "InvalidSubnetID.NotFound") {
		return nil
	}
	if err != nil {
		return p.failed("DeleteSubnet", region, awsID, err)
	}
	return nil
}
```

In `provider.go` add a `case typeSubnet:` to each switch: `return p.readSubnet(ctx, current.ProviderID, p.patience)`,
`return p.createSubnet(ctx, desired.Attrs)`, `return p.updateSubnet(ctx, current, desired)`,
`return p.deleteSubnet(ctx, current)`.

- [ ] **Step 4: Run**

Run: `go test -count=1 ./... && go vet ./... && gofmt -l .`
Expected: PASS. `TestAFailedAttributeAfterCreate…` prints one stderr line; that is the behaviour under test.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: treat `InvalidVpcID.NotFound` as a hard error on the first try (`TestASubnetWaitsFor…`); return the
`setMapPublicIP` error from `createSubnet` (`TestAFailedAttributeAfterCreate…`); set
`sub.MapPublicIpOnLaunch = aws.Bool(public)` whatever the call returned (same test — state claims `true` after AWS
refused); skip `setMapPublicIP` in `updateSubnet` (`TestMapPublicIP…`).

```bash
git add internal/awsprov/subnet.go internal/awsprov/provider.go internal/awsprov/subnet_test.go
git commit -m "aws.subnet: waits for a VPC AWS has not propagated, and never orphans a subnet it made" -- <same paths>
```

---
### Task 7: Discover and Import

**Files:**
- Create: `internal/awsprov/discover.go`
- Modify: `internal/awsprov/provider.go` (remove the `Discover`/`Import` stubs)
- Test: `internal/awsprov/discover_test.go`

**Interfaces:**
- Consumes: `vpcState`, `subnetState`, `readVPC`, `readSubnet`, `once`, `p.failed`.
- Produces: `Provider.Discover`, `Provider.Import`, `listVPCs(ctx, region)`, `listSubnets(ctx, region)`.

- [ ] **Step 1: Write the failing tests**

`internal/awsprov/discover_test.go`:

```go
package awsprov

import (
	"context"
	"strings"
	"testing"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/value"
)

// TestDiscoverFindsEveryRequestedTypeInEveryScannedRegionAcrossPages. The fixture is built so each
// plausible bug gives a different count: a page size of one (a paginator that stops early finds 1),
// a region not in discover_regions (a region leak finds 4), and a subnet (ignoring Types finds 4).
func TestDiscoverFindsEveryRequestedTypeInEveryScannedRegionAcrossPages(t *testing.T) {
	p, fake := fakeProvider(t, map[string]value.Value{"discover_regions": list("us-east-1", "eu-west-1")})
	fake.PageSize = 1
	a := fake.AddVPC("us-east-1", "10.0.0.0/16", map[string]string{"team": "a"})
	fake.AddVPC("us-east-1", "10.1.0.0/16", nil)
	fake.AddVPC("eu-west-1", "10.2.0.0/16", nil)
	fake.AddVPC("ap-south-1", "10.3.0.0/16", nil)
	if _, err := p.Create(context.Background(), desired(typeSubnet, subnetAttrs(a, nil))); err != nil {
		t.Fatal(err)
	}

	vpcs, err := p.Discover(context.Background(), provider.DiscoverRequest{Types: []string{typeVPC}})
	if err != nil {
		t.Fatal(err)
	}
	if len(vpcs) != 3 {
		t.Fatalf("discovered %d VPCs, want 3: %+v", len(vpcs), vpcs)
	}
	for _, d := range vpcs {
		if d.Type != typeVPC || !(strings.HasPrefix(d.ProviderID, "us-east-1/vpc-") || strings.HasPrefix(d.ProviderID, "eu-west-1/vpc-")) {
			t.Errorf("unexpected %s %s", d.Type, d.ProviderID)
		}
	}
	all, err := p.Discover(context.Background(), provider.DiscoverRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 4 {
		t.Errorf("discovering every type found %d, want 3 VPCs and 1 subnet", len(all))
	}
}

func TestDiscoverWithoutRegionsSaysWhatToAdd(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	fake.AddVPC("us-east-1", "10.0.0.0/16", nil)
	_, err := p.Discover(context.Background(), provider.DiscoverRequest{})
	if err == nil || !strings.Contains(err.Error(), "discover_regions") || !strings.Contains(err.Error(), `"test"`) {
		t.Fatalf("err = %v", err)
	}
}

// TestImportAdoptsByRegionalIDAndRefusesTheWrongType.
func TestImportAdoptsByRegionalIDAndRefusesTheWrongType(t *testing.T) {
	p, fake := fakeProvider(t, nil)
	ctx := context.Background()
	vpc := fake.AddVPC("eu-west-1", "10.0.0.0/16", nil)

	st, err := p.Import(ctx, typeVPC, "eu-west-1/"+vpc)
	if err != nil || st == nil || st.ProviderID != "eu-west-1/"+vpc {
		t.Fatalf("Import = %+v, %v", st, err)
	}
	if _, err := p.Import(ctx, typeSubnet, "eu-west-1/"+vpc); err == nil {
		t.Fatal("a VPC ID was imported as a subnet")
	}
	before := fake.Calls("DescribeVpcs")
	got, err := p.Import(ctx, typeVPC, "eu-west-1/vpc-0000000000000dead")
	if err != nil || got != nil {
		t.Fatalf("importing a missing VPC = %v, %v; want (nil, nil), which the host reports as no such resource", got, err)
	}
	if n := fake.Calls("DescribeVpcs") - before; n != 1 {
		t.Errorf("import made %d describes; it must not wait: its ID came from discovery", n)
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `go test -count=1 -run 'Discover|Import' ./internal/awsprov/`
Expected: FAIL — `not implemented`.

- [ ] **Step 3: Implement**

`internal/awsprov/discover.go`:

```go
package awsprov

import (
	"context"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
)

// Discover lists every requested type in every region the instance scans, including everything
// infrata did not create — which is what discovery is for. The regions come from `discover_regions`:
// infrata never sends one (DiscoverRequest.Region is always empty) and never sends `defaults:`.
func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	if len(p.config.DiscoverRegions) == 0 {
		return nil, fmt.Errorf("aws instance %q has no discover_regions, so there is nowhere to look: "+
			"add discover_regions: [us-east-1] to its providers: entry. `discover` takes no environment, so write "+
			"it as a literal, a variable with a default:, or pass it with --var", p.instance)
	}
	types := req.Types
	if len(types) == 0 {
		types = []string{typeSubnet, typeVPC}
	}
	var out []provider.DiscoveredResource
	for _, region := range p.config.DiscoverRegions {
		for _, typ := range types {
			var found []provider.DiscoveredResource
			var err error
			switch typ {
			case typeVPC:
				found, err = p.listVPCs(ctx, region)
			case typeSubnet:
				found, err = p.listSubnets(ctx, region)
			default:
				continue // not ours; the host skips undeclared types anyway
			}
			if err != nil {
				return nil, err
			}
			out = append(out, found...)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ProviderID < out[j].ProviderID })
	return out, nil
}

func discovered(st *resource.ResourceState) provider.DiscoveredResource {
	return provider.DiscoveredResource{Type: st.Type, ProviderID: st.ProviderID, Attributes: st.Attributes}
}

// listVPCs pages through DescribeVpcs. ctx is checked between pages: abandoning a read loses nothing.
// (The SDK also refuses to send with a cancelled context, so no test can tell the check apart from
// the SDK's; it is kept so the intent does not depend on middleware.)
func (p *Provider) listVPCs(ctx context.Context, region string) ([]provider.DiscoveredResource, error) {
	var out []provider.DiscoveredResource
	pg := ec2.NewDescribeVpcsPaginator(p.clients.ec2(region), &ec2.DescribeVpcsInput{})
	for pg.HasMorePages() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := pg.NextPage(ctx)
		if err != nil {
			return nil, p.failed("DescribeVpcs", region, "", err)
		}
		for _, v := range page.Vpcs {
			out = append(out, discovered(vpcState(region, v)))
		}
	}
	return out, nil
}

func (p *Provider) listSubnets(ctx context.Context, region string) ([]provider.DiscoveredResource, error) {
	var out []provider.DiscoveredResource
	pg := ec2.NewDescribeSubnetsPaginator(p.clients.ec2(region), &ec2.DescribeSubnetsInput{})
	for pg.HasMorePages() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := pg.NextPage(ctx)
		if err != nil {
			return nil, p.failed("DescribeSubnets", region, "", err)
		}
		for _, sub := range page.Subnets {
			out = append(out, discovered(subnetState(region, sub)))
		}
	}
	return out, nil
}

// Import reads one resource by its regional ID, once: the ID came from discovery, so it is visible.
// (nil, nil) for a well-formed ID that does not exist; the host words that for the user.
func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	switch resourceType {
	case typeVPC:
		return p.readVPC(ctx, id, once)
	case typeSubnet:
		return p.readSubnet(ctx, id, once)
	}
	return nil, unknownType(resourceType)
}
```

Delete the `Discover` and `Import` stubs from `provider.go`.

- [ ] **Step 4: Run**

Run: `go test -count=1 ./... && go vet ./...`
Expected: PASS.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: stop after the first page (1 VPC, not 3); ignore `req.Types` (4, not 3); scan a hard-coded
`us-east-1` only (2, not 3); give `Import` `p.patience` (4 describes, not 1); drop `parseID`'s prefix check
(the wrong-type import succeeds).

```bash
git add internal/awsprov/discover.go internal/awsprov/provider.go internal/awsprov/discover_test.go
git commit -m "aws: discover pages through every scanned region, and import refuses an ID of the wrong type" -- <same paths>
```

---

### Task 8: Through infrata's host — `pkg/plugintest`

**Files:**
- Modify: `internal/awsprov/protocol_test.go`

**Interfaces:**
- Consumes: `openHost` (Task 1), `isolateAWS`, `ec2fake`, `list`, `s`, `vpcAttrs`, `subnetAttrs`, `tagsValue`.
- Produces: nothing new; this task proves the host accepts what the plugin returns.

These tests assert what the HOST does (bookkeeping, sensitivity, undeclared attributes, classification across
the pipe), which Task 5's `TestAVPCIsCreatedWithItsTagsAndReadBack` deliberately asserts the plugin does NOT do.

- [ ] **Step 1: Write the tests**

Append to `internal/awsprov/protocol_test.go`:

```go
// hostProvider is a configured instance reached through the real protocol, against a fresh fake.
func hostProvider(t *testing.T) (provider.Provider, *ec2fake.Server) {
	t.Helper()
	fake := ec2fake.New()
	t.Cleanup(fake.Close)
	isolateAWS(t, fake.URL)
	prov, err := openHost(t).Configure(provider.Config{Instance: "main", Values: map[string]value.Value{
		"discover_regions": list("us-east-1"),
	}})
	if err != nil {
		t.Fatalf("Configure: %v", err)
	}
	return prov, fake
}

func TestACreateRoundTripsThroughTheHost(t *testing.T) {
	prov, _ := hostProvider(t)
	ctx := context.Background()
	st, err := prov.Create(ctx, desired(typeVPC, vpcAttrs(map[string]value.Value{"tags": tagsValue("team", "platform")})))
	if err != nil {
		t.Fatalf("Create through the host: %v", err)
	}
	if st.Address.Name != "r" || st.Provider != "main" {
		t.Errorf("host did not re-attach bookkeeping: address=%q provider=%q", st.Address, st.Provider)
	}
	got, err := prov.Read(ctx, st)
	if err != nil || got == nil {
		t.Fatalf("Read through the host = %v, %v", got, err)
	}
	if got.ProviderID != st.ProviderID {
		t.Errorf("provider ID changed across the pipe: %q → %q", st.ProviderID, got.ProviderID)
	}
}

// TestEveryDiscoveredAttributeIsDeclared. The host fails the whole discovery on one undeclared
// attribute, so a default VPC with tags and a subnet — the richest shapes — go through it.
func TestEveryDiscoveredAttributeIsDeclared(t *testing.T) {
	prov, fake := hostProvider(t)
	vpc := fake.AddVPC("us-east-1", "172.31.0.0/16", map[string]string{"Name": "default", "aws:reserved": "x"})
	if _, err := prov.Create(context.Background(), desired(typeSubnet, subnetAttrs(vpc, nil))); err != nil {
		t.Fatal(err)
	}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{})
	if err != nil {
		t.Fatalf("the host refused a discovery: %v", err)
	}
	if len(found) != 2 {
		t.Errorf("discovered %d through the host, want 2", len(found))
	}
}

func TestAThrottledCreateKeepsItsClassificationAcrossThePipe(t *testing.T) {
	prov, fake := hostProvider(t)
	fake.Inject(ec2fake.Fault{Action: "CreateVpc", Nth: 1, Status: 503, Code: "RequestLimitExceeded", Message: "slow down"})
	_, err := prov.Create(context.Background(), desired(typeVPC, vpcAttrs(nil)))
	if err == nil {
		t.Fatal("expected the throttle")
	}
	if got := prov.ClassifyError(err); got != provider.SafeToRetry {
		t.Errorf("classification after the pipe = %v, want SafeToRetry", got)
	}
	if !strings.Contains(err.Error(), "RequestLimitExceeded") {
		t.Errorf("the message lost AWS's code: %v", err)
	}
}

func TestImportingAMissingIDIsReportedByTheHost(t *testing.T) {
	prov, _ := hostProvider(t)
	_, err := prov.Import(context.Background(), typeVPC, "us-east-1/vpc-0000000000000dead")
	if err == nil || !strings.Contains(err.Error(), "no aws.vpc with id") {
		t.Fatalf("err = %v", err)
	}
}
```

(imports to add: `strings`, `github.com/infrata/infrata-provider-aws/internal/ec2fake`,
`github.com/infrata/infrata/pkg/provider`, `github.com/infrata/infrata/pkg/value`.)

Note: `hostProvider` does not zero the patience backoff — the plugin runs behind the pipe. No test here makes a
read miss, so no test waits.

- [ ] **Step 2: Run**

Run: `go test -count=1 -run 'ThroughTheHost|Discovered|AcrossThePipe|ReportedByTheHost' -v ./internal/awsprov/`
Expected: PASS. If `TestImportingAMissingID…` fails on wording, read infrata's `internal/pluginhost/adapter.go`
`Import` for the current text and assert on that; the test's point is that `(nil, nil)` reaches the user as a
sentence, not the exact words.

- [ ] **Step 3: Sabotage, then commit**

Sabotages: report an extra attribute `"state"` from `vpcState` (`TestEveryDiscoveredAttributeIsDeclared` — the host
refuses it); make `classify` return `NotSafeToRetry` for throttles (`…AcrossThePipe`); have `readVPC` return an
error instead of `(nil, nil)` when not found (`TestImportingAMissingID…`).

```bash
git add internal/awsprov/protocol_test.go
git commit -m "aws: what the plugin returns survives infrata's own host" -- internal/awsprov/protocol_test.go
```

---
### Task 9: A real infrata against the real binary and the fake endpoint (`-tags e2e`)

**Files:**
- Create: `e2e/e2e_test.go`, `e2e/testdata/basic/infra.yml`
- Test: the same file

**Interfaces:**
- Consumes: `ec2fake.New`, `AddVPC`, `SetTag`, `SetCIDR`, `Remove`, `VPCs`, `Subnets`; the binary from Task 1.
- Produces: `e2e/testdata/basic/infra.yml`, which the README quotes byte for byte (Task 12).

The harness is `infrata-provider-fake/e2e/e2e_test.go`'s: `TestMain` builds `infrata` from `$INFRATA_SRC`
(default `../../infrata`, skipping loudly if absent) and this plugin into a temp plugin dir; `infrata(...)` runs the
CLI with `--plugin-dir`; `expect(...)` checks exit code and substrings; `planOps(...)` reads `plan dev --output`.
Copy those four helpers unchanged except as below. The one structural change: the "cloud" is a fake server, so each
project owns one, and every command's environment points the plugin at it. The plugin inherits infrata's environment.

Do not copy `TestTheInfrataUnderTestSpeaksTheManifestsProtocol` here; it reads `plugin.yaml`, which Task 10 creates,
and Task 10 adds it.

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
    vpc_id: ${vpc.id}
    cidr: 10.0.1.0/24
    availability_zone: us-east-1a
```

The region reaches both resources through `defaults:`, from a variable with a `default:` so `discover` resolves it (D20); an environment file could override it.

- [ ] **Step 2: Write the suite**

`e2e/e2e_test.go` — the fake plugin's `TestMain` with the plugin build line changed to
`{"..", filepath.Join(pluginDir, "infrata-plugin-aws"), "./cmd/infrata-plugin-aws"}`, its `writeFile`
unchanged, and these replacing `project`, `infrata`, `expect` and `planOps` (now methods on a per-project `env`):

```go
//go:build e2e

package e2e

// (imports: encoding/json, fmt, os, os/exec, path/filepath, strings, testing,
//  github.com/infrata/infrata-provider-aws/internal/ec2fake)

// env is one project and the AWS account it talks to.
type env struct {
	dir  string
	fake *ec2fake.Server
}

func project(t *testing.T, body string) *env {
	t.Helper()
	if skipReason != "" {
		t.Skip(skipReason)
	}
	e := &env{dir: t.TempDir(), fake: ec2fake.New()}
	t.Cleanup(e.fake.Close)
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

// infrata runs the CLI in the project. Nothing from the developer's machine reaches the plugin: no
// plugins but ours, no AWS config or credentials but static test keys, no instance metadata, and
// EC2/STS pointed at this project's fake.
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
		"AWS_ENDPOINT_URL=", "AWS_ENDPOINT_URL_EC2="+e.fake.URL, "AWS_ENDPOINT_URL_STS="+e.fake.URL,
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

// TestTheWorkflow is AGENT.md's checklist against a VPC and a subnet.
func TestTheWorkflow(t *testing.T) {
	e := project(t, fixture(t, "basic"))
	var managedVPC string // set once apply has created it; every later subtest acts on this VPC
	vpcID := func() string {
		if managedVPC == "" {
			t.Fatal("no managed VPC yet: the apply subtest did not run or failed")
		}
		return managedVPC
	}

	t.Run("explain", func(t *testing.T) {
		e.expect(t, 0, []string{"aws.subnet", "(replaces on change)", "(default: false)", "aws.vpc", "<region>/<subnet id>"},
			"explain", "aws.subnet")
	})
	t.Run("plan proposes two creates", func(t *testing.T) {
		e.expect(t, 2, []string{"Plan: 2 to create", `region: "us-east-1"`}, "plan", "dev")
	})
	t.Run("apply creates them, and a re-plan is clean", func(t *testing.T) {
		e.expect(t, 2, []string{"Apply complete: 2 applied, 0 failed"}, "apply", "dev", "--auto-approve")
		if len(e.fake.VPCs()) != 1 || len(e.fake.Subnets()) != 1 || e.fake.Subnets()[0].VPCID != e.fake.VPCs()[0].ID {
			t.Fatalf("fake holds %+v / %+v", e.fake.VPCs(), e.fake.Subnets())
		}
		managedVPC = e.fake.VPCs()[0].ID
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("re-plan after apply proposes %v, want nothing — an attribute read back differs from what was created", ops)
		}
	})
	t.Run("a tag changed outside infrata is an update", func(t *testing.T) {
		e.fake.SetTag(vpcID(), "team", "someone-else")
		if kind := e.planOps(t)["vpc"]; kind != "update" {
			t.Fatalf("vpc plans as %q, want update", kind)
		}
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
		if got := e.fake.VPCs()[0].Tags["team"]; got != "platform" { // one VPC exists at this point
			t.Fatalf("tag after apply = %q", got)
		}
	})
	t.Run("an attribute changed in place converges", func(t *testing.T) {
		e.editConfig(t, "    availability_zone: us-east-1a\n", "    availability_zone: us-east-1a\n    map_public_ip_on_launch: true\n")
		if kind := e.planOps(t)["private_a"]; kind != "update" {
			t.Fatalf("private_a plans as %q, want update", kind)
		}
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
		if !e.fake.Subnets()[0].MapPublicIP {
			t.Fatal("map_public_ip_on_launch not applied")
		}
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after the update proposes %v", ops)
		}
	})
	t.Run("removing tags converges", func(t *testing.T) {
		e.editConfig(t, "    tags:\n      team: platform\n", "")
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after removing tags proposes %v — tags were merged, not replaced, or {} was reported", ops)
		}
	})
	t.Run("a CIDR changed outside infrata forces a replacement", func(t *testing.T) {
		e.fake.SetCIDR(vpcID(), "10.50.0.0/16")
		if kind := e.planOps(t)["vpc"]; kind != "replace" {
			t.Fatalf("vpc plans as %q, want replace", kind)
		}
	})
	t.Run("a subnet deleted outside infrata is recreated", func(t *testing.T) {
		// Restore the CIDR first so this subtest is about one thing.
		e.fake.SetCIDR(vpcID(), "10.0.0.0/16")
		e.fake.Remove(e.fake.Subnets()[0].ID)
		if kind := e.planOps(t)["private_a"]; kind != "create" {
			t.Fatalf("private_a plans as %q, want create", kind)
		}
		e.expect(t, 2, []string{"0 failed"}, "apply", "dev", "--auto-approve")
	})
	t.Run("discover and import adopt a VPC infrata did not create", func(t *testing.T) {
		other := e.fake.AddVPC("us-east-1", "172.16.0.0/16", map[string]string{"owner": "legacy"})
		e.expect(t, 0, []string{"aws.vpc", "us-east-1/" + other}, "discover")
		e.expect(t, 0, []string{"1 resource imported"}, "import", "dev", "aws.vpc.us-east-1/"+other, "--generate")
		if ops := e.planOps(t); len(ops) != 0 {
			t.Fatalf("plan after import --generate proposes %v", ops)
		}
	})
	t.Run("destroy removes everything infrata manages", func(t *testing.T) {
		e.expect(t, 2, []string{"0 failed"}, "destroy", "dev", "--auto-approve")
		if len(e.fake.VPCs()) != 0 || len(e.fake.Subnets()) != 0 {
			t.Fatalf("after destroy the fake holds %+v / %+v", e.fake.VPCs(), e.fake.Subnets())
		}
	})
}

// TestASubnetWithoutAVPCFailsBeforeAnyAPICall. The requirement is a pre-flight hint; this is what it
// buys a user.
func TestASubnetWithoutAVPCFailsBeforeAnyAPICall(t *testing.T) {
	e := project(t, `project: demo
environments:
  dev: {}
providers:
  - plugin: aws
    defaults:
      region: us-east-1
resources:
  orphan:
    type: aws.subnet
    vpc_id: vpc-0123456789abcdef0
    cidr: 10.0.1.0/24
    availability_zone: us-east-1a
`)
	e.expect(t, 1, []string{`"orphan" is missing required vpc`, "aws.vpc"}, "plan", "dev")
	if n := e.fake.Calls("CreateSubnet") + e.fake.Calls("DescribeSubnets"); n != 0 {
		t.Errorf("%d EC2 calls were made for a project that could not plan", n)
	}
}

// TestAMisspelledKeyIsRefusedAgainstTheProvidersEntry.
func TestAMisspelledKeyIsRefusedAgainstTheProvidersEntry(t *testing.T) {
	e := project(t, strings.Replace(fixture(t, "basic"), "discover_regions:", "discover_region:", 1))
	e.expect(t, 1, []string{`provider instance "aws" could not be configured`, `unknown configuration "discover_region"`}, "plan", "dev")
}

// TestProviderVariablesReachEnvironmentCommands pins infrata 5895f8a/76c3f28 for this plugin's model
// (D20). The region is set ONLY in environments/dev.yml: a `default:` would resolve without an
// environment and let a regression pass. It also pins F9: discover refuses the environment-only
// default even though it never uses it, and --var gets past that.
func TestProviderVariablesReachEnvironmentCommands(t *testing.T) {
	e := project(t, `project: demo
environments:
  dev: {}
variables:
  aws_region:
    type: string
providers:
  - plugin: aws
    discover_regions: [us-east-1]
    defaults:
      region: ${aws_region}
resources:
  vpc:
    type: aws.vpc
    cidr: 10.0.0.0/16
`)
	writeFile(t, filepath.Join(e.dir, "environments", "dev.yml"), "variables:\n  aws_region: us-east-1\n")
	unresolved := func(t *testing.T, out string) {
		for _, bad := range []string{"undefined variable", "could not be resolved"} {
			if strings.Contains(out, bad) {
				t.Fatalf("a providers: variable did not resolve (%q):\n%s", bad, out)
			}
		}
	}
	unresolved(t, e.expect(t, 2, []string{"Apply complete: 1 applied"}, "apply", "dev", "--auto-approve"))
	out, _ := e.infrata(t, "refresh", "dev")
	unresolved(t, out)
	other := e.fake.AddVPC("us-east-1", "172.16.0.0/16", nil)
	unresolved(t, e.expect(t, 0, []string{"1 resource imported"}, "import", "dev", "aws.vpc.us-east-1/"+other))
	e.expect(t, 1, []string{"could not be resolved", "--var"}, "discover") // F9; exit code as infrata reports a refused instance
	e.expect(t, 0, []string{"us-east-1/" + other}, "discover", "--var", "aws_region=us-east-1")
	unresolved(t, e.expect(t, 2, []string{"0 failed"}, "destroy", "dev", "--auto-approve"))
	if len(e.fake.VPCs()) != 0 {
		t.Fatalf("destroy left %+v", e.fake.VPCs())
	}
}
```

- [ ] **Step 3: Run**

Run: `go test -tags e2e -count=1 -v ./e2e/`
Expected: PASS (the recreate subtest takes about 4s: the patience before reporting gone). If a quoted output
line differs (for example `1 resource imported`, or the exit code of a refused `discover`), read the actual output
and assert on it — the wording is infrata's — and note it here. If infrata has since changed F9 so `discover`
succeeds without `--var`, flip that assertion and mark F9 answered.

- [ ] **Step 4: Sabotage, then commit**

Sabotages: report `tags: {}` for an untagged resource (the "removing tags converges" subtest); make
`availability_zone` optional in the schema and omit it from the fixture's subnet — the re-plan after apply is not
clean; drop the `aws:` tag filter and add `"aws:cloudformation:stack-name": "legacy"` to the `AddVPC` tags in the
import subtest (the post-import plan is not clean); have `vpcState` report `is_default` as a string (the host refuses
the attribute's kind, and apply fails).

```bash
git add e2e/e2e_test.go e2e/testdata/basic/infra.yml
git commit -m "e2e: infrata plans, applies, repairs drift, imports and destroys against the fake EC2" -- <same paths>
```

---

### Task 10: Manifest, release gate, CI

**Files:**
- Create: `plugin.yaml`, `scripts/release-check`, `scripts/build-release`, `scripts/scripts_test.go`
- Create: `internal/awsprov/manifest_test.go`
- Create: `.github/workflows/release.yml`, `.github/workflows/ci.yml`, `.github/workflows/bump-infrata.yml`

**Interfaces:**
- Consumes: `awsprov.Version`, `PluginName`; infrata `pkg/pluginmanifest.Parse`, `pkg/pluginproto.Version`.
- Produces: the release convention `infrata-plugin-aws_<version>_<goos>_<goarch>.tar.gz` (`.zip` on Windows).

Copy the fake plugin's files and change only what names the plugin. Each copy is listed with its exact
substitutions so nothing else drifts.

- [ ] **Step 1: `plugin.yaml`**

```yaml
# plugin.yaml: what this plugin is, and what it works with. infrata PLAN.md §31.2.
# Read at a release TAG, never at the default branch, which describes unreleased code.
manifest: 1
name: aws
version: 0.1.0
protocol: [1]
platforms: [linux/amd64, linux/arm64, linux/arm, linux/386, darwin/amd64, darwin/arm64, windows/amd64, windows/arm64]
description: The AWS provider for infrata.
source: https://github.com/infrata/infrata-provider-aws
```

No `infrata:` key: absent means unconstrained, and no infrata release is known to be incompatible.

- [ ] **Step 2: The manifest test (failing first — run it before creating `plugin.yaml` to see it fail)**

`internal/awsprov/manifest_test.go` — `infrata-provider-fake/internal/fake/manifest_test.go` with `package awsprov`
and nothing else changed (it already reads `../../plugin.yaml` and compares against `PluginName` and
`pluginproto.Version`).

Run: `go test -count=1 -run Manifest ./internal/awsprov/` — FAIL before Step 1's file exists, PASS after.

Also copy `TestTheInfrataUnderTestSpeaksTheManifestsProtocol` from `infrata-provider-fake/e2e/e2e_test.go` into
`e2e/e2e_test.go` unchanged (it reads `../plugin.yaml`), and add `e2e/e2e_test.go` to this task's commit. Run:
`go test -tags e2e -count=1 -run Manifest ./e2e/`.

- [ ] **Step 3: Scripts**

Copy `infrata-provider-fake/scripts/{release-check,build-release,scripts_test.go}` and substitute:

| In the fake's file | Here |
| --- | --- |
| `infrata-plugin-fake` | `infrata-plugin-aws` |
| `./cmd/infrata-plugin-fake` | `./cmd/infrata-plugin-aws` |
| `github.com/infrata/infrata-provider-fake/internal/fake.Version` | `github.com/infrata/infrata-provider-aws/internal/awsprov.Version` |
| `FAKE_VERSION_SYMBOL` | `PLUGIN_VERSION_SYMBOL` |
| `FAKE_MANIFEST` | `PLUGIN_MANIFEST` |
| `internal/fake.Version` (in the error message) | `internal/awsprov.Version` |

`chmod +x scripts/release-check scripts/build-release`.

Run: `go test -count=1 ./scripts/`
Expected: PASS, including the test that points `-X` at a nonexistent symbol and requires the gate to refuse with
`0.0.0-dev`. Cross-compiling eight platforms with the AWS SDK is slower than the fake's; if
`TestBuildReleaseNamesArchivesByTheInstallConvention` times out, it already uses `PLATFORMS` to build two.

- [ ] **Step 4: Workflows (D19)**

Every workflow fetches infrata as a module, so each starts with the same two steps: tell Go the module is private,
and give git the token. A pinned build must never see a `go.work` (none is committed; `GOWORK: off` makes that
explicit).

`.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
  workflow_dispatch:

permissions:
  contents: read

env:
  GOTOOLCHAIN: local
  GOPRIVATE: github.com/infrata/*

jobs:
  # BLOCKING. The infrata version go.mod requires — what a release ships against.
  pinned:
    runs-on: ubuntu-latest
    env:
      GOWORK: "off"
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - name: Let go fetch the private infrata module
        run: git config --global url."https://x-access-token:${{ secrets.INFRATA_CHECKOUT_TOKEN }}@github.com/infrata/".insteadOf "https://github.com/infrata/"
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: The infrata revision go.mod requires
        id: infrata
        working-directory: infrata-provider-aws
        run: |
          v="$(go list -m -f '{{.Version}}' github.com/infrata/infrata)"
          # A pseudo-version ends in a 12-character commit hash; a tag is used as is.
          if [[ "$v" =~ -([0-9a-f]{12})$ ]]; then v="${BASH_REMATCH[1]}"; fi
          echo "ref=$v" >> "$GITHUB_OUTPUT"
      # The e2e suite builds the infrata CLI from ../infrata: the same revision the plugin compiles against.
      - uses: actions/checkout@v7
        with:
          repository: infrata/infrata
          path: infrata
          fetch-depth: 0
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - run: git -C infrata checkout --detach "${{ steps.infrata.outputs.ref }}"
      - name: Format, vet, test
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          test -z "$(gofmt -l .)"
          go vet ./...
          go vet -tags e2e,live ./...
          go test -count=1 ./...
          go test -tags e2e -count=1 ./e2e/

  # NON-BLOCKING early warning: infrata's main, through a workspace. Red here means the next bump needs work.
  infrata-main:
    runs-on: ubuntu-latest
    continue-on-error: true
    steps:
      - uses: actions/checkout@v7
        with:
          path: infrata-provider-aws
      - name: Let go fetch the private infrata module
        run: git config --global url."https://x-access-token:${{ secrets.INFRATA_CHECKOUT_TOKEN }}@github.com/infrata/".insteadOf "https://github.com/infrata/"
      - uses: actions/checkout@v7
        with:
          repository: infrata/infrata
          path: infrata
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - uses: actions/setup-go@v7
        with:
          go-version: "1.27"
          cache-dependency-path: infrata-provider-aws/go.sum
      - name: Test against infrata main
        working-directory: infrata-provider-aws
        run: |
          set -euo pipefail
          go work init . ../infrata
          go test -count=1 ./...
          go test -tags e2e -count=1 ./e2e/
```

`.github/workflows/release.yml` — the fake plugin's file changed in four ways: `infrata-provider-fake` →
`infrata-provider-aws` in the checkout `path`, every `working-directory` and `cache-dependency-path`; top-level
`env` gains `GOPRIVATE: github.com/infrata/*` and `GOWORK: "off"`; the git `insteadOf` step is added before
`setup-go`; and the infrata checkout is replaced by `ci.yml`'s "revision go.mod requires" step, the `fetch-depth: 0`
checkout and the `git checkout --detach` (so a release is tested against the infrata it ships with, not `main`).

`.github/workflows/bump-infrata.yml`:

```yaml
name: Bump infrata

# infrata tags often until its first official release (James, 2026-09-13). This notices each tag and
# proposes it, having already run the suite: a PR opened with GITHUB_TOKEN gets CI runs that wait for
# approval, so the evidence is attached here instead.

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
        run: git config --global url."https://x-access-token:${{ secrets.INFRATA_CHECKOUT_TOKEN }}@github.com/infrata/".insteadOf "https://github.com/infrata/"
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
          # @upgrade never moves from a newer pseudo-version back to an older tag.
          go get github.com/infrata/infrata@upgrade
          go mod tidy
          after="$(go list -m -f '{{.Version}}' github.com/infrata/infrata)"
          echo "before=$before" >> "$GITHUB_OUTPUT"
          echo "after=$after" >> "$GITHUB_OUTPUT"
          if [[ "$before" == "$after" ]]; then echo "changed=false" >> "$GITHUB_OUTPUT"; else echo "changed=true" >> "$GITHUB_OUTPUT"; fi
      - uses: actions/checkout@v7
        if: steps.bump.outputs.changed == 'true'
        with:
          repository: infrata/infrata
          path: infrata
          ref: ${{ steps.bump.outputs.after }}
          token: ${{ secrets.INFRATA_CHECKOUT_TOKEN }}
      - name: Test against the new version
        if: steps.bump.outputs.changed == 'true'
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
          git commit -m "deps: infrata ${{ steps.bump.outputs.before }} -> ${{ steps.bump.outputs.after }}" -- go.mod go.sum
          git push origin "$branch"
          gh pr create --title "infrata ${{ steps.bump.outputs.after }}" \
            --body "Upgraded from ${{ steps.bump.outputs.before }}. vet, the plain suite and the e2e suite passed against it in this run."
```

The `@upgrade` step only runs a newer TAG through; the checkout `ref` is that tag. (A bump to a pseudo-version is
always a person's decision, made in a normal PR.) The repository setting "Allow GitHub Actions to create and approve
pull requests" must be on, which is James's to change.

Check locally what can be checked: `go vet -tags e2e,live ./...` (after Tasks 9 and 11 exist),
`scripts/release-check v0.1.0` (expects "tag, plugin.yaml and binary all say 0.1.0"), and each workflow's shell
logic by hand — `go list -m -f '{{.Version}}' github.com/infrata/infrata` and the pseudo-version regex on its
output. The workflows themselves are proven only by the first push, which is James's call.

- [ ] **Step 5: Sabotage, then commit**

Sabotages: `name: aws2` in `plugin.yaml` (manifest test); the ldflags symbol pointing at `internal/fake.Version`
(the scripts test for the stamp); `version: 0.1.1` with tag `v0.1.0` (release-check refuses).

```bash
git add plugin.yaml scripts/release-check scripts/build-release scripts/scripts_test.go \
  internal/awsprov/manifest_test.go e2e/e2e_test.go .github/workflows/release.yml .github/workflows/ci.yml \
  .github/workflows/bump-infrata.yml
git commit -m "release: the same three-way version gate as the fake plugin, and CI on every push" -- <same paths>
```

---

### Task 11: The live suite — real AWS, opt-in

**Files:**
- Create: `live/live_test.go`, `live/README.md`

**Interfaces:**
- Consumes: the plugin through `pkg/plugintest` (so the host's rules apply, as in production).
- Produces: nothing other suites use.

Run by hand only. Never in CI (Q2). Everything it creates is free (VPCs, subnets).

- [ ] **Step 1: Write the suite**

`live/live_test.go`:

```go
//go:build live

// Package live runs the plugin against a real AWS account. It is the only suite that can see real
// eventual consistency and IAM behaviour, and the only one that costs anything if it leaks, so it
// refuses to run unless it is told which account it may use and the credentials really are that account.
//
//	INFRATA_AWS_LIVE_PROFILE=sandbox INFRATA_AWS_LIVE_ACCOUNT=111111111111 go test -tags live -count=1 -v ./live/
package live

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/infrata/infrata-provider-aws/internal/awsprov"
	"github.com/infrata/infrata/pkg/address"
	"github.com/infrata/infrata/pkg/plugintest"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

const runTag = "infrata-live-run"

func env(t *testing.T) (profile, account, region string) {
	t.Helper()
	profile, account = os.Getenv("INFRATA_AWS_LIVE_PROFILE"), os.Getenv("INFRATA_AWS_LIVE_ACCOUNT")
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
	return profile, account, region
}

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }

func TestTheLifecycleAgainstRealAWS(t *testing.T) {
	profile, _, region := env(t)
	ctx := context.Background()
	run := fmt.Sprintf("%d", time.Now().UnixNano())
	tags := value.Map(map[string]value.Value{runTag: s(run)}, value.SourceExplicit)

	host, err := plugintest.Open(ctx, awsprov.NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer host.Close()
	prov, err := host.Configure(provider.Config{Instance: "live", Values: map[string]value.Value{
		"profile":          s(profile),
		"discover_regions": value.List([]value.Value{s(region)}, value.SourceExplicit),
	}})
	if err != nil {
		t.Fatal(err)
	}

	vpc, err := prov.Create(ctx, &resource.DesiredResource{Address: address.Address{Name: "vpc"}, Type: "aws.vpc",
		Attrs: map[string]value.Value{"region": s(region), "cidr": s("10.99.0.0/16"), "tags": tags}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = prov.Delete(context.Background(), vpc) })

	// Immediately: the eventual-consistency cases the fake can only simulate.
	if got, err := prov.Read(ctx, vpc); err != nil || got == nil {
		t.Fatalf("read straight after create = %v, %v", got, err)
	}
	vpcID, _ := vpc.Attributes["id"].AsString()
	subnet, err := prov.Create(ctx, &resource.DesiredResource{Address: address.Address{Name: "subnet"}, Type: "aws.subnet",
		Attrs: map[string]value.Value{"region": s(region), "vpc_id": s(vpcID), "cidr": s("10.99.1.0/24"),
			"availability_zone": s(region + "a"), "map_public_ip_on_launch": value.Bool(true, value.SourceExplicit), "tags": tags}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = prov.Delete(context.Background(), subnet) })
	if b, _ := subnet.Attributes["map_public_ip_on_launch"].AsBool(); !b {
		t.Error("map_public_ip_on_launch was not applied straight after create")
	}

	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: []string{"aws.subnet"}})
	if err != nil {
		t.Fatal(err)
	}
	var seen bool
	for _, d := range found {
		seen = seen || d.ProviderID == subnet.ProviderID
	}
	if !seen {
		t.Error("discovery did not find the subnet it just made (allow for propagation before treating as a bug)")
	}
	if _, err := prov.Import(ctx, "aws.vpc", subnet.ProviderID); err == nil {
		t.Error("a subnet ID imported as a VPC")
	}

	if err := prov.Delete(ctx, subnet); err != nil {
		t.Fatal(err)
	}
	if err := prov.Delete(ctx, vpc); err != nil {
		t.Fatal(err) // a DependencyViolation here means the subnet's deletion had not propagated
	}
	if got, err := prov.Read(ctx, vpc); err != nil || got != nil {
		t.Errorf("read after delete = %v, %v; want gone", got, err)
	}
}

// TestSweepLeftovers deletes anything this suite tagged more than an hour ago: a crashed run's debris.
func TestSweepLeftovers(t *testing.T) {
	profile, _, region := env(t)
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		t.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)
	cutoff := time.Now().Add(-time.Hour).UnixNano()
	filter := []ec2types.Filter{{Name: aws.String("tag-key"), Values: []string{runTag}}}
	old := func(tags []ec2types.Tag) bool {
		for _, tg := range tags {
			if aws.ToString(tg.Key) == runTag {
				var n int64
				_, _ = fmt.Sscan(aws.ToString(tg.Value), &n)
				return n < cutoff
			}
		}
		return false
	}
	subs, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{Filters: filter})
	if err != nil {
		t.Fatal(err)
	}
	for _, sub := range subs.Subnets {
		if old(sub.Tags) {
			t.Logf("deleting %s", aws.ToString(sub.SubnetId))
			if _, err := client.DeleteSubnet(ctx, &ec2.DeleteSubnetInput{SubnetId: sub.SubnetId}); err != nil && !strings.Contains(err.Error(), "NotFound") {
				t.Error(err)
			}
		}
	}
	vpcs, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{Filters: filter})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range vpcs.Vpcs {
		if old(v.Tags) {
			t.Logf("deleting %s", aws.ToString(v.VpcId))
			if _, err := client.DeleteVpc(ctx, &ec2.DeleteVpcInput{VpcId: v.VpcId}); err != nil && !strings.Contains(err.Error(), "NotFound") {
				t.Error(err)
			}
		}
	}
}
```

`live/README.md`: what the suite needs (a dedicated account, a profile for it, the two variables), what it creates
(one VPC `10.99.0.0/16` and one subnet, tagged `infrata-live-run`), that it refuses to run on the wrong account,
how to run the sweeper, and that it is never run in CI.

- [ ] **Step 2: Check it compiles and skips**

Run: `go vet -tags live ./live/ && go test -tags live -count=1 -v ./live/`
Expected: both tests SKIP with the "set INFRATA_AWS_LIVE_PROFILE…" message. Running it for real waits on Q2.

- [ ] **Step 3: Record what a real run found**

When it is first run against a real account, record in the Verification log: whether `InvalidSubnetID.NotFound`
is the real code, whether a read straight after create needed the patience, and whether any other claim here
came back different. Until then those rows stay marked unconfirmed.

- [ ] **Step 4: Commit**

The sabotage for a suite nobody runs yet is the account guard: set `INFRATA_AWS_LIVE_ACCOUNT` to a wrong account
with a real profile and confirm it refuses before creating anything. Record whether that was done.

```bash
git add live/live_test.go live/README.md
git commit -m "live: an opt-in suite against real AWS that refuses any account it was not told to use" -- <same paths>
```

---

### Task 12: README, and closing out

**Files:**
- Create: `README.md`, `internal/awsprov/readme_test.go`
- Modify: `CLAUDE.md` ("Current state"), this plan (Verification log rows confirmed during execution)
- Vault: `projects/labs/infra-tool.md`, `projects/labs/daily/<date>.md`

- [ ] **Step 1: README, tied to the tested fixture**

`README.md` must cover, with real commands and output from the e2e run:
- what the plugin is, in two sentences
- how to build it and where infrata finds it (`--plugin-dir`, `.infra/plugins/`, `~/.local/share/infrata/plugins/`, `$PATH`)
- credentials: the default chain, `profile`, `assume_role_arn`; one instance per account; `import --provider <instance>`
  to adopt from one account when two are configured
- regions: `defaults: {region: ${aws_region}}` with a `default:` and per-environment override, per-resource override,
  the warning that changing the default region replaces every resource that inherits it, and `discover_regions`
  (literal — `discover` has no environment, and refuses an environment-only value even in `defaults:`)
- building: `go work init . ../infrata` for local work, `GOWORK=off` plus credentials for the pinned build
- the `e2e/testdata/basic/infra.yml` project quoted byte for byte
- every type, its attributes, and which are computed, force-new or defaulted; the `<region>/<id>` import form
- what the plugin does about eventual consistency and retries, in user terms (a read may take up to ~4s to report
  a resource gone; a failed create is not retried if AWS might have acted)
- how to run the three suites, and a pointer to the fake plugin's `AGENT.md` for plugin authors

`internal/awsprov/readme_test.go` — the fake's `readme_test.go` with the fixture path unchanged and the
must-mention list replaced by:
`"discover_regions", "assume_role_arn", "profile", "--provider", "defaults:", "${aws_region}", "us-east-1/vpc-", "-tags e2e", "-tags live", "go work init", "GOWORK=off", "plugin.yaml", "scripts/release-check", "0.0.0-dev"`.

Run: `go test -count=1 -run Readme ./internal/awsprov/` — FAIL before the README exists, PASS after. Sabotage: change
one character in the README's quoted fixture.

- [ ] **Step 2: Whole-suite verification**

```bash
git -C ../infrata status --short && git -C ../infrata log -1 --format=%h
gofmt -l . ; go vet ./... ; go vet -tags e2e,live ./...
go test -count=1 ./...
go test -tags e2e -count=1 -v ./e2e/
go test -tags live -count=1 ./live/      # skips without the live variables
```

All green, output kept for the close-out note, and the infrata commit it ran against recorded.

- [ ] **Step 3: Documentation**

- `CLAUDE.md`: "Current state" becomes "Built", naming the two types, the suites and their status, and what is not
  built yet (VPC DNS attributes, other types, the live suite's first run).
- This plan: every Verification log row still marked unconfirmed either confirmed or left marked, with a reason.
- Vault `projects/labs/infra-tool.md`: the AWS provider section's status, decisions that changed during execution,
  new follow-ups; today's labs daily note gets one bullet linking it.

- [ ] **Step 4: Commit, then ask**

```bash
git add README.md internal/awsprov/readme_test.go CLAUDE.md docs/plans/2026-09-13-first-slice-vpc-subnet.md
git commit -m "docs: a README whose example is the one the e2e suite runs" -- <same paths>
```

Ask James before pushing, and before any tag (a tag starts the release workflow).
