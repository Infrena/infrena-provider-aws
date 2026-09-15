# aws.vectorbucketpolicy

**CloudFormation type:** `AWS::S3Vectors::VectorBucketPolicy`

Resource Type definition for AWS::S3Vectors::VectorBucketPolicy

Region attribute: `region`

**Import ID:** `<region>/VectorBucketArn` (AWS::S3Vectors::VectorBucketPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Policy` |  | `string` | required |  | A policy document containing permissions to add to the specified vector bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. |
| `VectorBucketArn` | vector_bucket_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.vectorbucket.VectorBucketArn | The Amazon Resource Name (ARN) of the vector bucket. |
| `VectorBucketName` | vector_bucket_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the vector bucket |

Supports update: yes

Discovery: supported
