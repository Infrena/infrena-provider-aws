# aws.glue.database

**CloudFormation type:** `AWS::Glue::Database`

Resource Type definition for AWS::Glue::Database

Region attribute: `region`

**Import ID:** `<region>/DatabaseName` (AWS::Glue::Database)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CatalogId` | catalog_id | `string` | required | aws.catalog.CatalogId | The AWS account ID for the account in which to create the catalog object. |
| `DatabaseInput` | database_input | `map` | required |  | The structure used to create or update a database. |
| `DatabaseName` | database_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the database. For hive compatibility, this is folded to lowercase when it is store. |

Supports update: yes

Discovery: supported
