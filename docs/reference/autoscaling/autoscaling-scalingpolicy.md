# aws.autoscaling.scalingpolicy

**CloudFormation type:** `AWS::AutoScaling::ScalingPolicy`

The AWS::AutoScaling::ScalingPolicy resource specifies an Amazon EC2 Auto Scaling scaling policy so that the Auto Scaling group can scale the number of instances available for your application.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AutoScaling::ScalingPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdjustmentType` | adjustment_type | `string` | optional, computed, provider-chosen |  | Specifies how the scaling adjustment is interpreted. The valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity. |
| `Arn` |  | `string` | computed |  | The ARN of the AutoScaling scaling policy |
| `AutoScalingGroupName` | auto_scaling_group_name | `string` | required, replaces on change |  | The name of the Auto Scaling group. |
| `Cooldown` |  | `string` | optional, computed, provider-chosen |  | The duration of the policy's cooldown period, in seconds. When a cooldown period is specified here, it overrides the default cooldown period defined for the Auto Scaling group. |
| `EstimatedInstanceWarmup` | estimated_instance_warmup | `integer` | optional, computed, provider-chosen |  | The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics. If not provided, the default is to use the value from the default cooldown period for the Auto Scaling group. Valid only if the policy type is TargetTrackingScaling or StepScaling. |
| `MetricAggregationType` | metric_aggregation_type | `string` | optional, computed, provider-chosen |  | The aggregation type for the CloudWatch metrics. The valid values are Minimum, Maximum, and Average. If the aggregation type is null, the value is treated as Average. Valid only if the policy type is StepScaling. |
| `MinAdjustmentMagnitude` | min_adjustment_magnitude | `integer` | optional, computed, provider-chosen |  | The minimum value to scale by when the adjustment type is PercentChangeInCapacity. For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances. |
| `PolicyName` | policy_name | `string` | computed |  |  |
| `PolicyType` | policy_type | `string` | optional, computed, provider-chosen |  | One of the following policy types: TargetTrackingScaling, StepScaling, SimpleScaling (default), PredictiveScaling |
| `PredictiveScalingConfiguration` | predictive_scaling_configuration | `map` | optional, computed, provider-chosen |  | A predictive scaling policy. Includes support for predefined metrics only. |
| `ScalingAdjustment` | scaling_adjustment | `integer` | optional, computed, provider-chosen |  | The amount by which to scale, based on the specified adjustment type. A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value. Required if the policy type is SimpleScaling. (Not used with any other policy type.) |
| `StepAdjustments` | step_adjustments | `list` | optional, computed, provider-chosen |  | A set of adjustments that enable you to scale based on the size of the alarm breach. Required if the policy type is StepScaling. (Not used with any other policy type.) |
| `TargetTrackingConfiguration` | target_tracking_configuration | `map` | optional, computed, provider-chosen |  | A target tracking scaling policy. Includes support for predefined or customized metrics. |

Supports update: yes

Discovery: supported
