# aws.s3tables.namespace

**CloudFormation type:** `AWS::S3Tables::Namespace`

Resource Type definition for AWS::S3Tables::Namespace

Region attribute: `region`

**Import ID:** `<region>/TableBucketARN|Namespace` (AWS::S3Tables::Namespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Namespace` |  | `string` | required, replaces on change |  | A name for the namespace. |
| `TableBucketARN` | table_bucket_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the specified table bucket. |

Supports update: no

Discovery: supported (parent resource required)
