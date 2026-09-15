# aws.room

**CloudFormation type:** `AWS::IVSChat::Room`

Resource type definition for AWS::IVSChat::Room.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVSChat::Room)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Room ARN is automatically generated on creation and assigned as the unique identifier. |
| `Id` |  | `string` | computed |  | The system-generated ID of the room. |
| `LoggingConfigurationIdentifiers` | logging_configuration_identifiers | `list` | optional, computed, provider-chosen |  | Array of logging configuration identifiers attached to the room. |
| `MaximumMessageLength` | maximum_message_length | `integer` | optional, computed, provider-chosen |  | The maximum number of characters in a single message. |
| `MaximumMessageRatePerSecond` | maximum_message_rate_per_second | `integer` | optional, computed, provider-chosen |  | The maximum number of messages per second that can be sent to the room. |
| `MessageReviewHandler` | message_review_handler | `map` | optional, computed, provider-chosen |  | Configuration information for optional review of messages. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the room. The value does not need to be unique. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
