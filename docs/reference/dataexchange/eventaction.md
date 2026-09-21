# aws.eventaction

**CloudFormation type:** `AWS::DataExchange::EventAction`

An event action is an AWS Data Exchange resource that automatically exports data set revisions to Amazon S3 when a revision is published.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::DataExchange::EventAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `map` | required |  | What occurs after a certain event. |
| `Arn` |  | `string` | computed |  | The ARN for the event action. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the event action was created, in ISO 8601 format. |
| `Event` |  | `map` | required, replaces on change |  | What occurs to start an action. |
| `EventActionId` | event_action_id | `string` | computed |  | The unique identifier for the event action. |
| `Tags` |  | `map` | tags map |  | The tags for the event action. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time that the event action was last updated, in ISO 8601 format. |

Supports update: yes

Discovery: supported
