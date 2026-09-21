# aws.channelassociation

**CloudFormation type:** `AWS::Notifications::ChannelAssociation`

Resource Type definition for AWS::Notifications::ChannelAssociation

Region attribute: `region`

**Import ID:** `<region>/Arn|NotificationConfigurationArn` (AWS::Notifications::ChannelAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | required, replaces on change |  | ARN identifier of the channel. |
| `NotificationConfigurationArn` | notification_configuration_arn | `string` | required, replaces on change | aws.notificationconfiguration.Arn | ARN identifier of the NotificationConfiguration. |

Supports update: no

Discovery: supported (parent resource required)
