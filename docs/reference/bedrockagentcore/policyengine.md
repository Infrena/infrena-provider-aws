# aws.policyengine

**CloudFormation type:** `AWS::BedrockAgentCore::PolicyEngine`

Resource Type definition for AWS::BedrockAgentCore::PolicyEngine

Region attribute: `region`

**Import ID:** `<region>/PolicyEngineArn` (AWS::BedrockAgentCore::PolicyEngine)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the policy engine was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A human-readable description of the policy engine's purpose and scope |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the KMS key used to encrypt the policy engine data |
| `Name` |  | `string` | required, replaces on change |  | The customer-assigned immutable name for the policy engine |
| `PolicyEngineArn` | policy_engine_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the policy engine |
| `PolicyEngineId` | policy_engine_id | `string` | computed |  | The unique identifier for the policy engine |
| `Status` |  | `string` | computed |  | The current status of the policy engine |
| `StatusReasons` | status_reasons | `list` | computed |  | Additional information about the policy engine status |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to assign to the policy engine. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the policy engine was last updated |

Supports update: yes

Discovery: supported
