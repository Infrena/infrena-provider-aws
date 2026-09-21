# aws.certificateauthorityactivation

**CloudFormation type:** `AWS::ACMPCA::CertificateAuthorityActivation`

Used to install the certificate authority certificate and update the certificate authority status.

Region attribute: `region`

**Import ID:** `<region>/CertificateAuthorityArn` (AWS::ACMPCA::CertificateAuthorityActivation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Certificate` |  | `string` | required, write-only |  | Certificate Authority certificate that will be installed in the Certificate Authority. |
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | required, replaces on change | aws.acmpca.certificateauthority.Arn | Arn of the Certificate Authority. |
| `CertificateChain` | certificate_chain | `string` | optional, computed, provider-chosen, write-only |  | Certificate chain for the Certificate Authority certificate. |
| `CompleteCertificateChain` | complete_certificate_chain | `string` | computed |  | The complete certificate chain, including the Certificate Authority certificate. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the Certificate Authority. |

Supports update: yes

Discovery: not supported
