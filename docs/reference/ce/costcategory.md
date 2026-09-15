# aws.costcategory

**CloudFormation type:** `AWS::CE::CostCategory`

Resource Type definition for AWS::CE::CostCategory. Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CE::CostCategory)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Cost category ARN |
| `DefaultValue` | default_value | `string` | optional, computed, provider-chosen |  | The default value for the cost category |
| `EffectiveStart` | effective_start | `string` | computed |  | ISO 8601 date time with offset format |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RuleVersion` | rule_version | `string` | required |  |  |
| `Rules` |  | `string` | required |  | JSON array format of Expression in Billing and Cost Management API |
| `SplitChargeRules` | split_charge_rules | `string` | optional, computed, provider-chosen |  | Json array format of CostCategorySplitChargeRule in Billing and Cost Management API |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the cost category. |

Supports update: yes

Discovery: supported
