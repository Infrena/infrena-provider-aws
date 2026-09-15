# aws.customlineitem

**CloudFormation type:** `AWS::BillingConductor::CustomLineItem`

A custom line item is an one time charge that is applied to a specific billing group's bill.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BillingConductor::CustomLineItem)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | optional, computed, provider-chosen, replaces on change |  | The account which this custom line item will be charged to |
| `Arn` |  | `string` | computed |  | ARN |
| `AssociationSize` | association_size | `integer` | computed |  | Number of source values associated to this custom line item |
| `BillingGroupArn` | billing_group_arn | `string` | required, replaces on change | aws.billingconductor.billinggroup.Arn | Billing Group ARN |
| `BillingPeriodRange` | billing_period_range | `map` | optional, computed, provider-chosen |  |  |
| `ComputationRule` | computation_rule | `string` | optional, computed, provider-chosen, replaces on change |  | The display settings of the Custom Line Item. |
| `CreationTime` | creation_time | `integer` | computed |  | Creation timestamp in UNIX epoch time format |
| `CurrencyCode` | currency_code | `string` | computed |  |  |
| `CustomLineItemChargeDetails` | custom_line_item_charge_details | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedTime` | last_modified_time | `integer` | computed |  | Latest modified timestamp in UNIX epoch time format |
| `Name` |  | `string` | required |  |  |
| `PresentationDetails` | presentation_details | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ProductCode` | product_code | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
