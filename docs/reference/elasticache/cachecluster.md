# aws.cachecluster

**CloudFormation type:** `AWS::ElastiCache::CacheCluster`

Resource type definition for AWS::ElastiCache::CacheCluster

Region attribute: `region`

**Import ID:** `<region>/ClusterName` (AWS::ElastiCache::CacheCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AZMode` | az_mode | `string` | optional, computed, provider-chosen |  | Specifies whether the nodes in this Memcached cluster are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. |
| `AutoMinorVersionUpgrade` | auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  | If you are running Redis engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign. |
| `CacheNodeType` | cache_node_type | `string` | required |  | The compute and memory capacity of the nodes in the node group (shard). |
| `CacheParameterGroupName` | cache_parameter_group_name | `string` | optional, computed, provider-chosen |  | The name of the parameter group to associate with this cluster. |
| `CacheSecurityGroupNames` | cache_security_group_names | `list` | optional, computed, provider-chosen, write-only |  | A list of security group names to associate with this cluster. |
| `CacheSubnetGroupName` | cache_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the subnet group to be used for the cluster. |
| `ClusterName` | cluster_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the cache cluster. |
| `ConfigurationEndpoint` | configuration_endpoint | `map` | computed |  | Specifies the ConfigurationEndpoint address and port |
| `Engine` |  | `string` | required, replaces on change |  | The name of the cache engine to be used for this cluster. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | The version number of the cache engine to be used for this cluster |
| `IpDiscovery` | ip_discovery | `string` | optional, computed, provider-chosen |  | The Ip Discovery parameter for cachecluster. |
| `LogDeliveryConfigurations` | log_delivery_configurations | `list` | optional, computed, provider-chosen |  | Specifies the destination, format and type of the logs |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The network type parameter for cachecluster. |
| `NotificationTopicArn` | notification_topic_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent. |
| `NumCacheNodes` | num_cache_nodes | `integer` | required |  | The number of cache nodes that the cache cluster should have. |
| `Port` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The port number on which each of the cache nodes accepts connections. |
| `PreferredAvailabilityZone` | preferred_availability_zone | `string` | optional, computed, provider-chosen, write-only |  | The EC2 Availability Zone in which the cluster is created. |
| `PreferredAvailabilityZones` | preferred_availability_zones | `list` | optional, computed, provider-chosen |  | A list of the Availability Zones in which cache nodes are created. The order of the zones in the list is not important. |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | Specifies the weekly time range during which maintenance on the cluster is performed. |
| `RedisEndpoint` | redis_endpoint | `map` | computed |  | Specifies the RedisEndPoint address and port |
| `SnapshotArns` | snapshot_arns | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | A single-element string list containing an Amazon Resource Name (ARN) that uniquely identifies a Redis RDB snapshot file stored in Amazon S3. |
| `SnapshotName` | snapshot_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of a Redis snapshot from which to restore data into the new node group (shard). |
| `SnapshotRetentionLimit` | snapshot_retention_limit | `integer` | optional, computed, provider-chosen |  | The number of days for which ElastiCache retains automatic snapshots before deleting them. |
| `SnapshotWindow` | snapshot_window | `string` | optional, computed, provider-chosen |  | The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard). |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to be added to this resource. |
| `TransitEncryptionEnabled` | transit_encryption_enabled | `boolean` | optional, computed, provider-chosen |  | A flag that enables in-transit encryption when set to true. You cannot modify the value of TransitEncryptionEnabled after the cluster is created |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | One or more VPC security groups associated with the cluster. |

Supports update: yes

Discovery: supported
