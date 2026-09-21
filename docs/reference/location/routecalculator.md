# aws.routecalculator

**CloudFormation type:** `AWS::Location::RouteCalculator`

Definition of AWS::Location::RouteCalculator Resource Type

Region attribute: `region`

**Import ID:** `<region>/CalculatorName` (AWS::Location::RouteCalculator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CalculatorArn` | calculator_arn | `string` | computed |  |  |
| `CalculatorName` | calculator_name | `string` | required, replaces on change |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `DataSource` | data_source | `string` | required, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `PricingPlan` | pricing_plan | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
