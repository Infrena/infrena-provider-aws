# aws.inspectorv2.filter

**CloudFormation type:** `AWS::InspectorV2::Filter`

Inspector Filter resource schema

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::InspectorV2::Filter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Findings filter ARN. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Findings filter description. |
| `FilterAction` | filter_action | `string` | required |  | Findings filter action. |
| `FilterCriteria` | filter_criteria | `map` | required |  | Findings filter criteria. |
| `Name` |  | `string` | required |  | Findings filter name. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
