# aws.b2bi.capability

**CloudFormation type:** `AWS::B2BI::Capability`

Definition of AWS::B2BI::Capability Resource Type

Region attribute: `region`

**Import ID:** `<region>/CapabilityId` (AWS::B2BI::Capability)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapabilityArn` | capability_arn | `string` | computed |  |  |
| `CapabilityId` | capability_id | `string` | computed |  |  |
| `Configuration` |  | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `InstructionsDocuments` | instructions_documents | `list` | optional, computed, provider-chosen |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
