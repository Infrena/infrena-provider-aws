# aws.serverlesscachesnapshot

**CloudFormation type:** `AWS::ElastiCache::ServerlessCacheSnapshot`

Resource Type definition for AWS::ElastiCache::ServerlessCacheSnapshot. A serverless cache snapshot is a point-in-time backup of an ElastiCache serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.

Region attribute: `region`

**Import ID:** `<region>/ARN` (AWS::ElastiCache::ServerlessCacheSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARN` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the serverless cache snapshot. |
| `BytesUsedForCache` | bytes_used_for_cache | `string` | computed |  | The total size of the serverless cache snapshot, in bytes. |
| `CreateTime` | create_time | `string` | computed |  | The date and time that the source serverless cache's metadata and cache data set was obtained for the snapshot. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the AWS KMS key used to encrypt the snapshot. Provide the key ARN: the resource returns the key ARN on read, so supplying a bare key ID or alias for this createOnly property may be reported as drift by CloudFormation. |
| `ServerlessCacheConfiguration` | serverless_cache_configuration | `map` | computed |  | The configuration of the serverless cache, at the time the snapshot was taken. |
| `ServerlessCacheName` | serverless_cache_name | `string` | required, replaces on change |  | The name of an existing serverless cache. The snapshot is created from this cache. |
| `ServerlessCacheSnapshotName` | serverless_cache_snapshot_name | `string` | required, replaces on change |  | The name of the serverless cache snapshot. Must be unique for the customer account. This value is stored as a lowercase string. |
| `SnapshotType` | snapshot_type | `string` | computed |  | The type of snapshot of the serverless cache. |
| `Status` |  | `string` | computed |  | The current status of the serverless cache snapshot. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to be added to the serverless cache snapshot resource. |

Supports update: yes

Discovery: supported
