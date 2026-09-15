# aws.aiagentversion

**CloudFormation type:** `AWS::Wisdom::AIAgentVersion`

Definition of AWS::Wisdom::AIAgentVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssistantId|AIAgentId|VersionNumber` (AWS::Wisdom::AIAgentVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIAgentArn` | ai_agent_arn | `string` | computed |  |  |
| `AIAgentId` | ai_agent_id | `string` | required, replaces on change | aws.aiagent.AIAgentId |  |
| `AIAgentVersionId` | ai_agent_version_id | `string` | computed |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | optional, computed, provider-chosen, replaces on change |  |  |
| `VersionNumber` | version_number | `float` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
