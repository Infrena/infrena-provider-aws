# aws.neptune.dbcluster

**CloudFormation type:** `AWS::Neptune::DBCluster`

The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.

Region attribute: `region`

**Import ID:** `<region>/DBClusterIdentifier` (AWS::Neptune::DBCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatedRoles` | associated_roles | `list` | optional, computed, provider-chosen |  | Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. |
| `AvailabilityZones` | availability_zones | `list` | optional, computed, provider-chosen, replaces on change |  | Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in. |
| `BackupRetentionPeriod` | backup_retention_period | `integer` | optional, computed, provider-chosen |  | Specifies the number of days for which automatic DB snapshots are retained. |
| `ClusterResourceId` | cluster_resource_id | `string` | computed |  | The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies. |
| `CopyTagsToSnapshot` | copy_tags_to_snapshot | `boolean` | optional, computed, provider-chosen |  | A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them. |
| `DBClusterIdentifier` | db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string. |
| `DBClusterParameterGroupName` | db_cluster_parameter_group_name | `string` | optional, computed, provider-chosen |  | Provides the name of the DB cluster parameter group. |
| `DBInstanceParameterGroupName` | db_instance_parameter_group_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the DB parameter group to apply to all instances of the DB cluster. Used only in case of a major EngineVersion upgrade request. |
| `DBPort` | db_port | `integer` | optional, computed, provider-chosen |  | The port number on which the DB instances in the DB cluster accept connections. |
| `DBSubnetGroupName` | db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group. |
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. |
| `EnableCloudwatchLogsExports` | enable_cloudwatch_logs_exports | `list` | optional, computed, provider-chosen |  | Specifies a list of log types that are enabled for export to CloudWatch Logs. |
| `Endpoint` |  | `string` | computed |  | The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com` |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | Indicates the database engine version. |
| `GlobalClusterIdentifier` | global_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the Neptune global database to which this new DB cluster should be added. |
| `IamAuthEnabled` | iam_auth_enabled | `boolean` | optional, computed, provider-chosen |  | True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef. If you enable the StorageEncrypted property but don't specify this property, the default KMS key is used. If you specify this property, you must set the StorageEncrypted property to true. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen |  | The network type of the DB cluster. |
| `Port` |  | `string` | computed |  | The port number on which the DB cluster accepts connections. For example: `8182`. |
| `PreferredBackupWindow` | preferred_backup_window | `string` | optional, computed, provider-chosen |  | Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod. |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC). |
| `ReadEndpoint` | read_endpoint | `string` | computed |  | The reader endpoint for the DB cluster. For example: `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com` |
| `RestoreToTime` | restore_to_time | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Creates a new DB cluster from a DB snapshot or DB cluster snapshot. |
| `RestoreType` | restore_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Creates a new DB cluster from a DB snapshot or DB cluster snapshot. |
| `ServerlessScalingConfiguration` | serverless_scaling_configuration | `map` | optional, computed, provider-chosen |  | Contains the scaling configuration of an Neptune Serverless DB cluster. |
| `SnapshotIdentifier` | snapshot_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot. |
| `SourceDBClusterIdentifier` | source_db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Creates a new DB cluster from a DB snapshot or DB cluster snapshot. |
| `StorageEncrypted` | storage_encrypted | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the DB cluster is encrypted. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to this cluster. |
| `UseLatestRestorableTime` | use_latest_restorable_time | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Creates a new DB cluster from a DB snapshot or DB cluster snapshot. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | Provides a list of VPC security groups that the DB cluster belongs to. |

Supports update: yes

Discovery: supported
