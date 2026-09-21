# aws.notification

**CloudFormation type:** `AWS::Connect::Notification`

Resource Type definition for AWS::Connect::Notification

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Connect::Notification)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) for the notification. |
| `Content` |  | `map` | required |  | The content of a notification |
| `CreatedAt` | created_at | `string` | computed |  | The time a notification was created |
| `ExpiresAt` | expires_at | `string` | optional, computed, provider-chosen, replaces on change |  | The time a notification will expire |
| `Id` |  | `string` | computed |  | The identifier of the notification. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Priority` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The priority of notification. In the Amazon Connect console, when you create a notification, you are prompted to assign one of the following priorities: High (HIGH) or LOW (LOW) |
| `Recipients` |  | `list` | optional, computed, provider-chosen, replaces on change |  | The recipients of the notification. |
| `Tags` |  | `map` | tags map |  | One or more tags. |

Supports update: yes

Discovery: supported (parent resource required)
