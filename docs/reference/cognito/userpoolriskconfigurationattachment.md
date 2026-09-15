# aws.userpoolriskconfigurationattachment

**CloudFormation type:** `AWS::Cognito::UserPoolRiskConfigurationAttachment`

Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|ClientId` (AWS::Cognito::UserPoolRiskConfigurationAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountTakeoverRiskConfiguration` | account_takeover_risk_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ClientId` | client_id | `string` | required, replaces on change |  |  |
| `CompromisedCredentialsRiskConfiguration` | compromised_credentials_risk_configuration | `map` | optional, computed, provider-chosen |  |  |
| `RiskExceptionConfiguration` | risk_exception_configuration | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: not supported
