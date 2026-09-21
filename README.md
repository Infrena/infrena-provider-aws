# infrena-provider-aws

The AWS provider for [infrena](https://github.com/infrena/infrena), distributed as a plugin binary:
`infrena-plugin-aws`. It is one generic provider serving every resource type AWS Cloud Control API
supports, generated from AWS's published CloudFormation resource schemas — there is no handwritten
`aws.vpc` or `aws.subnet` anymore. `infrena explain <type>` is the reference for any one of them: it
prints every attribute, its spellings, and the type's import ID shape.

**Documentation:** [service guides, a reference page for every type, examples and modules](docs/README.md).

## Installing it

Every release publishes a binary for each supported platform, named the way `infrena plugins install`
expects: `infrena-plugin-aws_<version>_<goos>_<goarch>.tar.gz`, `.zip` for Windows, alongside a
`SHA256SUMS` file covering all of them.

```bash
infrena plugins install aws
```

To install by hand, download the archive for your platform from the
[releases page](https://github.com/infrena/infrena-provider-aws/releases), check it against
`SHA256SUMS`, and put the extracted `infrena-plugin-aws` binary where infrena looks for plugins:

```bash
sha256sum --check --ignore-missing SHA256SUMS
tar -xzf infrena-plugin-aws_<version>_linux_amd64.tar.gz
mkdir -p ~/.local/share/infrena/plugins
mv infrena-plugin-aws_<version>_linux_amd64/infrena-plugin-aws ~/.local/share/infrena/plugins/
```

infrena looks for `infrena-plugin-aws` in, in order: `--plugin-dir`, `INFRENA_PLUGIN_PATH`,
`<project>/.infrena/plugins`, `~/.local/share/infrena/plugins`, then `$PATH`.

Each release states the infrena versions it works with. `plugin.yaml`'s `infrena:` floor is the oldest
host that accepts the binary, which is not the same as the version the plugin is built against — see
[the roadmap](docs/ROADMAP.md) for what each release changed.

## Building it from source

```bash
go work init . ../infrena          # once, for local work against a sibling infrena checkout (go.work is gitignored)
go build ./cmd/infrena-plugin-aws
```

Put the built binary in one of the directories above, or pass `--plugin-dir` pointing at it.

## Type names

`aws.<resource>` when the CloudFormation resource segment is unique across every supported type
(`AWS::EC2::VPC` → `aws.vpc`), otherwise `aws.<service>.<resource>` (`AWS::EC2::Instance` →
`aws.ec2.instance`). A name never changes once released: a later AWS type that would clash with an
already-assigned short name gets the qualified form instead, and the short name keeps its original
owner.

## Attribute names

Every attribute accepts AWS's own property name in any case (`CidrBlock`), its generated snake_case
form (`cidr_block`), and a friendly alias where one is curated (`cidr`). Plans, `infrena explain` and
`import --generate` show the friendly alias when one exists, otherwise the snake_case form. A property
whose name would collide with an infrena resource keyword shows as `type_value`, `provider_value` or
`lifecycle_value` instead of `type`, `provider` or `lifecycle`. A type that has its own AWS `Region`
property takes the plugin's region under `aws_region` instead of `region`. Nested keys (inside objects
and lists) accept the same spellings as top level. `tags:` is always a map (`{key: value}`), even for
types whose underlying AWS property is a list of `{Key, Value}` pairs.

## Values AWS chooses

An attribute you leave unset keeps whatever value AWS assigns or defaults it to, and never plans a
change on later runs. Removing an attribute from configuration after AWS has set it keeps AWS's current
value rather than resetting it — infrena has no way to tell "never set" from "no longer configured," so
the provider treats both the same way.

## Credentials

The default AWS credential chain (environment, shared config, SSO, instance role) applies unless the
provider instance sets `profile` (a name in `~/.aws/config`) or `assume_role_arn` (a role ARN to assume
on top of the resolved credentials). Use one provider instance per AWS account. `infrena import
--provider <instance> <type> <id>` imports through a specific instance when more than one is configured.

## Regions

```yaml
providers:
  - plugin: aws
    defaults:
      region: ${var.aws_region}
```

`aws_region` is a variable with a `default:`, overridden per environment — `defaults:` must resolve
without an environment because `discover` takes none. Any resource can override `region:` on its own.
Changing the default region replaces every resource that inherits it, since region is `ForceNew`. Global
types (IAM, Organizations, CloudFront, Route 53) have no region attribute at all; their provider IDs are
`global/<identifier>`, and Cloud Control is always called for them in `us-east-1`.

## Discovery

```yaml
    discover_regions: [us-east-1]
    discover_types: [aws.vpc, aws.subnet, aws.securitygroup]
```

`discover_regions` is a literal list of regions to scan. `discover_types` is a literal list of infrena
type names; when unset, discovery scans a curated default set (VPCs, subnets, security groups, S3
buckets, IAM roles, and similar commonly-used types) rather than every type the catalog knows, because
an unfiltered scan is roughly 1,500 `ListResources` calls per region per run. `import` discovers first,
so the same set and the same `discover_types` override apply there. Types that only list under a parent
resource (for example a resource that only exists nested under another) are never discovered; import
them directly by ID instead. infrena names a discovered resource from its `Name` tag when it has one
(`vpc-app1` for a VPC tagged `Name: app1`), falling back to its sanitised provider ID otherwise; either
way the name is prefixed with the type's last segment. `import --generate` writes a reference (for
example `vpc: ${vpc-app1}`) instead of a literal id wherever an imported resource's attribute points at
another resource imported in the same run.

