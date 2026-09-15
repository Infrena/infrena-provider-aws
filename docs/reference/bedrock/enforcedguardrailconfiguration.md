# aws.enforcedguardrailconfiguration

**CloudFormation type:** `AWS::Bedrock::EnforcedGuardrailConfiguration`

Definition of AWS::Bedrock::EnforcedGuardrailConfiguration Resource Type

Region attribute: `region`

**Import ID:** `<region>/ConfigId` (AWS::Bedrock::EnforcedGuardrailConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigId` | config_id | `string` | computed |  | Unique ID for the account enforced configuration |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the configuration was created |
| `CreatedBy` | created_by | `string` | computed |  | The ARN of the role used to create the configuration |
| `GuardrailArn` | guardrail_arn | `string` | computed |  | ARN representation for the guardrail |
| `GuardrailId` | guardrail_id | `string` | computed |  | Unique ID for the guardrail |
| `GuardrailIdentifier` | guardrail_identifier | `string` | required |  | Identifier for the guardrail, could be the ID or the ARN |
| `GuardrailVersion` | guardrail_version | `string` | required |  | Numerical guardrail version (not DRAFT) |
| `ModelEnforcement` | model_enforcement | `map` | optional, computed, provider-chosen |  | Model-specific information for the enforced guardrail configuration. If not present, the configuration is enforced on all models |
| `Owner` |  | `string` | computed |  | Configuration owner type |
| `SelectiveContentGuarding` | selective_content_guarding | `map` | optional, computed, provider-chosen |  | Selective content guarding controls for enforced guardrails |
| `UpdatedAt` | updated_at | `string` | computed |  | Timestamp when the configuration was last updated |
| `UpdatedBy` | updated_by | `string` | computed |  | The ARN of the role used to update the configuration |

Supports update: yes

Discovery: supported
