# aws.serverlesscache

**CloudFormation type:** `AWS::ElastiCache::ServerlessCache`

The AWS::ElastiCache::ServerlessCache resource creates an Amazon ElastiCache Serverless Cache.

Region attribute: `region`

**Import ID:** `<region>/ServerlessCacheName` (AWS::ElastiCache::ServerlessCache)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARN` |  | `string` | computed |  | The ARN of the Serverless Cache. |
| `CacheUsageLimits` | cache_usage_limits | `map` | optional, computed, provider-chosen |  | The cache capacity limit of the Serverless Cache. |
| `CreateTime` | create_time | `string` | computed |  | The creation time of the Serverless Cache. |
| `DailySnapshotTime` | daily_snapshot_time | `string` | optional, computed, provider-chosen |  | The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Serverless Cache. |
| `Endpoint` |  | `map` | optional, computed, provider-chosen |  | The address and the port. |
| `Engine` |  | `string` | required |  | The engine name of the Serverless Cache. |
| `FinalSnapshotName` | final_snapshot_name | `string` | optional, computed, provider-chosen, write-only |  | The final snapshot name which is taken before Serverless Cache is deleted. |
| `FullEngineVersion` | full_engine_version | `string` | computed |  | The full engine version of the Serverless Cache. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the KMS key used to encrypt the cluster. |
| `MajorEngineVersion` | major_engine_version | `string` | optional, computed, provider-chosen |  | The major engine version of the Serverless Cache. |
| `ReaderEndpoint` | reader_endpoint | `map` | optional, computed, provider-chosen |  | The address and the port. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | One or more Amazon VPC security groups associated with this Serverless Cache. |
| `ServerlessCacheName` | serverless_cache_name | `string` | required, replaces on change |  | The name of the Serverless Cache. This value must be unique. |
| `SnapshotArnsToRestore` | snapshot_arns_to_restore | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN's of snapshot to restore Serverless Cache. |
| `SnapshotRetentionLimit` | snapshot_retention_limit | `integer` | optional, computed, provider-chosen |  | The snapshot retention limit of the Serverless Cache. |
| `Status` |  | `string` | computed |  | The status of the Serverless Cache. |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | The subnet id's of the Serverless Cache. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this Serverless Cache. |
| `UserGroupId` | user_group_id | `string` | optional, computed, provider-chosen | aws.usergroup.UserGroupId | The ID of the user group. |

Supports update: yes

Discovery: supported
