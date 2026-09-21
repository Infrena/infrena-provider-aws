# aws.schemamapping

**CloudFormation type:** `AWS::EntityResolution::SchemaMapping`

SchemaMapping defined in AWS Entity Resolution service

Region attribute: `region`

**Import ID:** `<region>/SchemaName` (AWS::EntityResolution::SchemaMapping)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this SchemaMapping got created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the SchemaMapping |
| `HasWorkflows` | has_workflows | `boolean` | computed |  | The boolean value that indicates whether or not a SchemaMapping has MatchingWorkflows that are associated with |
| `MappedInputFields` | mapped_input_fields | `list` | required |  | The SchemaMapping attributes input |
| `SchemaArn` | schema_arn | `string` | computed |  | The SchemaMapping arn associated with the Schema |
| `SchemaName` | schema_name | `string` | required, replaces on change |  | The name of the SchemaMapping |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | The time of this SchemaMapping got last updated at |

Supports update: yes

Discovery: supported
