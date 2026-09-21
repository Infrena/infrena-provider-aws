# aws.rds.dbcluster

**CloudFormation type:** `AWS::RDS::DBCluster`

The ``AWS::RDS::DBCluster`` resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.

Region attribute: `region`

**Import ID:** `<region>/DBClusterIdentifier` (AWS::RDS::DBCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedStorage` | allocated_storage | `integer` | optional, computed, provider-chosen |  | The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster. |
| `AssociatedRoles` | associated_roles | `list` | optional, computed, provider-chosen |  | Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf. |
| `AutoMinorVersionUpgrade` | auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  | Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically. |
| `AvailabilityZones` | availability_zones | `list` | optional, computed, provider-chosen |  | A list of Availability Zones (AZs) where instances in the DB cluster can be created. For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide*. |
| `BacktrackWindow` | backtrack_window | `integer` | optional, computed, provider-chosen |  | The target backtrack window, in seconds. To disable backtracking, set this value to ``0``. |
| `BackupRetentionPeriod` | backup_retention_period | `integer` | optional, computed, provider-chosen |  | The number of days for which automated backups are retained. |
| `ClusterScalabilityType` | cluster_scalability_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the scalability mode of the Aurora DB cluster. When set to ``limitless``, the cluster operates as an Aurora Limitless Database, allowing you to create a DB shard group for horizontal scaling (sharding) capabilities. When set to ``standard`` (the default), the cluster uses normal DB instance creation. |
| `CopyTagsToSnapshot` | copy_tags_to_snapshot | `boolean` | optional, computed, provider-chosen |  | A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them. |
| `DBClusterArn` | db_cluster_arn | `string` | computed |  |  |
| `DBClusterIdentifier` | db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The DB cluster identifier. This parameter is stored as a lowercase string. |
| `DBClusterInstanceClass` | db_cluster_instance_class | `string` | optional, computed, provider-chosen |  | The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example ``db.m6gd.xlarge``. Not all DB instance classes are available in all AWS-Regions, or for all database engines. |
| `DBClusterParameterGroupName` | db_cluster_parameter_group_name | `string` | optional, computed, provider-chosen |  | The name of the DB cluster parameter group to associate with this DB cluster. |
| `DBClusterResourceId` | db_cluster_resource_id | `string` | computed |  |  |
| `DBInstanceParameterGroupName` | db_instance_parameter_group_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the DB parameter group to apply to all instances of the DB cluster. |
| `DBSubnetGroupName` | db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | A DB subnet group that you want to associate with this DB cluster. |
| `DBSystemId` | db_system_id | `string` | optional, computed, provider-chosen, replaces on change |  | Reserved for future use. |
| `DatabaseInsightsMode` | database_insights_mode | `string` | optional, computed, provider-chosen |  | The mode of Database Insights to enable for the DB cluster. |
| `DatabaseName` | database_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of your database. If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide*. |
| `DeleteAutomatedBackups` | delete_automated_backups | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to remove automated backups immediately after the DB cluster is deleted. This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB cluster is deleted, unless the AWS Backup policy specifies a point-in-time restore rule. |
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled. |
| `Domain` |  | `string` | optional, computed, provider-chosen |  | Indicates the directory ID of the Active Directory to create the DB cluster. |
| `DomainIAMRoleName` | domain_iam_role_name | `string` | optional, computed, provider-chosen |  | Specifies the name of the IAM role to use when making API calls to the Directory Service. |
| `EnableCloudwatchLogsExports` | enable_cloudwatch_logs_exports | `list` | optional, computed, provider-chosen |  | The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide*. |
| `EnableGlobalWriteForwarding` | enable_global_write_forwarding | `boolean` | optional, computed, provider-chosen |  | Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database. |
| `EnableHttpEndpoint` | enable_http_endpoint | `boolean` | optional, computed, provider-chosen |  | Specifies whether to enable the HTTP endpoint for the DB cluster. By default, the HTTP endpoint isn't enabled. |
| `EnableIAMDatabaseAuthentication` | enable_iam_database_authentication | `boolean` | optional, computed, provider-chosen |  | A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled. |
| `EnableLocalWriteForwarding` | enable_local_write_forwarding | `boolean` | optional, computed, provider-chosen |  | Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster. By default, write operations aren't allowed on reader DB instances. |
| `Endpoint` |  | `map` | computed |  | The ``Endpoint`` return value specifies the connection endpoint for the primary instance of the DB cluster. |
| `Engine` |  | `string` | optional, computed, provider-chosen |  | The name of the database engine to be used for this DB cluster. |
| `EngineLifecycleSupport` | engine_lifecycle_support | `string` | optional, computed, provider-chosen |  | The lifecycle type for this DB cluster. |
| `EngineMode` | engine_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The DB engine mode of the DB cluster, either ``provisioned`` or ``serverless``. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | The version number of the database engine to use. |
| `GlobalClusterIdentifier` | global_cluster_identifier | `string` | optional, computed, provider-chosen |  | If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster. To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource. |
| `Iops` |  | `integer` | optional, computed, provider-chosen |  | The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as ``arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef``. If you enable the ``StorageEncrypted`` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the ``StorageEncrypted`` property to ``true``. |
| `ManageMasterUserPassword` | manage_master_user_password | `boolean` | optional, computed, provider-chosen |  | Specifies whether to manage the master user password with AWS Secrets Manager. |
| `MasterUserAuthenticationType` | master_user_authentication_type | `string` | optional, computed, provider-chosen, write-only |  | Specifies the authentication type for the master user. With IAM master user authentication, you can configure the master DB user with IAM database authentication when you create a DB cluster. |
| `MasterUserPassword` | master_user_password | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The master password for the DB instance. |
| `MasterUserSecret` | master_user_secret | `map` | optional, computed, provider-chosen |  | The ``MasterUserSecret`` return value specifies the secret managed by RDS in AWS Secrets Manager for the master user password. |
| `MasterUsername` | master_username | `string` | optional, computed, provider-chosen |  | The name of the master user for the DB cluster. |
| `MonitoringInterval` | monitoring_interval | `integer` | optional, computed, provider-chosen |  | The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify ``0``. |
| `MonitoringRoleArn` | monitoring_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is ``arn:aws:iam:123456789012:role/emaccess``. For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide*. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen |  | The network type of the DB cluster. |
| `PerformanceInsightsEnabled` | performance_insights_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether to turn on Performance Insights for the DB cluster. |
| `PerformanceInsightsKmsKeyId` | performance_insights_kms_key_id | `string` | optional, computed, provider-chosen |  | The AWS KMS key identifier for encryption of Performance Insights data. |
| `PerformanceInsightsRetentionPeriod` | performance_insights_retention_period | `integer` | optional, computed, provider-chosen |  | The number of days to retain Performance Insights data. When creating a DB cluster without enabling Performance Insights, you can't specify the parameter ``PerformanceInsightsRetentionPeriod``. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port number on which the DB instances in the DB cluster accept connections. |
| `PreferredBackupWindow` | preferred_backup_window | `string` | optional, computed, provider-chosen |  | The daily time range during which automated backups are created. For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.* |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC). |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the DB cluster is publicly accessible. |
| `ReadEndpoint` | read_endpoint | `map` | computed |  | The ``ReadEndpoint`` return value specifies the reader endpoint for the DB cluster. |
| `ReplicationSourceIdentifier` | replication_source_identifier | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica. |
| `RestoreToTime` | restore_to_time | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The date and time to restore the DB cluster to. |
| `RestoreType` | restore_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The type of restore to be performed. You can specify one of the following values: |
| `ScalingConfiguration` | scaling_configuration | `map` | optional, computed, provider-chosen |  | The ``ScalingConfiguration`` property type specifies the scaling configuration of an Aurora Serverless v1 DB cluster. |
| `ServerlessV2ScalingConfiguration` | serverless_v2_scaling_configuration | `map` | optional, computed, provider-chosen |  | The ``ServerlessV2ScalingConfiguration`` property type specifies the scaling configuration of an Aurora Serverless V2 DB cluster. For more information, see [Using Amazon Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html) in the *Amazon Aurora User Guide*. |
| `SnapshotIdentifier` | snapshot_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier for the DB snapshot or DB cluster snapshot to restore from. |
| `SourceDBClusterIdentifier` | source_db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore. |
| `SourceDbClusterResourceId` | source_db_cluster_resource_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The resource ID of the source DB cluster from which to restore. |
| `SourceRegion` | source_region | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, ``us-east-1``. |
| `StorageEncrypted` | storage_encrypted | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the DB cluster is encrypted. |
| `StorageEncryptionType` | storage_encryption_type | `string` | computed |  |  |
| `StorageThroughput` | storage_throughput | `integer` | computed |  |  |
| `StorageType` | storage_type | `string` | optional, computed, provider-chosen |  | The storage type to associate with the DB cluster. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the DB cluster. |
| `UseLatestRestorableTime` | use_latest_restorable_time | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | A value that indicates whether to restore the DB cluster to the latest restorable backup time. By default, the DB cluster is not restored to the latest restorable backup time. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | A list of EC2 VPC security groups to associate with this DB cluster. |

Supports update: yes

Discovery: supported
