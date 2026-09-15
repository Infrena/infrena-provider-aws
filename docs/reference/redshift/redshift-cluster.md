# aws.redshift.cluster

**CloudFormation type:** `AWS::Redshift::Cluster`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/ClusterIdentifier` (AWS::Redshift::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowVersionUpgrade` | allow_version_upgrade | `boolean` | optional, computed, provider-chosen |  | Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True |
| `AquaConfigurationStatus` | aqua_configuration_status | `string` | optional, computed, provider-chosen |  | The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following. |
| `AutomatedSnapshotRetentionPeriod` | automated_snapshot_retention_period | `integer` | optional, computed, provider-chosen |  | The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1 |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen |  | The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint |
| `AvailabilityZoneRelocation` | availability_zone_relocation | `boolean` | optional, computed, provider-chosen |  | The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete. |
| `AvailabilityZoneRelocationStatus` | availability_zone_relocation_status | `string` | optional, computed, provider-chosen |  | The availability zone relocation status of the cluster |
| `Classic` |  | `boolean` | optional, computed, provider-chosen, write-only |  | A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic. |
| `ClusterIdentifier` | cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account |
| `ClusterNamespaceArn` | cluster_namespace_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the cluster namespace. |
| `ClusterParameterGroupName` | cluster_parameter_group_name | `string` | optional, computed, provider-chosen |  | The name of the parameter group to be associated with this cluster. |
| `ClusterSecurityGroups` | cluster_security_groups | `list` | optional, computed, provider-chosen |  | A list of security groups to be associated with this cluster. |
| `ClusterSubnetGroupName` | cluster_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of a cluster subnet group to be associated with this cluster. |
| `ClusterType` | cluster_type | `string` | required |  | The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required |
| `ClusterVersion` | cluster_version | `string` | optional, computed, provider-chosen |  | The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster. |
| `DBName` | db_name | `string` | required, replaces on change |  | The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database. |
| `DeferMaintenance` | defer_maintenance | `boolean` | optional, computed, provider-chosen, write-only |  | A boolean indicating whether to enable the deferred maintenance window. |
| `DeferMaintenanceDuration` | defer_maintenance_duration | `integer` | optional, computed, provider-chosen, write-only |  | An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 60 days or less. |
| `DeferMaintenanceEndTime` | defer_maintenance_end_time | `string` | optional, computed, provider-chosen |  | A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration. |
| `DeferMaintenanceIdentifier` | defer_maintenance_identifier | `string` | computed |  | A unique identifier for the deferred maintenance window. |
| `DeferMaintenanceStartTime` | defer_maintenance_start_time | `string` | optional, computed, provider-chosen |  | A timestamp indicating the start time for the deferred maintenance window. |
| `DestinationRegion` | destination_region | `string` | optional, computed, provider-chosen |  | The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference |
| `ElasticIp` | elastic_ip | `string` | optional, computed, provider-chosen |  | The Elastic IP (EIP) address for the cluster. |
| `Encrypted` |  | `boolean` | optional, computed, provider-chosen |  | If true, the data in the cluster is encrypted at rest. |
| `Endpoint` |  | `map` | optional, computed, provider-chosen |  |  |
| `EnhancedVpcRouting` | enhanced_vpc_routing | `boolean` | optional, computed, provider-chosen |  | An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide. |
| `HsmClientCertificateIdentifier` | hsm_client_certificate_identifier | `string` | optional, computed, provider-chosen |  | Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM |
| `HsmConfigurationIdentifier` | hsm_configuration_identifier | `string` | optional, computed, provider-chosen |  | Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM. |
| `IamRoles` | iam_roles | `list` | optional, computed, provider-chosen |  | A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster. |
| `LoggingProperties` | logging_properties | `map` | optional, computed, provider-chosen |  |  |
| `MaintenanceTrackName` | maintenance_track_name | `string` | optional, computed, provider-chosen |  | The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied. |
| `ManageMasterPassword` | manage_master_password | `boolean` | optional, computed, provider-chosen, write-only |  | A boolean indicating if the redshift cluster's admin user credentials is managed by Redshift or not. You can't use MasterUserPassword if ManageMasterPassword is true. If ManageMasterPassword is false or not set, Amazon Redshift uses MasterUserPassword for the admin user account's password. |
| `ManualSnapshotRetentionPeriod` | manual_snapshot_retention_period | `integer` | optional, computed, provider-chosen |  | The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely. |
| `MasterPasswordSecretArn` | master_password_secret_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the cluster's admin user credentials secret. |
| `MasterPasswordSecretKmsKeyId` | master_password_secret_kms_key_id | `string` | optional, computed, provider-chosen |  | The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret. |
| `MasterUserPassword` | master_user_password | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password associated with the master user account for the cluster that is being created. You can't use MasterUserPassword if ManageMasterPassword is true. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character. |
| `MasterUsername` | master_username | `string` | required, replaces on change |  | The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter. |
| `MultiAZ` | multi_az | `boolean` | optional, computed, provider-chosen |  | A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az. |
| `NamespaceResourcePolicy` | namespace_resource_policy | `map` | optional, computed, provider-chosen |  | The namespace resource policy document that will be attached to a Redshift cluster. |
| `NodeType` | node_type | `string` | required |  | The node type to be provisioned for the cluster.Valid Values: ds2.xlarge \| ds2.8xlarge \| dc1.large \| dc1.8xlarge \| dc2.large \| dc2.8xlarge \| ra3.large \| ra3.4xlarge \| ra3.16xlarge \| rg.large \| rg.xlarge \| rg.4xlarge \| rg.12xlarge |
| `NumberOfNodes` | number_of_nodes | `integer` | optional, computed, provider-chosen |  | The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. |
| `OwnerAccount` | owner_account | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | The weekly time range (in UTC) during which automated cluster maintenance can occur. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen |  | If true, the cluster can be accessed from a public network. |
| `ResourceAction` | resource_action | `string` | optional, computed, provider-chosen |  | The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs |
| `RevisionTarget` | revision_target | `string` | optional, computed, provider-chosen |  | The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request. |
| `RotateEncryptionKey` | rotate_encryption_key | `boolean` | optional, computed, provider-chosen |  | A boolean indicating if we want to rotate Encryption Keys. |
| `SnapshotClusterIdentifier` | snapshot_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name. |
| `SnapshotCopyGrantName` | snapshot_copy_grant_name | `string` | optional, computed, provider-chosen |  | The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region. |
| `SnapshotCopyManual` | snapshot_copy_manual | `boolean` | optional, computed, provider-chosen |  | Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots. |
| `SnapshotCopyRetentionPeriod` | snapshot_copy_retention_period | `integer` | optional, computed, provider-chosen |  | The number of days to retain automated snapshots in the destination region after they are copied from the source region. |
| `SnapshotIdentifier` | snapshot_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The list of tags for the cluster parameter group. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster. |

Supports update: yes

Discovery: supported
