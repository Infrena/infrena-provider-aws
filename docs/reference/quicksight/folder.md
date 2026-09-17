# aws.folder

**CloudFormation type:** `AWS::QuickSight::Folder`

Definition of the AWS::QuickSight::Folder Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|FolderId` (AWS::QuickSight::Folder)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) for the folder.</p> |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that the folder was created.</p> |
| `FolderId` | folder_id | `string` | optional, computed, provider-chosen, replaces on change | aws.folder.FolderId |  |
| `FolderType` | folder_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The time that the folder was last updated.</p> |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `ParentFolderArn` | parent_folder_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.folder.Arn |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `SharingModel` | sharing_model | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
