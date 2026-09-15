# aws.assistantassociation

**CloudFormation type:** `AWS::Wisdom::AssistantAssociation`

Definition of AWS::Wisdom::AssistantAssociation Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssistantAssociationId|AssistantId` (AWS::Wisdom::AssistantAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssistantArn` | assistant_arn | `string` | computed |  |  |
| `AssistantAssociationArn` | assistant_association_arn | `string` | computed |  |  |
| `AssistantAssociationId` | assistant_association_id | `string` | computed |  |  |
| `AssistantId` | assistant_id | `string` | required, replaces on change | aws.assistant.AssistantId |  |
| `Association` |  | `string` | required, replaces on change |  |  |
| `AssociationType` | association_type | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
