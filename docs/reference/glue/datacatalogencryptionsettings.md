# aws.datacatalogencryptionsettings

**CloudFormation type:** `AWS::Glue::DataCatalogEncryptionSettings`

Resource Type definition for AWS::Glue::DataCatalogEncryptionSettings

Region attribute: `region`

**Import ID:** `<region>/CatalogId` (AWS::Glue::DataCatalogEncryptionSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CatalogId` | catalog_id | `string` | required, replaces on change | aws.catalog.CatalogId | The ID of the Data Catalog in which the settings are created. |
| `DataCatalogEncryptionSettings` | data_catalog_encryption_settings | `map` | required |  | Contains configuration information for maintaining Data Catalog security. |

Supports update: yes

Discovery: not supported
