# aws.aps.resourcepolicy

**CloudFormation type:** `AWS::APS::ResourcePolicy`

Resource Type definition for AWS::APS::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/WorkspaceArn` (AWS::APS::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `string` | required |  | The JSON to use as the Resource-based Policy. |
| `WorkspaceArn` | workspace_arn | `string` | required, replaces on change | aws.aps.workspace.Arn | The Arn of an APS Workspace that the PolicyDocument will be attached to. |

Supports update: yes

Discovery: supported
