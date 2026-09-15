# aws.qbusiness.datasource

**CloudFormation type:** `AWS::QBusiness::DataSource`

Definition of AWS::QBusiness::DataSource Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|DataSourceId|IndexId` (AWS::QBusiness::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `Configuration` |  | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DataSourceArn` | data_source_arn | `string` | computed |  |  |
| `DataSourceId` | data_source_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `DocumentEnrichmentConfiguration` | document_enrichment_configuration | `map` | optional, computed, provider-chosen |  |  |
| `IndexId` | index_id | `string` | required, replaces on change | aws.qbusiness.index.IndexId |  |
| `MediaExtractionConfiguration` | media_extraction_configuration | `map` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `SyncSchedule` | sync_schedule | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |
| `VpcConfiguration` | vpc_configuration | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
