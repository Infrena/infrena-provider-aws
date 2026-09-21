# aws.tablepolicy

**CloudFormation type:** `AWS::S3Tables::TablePolicy`

Resource Type definition for AWS::S3Tables::TablePolicy

Region attribute: `region`

**Import ID:** `<region>/TableARN` (AWS::S3Tables::TablePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Namespace` |  | `string` | computed |  | The namespace that the table belongs to. |
| `ResourcePolicy` | resource_policy | `string` | required |  | A policy document containing permissions to add to the specified table. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. |
| `TableARN` | table_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the specified table. |
| `TableBucketARN` | table_bucket_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified table bucket. |
| `TableName` | table_name | `string` | computed |  | The name for the table. |

Supports update: yes

Discovery: supported (parent resource required)
