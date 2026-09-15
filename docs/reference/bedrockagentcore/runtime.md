# aws.runtime

**CloudFormation type:** `AWS::BedrockAgentCore::Runtime`

Resource Type definition for AWS::BedrockAgentCore::Runtime

Region attribute: `region`

**Import ID:** `<region>/AgentRuntimeId` (AWS::BedrockAgentCore::Runtime)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentRuntimeArn` | agent_runtime_arn | `string` | computed |  | The Amazon Resource Name(ARN) that uniquely identifies the Agent |
| `AgentRuntimeArtifact` | agent_runtime_artifact | `map` | required |  | The artifact of the agent |
| `AgentRuntimeId` | agent_runtime_id | `string` | computed |  | Identifier for a resource |
| `AgentRuntimeName` | agent_runtime_name | `string` | required, replaces on change |  | Name for a resource |
| `AgentRuntimeVersion` | agent_runtime_version | `string` | computed |  | Version of the Agent |
| `AuthorizerConfiguration` | authorizer_configuration | `map` | optional, computed, provider-chosen |  | Configuration for the authorizer |
| `CapacityProviderConfiguration` | capacity_provider_configuration | `map` | optional, computed, provider-chosen |  | Configuration for a capacity provider |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the Agent was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the resource |
| `EnvironmentVariables` | environment_variables | `map` | optional, computed, provider-chosen |  | Environment variable attributes |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for failure if the agent is in a failed state. |
| `FilesystemConfigurations` | filesystem_configurations | `list` | optional, computed, provider-chosen |  | List of filesystem configurations |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | When resource was last updated |
| `LifecycleConfiguration` | lifecycle_configuration | `map` | optional, computed, provider-chosen |  | Configuration for managing the lifecycle of runtime sessions and resources |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  | Network access configuration for the Agent |
| `ProtocolConfiguration` | protocol_configuration | `string` | optional, computed, provider-chosen |  | Protocol configuration for the agent runtime |
| `RequestHeaderConfiguration` | request_header_configuration | `map` | optional, computed, provider-chosen |  | Configuration for HTTP request headers |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Amazon Resource Name (ARN) of an IAM role |
| `Status` |  | `string` | computed |  | Current status of the agent |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `WorkloadIdentityDetails` | workload_identity_details | `map` | computed |  | Configuration for workload identity |

Supports update: yes

Discovery: supported
