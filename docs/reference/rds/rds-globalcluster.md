# aws.rds.globalcluster

**CloudFormation type:** `AWS::RDS::GlobalCluster`

Resource Type definition for AWS::RDS::GlobalCluster

Region attribute: `region`

**Import ID:** `<region>/GlobalClusterIdentifier` (AWS::RDS::GlobalCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled. |
| `Engine` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora). |
| `EngineLifecycleSupport` | engine_lifecycle_support | `string` | optional, computed, provider-chosen |  | The life cycle type of the global cluster. You can use this setting to enroll your global cluster into Amazon RDS Extended Support. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster. |
| `GlobalClusterIdentifier` | global_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string. |
| `GlobalEndpoint` | global_endpoint | `map` | computed |  |  |
| `SourceDBClusterIdentifier` | source_db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string. |
| `StorageEncrypted` | storage_encrypted | `boolean` | optional, computed, provider-chosen, replaces on change |  | The storage encryption setting for the new global database cluster. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
