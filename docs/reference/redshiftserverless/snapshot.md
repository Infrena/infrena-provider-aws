# aws.snapshot

**CloudFormation type:** `AWS::RedshiftServerless::Snapshot`

Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.

Region attribute: `region`

**Import ID:** `<region>/SnapshotName` (AWS::RedshiftServerless::Snapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `NamespaceName` | namespace_name | `string` | optional, computed, provider-chosen, replaces on change |  | The namespace the snapshot is associated with. |
| `OwnerAccount` | owner_account | `string` | computed |  | The owner account of the snapshot. |
| `RetentionPeriod` | retention_period | `integer` | optional, computed, provider-chosen |  | The retention period of the snapshot. |
| `Snapshot` |  | `map` | computed |  | Definition for snapshot resource |
| `SnapshotName` | snapshot_name | `string` | required, replaces on change |  | The name of the snapshot. |
| `Tags` |  | `map` | replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