### Resources something else owns

Discovery marks the resources whose lifecycle belongs to something other than this project — AWS itself, a
CloudFormation stack, an Auto Scaling group — and `import` leaves a marked resource out unless you name it:
`infrena import dev` never adopts your default VPC by accident, while
`infrena import dev aws.vpc.us-east-1/vpc-0abc123` still adopts it deliberately. Each one is reported with the
evidence behind the claim:

| Marked | Evidence |
| --- | --- |
| the default VPC | EC2 reports it as the region's default (`DescribeVpcs`, filtered `is-default`) |
| a default subnet | EC2 reports `DefaultForAz` for it (`DescribeSubnets`) |
| the default security group | its `GroupName` is `default` |
| a service-linked role | its `Path` starts `/aws-service-role/` |
| anything a CloudFormation stack owns | the `aws:cloudformation:stack-id` tag AWS writes on it |
| an instance a fleet or Auto Scaling group launched | one of the `aws:autoscaling:groupName`, `aws:ec2:fleet-id`, `aws:ec2spot:fleet-request-id` tags AWS writes on it |

Adopting a fleet-launched instance is the one that looks like a bug days later rather than immediately: the
controller replaces the instance on its own schedule — minutes, on a spot fleet draining short jobs — so the
recorded instance id stops existing, and every plan afterwards proposes creating it again. infrena cannot manage a
lifecycle it does not control. A spot instance you launched yourself carries none of these tags and is adopted
normally.

Nothing is marked on a guess: a VPC is not a default because its CIDR is 172.31.0.0/16, or because of what it is
called; an instance is not fleet-launched because it is spot, or because it was launched from a launch template.
Cloud Control cannot answer the VPC and subnet questions at all — AWS's published schemas carry no
`IsDefault` on `AWS::EC2::VPC` and no `DefaultForAz` on `AWS::EC2::Subnet` — so the plugin asks EC2, at most once
per region per run and only while a VPC or subnet is being discovered. If that call fails, because the credentials
lack `ec2:DescribeVpcs` or `ec2:DescribeSubnets`, discovery still returns everything it found; it marks nothing and
says so on stderr.

## Import IDs

`<region>/<identifier>` for a regional type (`us-east-1/vpc-0abc123`), `global/<identifier>` for a
global type. A type with a composite primary identifier joins its parts with `|`
(`us-east-1/cert-authority-arn|certificate-arn`). `infrena explain <type>` prints the exact shape under
its "Import ID" section.

## A worked example

This is `e2e/testdata/basic/infrena.yml`, exactly as the e2e suite runs it:

