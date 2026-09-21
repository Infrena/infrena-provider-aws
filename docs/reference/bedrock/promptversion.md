# aws.promptversion

**CloudFormation type:** `AWS::Bedrock::PromptVersion`

Definition of AWS::Bedrock::PromptVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Bedrock::PromptVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of a prompt version resource |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `CustomerEncryptionKeyArn` | customer_encryption_key_arn | `string` | computed |  | A KMS key ARN |
| `DefaultVariant` | default_variant | `string` | computed |  | Name for a variant. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description for a prompt version resource. |
| `Name` |  | `string` | computed |  | Name for a prompt resource. |
| `PromptArn` | prompt_arn | `string` | required, replaces on change | aws.bedrock.prompt.Arn | ARN of a prompt resource possibly with a version |
| `PromptId` | prompt_id | `string` | computed |  | Identifier for a Prompt |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |
| `Variants` |  | `list` | computed |  | List of prompt variants |
| `Version` |  | `string` | computed |  | Version. |

Supports update: no

Discovery: supported (parent resource required)
