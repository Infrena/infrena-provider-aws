# aws.faq

**CloudFormation type:** `AWS::Kendra::Faq`

A Kendra FAQ resource

Region attribute: `region`

**Import ID:** `<region>/Id|IndexId` (AWS::Kendra::Faq)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the FAQ |
| `FileFormat` | file_format | `string` | optional, computed, provider-chosen, replaces on change |  | Format of the input file |
| `Id` |  | `string` | computed |  | Unique ID of the FAQ |
| `IndexId` | index_id | `string` | required, replaces on change | aws.kendra.index.Id | Unique ID of Index |
| `LanguageCode` | language_code | `string` | optional, computed, provider-chosen |  | The code for a language. |
| `Name` |  | `string` | required, replaces on change |  | FAQ name |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | FAQ role ARN |
| `S3Path` | s3_path | `map` | required, replaces on change |  | FAQ S3 path |
| `Tags` |  | `map` | tags map |  | List of tags |

Supports update: yes

Discovery: supported
