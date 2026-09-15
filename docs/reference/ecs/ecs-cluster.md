# aws.ecs.cluster

**CloudFormation type:** `AWS::ECS::Cluster`

The ``AWS::ECS::Cluster`` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.

Region attribute: `region`

**Import ID:** `<region>/ClusterName` (AWS::ECS::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CapacityProviders` | capacity_providers | `list` | optional, computed, provider-chosen |  | The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions. |
| `ClusterName` | name, cluster_name | `string` | optional, computed, provider-chosen, replaces on change |  | A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name. |
| `ClusterSettings` | cluster_settings | `list` | optional, computed, provider-chosen |  | The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights with enhanced observability or CloudWatch Container Insights for a cluster. |
| `Configuration` |  | `map` | optional, computed, provider-chosen |  | The execute command and managed storage configuration for the cluster. |
| `DefaultCapacityProviderStrategy` | default_capacity_provider_strategy | `list` | optional, computed, provider-chosen |  | The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used. |
| `ServiceConnectDefaults` | service_connect_defaults | `map` | optional, computed, provider-chosen, write-only |  | Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both. |

Supports update: yes

Discovery: supported
