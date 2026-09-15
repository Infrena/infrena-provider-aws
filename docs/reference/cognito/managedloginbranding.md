# aws.managedloginbranding

**CloudFormation type:** `AWS::Cognito::ManagedLoginBranding`

Resource Type definition for AWS::Cognito::ManagedLoginBranding

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|ManagedLoginBrandingId` (AWS::Cognito::ManagedLoginBranding)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Assets` |  | `list` | optional, computed, provider-chosen |  |  |
| `ClientId` | client_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ManagedLoginBrandingId` | managed_login_branding_id | `string` | computed |  |  |
| `ReturnMergedResources` | return_merged_resources | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `Settings` |  | `map` | optional, computed, provider-chosen |  |  |
| `UseCognitoProvidedValues` | use_cognito_provided_values | `boolean` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: not supported