```yaml
project: demo

infrena: ">= 0.5"

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
      region: ${var.aws_region}

resources:
  vpc:
    type: aws.vpc
    cidr: 10.0.0.0/16
    tags:
      team: platform

  private_a:
    type: aws.subnet
    vpc: ${vpc}
    cidr: 10.0.1.0/24
    az: us-east-1a

  web:
    type: aws.securitygroup
    description: web servers
    vpc_id: ${vpc}
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

### Passing a whole resource

`VpcId: ${vpc}` passes the VPC whole. The plugin declares that a subnet's `VpcId` holds a VPC's `VpcId`, so
infrena reads that attribute for you, and a VPC's id cannot be confused with its arn. Writing the attribute out,
`${vpc.vpc_id}`, still works, and either spelling is refused at compile time if `vpc` is not an `aws.vpc`. Where
the plugin declares no reference, `${vpc}` is a compile error telling you to name the attribute you mean. The
declarations are derived from AWS's property names and reviewed; `infrena explain <type>` shows each one as
"refers to". A list of ids takes whole resources too (`SubnetIds: [...]` with `- ${private_a}` items).

A whole-resource reference into an attribute written with a friendly alias or snake_case name (`vpc: ${vpc}`,
`vpc_id: ${vpc}`) needs infrena v0.6.2 or newer. On infrena 0.6.0 and 0.6.1, write the attribute receiving `${vpc}`
with AWS's own name instead (`VpcId: ${vpc}`, any case): those releases look up the receiving attribute's
declaration before folding aliases, so the aliased spellings wrongly report that the attribute declares no
reference.

## Breaking changes in 0.2.0

- Needs infrena 0.7.0 or newer (plugin protocol 4).
- Attributes that now declare a reference type-check explicit references, so a mismatched target is an error where
  it previously was not: for example `vpc_id: ${subnet.arn}` is refused, because a security group's `vpc_id` refers
  to a VPC, not a subnet. The full list of attributes that declare a reference is `gen/references.lock.json`'s
  `accepted` and `approved` entries.

## Behaviour worth knowing

- Creates and deletes wait for AWS to finish the operation before returning; a VPC takes about 12
  seconds for real (a fake create/delete in tests is instant).
- A create that fails after AWS has actually created something is still recorded as existing, not left
  orphaned: the provider reads the resource back and reports its true state, logging the failure to
  stderr.
- A read may retry for up to about 4 seconds before reporting a resource gone, to cover AWS's eventual
  consistency on a resource created moments earlier.

## Regenerating the catalog

The catalog (`internal/catalog/catalog.json.gz`) is generated, not written by hand:

```bash
scripts/fetch-schemas              # downloads AWS's schema bundle into schemas/ (gitignored)
go run ./cmd/gen-cloudcontrol       # reads it plus gen/overlay.yaml and gen/names.lock.json, writes the catalog
```

`gen/overlay.yaml` is the hand-curated part: which types are global, friendly aliases, sensitive
properties, and the default discovery set. `gen/names.lock.json` is the committed type-name map; it only
grows, and a name already in it never moves. A weekly workflow re-runs both scripts and opens a PR with
the diff when AWS's schemas changed.

## Building and testing

```bash
go work init . ../infrena                                            # local work against a sibling checkout
go build ./cmd/infrena-plugin-aws
go test -count=1 ./...                                                # unit + ccfake + protocol tests, no AWS account
go test -tags e2e -count=1 -v ./e2e/                                  # a real infrena binary against this plugin
go test -tags live -count=1 -v ./live/                                # REAL AWS; needs INFRENA_AWS_LIVE_* and a deliberate decision to spend
GOWORK=off go test -count=1 ./...                                     # the pinned build CI blocks on
```

## Releasing

`plugin.yaml` carries the manifest: version, protocol, platforms, and the `infrena: ">= 0.5.0"` floor.
`scripts/release-check vX.Y.Z` refuses to release unless the git tag, `plugin.yaml`'s version, and the
built binary's own reported version all agree. Outside a stamped release the binary reports version
`0.0.0-dev`.

## Writing another plugin

This repository is one worked example. For the authoring guide and a smaller reference
implementation, see the fake plugin's
[docs/writing-a-provider.md](https://github.com/infrena/infrena-provider-fake/blob/main/docs/writing-a-provider.md).

---

## Licence

Apache 2.0 — see [LICENSE](LICENSE).

Contributions are welcome; see [CONTRIBUTING.md](CONTRIBUTING.md), which explains the
[Contributor License Agreement](CLA.md) and why an open core project asks for one. Everyone
taking part is expected to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

Found a security problem? Please do not open an issue — [SECURITY.md](SECURITY.md) has the
private reporting path.
