# aws.devopsagent.agentspace

**CloudFormation type:** `AWS::DevOpsAgent::AgentSpace`

Resource Type definition for AWS::DevOpsAgent::AgentSpace

Region attribute: `region`

**Import ID:** `<region>/AgentSpaceId` (AWS::DevOpsAgent::AgentSpace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentSpaceId` | agent_space_id | `string` | computed |  | The unique identifier of the AgentSpace |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the AgentSpace. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the resource was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the AgentSpace. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the KMS key to use for encryption. |
| `Locale` |  | `string` | optional, computed, provider-chosen |  | The locale for the AgentSpace, which determines the language used in agent responses. |
| `Name` |  | `string` | required |  | The name of the AgentSpace. |
| `OperatorApp` | operator_app | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the resource was last updated. |

Supports update: yes

Discovery: supported
