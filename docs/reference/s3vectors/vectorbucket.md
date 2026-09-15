# aws.vectorbucket

**CloudFormation type:** `AWS::S3Vectors::VectorBucket`

Resource Type definition for AWS::S3Vectors::VectorBucket

Region attribute: `region`

**Import ID:** `<region>/VectorBucketArn` (AWS::S3Vectors::VectorBucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Date and time when the vector bucket was created. |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The encryption configuration for the vector bucket. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | User tags (key-value pairs) to associate with the vector bucket. |
| `VectorBucketArn` | vector_bucket_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the vector bucket. |
| `VectorBucketName` | vector_bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the vector bucket. |

Supports update: yes

Discovery: supported
