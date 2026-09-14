# spikes/generic-aws — THROWAWAY

Proof-of-concept code for the investigation in
[`docs/investigations/2026-09-13-generic-aws-provider.md`](../../docs/investigations/2026-09-13-generic-aws-provider.md).
Nothing here is production code or a design to copy line by line. It exists to prove or disprove whether the
AWS plugin can avoid handwritten per-resource code.

It is its own Go module so its dependencies (Cloud Control, CloudFormation) never enter the plugin's
`go.mod`. It borrows `internal/ec2fake` from the plugin through a `replace`.

| Package | Approach | What it proves |
| --- | --- | --- |
| `reflectinvoke/` | B — reflection over generated SDK clients | Any EC2 operation can be called by name with a map in and a map out. A misspelt member is silently dropped. |
| `smithyinvoke/` | A/C — runtime Smithy model → `smithy.Schema` → smithy-go's own `ec2query` codec | Any EC2 operation can be called by name with **no generated type at all**, validated against the model before sending. |
| `cmd/smithyinvoke-live/` | same, against real EC2 | Read-only `DescribeVpcs`/`DescribeAvailabilityZones` signed and accepted by AWS. |
| `cfnschema/` | E — CloudFormation resource schemas → infrata definitions | Real `AWS::EC2::VPC`/`AWS::EC2::Subnet` schemas map to the flags that were handwritten in Task 1, and infrata's host accepts them. |
| `cmd/cloudcontrol-poc/` | E — Cloud Control API, live | One VPC created, updated (JSON Patch), read and deleted from a type-name string and a JSON document. |

```bash
cd spikes/generic-aws
export GOWORK=off GOPRIVATE='github.com/infrata/*'
scripts/fetch-models                                   # the 8 MB EC2 Smithy model, gitignored
go test -count=1 ./...                                 # fake EC2 only; no AWS account
go run ./cmd/cloudcontrol-poc -profile infrata         # read-only: schema → definition
go run ./cmd/cloudcontrol-poc -profile infrata -live   # CREATES, updates and deletes one VPC
go run ./cmd/smithyinvoke-live -profile infrata        # read-only calls against real EC2
```
