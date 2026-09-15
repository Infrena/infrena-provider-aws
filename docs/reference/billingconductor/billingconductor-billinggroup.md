# aws.billingconductor.billinggroup

**CloudFormation type:** `AWS::BillingConductor::BillingGroup`

A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BillingConductor::BillingGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountGrouping` | account_grouping | `map` | required |  |  |
| `Arn` |  | `string` | computed |  | Billing Group ARN |
| `ComputationPreference` | computation_preference | `map` | required |  |  |
| `CreationTime` | creation_time | `integer` | computed |  | Creation timestamp in UNIX epoch time format |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedTime` | last_modified_time | `integer` | computed |  | Latest modified timestamp in UNIX epoch time format |
| `Name` |  | `string` | required |  |  |
| `PrimaryAccountId` | primary_account_id | `string` | optional, computed, provider-chosen, replaces on change |  | This account will act as a virtual payer account of the billing group |
| `Size` |  | `integer` | computed |  | Number of accounts in the billing group |
| `Status` |  | `string` | computed |  |  |
| `StatusReason` | status_reason | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
