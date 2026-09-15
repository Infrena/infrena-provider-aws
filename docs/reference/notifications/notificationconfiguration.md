# aws.notificationconfiguration

**CloudFormation type:** `AWS::Notifications::NotificationConfiguration`

Resource Type definition for AWS::Notifications::NotificationConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Notifications::NotificationConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AggregationDuration` | aggregation_duration | `string` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | required |  |  |
| `Name` |  | `string` | required |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | A list of tags that are attached to the role. |

Supports update: yes

Discovery: supported
