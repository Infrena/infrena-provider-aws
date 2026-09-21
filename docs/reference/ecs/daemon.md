# aws.daemon

**CloudFormation type:** `AWS::ECS::Daemon`

Information about a daemon resource.

Region attribute: `region`

**Import ID:** `<region>/DaemonArn` (AWS::ECS::Daemon)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityProviderArns` | capacity_provider_arns | `list` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Names (ARNs) of the capacity providers associated with the daemon. |
| `ClusterArn` | cluster_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.ecs.cluster.Arn | The Amazon Resource Name (ARN) of the cluster that the daemon is running in. |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DaemonArn` | daemon_arn | `string` | computed |  |  |
| `DaemonName` | daemon_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `DaemonStatus` | daemon_status | `string` | computed |  |  |
| `DaemonTaskDefinitionArn` | daemon_task_definition_arn | `string` | optional, computed, provider-chosen, write-only | aws.daemontaskdefinition.DaemonTaskDefinitionArn | The Amazon Resource Name (ARN) of the daemon task definition used by this revision. |
| `DeploymentArn` | deployment_arn | `string` | computed |  |  |
| `DeploymentConfiguration` | deployment_configuration | `map` | optional, computed, provider-chosen, write-only |  | Optional deployment parameters that control how a daemon rolls out updates across container instances. |
| `EnableECSManagedTags` | enable_ecs_managed_tags | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether Amazon ECS managed tags are turned on for the daemon tasks. |
| `EnableExecuteCommand` | enable_execute_command | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether the execute command functionality is turned on for the daemon tasks. |
| `PropagateTags` | propagate_tags | `string` | optional, computed, provider-chosen, write-only |  | Specifies whether tags are propagated from the daemon to the daemon tasks. |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
