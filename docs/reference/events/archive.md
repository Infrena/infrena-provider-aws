# aws.archive

**CloudFormation type:** `AWS::Events::Archive`

Resource Type definition for AWS::Events::Archive

Region attribute: `region`

**Import ID:** `<region>/ArchiveName` (AWS::Events::Archive)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArchiveName` | archive_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EventPattern` | event_pattern | `map` | optional, computed, provider-chosen |  |  |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen |  |  |
| `RetentionDays` | retention_days | `integer` | optional, computed, provider-chosen |  |  |
| `SourceArn` | source_arn | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
