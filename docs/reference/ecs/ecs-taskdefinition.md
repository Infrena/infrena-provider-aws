# aws.ecs.taskdefinition

**CloudFormation type:** `AWS::ECS::TaskDefinition`

Registers a new task definition from the supplied ``family`` and ``containerDefinitions``. Optionally, you can add data volumes to your containers with the ``volumes`` parameter. For more information about task definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/TaskDefinitionArn` (AWS::ECS::TaskDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContainerDefinitions` | container_definitions | `list` | optional, computed, provider-chosen, replaces on change |  | A list of container definitions in JSON format that describe the different containers that make up your task. For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `Cpu` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The number of ``cpu`` units used by the task. If you use the EC2 launch type, this field is optional. Any value can be used. If you use the Fargate launch type, this field is required. You must use one of the following values. The value that you choose determines your range of valid values for the ``memory`` parameter. |
| `EnableFaultInjection` | enable_fault_injection | `boolean` | optional, computed, provider-chosen, replaces on change |  | Enables fault injection and allows for fault injection requests to be accepted from the task's containers. The default value is ``false``. |
| `EphemeralStorage` | ephemeral_storage | `map` | optional, computed, provider-chosen, replaces on change |  | The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on FARGATElong. For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon ECS Developer Guide;*. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf. For informationabout the required IAM roles for Amazon ECS, see [IAM roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/security-ecs-iam-role-overview.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `Family` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of a family that this task definition is registered to. Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed. |
| `InferenceAccelerators` | inference_accelerators | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `IpcMode` | ipc_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The IPC resource namespace to use for the containers in the task. The valid values are ``host``, ``task``, or ``none``. If ``host`` is specified, then all containers within the tasks that specified the ``host`` IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance. If ``task`` is specified, all containers within the specified task share the same IPC resources. If ``none`` is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance. If no value is specified, then the IPC resource namespace sharing depends on the Docker daemon setting on the container instance. |
| `Memory` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The amount (in MiB) of memory used by the task. |
| `NetworkMode` | network_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The Docker networking mode to use for the containers in the task. The valid values are ``none``, ``bridge``, ``awsvpc``, and ``host``. If no network mode is specified, the default is ``bridge``. |
| `PidMode` | pid_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The process namespace to use for the containers in the task. The valid values are ``host`` or ``task``. On Fargate for Linux containers, the only valid value is ``task``. For example, monitoring sidecars might need ``pidMode`` to access information about other containers running in the same task. |
| `PlacementConstraints` | placement_constraints | `list` | optional, computed, provider-chosen, replaces on change |  | An array of placement constraint objects to use for tasks. |
| `ProxyConfiguration` | proxy_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration details for the App Mesh proxy. |
| `RequiresCompatibilities` | requires_compatibilities | `list` | optional, computed, provider-chosen, replaces on change |  | The task launch types the task definition was validated against. The valid values are ``MANAGED_INSTANCES``, ``EC2``, ``FARGATE``, and ``EXTERNAL``. For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `RuntimePlatform` | runtime_platform | `map` | optional, computed, provider-chosen, replaces on change |  | Information about the platform for the Amazon ECS service or task. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The metadata that you apply to the task definition to help you categorize and organize them. Each tag consists of a key and an optional value. You define both of them. |
| `TaskDefinitionArn` | task_definition_arn | `string` | computed |  |  |
| `TaskRoleArn` | task_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The short name or full Amazon Resource Name (ARN) of the IAMlong role that grants containers in the task permission to call AWS APIs on your behalf. For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `Volumes` |  | `list` | optional, computed, provider-chosen, replaces on change |  | The list of data volume definitions for the task. For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide*. |

Supports update: yes

Discovery: supported
