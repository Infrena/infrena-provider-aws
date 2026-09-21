# aws.maintenancewindow

**CloudFormation type:** `AWS::SSM::MaintenanceWindow`

Resource type definition for AWS::SSM::MaintenanceWindow

Region attribute: `region`

**Import ID:** `<region>/WindowId` (AWS::SSM::MaintenanceWindow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowUnassociatedTargets` | allow_unassociated_targets | `boolean` | required |  | Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets. If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window. |
| `Cutoff` |  | `integer` | required |  | The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the maintenance window. |
| `Duration` |  | `integer` | required |  | The duration of the maintenance window in hours. |
| `EndDate` | end_date | `string` | optional, computed, provider-chosen |  | The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive. |
| `Name` |  | `string` | required |  | The name of the maintenance window. |
| `Schedule` |  | `string` | required |  | The schedule of the maintenance window in the form of a cron or rate expression. |
| `ScheduleOffset` | schedule_offset | `integer` | optional, computed, provider-chosen |  | The number of days to wait to run a maintenance window after the scheduled cron expression date and time. |
| `ScheduleTimezone` | schedule_timezone | `string` | optional, computed, provider-chosen |  | The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format. |
| `StartDate` | start_date | `string` | optional, computed, provider-chosen |  | The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active. StartDate allows you to delay activation of the maintenance window until the specified future date. |
| `Tags` |  | `map` | tags map |  | Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in. |
| `WindowId` | window_id | `string` | computed |  | The ID of the maintenance window. |

Supports update: yes

Discovery: supported
