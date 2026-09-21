# aws.channelplacementgroup

**CloudFormation type:** `AWS::MediaLive::ChannelPlacementGroup`

Definition of AWS::MediaLive::ChannelPlacementGroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id|ClusterId` (AWS::MediaLive::ChannelPlacementGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the channel placement group. |
| `Channels` |  | `list` | computed |  | List of channel IDs added to the channel placement group. |
| `ClusterId` | cluster_id | `string` | optional, computed, provider-chosen, replaces on change | aws.medialive.cluster.Id | The ID of the cluster the node is on. |
| `Id` |  | `string` | computed |  | Unique internal identifier. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the channel placement group. |
| `Nodes` |  | `list` | optional, computed, provider-chosen |  | List of nodes added to the channel placement group |
| `State` |  | `string` | computed |  | The current state of the ChannelPlacementGroupState |
| `Tags` |  | `map` | tags map |  | A collection of key-value pairs. |

Supports update: yes

Discovery: supported (parent resource required)
