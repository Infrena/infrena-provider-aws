# aws.iotsitewise.task

**CloudFormation type:** `AWS::IoTSiteWise::Task`

Resource schema for AWS::IoTSiteWise::Task. A task defines a reusable containerized compute workload that can be referenced by one or more pipeline compute nodes.

Region attribute: `region`

**Import ID:** `<region>/TaskArn` (AWS::IoTSiteWise::Task)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the task. |
| `Status` |  | `string` | computed |  | The current lifecycle status of the task. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TaskArn` | task_arn | `string` | computed |  | The ARN of the task. |
| `TaskConfiguration` | task_configuration | `map` | required |  | The task execution configuration. |
| `TaskName` | task_name | `string` | required, replaces on change |  | The name of the task. Must be unique within the workspace. |
| `WorkspaceName` | workspace_name | `string` | required, replaces on change |  | The name of the workspace. |

Supports update: yes

Discovery: supported (parent resource required)
