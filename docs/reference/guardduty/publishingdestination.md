# aws.publishingdestination

**CloudFormation type:** `AWS::GuardDuty::PublishingDestination`

Resource Type definition for AWS::GuardDuty::PublishingDestination.

Region attribute: `region`

**Import ID:** `<region>/DetectorId|Id` (AWS::GuardDuty::PublishingDestination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DestinationProperties` | destination_properties | `map` | required |  |  |
| `DestinationType` | destination_type | `string` | required |  | The type of resource for the publishing destination. Currently only Amazon S3 buckets are supported. |
| `DetectorId` | detector_id | `string` | required, replaces on change | aws.guardduty.detector.Id | The ID of the GuardDuty detector associated with the publishing destination. |
| `Id` |  | `string` | computed |  | The ID of the publishing destination. |
| `PublishingFailureStartTimestamp` | publishing_failure_start_timestamp | `string` | computed |  | The time, in epoch millisecond format, at which GuardDuty was first unable to publish findings to the destination. |
| `Status` |  | `string` | computed |  | The status of the publishing destination. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
