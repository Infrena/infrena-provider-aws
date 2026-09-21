# aws.threatintelset

**CloudFormation type:** `AWS::GuardDuty::ThreatIntelSet`

Resource Type definition for AWS::GuardDuty::ThreatIntelSet

Region attribute: `region`

**Import ID:** `<region>/Id|DetectorId` (AWS::GuardDuty::ThreatIntelSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Activate` |  | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `DetectorId` | detector_id | `string` | optional, computed, provider-chosen, replaces on change | aws.guardduty.detector.Id |  |
| `ExpectedBucketOwner` | expected_bucket_owner | `string` | optional, computed, provider-chosen |  |  |
| `Format` |  | `string` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Location` |  | `string` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
