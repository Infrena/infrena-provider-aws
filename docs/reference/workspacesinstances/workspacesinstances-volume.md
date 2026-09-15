# aws.workspacesinstances.volume

**CloudFormation type:** `AWS::WorkspacesInstances::Volume`

Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources

Region attribute: `region`

**Import ID:** `<region>/VolumeId` (AWS::WorkspacesInstances::Volume)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | required, replaces on change |  | The Availability Zone in which to create the volume |
| `Encrypted` |  | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the volume should be encrypted |
| `Iops` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of I/O operations per second (IOPS) |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption |
| `SizeInGB` | size_in_gb | `integer` | optional, computed, provider-chosen, replaces on change |  | The size of the volume, in GiBs |
| `SnapshotId` | snapshot_id | `string` | optional, computed, provider-chosen, replaces on change |  | The snapshot from which to create the volume |
| `TagSpecifications` | tag_specifications | `list` | optional, computed, provider-chosen, replaces on change |  | The tags passed to EBS volume |
| `Throughput` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The throughput to provision for a volume, with a maximum of 1,000 MiB/s |
| `VolumeId` | volume_id | `string` | computed |  | Unique identifier for the volume |
| `VolumeType` | volume_type | `string` | optional, computed, provider-chosen, replaces on change |  | The volume type |

Supports update: no

Discovery: supported
