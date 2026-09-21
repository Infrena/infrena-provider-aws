# aws.cacertificate

**CloudFormation type:** `AWS::IoT::CACertificate`

Registers a CA Certificate in IoT.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoT::CACertificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AutoRegistrationStatus` | auto_registration_status | `string` | optional, computed, provider-chosen |  |  |
| `CACertificatePem` | ca_certificate_pem | `string` | required, replaces on change |  |  |
| `CertificateMode` | certificate_mode | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `RegistrationConfig` | registration_config | `map` | optional, computed, provider-chosen |  |  |
| `RemoveAutoRegistration` | remove_auto_registration | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `Status` |  | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `VerificationCertificatePem` | verification_certificate_pem | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The private key verification certificate. |

Supports update: yes

Discovery: supported
