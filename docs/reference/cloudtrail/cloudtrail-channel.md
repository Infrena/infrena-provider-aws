# aws.cloudtrail.channel

**CloudFormation type:** `AWS::CloudTrail::Channel`

A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.

Region attribute: `region`

**Import ID:** `<region>/ChannelArn` (AWS::CloudTrail::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelArn` | channel_arn | `string` | computed |  | The Amazon Resource Name (ARN) of a channel. |
| `Destinations` |  | `list` | optional, computed, provider-chosen |  | One or more resources to which events arriving through a channel are logged and stored. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the channel. |
| `Source` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of an on-premises storage solution or application, or a partner event source. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
