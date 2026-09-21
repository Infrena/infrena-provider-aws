# aws.customdbengineversion

**CloudFormation type:** `AWS::RDS::CustomDBEngineVersion`

Creates a custom DB engine version (CEV).

Region attribute: `region`

**Import ID:** `<region>/Engine|EngineVersion` (AWS::RDS::CustomDBEngineVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBEngineVersionArn` | db_engine_version_arn | `string` | computed |  |  |
| `DatabaseInstallationFiles` | database_installation_files | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `DatabaseInstallationFilesS3BucketName` | database_installation_files_s3_bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is ``my-custom-installation-files``. |
| `DatabaseInstallationFilesS3Prefix` | database_installation_files_s3_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is ``123456789012/cev1``. If this setting isn't specified, no prefix is assumed. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description of your CEV. |
| `Engine` |  | `string` | required, replaces on change |  | The database engine to use for your custom engine version (CEV). |
| `EngineVersion` | engine_version | `string` | required, replaces on change |  | The name of your CEV. The name format is ``major version.customized_string``. For example, a valid CEV name is ``19.my_cev1``. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of ``Engine`` and ``EngineVersion`` is unique per customer per Region. |
| `ImageId` | image_id | `string` | optional, computed, provider-chosen, replaces on change |  | A value that indicates the ID of the AMI. |
| `KMSKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS KMS key identifier for an encrypted CEV. A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS. |
| `Manifest` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed. |
| `SourceCustomDbEngineVersionIdentifier` | source_custom_db_engine_version_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN of a CEV to use as a source for creating a new CEV. You can specify a different Amazon Machine Imagine (AMI) by using either ``Source`` or ``UseAwsProvidedLatestImage``. You can't specify a different JSON manifest when you specify ``SourceCustomDbEngineVersionIdentifier``. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | A value that indicates the status of a custom engine version (CEV). |
| `Tags` |  | `map` | tags map |  | A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.* |
| `UseAwsProvidedLatestImage` | use_aws_provided_latest_image | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV. If you specify ``UseAwsProvidedLatestImage``, you can't also specify ``ImageId``. |

Supports update: yes

Discovery: supported
