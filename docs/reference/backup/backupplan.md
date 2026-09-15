# aws.backupplan

**CloudFormation type:** `AWS::Backup::BackupPlan`

Resource Type definition for AWS::Backup::BackupPlan

Region attribute: `region`

**Import ID:** `<region>/BackupPlanId` (AWS::Backup::BackupPlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BackupPlan` | backup_plan | `map` | required |  |  |
| `BackupPlanArn` | backup_plan_arn | `string` | computed |  |  |
| `BackupPlanId` | backup_plan_id | `string` | computed |  |  |
| `BackupPlanTags` | backup_plan_tags | `map` | optional, computed, provider-chosen |  |  |
| `VersionId` | version_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
