# aws.awslogsource

**CloudFormation type:** `AWS::SecurityLake::AwsLogSource`

Resource Type definition for AWS::SecurityLake::AwsLogSource

Region attribute: `region`

**Import ID:** `<region>/SourceName|SourceVersion` (AWS::SecurityLake::AwsLogSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Accounts` |  | `list` | optional, computed, provider-chosen |  | AWS account where you want to collect logs from. |
| `DataLakeArn` | data_lake_arn | `string` | required, replaces on change | aws.datalake.Arn | The ARN for the data lake. |
| `SourceName` | source_name | `string` | required, replaces on change |  | The name for a AWS source. This must be a Regionally unique value. |
| `SourceVersion` | source_version | `string` | required, replaces on change |  | The version for a AWS source. This must be a Regionally unique value. |

Supports update: yes

Discovery: supported
