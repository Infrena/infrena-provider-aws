# aws.limit

**CloudFormation type:** `AWS::Deadline::Limit`

Resource Type definition for AWS::Deadline::Limit

Region attribute: `region`

**Import ID:** `<region>/FarmId|LimitId` (AWS::Deadline::Limit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmountRequirementName` | amount_requirement_name | `string` | required, replaces on change |  |  |
| `CurrentCount` | current_count | `integer` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `LimitId` | limit_id | `string` | computed |  |  |
| `MaxCount` | max_count | `integer` | required |  |  |

Supports update: yes

Discovery: supported (parent resource required)
