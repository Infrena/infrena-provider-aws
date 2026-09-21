# aws.elasticache.subnetgroup

**CloudFormation type:** `AWS::ElastiCache::SubnetGroup`

Resource Type definition for AWS::ElastiCache::SubnetGroup

Region attribute: `region`

**Import ID:** `<region>/CacheSubnetGroupName` (AWS::ElastiCache::SubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CacheSubnetGroupName` | cache_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the cache subnet group. This value is stored as a lowercase string. |
| `Description` |  | `string` | required |  | The description for the cache subnet group. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The EC2 subnet IDs for the cache subnet group. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
