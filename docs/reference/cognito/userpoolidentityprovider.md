# aws.userpoolidentityprovider

**CloudFormation type:** `AWS::Cognito::UserPoolIdentityProvider`

Resource Type definition for AWS::Cognito::UserPoolIdentityProvider

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|ProviderName` (AWS::Cognito::UserPoolIdentityProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttributeMapping` | attribute_mapping | `map` | optional, computed, provider-chosen |  |  |
| `IdpIdentifiers` | idp_identifiers | `list` | optional, computed, provider-chosen |  |  |
| `ProviderDetails` | provider_details | `map` | required |  |  |
| `ProviderName` | provider_name | `string` | required, replaces on change |  |  |
| `ProviderType` | provider_type | `string` | required, replaces on change |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: supported (parent resource required)
