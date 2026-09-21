# aws.querysuggestionsblocklist

**CloudFormation type:** `AWS::Kendra::QuerySuggestionsBlockList`

A block list used for query suggestions for an Amazon Kendra index. A block list contains words or phrases that should not appear as query suggestions.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Kendra::QuerySuggestionsBlockList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the query suggestions block list. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the block list. |
| `Id` |  | `string` | computed |  | The identifier of the block list. |
| `IndexId` | index_id | `string` | required, replaces on change | aws.kendra.index.Id | The identifier of the index for the block list. |
| `Name` |  | `string` | required |  | The name of the block list. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role with permission to access the S3 bucket that contains the block list text file. |
| `SourceS3Path` | source_s3_path | `map` | required |  | Information required to find a specific file in an Amazon S3 bucket. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that identify or categorize the block list. |

Supports update: yes

Discovery: supported (parent resource required)
