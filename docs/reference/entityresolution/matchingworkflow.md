# aws.matchingworkflow

**CloudFormation type:** `AWS::EntityResolution::MatchingWorkflow`

MatchingWorkflow defined in AWS Entity Resolution service

Region attribute: `region`

**Import ID:** `<region>/WorkflowName` (AWS::EntityResolution::MatchingWorkflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this MatchingWorkflow got created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the MatchingWorkflow |
| `IncrementalRunConfig` | incremental_run_config | `map` | optional, computed, provider-chosen |  |  |
| `InputSourceConfig` | input_source_config | `list` | required |  |  |
| `OutputSourceConfig` | output_source_config | `list` | required |  |  |
| `ResolutionTechniques` | resolution_techniques | `map` | required |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | The time of this MatchingWorkflow got last updated at |
| `WorkflowArn` | workflow_arn | `string` | computed |  | The default MatchingWorkflow arn |
| `WorkflowName` | workflow_name | `string` | required, replaces on change |  | The name of the MatchingWorkflow |

Supports update: yes

Discovery: supported
