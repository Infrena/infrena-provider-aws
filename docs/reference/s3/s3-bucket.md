# aws.s3.bucket

**CloudFormation type:** `AWS::S3::Bucket`

The ``AWS::S3::Bucket`` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.

Region attribute: `region`

**Import ID:** `<region>/BucketName` (AWS::S3::Bucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AbacStatus` | abac_status | `string` | optional, computed, provider-chosen |  | The ABAC status of the general purpose bucket. When ABAC is enabled for the general purpose bucket, you can use tags to manage access to the general purpose buckets as well as for cost tracking purposes. When ABAC is disabled for the general purpose buckets, you can only use tags for cost tracking purposes. For more information, see [Using tags with S3 general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging.html). |
| `AccelerateConfiguration` | accelerate_configuration | `map` | optional, computed, provider-chosen |  | Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*. |
| `AccessControl` | access_control | `string` | optional, computed, provider-chosen, write-only |  | This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*. |
| `AnalyticsConfigurations` | analytics_configurations | `list` | optional, computed, provider-chosen |  | Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket. |
| `Arn` |  | `string` | computed |  | the Amazon Resource Name (ARN) of the specified bucket. |
| `BucketEncryption` | bucket_encryption | `map` | optional, computed, provider-chosen |  | Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*. |
| `BucketName` | name, bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) in the *Amazon S3 User Guide*. |
| `BucketNamePrefix` | bucket_name_prefix | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `BucketNamespace` | bucket_namespace | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CorsConfiguration` | cors_configuration | `map` | optional, computed, provider-chosen |  | Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*. |
| `DomainName` | domain_name | `string` | computed |  |  |
| `DualStackDomainName` | dual_stack_domain_name | `string` | computed |  |  |
| `IntelligentTieringConfigurations` | intelligent_tiering_configurations | `list` | optional, computed, provider-chosen |  | Defines how Amazon S3 handles Intelligent-Tiering storage. |
| `InventoryConfigurations` | inventory_configurations | `list` | optional, computed, provider-chosen |  | Specifies the S3 Inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*. |
| `LifecycleConfiguration` | lifecycle_configuration | `map` | optional, computed, provider-chosen |  | Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*. |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  | Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For examples and more information, see [PUT Bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the *Amazon S3 API Reference*. |
| `MetadataConfiguration` | metadata_configuration | `map` | optional, computed, provider-chosen |  | Creates a V2 S3 Metadata configuration of a general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the *Amazon S3 User Guide*. |
| `MetadataTableConfiguration` | metadata_table_configuration | `map` | optional, computed, provider-chosen |  | We recommend that you create your S3 Metadata configurations by using the V2 [MetadataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadataconfiguration.html) resource type. We no longer recommend using the V1 ``MetadataTableConfiguration`` resource type. |
| `MetricsConfigurations` | metrics_configurations | `list` | optional, computed, provider-chosen |  | Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html). |
| `NotificationConfiguration` | notification_configuration | `map` | optional, computed, provider-chosen |  | Describes the notification configuration for an Amazon S3 bucket. |
| `ObjectLockConfiguration` | object_lock_configuration | `map` | optional, computed, provider-chosen |  | Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). |
| `ObjectLockEnabled` | object_lock_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket. |
| `OwnershipControls` | ownership_controls | `map` | optional, computed, provider-chosen |  | Specifies the container element for Object Ownership rules. |
| `PublicAccessBlockConfiguration` | public_access_block_configuration | `map` | optional, computed, provider-chosen |  | The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Bucket-level settings work alongside account-level settings (which may inherit from organization-level policies). For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide*. |
| `RegionalDomainName` | regional_domain_name | `string` | computed |  |  |
| `ReplicationConfiguration` | replication_configuration | `map` | optional, computed, provider-chosen |  | A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB. The latest version of the replication configuration XML is V2. For more information about XML V2 replication configurations, see [Replication configuration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication-add-config.html) in the *Amazon S3 User Guide*. |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this S3 bucket. |
| `VersioningConfiguration` | versioning_configuration | `map` | optional, computed, provider-chosen |  | Describes the versioning state of an Amazon S3 bucket. For more information, see [PUT Bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the *Amazon S3 API Reference*. |
| `WebsiteConfiguration` | website_configuration | `map` | optional, computed, provider-chosen |  | Specifies website configuration parameters for an Amazon S3 bucket. |
| `WebsiteURL` | website_url | `string` | computed |  |  |

Supports update: yes

Discovery: supported
