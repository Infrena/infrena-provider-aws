# aws.idmappingworkflow

**CloudFormation type:** `AWS::EntityResolution::IdMappingWorkflow`

IdMappingWorkflow defined in AWS Entity Resolution service

Region attribute: `region`

**Import ID:** `<region>/WorkflowName` (AWS::EntityResolution::IdMappingWorkflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this IdMappingWorkflow got created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the IdMappingWorkflow |
| `IdMappingIncrementalRunConfig` | id_mapping_incremental_run_config | `map` | optional, computed, provider-chosen |  |  |
| `IdMappingTechniques` | id_mapping_techniques | `map` | required |  |  |
| `InputSourceConfig` | input_source_config | `list` | required |  |  |
| `OutputSourceConfig` | output_source_config | `list` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | The time of this IdMappingWorkflow got last updated at |
| `WorkflowArn` | workflow_arn | `string` | computed |  | The default IdMappingWorkflow arn |
| `WorkflowName` | workflow_name | `string` | required, replaces on change |  | The name of the IdMappingWorkflow |

Supports update: yes

Discovery: supported
