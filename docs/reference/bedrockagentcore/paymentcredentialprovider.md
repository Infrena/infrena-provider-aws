# aws.paymentcredentialprovider

**CloudFormation type:** `AWS::BedrockAgentCore::PaymentCredentialProvider`

Resource Type definition for AWS::BedrockAgentCore::PaymentCredentialProvider

Region attribute: `region`

**Import ID:** `<region>/CredentialProviderArn` (AWS::BedrockAgentCore::PaymentCredentialProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTime` | created_time | `string` | computed |  | The timestamp when the credential provider was created |
| `CredentialProviderArn` | credential_provider_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the payment credential provider |
| `CredentialProviderVendor` | credential_provider_vendor | `string` | required, replaces on change |  | Supported vendor types for payment providers |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the credential provider was last updated |
| `Name` |  | `string` | required, replaces on change |  | Unique name for the payment credential provider |
| `ProviderConfigurationInput` | provider_configuration_input | `map` | optional, computed, provider-chosen, write-only |  | Provider configuration input containing secrets for creation/update |
| `ProviderConfigurationOutput` | provider_configuration_output | `map` | computed |  | Provider configuration output containing secret ARNs (no raw secrets) |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the payment credential provider |

Supports update: yes

Discovery: supported
