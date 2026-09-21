# aws.aiagent

**CloudFormation type:** `AWS::Wisdom::AIAgent`

Definition of AWS::Wisdom::AIAgent Resource Type

Region attribute: `region`

**Import ID:** `<region>/AIAgentId|AssistantId` (AWS::Wisdom::AIAgent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIAgentArn` | ai_agent_arn | `string` | computed |  |  |
| `AIAgentId` | ai_agent_id | `string` | computed |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `Configuration` |  | `string` | required |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
