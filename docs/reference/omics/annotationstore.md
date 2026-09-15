# aws.annotationstore

**CloudFormation type:** `AWS::Omics::AnnotationStore`

Definition of AWS::Omics::AnnotationStore Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Omics::AnnotationStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Reference` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SseConfig` | sse_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `StoreArn` | store_arn | `string` | computed |  |  |
| `StoreFormat` | store_format | `string` | required, replaces on change |  |  |
| `StoreOptions` | store_options | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StoreSizeBytes` | store_size_bytes | `float` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `UpdateTime` | update_time | `string` | computed |  |  |

Supports update: yes

Discovery: supported
