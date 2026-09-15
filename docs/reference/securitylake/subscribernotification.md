# aws.subscribernotification

**CloudFormation type:** `AWS::SecurityLake::SubscriberNotification`

Resource Type definition for AWS::SecurityLake::SubscriberNotification

Region attribute: `region`

**Import ID:** `<region>/SubscriberArn` (AWS::SecurityLake::SubscriberNotification)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `NotificationConfiguration` | notification_configuration | `map` | required |  |  |
| `SubscriberArn` | subscriber_arn | `string` | required, replaces on change | aws.subscriber.SubscriberArn | The ARN for the subscriber |
| `SubscriberEndpoint` | subscriber_endpoint | `string` | computed |  | The endpoint the subscriber should listen to for notifications |

Supports update: yes

Discovery: supported
