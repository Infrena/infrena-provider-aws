# aws.mailmanagerarchive

**CloudFormation type:** `AWS::SES::MailManagerArchive`

Definition of AWS::SES::MailManagerArchive Resource Type

Region attribute: `region`

**Import ID:** `<region>/ArchiveId` (AWS::SES::MailManagerArchive)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArchiveArn` | archive_arn | `string` | computed |  |  |
| `ArchiveId` | archive_id | `string` | computed |  |  |
| `ArchiveName` | archive_name | `string` | optional, computed, provider-chosen |  |  |
| `ArchiveState` | archive_state | `string` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Retention` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
