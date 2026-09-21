# aws.compositealarm

**CloudFormation type:** `AWS::CloudWatch::CompositeAlarm`

The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression

Region attribute: `region`

**Import ID:** `<region>/AlarmName` (AWS::CloudWatch::CompositeAlarm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionsEnabled` | actions_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE. |
| `ActionsSuppressor` | actions_suppressor | `string` | optional, computed, provider-chosen |  | Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. |
| `ActionsSuppressorExtensionPeriod` | actions_suppressor_extension_period | `integer` | optional, computed, provider-chosen |  | Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds. |
| `ActionsSuppressorWaitPeriod` | actions_suppressor_wait_period | `integer` | optional, computed, provider-chosen |  | Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds. |
| `AlarmActions` | alarm_actions | `list` | optional, computed, provider-chosen |  | The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). |
| `AlarmDescription` | alarm_description | `string` | optional, computed, provider-chosen |  | The description of the alarm |
| `AlarmName` | alarm_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Composite Alarm |
| `AlarmRule` | alarm_rule | `string` | required |  | Expression which aggregates the state of other Alarms (Metric or Composite Alarms) |
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the alarm |
| `InsufficientDataActions` | insufficient_data_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN). |
| `OKActions` | ok_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN). |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to associate with the composite alarm. You can associate as many as 50 tags with an alarm. |

Supports update: yes

Discovery: supported
