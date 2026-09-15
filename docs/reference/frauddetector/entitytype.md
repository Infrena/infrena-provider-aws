# aws.entitytype

**CloudFormation type:** `AWS::FraudDetector::EntityType`

An entity type for fraud detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::EntityType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The entity type ARN. |
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the entity type was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The entity type description. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the entity type was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the entity type. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with this entity type. |

Supports update: yes

Discovery: supported
