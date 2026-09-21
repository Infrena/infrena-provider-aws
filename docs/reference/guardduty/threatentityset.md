# aws.threatentityset

**CloudFormation type:** `AWS::GuardDuty::ThreatEntitySet`

Resource Type definition for AWS::GuardDuty::ThreatEntitySet

Region attribute: `region`

**Import ID:** `<region>/Id|DetectorId` (AWS::GuardDuty::ThreatEntitySet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Activate` |  | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DetectorId` | detector_id | `string` | optional, computed, provider-chosen, replaces on change | aws.guardduty.detector.Id |  |
| `ErrorDetails` | error_details | `string` | computed |  |  |
| `ExpectedBucketOwner` | expected_bucket_owner | `string` | optional, computed, provider-chosen |  |  |
| `Format` |  | `string` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Location` |  | `string` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
