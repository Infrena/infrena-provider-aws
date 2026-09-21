# aws.clustersnapshot

**CloudFormation type:** `AWS::RDS::ClusterSnapshot`

Creates a snapshot of a DB cluster.

Region attribute: `region`

**Import ID:** `<region>/DBClusterSnapshotArn` (AWS::RDS::ClusterSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedStorage` | allocated_storage | `integer` | computed |  | The allocated storage size of the DB cluster snapshot in gibibytes (GiB). |
| `AvailabilityZones` | availability_zones | `list` | computed |  | The list of Availability Zones where instances in the DB cluster snapshot can be restored. |
| `ClusterCreateTime` | cluster_create_time | `string` | computed |  | The time when the DB cluster was created, in Universal Coordinated Time (UTC). |
| `DBClusterIdentifier` | db_cluster_identifier | `string` | required, replaces on change |  | The identifier of the DB cluster to create a snapshot for. |
| `DBClusterSnapshotArn` | db_cluster_snapshot_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the DB cluster snapshot. |
| `DBClusterSnapshotIdentifier` | db_cluster_snapshot_identifier | `string` | required, replaces on change |  | The identifier for the DB cluster snapshot. Must contain from 1 to 63 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens. |
| `DbClusterResourceId` | db_cluster_resource_id | `string` | computed |  | The resource ID of the DB cluster that this DB cluster snapshot was created from. |
| `Engine` |  | `string` | computed |  | The name of the database engine for this DB cluster snapshot. |
| `EngineMode` | engine_mode | `string` | computed |  | The engine mode of the database engine for this DB cluster snapshot. |
| `EngineVersion` | engine_version | `string` | computed |  | The version of the database engine for this DB cluster snapshot. |
| `IAMDatabaseAuthenticationEnabled` | iam_database_authentication_enabled | `boolean` | computed |  | Indicates whether mapping of AWS IAM accounts to database accounts is enabled. |
| `KmsKeyId` | kms_key_id | `string` | computed |  | If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot. |
| `LicenseModel` | license_model | `string` | computed |  | The license model information for this DB cluster snapshot. |
| `MasterUsername` | master_username | `string` | computed |  | The master username for this DB cluster snapshot. |
| `Port` |  | `integer` | computed |  | The port that the DB cluster was listening on at the time of the snapshot. |
| `SnapshotCreateTime` | snapshot_create_time | `string` | computed |  | The time when the snapshot was taken, in Universal Coordinated Time (UTC). |
| `SnapshotType` | snapshot_type | `string` | computed |  | The type of the DB cluster snapshot. |
| `Status` |  | `string` | computed |  | The status of this DB cluster snapshot. |
| `StorageEncrypted` | storage_encrypted | `boolean` | computed |  | Indicates whether the DB cluster snapshot is encrypted. |
| `Tags` |  | `map` | tags map |  | The tags to be assigned to the DB cluster snapshot. |
| `VpcId` | vpc_id | `string` | computed |  | The VPC ID associated with the DB cluster snapshot. |

Supports update: yes

Discovery: supported
