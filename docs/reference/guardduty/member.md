# aws.member

**CloudFormation type:** `AWS::GuardDuty::Member`

Resource Type definition for AWS::GuardDuty::Member

Region attribute: `region`

**Import ID:** `<region>/DetectorId|MemberId` (AWS::GuardDuty::Member)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DetectorId` | detector_id | `string` | optional, computed, provider-chosen, replaces on change | aws.guardduty.detector.Id |  |
| `DisableEmailNotification` | disable_email_notification | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `Email` |  | `string` | required |  |  |
| `MemberId` | member_id | `string` | optional, computed, provider-chosen, replaces on change | aws.member.MemberId |  |
| `Message` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
