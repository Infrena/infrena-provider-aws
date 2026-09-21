# aws.aiprompt

**CloudFormation type:** `AWS::Wisdom::AIPrompt`

Definition of AWS::Wisdom::AIPrompt Resource Type

Region attribute: `region`

**Import ID:** `<region>/AIPromptId|AssistantId` (AWS::Wisdom::AIPrompt)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIPromptArn` | ai_prompt_arn | `string` | computed |  |  |
| `AIPromptId` | ai_prompt_id | `string` | computed |  |  |
| `ApiFormat` | api_format | `string` | required, replaces on change |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | optional, computed, provider-chosen, replaces on change | aws.assistant.AssistantId |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ModelId` | model_id | `string` | required |  |  |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `TemplateConfiguration` | template_configuration | `map` | required |  |  |
| `TemplateType` | template_type | `string` | required, replaces on change |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
