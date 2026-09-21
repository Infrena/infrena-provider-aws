# aws.vocabularyfilter

**CloudFormation type:** `AWS::Transcribe::VocabularyFilter`

Creates a custom vocabulary filter that you can use to mask, delete, or flag specific words from your transcript.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Transcribe::VocabularyFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the vocabulary filter. |
| `DataAccessRoleArn` | data_access_role_arn | `string` | optional, computed, provider-chosen, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files. |
| `LanguageCode` | language_code | `string` | required, replaces on change |  | The language code that represents the language of the entries in your vocabulary filter. |
| `Tags` |  | `map` | tags map |  | Tags associated with the vocabulary filter. |
| `VocabularyFilterFileUri` | vocabulary_filter_file_uri | `string` | optional, computed, provider-chosen, write-only |  | The Amazon S3 location of the text file that contains your custom vocabulary filter terms. |
| `VocabularyFilterName` | vocabulary_filter_name | `string` | required, replaces on change |  | A unique name, chosen by you, for your custom vocabulary filter. |
| `Words` |  | `list` | optional, computed, provider-chosen, write-only |  | Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request. |

Supports update: yes

Discovery: supported
