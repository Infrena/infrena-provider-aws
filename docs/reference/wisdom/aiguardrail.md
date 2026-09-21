# aws.aiguardrail

**CloudFormation type:** `AWS::Wisdom::AIGuardrail`

Definition of AWS::Wisdom::AIGuardrail Resource Type

Region attribute: `region`

**Import ID:** `<region>/AIGuardrailId|AssistantId` (AWS::Wisdom::AIGuardrail)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIGuardrailArn` | ai_guardrail_arn | `string` | computed |  |  |
| `AIGuardrailId` | ai_guardrail_id | `string` | computed |  |  |
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `BlockedInputMessaging` | blocked_input_messaging | `string` | required |  | Messaging for when violations are detected in text |
| `BlockedOutputsMessaging` | blocked_outputs_messaging | `string` | required |  | Messaging for when violations are detected in text |
| `ContentPolicyConfig` | content_policy_config | `map` | optional, computed, provider-chosen |  | Content policy config for a guardrail. |
| `ContextualGroundingPolicyConfig` | contextual_grounding_policy_config | `map` | optional, computed, provider-chosen |  | Contextual grounding policy config for a guardrail. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the guardrail or its version |
| `ModifiedTimeSeconds` | modified_time_seconds | `float` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SensitiveInformationPolicyConfig` | sensitive_information_policy_config | `map` | optional, computed, provider-chosen |  | Sensitive information policy config for a guardrail. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `TopicPolicyConfig` | topic_policy_config | `map` | optional, computed, provider-chosen |  | Topic policy config for a guardrail. |
| `WordPolicyConfig` | word_policy_config | `map` | optional, computed, provider-chosen |  | Word policy config for a guardrail. |

Supports update: yes

Discovery: supported (parent resource required)
