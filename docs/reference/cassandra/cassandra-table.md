# aws.cassandra.table

**CloudFormation type:** `AWS::Cassandra::Table`

Resource schema for AWS::Cassandra::Table

Region attribute: `region`

**Import ID:** `<region>/KeyspaceName|TableName` (AWS::Cassandra::Table)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingSpecifications` | auto_scaling_specifications | `map` | optional, computed, provider-chosen, write-only |  | Represents the read and write settings used for AutoScaling. |
| `BillingMode` | billing_mode | `map` | optional, computed, provider-chosen |  |  |
| `CdcSpecification` | cdc_specification | `map` | optional, computed, provider-chosen |  | Represents the CDC configuration for the table |
| `ClientSideTimestampsEnabled` | client_side_timestamps_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again. |
| `ClusteringKeyColumns` | clustering_key_columns | `list` | optional, computed, provider-chosen, replaces on change |  | Clustering key columns of the table |
| `DefaultTimeToLive` | default_time_to_live | `integer` | optional, computed, provider-chosen |  | Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column. |
| `EncryptionSpecification` | encryption_specification | `map` | optional, computed, provider-chosen |  | Represents the settings used to enable server-side encryption |
| `KeyspaceName` | keyspace_name | `string` | required, replaces on change |  | Name for Cassandra keyspace |
| `PartitionKeyColumns` | partition_key_columns | `list` | required, replaces on change |  | Partition key columns of the table |
| `PointInTimeRecoveryEnabled` | point_in_time_recovery_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether point in time recovery is enabled (true) or disabled (false) on the table |
| `RegularColumns` | regular_columns | `list` | optional, computed, provider-chosen |  | Non-key columns of the table |
| `ReplicaSpecifications` | replica_specifications | `list` | optional, computed, provider-chosen, write-only |  |  |
| `TableName` | table_name | `string` | optional, computed, provider-chosen, replaces on change |  | Name for Cassandra table |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource |
| `WarmThroughput` | warm_throughput | `map` | optional, computed, provider-chosen |  | Warm throughput configuration for the table |

Supports update: yes

Discovery: supported
