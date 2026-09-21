# aws.signalingchannel

**CloudFormation type:** `AWS::KinesisVideo::SignalingChannel`

Resource Type Definition for AWS::KinesisVideo::SignalingChannel

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::KinesisVideo::SignalingChannel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel. |
| `MessageTtlSeconds` | message_ttl_seconds | `integer` | optional, computed, provider-chosen |  | The period of time a signaling channel retains undelivered messages before they are discarded. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Kinesis Video Signaling Channel. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type. |

Supports update: yes

Discovery: supported
