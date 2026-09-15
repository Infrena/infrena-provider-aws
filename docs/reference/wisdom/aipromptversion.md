# aws.aipromptversion

**CloudFormation type:** `AWS::Wisdom::AIPromptVersion`

Definition of AWS::Wisdom::AIPromptVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssistantId|AIPromptId|VersionNumber` (AWS::Wisdom::AIPromptVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIPromptArn` | ai_prompt_arn | `string` | computed |  |  |
| `AIPromptId` | ai_prompt_id | `string` | required, replaces on change | aws.aiprompt.AIPromptId |  |
| `AIPromptVersionId` | ai_prompt_version_id | `string` | computed |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | optional, computed, provider-chosen, replaces on change |  |  |
| `VersionNumber` | version_number | `float` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
