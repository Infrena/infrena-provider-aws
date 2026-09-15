# aws.docdb.eventsubscription

**CloudFormation type:** `AWS::DocDB::EventSubscription`

Resource Type definition for AWS::DocDB::EventSubscription

Region attribute: `region`

**Import ID:** `<region>/SubscriptionName` (AWS::DocDB::EventSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | A Boolean value; set to true to activate the subscription, set to false to create the subscription but not active it. |
| `EventCategories` | event_categories | `list` | optional, computed, provider-chosen |  | A list of event categories for a SourceType that you want to subscribe to. |
| `SnsTopicArn` | sns_topic_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the SNS topic created for event notification. Amazon SNS creates the ARN when you create a topic and subscribe to it. |
| `SourceIds` | source_ids | `list` | optional, computed, provider-chosen |  | The list of identifiers of the event sources for which events are returned |
| `SourceType` | source_type | `string` | optional, computed, provider-chosen |  | The type of source that is generating the events. |
| `SubscriptionName` | subscription_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the subscription. |

Supports update: yes

Discovery: supported
