# aws.backupselection

**CloudFormation type:** `AWS::Backup::BackupSelection`

Resource Type definition for AWS::Backup::BackupSelection

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Backup::BackupSelection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BackupPlanId` | backup_plan_id | `string` | required, replaces on change | aws.backupplan.BackupPlanId |  |
| `BackupSelection` | backup_selection | `map` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `SelectionId` | selection_id | `string` | computed |  |  |

Supports update: no

Discovery: supported
