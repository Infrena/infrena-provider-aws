# aws.productsubscription

**CloudFormation type:** `AWS::SecurityHub::ProductSubscription`

The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.

Region attribute: `region`

**Import ID:** `<region>/ProductSubscriptionArn` (AWS::SecurityHub::ProductSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ProductArn` | product_arn | `string` | required, replaces on change |  | The generic ARN of the product being subscribed to |
| `ProductSubscriptionArn` | product_subscription_arn | `string` | computed |  | The ARN of the product subscription for the account |

Supports update: no

Discovery: supported
