# aws.timestream.table

**CloudFormation type:** `AWS::Timestream::Table`

The AWS::Timestream::Table resource creates a Timestream Table.

Region attribute: `region`

**Import ID:** `<region>/DatabaseName|TableName` (AWS::Timestream::Table)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DatabaseName` | database_name | `string` | required, replaces on change |  | The name for the database which the table to be created belongs to. |
| `MagneticStoreWriteProperties` | magnetic_store_write_properties | `map` | optional, computed, provider-chosen |  | The properties that determine whether magnetic store writes are enabled. |
| `Name` |  | `string` | computed |  | The table name exposed as a read-only attribute. |
| `RetentionProperties` | retention_properties | `map` | optional, computed, provider-chosen |  | The retention duration of the memory store and the magnetic store. |
| `Schema` |  | `map` | optional, computed, provider-chosen |  | A Schema specifies the expected data model of the table. |
| `TableName` | table_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
