# aws.harness

**CloudFormation type:** `AWS::BedrockAgentCore::Harness`

Resource Type definition for AWS::BedrockAgentCore::Harness - a managed agentic loop service that provides a turnkey solution for running stateful, tool-equipped AI agents.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BedrockAgentCore::Harness)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedTools` | allowed_tools | `list` | optional, computed, provider-chosen |  | The tools that the agent is allowed to use. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the harness. |
| `AuthorizerConfiguration` | authorizer_configuration | `map` | optional, computed, provider-chosen |  | The inbound authorization configuration for authenticating incoming requests. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the harness was created. |
| `Environment` |  | `map` | optional, computed, provider-chosen |  | The compute environment configuration for the harness, including underlying runtime information. |
| `EnvironmentArtifact` | environment_artifact | `map` | optional, computed, provider-chosen |  | The environment artifact for the harness, such as a custom container image. |
| `EnvironmentVariables` | environment_variables | `map` | optional, computed, provider-chosen |  | Environment variables to set in the harness runtime environment. |
| `ExecutionRoleArn` | execution_role_arn | `string` | required | aws.role.Arn | The ARN of the IAM role that the harness assumes when running. |
| `HarnessId` | harness_id | `string` | computed |  | The unique identifier of the harness. |
| `HarnessName` | harness_name | `string` | required, replaces on change |  | The name of the harness. |
| `MaxIterations` | max_iterations | `integer` | optional, computed, provider-chosen |  | The maximum number of iterations the agent loop can execute per invocation. |
| `MaxTokens` | max_tokens | `integer` | optional, computed, provider-chosen |  | The maximum number of tokens the agent can generate per iteration. |
| `Memory` |  | `map` | optional, computed, provider-chosen |  | The AgentCore Memory configuration for persisting conversation context. |
| `Model` |  | `map` | required |  | The model configuration for the harness. |
| `Skills` |  | `list` | optional, computed, provider-chosen |  | The skills available to the agent. |
| `Status` |  | `string` | computed |  | The current status of the harness. |
| `SystemPrompt` | system_prompt | `list` | optional, computed, provider-chosen |  | The system prompt that defines the agent's behavior. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the harness resource. |
| `TimeoutSeconds` | timeout_seconds | `integer` | optional, computed, provider-chosen |  | The maximum duration in seconds for the agent loop execution per invocation. |
| `Tools` |  | `list` | optional, computed, provider-chosen |  | The tools available to the agent. |
| `Truncation` |  | `map` | optional, computed, provider-chosen |  | The truncation configuration for managing conversation context. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the harness was last updated. |
| `Version` |  | `string` | computed |  | The version of the harness. Incremented on every successful update. |

Supports update: yes

Discovery: supported
