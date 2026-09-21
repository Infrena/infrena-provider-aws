# aws.thesaurus

**CloudFormation type:** `AWS::Kendra::Thesaurus`

A thesaurus for an Amazon Kendra index. The thesaurus contains a list of synonyms in Solr format.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Kendra::Thesaurus)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the thesaurus. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the thesaurus. |
| `Id` |  | `string` | computed |  | The identifier of the thesaurus. |
| `IndexId` | index_id | `string` | required, replaces on change | aws.kendra.index.Id | The identifier of the index for the thesaurus. |
| `Name` |  | `string` | required |  | A name for the thesaurus. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | An IAM role that gives Amazon Kendra permissions to access the thesaurus file specified in SourceS3Path. |
| `SourceS3Path` | source_s3_path | `map` | required |  | Information required to find a specific file in an Amazon S3 bucket. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that identify or categorize the thesaurus. |

Supports update: yes

Discovery: supported (parent resource required)
