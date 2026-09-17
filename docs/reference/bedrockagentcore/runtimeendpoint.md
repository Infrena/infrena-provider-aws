# aws.runtimeendpoint

**CloudFormation type:** `AWS::BedrockAgentCore::RuntimeEndpoint`

Resource definition for AWS::BedrockAgentCore::RuntimeEndpoint

Region attribute: `region`

**Import ID:** `<region>/AgentRuntimeEndpointArn` (AWS::BedrockAgentCore::RuntimeEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentRuntimeArn` | agent_runtime_arn | `string` | computed |  | The ARN of the Agent Runtime |
| `AgentRuntimeEndpointArn` | agent_runtime_endpoint_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the AgentCore Runtime. |
| `AgentRuntimeId` | agent_runtime_id | `string` | required, replaces on change | aws.runtime.AgentRuntimeId | The ID of the parent Agent Runtime |
| `AgentRuntimeVersion` | agent_runtime_version | `string` | optional, computed, provider-chosen |  | The version of the AgentCore Runtime to use for the endpoint. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the Agent Runtime Endpoint was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the AgentCore Runtime endpoint. |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for failure if the endpoint is in a failed state |
| `Id` |  | `string` | computed |  | The unique ID of the Agent Runtime Endpoint itself |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp when the Agent Runtime Endpoint was last updated |
| `LiveVersion` | live_version | `string` | computed |  | The Live version of the Agent Runtime |
| `Name` |  | `string` | required, replaces on change |  | The name of the Agent Runtime Endpoint |
| `Status` |  | `string` | computed |  | The status of the Agent Runtime Endpoint |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `TargetVersion` | target_version | `string` | computed |  | The target version of the AgentCore Runtime for the endpoint. |

Supports update: yes

Discovery: supported (parent resource required)
