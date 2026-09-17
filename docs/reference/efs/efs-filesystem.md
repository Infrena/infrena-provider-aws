# aws.efs.filesystem

**CloudFormation type:** `AWS::EFS::FileSystem`

The ``AWS::EFS::FileSystem`` resource creates a new, empty file system in EFSlong (EFS). You must create a mount target ([AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an EC2 or other AWS cloud compute resource.

Region attribute: `region`

**Import ID:** `<region>/FileSystemId` (AWS::EFS::FileSystem)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AvailabilityZoneName` | availability_zone_name | `string` | optional, computed, provider-chosen, replaces on change |  | For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*. |
| `BackupPolicy` | backup_policy | `map` | optional, computed, provider-chosen |  | The backup policy turns automatic backups for the file system on or off. |
| `BypassPolicyLockoutSafetyCheck` | bypass_policy_lockout_safety_check | `boolean` | optional, computed, provider-chosen, write-only |  | (Optional) A boolean that specifies whether or not to bypass the ``FileSystemPolicy`` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ``PutFileSystemPolicy`` requests on this file system. Set ``BypassPolicyLockoutSafetyCheck`` to ``True`` only when you intend to prevent the IAM principal that is making the request from making subsequent ``PutFileSystemPolicy`` requests on this file system. The default value is ``False``. |
| `Encrypted` |  | `boolean` | optional, computed, provider-chosen, replaces on change |  | A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ``/aws/elasticfilesystem``, is used to protect the encrypted file system. |
| `FileSystemId` | file_system_id | `string` | computed |  |  |
| `FileSystemPolicy` | file_system_policy | `map` | optional, computed, provider-chosen |  | The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*. |
| `FileSystemProtection` | file_system_protection | `map` | optional, computed, provider-chosen |  | Describes the protection on the file system. |
| `FileSystemTags` | file_system_tags | `map` | optional, computed, provider-chosen, tags map |  | Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``"Key":"Name","Value":"{value}"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats: |
| `LifecyclePolicies` | lifecycle_policies | `list` | optional, computed, provider-chosen |  | An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following: |
| `PerformanceMode` | performance_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The performance mode of the file system. We recommend ``generalPurpose`` performance mode for all file systems. File systems using the ``maxIO`` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ``maxIO`` mode is not supported on One Zone file systems. |
| `ProvisionedThroughputInMibps` | provisioned_throughput_in_mibps | `float` | optional, computed, provider-chosen |  | The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*. |
| `ReplicationConfiguration` | replication_configuration | `map` | optional, computed, provider-chosen, write-only |  | Describes the replication configuration for a specific file system. |
| `ThroughputMode` | throughput_mode | `string` | optional, computed, provider-chosen |  | Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. |

Supports update: yes

Discovery: supported
