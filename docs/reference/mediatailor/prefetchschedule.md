# aws.prefetchschedule

**CloudFormation type:** `AWS::MediaTailor::PrefetchSchedule`

Definition of AWS::MediaTailor::PrefetchSchedule Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaTailor::PrefetchSchedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the prefetch schedule. |
| `Consumption` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  | The name to assign to the prefetch schedule. |
| `PlaybackConfigurationName` | playback_configuration_name | `string` | required, replaces on change |  | The name of the playback configuration. |
| `RecurringPrefetchConfiguration` | recurring_prefetch_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Retrieval` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ScheduleType` | schedule_type | `string` | optional, computed, provider-chosen, replaces on change |  | The frequency that MediaTailor creates prefetch schedules. |
| `StreamId` | stream_id | `string` | optional, computed, provider-chosen, replaces on change |  | An optional stream identifier that MediaTailor uses to prefetch ads for multiple streams that use the same playback configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the prefetch schedule. |

Supports update: yes

Discovery: supported (parent resource required)
