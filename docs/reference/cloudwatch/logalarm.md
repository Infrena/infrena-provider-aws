# aws.logalarm

**CloudFormation type:** `AWS::CloudWatch::LogAlarm`

Resource Type definition for AWS::CloudWatch::LogAlarm. A LogAlarm evaluates scheduled query results from CloudWatch Logs and triggers actions when thresholds are breached.

Region attribute: `region`

**Import ID:** `<region>/AlarmName` (AWS::CloudWatch::LogAlarm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionLogLineCount` | action_log_line_count | `integer` | optional, computed, provider-chosen |  | The number of log lines to include in alarm notifications. Valid values are 0 to 50. |
| `ActionLogLineRoleArn` | action_log_line_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications. Required when ActionLogLineCount is greater than 0. |
| `ActionsEnabled` | actions_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE. |
| `AlarmActions` | alarm_actions | `list` | optional, computed, provider-chosen |  | The list of actions to execute when this alarm transitions into an ALARM state from any other state. |
| `AlarmDescription` | alarm_description | `string` | optional, computed, provider-chosen |  | The description of the log alarm. |
| `AlarmName` | alarm_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the log alarm. |
| `Arn` |  | `string` | computed |  | The ARN of the log alarm. |
| `ComparisonOperator` | comparison_operator | `string` | required |  | The arithmetic operation to use when comparing the specified threshold and the query results. Valid values are GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, and LessThanOrEqualToThreshold. |
| `InsufficientDataActions` | insufficient_data_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. |
| `OKActions` | ok_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the OK state from any other state. |
| `QueryResultsToAlarm` | query_results_to_alarm | `integer` | required |  | The number of query results that must be breaching to trigger the alarm. |
| `QueryResultsToEvaluate` | query_results_to_evaluate | `integer` | required |  | The number of query results over which data is compared to the specified threshold. |
| `ScheduledQueryConfiguration` | scheduled_query_configuration | `map` | required |  | The scheduled query configuration for the log alarm. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to associate with the log alarm. |
| `Threshold` |  | `float` | required |  | The value to compare against the results of the scheduled query evaluation. |
| `TreatMissingData` | treat_missing_data | `string` | optional, computed, provider-chosen |  | Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing. |
| `WarmUpConfiguration` | warm_up_configuration | `map` | optional, computed, provider-chosen |  | The warm-up configuration for an alarm. A warm-up period delays alarm evaluation after you create or update the alarm. This reduces alarm noise from missing data while a new resource or service starts up. During the warm-up period, the alarm stays in INSUFFICIENT_DATA and doesn't perform alarm actions. |

Supports update: yes

Discovery: supported
