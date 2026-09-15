# aws.taskset

**CloudFormation type:** `AWS::ECS::TaskSet`

Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.

Region attribute: `region`

**Import ID:** `<region>/Cluster|Service|Id` (AWS::ECS::TaskSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityProviderStrategy` | capacity_provider_strategy | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Cluster` |  | `string` | required, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in. |
| `ExternalId` | external_id | `string` | optional, computed, provider-chosen, replaces on change |  | An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. |
| `Id` |  | `string` | computed |  | The ID of the task set. |
| `LaunchType` | launch_type | `string` | optional, computed, provider-chosen, replaces on change |  | The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. |
| `LoadBalancers` | load_balancers | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing the network configuration for a task or service. |
| `PlatformVersion` | platform_version | `string` | optional, computed, provider-chosen, replaces on change |  | The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default. |
| `Scale` |  | `map` | optional, computed, provider-chosen |  | A floating-point percentage of the desired number of tasks to place and keep running in the task set. |
| `Service` |  | `string` | required, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the service to create the task set in. |
| `ServiceRegistries` | service_registries | `list` | optional, computed, provider-chosen, replaces on change |  | The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TaskDefinition` | task_definition | `string` | required, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use. |

Supports update: yes

Discovery: not supported
