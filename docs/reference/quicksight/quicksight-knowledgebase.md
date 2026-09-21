# aws.quicksight.knowledgebase

**CloudFormation type:** `AWS::QuickSight::KnowledgeBase`

Definition of AWS::QuickSight::KnowledgeBase Resource Type

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|KnowledgeBaseId` (AWS::QuickSight::KnowledgeBase)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessControlConfiguration` | access_control_configuration | `map` | optional, computed, provider-chosen |  |  |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DataSourceArn` | data_source_arn | `string` | required, replaces on change | aws.quicksight.datasource.Arn |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DocumentCount` | document_count | `float` | computed |  |  |
| `IsEmailNotificationOptedForIngestionFailures` | is_email_notification_opted_for_ingestion_failures | `boolean` | optional, computed, provider-chosen |  |  |
| `KnowledgeBaseArn` | knowledge_base_arn | `string` | computed |  |  |
| `KnowledgeBaseConfiguration` | knowledge_base_configuration | `map` | required |  |  |
| `KnowledgeBaseId` | knowledge_base_id | `string` | required, replaces on change | aws.quicksight.knowledgebase.KnowledgeBaseId |  |
| `KnowledgeBaseSizeBytes` | knowledge_base_size_bytes | `float` | computed |  |  |
| `MediaExtractionConfiguration` | media_extraction_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `PrimaryOwnerArn` | primary_owner_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PrimaryOwnerUsername` | primary_owner_username | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
