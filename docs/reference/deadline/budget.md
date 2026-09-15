# aws.budget

**CloudFormation type:** `AWS::Deadline::Budget`

Creates a budget to set spending thresholds for your rendering activity.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Budget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | The budget actions to specify what happens when the budget runs out. |
| `ApproximateDollarLimit` | approximate_dollar_limit | `float` | required |  | The dollar limit based on consumed usage. |
| `Arn` |  | `string` | computed |  | The ARN of the budget. |
| `BudgetId` | budget_id | `string` | computed |  | The budget ID. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the budget. |
| `DisplayName` | display_name | `string` | required |  | The display name of the budget. |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId | The farm ID. |
| `Schedule` |  | `map` | required |  | The start and end time of the budget. |
| `Status` |  | `string` | computed |  | The status of the budget. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UsageTrackingResource` | usage_tracking_resource | `map` | required, replaces on change |  | The usage details of the allotted budget. |

Supports update: yes

Discovery: supported (parent resource required)
