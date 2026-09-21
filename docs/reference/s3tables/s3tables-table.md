# aws.s3tables.table

**CloudFormation type:** `AWS::S3Tables::Table`

Resource Type definition for AWS::S3Tables::Table

Region attribute: `region`

**Import ID:** `<region>/TableARN` (AWS::S3Tables::Table)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Compaction` |  | `map` | optional, computed, provider-chosen |  | Settings governing the Compaction maintenance action. Contains details about the compaction settings for an Iceberg table. |
| `IcebergMetadata` | iceberg_metadata | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Contains details about the metadata for an Iceberg table. Specify either IcebergSchema (for simple flat schemas with primitive types only) or IcebergSchemaV2 (for schemas with nested types like struct, list, map), but not both. |
| `Namespace` |  | `string` | required |  | The namespace that the table belongs to. |
| `OpenTableFormat` | open_table_format | `string` | required, replaces on change |  | Format of the table. |
| `SnapshotManagement` | snapshot_management | `map` | optional, computed, provider-chosen |  | Contains details about the snapshot management settings for an Iceberg table. A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours. |
| `StorageClassConfiguration` | storage_class_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies storage class settings for the table |
| `TableARN` | table_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified table. |
| `TableBucketARN` | table_bucket_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the specified table bucket. |
| `TableName` | table_name | `string` | required |  | The name for the table. |
| `Tags` |  | `map` | tags map |  | User tags (key-value pairs) to associate with the table. |
| `VersionToken` | version_token | `string` | computed |  | The version token of the table |
| `WarehouseLocation` | warehouse_location | `string` | computed |  | The warehouse location of the table. |
| `WithoutMetadata` | without_metadata | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Indicates that you don't want to specify a schema for the table. This property is mutually exclusive to 'IcebergMetadata', and its only possible value is 'Yes'. |

Supports update: yes

Discovery: supported (parent resource required)
