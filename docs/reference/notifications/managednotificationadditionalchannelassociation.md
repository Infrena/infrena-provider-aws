# aws.managednotificationadditionalchannelassociation

**CloudFormation type:** `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation`

Resource Type definition for AWS::Notifications::ManagedNotificationAdditionalChannelAssociation

Region attribute: `region`

**Import ID:** `<region>/ChannelArn|ManagedNotificationConfigurationArn` (AWS::Notifications::ManagedNotificationAdditionalChannelAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelArn` | channel_arn | `string` | required, replaces on change |  | ARN identifier of the channel. |
| `IsSensitiveEventsSubscribed` | is_sensitive_events_subscribed | `boolean` | optional, computed, provider-chosen |  | Whether the channel association is subscribed to sensitive events. Access to sensitive events is gated by the SubscribeSensitiveEvents virtual IAM action. |
| `ManagedNotificationConfigurationArn` | managed_notification_configuration_arn | `string` | required, replaces on change |  | ARN identifier of the Managed Notification. |

Supports update: yes

Discovery: supported (parent resource required)
