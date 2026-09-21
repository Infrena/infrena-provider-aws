# aws.targetdomain

**CloudFormation type:** `AWS::SecurityAgent::TargetDomain`

Resource Type definition for AWS::SecurityAgent::TargetDomain

Region attribute: `region`

**Import ID:** `<region>/TargetDomainId` (AWS::SecurityAgent::TargetDomain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the target domain was registered |
| `Tags` |  | `map` | tags map |  | Tags for the target domain |
| `TargetDomainId` | target_domain_id | `string` | computed |  | Unique identifier of the target domain |
| `TargetDomainName` | target_domain_name | `string` | required, replaces on change |  | Domain name of the target domain |
| `VerificationDetails` | verification_details | `map` | computed |  | Verification details to verify registered target domain |
| `VerificationMethod` | verification_method | `string` | required |  | Verification method for the target domain |
| `VerificationStatus` | verification_status | `string` | computed |  | Current verification status of the registered target domain |
| `VerificationStatusReason` | verification_status_reason | `string` | computed |  | Reason for the current target domain verification status |
| `VerifiedAt` | verified_at | `string` | computed |  | Timestamp when the target domain was last successfully verified |

Supports update: yes

Discovery: supported
