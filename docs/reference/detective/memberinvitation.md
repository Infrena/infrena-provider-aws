# aws.memberinvitation

**CloudFormation type:** `AWS::Detective::MemberInvitation`

Resource schema for AWS::Detective::MemberInvitation

Region attribute: `region`

**Import ID:** `<region>/GraphArn|MemberId` (AWS::Detective::MemberInvitation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DisableEmailNotification` | disable_email_notification | `boolean` | optional, computed, provider-chosen, write-only |  | When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect. |
| `GraphArn` | graph_arn | `string` | required, replaces on change | aws.detective.graph.Arn | The ARN of the graph to which the member account will be invited |
| `MemberEmailAddress` | member_email_address | `string` | required |  | The root email address for the account to be invited, for validation. Updating this field has no effect. |
| `MemberId` | member_id | `string` | required, replaces on change |  | The AWS account ID to be invited to join the graph as a member |
| `Message` |  | `string` | optional, computed, provider-chosen, write-only |  | A message to be included in the email invitation sent to the invited account. Updating this field has no effect. |

Supports update: yes

Discovery: supported
