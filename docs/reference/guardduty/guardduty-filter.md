# aws.guardduty.filter

**CloudFormation type:** `AWS::GuardDuty::Filter`

Resource Type definition for AWS::GuardDuty::Filter

Region attribute: `region`

**Import ID:** `<region>/DetectorId|Name` (AWS::GuardDuty::Filter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DetectorId` | detector_id | `string` | required, replaces on change | aws.guardduty.detector.Id |  |
| `FindingCriteria` | finding_criteria | `map` | required |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Rank` |  | `integer` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
