# aws.channelflow

**CloudFormation type:** `AWS::Chime::ChannelFlow`

Creates a channel flow in the Amazon Chime SDK Messaging service. A channel flow is a container for processors (Lambda functions) that perform actions on chat messages.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Chime::ChannelFlow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppInstanceArn` | app_instance_arn | `string` | required, replaces on change | aws.appinstance.AppInstanceArn | The ARN of the app instance. |
| `AppInstanceId` | app_instance_id | `string` | computed |  | The ID of the app instance, extracted from the channel flow ARN. |
| `Arn` |  | `string` | computed |  | The ARN of the channel flow. |
| `ChannelFlowId` | channel_flow_id | `string` | computed |  | The ID of the channel flow, extracted from the channel flow ARN. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The time at which the channel flow was created. |
| `LastUpdatedTimestamp` | last_updated_timestamp | `string` | computed |  | The time at which the channel flow was last updated. |
| `Name` |  | `string` | required |  | The name of the channel flow. |
| `Processors` |  | `list` | required |  | Information about the processor Lambda functions. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the channel flow. |

Supports update: yes

Discovery: supported (parent resource required)
