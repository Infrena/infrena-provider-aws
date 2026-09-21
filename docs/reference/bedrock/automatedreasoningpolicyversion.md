# aws.automatedreasoningpolicyversion

**CloudFormation type:** `AWS::Bedrock::AutomatedReasoningPolicyVersion`

Definition of AWS::Bedrock::AutomatedReasoningPolicyVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/PolicyArn|Version` (AWS::Bedrock::AutomatedReasoningPolicyVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Time this policy version was created |
| `DefinitionHash` | definition_hash | `string` | computed |  | The hash for this version |
| `Description` |  | `string` | computed |  | The description inherited from the policy |
| `LastUpdatedDefinitionHash` | last_updated_definition_hash | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The hash for this version |
| `Name` |  | `string` | computed |  | The name inherited from the policy |
| `PolicyArn` | policy_arn | `string` | required, replaces on change |  | Arn of the policy |
| `PolicyId` | policy_id | `string` | computed |  | The id of the associated policy |
| `Tags` |  | `map` | replaces on change, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | Time this policy was last updated |
| `Version` |  | `string` | computed |  | The version of the policy |

Supports update: no

Discovery: supported
