# aws.variantstore

**CloudFormation type:** `AWS::Omics::VariantStore`

Definition of AWS::Omics::VariantStore Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Omics::VariantStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Reference` |  | `map` | required, replaces on change |  |  |
| `SseConfig` | sse_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `StoreArn` | store_arn | `string` | computed |  |  |
| `StoreSizeBytes` | store_size_bytes | `float` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `UpdateTime` | update_time | `string` | computed |  |  |

Supports update: yes

Discovery: supported
