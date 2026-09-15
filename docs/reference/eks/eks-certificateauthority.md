# aws.eks.certificateauthority

**CloudFormation type:** `AWS::EKS::CertificateAuthority`

Resource Type definition for EKS CertificateAuthority.

Region attribute: `region`

**Import ID:** `<region>/ClusterName|Id` (AWS::EKS::CertificateAuthority)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActivatedAt` | activated_at | `string` | computed |  | The timestamp when the certificate authority was activated. |
| `ActivatedBy` | activated_by | `string` | computed |  | The entity that activated the certificate authority. |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | The name of the EKS cluster that the certificate authority belongs to. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the certificate authority was created. |
| `CreatedBy` | created_by | `string` | computed |  | The entity that created the certificate authority. |
| `Data` |  | `string` | computed |  | The Base64 encoded certificate authority data. |
| `DistributionStatus` | distribution_status | `string` | computed |  | The distribution status of the certificate authority. |
| `Id` |  | `string` | computed |  | The unique identifier of the certificate authority. |
| `RollbackAvailable` | rollback_available | `boolean` | computed |  | Whether activation of this certificate authority can still be rolled back. |
| `ScheduledEvents` | scheduled_events | `map` | computed |  | The scheduled auto-activation events for the certificate authority, computed from its validity window. |
| `SigningStatus` | signing_status | `string` | computed |  | The signing status of the certificate authority. |
| `Validity` |  | `map` | computed |  | The validity period of the certificate authority. |

Supports update: no

Discovery: supported (parent resource required)
