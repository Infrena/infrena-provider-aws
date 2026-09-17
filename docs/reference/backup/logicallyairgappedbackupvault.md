# aws.logicallyairgappedbackupvault

**CloudFormation type:** `AWS::Backup::LogicallyAirGappedBackupVault`

Resource Type definition for AWS::Backup::LogicallyAirGappedBackupVault

Region attribute: `region`

**Import ID:** `<region>/BackupVaultName` (AWS::Backup::LogicallyAirGappedBackupVault)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPolicy` | access_policy | `string` | optional, computed, provider-chosen |  |  |
| `BackupVaultArn` | backup_vault_arn | `string` | computed |  |  |
| `BackupVaultName` | backup_vault_name | `string` | required, replaces on change |  |  |
| `BackupVaultTags` | backup_vault_tags | `map` | optional, computed, provider-chosen |  |  |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `MaxRetentionDays` | max_retention_days | `integer` | required, replaces on change |  |  |
| `MinRetentionDays` | min_retention_days | `integer` | required, replaces on change |  |  |
| `MpaApprovalTeamArn` | mpa_approval_team_arn | `string` | optional, computed, provider-chosen | aws.approvalteam.Arn |  |
| `Notifications` |  | `map` | optional, computed, provider-chosen |  |  |
| `VaultState` | vault_state | `string` | computed |  |  |
| `VaultType` | vault_type | `string` | computed |  |  |

Supports update: yes

Discovery: supported
