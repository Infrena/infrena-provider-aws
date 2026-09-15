# aws.acmedomainvalidation

**CloudFormation type:** `AWS::CertificateManager::AcmeDomainValidation`

Resource Type definition for AWS::CertificateManager::AcmeDomainValidation

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CertificateManager::AcmeDomainValidation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcmeEndpointArn` | acme_endpoint_arn | `string` | required, replaces on change | aws.acmeendpoint.AcmeEndpointArn | The ARN of the ACME endpoint this domain validation is associated with. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the domain validation. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The domain name to validate. |
| `PrevalidationOptions` | prevalidation_options | `map` | required |  | Prevalidation method configuration. Currently only DNS-based prevalidation is supported. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with the domain validation. |

Supports update: yes

Discovery: supported (parent resource required)
