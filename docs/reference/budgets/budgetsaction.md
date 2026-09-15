# aws.budgetsaction

**CloudFormation type:** `AWS::Budgets::BudgetsAction`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/ActionId|BudgetName` (AWS::Budgets::BudgetsAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionId` | action_id | `string` | computed |  |  |
| `ActionThreshold` | action_threshold | `map` | required |  |  |
| `ActionType` | action_type | `string` | required, replaces on change |  |  |
| `ApprovalModel` | approval_model | `string` | optional, computed, provider-chosen |  |  |
| `BudgetName` | budget_name | `string` | required, replaces on change |  |  |
| `Definition` |  | `map` | required |  |  |
| `ExecutionRoleArn` | execution_role_arn | `string` | required | aws.role.Arn |  |
| `NotificationType` | notification_type | `string` | required |  |  |
| `ResourceTags` | resource_tags | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Subscribers` |  | `list` | required |  |  |

Supports update: yes

Discovery: supported
