# aws.identitypoolroleattachment

**CloudFormation type:** `AWS::Cognito::IdentityPoolRoleAttachment`

Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Cognito::IdentityPoolRoleAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `IdentityPoolId` | identity_pool_id | `string` | required, replaces on change | aws.identitypool.Id |  |
| `RoleMappings` | role_mappings | `map` | optional, computed, provider-chosen |  |  |
| `Roles` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
