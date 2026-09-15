# aws.rolealias

**CloudFormation type:** `AWS::IoT::RoleAlias`

Use the AWS::IoT::RoleAlias resource to declare an AWS IoT RoleAlias.

Region attribute: `region`

**Import ID:** `<region>/RoleAlias` (AWS::IoT::RoleAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CredentialDurationSeconds` | credential_duration_seconds | `integer` | optional, computed, provider-chosen |  |  |
| `RoleAlias` | role_alias | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RoleAliasArn` | role_alias_arn | `string` | computed |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
