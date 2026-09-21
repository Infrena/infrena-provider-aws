# aws.acmeendpoint

**CloudFormation type:** `AWS::CertificateManager::AcmeEndpoint`

Resource Type definition for AWS::CertificateManager::AcmeEndpoint

Region attribute: `region`

**Import ID:** `<region>/AcmeEndpointArn` (AWS::CertificateManager::AcmeEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcmeEndpointArn` | acme_endpoint_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the ACME endpoint. |
| `AuthorizationBehavior` | authorization_behavior | `string` | required, replaces on change |  | The authorization behavior for the ACME endpoint. |
| `CertificateAuthority` | certificate_authority | `map` | required |  | The certificate authority configuration for the ACME endpoint. |
| `CertificateTags` | certificate_tags | `list` | optional, computed, provider-chosen, replaces on change |  | Tags applied to certificates issued via this endpoint. |
| `Contact` |  | `string` | optional, computed, provider-chosen |  | Whether contact information is required for the ACME endpoint. |
| `EndpointUrl` | endpoint_url | `string` | computed |  | The ACME directory URL for the endpoint. |
| `Tags` |  | `map` | tags map |  | Tags associated with the ACME endpoint. |

Supports update: yes

Discovery: supported
