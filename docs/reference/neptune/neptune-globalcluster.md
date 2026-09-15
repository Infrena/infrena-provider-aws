# aws.neptune.globalcluster

**CloudFormation type:** `AWS::Neptune::GlobalCluster`

Resource Type definition for AWS::Neptune::GlobalCluster

Region attribute: `region`

**Import ID:** `<region>/GlobalClusterIdentifier` (AWS::Neptune::GlobalCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | Whether deletion protection is enabled. |
| `Engine` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the database engine. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | The version number of the database engine. |
| `GlobalClusterIdentifier` | global_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The cluster identifier of the global database cluster. |
| `SourceDBClusterIdentifier` | source_db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of an existing Neptune DB cluster to use as the primary cluster of the new global database. |
| `StorageEncrypted` | storage_encrypted | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether the global database cluster is storage encrypted. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
