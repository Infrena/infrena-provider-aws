# aws.tablebucketpolicy

**CloudFormation type:** `AWS::S3Tables::TableBucketPolicy`

Applies an IAM resource policy to a table bucket.

Region attribute: `region`

**Import ID:** `<region>/TableBucketARN` (AWS::S3Tables::TableBucketPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourcePolicy` | resource_policy | `string` | required |  | A policy document containing permissions to add to the specified table bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. |
| `TableBucketARN` | table_bucket_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the table bucket to which the policy applies. |

Supports update: yes

Discovery: supported
