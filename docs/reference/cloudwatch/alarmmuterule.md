# aws.alarmmuterule

**CloudFormation type:** `AWS::CloudWatch::AlarmMuteRule`

Resource Type definition for AWS::CloudWatch::AlarmMuteRule that allows defining a rule and targeting alarms to mute their actions during the specified window.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudWatch::AlarmMuteRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the AlarmMuteRule |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the AlarmMuteRule |
| `ExpireDate` | expire_date | `string` | optional, computed, provider-chosen |  | The date, with the same timezone offset as "ScheduleTimezone" after which the alarm mute rule will be expired. |
| `LastUpdatedTimestamp` | last_updated_timestamp | `string` | computed |  | The last update timestamp of the alarm mute schedule |
| `MuteTargets` | mute_targets | `map` | optional, computed, provider-chosen |  | Targets to be muted |
| `MuteType` | mute_type | `string` | computed |  | The mute type of the alarm mute |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the AlarmMuteRule |
| `Rule` |  | `map` | required |  | The rule for the mute |
| `StartDate` | start_date | `string` | optional, computed, provider-chosen |  | The date, with the same timezone offset as "ScheduleTimezone", after which the alarm mute rule will become active. |
| `Status` |  | `string` | computed |  | The current status of the AlarmMuteRule |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
