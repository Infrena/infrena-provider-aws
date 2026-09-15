# aws.onlineevaluationconfig

**CloudFormation type:** `AWS::BedrockAgentCore::OnlineEvaluationConfig`

Resource Type definition for AWS::BedrockAgentCore::OnlineEvaluationConfig - Creates an online evaluation configuration for continuous monitoring of agent performance.

Region attribute: `region`

**Import ID:** `<region>/OnlineEvaluationConfigArn` (AWS::BedrockAgentCore::OnlineEvaluationConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusteringConfig` | clustering_config | `map` | optional, computed, provider-chosen |  | The configuration for clustering analysis of evaluation results. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the online evaluation configuration was created. |
| `DataSourceConfig` | data_source_config | `map` | required |  | The configuration that specifies where to read agent traces for online evaluation. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the online evaluation configuration. |
| `EvaluationExecutionRoleArn` | evaluation_execution_role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role that grants permissions for evaluation. |
| `Evaluators` |  | `list` | optional, computed, provider-chosen |  | The list of evaluators to apply during online evaluation. |
| `ExecutionStatus` | execution_status | `string` | optional, computed, provider-chosen |  | The execution status indicating whether the online evaluation is currently running. |
| `Insights` |  | `list` | optional, computed, provider-chosen |  | The list of insights to enable for failure analysis. |
| `OnlineEvaluationConfigArn` | online_evaluation_config_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the online evaluation configuration. |
| `OnlineEvaluationConfigId` | online_evaluation_config_id | `string` | computed |  | The unique identifier of the online evaluation configuration. |
| `OnlineEvaluationConfigName` | online_evaluation_config_name | `string` | required, replaces on change |  | The name of the online evaluation configuration. Must be unique within your account. |
| `OutputConfig` | output_config | `map` | computed |  | The configuration that specifies where evaluation results should be written. |
| `Rule` |  | `map` | required |  | The evaluation rule that defines sampling configuration, filtering criteria, and session detection settings. |
| `Status` |  | `string` | computed |  | The status of the online evaluation configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to assign to the online evaluation configuration. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the online evaluation configuration was last updated. |

Supports update: yes

Discovery: supported
