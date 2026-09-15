# aws.acmpca.certificateauthority

**CloudFormation type:** `AWS::ACMPCA::CertificateAuthority`

Private certificate authority.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ACMPCA::CertificateAuthority)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the certificate authority. |
| `CertificateSigningRequest` | certificate_signing_request | `string` | computed |  | The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate. |
| `CsrExtensions` | csr_extensions | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Structure that contains CSR pass though extensions information. |
| `KeyAlgorithm` | key_algorithm | `string` | required, replaces on change |  | Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate. |
| `KeyStorageSecurityStandard` | key_storage_security_standard | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys. |
| `RevocationConfiguration` | revocation_configuration | `map` | optional, computed, provider-chosen, write-only |  | Certificate Authority revocation information. |
| `SigningAlgorithm` | signing_algorithm | `string` | required, replaces on change |  | Algorithm your CA uses to sign certificate requests. |
| `Subject` |  | `map` | required, replaces on change, write-only |  | Structure that contains X.500 distinguished name information for your CA. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the certificate authority. |
| `UsageMode` | usage_mode | `string` | optional, computed, provider-chosen, replaces on change |  | Usage mode of the ceritificate authority. |

Supports update: yes

Discovery: supported
