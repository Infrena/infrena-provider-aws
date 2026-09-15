# aws.variable

**CloudFormation type:** `AWS::FraudDetector::Variable`

A resource schema for a Variable in Amazon Fraud Detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::Variable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the variable. |
| `CreatedTime` | created_time | `string` | computed |  | The time when the variable was created. |
| `DataSource` | data_source | `string` | required |  | The source of the data. |
| `DataType` | data_type | `string` | required |  | The data type. |
| `DefaultValue` | default_value | `string` | required |  | The default value for the variable when no value is received. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The time when the variable was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the variable. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with this variable. |
| `VariableType` | variable_type | `string` | optional, computed, provider-chosen |  | The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types |

Supports update: yes

Discovery: supported
