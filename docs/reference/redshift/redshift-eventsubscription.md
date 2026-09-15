# aws.redshift.eventsubscription

**CloudFormation type:** `AWS::Redshift::EventSubscription`

The `AWS::Redshift::EventSubscription` resource creates an Amazon Redshift Event Subscription.

Region attribute: `region`

**Import ID:** `<region>/SubscriptionName` (AWS::Redshift::EventSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CustSubscriptionId` | cust_subscription_id | `string` | computed |  | The name of the Amazon Redshift event notification subscription. |
| `CustomerAwsId` | customer_aws_id | `string` | computed |  | The AWS account associated with the Amazon Redshift event notification subscription. |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it. |
| `EventCategories` | event_categories | `list` | optional, computed, provider-chosen |  | Specifies the Amazon Redshift event categories to be published by the event notification subscription. |
| `EventCategoriesList` | event_categories_list | `list` | computed |  | The list of Amazon Redshift event categories specified in the event notification subscription. |
| `Severity` |  | `string` | optional, computed, provider-chosen |  | Specifies the Amazon Redshift event severity to be published by the event notification subscription. |
| `SnsTopicArn` | sns_topic_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications. |
| `SourceIds` | source_ids | `list` | optional, computed, provider-chosen |  | A list of one or more identifiers of Amazon Redshift source objects. |
| `SourceIdsList` | source_ids_list | `list` | computed |  | A list of the sources that publish events to the Amazon Redshift event notification subscription. |
| `SourceType` | source_type | `string` | optional, computed, provider-chosen |  | The type of source that will be generating the events. |
| `Status` |  | `string` | computed |  | The status of the Amazon Redshift event notification subscription. |
| `SubscriptionCreationTime` | subscription_creation_time | `string` | computed |  | The date and time the Amazon Redshift event notification subscription was created. |
| `SubscriptionName` | subscription_name | `string` | required, replaces on change |  | The name of the Amazon Redshift event notification subscription |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
