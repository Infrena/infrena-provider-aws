# aws.certificatemanager.certificate

**CloudFormation type:** `AWS::CertificateManager::Certificate`

Resource Type definition for AWS::CertificateManager::Certificate

Region attribute: `region`

**Import ID:** `<region>/CertificateArn` (AWS::CertificateManager::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate. |
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate. |
| `CertificateExport` | certificate_export | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the certificate can be exported. ENABLED allows the certificate to be exported, DISABLED prevents export. |
| `CertificateTransparencyLoggingPreference` | certificate_transparency_logging_preference | `string` | optional, computed, provider-chosen |  | You can opt out of certificate transparency logging by specifying the DISABLED option. Opt in by specifying ENABLED. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate |
| `DomainValidationOptions` | domain_validation_options | `list` | optional, computed, provider-chosen, replaces on change |  | Domain information that domain name registrars use to verify your identity. |
| `KeyAlgorithm` | key_algorithm | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data. |
| `SubjectAlternativeNames` | subject_alternative_names | `list` | optional, computed, provider-chosen, replaces on change |  | Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can identify the certificate. |
| `ValidationMethod` | validation_method | `string` | optional, computed, provider-chosen, write-only |  | The method you want to use to validate that you own or control the domain associated with a public certificate. Valid values are DNS, EMAIL or HTTP |

Supports update: yes

Discovery: supported
