# aws.notificationrule

**CloudFormation type:** `AWS::CodeStarNotifications::NotificationRule`

Resource Type definition for AWS::CodeStarNotifications::NotificationRule

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeStarNotifications::NotificationRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedBy` | created_by | `string` | optional, computed, provider-chosen |  |  |
| `DetailType` | detail_type | `string` | required |  |  |
| `EventTypeId` | event_type_id | `string` | optional, computed, provider-chosen, write-only |  |  |
| `EventTypeIds` | event_type_ids | `list` | required |  |  |
| `Name` |  | `string` | required |  |  |
| `Resource` |  | `string` | required, replaces on change |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `TargetAddress` | target_address | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Targets` |  | `list` | required |  |  |

Supports update: yes

Discovery: supported
