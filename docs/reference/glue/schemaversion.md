# aws.schemaversion

**CloudFormation type:** `AWS::Glue::SchemaVersion`

This resource represents an individual schema version of a schema defined in Glue Schema Registry.

Region attribute: `region`

**Import ID:** `<region>/VersionId` (AWS::Glue::SchemaVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Schema` |  | `map` | required, replaces on change |  | Identifier for the schema where the schema version will be created. |
| `SchemaDefinition` | schema_definition | `string` | required, replaces on change |  | Complete definition of the schema in plain-text. |
| `VersionId` | version_id | `string` | computed |  | Represents the version ID associated with the schema version. |

Supports update: no

Discovery: supported (parent resource required)
