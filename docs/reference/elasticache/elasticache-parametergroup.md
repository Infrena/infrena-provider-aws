# aws.elasticache.parametergroup

**CloudFormation type:** `AWS::ElastiCache::ParameterGroup`

Resource Type definition for AWS::ElastiCache::ParameterGroup

Region attribute: `region`

**Import ID:** `<region>/CacheParameterGroupName` (AWS::ElastiCache::ParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CacheParameterGroupFamily` | cache_parameter_group_family | `string` | required, replaces on change |  | The name of the cache parameter group family that this cache parameter group is compatible with. |
| `CacheParameterGroupName` | cache_parameter_group_name | `string` | computed |  | The name of the Cache Parameter Group. |
| `Description` |  | `string` | required |  | The description for this cache parameter group. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | A comma-delimited list of parameter name/value pairs. For more information see ModifyCacheParameterGroup in the Amazon ElastiCache API Reference Guide. |
| `Tags` |  | `map` | tags map |  | Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted. |

Supports update: yes

Discovery: supported
