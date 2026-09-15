# aws.iot.certificate

**CloudFormation type:** `AWS::IoT::Certificate`

Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoT::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CACertificatePem` | ca_certificate_pem | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CertificateMode` | certificate_mode | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CertificatePem` | certificate_pem | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CertificateSigningRequest` | certificate_signing_request | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Id` |  | `string` | computed |  |  |
| `Status` |  | `string` | required |  |  |

Supports update: yes

Discovery: supported
