# aws.apikeycredentialprovider

**CloudFormation type:** `AWS::BedrockAgentCore::ApiKeyCredentialProvider`

Resource Type definition for AWS::BedrockAgentCore::ApiKeyCredentialProvider

Region attribute: `region`

**Import ID:** `<region>/CredentialProviderArn` (AWS::BedrockAgentCore::ApiKeyCredentialProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiKey` | api_key | `string` | optional, computed, provider-chosen, write-only |  | The API key to use for authentication |
| `ApiKeySecretArn` | api_key_secret_arn | `map` | computed |  | Contains information about the API key secret in AWS Secrets Manager |
| `ApiKeySecretConfig` | api_key_secret_config | `map` | optional, computed, provider-chosen, write-only |  | A reference to a customer-provided secret stored in AWS Secrets Manager |
| `ApiKeySecretJsonKey` | api_key_secret_json_key | `string` | computed |  | The JSON key within the secret that contains the API key value |
| `ApiKeySecretSource` | api_key_secret_source | `string` | optional, computed, provider-chosen, write-only |  | The source of the API key secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets. |
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the credential provider was created |
| `CredentialProviderArn` | credential_provider_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the API key credential provider |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the credential provider was last updated |
| `Name` |  | `string` | required, replaces on change |  | The name of the API key credential provider |
| `Tags` |  | `map` | tags map |  | Tags to assign to the API key credential provider |

Supports update: yes

Discovery: supported
