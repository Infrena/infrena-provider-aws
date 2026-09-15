# aws.harnessendpoint

**CloudFormation type:** `AWS::BedrockAgentCore::HarnessEndpoint`

Resource Type definition for AWS::BedrockAgentCore::HarnessEndpoint - a named, stable reference to a specific version of a Harness that callers invoke, allowing the underlying version to be updated without changing how the agent is invoked.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BedrockAgentCore::HarnessEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the endpoint. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the endpoint was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the endpoint. |
| `EndpointName` | endpoint_name | `string` | required, replaces on change |  | The name of the endpoint. Must start with a letter and contain only alphanumeric characters and underscores. |
| `HarnessId` | harness_id | `string` | required, replaces on change | aws.harness.HarnessId | The ID of the harness that the endpoint belongs to. |
| `HarnessName` | harness_name | `string` | computed |  | The name of the harness that the endpoint belongs to. |
| `LiveVersion` | live_version | `string` | computed |  | The harness version that the endpoint is currently serving. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the harness endpoint resource. |
| `TargetVersion` | target_version | `string` | optional, computed, provider-chosen, write-only |  | The harness version that the endpoint points to and serves invocations from. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the endpoint was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
