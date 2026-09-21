# aws.paralleldata

**CloudFormation type:** `AWS::Translate::ParallelData`

A parallel data resource in Amazon Translate used to customize machine translation output.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Translate::ParallelData)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the parallel data resource. |
| `CreatedAt` | created_at | `string` | computed |  | The time at which the parallel data resource was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A custom description for the parallel data resource. |
| `EncryptionKey` | encryption_key | `map` | optional, computed, provider-chosen, replaces on change |  | The encryption key used to encrypt this object. |
| `FailedRecordCount` | failed_record_count | `integer` | computed |  | The number of records unsuccessfully imported. |
| `ImportedDataSize` | imported_data_size | `integer` | computed |  | The number of UTF-8 characters imported from the parallel data input file. |
| `ImportedRecordCount` | imported_record_count | `integer` | computed |  | The number of records successfully imported. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The time at which the parallel data resource was last updated. |
| `Name` |  | `string` | required, replaces on change |  | A custom name for the parallel data resource. Must be unique in the account and region. |
| `ParallelDataConfig` | parallel_data_config | `map` | required |  | Specifies the format and S3 location of the parallel data input file. |
| `SkippedRecordCount` | skipped_record_count | `integer` | computed |  | The number of items skipped during import. |
| `SourceLanguageCode` | source_language_code | `string` | computed |  | The source language of the translations in the parallel data file. |
| `Status` |  | `string` | computed |  | The status of the parallel data resource. |
| `Tags` |  | `map` | tags map |  | Tags associated with the parallel data resource. |
| `TargetLanguageCodes` | target_language_codes | `list` | computed |  | The language codes for the target languages available in the parallel data file. |

Supports update: yes

Discovery: supported
