# aws.locations3

**CloudFormation type:** `AWS::DataSync::LocationS3`

Resource schema for AWS::DataSync::LocationS3

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationS3)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon S3 bucket location. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the S3 location that was described. |
| `S3BucketArn` | s3_bucket_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) of the Amazon S3 bucket. |
| `S3Config` | s3_config | `map` | required |  | The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket. |
| `S3StorageClass` | s3_storage_class | `string` | optional, computed, provider-chosen |  | The Amazon S3 storage class you want to store your files in when this location is used as a task destination. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
