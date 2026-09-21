# aws.userpoolusertogroupattachment

**CloudFormation type:** `AWS::Cognito::UserPoolUserToGroupAttachment`

Resource Type definition for AWS::Cognito::UserPoolUserToGroupAttachment

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|GroupName|Username` (AWS::Cognito::UserPoolUserToGroupAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupName` | group_name | `string` | required, replaces on change |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |
| `Username` |  | `string` | required, replaces on change |  |  |

Supports update: no

Discovery: not supported
