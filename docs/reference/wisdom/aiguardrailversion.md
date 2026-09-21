# aws.aiguardrailversion

**CloudFormation type:** `AWS::Wisdom::AIGuardrailVersion`

Definition of AWS::Wisdom::AIGuardrailVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssistantId|AIGuardrailId|VersionNumber` (AWS::Wisdom::AIGuardrailVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIGuardrailArn` | ai_guardrail_arn | `string` | computed |  |  |
| `AIGuardrailId` | ai_guardrail_id | `string` | required, replaces on change | aws.aiguardrail.AIGuardrailId |  |
| `AIGuardrailVersionId` | ai_guardrail_version_id | `string` | computed |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | optional, computed, provider-chosen, replaces on change |  |  |
| `VersionNumber` | version_number | `float` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
