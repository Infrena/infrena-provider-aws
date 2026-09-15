# aws.scheduledaudit

**CloudFormation type:** `AWS::IoT::ScheduledAudit`

Scheduled audits can be used to specify the checks you want to perform during an audit and how often the audit should be run.

Region attribute: `region`

**Import ID:** `<region>/ScheduledAuditName` (AWS::IoT::ScheduledAudit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DayOfMonth` | day_of_month | `string` | optional, computed, provider-chosen |  | The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY. |
| `DayOfWeek` | day_of_week | `string` | optional, computed, provider-chosen |  | The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY. |
| `Frequency` |  | `string` | required |  | How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY. |
| `ScheduledAuditArn` | scheduled_audit_arn | `string` | computed |  | The ARN (Amazon resource name) of the scheduled audit. |
| `ScheduledAuditName` | scheduled_audit_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name you want to give to the scheduled audit. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetCheckNames` | target_check_names | `list` | required |  | Which checks are performed during the scheduled audit. Checks must be enabled for your account. |

Supports update: yes

Discovery: supported
