# aws.restoretestingplan

**CloudFormation type:** `AWS::Backup::RestoreTestingPlan`

Definition of AWS::Backup::RestoreTestingPlan Resource Type

Region attribute: `region`

**Import ID:** `<region>/RestoreTestingPlanName` (AWS::Backup::RestoreTestingPlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RecoveryPointSelection` | recovery_point_selection | `map` | required |  |  |
| `RestoreTestingPlanArn` | restore_testing_plan_arn | `string` | computed |  |  |
| `RestoreTestingPlanName` | restore_testing_plan_name | `string` | required, replaces on change |  |  |
| `ScheduleExpression` | schedule_expression | `string` | required |  |  |
| `ScheduleExpressionTimezone` | schedule_expression_timezone | `string` | optional, computed, provider-chosen |  |  |
| `StartWindowHours` | start_window_hours | `integer` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
