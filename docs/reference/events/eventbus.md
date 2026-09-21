# aws.eventbus

**CloudFormation type:** `AWS::Events::EventBus`

Resource type definition for AWS::Events::EventBus

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Events::EventBus)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) for the event bus. |
| `DeadLetterConfig` | dead_letter_config | `map` | optional, computed, provider-chosen |  | Dead Letter Queue for the event bus. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the event bus. |
| `EventSourceName` | event_source_name | `string` | optional, computed, provider-chosen, write-only |  | If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen |  | Kms Key Identifier used to encrypt events at rest in the event bus. |
| `LogConfig` | log_config | `map` | optional, computed, provider-chosen |  | The logging configuration settings for vended logs. |
| `Name` |  | `string` | required, replaces on change |  | The name of the event bus. |
| `Policy` |  | `string` | optional, computed, provider-chosen |  | A JSON string that describes the permission policy statement for the event bus. |
| `Tags` |  | `map` | tags map |  | Any tags assigned to the event bus. |

Supports update: yes

Discovery: supported
