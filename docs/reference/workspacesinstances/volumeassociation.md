# aws.volumeassociation

**CloudFormation type:** `AWS::WorkspacesInstances::VolumeAssociation`

Resource Type definition for AWS::WorkspacesInstances::VolumeAssociation

Region attribute: `region`

**Import ID:** `<region>/WorkspaceInstanceId|VolumeId|Device` (AWS::WorkspacesInstances::VolumeAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Device` |  | `string` | required, replaces on change |  | The device name for the volume attachment |
| `DisassociateMode` | disassociate_mode | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Mode to use when disassociating the volume |
| `VolumeId` | volume_id | `string` | required, replaces on change | aws.workspacesinstances.volume.VolumeId | ID of the volume to attach to the workspace instance |
| `WorkspaceInstanceId` | workspace_instance_id | `string` | required, replaces on change | aws.workspaceinstance.WorkspaceInstanceId | ID of the workspace instance to associate with the volume |

Supports update: no

Discovery: supported
