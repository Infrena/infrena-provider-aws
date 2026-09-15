# aws.tableoptimizer

**CloudFormation type:** `AWS::Glue::TableOptimizer`

Resource Type definition for AWS::Glue::TableOptimizer

Region attribute: `region`

**Import ID:** `<region>/TableName|DatabaseName|Type|CatalogId` (AWS::Glue::TableOptimizer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CatalogId` | catalog_id | `string` | required, replaces on change | aws.catalog.CatalogId | The catalog ID of the table |
| `DatabaseName` | database_name | `string` | required, replaces on change |  | The name of the database. For Hive compatibility, this is folded to lowercase when it is stored. |
| `TableName` | table_name | `string` | required, replaces on change |  | The table name. For Hive compatibility, this must be entirely lowercase. |
| `TableOptimizerConfiguration` | table_optimizer_configuration | `map` | required |  | Specifies configuration details of a table optimizer. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of table optimizer. |

Supports update: yes

Discovery: supported (parent resource required)
