# aws.expressgatewayservice

**CloudFormation type:** `AWS::ECS::ExpressGatewayService`

Resource Type definition for AWS::ECS::ExpressGatewayService

Region attribute: `region`

**Import ID:** `<region>/ServiceArn` (AWS::ECS::ExpressGatewayService)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActiveConfigurations` | active_configurations | `list` | computed |  |  |
| `Cluster` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Cpu` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `ECSManagedResourceArns` | ecs_managed_resource_arns | `map` | computed |  |  |
| `Endpoint` |  | `string` | computed |  |  |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, write-only | aws.role.Arn |  |
| `HealthCheckPath` | health_check_path | `string` | optional, computed, provider-chosen, write-only |  |  |
| `InfrastructureRoleArn` | infrastructure_role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `Memory` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `PrimaryContainer` | primary_container | `map` | optional, computed, provider-chosen, write-only |  |  |
| `ScalingTarget` | scaling_target | `map` | optional, computed, provider-chosen, write-only |  |  |
| `ServiceArn` | service_arn | `string` | computed |  |  |
| `ServiceName` | service_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `map` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  |  |
| `TaskDefinitionArn` | task_definition_arn | `string` | optional, computed, provider-chosen, write-only | aws.ecs.taskdefinition.TaskDefinitionArn |  |
| `TaskRoleArn` | task_role_arn | `string` | optional, computed, provider-chosen, write-only | aws.role.Arn |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: not supported
