# aws.identitypool

**CloudFormation type:** `AWS::Cognito::IdentityPool`

Resource Type definition for AWS::Cognito::IdentityPool

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Cognito::IdentityPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowClassicFlow` | allow_classic_flow | `boolean` | optional, computed, provider-chosen |  |  |
| `AllowUnauthenticatedIdentities` | allow_unauthenticated_identities | `boolean` | required |  |  |
| `CognitoEvents` | cognito_events | `map` | optional, computed, provider-chosen, write-only |  |  |
| `CognitoIdentityProviders` | cognito_identity_providers | `list` | optional, computed, provider-chosen |  |  |
| `CognitoStreams` | cognito_streams | `map` | optional, computed, provider-chosen, write-only |  |  |
| `DeveloperProviderName` | developer_provider_name | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `IdentityPoolName` | identity_pool_name | `string` | optional, computed, provider-chosen |  |  |
| `IdentityPoolTags` | identity_pool_tags | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Name` |  | `string` | computed |  |  |
| `OpenIdConnectProviderARNs` | open_id_connect_provider_ar_ns | `list` | optional, computed, provider-chosen |  |  |
| `PushSync` | push_sync | `map` | optional, computed, provider-chosen, write-only |  |  |
| `SamlProviderARNs` | saml_provider_ar_ns | `list` | optional, computed, provider-chosen |  |  |
| `SupportedLoginProviders` | supported_login_providers | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
