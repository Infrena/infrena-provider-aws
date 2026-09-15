# aws.eventtrigger

**CloudFormation type:** `AWS::CustomerProfiles::EventTrigger`

An event trigger resource of Amazon Connect Customer Profiles

Region attribute: `region`

**Import ID:** `<region>/DomainName|EventTriggerName` (AWS::CustomerProfiles::EventTrigger)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the event trigger was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the event trigger. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `EventTriggerConditions` | event_trigger_conditions | `list` | required |  | A list of conditions that determine when an event should trigger the destination. |
| `EventTriggerLimits` | event_trigger_limits | `map` | optional, computed, provider-chosen |  | Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods. |
| `EventTriggerName` | event_trigger_name | `string` | required, replaces on change |  | The unique name of the event trigger. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the event trigger was most recently updated. |
| `ObjectTypeName` | object_type_name | `string` | required |  | The unique name of the object type. |
| `SegmentFilter` | segment_filter | `string` | optional, computed, provider-chosen |  | The destination is triggered only for profiles that meet the criteria of a segment definition. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
