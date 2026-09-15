# aws.findingsfilter

**CloudFormation type:** `AWS::Macie::FindingsFilter`

Macie FindingsFilter resource schema.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Macie::FindingsFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | optional, computed, provider-chosen |  | Findings filter action. |
| `Arn` |  | `string` | computed |  | Findings filter ARN. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Findings filter description |
| `FindingCriteria` | finding_criteria | `map` | required |  | Findings filter criteria. |
| `Id` |  | `string` | computed |  | Findings filter ID. |
| `Name` |  | `string` | required |  | Findings filter name |
| `Position` |  | `integer` | optional, computed, provider-chosen |  | Findings filter position. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
