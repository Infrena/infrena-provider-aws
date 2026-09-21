# aws.executionplan

**CloudFormation type:** `AWS::KendraRanking::ExecutionPlan`

A KendraRanking Rescore execution plan

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::KendraRanking::ExecutionPlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CapacityUnits` | capacity_units | `map` | optional, computed, provider-chosen |  | Capacity units |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the execution plan |
| `Id` |  | `string` | computed |  | Unique ID of rescore execution plan |
| `Name` |  | `string` | required |  | Name of kendra ranking rescore execution plan |
| `Tags` |  | `map` | tags map |  | List of tags |

Supports update: yes

Discovery: supported
