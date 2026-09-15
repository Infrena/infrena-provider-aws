# aws.autoscaling.scheduledaction

**CloudFormation type:** `AWS::AutoScaling::ScheduledAction`

The AWS::AutoScaling::ScheduledAction resource specifies an Amazon EC2 Auto Scaling scheduled action so that the Auto Scaling group can change the number of instances available for your application in response to predictable load changes.

Region attribute: `region`

**Import ID:** `<region>/ScheduledActionName|AutoScalingGroupName` (AWS::AutoScaling::ScheduledAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingGroupName` | auto_scaling_group_name | `string` | required, replaces on change |  | The name of the Auto Scaling group. |
| `DesiredCapacity` | desired_capacity | `integer` | optional, computed, provider-chosen |  | The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. |
| `EndTime` | end_time | `string` | optional, computed, provider-chosen |  | The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored. |
| `MaxSize` | max_size | `integer` | optional, computed, provider-chosen |  | The minimum size of the Auto Scaling group. |
| `MinSize` | min_size | `integer` | optional, computed, provider-chosen |  | The minimum size of the Auto Scaling group. |
| `Recurrence` |  | `string` | optional, computed, provider-chosen |  | The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops. |
| `ScheduledActionName` | scheduled_action_name | `string` | computed |  | Auto-generated unique identifier |
| `StartTime` | start_time | `string` | optional, computed, provider-chosen |  | The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored. |
| `TimeZone` | time_zone | `string` | optional, computed, provider-chosen |  | The time zone for the cron expression. |

Supports update: yes

Discovery: supported
