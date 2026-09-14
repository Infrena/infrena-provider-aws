# Generic Cloud Control provider — design

**Status:** draft for James's review, 2026-09-14. **Supersedes** Tasks 5–6 of
`docs/plans/2026-09-13-first-slice-vpc-subnet.md` (handwritten `aws.vpc`/`aws.subnet`). **Evidence:**
`docs/investigations/2026-09-13-generic-aws-provider.md` and `spikes/generic-aws/`. **Contract:** infrata **v0.3.0**
(`45deb30`): plugin protocol 2, `schema.Attribute.Optional` and `Aliases` (PLAN §14.1), `lifecycle: ignore_changes`
(§14.2, host-only).

## 1. Decisions already made (James, 2026-09-13 and 2026-09-14)

| # | Decision |
| --- | --- |
| J1 | **Cloud Control API is the core.** One generic provider manages every supported `AWS::` type. No handwritten resources. |
| J2 | **Type names** are `aws.<name>`, where `<name>` is the resource segment lowercased (`AWS::EC2::VPC` → `aws.vpc`) when that segment is unique across supported types, and `aws.<service>.<name>` otherwise (`AWS::EC2::Instance` → `aws.ec2.instance`). |
| J3 | **Names never move.** A committed name map locks each assigned name. A later AWS type whose segment clashes with a locked short name gets the qualified form; the locked name keeps its owner. |
| J4 | **Attributes accept several spellings:** AWS's property name case-insensitively, a generated snake_case form, and curated friendly aliases (`cidr`). |
| J5 | **Plans, `explain` and `import --generate` show the friendly alias** where one exists. |
| J6 | **Engine support comes from infrata**, not the plugin: done in v0.3.0. |

## 2. Facts this design rests on (measured 2026-09-14, us-east-1)

- The public schema bundle (`https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip`,
  2,998,086 bytes) holds 1,729 `AWS::` schemas. **Exactly 1,584 have `create`, `read` and `delete` handlers, the same
  set `list-types` reports as provisionable** (1,426 fully mutable and 159 immutable, 1,584 distinct; one type appears in both lists). The
  bundle alone is enough; no API call at generation time.
- Top-level properties: 14,971 (9.5 per type; largest `AWS::RDS::DBInstance`, 101).
- **Name folding is clean:** no type has two properties equal under lowercasing, or a generated snake_case form
  colliding with another property.
- Short names: 1,073 unique; 152 clashing segments cover 511 types; `service.resource` never clashes.
- Handler shape: 297 types lack `update` (every change replaces); 217 lack `list`; 381 declare a list `handlerSchema`
  (listing needs a parent `ResourceModel`); 306 have composite primary identifiers.
- Flags on **nested** pointers: 955 (infrata flags are top-level only).
- `writeOnlyProperties`: 1,170, of which 86 look secret by name (password, secret, token, private key, credential).
- Properties whose name folds to an infrata resource key: `Type` (95 types), `Provider` (8), `Region` (5, e.g.
  `AWS::Route53::RecordSet`), `Lifecycle` (1). infrata matches resource keys **exactly** (`type`, `depends_on`,
  `provider`, `skip`, `only`, `lifecycle` in `internal/config/decode.go`); reserved attributes are exactly
  `prevent_destroy` and `retain`, which no AWS property uses.
- Generated definitions as JSON: ~1.74 MB with one-line descriptions (~0.30 MB gzipped), ~0.90 MB without. infrata
  loads every plugin's schemas on **every** command, `validate` included.

## 3. Architecture

```text
schema bundle (pinned, fetched)          overlay.yaml (curated, committed)
            │                                        │
            └──────────► cmd/gen-cloudcontrol ◄──────┘
                                  │
          names.lock.json (committed)    internal/catalog/catalog.json.gz (committed, embedded)
                                  │
internal/catalog  ── infrata definitions + per-type runtime metadata (CFN type, identifier, write-only, tagging)
internal/cloudcontrol ── the generic Provider: one implementation, dispatch by catalog entry
internal/awsprov  ── Plugin, instance config, credentials, clients, error classification (Tasks 1–4, adapted)
internal/ccfake   ── in-process Cloud Control fake (awsJson1_0), replaces internal/ec2fake
```

### 3.1 Generator (`cmd/gen-cloudcontrol`)

Runs by hand or in the bump workflow, never at build or run time. Inputs: the bundle zip (downloaded by
`scripts/fetch-schemas`, its SHA-256 recorded in the output header), `overlay.yaml`, and the current
`names.lock.json`. Outputs are committed and reviewed as diffs.

**Types.** Include a schema only if it has `create`, `read` and `delete` handlers. Assign names:

1. Every type already in `names.lock.json` keeps its name, including types AWS has since removed (kept as tombstones so
   a name is never reused).
2. For new types: `aws.<segment>` if the lowercased segment is unused in the lock AND no other current type shares it;
   otherwise `aws.<service>.<segment>`.
3. Refuse (fail generation) if two types would get one name.

**Attributes**, per top-level property:

