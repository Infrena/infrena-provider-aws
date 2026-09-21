# aws.daemontaskdefinition

**CloudFormation type:** `AWS::ECS::DaemonTaskDefinition`

The details of a daemon task definition. A daemon task definition is a template that describes the containers that form a daemon. Daemons deploy cross-cutting software agents independently across your Amazon ECS infrastructure.

Region attribute: `region`

**Import ID:** `<region>/DaemonTaskDefinitionArn` (AWS::ECS::DaemonTaskDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContainerDefinitions` | container_definitions | `list` | optional, computed, provider-chosen, replaces on change |  | A list of container definitions in JSON format that describe the containers that make up the daemon task. |
| `Cpu` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The number of CPU units used by the daemon task. |
| `DaemonTaskDefinitionArn` | daemon_task_definition_arn | `string` | computed |  |  |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make Amazon Web Services API calls on your behalf. |
| `Family` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of a family that this daemon task definition is registered to. |
| `IpcMode` | ipc_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The IPC namespace mode for the daemon. The valid values are ``none`` and ``shared``. The default is ``none``. |
| `Memory` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The amount of memory (in MiB) used by the daemon task. |
| `PidMode` | pid_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The PID namespace mode for the daemon. The valid values are ``none`` and ``shared``. The default is ``none``. |
| `Tags` |  | `map` | tags map |  |  |
| `TaskRoleArn` | task_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The short name or full Amazon Resource Name (ARN) of the IAM role that grants containers in the daemon task permission to call Amazon Web Services APIs on your behalf. |
| `Volumes` |  | `list` | optional, computed, provider-chosen, replaces on change |  | The list of data volume definitions for the daemon task. |

Supports update: yes

Discovery: supported
