# aws.ecs.service

**CloudFormation type:** `AWS::ECS::Service`

The ``AWS::ECS::Service`` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.

Region attribute: `region`

**Import ID:** `<region>/ServiceArn|Cluster` (AWS::ECS::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZoneRebalancing` | availability_zone_rebalancing | `string` | optional, computed, provider-chosen |  | Indicates whether to use Availability Zone rebalancing for the service. |
| `CapacityProviderStrategy` | capacity_provider_strategy | `list` | optional, computed, provider-chosen |  | The capacity provider strategy to use for the service. |
| `Cluster` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on. If you do not specify a cluster, the default cluster is assumed. |
| `DeploymentConfiguration` | deployment_configuration | `map` | optional, computed, provider-chosen |  | Optional deployment parameters that control how many tasks run during a deployment and the ordering of stopping and starting tasks. |
| `DeploymentController` | deployment_controller | `map` | optional, computed, provider-chosen |  | The deployment controller to use for the service. |
| `DesiredCount` | desired_count | `integer` | optional, computed, provider-chosen |  | The number of instantiations of the specified task definition to place and keep running in your service. |
| `EnableECSManagedTags` | enable_ecs_managed_tags | `boolean` | optional, computed, provider-chosen |  | Specifies whether to turn on Amazon ECS managed tags for the tasks within the service. For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `EnableExecuteCommand` | enable_execute_command | `boolean` | optional, computed, provider-chosen |  | Determines whether the execute command functionality is turned on for the service. If ``true``, the execute command functionality is turned on for all containers in tasks as part of the service. |
| `ForceNewDeployment` | force_new_deployment | `map` | optional, computed, provider-chosen, write-only |  | Determines whether to force a new deployment of the service. By default, deployments aren't forced. You can use this option to start a new deployment with no service definition changes. For example, you can update a service's tasks to use a newer Docker image with the same image/tag combination (``my_image:latest``) or to roll Fargate tasks onto a newer platform version. |
| `HealthCheckGracePeriodSeconds` | health_check_grace_period_seconds | `integer` | optional, computed, provider-chosen |  | The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task has first started. If you do not specify a health check grace period value, the default value of 0 is used. If you do not use any of the health checks, then ``healthCheckGracePeriodSeconds`` is unused. |
| `LaunchType` | launch_type | `string` | optional, computed, provider-chosen |  | The launch type on which to run your service. For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `LoadBalancers` | load_balancers | `list` | optional, computed, provider-chosen |  | A list of load balancer objects to associate with the service. If you specify the ``Role`` property, ``LoadBalancers`` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `Monitoring` |  | `map` | optional, computed, provider-chosen, write-only |  | The optional monitoring configuration for a service, which defines the resolution for the service-level ``CPUUtilization`` and ``MemoryUtilization`` Amazon CloudWatch metrics. When not specified, Amazon ECS uses the default resolution of ``60`` seconds. |
| `Name` |  | `string` | computed |  |  |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  | The network configuration for a task or service. |
| `PlacementConstraints` | placement_constraints | `list` | optional, computed, provider-chosen |  | An array of placement constraint objects to use for tasks in your service. You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime. |
| `PlacementStrategies` | placement_strategies | `list` | optional, computed, provider-chosen |  | The placement strategy objects to use for tasks in your service. You can specify a maximum of 5 strategy rules for each service. |
| `PlatformVersion` | platform_version | `string` | optional, computed, provider-chosen |  | The platform version that your tasks in the service are running on. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the ``LATEST`` platform version is used. For more information, see [platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide*. |
| `PropagateTags` | propagate_tags | `string` | optional, computed, provider-chosen |  | Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags aren't propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action. |
| `Role` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the ``awsvpc`` network mode. If you specify the ``role`` parameter, you must also specify a load balancer object with the ``loadBalancers`` parameter. |
| `SchedulingStrategy` | scheduling_strategy | `string` | optional, computed, provider-chosen, replaces on change |  | The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html). |
| `ServiceArn` | service_arn | `string` | computed |  |  |
| `ServiceConnectConfiguration` | service_connect_configuration | `map` | optional, computed, provider-chosen, write-only |  | The Service Connect configuration of your Amazon ECS service. The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace. |
| `ServiceName` | service_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of your service. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions. |
| `ServiceRegistries` | service_registries | `list` | optional, computed, provider-chosen |  | The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html). |
| `Tags` |  | `map` | tags map |  | The metadata that you apply to the service to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well. |
| `TaskDefinition` | task_definition | `string` | optional, computed, provider-chosen |  | The ``family`` and ``revision`` (``family:revision``) or full ARN of the task definition to run in your service. If a ``revision`` isn't specified, the latest ``ACTIVE`` revision is used. |
| `VolumeConfigurations` | volume_configurations | `list` | optional, computed, provider-chosen, write-only |  | The configuration for a volume specified in the task definition as a volume that is configured at launch time. Currently, the only supported volume type is an Amazon EBS volume. |
| `VpcLatticeConfigurations` | vpc_lattice_configurations | `list` | optional, computed, provider-chosen |  | The VPC Lattice configuration for the service being created. |

Supports update: yes

Discovery: supported
