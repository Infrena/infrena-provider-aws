# aws.multiregionaccesspointpolicy

**CloudFormation type:** `AWS::S3::MultiRegionAccessPointPolicy`

The policy to be attached to a Multi Region Access Point

Region attribute: `region`

**Import ID:** `<region>/MrapName` (AWS::S3::MultiRegionAccessPointPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `MrapName` | mrap_name | `string` | required, replaces on change |  | The name of the Multi Region Access Point to apply policy |
| `Policy` |  | `map` | required |  | Policy document to apply to a Multi Region Access Point |
| `PolicyStatus` | policy_status | `map` | computed |  | The Policy Status associated with this Multi Region Access Point |

Supports update: yes

Discovery: supported
