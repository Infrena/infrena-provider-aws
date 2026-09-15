# aws.pcs.queue

**CloudFormation type:** `AWS::PCS::Queue`

AWS::PCS::Queue resource creates an AWS PCS queue.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::PCS::Queue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The unique Amazon Resource Name (ARN) of the queue. |
| `ClusterId` | cluster_id | `string` | required, replaces on change | aws.pcs.cluster.Id | The ID of the cluster of the queue. |
| `ComputeNodeGroupConfigurations` | compute_node_group_configurations | `list` | optional, computed, provider-chosen |  | The list of compute node group configurations associated with the queue. Queues assign jobs to associated compute node groups. |
| `ErrorInfo` | error_info | `list` | computed |  | The list of errors that occurred during queue provisioning. |
| `Id` |  | `string` | computed |  | The generated unique ID of the queue. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name that identifies the queue. |
| `SlurmConfiguration` | slurm_configuration | `map` | optional, computed, provider-chosen |  | The Slurm configuration for the queue. |
| `Status` |  | `string` | computed |  | The provisioning status of the queue. The provisioning status doesn't indicate the overall health of the queue. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | 1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string. |

Supports update: yes

Discovery: supported (parent resource required)
