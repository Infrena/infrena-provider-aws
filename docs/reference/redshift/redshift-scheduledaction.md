# aws.redshift.scheduledaction

**CloudFormation type:** `AWS::Redshift::ScheduledAction`

The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.

Region attribute: `region`

**Import ID:** `<region>/ScheduledActionName` (AWS::Redshift::ScheduledAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Enable` |  | `boolean` | optional, computed, provider-chosen |  | If true, the schedule is enabled. If false, the scheduled action does not trigger. |
| `EndTime` | end_time | `string` | optional, computed, provider-chosen |  | The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger. |
| `IamRole` | iam_role | `string` | optional, computed, provider-chosen |  | The IAM role to assume to run the target action. |
| `NextInvocations` | next_invocations | `list` | computed |  | List of times when the scheduled action will run. |
| `Schedule` |  | `string` | optional, computed, provider-chosen |  | The schedule in `at( )` or `cron( )` format. |
| `ScheduledActionDescription` | scheduled_action_description | `string` | optional, computed, provider-chosen |  | The description of the scheduled action. |
| `ScheduledActionName` | scheduled_action_name | `string` | required, replaces on change |  | The name of the scheduled action. The name must be unique within an account. |
| `StartTime` | start_time | `string` | optional, computed, provider-chosen |  | The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger. |
| `State` |  | `string` | computed |  | The state of the scheduled action. |
| `TargetAction` | target_action | `map` | optional, computed, provider-chosen |  | A JSON format string of the Amazon Redshift API operation with input parameters. |

Supports update: yes

Discovery: supported
