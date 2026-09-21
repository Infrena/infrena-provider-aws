# aws.acmpca.permission

**CloudFormation type:** `AWS::ACMPCA::Permission`

Permission set on private certificate authority

Region attribute: `region`

**Import ID:** `<region>/CertificateAuthorityArn|Principal` (AWS::ACMPCA::Permission)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required, replaces on change |  | The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided. |
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | required, replaces on change | aws.acmpca.certificateauthority.Arn | The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission. |
| `Principal` |  | `string` | required, replaces on change |  | The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com. |
| `SourceAccount` | source_account | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the calling account. |

Supports update: no

Discovery: not supported
