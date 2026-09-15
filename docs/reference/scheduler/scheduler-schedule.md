# aws.scheduler.schedule

**CloudFormation type:** `AWS::Scheduler::Schedule`

Definition of AWS::Scheduler::Schedule Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Scheduler::Schedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the schedule. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the schedule. |
| `EndDate` | end_date | `string` | optional, computed, provider-chosen |  | The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify. |
| `FlexibleTimeWindow` | flexible_time_window | `map` | required |  | Flexible time window allows configuration of a window within which a schedule can be invoked |
| `GroupName` | group_name | `string` | optional, computed, provider-chosen |  | The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The ARN for a KMS Key that will be used to encrypt customer data. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ScheduleExpression` | schedule_expression | `string` | required |  | The scheduling expression. |
| `ScheduleExpressionTimezone` | schedule_expression_timezone | `string` | optional, computed, provider-chosen |  | The timezone in which the scheduling expression is evaluated. |
| `StartDate` | start_date | `string` | optional, computed, provider-chosen |  | The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify. |
| `State` |  | `string` | optional, computed, provider-chosen |  | Specifies whether the schedule is enabled or disabled. |
| `Target` |  | `map` | required |  | The schedule target. |

Supports update: yes

Discovery: supported
