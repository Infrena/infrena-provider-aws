# aws.label

**CloudFormation type:** `AWS::FraudDetector::Label`

An label for fraud detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::Label)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The label ARN. |
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the label was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The label description. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the label was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the label. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with this label. |

Supports update: yes

Discovery: supported
