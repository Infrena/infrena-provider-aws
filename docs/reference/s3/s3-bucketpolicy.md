# aws.s3.bucketpolicy

**CloudFormation type:** `AWS::S3::BucketPolicy`

Applies an Amazon S3 bucket policy to an Amazon S3 bucket. If you are using an identity other than the root user of the AWS-account that owns the bucket, the calling identity must have the ``PutBucketPolicy`` permissions on the specified bucket and belong to the bucket owner's account in order to use this operation.

Region attribute: `region`

**Import ID:** `<region>/Bucket` (AWS::S3::BucketPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Bucket` |  | `string` | required, replaces on change |  | The name of the Amazon S3 bucket to which the policy applies. |
| `PolicyDocument` | policy_document | `string` | required |  | A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide*. |

Supports update: yes

Discovery: supported
