# aws.schedulegroup

**CloudFormation type:** `AWS::Scheduler::ScheduleGroup`

Definition of AWS::Scheduler::ScheduleGroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Scheduler::ScheduleGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the schedule group. |
| `CreationDate` | creation_date | `string` | computed |  | The time at which the schedule group was created. |
| `LastModificationDate` | last_modification_date | `string` | computed |  | The time at which the schedule group was last modified. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `State` |  | `string` | computed |  | Specifies the state of the schedule group. |
| `Tags` |  | `map` | tags map |  | The list of tags to associate with the schedule group. |

Supports update: yes

Discovery: supported
