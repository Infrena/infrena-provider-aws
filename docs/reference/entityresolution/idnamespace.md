# aws.idnamespace

**CloudFormation type:** `AWS::EntityResolution::IdNamespace`

IdNamespace defined in AWS Entity Resolution service

Region attribute: `region`

**Import ID:** `<region>/IdNamespaceName` (AWS::EntityResolution::IdNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The date and time when the IdNamespace was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IdMappingWorkflowProperties` | id_mapping_workflow_properties | `list` | optional, computed, provider-chosen |  |  |
| `IdNamespaceArn` | id_namespace_arn | `string` | computed |  | The arn associated with the IdNamespace |
| `IdNamespaceName` | id_namespace_name | `string` | required, replaces on change |  |  |
| `InputSourceConfig` | input_source_config | `list` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | required |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time when the IdNamespace was updated |

Supports update: yes

Discovery: supported
