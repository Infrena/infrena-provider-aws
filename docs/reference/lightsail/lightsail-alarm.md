# aws.lightsail.alarm

**CloudFormation type:** `AWS::Lightsail::Alarm`

Resource Type definition for AWS::Lightsail::Alarm

Region attribute: `region`

**Import ID:** `<region>/AlarmName` (AWS::Lightsail::Alarm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlarmArn` | alarm_arn | `string` | computed |  |  |
| `AlarmName` | alarm_name | `string` | required, replaces on change |  | The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm. |
| `ComparisonOperator` | comparison_operator | `string` | required |  | The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand. |
| `ContactProtocols` | contact_protocols | `list` | optional, computed, provider-chosen |  | The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both. |
| `DatapointsToAlarm` | datapoints_to_alarm | `integer` | optional, computed, provider-chosen |  | The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an "M out of N" alarm, this value (datapointsToAlarm) is the M. |
| `EvaluationPeriods` | evaluation_periods | `integer` | required |  | The number of most recent periods over which data is compared to the specified threshold. If you are setting an "M out of N" alarm, this value (evaluationPeriods) is the N. |
| `MetricName` | metric_name | `string` | required, replaces on change |  | The name of the metric to associate with the alarm. |
| `MonitoredResourceName` | monitored_resource_name | `string` | required, replaces on change |  | The name of the Lightsail resource that the alarm monitors. |
| `NotificationEnabled` | notification_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter. |
| `NotificationTriggers` | notification_triggers | `list` | optional, computed, provider-chosen |  | The alarm states that trigger a notification. |
| `State` |  | `string` | computed |  | The current state of the alarm. |
| `Threshold` |  | `float` | required |  | The value against which the specified statistic is compared. |
| `TreatMissingData` | treat_missing_data | `string` | optional, computed, provider-chosen |  | Sets how this alarm will handle missing data points. |

Supports update: yes

Discovery: supported
