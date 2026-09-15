# aws.kendra.datasource

**CloudFormation type:** `AWS::Kendra::DataSource`

Kendra DataSource

Region attribute: `region`

**Import ID:** `<region>/Id|IndexId` (AWS::Kendra::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CustomDocumentEnrichmentConfiguration` | custom_document_enrichment_configuration | `map` | optional, computed, provider-chosen |  |  |
| `DataSourceConfiguration` | data_source_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of data source |
| `Id` |  | `string` | computed |  | ID of data source |
| `IndexId` | index_id | `string` | required, replaces on change | aws.kendra.index.Id | ID of Index |
| `LanguageCode` | language_code | `string` | optional, computed, provider-chosen |  | The code for a language. |
| `Name` |  | `string` | required |  | Name of data source |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | Role ARN |
| `Schedule` |  | `string` | optional, computed, provider-chosen |  | Schedule |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of tags |
| `Type` | type_value | `string` | required, replaces on change |  | Data source type |

Supports update: yes

Discovery: supported
