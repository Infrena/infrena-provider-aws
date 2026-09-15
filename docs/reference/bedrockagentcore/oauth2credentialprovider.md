# aws.oauth2credentialprovider

**CloudFormation type:** `AWS::BedrockAgentCore::OAuth2CredentialProvider`

Resource Type definition for AWS::BedrockAgentCore::OAuth2CredentialProvider

Region attribute: `region`

**Import ID:** `<region>/CredentialProviderArn` (AWS::BedrockAgentCore::OAuth2CredentialProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CallbackUrl` | callback_url | `string` | computed |  | The callback URL for the OAuth2 authorization flow |
| `ClientSecretArn` | client_secret_arn | `map` | computed |  | Contains information about a secret in AWS Secrets Manager |
| `ClientSecretJsonKey` | client_secret_json_key | `string` | computed |  | The JSON key within the secret that contains the client secret value |
| `ClientSecretSource` | client_secret_source | `string` | computed |  | The source of the client secret |
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the credential provider was created |
| `CredentialProviderArn` | credential_provider_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the OAuth2 credential provider |
| `CredentialProviderVendor` | credential_provider_vendor | `string` | required, replaces on change |  | The vendor of the OAuth2 credential provider |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the credential provider was last updated |
| `Name` |  | `string` | required, replaces on change |  | The name of the OAuth2 credential provider |
| `Oauth2ProviderConfigInput` | oauth2_provider_config_input | `map` | optional, computed, provider-chosen, write-only |  | Input configuration for an OAuth2 provider |
| `Oauth2ProviderConfigOutput` | oauth2_provider_config_output | `map` | computed |  | Output configuration for an OAuth2 provider |
| `Status` |  | `string` | computed |  | The current status of the OAuth2 credential provider |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the OAuth2 credential provider |

Supports update: yes

Discovery: supported
