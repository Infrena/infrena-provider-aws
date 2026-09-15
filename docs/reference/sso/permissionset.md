# aws.permissionset

**CloudFormation type:** `AWS::SSO::PermissionSet`

Resource Type definition for SSO PermissionSet

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|PermissionSetArn` (AWS::SSO::PermissionSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CustomerManagedPolicyReferences` | customer_managed_policy_references | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The permission set description. |
| `InlinePolicy` | inline_policy | `string` | optional, computed, provider-chosen |  | The inline policy to put in permission set. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.sso.instance.InstanceArn | The sso instance arn that the permission set is owned. |
| `ManagedPolicies` | managed_policies | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  | The name you want to assign to this permission set. |
| `PermissionSetArn` | permission_set_arn | `string` | computed |  | The permission set that the policy will be attached to |
| `PermissionsBoundary` | permissions_boundary | `map` | optional, computed, provider-chosen |  |  |
| `RelayStateType` | relay_state_type | `string` | optional, computed, provider-chosen |  | The relay state URL that redirect links to any service in the AWS Management Console. |
| `SessionDuration` | session_duration | `string` | optional, computed, provider-chosen |  | The length of time that a user can be signed in to an AWS account. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
