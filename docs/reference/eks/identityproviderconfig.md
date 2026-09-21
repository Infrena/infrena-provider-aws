# aws.identityproviderconfig

**CloudFormation type:** `AWS::EKS::IdentityProviderConfig`

An object representing an Amazon EKS IdentityProviderConfig.

Region attribute: `region`

**Import ID:** `<region>/IdentityProviderConfigName|ClusterName|Type` (AWS::EKS::IdentityProviderConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | The name of the identity provider configuration. |
| `IdentityProviderConfigArn` | identity_provider_config_arn | `string` | computed |  | The ARN of the configuration. |
| `IdentityProviderConfigName` | identity_provider_config_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the OIDC provider configuration. |
| `Oidc` |  | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing an OpenID Connect (OIDC) configuration. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the identity provider configuration. |

Supports update: yes

Discovery: supported (parent resource required)
