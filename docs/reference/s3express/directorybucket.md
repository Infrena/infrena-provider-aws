# aws.directorybucket

**CloudFormation type:** `AWS::S3Express::DirectoryBucket`

Resource Type definition for AWS::S3Express::DirectoryBucket.

Region attribute: `region`

**Import ID:** `<region>/BucketName` (AWS::S3Express::DirectoryBucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified bucket. |
| `AvailabilityZoneName` | availability_zone_name | `string` | computed |  | Returns the code for the Availability Zone or Local Zone where the directory bucket was created. An example for the code of an Availability Zone is 'us-east-1f'. |
| `BucketEncryption` | bucket_encryption | `map` | optional, computed, provider-chosen |  | Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS). |
| `BucketName` | bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone or Local Zone. The bucket name must also follow the format 'bucket_base_name--zone_id--x-s3'. The zone_id can be the ID of an Availability Zone or a Local Zone. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name. |
| `DataRedundancy` | data_redundancy | `string` | required, replaces on change |  | Specifies the number of Availability Zone or Local Zone that's used for redundancy for the bucket. |
| `InventoryConfigurations` | inventory_configurations | `list` | optional, computed, provider-chosen |  | The inventory configuration for an Amazon S3 Express bucket. |
| `LifecycleConfiguration` | lifecycle_configuration | `map` | optional, computed, provider-chosen |  | Lifecycle rules that define how Amazon S3 Express manages objects during their lifetime. |
| `LocationName` | location_name | `string` | required, replaces on change |  | Specifies the Zone ID of the Availability Zone or Local Zone where the directory bucket will be created. An example Availability Zone ID value is 'use1-az5'. |
| `MetricsConfigurations` | metrics_configurations | `list` | optional, computed, provider-chosen |  | Specifies the metrics configurations for the Amazon S3 Express bucket. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
