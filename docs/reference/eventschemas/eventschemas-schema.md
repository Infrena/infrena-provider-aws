# aws.eventschemas.schema

**CloudFormation type:** `AWS::EventSchemas::Schema`

Resource Type definition for AWS::EventSchemas::Schema

Region attribute: `region`

**Import ID:** `<region>/SchemaArn` (AWS::EventSchemas::Schema)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Content` |  | `string` | required |  | The source of the schema definition. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the schema. |
| `LastModified` | last_modified | `string` | computed |  | The last modified time of the schema. |
| `RegistryName` | registry_name | `string` | required, replaces on change |  | The name of the schema registry. |
| `SchemaArn` | schema_arn | `string` | computed |  | The ARN of the schema. |
| `SchemaName` | schema_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the schema. |
| `SchemaVersion` | schema_version | `string` | computed |  | The version number of the schema. |
| `Tags` |  | `map` | tags map |  | Tags associated with the resource. |
| `Type` | type_value | `string` | required |  | The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4. |
| `VersionCreatedDate` | version_created_date | `string` | computed |  | The date the schema version was created. |

Supports update: yes

Discovery: supported (parent resource required)
