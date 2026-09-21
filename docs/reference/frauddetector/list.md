# aws.list

**CloudFormation type:** `AWS::FraudDetector::List`

A resource schema for a List in Amazon Fraud Detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::List)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The list ARN. |
| `CreatedTime` | created_time | `string` | computed |  | The time when the list was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the list. |
| `Elements` |  | `list` | optional, computed, provider-chosen |  | The elements in this list. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The time when the list was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the list. |
| `Tags` |  | `map` | tags map |  | Tags associated with this list. |
| `VariableType` | variable_type | `string` | optional, computed, provider-chosen |  | The variable type of the list. |

Supports update: yes

Discovery: supported
