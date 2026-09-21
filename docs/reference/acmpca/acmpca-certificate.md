# aws.acmpca.certificate

**CloudFormation type:** `AWS::ACMPCA::Certificate`

The ``AWS::ACMPCA::Certificate`` resource is used to issue a certificate using your private certificate authority. For more information, see the [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html) action.

Region attribute: `region`

**Import ID:** `<region>/Arn|CertificateAuthorityArn` (AWS::ACMPCA::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiPassthrough` | api_passthrough | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Contains X.509 certificate information to be placed in an issued certificate. An ``APIPassthrough`` or ``APICSRPassthrough`` template variant must be selected, or else this parameter is ignored. |
| `Arn` |  | `string` | computed |  |  |
| `Certificate` |  | `string` | computed |  |  |
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | required, replaces on change | aws.acmpca.certificateauthority.Arn | The Amazon Resource Name (ARN) for the private CA issues the certificate. |
| `CertificateSigningRequest` | certificate_signing_request | `string` | required, replaces on change, write-only |  | The certificate signing request (CSR) for the certificate. |
| `SigningAlgorithm` | signing_algorithm | `string` | required, replaces on change, write-only |  | The name of the algorithm that will be used to sign the certificate to be issued. |
| `TemplateArn` | template_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, PCAshort defaults to the ``EndEntityCertificate/V1`` template. For more information about PCAshort templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html). |
| `Validity` |  | `map` | required, replaces on change, write-only |  | Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years. You can issue a certificate by calling the ``IssueCertificate`` operation. |
| `ValidityNotBefore` | validity_not_before | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years. You can issue a certificate by calling the ``IssueCertificate`` operation. |

Supports update: no

Discovery: not supported
