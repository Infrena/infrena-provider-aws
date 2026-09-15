# aws.organizationalunit

**CloudFormation type:** `AWS::Organizations::OrganizationalUnit`

You can use organizational units (OUs) to group accounts together to administer as a single unit. This greatly simplifies the management of your accounts. For example, you can attach a policy-based control to an OU, and all accounts within the OU automatically inherit the policy. You can create multiple OUs within a single organization, and you can create OUs within other OUs. Each OU can contain multiple accounts, and you can move accounts from one OU to another. However, OU names must be unique within a parent OU or root.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Organizations::OrganizationalUnit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of this OU. |
| `Id` |  | `string` | computed |  | The unique identifier (ID) associated with this OU. |
| `Name` |  | `string` | required |  | The friendly name of this OU. |
| `ParentId` | parent_id | `string` | required, replaces on change |  | The unique identifier (ID) of the parent root or OU that you want to create the new OU in. |
| `Path` |  | `string` | computed |  | The path in the organization where this OU exists. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags that you want to attach to the newly created OU. |

Supports update: yes

Discovery: supported (parent resource required)
