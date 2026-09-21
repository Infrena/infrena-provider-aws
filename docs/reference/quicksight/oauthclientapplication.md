# aws.oauthclientapplication

**CloudFormation type:** `AWS::QuickSight::OAuthClientApplication`

Definition of AWS::QuickSight::OAuthClientApplication Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::QuickSight::OAuthClientApplication)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClientId` | client_id | `string` | optional, computed, provider-chosen, write-only |  |  |
| `ClientSecret` | client_secret | `string` | optional, computed, provider-chosen, sensitive, write-only |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `DataSourceType` | data_source_type | `string` | optional, computed, provider-chosen |  |  |
| `IdentityProviderVpcConnectionProperties` | identity_provider_vpc_connection_properties | `map` | optional, computed, provider-chosen |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `OAuthAuthorizationEndpointUrl` | o_auth_authorization_endpoint_url | `string` | optional, computed, provider-chosen |  |  |
| `OAuthClientApplicationId` | o_auth_client_application_id | `string` | required, replaces on change | aws.oauthclientapplication.OAuthClientApplicationId |  |
| `OAuthClientAuthenticationType` | o_auth_client_authentication_type | `string` | required, replaces on change |  |  |
| `OAuthScopes` | o_auth_scopes | `string` | optional, computed, provider-chosen |  |  |
| `OAuthTokenEndpointUrl` | o_auth_token_endpoint_url | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