| Schema says | infrata attribute |
| --- | --- |
| `readOnlyProperties` | `Computed` |
| in `required` | `Required` |
| anything else | `Optional` + `Computed` (configuration may set it; AWS's value is kept when unset) |
| `createOnlyProperties` (top level), or the type has no `update` handler | `ForceNew` |
| overlay `sensitive:` | `Sensitive` |
| JSON type `string`/`integer`/`number`/`boolean`/`array`/`object` (following `$ref`) | `KindString`/`Int`/`Float`/`Bool`/`List`/`Map`; unions and untyped → `KindString` + reported |

Name and spellings: **canonical = AWS's property name** (`CidrBlock`); `Aliases` = curated aliases from the overlay
first, then the snake_case form (`cidr_block`) if it differs from the lowercased name. infrata's `Display` shows the
first alias.

Collisions the generator must resolve rather than emit (infrata's `Validate` would refuse the plugin):

- A property folding to `region` on a **regional** type (5 types have one; 4 if `AWS::Route53::RecordSet` is global per the overlay) clashes with the plugin's own `region`. The
  plugin's attribute stays `region` everywhere else; on these types it is `aws_region` and the property keeps its name.
- A property folding to `type`, `provider` or `lifecycle`: the canonical name is legal, but its lowercase spelling is
  claimed by infrata's resource keys. The generator gives it a first alias `<name>_value` (e.g. `type_value`) so
  `Display` never renders something that decodes as a resource key.

**Per-type runtime metadata** (not sent to infrata): CFN type name, primary identifier pointers, write-only
properties, whether `update`/`list` exist, list `handlerSchema` inputs, the tagging property, and regional vs global
(overlay).

**Output size:** one-line descriptions only; gzip-embedded. Task 1 measures `infrata validate` with the full catalog
against a no-op plugin before anything else is built (§6).

### 3.2 Overlay (`overlay.yaml`)

Hand-maintained, reviewed like code. Only what the schemas cannot say:

- `global:` type-name patterns that are not regional (`AWS::IAM::*`, `AWS::Organizations::*`, `AWS::CloudFront::*`,
  and individually verified others). Global types get no `region` attribute and are called in `us-east-1`.
- `aliases:` per type, per property (`AWS::EC2::VPC: {CidrBlock: [cidr]}`), seeded for the core types first (§7 Q4).
- `sensitive:` per type, per property; the generator prints the 86 name-based candidates for review instead of
  guessing.
- `requirements:` optional pre-flight hints (`AWS::EC2::Subnet` needs `AWS::EC2::VPC`).

### 3.3 Generic provider (`internal/cloudcontrol`)

Every method looks the type up in the catalog, then:

- **Create:** strip `region`, translate attribute keys to CFN names (they already are canonical), apply the tagging
  transform (§3.4), `CreateResource` with a fresh `ClientToken`, then **await** the request by polling
  `GetResourceRequestStatus` (honouring `RetryAfter`, bounded by the handler's `timeoutInMinutes`, default 120).
  Once sent, cancellation never abandons the wait: the host waits for the real answer. On `SUCCESS`, read back with
  `GetResource` (patience for propagation) and return that state. On `FAILED` **with an identifier**, read back; if
  the resource exists, return its state (the host drops an errored create's result, so an error would orphan it) and
  log the failure to stderr; if it does not, return the error.
- **Read:** `GetResource`. `ResourceNotFoundException` → `(nil, nil)` after bounded patience. Write-only properties
  are carried forward from current state (AWS never returns them).
- **Update:** compute an RFC 6902 patch at top-level property granularity: `add` or `replace` for each property in
  the desired state whose value differs from current. No `remove`: every settable property is Optional+Computed, so
  a property dropped from configuration produces no diff in infrata and its current AWS value is kept, as PLAN §14.1
  specifies (state cannot say whether configuration ever set it). Skip `readOnly`. `UpdateResource` with a
  `ClientToken`, await, read back.
- **Delete:** `DeleteResource` with a `ClientToken`, await. `NotFound` at any point is success.
- **Discover:** for each `discover_regions` × requested type with a `list` handler and no required
  `handlerSchema`: `ListResources` (paginated, `ctx` checked between pages), then `GetResource` per identifier (list
  results carry only the identifier for many types). Types that need a parent model are skipped and named once on
  stderr.
- **Import:** `GetResource` for `<region>/<identifier>`.
- **Provider ID:** `<region>/<identifier>`, split at the first `/` only, because identifiers can be ARNs and
  composite identifiers use `|`.
- **ClassifyError:** Cloud Control exceptions and handler error codes → infrata classes. `ThrottlingException`,
  `ConcurrentOperationException`, `ResourceConflictException` → `SafeToRetry` (refused before acting; a create also
  carries a token). `HandlerInternalFailureException`, `ServiceInternalErrorException`, `NetworkFailureException`,
  `NotStabilizedException`, HTTP ≥ 500, timeouts → `ConditionallyRetryable`. Everything else → `NotSafeToRetry`.
  The SDK's own retryer stays on for every Cloud Control call: every mutation carries a `ClientToken`.

### 3.4 Values

- **Top level:** infrata values ↔ JSON directly (`KindMap` ↔ object, `KindList` ↔ array).
- **Nested keys are AWS's exact names** (`{Key: team, Value: platform}`): infrata does not canonicalise inside map
  and list values, and a spelling that differs from what `GetResource` returns would diff forever. The provider
  refuses an unknown nested key with the accepted names (§7 Q2).
- **Tags:** when a schema's `tagging.tagProperty` is a list of `{Key, Value}`, the attribute is exposed as a
  `KindMap` (`tags: {team: platform}`) and translated both ways, because every tagged type has the same shape and
  a map is what people write.
- Values AWS normalises (casing, CIDR forms) and `propertyTransform` rules are a known source of perpetual diffs;
  handled per type as found, recorded in the overlay.

### 3.5 Kept from `first-slice`

Instance configuration (`profile`, `assume_role_arn`, `discover_regions`, fail-closed keys), credentials loading,
per-region clients (now Cloud Control clients), the `region` model (Required + ForceNew, from `defaults:`), the
patience helper, the error-message wrapper, the release/CI/manifest plan, and the live-suite guard. Removed: the
handwritten `aws.vpc`/`aws.subnet` code, `internal/ec2fake`, EC2-specific classification tests.

### 3.6 Testing

- **Generator:** golden tests on a handful of committed real schemas (VPC, Subnet, S3 Bucket, IAM Role, RDS
  DBInstance, one with `Region`, one with `Type`), asserting flags, spellings, name assignment and lock stability
  (adding a clashing type never renames a locked one).
- **Catalog:** every definition passes `schema.Validate` and `plugintest.Open` (the whole catalog loads in infrata's
  host).
- **Provider:** against `internal/ccfake`, an awsJson1_0 server (`X-Amz-Target: CloudApiService.<Op>`,
  `application/x-amz-json-1.0`) holding resources in memory, returning `IN_PROGRESS` then `SUCCESS` across polls,
  adding provider-chosen properties on read (to exercise Optional+Computed), honouring `ClientToken` idempotency,
  with fault and delay injection. The real SDK is the oracle for its wire format, as with `ec2fake`.
- **e2e:** infrata v0.3.0 built from source against the binary and the fake (`AWS_ENDPOINT_URL_CLOUDCONTROL`, checked in `service/cloudcontrol` v1.38.0 `endpoints.go`),
  covering plan/apply/re-plan clean, an unset provider-chosen attribute staying clean, drift, replacement from a
  create-only change, tags as a map, discover and import, destroy.
- **Live** (`-tags live`, account guard, never CI): VPC + Subnet + Security Group through the real API, then IAM Role
  (a global type).

## 4. Out of scope for the first release

Types needing a parent model for discovery (they still work for create/read/update/delete); nested create-only
pointers and `conditionalCreateOnly`; the Smithy escape hatch for non-provisionable types; private and third-party
types; per-region schema differences (the us-east-1 bundle is authoritative).

## 5. Risks

- **Load cost on every command** (~1.7 MB of schemas through the pipe). Measured first; fallback is dropping
  descriptions (~0.9 MB) or asking infrata for lazy schema loading.
- **Perpetual diffs** from AWS normalising values or returning nested defaults inside set objects. Mitigation per type,
  found by the e2e and live suites; Optional+Computed only covers top-level unset properties.
- **Cloud Control latency** (12.5 s for a VPC create in the spike) and handler quality varying by type.
- **Region-specific schemas** differ from us-east-1's.

## 6. Build order (becomes the plan)

1. Measure load cost with a generated catalog behind a no-op plugin. Stop and report if unacceptable.
2. Generator + lock + overlay + golden tests; commit the generated catalog.
3. Adapt Tasks 1–4 onto infrata v0.3.0 (`protocol: [2]`), catalog-backed `Definitions`.
4. `internal/ccfake`.
5. Generic provider: create/read/delete, then update (patch), then discover/import, then classification.
6. Tags transform and value translation.
7. e2e, then release plumbing, then the live suite.

## 7. Open questions for James

- **Q1 — displayed name when there is no curated alias.** infrata's `Display` shows the FIRST alias, and every property
  gets a snake_case alias, so without a curated alias the display is `cidr_block`, not `CidrBlock`. Recommended:
  accept that (lowercase, matches infrata's own examples). The alternative, showing AWS's name, means not generating
  snake_case aliases or asking infrata for a separate display field.
- **Q2 — nested keys.** Recommended: AWS's exact names inside objects and lists for the first release, plus the
  generic tags-as-map transform. Accepting other spellings inside nested values would need the plugin to rewrite them
  to match what AWS returns, which is doable later.
- **Q3 — clashes with infrata keywords.** Recommended: `type_value`/`provider_value`/`lifecycle_value` as the shown name
  for those properties, and `aws_region` for the plugin's region on the 4 regional types that have their own `Region`.
- **Q4 — which types get curated aliases first.** Recommended: EC2 networking (VPC, Subnet, SecurityGroup,
  InternetGateway, RouteTable, Route), EC2 Instance, S3 Bucket, IAM Role/Policy, RDS DBInstance/DBSubnetGroup, Lambda
  Function, ECS Cluster/Service; everything else ships with AWS names and snake_case until someone needs better.
