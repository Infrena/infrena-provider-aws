# aws.mailmanageraddoninstance

**CloudFormation type:** `AWS::SES::MailManagerAddonInstance`

Definition of AWS::SES::MailManagerAddonInstance Resource Type

Region attribute: `region`

**Import ID:** `<region>/AddonInstanceId` (AWS::SES::MailManagerAddonInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddonInstanceArn` | addon_instance_arn | `string` | computed |  |  |
| `AddonInstanceId` | addon_instance_id | `string` | computed |  |  |
| `AddonName` | addon_name | `string` | computed |  |  |
| `AddonSubscriptionId` | addon_subscription_id | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
