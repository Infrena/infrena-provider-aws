# aws.wisdom.knowledgebase

**CloudFormation type:** `AWS::Wisdom::KnowledgeBase`

Definition of AWS::Wisdom::KnowledgeBase Resource Type

Region attribute: `region`

**Import ID:** `<region>/KnowledgeBaseId` (AWS::Wisdom::KnowledgeBase)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `KnowledgeBaseArn` | knowledge_base_arn | `string` | computed |  |  |
| `KnowledgeBaseId` | knowledge_base_id | `string` | computed |  |  |
| `KnowledgeBaseType` | knowledge_base_type | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RenderingConfiguration` | rendering_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ServerSideEncryptionConfiguration` | server_side_encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SourceConfiguration` | source_configuration | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | replaces on change, tags map |  |  |
| `VectorIngestionConfiguration` | vector_ingestion_configuration | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
