# aws.s3vectors.index

**CloudFormation type:** `AWS::S3Vectors::Index`

Resource Type definition for AWS::S3Vectors::Index

Region attribute: `region`

**Import ID:** `<region>/IndexArn` (AWS::S3Vectors::Index)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Date and time when the vector index was created. |
| `DataType` | data_type | `string` | required, replaces on change |  | The data type of the vectors to be inserted into the vector index. |
| `Dimension` |  | `integer` | required, replaces on change |  | The dimensions of the vectors to be inserted into the vector index. |
| `DistanceMetric` | distance_metric | `string` | required, replaces on change |  | The distance metric to be used for similarity search. |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The encryption configuration for the index. |
| `IndexArn` | index_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the index |
| `IndexName` | index_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the vector index to create. |
| `MetadataConfiguration` | metadata_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The metadata configuration for the vector index. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | User tags (key-value pairs) to associate with the index. |
| `VectorBucketArn` | vector_bucket_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.vectorbucket.VectorBucketArn | The Amazon Resource Name (ARN) of the vector bucket. |
| `VectorBucketName` | vector_bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the vector bucket that contains the vector index. |

Supports update: yes

Discovery: supported (parent resource required)
