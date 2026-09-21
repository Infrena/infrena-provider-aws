# aws.glue.schema

**CloudFormation type:** `AWS::Glue::Schema`

This resource represents a schema of Glue Schema Registry.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Glue::Schema)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name for the Schema. |
| `CheckpointVersion` | checkpoint_version | `map` | optional, computed, provider-chosen |  | Specify checkpoint version for update. This is only required to update the Compatibility. |
| `Compatibility` |  | `string` | required |  | Compatibility setting for the schema. |
| `DataFormat` | data_format | `string` | required, replaces on change |  | Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF' |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the schema. If description is not provided, there will not be any default value for this. |
| `InitialSchemaVersionId` | initial_schema_version_id | `string` | computed |  | Represents the version ID associated with the initial schema version. |
| `Name` |  | `string` | required, replaces on change |  | Name of the schema. |
| `Registry` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Identifier for the registry which the schema is part of. |
| `SchemaDefinition` | schema_definition | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Definition for the initial schema version in plain-text. |
| `Tags` |  | `map` | tags map |  | List of tags to tag the schema |

Supports update: yes

Discovery: supported
