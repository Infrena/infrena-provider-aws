# aws.referencestore

**CloudFormation type:** `AWS::Omics::ReferenceStore`

Definition of AWS::Omics::ReferenceStore Resource Type

Region attribute: `region`

**Import ID:** `<region>/ReferenceStoreId` (AWS::Omics::ReferenceStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The store's ARN. |
| `CreationTime` | creation_time | `string` | computed |  | When the store was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description for the store. |
| `Name` |  | `string` | required, replaces on change |  | A name for the store. |
| `ReferenceStoreId` | reference_store_id | `string` | computed |  |  |
| `SseConfig` | sse_config | `map` | optional, computed, provider-chosen, replaces on change |  | Server-side encryption (SSE) settings for a store. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported
