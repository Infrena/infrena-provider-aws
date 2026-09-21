# aws.ec2.volume

**CloudFormation type:** `AWS::EC2::Volume`

Specifies an Amazon Elastic Block Store (Amazon EBS) volume. You can create an empty volume, a volume from a snapshot, or a volume copy from an existing source volume.

Region attribute: `region`

**Import ID:** `<region>/VolumeId` (AWS::EC2::Volume)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoEnableIO` | auto_enable_io | `boolean` | optional, computed, provider-chosen |  | Indicates whether the volume is auto-enabled for I/O operations. By default, EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O. |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen |  | The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``. |
| `AvailabilityZoneId` | availability_zone_id | `string` | optional, computed, provider-chosen |  | The ID of the Availability Zone in which to create the volume. For example, ``use1-az1``. |
| `Encrypted` |  | `boolean` | optional, computed, provider-chosen |  | Indicates whether the volume should be encrypted. The effect of setting the encryption state to ``true`` depends on the volume origin (new, from a snapshot, or from an existing volume), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*. |
| `Iops` |  | `integer` | optional, computed, provider-chosen |  | The number of I/O operations per second (IOPS) to provision for the volume. Required for ``io1`` and ``io2`` volumes. Optional for ``gp3`` volumes. Omit for all other volume types. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The identifier of the kms-key-long to use for Amazon EBS encryption. If ``KmsKeyId`` is specified, the encrypted state must be ``true``. |
| `MultiAttachEnabled` | multi_attach_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether Amazon EBS Multi-Attach is enabled. |
| `OutpostArn` | outpost_arn | `string` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Name (ARN) of the Outpost on which to create the volume. |
| `Size` |  | `integer` | optional, computed, provider-chosen |  | The size of the volume, in GiBs. |
| `SnapshotId` | snapshot_id | `string` | optional, computed, provider-chosen |  | The snapshot from which to create the volume. Only specify to create a volume from a snapshot. To create a new empty volume, omit this parameter and specify a value for ``Size`` instead. To create a volume copy, omit this parameter and specify ``SourceVolumeId`` instead. |
| `SourceVolumeId` | source_volume_id | `string` | optional, computed, provider-chosen | aws.ec2.volume.VolumeId | The ID of the source EBS volume to copy. When specified, the volume is created as an exact copy of the specified volume. Only specify to create a volume copy. To create a new empty volume or to create a volume from a snapshot, omit this parameter, |
| `Tags` |  | `map` | tags map |  | The tags to apply to the volume during creation. |
| `Throughput` |  | `integer` | optional, computed, provider-chosen |  | The throughput to provision for a volume, with a maximum of 2,000 MiB/s. |
| `VolumeId` | volume_id | `string` | computed |  |  |
| `VolumeInitializationRate` | volume_initialization_rate | `integer` | optional, computed, provider-chosen |  | Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download the snapshot blocks from Amazon S3 to the volume. This is also known as *volume initialization*. Specifying a volume initialization rate ensures that the volume is initialized at a predictable and consistent rate after creation. |
| `VolumeType` | volume_type | `string` | optional, computed, provider-chosen |  | The volume type. This parameter can be one of the following values: |

Supports update: yes

Discovery: supported
