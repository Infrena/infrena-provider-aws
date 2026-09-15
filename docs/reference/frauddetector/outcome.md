# aws.outcome

**CloudFormation type:** `AWS::FraudDetector::Outcome`

An outcome for rule evaluation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::Outcome)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The outcome ARN. |
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the outcome was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The outcome description. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the outcome was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the outcome. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with this outcome. |

Supports update: yes

Discovery: supported
