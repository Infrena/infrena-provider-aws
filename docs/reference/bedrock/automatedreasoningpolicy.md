# aws.automatedreasoningpolicy

**CloudFormation type:** `AWS::Bedrock::AutomatedReasoningPolicy`

Definition of AWS::Bedrock::AutomatedReasoningPolicy Resource Type

Region attribute: `region`

**Import ID:** `<region>/PolicyArn` (AWS::Bedrock::AutomatedReasoningPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Time this policy was created |
| `DefinitionHash` | definition_hash | `string` | computed |  | The hash for this version |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ForceDelete` | force_delete | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to force delete the automated reasoning policy even if it has active resources. When false , Amazon Bedrock validates if all artifacts have been deleted (e.g. policy version, test case, test result) for a policy before deletion. When true , Amazon Bedrock will delete the policy and all its artifacts without validation. Default is false |
| `KmsKeyArn` | kms_key_arn | `string` | computed |  | The KMS key with which the Policy's assets will be encrypted at rest. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The KMS key with which the Policy's assets will be encrypted at rest. |
| `Name` |  | `string` | required |  |  |
| `PolicyArn` | policy_arn | `string` | computed |  |  |
| `PolicyDefinition` | policy_definition | `map` | optional, computed, provider-chosen, write-only |  |  |
| `PolicyId` | policy_id | `string` | computed |  | The id of the policy |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | Time this policy was last updated |
| `Version` |  | `string` | computed |  | Version of the policy that was created. This will always be `DRAFT` |

Supports update: yes

Discovery: supported
