# aws.userpooluicustomizationattachment

**CloudFormation type:** `AWS::Cognito::UserPoolUICustomizationAttachment`

Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|ClientId` (AWS::Cognito::UserPoolUICustomizationAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CSS` |  | `string` | optional, computed, provider-chosen |  |  |
| `ClientId` | client_id | `string` | required, replaces on change |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: not supported
