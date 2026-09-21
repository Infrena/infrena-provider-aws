# aws.workspaceinstance

**CloudFormation type:** `AWS::WorkspacesInstances::WorkspaceInstance`

Resource Type definition for AWS::WorkspacesInstances::WorkspaceInstance

Region attribute: `region`

**Import ID:** `<region>/WorkspaceInstanceId` (AWS::WorkspacesInstances::WorkspaceInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EC2ManagedInstance` | ec2_managed_instance | `map` | computed |  |  |
| `ManagedInstance` | managed_instance | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ProvisionState` | provision_state | `string` | computed |  | The current state of the workspace instance |
| `Tags` |  | `map` | tags map |  |  |
| `WorkspaceInstanceId` | workspace_instance_id | `string` | computed |  | Unique identifier for the workspace instance |

Supports update: yes

Discovery: supported
