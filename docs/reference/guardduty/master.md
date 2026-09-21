# aws.master

**CloudFormation type:** `AWS::GuardDuty::Master`

GuardDuty Master resource schema

Region attribute: `region`

**Import ID:** `<region>/DetectorId|MasterId` (AWS::GuardDuty::Master)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DetectorId` | detector_id | `string` | required, replaces on change | aws.guardduty.detector.Id | Unique ID of the detector of the GuardDuty member account. |
| `InvitationId` | invitation_id | `string` | optional, computed, provider-chosen, replaces on change |  | Value used to validate the master account to the member account. |
| `MasterId` | master_id | `string` | required, replaces on change | aws.master.MasterId | ID of the account used as the master account. |

Supports update: no

Discovery: supported
