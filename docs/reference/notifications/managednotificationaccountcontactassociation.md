# aws.managednotificationaccountcontactassociation

**CloudFormation type:** `AWS::Notifications::ManagedNotificationAccountContactAssociation`

Resource Type definition for ManagedNotificationAccountContactAssociation

Region attribute: `region`

**Import ID:** `<region>/ManagedNotificationConfigurationArn|ContactIdentifier` (AWS::Notifications::ManagedNotificationAccountContactAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactIdentifier` | contact_identifier | `string` | required, replaces on change |  | This unique identifier for Contact |
| `IsSensitiveEventsSubscribed` | is_sensitive_events_subscribed | `boolean` | optional, computed, provider-chosen |  | Whether the account contact association is subscribed to sensitive events. Access to sensitive events is gated by the SubscribeSensitiveEvents virtual IAM action. |
| `ManagedNotificationConfigurationArn` | managed_notification_configuration_arn | `string` | required, replaces on change |  | The managed notification configuration ARN, against which the account contact association will be created |

Supports update: yes

Discovery: supported (parent resource required)
