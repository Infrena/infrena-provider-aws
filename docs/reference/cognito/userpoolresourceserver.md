# aws.userpoolresourceserver

**CloudFormation type:** `AWS::Cognito::UserPoolResourceServer`

Resource Type definition for AWS::Cognito::UserPoolResourceServer

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|Identifier` (AWS::Cognito::UserPoolResourceServer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Identifier` |  | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required |  |  |
| `Scopes` |  | `list` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: supported (parent resource required)
