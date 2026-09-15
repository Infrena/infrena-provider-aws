# aws.eventtype

**CloudFormation type:** `AWS::FraudDetector::EventType`

A resource schema for an EventType in Amazon Fraud Detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::EventType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the event type. |
| `CreatedTime` | created_time | `string` | computed |  | The time when the event type was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the event type. |
| `EntityTypes` | entity_types | `list` | required |  |  |
| `EventVariables` | event_variables | `list` | required |  |  |
| `Labels` |  | `list` | required |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The time when the event type was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name for the event type |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Tags associated with this event type. |

Supports update: yes

Discovery: supported
