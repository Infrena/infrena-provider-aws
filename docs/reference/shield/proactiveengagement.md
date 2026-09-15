# aws.proactiveengagement

**CloudFormation type:** `AWS::Shield::ProactiveEngagement`

Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::Shield::ProactiveEngagement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  |  |
| `EmergencyContactList` | emergency_contact_list | `list` | required |  | A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support. |
| `ProactiveEngagementStatus` | proactive_engagement_status | `string` | required |  | If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support. |

Supports update: yes

Discovery: supported
