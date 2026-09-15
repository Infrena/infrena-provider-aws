# aws.transfer.workflow

**CloudFormation type:** `AWS::Transfer::Workflow`

Resource Type definition for AWS::Transfer::Workflow

Region attribute: `region`

**Import ID:** `<region>/WorkflowId` (AWS::Transfer::Workflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the workflow. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A textual description for the workflow. |
| `OnExceptionSteps` | on_exception_steps | `list` | optional, computed, provider-chosen, replaces on change |  | Specifies the steps (actions) to take if any errors are encountered during execution of the workflow. |
| `Steps` |  | `list` | required, replaces on change |  | Specifies the details for the steps that are in the specified workflow. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose. |
| `WorkflowId` | workflow_id | `string` | computed |  | A unique identifier for the workflow. |

Supports update: yes

Discovery: supported
