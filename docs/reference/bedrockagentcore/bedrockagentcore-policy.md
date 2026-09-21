# aws.bedrockagentcore.policy

**CloudFormation type:** `AWS::BedrockAgentCore::Policy`

Resource Type definition for AWS::BedrockAgentCore::Policy

Region attribute: `region`

**Import ID:** `<region>/PolicyArn` (AWS::BedrockAgentCore::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the policy was created. |
| `Definition` |  | `map` | required |  | The definition structure for policies. Encapsulates different policy formats. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A human-readable description of the policy's purpose and functionality. |
| `EnforcementMode` | enforcement_mode | `string` | optional, computed, provider-chosen |  | Whether the policy contributes to the enforce decision returned to Gateway. LOG_ONLY policies are still evaluated but their decisions are observed only, allowing customers to validate a policy against real traffic before promoting it. |
| `Name` |  | `string` | required, replaces on change |  | The customer-assigned immutable name for the policy. Must be unique within the policy engine. |
| `PolicyArn` | policy_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the policy. |
| `PolicyEngineId` | policy_engine_id | `string` | required, replaces on change | aws.policyengine.PolicyEngineId | The identifier of the policy engine which contains this policy. |
| `PolicyId` | policy_id | `string` | computed |  | The unique identifier for the policy. |
| `Status` |  | `string` | computed |  | The current status of the policy. |
| `StatusReasons` | status_reasons | `list` | computed |  | Additional information about the policy status. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the policy was last updated. |
| `ValidationMode` | validation_mode | `string` | optional, computed, provider-chosen, write-only |  | The validation mode for the policy. Determines how Cedar analyzer validation results are handled. |

Supports update: yes

Discovery: supported (parent resource required)
