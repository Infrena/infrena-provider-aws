# aws.bedrock.prompt

**CloudFormation type:** `AWS::Bedrock::Prompt`

Definition of AWS::Bedrock::Prompt Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Bedrock::Prompt)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of a prompt resource possibly with a version |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `CustomerEncryptionKeyArn` | customer_encryption_key_arn | `string` | optional, computed, provider-chosen |  | A KMS key ARN |
| `DefaultVariant` | default_variant | `string` | optional, computed, provider-chosen |  | Name for a variant. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Name for a prompt resource. |
| `Id` |  | `string` | computed |  | Identifier for a Prompt |
| `Name` |  | `string` | required |  | Name for a prompt resource. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |
| `Variants` |  | `list` | optional, computed, provider-chosen |  | List of prompt variants |
| `Version` |  | `string` | computed |  | Draft Version. |

Supports update: yes

Discovery: supported
