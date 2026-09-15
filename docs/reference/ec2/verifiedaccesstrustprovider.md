# aws.verifiedaccesstrustprovider

**CloudFormation type:** `AWS::EC2::VerifiedAccessTrustProvider`

The AWS::EC2::VerifiedAccessTrustProvider type describes a verified access trust provider

Region attribute: `region`

**Import ID:** `<region>/VerifiedAccessTrustProviderId` (AWS::EC2::VerifiedAccessTrustProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The creation time. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the Amazon Web Services Verified Access trust provider. |
| `DeviceOptions` | device_options | `map` | optional, computed, provider-chosen, replaces on change |  | The options for device identity based trust providers. |
| `DeviceTrustProviderType` | device_trust_provider_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of device-based trust provider. Possible values: jamf\|crowdstrike |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The last updated time. |
| `NativeApplicationOidcOptions` | native_application_oidc_options | `map` | optional, computed, provider-chosen |  | The OpenID Connect details for an oidc -type, user-identity based trust provider for L4. |
| `OidcOptions` | oidc_options | `map` | optional, computed, provider-chosen |  | The OpenID Connect details for an oidc -type, user-identity based trust provider. |
| `PolicyReferenceName` | policy_reference_name | `string` | required, replaces on change |  | The identifier to be used when working with policy rules. |
| `SseSpecification` | sse_specification | `map` | optional, computed, provider-chosen |  | The configuration options for customer provided KMS encryption. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TrustProviderType` | trust_provider_type | `string` | required, replaces on change |  | Type of trust provider. Possible values: user\|device |
| `UserTrustProviderType` | user_trust_provider_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of device-based trust provider. Possible values: oidc\|iam-identity-center |
| `VerifiedAccessTrustProviderId` | verified_access_trust_provider_id | `string` | computed |  | The ID of the Amazon Web Services Verified Access trust provider. |

Supports update: yes

Discovery: supported
