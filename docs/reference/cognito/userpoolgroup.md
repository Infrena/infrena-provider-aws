# aws.userpoolgroup

**CloudFormation type:** `AWS::Cognito::UserPoolGroup`

Resource Type definition for AWS::Cognito::UserPoolGroup

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|GroupName` (AWS::Cognito::UserPoolGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `GroupName` | group_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Precedence` |  | `integer` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: supported (parent resource required)
