# aws.bedrockagentcore.capacityprovider

**CloudFormation type:** `AWS::BedrockAgentCore::CapacityProvider`

Resource Type definition for AWS::BedrockAgentCore::CapacityProvider. A capacity provider defines the compute resources (EC2) used to run Amazon Bedrock AgentCore agent runtimes.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BedrockAgentCore::CapacityProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the capacity provider. |
| `CapacityProviderId` | capacity_provider_id | `string` | computed |  | The unique identifier of the capacity provider. |
| `ComputeConfiguration` | compute_configuration | `map` | required, replaces on change |  | The capacity configuration for the capacity provider. Defines the compute resources for this capacity provider. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the capacity provider was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description of the capacity provider. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp when the capacity provider was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the capacity provider. |
| `PermissionsConfiguration` | permissions_configuration | `map` | required, replaces on change |  | Configuration for permissions associated with a capacity provider. |
| `Status` |  | `string` | computed |  | The current status of the capacity provider. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to the capacity provider. |

Supports update: yes

Discovery: supported
