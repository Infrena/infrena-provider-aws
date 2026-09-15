# aws.s3outposts.bucket

**CloudFormation type:** `AWS::S3Outposts::Bucket`

Resource Type Definition for AWS::S3Outposts::Bucket

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::S3Outposts::Bucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified bucket. |
| `BucketName` | bucket_name | `string` | required, replaces on change |  | A name for the bucket. |
| `LifecycleConfiguration` | lifecycle_configuration | `map` | optional, computed, provider-chosen |  | Rules that define how Amazon S3Outposts manages objects during their lifetime. |
| `OutpostId` | outpost_id | `string` | required, replaces on change |  | The id of the customer outpost on which the bucket resides. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this S3Outposts bucket. |

Supports update: yes

Discovery: supported (parent resource required)
