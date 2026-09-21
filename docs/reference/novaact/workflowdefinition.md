# aws.workflowdefinition

**CloudFormation type:** `AWS::NovaAct::WorkflowDefinition`

Definition of AWS::NovaAct::WorkflowDefinition Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::NovaAct::WorkflowDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the workflow definition. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the workflow definition was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | An optional description of the workflow definition's purpose and functionality. |
| `ExportConfig` | export_config | `map` | optional, computed, provider-chosen, replaces on change |  | Configuration settings for exporting workflow execution data and logs to Amazon S3. |
| `Name` |  | `string` | required, replaces on change |  | The name of the workflow definition. Must be unique within your account and region. |
| `Status` |  | `string` | computed |  | The current status of the workflow definition. |

Supports update: no

Discovery: supported
