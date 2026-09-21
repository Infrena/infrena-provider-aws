# aws.pcs.cluster

**CloudFormation type:** `AWS::PCS::Cluster`

AWS::PCS::Cluster resource creates an AWS PCS cluster.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::PCS::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The unique Amazon Resource Name (ARN) of the cluster. |
| `Endpoints` |  | `list` | computed |  | The list of endpoints available for interaction with the scheduler. |
| `ErrorInfo` | error_info | `list` | computed |  | The list of errors that occurred during cluster provisioning. |
| `Id` |  | `string` | computed |  | The generated unique ID of the cluster. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name that identifies the cluster. |
| `Networking` |  | `map` | required, replaces on change |  | The networking configuration for the cluster's control plane. |
| `Scheduler` |  | `map` | required |  | The cluster management and job scheduling software associated with the cluster. |
| `Size` |  | `string` | required, replaces on change |  | The size of the cluster. |
| `SlurmConfiguration` | slurm_configuration | `map` | optional, computed, provider-chosen |  | Additional options related to the Slurm scheduler. |
| `Status` |  | `string` | computed |  | The provisioning status of the cluster. The provisioning status doesn't indicate the overall health of the cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | 1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string. |

Supports update: yes

Discovery: supported
