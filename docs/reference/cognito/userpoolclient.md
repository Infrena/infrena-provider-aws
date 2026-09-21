# aws.userpoolclient

**CloudFormation type:** `AWS::Cognito::UserPoolClient`

Resource Type definition for AWS::Cognito::UserPoolClient

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|ClientId` (AWS::Cognito::UserPoolClient)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessTokenValidity` | access_token_validity | `integer` | optional, computed, provider-chosen |  |  |
| `AllowedOAuthFlows` | allowed_o_auth_flows | `list` | optional, computed, provider-chosen |  |  |
| `AllowedOAuthFlowsUserPoolClient` | allowed_o_auth_flows_user_pool_client | `boolean` | optional, computed, provider-chosen |  |  |
| `AllowedOAuthScopes` | allowed_o_auth_scopes | `list` | optional, computed, provider-chosen |  |  |
| `AnalyticsConfiguration` | analytics_configuration | `map` | optional, computed, provider-chosen |  |  |
| `AuthSessionValidity` | auth_session_validity | `integer` | optional, computed, provider-chosen |  |  |
| `CallbackURLs` | callback_ur_ls | `list` | optional, computed, provider-chosen |  |  |
| `ClientId` | client_id | `string` | computed |  |  |
| `ClientName` | client_name | `string` | optional, computed, provider-chosen |  |  |
| `ClientSecret` | client_secret | `string` | computed |  |  |
| `DefaultRedirectURI` | default_redirect_uri | `string` | optional, computed, provider-chosen |  |  |
| `EnablePropagateAdditionalUserContextData` | enable_propagate_additional_user_context_data | `boolean` | optional, computed, provider-chosen |  |  |
| `EnableTokenRevocation` | enable_token_revocation | `boolean` | optional, computed, provider-chosen |  |  |
| `ExplicitAuthFlows` | explicit_auth_flows | `list` | optional, computed, provider-chosen |  |  |
| `GenerateSecret` | generate_secret | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `IdTokenValidity` | id_token_validity | `integer` | optional, computed, provider-chosen |  |  |
| `LogoutURLs` | logout_ur_ls | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | computed |  |  |
| `PreventUserExistenceErrors` | prevent_user_existence_errors | `string` | optional, computed, provider-chosen |  |  |
| `ReadAttributes` | read_attributes | `list` | optional, computed, provider-chosen |  |  |
| `RefreshTokenRotation` | refresh_token_rotation | `map` | optional, computed, provider-chosen |  |  |
| `RefreshTokenValidity` | refresh_token_validity | `integer` | optional, computed, provider-chosen |  |  |
| `SupportedIdentityProviders` | supported_identity_providers | `list` | optional, computed, provider-chosen |  |  |
| `TokenValidityUnits` | token_validity_units | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |
| `WriteAttributes` | write_attributes | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
