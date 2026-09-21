# aws.glue.trigger

**CloudFormation type:** `AWS::Glue::Trigger`

Resource Type definition for AWS::Glue::Trigger

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::Trigger)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | The actions initiated by this trigger. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of this trigger. |
| `EventBatchingCondition` | event_batching_condition | `map` | optional, computed, provider-chosen |  | Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the trigger. |
| `Predicate` |  | `map` | optional, computed, provider-chosen |  | The predicate of this trigger, which defines when it will fire. |
| `Schedule` |  | `string` | optional, computed, provider-chosen |  | A cron expression used to specify the schedule. |
| `StartOnCreation` | start_on_creation | `boolean` | optional, computed, provider-chosen, write-only |  | Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags to use with this trigger. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of trigger that this is. |
| `WorkflowName` | workflow_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the workflow associated with the trigger. |

Supports update: yes

Discovery: supported
