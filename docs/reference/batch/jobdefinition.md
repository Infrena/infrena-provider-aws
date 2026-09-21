# aws.jobdefinition

**CloudFormation type:** `AWS::Batch::JobDefinition`

Resource Type definition for AWS::Batch::JobDefinition

Region attribute: `region`

**Import ID:** `<region>/JobDefinitionName` (AWS::Batch::JobDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConsumableResourceProperties` | consumable_resource_properties | `map` | optional, computed, provider-chosen |  |  |
| `ContainerProperties` | container_properties | `map` | optional, computed, provider-chosen |  |  |
| `EcsProperties` | ecs_properties | `map` | optional, computed, provider-chosen |  |  |
| `EksProperties` | eks_properties | `map` | optional, computed, provider-chosen |  |  |
| `JobDefinitionArn` | job_definition_arn | `string` | computed |  |  |
| `JobDefinitionName` | job_definition_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `NodeProperties` | node_properties | `map` | optional, computed, provider-chosen |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  |  |
| `PlatformCapabilities` | platform_capabilities | `list` | optional, computed, provider-chosen |  |  |
| `PropagateTags` | propagate_tags | `boolean` | optional, computed, provider-chosen |  |  |
| `ResourceRetentionPolicy` | resource_retention_policy | `map` | optional, computed, provider-chosen, write-only |  |  |
| `RetryStrategy` | retry_strategy | `map` | optional, computed, provider-chosen |  |  |
| `SchedulingPriority` | scheduling_priority | `integer` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `Timeout` |  | `map` | optional, computed, provider-chosen |  |  |
| `Type` | type_value | `string` | required |  |  |

Supports update: yes

Discovery: supported
