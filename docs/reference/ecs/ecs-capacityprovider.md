# aws.ecs.capacityprovider

**CloudFormation type:** `AWS::ECS::CapacityProvider`

Resource Type definition for AWS::ECS::CapacityProvider.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::ECS::CapacityProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingGroupProvider` | auto_scaling_group_provider | `map` | optional, computed, provider-chosen |  |  |
| `ClusterName` | cluster_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ManagedInstancesProvider` | managed_instances_provider | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
