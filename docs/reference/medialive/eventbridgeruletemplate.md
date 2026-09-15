# aws.eventbridgeruletemplate

**CloudFormation type:** `AWS::MediaLive::EventBridgeRuleTemplate`

Definition of AWS::MediaLive::EventBridgeRuleTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/Identifier` (AWS::MediaLive::EventBridgeRuleTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | An eventbridge rule template's ARN (Amazon Resource Name) |
| `CreatedAt` | created_at | `string` | computed |  | Placeholder documentation for __timestampIso8601 |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A resource's optional description. |
| `EventTargets` | event_targets | `list` | optional, computed, provider-chosen |  | Placeholder documentation for __listOfEventBridgeRuleTemplateTarget |
| `EventType` | event_type | `string` | required |  | The type of event to match with the rule. |
| `GroupId` | group_id | `string` | computed |  | An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-` |
| `GroupIdentifier` | group_identifier | `string` | optional, computed, provider-chosen, write-only |  | An eventbridge rule template group's identifier. Can be either be its id or current name. |
| `Id` |  | `string` | computed |  | An eventbridge rule template's id. AWS provided templates have ids that start with `aws-` |
| `Identifier` |  | `string` | computed |  | Placeholder documentation for __string |
| `ModifiedAt` | modified_at | `string` | computed |  | Placeholder documentation for __timestampIso8601 |
| `Name` |  | `string` | required |  | A resource's name. Names must be unique within the scope of a resource type in a specific region. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Represents the tags associated with a resource. |

Supports update: yes

Discovery: supported
