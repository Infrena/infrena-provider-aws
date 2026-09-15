# aws.guardrail

**CloudFormation type:** `AWS::Bedrock::Guardrail`

Definition of AWS::Bedrock::Guardrail Resource Type

Region attribute: `region`

**Import ID:** `<region>/GuardrailArn` (AWS::Bedrock::Guardrail)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutomatedReasoningPolicyConfig` | automated_reasoning_policy_config | `map` | optional, computed, provider-chosen |  | Optional configuration for integrating Automated Reasoning policies with the guardrail. |
| `BlockedInputMessaging` | blocked_input_messaging | `string` | required |  | Messaging for when violations are detected in text |
| `BlockedOutputsMessaging` | blocked_outputs_messaging | `string` | required |  | Messaging for when violations are detected in text |
| `ContentPolicyConfig` | content_policy_config | `map` | optional, computed, provider-chosen |  | Content policy config for a guardrail. |
| `ContextualGroundingPolicyConfig` | contextual_grounding_policy_config | `map` | optional, computed, provider-chosen |  | Contextual grounding policy config for a guardrail. |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp |
| `CrossRegionConfig` | cross_region_config | `map` | optional, computed, provider-chosen |  | The system-defined guardrail profile that you’re using with your guardrail |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the guardrail or its version |
| `FailureRecommendations` | failure_recommendations | `list` | computed |  | List of failure recommendations |
| `GuardrailArn` | guardrail_arn | `string` | computed |  | Arn representation for the guardrail |
| `GuardrailId` | guardrail_id | `string` | computed |  | Unique id for the guardrail |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The KMS key with which the guardrail was encrypted at rest |
| `Name` |  | `string` | required |  | Name of the guardrail |
| `SensitiveInformationPolicyConfig` | sensitive_information_policy_config | `map` | optional, computed, provider-chosen |  | Sensitive information policy config for a guardrail. |
| `Status` |  | `string` | computed |  | Status of the guardrail |
| `StatusReasons` | status_reasons | `list` | computed |  | List of status reasons |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of Tags |
| `TopicPolicyConfig` | topic_policy_config | `map` | optional, computed, provider-chosen |  | Topic policy config for a guardrail. |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp |
| `Version` |  | `string` | computed |  | Guardrail version |
| `WordPolicyConfig` | word_policy_config | `map` | optional, computed, provider-chosen |  | Word policy config for a guardrail. |

Supports update: yes

Discovery: supported
