# Can the AWS plugin be generic instead of handwritten per resource?

**Date:** 2026-09-13. **Branch:** `spike-generic-aws` (off `first-slice`). **Proofs of concept:** `spikes/generic-aws/`,
all throwaway. **Pinned versions:** aws-sdk-go-v2 v1.47.0, service/ec2 v1.332.0, service/cloudcontrol v1.38.0,
service/cloudformation v1.81.0, smithy-go v1.28.1, infrata v0.2.0 (host tests) / `main` at `5bc47a3`.

Every claim below was checked against source, the account, or AWS documentation. The verification log is at the end.
Claims that came back different from what was assumed are marked ✘ there.

---

## 1. Conclusion

**Yes. A generic AWS provider is possible, and it has already been proven live.** The layer that makes it possible is
not the one the question assumed, though.

- **Smithy models describe OPERATIONS, not RESOURCES.** smithy-go can invoke any AWS operation generically at runtime:
  the published model, turned into smithy-go's own runtime schemas, drives smithy-go's own protocol codecs. That was
  proven against real EC2. But the models almost never say which operations create, read, update and delete a given
  resource, which is the thing infrata needs. Of the models checked, only Lambda declares any `resource` shapes (11).
  EC2, S3, IAM, DynamoDB and Cloud Control declare none. An operation-level generic layer alone would still need a
  handwritten mapping for every resource. **That is the premise that is wrong.**
- **AWS already publishes the resource layer: CloudFormation resource type schemas, served through the AWS Cloud
  Control API.** Cloud Control is one generic create/read/update/delete/list API over 1,426 fully mutable `AWS::`
  resource types, plus 159 immutable (create/delete only) types. Each type's JSON schema declares its read-only,
  create-only and write-only properties, its required properties and its primary identifier, which map almost directly
  onto infrata's `Computed`, `ForceNew` and `Required` flags. **Proven live:** one `AWS::EC2::VPC` was created, patched,
  read and deleted from a type-name string and a JSON document, with no VPC code. The real VPC and subnet schemas
  mechanically produced the same flags Task 1 hand-wrote, and infrata's own plugin host accepted the generated
  definitions. This is how HashiCorp's `terraform-provider-awscc` and Pulumi's `aws-native` providers are built.
- **The main obstacle found is on infrata's side, not AWS's:** infrata has no *optional + computed* attribute (§5, §8).

## 2. Resource vs operation — where the distinction matters

| | AWS API operation | AWS resource |
| --- | --- | --- |
| What it is | One RPC: `CreateVpc`, `ModifyVpcAttribute`, `DescribeVpcs`, `CreateTags` | A thing with identity and lifecycle: a VPC |
| Where AWS describes it | Smithy model (`aws/api-models-aws`; vendored in `aws-sdk-go-v2/codegen/sdk-codegen/aws-models`, 431 services) | CloudFormation resource type schema (registry; `DescribeType`, or the public bundle `https://schema.cloudformation.<region>.amazonaws.com/CloudformationSchema.zip`) |
| What it knows | Input/output shapes, protocol traits, errors, paginators, waiters (EC2: 802 operations, 170 paginated, 26 waitable) | Properties, which are read-only / create-only / write-only, identifier, required, handlers and their IAM permissions |
| What it does NOT know | Which operations make up one resource's create, read, update and delete. Which fields force replacement. Which fields AWS fills in. How to read back what a create made. | Wire formats (it rides on Cloud Control), operations that are not resources (e.g. `RebootInstances`) |

Why it matters to infrata: infrata's `Provider` interface is resource-shaped (`Create/Read/Update/Delete/Discover/Import`,
per type). One infrata resource is usually **several** AWS operations. Creating a VPC with DNS hostnames on is
`CreateVpc` then `ModifyVpcAttribute`, and reading it back is `DescribeVpcs` plus `DescribeVpcAttribute` per attribute.
Cloud Control's real VPC schema lists exactly that, in its handler permissions:
`create: ec2:CreateVpc, ec2:ModifyVpcAttribute, ec2:DescribeVpcAttribute, …`. An operation invoker removes the
per-operation Go code. Only a resource layer removes the per-resource *knowledge*.

## 3. The approaches, evaluated

### A. Runtime discovery and invocation with the SDK and Smithy runtime

