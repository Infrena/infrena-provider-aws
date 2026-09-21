# aws.replicationset

**CloudFormation type:** `AWS::SSMIncidents::ReplicationSet`

Resource type definition for AWS::SSMIncidents::ReplicationSet

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMIncidents::ReplicationSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the ReplicationSet. |
| `DeletionProtected` | deletion_protected | `boolean` | optional, computed, provider-chosen |  | Configures the ReplicationSet deletion protection. |
| `Regions` |  | `list` | required |  | The ReplicationSet configuration. |
| `Tags` |  | `map` | tags map |  | The tags to apply to the replication set. |

Supports update: yes

Discovery: supported
