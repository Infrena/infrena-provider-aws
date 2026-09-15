# aws.bedrock.datasource

**CloudFormation type:** `AWS::Bedrock::DataSource`

Definition of AWS::Bedrock::DataSource Resource Type

Region attribute: `region`

**Import ID:** `<region>/KnowledgeBaseId|DataSourceId` (AWS::Bedrock::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time at which the data source was created. |
| `DataDeletionPolicy` | data_deletion_policy | `string` | optional, computed, provider-chosen |  | The deletion policy for the data source. |
| `DataSourceConfiguration` | data_source_configuration | `map` | required |  | Specifies a raw data source location to ingest. |
| `DataSourceId` | data_source_id | `string` | computed |  | Identifier for a resource. |
| `DataSourceStatus` | data_source_status | `string` | computed |  | The status of a data source. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Resource. |
| `FailureReasons` | failure_reasons | `list` | computed |  | The details of the failure reasons related to the data source. |
| `KnowledgeBaseId` | knowledge_base_id | `string` | required, replaces on change | aws.bedrock.knowledgebase.KnowledgeBaseId | The unique identifier of the knowledge base to which to add the data source. |
| `Name` |  | `string` | required |  | The name of the data source. |
| `ServerSideEncryptionConfiguration` | server_side_encryption_configuration | `map` | optional, computed, provider-chosen |  | Contains details about the server-side encryption for the data source. |
| `UpdatedAt` | updated_at | `string` | computed |  | The time at which the knowledge base was last updated. |
| `VectorIngestionConfiguration` | vector_ingestion_configuration | `map` | optional, computed, provider-chosen |  | Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. |

Supports update: yes

Discovery: supported (parent resource required)