**How it works (proven, `spikes/generic-aws/smithyinvoke`).** Load a service's Smithy JSON model at runtime. Build
`*smithy.Schema` values for an operation's input and output with smithy-go's public `smithy.NewSchema` and
`Schema.AddMember`, carrying the protocol traits (`xmlName`, `ec2QueryName`, `xmlFlattened`, `jsonName`). A small generic
value implements `smithy.Serializable` / `smithy.Deserializable` by walking a `map[string]any` against the schema.
smithy-go's own `transport/http/protocol/ec2query` codec (public; `awsjson`, `awsquery`, `restjson1`, `restxml` and
`rpcv2` exist beside it) builds the request. The request is then signed with `aws/signer/v4` and sent.

- **What is available:** everything about an operation. smithy-go v1.27.0 added "APIs for schema-based serialization"
  and "support for all current AWS and Smithy protocols". The newest generated clients already run on it: Cloud Control
  v1.33.0 (2026-08-25) "Enable schema-based (de)serialization for this service", and ships `schemas/` plus a
  `TypeRegistry`. EC2 v1.332.0 has not been switched yet, which is why the spike builds its schemas from the model.
- **Proven:** `CreateVpc` with nested `TagSpecifications`, `DescribeVpcs`, `DeleteVpc`, `CreateSubnet`, and
  `ModifySubnetAttribute` with a nested struct, all against the fake. `DescribeVpcs` and `DescribeAvailabilityZones`
  against **real EC2**, signed and accepted, with the nested output decoded. A real `InvalidVpcID.NotFound` came back
  as a `smithy.APIError`. A misspelt member is refused before any request, from the model. Sabotage: dropping the
  `xmlName`/`ec2QueryName` trait conversion silently stopped tags serialising, so trait fidelity is load-bearing.
- **What it cannot discover:** resources (above). There are also no SDK middleware conveniences: retries, endpoint
  rules (S3's are complex), checksums, presigning, and error type registries. A generated client assembles those per
  service, and a runtime invoker must reassemble or give them up.
- **Code to maintain:** the invoker is ~400 lines for one protocol, plus per-protocol wiring and middleware. It needs
  **no** code per operation.
- **New services/operations:** update the model file. No Go changes.
- **Arbitrary operations:** yes, for the protocols smithy-go ships. Event streams and streaming blobs need more.
- **Practical for production?** As an **escape hatch** for operations, yes. As the resource provider, no: it answers the
  wrong question. Its risk is maturity. The schema-serde APIs are about three months old, and a model file (EC2's is
  8 MB) must be shipped or embedded.

### B. Go reflection over generated SDK types

**How it works (proven, `spikes/generic-aws/reflectinvoke`).** `reflect.ValueOf(client).MethodByName("CreateVpc")`, build
the `*XInput` with `reflect.New`, fill it from a map through `encoding/json` (field names match, case-insensitively),
call, and marshal the output back to a map.

- **Available:** exactly what the generated Go types expose. There is no model metadata: no traits, required-ness,
  enums as values, or documentation.
- **Cannot discover:** resources. Nor can it validate input: **a misspelt member (`CidrBlok`) was silently dropped and
  the call still went out** (shown in the test log). Enum values, unions and timestamps depend on JSON coercions lining
  up, and nothing verifies they do.
- **Maintenance:** small code. But the plugin must import every service module it could ever call. EC2 alone is 30 MB
  of source, 803 generated operation files and ~354k lines of serde and types. Binary size grows with coverage.
- **New services:** add an import and rebuild. Operations appear automatically.
- **Production?** No. It is strictly worse than A: unvalidated input, a huge binary, and no model information.

### C. Build-time code generation from Smithy models

**How it works.** Run a generator over the Smithy models, as AWS's own `smithy-go-codegen` does, to emit per-operation
serde or per-resource Go code.

- **Available:** everything in the model, at build time.
- **Cannot discover:** resources, for the same reason as A. Generating code from operations still leaves the resource
  mapping to be written by hand.
- **Maintenance:** a generator, plus generated code in bulk; AWS's own generator is a JVM Smithy build.
- **Compared with A:** A now gets the same result at runtime, through smithy-go's public schema runtime, without a
  generator. C is only worth it to avoid shipping model files, and AWS's own schema-serde clients (C's output) already
  exist to be used instead.
