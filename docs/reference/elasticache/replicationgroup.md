# aws.replicationgroup

**CloudFormation type:** `AWS::ElastiCache::ReplicationGroup`

Resource type definition for AWS::ElastiCache::ReplicationGroup

Region attribute: `region`

**Import ID:** `<region>/ReplicationGroupId` (AWS::ElastiCache::ReplicationGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AtRestEncryptionEnabled` | at_rest_encryption_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | A flag that enables encryption at rest when set to true.AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group. |
| `AuthToken` | auth_token | `string` | optional, computed, provider-chosen, sensitive, write-only |  | Reserved parameter. The password used to access a password protected server.AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true. For more information. |
| `AutoMinorVersionUpgrade` | auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  | This parameter is currently disabled. |
| `AutomaticFailoverEnabled` | automatic_failover_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails. AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups. |
| `CacheNodeType` | cache_node_type | `string` | optional, computed, provider-chosen |  | The compute and memory capacity of the nodes in the node group (shard). |
| `CacheParameterGroupName` | cache_parameter_group_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used. |
| `CacheSecurityGroupNames` | cache_security_group_names | `list` | optional, computed, provider-chosen, write-only |  | A list of cache security group names to associate with this replication group. |
| `CacheSubnetGroupName` | cache_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the cache subnet group to be used for the replication group. |
| `ClusterMode` | cluster_mode | `string` | optional, computed, provider-chosen |  | Enabled or Disabled. To modify cluster mode from Disabled to Enabled, you must first set the cluster mode to Compatible. Compatible mode allows your Redis OSS clients to connect using both cluster mode enabled and cluster mode disabled. After you migrate all Redis OSS clients to use cluster mode enabled, you can then complete cluster mode configuration and set the cluster mode to Enabled. For more information, see Modify cluster mode. |
| `ConfigurationEndPoint` | configuration_end_point | `map` | computed |  | The configuration details of the replication group. |
| `DataTieringEnabled` | data_tiering_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. |
| `Durability` |  | `string` | optional, computed, provider-chosen |  | The durability setting for the replication group. Valid values: default, async, sync, disabled. Enabling durability on an existing non-durable cluster or disabling durability on an existing durable cluster is not currently supported and will result in an error; specify the desired durability at create time. The resolved state is returned in EffectiveDurability. |
| `EffectiveDurability` | effective_durability | `string` | computed |  | The resolved durability state of the replication group after resolving the default value. This is a read-only property. |
| `Engine` |  | `string` | optional, computed, provider-chosen |  | The name of the cache engine to be used for the clusters in this replication group. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen, write-only |  | The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation. |
| `GlobalReplicationGroupId` | global_replication_group_id | `string` | optional, computed, provider-chosen, replaces on change | aws.globalreplicationgroup.GlobalReplicationGroupId | The name of the Global datastore |
| `IpDiscovery` | ip_discovery | `string` | optional, computed, provider-chosen |  | The network type you choose when creating a replication group, either ipv4 \| ipv6. IPv6 is supported for workloads using Redis OSS engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the Nitro system. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the KMS key used to encrypt the disk on the cluster. |
| `LogDeliveryConfigurations` | log_delivery_configurations | `list` | optional, computed, provider-chosen |  | Specifies the destination, format and type of the logs. |
| `MultiAZEnabled` | multi_az_enabled | `boolean` | optional, computed, provider-chosen |  | A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | Must be either ipv4 \| ipv6 \| dual_stack. IPv6 is supported for workloads using Redis OSS engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the Nitro system |
| `NodeGroupConfiguration` | node_group_configuration | `list` | optional, computed, provider-chosen, write-only |  | NodeGroupConfiguration is a property of the AWS::ElastiCache::ReplicationGroup resource that configures an Amazon ElastiCache (ElastiCache) Redis cluster node group. |
| `NotificationTopicArn` | notification_topic_arn | `string` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent. |
| `NumCacheClusters` | num_cache_clusters | `integer` | optional, computed, provider-chosen |  | The number of clusters this replication group initially has.This parameter is not used if there is more than one node group (shard). You should use ReplicasPerNodeGroup instead. |
| `NumNodeGroups` | num_node_groups | `integer` | optional, computed, provider-chosen, write-only |  | An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1. |
| `Port` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The port number on which each member of the replication group accepts connections. |
| `PreferredCacheClusterAZs` | preferred_cache_cluster_a_zs | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list. This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead. |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen, write-only |  | Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. |
| `PrimaryClusterId` | primary_cluster_id | `string` | optional, computed, provider-chosen, write-only |  | The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available. |
| `PrimaryEndPoint` | primary_end_point | `map` | computed |  | The primary endpoint configuration |
| `ReadEndPoint` | read_end_point | `map` | computed |  |  |
| `ReaderEndPoint` | reader_end_point | `map` | computed |  | The endpoint of the reader node in the replication group. |
| `ReplicasPerNodeGroup` | replicas_per_node_group | `integer` | optional, computed, provider-chosen |  | An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5. |
| `ReplicationGroupDescription` | replication_group_description | `string` | required |  | A user-created description for the replication group. |
| `ReplicationGroupId` | replication_group_id | `string` | optional, computed, provider-chosen, replaces on change | aws.replicationgroup.ReplicationGroupId | The replication group identifier. This parameter is stored as a lowercase string. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, write-only | aws.securitygroup.Id | One or more Amazon VPC security groups associated with this replication group. |
| `SnapshotArns` | snapshot_arns | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3. |
| `SnapshotName` | snapshot_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created. |
| `SnapshotRetentionLimit` | snapshot_retention_limit | `integer` | optional, computed, provider-chosen |  | The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted. |
| `SnapshotWindow` | snapshot_window | `string` | optional, computed, provider-chosen |  | The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard). |
| `SnapshottingClusterId` | snapshotting_cluster_id | `string` | optional, computed, provider-chosen |  | The cluster ID that is used as the daily snapshot source for the replication group. This parameter cannot be set for Redis (cluster mode enabled) replication groups. |
| `Tags` |  | `map` | tags map |  | A list of cost allocation tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue. |
| `TransitEncryptionEnabled` | transit_encryption_enabled | `boolean` | optional, computed, provider-chosen |  | A flag that enables in-transit encryption when set to true. |
| `TransitEncryptionMode` | transit_encryption_mode | `string` | optional, computed, provider-chosen |  | A setting that allows you to migrate your clients to use in-transit encryption, with no downtime. When setting TransitEncryptionEnabled to true, you can set your TransitEncryptionMode to preferred in the same request, to allow both encrypted and unencrypted connections at the same time. Once you migrate all your Redis OSS clients to use encrypted connections you can modify the value to required to allow encrypted connections only. Setting TransitEncryptionMode to required is a two-step process that requires you to first set the TransitEncryptionMode to preferred, after that you can set TransitEncryptionMode to required. This process will not trigger the replacement of the replication group. |
| `UserGroupIds` | user_group_ids | `list` | optional, computed, provider-chosen | aws.usergroup.UserGroupId | The ID of user group to associate with the replication group. |

Supports update: yes

Discovery: supported
