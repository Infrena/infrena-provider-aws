# aws.dbshardgroup

**CloudFormation type:** `AWS::RDS::DBShardGroup`

Creates a new DB shard group for Aurora Limitless Database. You must enable Aurora Limitless Database to create a DB shard group.

Region attribute: `region`

**Import ID:** `<region>/DBShardGroupIdentifier` (AWS::RDS::DBShardGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputeRedundancy` | compute_redundancy | `integer` | optional, computed, provider-chosen |  | Specifies whether to create standby standby DB data access shard for the DB shard group. Valid values are the following: |
| `DBClusterIdentifier` | db_cluster_identifier | `string` | required, replaces on change |  | The name of the primary DB cluster for the DB shard group. |
| `DBShardGroupIdentifier` | db_shard_group_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the DB shard group. |
| `DBShardGroupResourceId` | db_shard_group_resource_id | `string` | computed |  |  |
| `Endpoint` |  | `string` | computed |  |  |
| `MaxACU` | max_acu | `float` | required |  | The maximum capacity of the DB shard group in Aurora capacity units (ACUs). |
| `MinACU` | min_acu | `float` | optional, computed, provider-chosen, write-only |  | The minimum capacity of the DB shard group in Aurora capacity units (ACUs). |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the DB shard group is publicly accessible. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group. |

Supports update: yes

Discovery: supported
