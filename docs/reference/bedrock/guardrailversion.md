# aws.guardrailversion

**CloudFormation type:** `AWS::Bedrock::GuardrailVersion`

Definition of AWS::Bedrock::GuardrailVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/GuardrailId|Version` (AWS::Bedrock::GuardrailVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the Guardrail version |
| `GuardrailArn` | guardrail_arn | `string` | computed |  | Arn representation for the guardrail |
| `GuardrailId` | guardrail_id | `string` | computed |  | Unique id for the guardrail |
| `GuardrailIdentifier` | guardrail_identifier | `string` | required, replaces on change, write-only |  | Identifier (GuardrailId or GuardrailArn) for the guardrail |
| `Version` |  | `string` | computed |  | Guardrail version |

Supports update: no

Discovery: not supported
