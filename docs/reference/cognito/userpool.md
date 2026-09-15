# aws.userpool

**CloudFormation type:** `AWS::Cognito::UserPool`

Definition of AWS::Cognito::UserPool Resource Type

Region attribute: `region`

**Import ID:** `<region>/UserPoolId` (AWS::Cognito::UserPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountRecoverySetting` | account_recovery_setting | `map` | optional, computed, provider-chosen |  |  |
| `AdminCreateUserConfig` | admin_create_user_config | `map` | optional, computed, provider-chosen |  |  |
| `AliasAttributes` | alias_attributes | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AutoVerifiedAttributes` | auto_verified_attributes | `list` | optional, computed, provider-chosen |  |  |
| `DeletionProtection` | deletion_protection | `string` | optional, computed, provider-chosen |  |  |
| `DeviceConfiguration` | device_configuration | `map` | optional, computed, provider-chosen |  |  |
| `EmailAuthenticationMessage` | email_authentication_message | `string` | optional, computed, provider-chosen |  |  |
| `EmailAuthenticationSubject` | email_authentication_subject | `string` | optional, computed, provider-chosen |  |  |
| `EmailConfiguration` | email_configuration | `map` | optional, computed, provider-chosen |  |  |
| `EmailVerificationMessage` | email_verification_message | `string` | optional, computed, provider-chosen |  |  |
| `EmailVerificationSubject` | email_verification_subject | `string` | optional, computed, provider-chosen |  |  |
| `EnabledMfas` | enabled_mfas | `list` | optional, computed, provider-chosen, write-only |  |  |
| `IssuerConfiguration` | issuer_configuration | `map` | optional, computed, provider-chosen |  |  |
| `KeyConfiguration` | key_configuration | `map` | optional, computed, provider-chosen |  |  |
| `LambdaConfig` | lambda_config | `map` | optional, computed, provider-chosen |  |  |
| `MfaConfiguration` | mfa_configuration | `string` | optional, computed, provider-chosen |  |  |
| `Policies` |  | `map` | optional, computed, provider-chosen |  |  |
| `ProviderName` | provider_name | `string` | computed |  |  |
| `ProviderURL` | provider_url | `string` | computed |  |  |
| `Schema` |  | `list` | optional, computed, provider-chosen |  |  |
| `SmsAuthenticationMessage` | sms_authentication_message | `string` | optional, computed, provider-chosen |  |  |
| `SmsConfiguration` | sms_configuration | `map` | optional, computed, provider-chosen |  |  |
| `SmsVerificationMessage` | sms_verification_message | `string` | optional, computed, provider-chosen |  |  |
| `UserAttributeUpdateSettings` | user_attribute_update_settings | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolAddOns` | user_pool_add_ons | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | computed |  |  |
| `UserPoolName` | user_pool_name | `string` | optional, computed, provider-chosen |  |  |
| `UserPoolTags` | user_pool_tags | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolTier` | user_pool_tier | `string` | optional, computed, provider-chosen |  |  |
| `UsernameAttributes` | username_attributes | `list` | optional, computed, provider-chosen |  |  |
| `UsernameConfiguration` | username_configuration | `map` | optional, computed, provider-chosen |  |  |
| `VerificationMessageTemplate` | verification_message_template | `map` | optional, computed, provider-chosen |  |  |
| `WebAuthnFactorConfiguration` | web_authn_factor_configuration | `string` | optional, computed, provider-chosen |  |  |
| `WebAuthnRelyingPartyID` | web_authn_relying_party_id | `string` | optional, computed, provider-chosen |  |  |
| `WebAuthnUserVerification` | web_authn_user_verification | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
