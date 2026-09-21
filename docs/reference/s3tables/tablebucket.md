# aws.tablebucket

**CloudFormation type:** `AWS::S3Tables::TableBucket`

Creates an Amazon S3 Tables table bucket in the same AWS Region where you create the AWS CloudFormation stack.

Region attribute: `region`

**Import ID:** `<region>/TableBucketARN` (AWS::S3Tables::TableBucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen |  | Specifies encryption settings for the table bucket |
| `MetricsConfiguration` | metrics_configuration | `map` | optional, computed, provider-chosen |  | Settings governing the Metric configuration for the table bucket. |
| `ReplicationConfiguration` | replication_configuration | `map` | optional, computed, provider-chosen |  | Specifies replication configuration for the table bucket |
| `StorageClassConfiguration` | storage_class_configuration | `map` | optional, computed, provider-chosen |  | Specifies storage class settings for the table bucket |
| `TableBucketARN` | table_bucket_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified table bucket. |
| `TableBucketName` | table_bucket_name | `string` | required, replaces on change |  | A name for the table bucket. |
| `Tags` |  | `map` | tags map |  | User tags (key-value pairs) to associate with the table bucket. |
| `UnreferencedFileRemoval` | unreferenced_file_removal | `map` | optional, computed, provider-chosen |  | Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. |

Supports update: yes

Discovery: supported
