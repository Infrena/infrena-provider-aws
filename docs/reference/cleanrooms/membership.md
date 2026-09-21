# aws.membership

**CloudFormation type:** `AWS::CleanRooms::Membership`

Represents an AWS account that is a part of a collaboration

Region attribute: `region`

**Import ID:** `<region>/MembershipIdentifier` (AWS::CleanRooms::Membership)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationCreatorAccountId` | collaboration_creator_account_id | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | required, replaces on change |  |  |
| `DefaultJobResultConfiguration` | default_job_result_configuration | `map` | optional, computed, provider-chosen |  |  |
| `DefaultResultConfiguration` | default_result_configuration | `map` | optional, computed, provider-chosen |  |  |
| `IsMetricsEnabled` | is_metrics_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `JobLogStatus` | job_log_status | `string` | optional, computed, provider-chosen |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | computed |  |  |
| `PaymentConfiguration` | payment_configuration | `map` | optional, computed, provider-chosen |  |  |
| `QueryLogStatus` | query_log_status | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms membership. |

Supports update: yes

Discovery: supported
