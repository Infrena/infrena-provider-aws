# aws.datacellsfilter

**CloudFormation type:** `AWS::LakeFormation::DataCellsFilter`

Resource Type definition for AWS::LakeFormation::DataCellsFilter

Region attribute: `region`

**Import ID:** `<region>/TableCatalogId|DatabaseName|TableName|Name` (AWS::LakeFormation::DataCellsFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ColumnNames` | column_names | `list` | optional, computed, provider-chosen, replaces on change |  | A list of column names. |
| `ColumnWildcard` | column_wildcard | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required. |
| `DatabaseName` | database_name | `string` | required, replaces on change |  | A string representing a resource's name. |
| `Name` |  | `string` | required, replaces on change |  | A string representing a resource's name. |
| `RowFilter` | row_filter | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required. |
| `TableCatalogId` | table_catalog_id | `string` | required, replaces on change |  | A string representing the Catalog Id. |
| `TableName` | table_name | `string` | required, replaces on change |  | A string representing a resource's name. |

Supports update: no

Discovery: supported
