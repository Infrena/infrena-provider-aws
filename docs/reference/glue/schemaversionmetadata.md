# aws.schemaversionmetadata

**CloudFormation type:** `AWS::Glue::SchemaVersionMetadata`

This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.

Region attribute: `region`

**Import ID:** `<region>/SchemaVersionId|Key|Value` (AWS::Glue::SchemaVersionMetadata)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Key` |  | `string` | required, replaces on change |  | Metadata key |
| `SchemaVersionId` | schema_version_id | `string` | required, replaces on change | aws.schemaversion.VersionId | Represents the version ID associated with the schema version. |
| `Value` |  | `string` | required, replaces on change |  | Metadata value |

Supports update: no

Discovery: supported (parent resource required)
