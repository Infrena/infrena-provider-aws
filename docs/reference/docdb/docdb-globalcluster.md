# aws.docdb.globalcluster

**CloudFormation type:** `AWS::DocDB::GlobalCluster`

The AWS::DocDB::GlobalCluster resource represents an Amazon DocumentDB Global Cluster.

Region attribute: `region`

**Import ID:** `<region>/GlobalClusterIdentifier` (AWS::DocDB::GlobalCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | Indicates whether the global cluster has deletion protection enabled. The global cluster can't be deleted when deletion protection is enabled. |
| `Engine` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The database engine to use for this global cluster. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen, replaces on change |  | The engine version to use for this global cluster. |
| `GlobalClusterArn` | global_cluster_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the global cluster. |
| `GlobalClusterIdentifier` | global_cluster_identifier | `string` | required, replaces on change |  | The cluster identifier of the global cluster. |
| `GlobalClusterResourceId` | global_cluster_resource_id | `string` | computed |  | The AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed. |
| `SourceDBClusterIdentifier` | source_db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) to use as the primary cluster of the global cluster. You may also choose to instead specify the DBClusterIdentifier. If you provide a value for this parameter, don't specify values for the following settings because Amazon DocumentDB uses the values from the specified source DB cluster: Engine, EngineVersion, StorageEncrypted |
| `StorageEncrypted` | storage_encrypted | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the global cluster has storage encryption enabled. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to be assigned to the Amazon DocumentDB resource. |

Supports update: yes

Discovery: supported
