# aws.lifecyclehook

**CloudFormation type:** `AWS::AutoScaling::LifecycleHook`

Resource Type definition for AWS::AutoScaling::LifecycleHook

Region attribute: `region`

**Import ID:** `<region>/AutoScalingGroupName|LifecycleHookName` (AWS::AutoScaling::LifecycleHook)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingGroupName` | auto_scaling_group_name | `string` | required, replaces on change |  | The name of the Auto Scaling group for the lifecycle hook. |
| `DefaultResult` | default_result | `string` | optional, computed, provider-chosen |  | The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default). |
| `HeartbeatTimeout` | heartbeat_timeout | `integer` | optional, computed, provider-chosen |  | The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property. |
| `LifecycleHookName` | lifecycle_hook_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the lifecycle hook. |
| `LifecycleTransition` | lifecycle_transition | `string` | required |  | The instance state to which you want to attach the lifecycle hook. |
| `NotificationMetadata` | notification_metadata | `string` | optional, computed, provider-chosen |  | Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target. |
| `NotificationTargetARN` | notification_target_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata. |
| `RoleARN` | role_arn | `string` | optional, computed, provider-chosen |  | The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue. |

Supports update: yes

Discovery: supported
