# aws.anomalysubscription

**CloudFormation type:** `AWS::CE::AnomalySubscription`

AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified

Region attribute: `region`

**Import ID:** `<region>/SubscriptionArn` (AWS::CE::AnomalySubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The accountId |
| `Frequency` |  | `string` | required |  | The frequency at which anomaly reports are sent over email. |
| `MonitorArnList` | monitor_arn_list | `list` | required |  | A list of cost anomaly monitors. |
| `ResourceTags` | resource_tags | `map` | replaces on change, tags map |  | Tags to assign to subscription. |
| `Subscribers` |  | `list` | required |  | A list of subscriber |
| `SubscriptionArn` | subscription_arn | `string` | computed |  | Subscription ARN |
| `SubscriptionName` | subscription_name | `string` | required |  | The name of the subscription. |
| `Threshold` |  | `float` | optional, computed, provider-chosen |  | The dollar value that triggers a notification if the threshold is exceeded. |
| `ThresholdExpression` | threshold_expression | `string` | optional, computed, provider-chosen |  | An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for. |

Supports update: yes

Discovery: supported
