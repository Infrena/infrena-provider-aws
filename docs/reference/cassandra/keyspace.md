# aws.keyspace

**CloudFormation type:** `AWS::Cassandra::Keyspace`

Resource schema for AWS::Cassandra::Keyspace

Region attribute: `region`

**Import ID:** `<region>/KeyspaceName` (AWS::Cassandra::Keyspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClientSideTimestampsEnabled` | client_side_timestamps_enabled | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace. To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can’t disable it again. |
| `KeyspaceName` | keyspace_name | `string` | optional, computed, provider-chosen, replaces on change |  | Name for Cassandra keyspace |
| `ReplicationSpecification` | replication_specification | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
