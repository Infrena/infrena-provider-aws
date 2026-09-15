# aws.maintenancewindowtask

**CloudFormation type:** `AWS::SSM::MaintenanceWindowTask`

Resource Type definition for AWS::SSM::MaintenanceWindowTask

Region attribute: `region`

**Import ID:** `<region>/WindowId|WindowTaskId` (AWS::SSM::MaintenanceWindowTask)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CutoffBehavior` | cutoff_behavior | `string` | optional, computed, provider-chosen |  | The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the task. |
| `LoggingInfo` | logging_info | `map` | optional, computed, provider-chosen |  | Information about an Amazon S3 bucket to write Run Command task-level logs to. |
| `MaxConcurrency` | max_concurrency | `string` | optional, computed, provider-chosen |  | The maximum number of targets this task can be run for, in parallel. |
| `MaxErrors` | max_errors | `string` | optional, computed, provider-chosen |  | The maximum number of errors allowed before this task stops being scheduled. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The task name. |
| `Priority` |  | `integer` | required |  | The priority of the task in the maintenance window. The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel. |
| `ServiceRoleArn` | service_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM service role for AWS Systems Manager to assume when running a maintenance window task. |
| `Targets` |  | `list` | optional, computed, provider-chosen |  | The targets (either instances or window target ids). |
| `TaskArn` | task_arn | `string` | required |  | The resource that the task uses during execution. |
| `TaskInvocationParameters` | task_invocation_parameters | `map` | optional, computed, provider-chosen |  | The parameters to pass to the task when it runs. Populate only the fields that match the task type. All other fields should be empty. |
| `TaskParameters` | task_parameters | `map` | optional, computed, provider-chosen |  | The parameters to pass to the task when it runs. |
| `TaskType` | task_type | `string` | required, replaces on change |  | The type of task. |
| `WindowId` | window_id | `string` | required, replaces on change |  | The ID of the maintenance window where the task is registered. |
| `WindowTaskId` | window_task_id | `string` | computed |  | Unique identifier of the maintenance window task. |

Supports update: yes

Discovery: supported (parent resource required)
