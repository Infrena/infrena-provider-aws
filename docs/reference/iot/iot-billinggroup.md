# aws.iot.billinggroup

**CloudFormation type:** `AWS::IoT::BillingGroup`

Resource Type definition for AWS::IoT::BillingGroup

Region attribute: `region`

**Import ID:** `<region>/BillingGroupName` (AWS::IoT::BillingGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BillingGroupName` | billing_group_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `BillingGroupProperties` | billing_group_properties | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
