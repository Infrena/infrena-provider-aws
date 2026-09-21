# aws.syncjob

**CloudFormation type:** `AWS::IoTTwinMaker::SyncJob`

Resource schema for AWS::IoTTwinMaker::SyncJob

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId|SyncSource` (AWS::IoTTwinMaker::SyncJob)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the SyncJob. |
| `CreationDateTime` | creation_date_time | `string` | computed |  | The date and time when the sync job was created. |
| `State` |  | `string` | computed |  | The state of SyncJob. |
| `SyncRole` | sync_role | `string` | required, replaces on change |  | The IAM Role that execute SyncJob. |
| `SyncSource` | sync_source | `string` | required, replaces on change |  | The source of the SyncJob. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A key-value pair to associate with a resource. |
| `UpdateDateTime` | update_date_time | `string` | computed |  | The date and time when the sync job was updated. |
| `WorkspaceId` | workspace_id | `string` | required, replaces on change | aws.iottwinmaker.workspace.WorkspaceId | The ID of the workspace. |

Supports update: no

Discovery: supported (parent resource required)
