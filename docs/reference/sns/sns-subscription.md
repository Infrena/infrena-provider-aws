# aws.sns.subscription

**CloudFormation type:** `AWS::SNS::Subscription`

Resource Type definition for AWS::SNS::Subscription

Region attribute: `aws_region`

**Import ID:** `<region>/Arn` (AWS::SNS::Subscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn of the subscription |
| `DeliveryPolicy` | delivery_policy | `string` | optional, computed, provider-chosen |  | The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic. |
| `Endpoint` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The subscription's endpoint. The endpoint value depends on the protocol that you specify. |
| `FilterPolicy` | filter_policy | `string` | optional, computed, provider-chosen |  | The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages. |
| `FilterPolicyScope` | filter_policy_scope | `string` | optional, computed, provider-chosen |  | This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody. |
| `Protocol` |  | `string` | required, replaces on change |  | The subscription's protocol. |
| `RawMessageDelivery` | raw_message_delivery | `boolean` | optional, computed, provider-chosen |  | When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. |
| `RedrivePolicy` | redrive_policy | `string` | optional, computed, provider-chosen |  | When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing. |
| `Region` |  | `string` | optional, computed, provider-chosen, write-only |  | For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default. |
| `ReplayPolicy` | replay_policy | `string` | optional, computed, provider-chosen |  | Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes. |
| `SubscriptionRoleArn` | subscription_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | This property applies only to Amazon Data Firehose delivery stream subscriptions. |
| `TopicArn` | topic_arn | `string` | required, replaces on change | aws.sns.topic.TopicArn | The ARN of the topic to subscribe to. |

Supports update: yes

Discovery: supported
