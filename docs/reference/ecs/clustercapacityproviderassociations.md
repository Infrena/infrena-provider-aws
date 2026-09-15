# aws.clustercapacityproviderassociations

**CloudFormation type:** `AWS::ECS::ClusterCapacityProviderAssociations`

Associate a set of ECS Capacity Providers with a specified ECS Cluster

Region attribute: `region`

**Import ID:** `<region>/Cluster` (AWS::ECS::ClusterCapacityProviderAssociations)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityProviders` | capacity_providers | `list` | optional, computed, provider-chosen |  | List of capacity providers to associate with the cluster |
| `Cluster` |  | `string` | required, replaces on change |  | The name of the cluster |
| `DefaultCapacityProviderStrategy` | default_capacity_provider_strategy | `list` | required |  | List of capacity providers to associate with the cluster |

Supports update: yes

Discovery: supported
