# aws.dms.certificate

**CloudFormation type:** `AWS::DMS::Certificate`

Resource Type definition for AWS::DMS::Certificate

Region attribute: `region`

**Import ID:** `<region>/CertificateArn` (AWS::DMS::Certificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | computed |  | The certificate Arn |
| `CertificateIdentifier` | certificate_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The certificate Identifier |
| `CertificatePem` | certificate_pem | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The certificate Pem |
| `CertificateWallet` | certificate_wallet | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The certificate Wallet |

Supports update: no

Discovery: supported
