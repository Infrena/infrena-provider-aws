# aws.medialive.cluster

**CloudFormation type:** `AWS::MediaLive::Cluster`

Definition of AWS::MediaLive::Cluster Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaLive::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Cluster. |
| `ChannelIds` | channel_ids | `list` | computed |  | The MediaLive Channels that are currently running on Nodes in this Cluster. |
| `ClusterType` | cluster_type | `string` | optional, computed, provider-chosen, replaces on change |  | The hardware type for the cluster. |
| `Id` |  | `string` | computed |  | The unique ID of the Cluster. |
| `InstanceRoleArn` | instance_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The IAM role your nodes will use. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The user-specified name of the Cluster to be created. |
| `NetworkSettings` | network_settings | `map` | optional, computed, provider-chosen |  | On premises settings which will have the interface network mappings and default Output logical interface |
| `State` |  | `string` | computed |  | The current state of the Cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of key-value pairs. |

Supports update: yes

Discovery: supported
