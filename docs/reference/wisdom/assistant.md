# aws.assistant

**CloudFormation type:** `AWS::Wisdom::Assistant`

Definition of AWS::Wisdom::Assistant Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssistantId` (AWS::Wisdom::Assistant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `ServerSideEncryptionConfiguration` | server_side_encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