- **Production?** Only as a way to pre-compute A's schemas. It isn't a separate architecture.

### D. Hybrid: Smithy-generated adapters plus handwritten resource metadata

**How it works.** A or C does the operation plumbing. For each resource, handwritten metadata names which operations
implement create, read, update and delete, which fields force replacement, how to read back each attribute, and waiters.

- **Available / cannot discover:** as A, plus whatever the metadata says.
- **Maintenance:** **metadata per resource**. For EC2 VPC that metadata is essentially what Task 5 hand-wrote in code:
  `CreateVpc`, `ModifyVpcAttribute` per DNS attribute, `DescribeVpcAttribute` per attribute on read, tags via
  `CreateTags`/`DeleteTags`, ForceNew on the CIDR. Moving it from Go into YAML shrinks it without removing it.
- **New resources:** write their metadata.
- **Production?** Viable, but it re-derives what CloudFormation's resource handlers already are. **Worth it only for
  the gaps Cloud Control does not cover** (§8).

### E. AWS Cloud Control API + CloudFormation resource type schemas (not in the original list; recommended)

**How it works (proven live, `spikes/generic-aws/cmd/cloudcontrol-poc` and `cfnschema`).**

- **Operations:** eight operations cover every supported resource type: `CreateResource(TypeName, DesiredState JSON,
  ClientToken)`, `GetResource(TypeName, Identifier)`, `UpdateResource(TypeName, Identifier, PatchDocument = RFC 6902
  JSON Patch, ClientToken)`, `DeleteResource`, `ListResources(TypeName, ResourceModel?)`, and
  `GetResourceRequestStatus` / `CancelResourceRequest` / `ListResourceRequests` for the asynchronous requests.
- **Schema → infrata definition:**

  | CloudFormation schema | infrata |
  | --- | --- |
  | `readOnlyProperties` | `Computed` |
  | `createOnlyProperties`, or no `update` handler | `ForceNew` |
  | `required` | `Required` |
  | `writeOnlyProperties` | carry forward from state on read (the planner needs it; secrets are typical) |
  | `primaryIdentifier` | provider ID |

**Live result (account `infrata`, us-east-1).**

| Step | What happened | Time |
| --- | --- | --- |
| Create | `{"CidrBlock":"10.99.0.0/16","Tags":[…]}` | 12.5s |
| Update | JSON Patch setting `EnableDnsHostnames` and replacing tags, confirmed by read | 2.5s |
| Delete | removed | 2.1s |
| Read after delete | `ResourceNotFoundException` | — |

The account was checked afterwards: only its default VPC remains.

- **Available:** 1,426 fully mutable and 159 immutable `AWS::` types in us-east-1. Every core type checked is fully
  mutable: `AWS::EC2::VPC`, `Subnet`, `SecurityGroup`, `Instance`, `AWS::S3::Bucket`, `AWS::RDS::DBInstance`,
  `AWS::IAM::Role`, `AWS::Route53::RecordSet`, `AWS::ECS::Service`, `AWS::Lambda::Function`.
- **Creates are safe to retry:** `ClientToken` makes a create or delete idempotent for 36 hours. That fixes the problem
  Task 5 had to design around: raw `CreateVpc` has no token.
- **Price:** no additional charge for AWS types.
- **Cannot discover:** see §5. In short:
  - **288** `AWS::` types are `NON_PROVISIONABLE` (e.g. App Mesh, `AWS::AppStream::Fleet`).
  - Which types are global (IAM, Route 53) rather than regional.
  - Which properties are secret.
  - When a `conditionalCreateOnly` property replaces.
  - Create-only flags on *nested* properties (37 on VPC alone).
  - Which properties AWS fills in when unset — the optional + computed problem.
