# aws.graphsnapshot

**CloudFormation type:** `AWS::NeptuneGraph::GraphSnapshot`

Resource Type definition for AWS::NeptuneGraph::GraphSnapshot

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::NeptuneGraph::GraphSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the graph snapshot. |
| `GraphIdentifier` | graph_identifier | `string` | required, replaces on change |  | The unique identifier of the Neptune Analytics graph to create the snapshot from. |
| `Id` |  | `string` | computed |  | The unique identifier of the graph snapshot. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | computed |  | The ID of the KMS key used to encrypt and decrypt the snapshot. |
| `SnapshotCreateTime` | snapshot_create_time | `string` | computed |  | The time when the snapshot was created. |
| `SnapshotName` | snapshot_name | `string` | required, replaces on change |  | The snapshot name. |
| `Status` |  | `string` | computed |  | The current status of the graph snapshot. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
