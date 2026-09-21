# aws.legalhold

**CloudFormation type:** `AWS::Backup::LegalHold`

Creates a legal hold on recovery points (backups). A legal hold prevents backups from being deleted while under hold.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Backup::LegalHold)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the legal hold. |
| `CreationDate` | creation_date | `string` | computed |  | The time when the legal hold was created. |
| `Description` |  | `string` | required, replaces on change |  | The description of the legal hold. |
| `LegalHoldId` | legal_hold_id | `string` | computed |  | The ID of the legal hold. |
| `RecoveryPointSelection` | recovery_point_selection | `map` | required, replaces on change |  | The criteria to assign a set of resources, such as resource types or backup vaults. |
| `Status` |  | `string` | computed |  | The status of the legal hold. |
| `Tags` |  | `map` | tags map |  | Optional tags to include. |
| `Title` |  | `string` | required, replaces on change |  | The title of the legal hold. |

Supports update: yes

Discovery: supported
