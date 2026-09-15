# aws.pricingrule

**CloudFormation type:** `AWS::BillingConductor::PricingRule`

A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BillingConductor::PricingRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Pricing rule ARN |
| `AssociatedPricingPlanCount` | associated_pricing_plan_count | `integer` | computed |  | The number of pricing plans associated with pricing rule |
| `BillingEntity` | billing_entity | `string` | optional, computed, provider-chosen, replaces on change |  | The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL. |
| `CreationTime` | creation_time | `integer` | computed |  | Creation timestamp in UNIX epoch time format |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Pricing rule description |
| `LastModifiedTime` | last_modified_time | `integer` | computed |  | Latest modified timestamp in UNIX epoch time format |
| `ModifierPercentage` | modifier_percentage | `float` | optional, computed, provider-chosen |  | Pricing rule modifier percentage |
| `Name` |  | `string` | required |  | Pricing rule name |
| `Operation` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The Operation which a SKU pricing rule is modifying |
| `Scope` |  | `string` | required, replaces on change |  | A term used to categorize the granularity of a Pricing Rule. |
| `Service` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The service which a pricing rule is applied on |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Tiering` |  | `map` | optional, computed, provider-chosen |  | The set of tiering configurations for the pricing rule. |
| `Type` | type_value | `string` | required |  | One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule. |
| `UsageType` | usage_type | `string` | optional, computed, provider-chosen, replaces on change |  | The UsageType which a SKU pricing rule is modifying |

Supports update: yes

Discovery: supported
