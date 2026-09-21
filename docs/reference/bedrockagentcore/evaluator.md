# aws.evaluator

**CloudFormation type:** `AWS::BedrockAgentCore::Evaluator`

Resource Type definition for AWS::BedrockAgentCore::Evaluator - Creates a custom evaluator for agent quality assessment using LLM-as-a-Judge configurations.

Region attribute: `region`

**Import ID:** `<region>/EvaluatorArn` (AWS::BedrockAgentCore::Evaluator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the evaluator was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the evaluator. |
| `EvaluatorArn` | evaluator_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the evaluator. |
| `EvaluatorConfig` | evaluator_config | `map` | required |  | The configuration that defines how an evaluator assesses agent performance. |
| `EvaluatorId` | evaluator_id | `string` | computed |  | The unique identifier of the evaluator. |
| `EvaluatorName` | evaluator_name | `string` | required, replaces on change |  | The name of the evaluator. Must be unique within your account. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The ARN of the KMS key used to encrypt evaluator data. |
| `Level` |  | `string` | required |  | The evaluation level that determines the scope of evaluation. |
| `Status` |  | `string` | computed |  | The current status of the evaluator. |
| `Tags` |  | `map` | tags map |  | A list of tags to assign to the evaluator. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the evaluator was last updated. |

Supports update: yes

Discovery: supported
