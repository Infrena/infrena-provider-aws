# aws.userhierarchygroup

**CloudFormation type:** `AWS::Connect::UserHierarchyGroup`

Resource Type definition for AWS::Connect::UserHierarchyGroup

Region attribute: `region`

**Import ID:** `<region>/UserHierarchyGroupArn` (AWS::Connect::UserHierarchyGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Name` |  | `string` | required |  | The name of the user hierarchy group. |
| `ParentGroupArn` | parent_group_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) for the User hierarchy group. |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `UserHierarchyGroupArn` | user_hierarchy_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the User hierarchy group. |

Supports update: yes

Discovery: supported (parent resource required)
