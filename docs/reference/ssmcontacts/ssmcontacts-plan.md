# aws.ssmcontacts.plan

**CloudFormation type:** `AWS::SSMContacts::Plan`

Engagement Plan for a SSM Incident Manager Contact.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMContacts::Plan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the contact. |
| `ContactId` | contact_id | `string` | optional, computed, provider-chosen, replaces on change |  | Contact ID for the AWS SSM Incident Manager Contact to associate the plan. |
| `RotationIds` | rotation_ids | `list` | optional, computed, provider-chosen, write-only |  | Rotation Ids to associate with Oncall Contact for engagement. |
| `Stages` |  | `list` | optional, computed, provider-chosen |  | The stages that an escalation plan or engagement plan engages contacts and contact methods in. |

Supports update: yes

Discovery: supported
