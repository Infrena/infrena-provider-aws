# aws.workspaces.workspace

**CloudFormation type:** `AWS::WorkSpaces::Workspace`

Resource Type definition for AWS::WorkSpaces::Workspace

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId` (AWS::WorkSpaces::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BundleId` | bundle_id | `string` | required |  |  |
| `DirectoryId` | directory_id | `string` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `RootVolumeEncryptionEnabled` | root_volume_encryption_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UserName` | user_name | `string` | required, replaces on change |  |  |
| `UserVolumeEncryptionEnabled` | user_volume_encryption_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `VolumeEncryptionKey` | volume_encryption_key | `string` | optional, computed, provider-chosen |  |  |
| `WorkspaceId` | workspace_id | `string` | computed |  |  |
| `WorkspaceProperties` | workspace_properties | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
