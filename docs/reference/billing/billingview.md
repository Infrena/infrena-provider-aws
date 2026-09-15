# aws.billingview

**CloudFormation type:** `AWS::Billing::BillingView`

A billing view is a container of cost & usage metadata.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Billing::BillingView)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BillingViewType` | billing_view_type | `string` | computed |  |  |
| `CreatedAt` | created_at | `float` | computed |  | The time when the billing view was created. |
| `DataFilterExpression` | data_filter_expression | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `OwnerAccountId` | owner_account_id | `string` | computed |  |  |
| `SourceViews` | source_views | `list` | required |  | An array of strings that define the billing view's source. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs associated to the billing view being created. |
| `UpdatedAt` | updated_at | `float` | computed |  | The time when the billing view was last updated. |

Supports update: yes

Discovery: supported
