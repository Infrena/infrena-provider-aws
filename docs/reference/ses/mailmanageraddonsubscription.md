# aws.mailmanageraddonsubscription

**CloudFormation type:** `AWS::SES::MailManagerAddonSubscription`

Definition of AWS::SES::MailManagerAddonSubscription Resource Type

Region attribute: `region`

**Import ID:** `<region>/AddonSubscriptionId` (AWS::SES::MailManagerAddonSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddonName` | addon_name | `string` | required, replaces on change |  |  |
| `AddonSubscriptionArn` | addon_subscription_arn | `string` | computed |  |  |
| `AddonSubscriptionId` | addon_subscription_id | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
