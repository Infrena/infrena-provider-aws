# aws.personalize.schema

**CloudFormation type:** `AWS::Personalize::Schema`

Resource schema for AWS::Personalize::Schema.

Region attribute: `region`

**Import ID:** `<region>/SchemaArn` (AWS::Personalize::Schema)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Domain` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The domain of a Domain dataset group. |
| `Name` |  | `string` | required, replaces on change |  | Name for the schema. |
| `Schema` |  | `string` | required, replaces on change |  | A schema in Avro JSON format. |
| `SchemaArn` | schema_arn | `string` | computed |  | Arn for the schema. |
| `Tags` |  | `map` | replaces on change, tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: no

Discovery: supported
