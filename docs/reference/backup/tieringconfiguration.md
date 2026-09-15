# aws.tieringconfiguration

**CloudFormation type:** `AWS::Backup::TieringConfiguration`

Resource Type definition for AWS::Backup::TieringConfiguration

Region attribute: `region`

**Import ID:** `<region>/TieringConfigurationName` (AWS::Backup::TieringConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BackupVaultName` | backup_vault_name | `string` | required |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  |  |
| `ResourceSelection` | resource_selection | `list` | required |  |  |
| `TieringConfigurationArn` | tiering_configuration_arn | `string` | computed |  |  |
| `TieringConfigurationName` | tiering_configuration_name | `string` | required, replaces on change |  |  |
| `TieringConfigurationTags` | tiering_configuration_tags | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
