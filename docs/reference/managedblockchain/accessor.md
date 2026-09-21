# aws.accessor

**CloudFormation type:** `AWS::ManagedBlockchain::Accessor`

Definition of AWS::ManagedBlockchain::com.amazonaws.taiga.webservice.api#Accessor Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ManagedBlockchain::Accessor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessorType` | accessor_type | `string` | required, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `BillingToken` | billing_token | `string` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen, write-only |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
