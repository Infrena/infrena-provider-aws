# aws.dbsnapshot

**CloudFormation type:** `AWS::RDS::DBSnapshot`

Creates a snapshot of a DB instance.

Region attribute: `region`

**Import ID:** `<region>/DBSnapshotArn` (AWS::RDS::DBSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedStorage` | allocated_storage | `integer` | computed |  | The allocated storage size in gibibytes (GiB). |
| `AvailabilityZone` | availability_zone | `string` | computed |  | The name of the Availability Zone the DB instance was located in at the time of the DB snapshot. |
| `DBInstanceIdentifier` | db_instance_identifier | `string` | required, replaces on change |  | The identifier of the DB instance that you want to create the snapshot of. |
| `DBSnapshotArn` | db_snapshot_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the DB snapshot. |
| `DBSnapshotIdentifier` | db_snapshot_identifier | `string` | required, replaces on change |  | The identifier for the DB snapshot. Must contain from 1 to 255 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens. |
| `DbiResourceId` | dbi_resource_id | `string` | computed |  | The identifier for the source DB instance, which is unique to an AWS Region. |
| `Encrypted` |  | `boolean` | computed |  | Indicates whether the DB snapshot is encrypted. |
| `Engine` |  | `string` | computed |  | The name of the database engine. |
| `EngineVersion` | engine_version | `string` | computed |  | The version of the database engine. |
| `IAMDatabaseAuthenticationEnabled` | iam_database_authentication_enabled | `boolean` | computed |  | Indicates whether mapping of AWS IAM accounts to database accounts is enabled. |
| `InstanceCreateTime` | instance_create_time | `string` | computed |  | The time when the DB instance was created, in UTC. |
| `Iops` |  | `integer` | computed |  | The Provisioned IOPS value of the DB instance at the time of the snapshot. |
| `KmsKeyId` | kms_key_id | `string` | computed |  | If Encrypted is true, the AWS KMS key identifier for the encrypted DB snapshot. |
| `LicenseModel` | license_model | `string` | computed |  | License model information for the restored DB instance. |
| `MasterUsername` | master_username | `string` | computed |  | The master username for the DB snapshot. |
| `OptionGroupName` | option_group_name | `string` | computed |  | The option group name for the DB snapshot. |
| `OriginalSnapshotCreateTime` | original_snapshot_create_time | `string` | computed |  | The time of the CreateDBSnapshot operation in UTC. Doesn't change when the snapshot is copied. |
| `Port` |  | `integer` | computed |  | The port that the database engine was listening on at the time of the snapshot. |
| `SnapshotCreateTime` | snapshot_create_time | `string` | computed |  | The time when the snapshot was taken, in UTC. |
| `SnapshotType` | snapshot_type | `string` | computed |  | The type of the DB snapshot. |
| `Status` |  | `string` | computed |  | The status of this DB snapshot. |
| `StorageThroughput` | storage_throughput | `integer` | computed |  | The storage throughput for the DB snapshot. |
| `StorageType` | storage_type | `string` | computed |  | The storage type associated with the DB snapshot. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to be assigned to the DB snapshot. |
| `VpcId` | vpc_id | `string` | computed |  | The VPC ID associated with the DB snapshot. |

Supports update: yes

Discovery: supported
