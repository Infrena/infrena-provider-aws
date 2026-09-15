# aws.pricingplan

**CloudFormation type:** `AWS::BillingConductor::PricingPlan`

Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BillingConductor::PricingPlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Pricing Plan ARN |
| `CreationTime` | creation_time | `integer` | computed |  | Creation timestamp in UNIX epoch time format |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedTime` | last_modified_time | `integer` | computed |  | Latest modified timestamp in UNIX epoch time format |
| `Name` |  | `string` | required |  |  |
| `PricingRuleArns` | pricing_rule_arns | `list` | optional, computed, provider-chosen | aws.pricingrule.Arn |  |
| `Size` |  | `integer` | computed |  | Number of associated pricing rules |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
