# aws.backupvault

**CloudFormation type:** `AWS::Backup::BackupVault`

Resource Type definition for AWS::Backup::BackupVault

Region attribute: `region`

**Import ID:** `<region>/BackupVaultName` (AWS::Backup::BackupVault)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPolicy` | access_policy | `string` | optional, computed, provider-chosen |  |  |
| `BackupVaultArn` | backup_vault_arn | `string` | computed |  |  |
| `BackupVaultName` | backup_vault_name | `string` | required, replaces on change |  |  |
| `BackupVaultTags` | backup_vault_tags | `map` | optional, computed, provider-chosen |  |  |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LockConfiguration` | lock_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Notifications` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
