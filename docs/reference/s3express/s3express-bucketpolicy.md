# aws.s3express.bucketpolicy

**CloudFormation type:** `AWS::S3Express::BucketPolicy`

Resource Type definition for AWS::S3Express::BucketPolicy.

Region attribute: `region`

**Import ID:** `<region>/Bucket` (AWS::S3Express::BucketPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Bucket` |  | `string` | required, replaces on change |  | The name of the S3 directory bucket to which the policy applies. |
| `PolicyDocument` | policy_document | `string` | required |  | A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. |

Supports update: yes

Discovery: supported