- **Maintenance:** one generic provider (the spike's mapping and live program are ~550 lines together), a schema bundle
  pinned at build time, and a small overlay (§6). Handwritten code, by contrast, took ~320 lines of implementation plus
  ~360 of tests for just two types on `first-slice`.
- **New resources:** update the pinned schema bundle, with no Go changes. AWS adds them upstream.
- **Arbitrary operations:** no. It covers resources only; approach A is the escape hatch.
- **Production?** **Yes**, with the caveats in §8, and with precedent: `terraform-provider-awscc` ("generated from the
  latest CloudFormation schemas, and will release weekly") and `pulumi-aws-native` ("covers all resources as
  supported by the AWS Cloud Control API. This does not yet include all AWS resources").

## 4. What the SDK / Smithy runtime can do dynamically

- **Build schemas at runtime:** `smithy.NewSchema`, `AddMember`, `NewOperationSchema`, `NewServiceSchema` are public.
  Recursive shapes work: a schema can be cached before its members are added, because member schemas share the
  target's member map.
- **Serialise and deserialise any `smithy.Serializable` / `Deserializable`** through public protocol codecs:
  `ec2query`, `awsquery`, `awsjson` (1.0/1.1), `restjson1`, `restxml`, `rpcv2`. The value type does not have to be
  generated.
- **Discover operations:** only from the model file. Generated clients expose no operation registry. `ServiceID` and
  `ServiceAPIVersion` are constants, and the operation name is available only inside a call, through
  `middleware.GetOperationName(ctx)`.
- **The model ships only with the generator, not the module:** the `service/ec2` module holds generated Go and a
  `generated.json` build manifest, not the Smithy model. Models come from `aws/api-models-aws` or from
  `aws-sdk-go-v2/codegen/sdk-codegen/aws-models`.
- **Signing:** `aws/signer/v4.Signer.SignHTTP` works on any request.
- **Deliberately not assembled by the spike:** retries (`aws/retry` is standalone and reusable), endpoint rules,
  checksums, and error registries. Those would be the production work of approach A.

## 5. What cannot be derived from AWS's models

| Needed by infrata | Smithy model | CloudFormation schema | So it comes from |
| --- | --- | --- | --- |
| Which operations are a resource's create, read, update and delete | ✘ (Lambda only, partially) | ✔ (handlers) | Cloud Control |
| ForceNew | ✘ | ✔ top-level `createOnlyProperties`. ✘ nested pointers (37 on VPC). ✘ `conditionalCreateOnly` (e.g. VPC `InstanceTenancy`, RDS `Engine`) | Schema, plus an infrata decision on nested and conditional flags |
| Computed | ✘ | ✔ `readOnlyProperties` | Schema |
| **Optional but AWS-defaulted** (e.g. VPC `EnableDnsSupport`, `InstanceTenancy`, returned by `GetResource` though never configured) | ✘ | ✘ | **infrata needs optional + computed**, or the plugin prunes |
| Sensitive | Smithy `@sensitive` trait on some shapes | ✘. `writeOnly` ≠ secret: VPC's write-only properties are `Ipv4IpamPoolId`, `Ipv4NetmaskLength`; RDS's include `MasterUserPassword` | Overlay list |
| Regional vs global type | ✘ | ✘ | Overlay: namespaces `IAM`, `Route53`, `CloudFront`, … |
| Requirements (subnet needs a VPC) | ✘ | ✘ (no cross-type references) | Overlay or none; `vpc_id` references still order things |
| Import ID shape | ✘ | ✔ `primaryIdentifier` (may be composite, `a|b`) | Schema, plus the plugin's `<region>/` prefix |
| Drift-noise transforms | ✘ | ✔ `propertyTransform` | Schema |
| Names users like (`cidr` vs `CidrBlock`) | ✘ | ✘ | A naming policy (recommend keeping AWS's names) |

## 6. What must be generated or pinned at build time

- **The CloudFormation schema bundle**, pinned per plugin release and turned into embedded infrata definitions.
  infrata requires `Definitions()` to be static and offline (every command, `validate` included, loads them), so schemas
  cannot be fetched from `DescribeType` at runtime. The bundle is a 2.9 MB zip per region (checked: us-east-1, 200 OK).
  Weekly regeneration, as `awscc` does, fits the `bump-infrata.yml` pattern already planned.
- **An overlay file** for what §5 says cannot be derived: global types, sensitive properties, optional + computed
  properties, and any requirement hints.
- **Nothing per operation.** For approach A, embed or ship only the model files for the services the escape hatch
  covers.

## 7. Proof-of-concept summary

| POC | Against | Result |
| --- | --- | --- |
| `cfnschema` tests | real VPC/Subnet schemas fetched with `DescribeType` | flags match Task 1's handwritten ones; `plugintest.Open` accepts the generated `aws.ec2.vpc`/`aws.ec2.subnet` |
| `cmd/cloudcontrol-poc -live` | real AWS | create / patch / read / delete of one VPC, timings above |
| `smithyinvoke` tests | `internal/ec2fake` | generic CreateVpc (nested tags), DescribeVpcs, DeleteVpc, CreateSubnet, ModifySubnetAttribute; strict input validation; typed AWS error codes |
| `cmd/smithyinvoke-live` | real EC2, read-only | DescribeVpcs, DescribeAvailabilityZones, NotFound error |
| `reflectinvoke` tests | `internal/ec2fake` | works; silently drops a misspelt member |

## 8. Recommendation: how to build the AWS plugin

**Build a generic Cloud Control provider as the core, with the Smithy runtime invoker as a narrow escape hatch. Stop
handwriting resources.**

1. **Codegen step** (`cmd/gen-definitions`): read the pinned CloudFormation schema bundle and the overlay, and emit
   embedded JSON definitions. Type names follow `aws.<service>.<resource>` (`AWS::EC2::VPC` → `aws.ec2.vpc`).
   Attribute names stay AWS's own (`CidrBlock`): renaming 1,500 types by hand is the handwritten work again. Every
   regional type gets the plugin's `region` attribute (Required + ForceNew); global types listed in the overlay do not.
2. **Generic `Provider`**, one implementation for every type:
   - **Create:** `CreateResource` with a `ClientToken`, then await the request, honouring `RetryAfter` and never
     abandoning a sent request on cancellation.
   - **Read:** `GetResource`. Write-only properties are carried forward from current state, and `NotFound` means gone.
   - **Update:** a JSON Patch computed from state vs desired, skipping read-only properties.
   - **Delete:** `DeleteResource`, with `NotFound` counting as success.
   - **Discover:** `ListResources` per `discover_regions`, then `GetResource` per identifier (the live VPC list returned
     only `VpcId`). `ResourceModel` covers types that need a parent.
   - **Import:** `GetResource`.
   - **ClassifyError:** map Cloud Control's `HandlerErrorCode`s (`Throttling`, `NotStabilized`,
     `ServiceInternalError`, …) and exceptions to infrata's three classes, generically.
3. **infrata change needed first — take to the infrata session:** an *optional + computed* attribute. Configuration may
   set it; if it doesn't, the provider's value is recorded and not diffed. Today the compiler refuses to set a computed
   attribute (`internal/compiler/schema.go:82`, "is computed and cannot be set"), and the planner diffs every returned
   non-computed attribute as "removed from configuration". Without it, every Cloud Control resource plans a change
   forever. (Terraform has exactly this: `Optional` + `Computed`.) The plugin-side fallback is to prune on `Read` any
   property absent from prior state, which costs drift detection on those properties. Also worth raising there: nested
   create-only flags and conditional create-only.
4. **Escape hatch (approach A), later and only when needed:** for the 288 non-provisionable types or non-resource
   actions, hand a small metadata entry (approach D) to the Smithy invoker rather than writing typed SDK code.
5. **What survives from `first-slice`:**
   - **Kept:** module layout, credentials and instance configuration (Task 2), the region model, the release, CI and
     manifest plan, error-classification discipline (Task 4's tests become Cloud Control's), and the `patience` idea.
   - **Retired:** `internal/ec2fake`'s role becomes a Cloud Control fake (a simpler `awsJson1_0` API), and the
     handwritten `aws.vpc`/`aws.subnet` (Tasks 5–6).
   - **Type names change** from `aws.vpc` to `aws.ec2.vpc`.

## 9. Risks and limitations

- **Coverage gaps:** 288 `AWS::` types are non-provisionable, and 159 are immutable (every change replaces).
  Coverage differs by region. Checks here were us-east-1 only.
- **Latency:** Cloud Control is asynchronous. The VPC create took 12.5s end to end, against a synchronous `CreateVpc`,
  so long applies get slower. Polling cost and Cloud Control quotas were not measured.
- **Schema quality varies by type:** `awscc` notes "some services use an older CloudFormation schema". Expect a
  per-type test matrix, not per-type code.
- **Property shapes follow CloudFormation, not EC2:** tags are a list of `{Key, Value}`, not a map, and nested objects are
  lists and maps in infrata values. The README and `explain` must show them as they are.
- **Identity:** Cloud Control identifiers can be composite (`a|b`). infrata import selectors must match them exactly.
- **Permissions:** the user needs Cloud Control permissions **and** each handler's underlying permissions (listed per
  type in the schema).
- **smithy-go schema-serde is new** (v1.27.0, about June 2026), and the escape hatch depends on it.
- **Everything here is one spike:** one resource type exercised live, in one region, on root credentials (to be
  replaced, per the vault follow-up).

## 10. Verification log

| Claim | Checked against | Result |
| --- | --- | --- |
| Smithy models declare resources, so a resource layer can be derived from them | shape counts in `ec2.json`, `s3.json`, `dynamodb.json`, `iam.json`, `lambda.json`, `cloudcontrol.json` (aws-sdk-go-v2 `codegen/sdk-codegen/aws-models`) | ✘ only Lambda declares resource shapes (11); others 0 |
| The SDK module ships the Smithy model | `service/ec2@v1.332.0` contents | ✘ only generated Go and `generated.json` |
| smithy-go has a public runtime schema and protocol codec API | `smithy-go@v1.28.1` `schema.go`, `serde.go`, `transport/http/protocol.go` (`ClientProtocol`), `transport/http/protocol/{ec2query,awsjson,awsquery,restjson1,restxml,rpcv2}`; CHANGELOG v1.27.0 | ✔ |
| Generated clients already use it | cloudcontrol v1.38.0 `schemas/`, `type_registry.go`, `api_client.go:220` `awsjson.New10`; CHANGELOG v1.33.0 | ✔ Cloud Control yes; EC2 v1.332.0 and CloudFormation v1.81.0 no |
| A runtime schema built from the model drives the real codec | `smithyinvoke` tests (fake) and `cmd/smithyinvoke-live` (real EC2, read-only) | ✔ |
| Protocol traits are required for correct wire output | sabotage: `xmlName`/`ec2QueryName` conversion removed | ✔ tags silently not serialised |
| Reflection can invoke generated operations generically | `reflectinvoke` test | ✔, and silently drops a misspelt member |
| `CreateVpc`/`CreateSubnet` accept a client token | ec2 model `idempotencyToken` trait scan (94 structures have one) | ✘ neither does |
| Cloud Control `CreateResource`/`UpdateResource`/`DeleteResource` take an idempotency token | cloudcontrol model; API reference "valid for 36 hours once used" | ✔ |
| Resource type schemas define read-only / create-only / write-only / identifier / handlers | CloudFormation CLI "Resource type schema" page; real `AWS::EC2::VPC` via `DescribeType` | ✔ |
| Schema flags match hand-derived ones for VPC/Subnet | `cfnschema` tests against fetched schemas | ✔ (Subnet's `AvailabilityZone` is optional in the schema; AWS picks one, which is the optional + computed case) |
| Generated definitions pass infrata's host | `plugintest.Open` in `cfnschema` tests, infrata v0.2.0 | ✔ |
| Cloud Control can create/update/read/delete a VPC generically | live run on account `infrata`, us-east-1 | ✔; `ListResourceRequests` shows CREATE/UPDATE/DELETE SUCCESS; only the default VPC remains |
| Coverage counts | `cloudformation list-types --visibility PUBLIC --type RESOURCE --provisioning-type …`, filtered to `AWS::` | 1,426 fully mutable / 159 immutable / 288 non-provisionable |
| Key types are supported | `describe-type` provisioning type for 10 types | ✔ all FULLY_MUTABLE |
| `ListResources` returns full properties | live `list-resources AWS::EC2::VPC` | ✘ identifier only for VPC; `GetResource` per item needed. `ResourceModel` input exists (model) |
| Write-only means secret | VPC and RDS schemas | ✘ not in general |
| Cloud Control is free for AWS types | aws.amazon.com/cloudcontrolapi/pricing | ✔ |
| `awscc` and `aws-native` are built on Cloud Control schemas | their GitHub READMEs | ✔ |
| Public schema bundle exists | `HEAD https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip` | ✔ 200, 2.9 MB |
| infrata lacks optional + computed | `internal/compiler/schema.go:82`; `pkg/schema/attribute.go`; `internal/planner/diff.go` `!inConfig` branch | ✔ gap confirmed |
