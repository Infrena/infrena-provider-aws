# aws.transfer.certificate

**CloudFormation type:** `AWS::Transfer::Certificate`

Resource Type definition for AWS::Transfer::Certificate

Region attribute: `region`

**Import ID:** `<region>/CertificateId` (AWS::Transfer::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActiveDate` | active_date | `string` | optional, computed, provider-chosen |  | Specifies the active date for the certificate. |
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the agreement. |
| `Certificate` |  | `string` | required, replaces on change |  | Specifies the certificate body to be imported. |
| `CertificateChain` | certificate_chain | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the certificate chain to be imported. |
| `CertificateId` | certificate_id | `string` | computed |  | A unique identifier for the certificate. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A textual description for the certificate. |
| `InactiveDate` | inactive_date | `string` | optional, computed, provider-chosen |  | Specifies the inactive date for the certificate. |
| `NotAfterDate` | not_after_date | `string` | computed |  | Specifies the not after date for the certificate. |
| `NotBeforeDate` | not_before_date | `string` | computed |  | Specifies the not before date for the certificate. |
| `PrivateKey` | private_key | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  | Specifies the private key for the certificate. |
| `Serial` |  | `string` | computed |  | Specifies Certificate's serial. |
| `Status` |  | `string` | computed |  | A status description for the certificate. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose. |
| `Type` | type_value | `string` | computed |  | Describing the type of certificate. With or without a private key. |
| `Usage` |  | `string` | required |  | Specifies the usage type for the certificate. |

Supports update: yes

Discovery: supported
