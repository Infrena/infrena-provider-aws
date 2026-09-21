# aws.node

**CloudFormation type:** `AWS::MediaLive::Node`

Definition of AWS::MediaLive::Node Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaLive::Node)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Node. It is automatically assigned when the Node is created. |
| `ChannelPlacementGroups` | channel_placement_groups | `list` | computed |  | An array of IDs. Each ID is one ChannelPlacementGroup that is associated with this Node. |
| `ClusterId` | cluster_id | `string` | required, replaces on change | aws.medialive.cluster.Id | The ID of the Cluster that the Node belongs to. |
| `ConnectionState` | connection_state | `string` | computed |  | The current connection state of the Node. |
| `Id` |  | `string` | computed |  | The unique ID of the Node. Unique in the Cluster. The ID is the resource-id portion of the ARN. |
| `InstanceArn` | instance_arn | `string` | computed |  | The ARN of the EC2 instance hosting the Node. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The user-specified name of the Node. |
| `NodeInterfaceMappings` | node_interface_mappings | `list` | optional, computed, provider-chosen, replaces on change |  | An array of interface mappings for the Node. |
| `Role` |  | `string` | optional, computed, provider-chosen |  | The role of the Node in the Cluster. ACTIVE means the Node is available for encoding. BACKUP means the Node is a redundant Node and might get used if an ACTIVE Node fails. |
| `SdiSourceMappings` | sdi_source_mappings | `list` | optional, computed, provider-chosen |  | An array of SDI source mappings. |
| `State` |  | `string` | computed |  | The current state of the Node. |
| `Tags` |  | `map` | tags map |  | A collection of key-value pairs. |

Supports update: yes

Discovery: supported (parent resource required)
