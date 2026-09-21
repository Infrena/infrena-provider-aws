# aws.iottwinmaker.workspace

**CloudFormation type:** `AWS::IoTTwinMaker::Workspace`

Resource schema for AWS::IoTTwinMaker::Workspace

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId` (AWS::IoTTwinMaker::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the workspace. |
| `CreationDateTime` | creation_date_time | `string` | computed |  | The date and time when the workspace was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the workspace. |
| `Role` |  | `string` | required |  | The ARN of the execution role associated with the workspace. |
| `S3Location` | s3_location | `string` | required |  | The ARN of the S3 bucket where resources associated with the workspace are stored. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of key-value pairs to associate with a resource. |
| `UpdateDateTime` | update_date_time | `string` | computed |  | The date and time of the current update. |
| `WorkspaceId` | workspace_id | `string` | required, replaces on change | aws.iottwinmaker.workspace.WorkspaceId | The ID of the workspace. |

Supports update: yes

Discovery: supported
