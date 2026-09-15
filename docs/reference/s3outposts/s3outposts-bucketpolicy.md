# aws.s3outposts.bucketpolicy

**CloudFormation type:** `AWS::S3Outposts::BucketPolicy`

Resource Type Definition for AWS::S3Outposts::BucketPolicy

Region attribute: `region`

**Import ID:** `<region>/Bucket` (AWS::S3Outposts::BucketPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Bucket` |  | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the specified bucket. |
| `PolicyDocument` | policy_document | `map` | required |  | A policy document containing permissions to add to the specified bucket. |

Supports update: yes

Discovery: not supported
