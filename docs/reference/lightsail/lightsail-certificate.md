# aws.lightsail.certificate

**CloudFormation type:** `AWS::Lightsail::Certificate`

Resource Type definition for AWS::Lightsail::Certificate.

Region attribute: `region`

**Import ID:** `<region>/CertificateName` (AWS::Lightsail::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | computed |  |  |
| `CertificateName` | certificate_name | `string` | required, replaces on change |  | The name for the certificate. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The domain name (e.g., example.com ) for the certificate. |
| `Status` |  | `string` | computed |  | The validation status of the certificate. |
| `SubjectAlternativeNames` | subject_alternative_names | `list` | optional, computed, provider-chosen, replaces on change |  | An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
