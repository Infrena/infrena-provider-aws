# aws.assignment

**CloudFormation type:** `AWS::SSO::Assignment`

Resource Type definition for SSO assignmet

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|TargetId|TargetType|PermissionSetArn|PrincipalType|PrincipalId` (AWS::SSO::Assignment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.sso.instance.InstanceArn | The sso instance that the permission set is owned. |
| `PermissionSetArn` | permission_set_arn | `string` | required, replaces on change | aws.permissionset.PermissionSetArn | The permission set that the assignment will be assigned |
| `PrincipalId` | principal_id | `string` | required, replaces on change |  | The assignee's identifier, user id/group id |
| `PrincipalType` | principal_type | `string` | required, replaces on change |  | The assignee's type, user/group |
| `TargetId` | target_id | `string` | required, replaces on change |  | The account id to be provisioned. |
| `TargetType` | target_type | `string` | required, replaces on change |  | The type of resource to be provisioned to, only aws account now |

Supports update: no

Discovery: supported (parent resource required)
