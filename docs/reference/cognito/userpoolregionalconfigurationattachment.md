# aws.userpoolregionalconfigurationattachment

**CloudFormation type:** `AWS::Cognito::UserPoolRegionalConfigurationAttachment`

Resource Type definition for AWS::Cognito::UserPoolRegionalConfigurationAttachment

Region attribute: `region`

**Import ID:** `<region>/UserPoolId` (AWS::Cognito::UserPoolRegionalConfigurationAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EmailConfiguration` | email_configuration | `map` | optional, computed, provider-chosen |  |  |
| `LambdaConfig` | lambda_config | `map` | optional, computed, provider-chosen |  |  |
| `SmsConfiguration` | sms_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the replica. Set to ACTIVE or INACTIVE. |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |
| `UserPoolTags` | user_pool_tags | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: not supported
