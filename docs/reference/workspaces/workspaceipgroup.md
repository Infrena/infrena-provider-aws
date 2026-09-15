# aws.workspaceipgroup

**CloudFormation type:** `AWS::WorkSpaces::WorkspaceIpGroup`

Resource Type definition for an IP access control group for Amazon WorkSpaces.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::WorkSpaces::WorkspaceIpGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the IP access control group. |
| `GroupDesc` | group_desc | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the group. |
| `GroupId` | group_id | `string` | computed |  | The identifier of the IP access control group. |
| `GroupName` | group_name | `string` | required, replaces on change |  | The name of the group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the IP access control group. |
| `UserRules` | user_rules | `list` | optional, computed, provider-chosen |  | The rules for the IP access control group. |

Supports update: yes

Discovery: supported
