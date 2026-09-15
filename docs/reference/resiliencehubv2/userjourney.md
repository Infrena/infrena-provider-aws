# aws.userjourney

**CloudFormation type:** `AWS::ResilienceHubV2::UserJourney`

Creates a user journey within a Resilience Hub system.

Region attribute: `region`

**Import ID:** `<region>/SystemIdentifier|UserJourneyId` (AWS::ResilienceHubV2::UserJourney)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the user journey was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the user journey. |
| `Name` |  | `string` | required, replaces on change |  | The name of the user journey. |
| `PolicyArn` | policy_arn | `string` | optional, computed, provider-chosen | aws.resiliencehubv2.policy.PolicyArn | The ARN of the resilience policy to associate with this user journey. |
| `SystemIdentifier` | system_identifier | `string` | required, replaces on change |  | The system ARN or system ID that owns this user journey. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the user journey was last updated. |
| `UserJourneyId` | user_journey_id | `string` | computed |  | The server-generated user journey ID. |

Supports update: yes

Discovery: supported (parent resource required)
