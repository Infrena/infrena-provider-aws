# aws.eventschemas.registry

**CloudFormation type:** `AWS::EventSchemas::Registry`

Resource Type definition for AWS::EventSchemas::Registry

Region attribute: `region`

**Import ID:** `<region>/RegistryArn` (AWS::EventSchemas::Registry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the registry to be created. |
| `RegistryArn` | registry_arn | `string` | computed |  | The ARN of the registry. |
| `RegistryName` | registry_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the schema registry. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with the resource. |

Supports update: yes

Discovery: supported
