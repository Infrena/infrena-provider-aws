# aws.eventrule

**CloudFormation type:** `AWS::Notifications::EventRule`

Resource Type definition for AWS::Notifications::EventRule

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Notifications::EventRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `EventPattern` | event_pattern | `string` | optional, computed, provider-chosen |  |  |
| `EventType` | event_type | `string` | required, replaces on change |  |  |
| `ManagedRules` | managed_rules | `list` | computed |  |  |
| `NotificationConfigurationArn` | notification_configuration_arn | `string` | required, replaces on change | aws.notificationconfiguration.Arn |  |
| `Regions` |  | `list` | required |  |  |
| `Source` |  | `string` | required, replaces on change |  |  |
| `StatusSummaryByRegion` | status_summary_by_region | `map` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
