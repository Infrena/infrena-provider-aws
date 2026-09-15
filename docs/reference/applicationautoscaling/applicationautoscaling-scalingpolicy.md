# aws.applicationautoscaling.scalingpolicy

**CloudFormation type:** `AWS::ApplicationAutoScaling::ScalingPolicy`

The ``AWS::ApplicationAutoScaling::ScalingPolicy`` resource defines a scaling policy that Application Auto Scaling uses to adjust the capacity of a scalable target. 

Region attribute: `region`

**Import ID:** `<region>/Arn|ScalableDimension` (AWS::ApplicationAutoScaling::ScalingPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the scaling policy. |
| `PolicyType` | policy_type | `string` | required |  | The scaling policy type. |
| `PredictiveScalingPolicyConfiguration` | predictive_scaling_policy_configuration | `map` | optional, computed, provider-chosen |  | Represents a predictive scaling policy configuration. Predictive scaling is supported on Amazon ECS services. |
| `ResourceId` | resource_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier. |
| `ScalableDimension` | scalable_dimension | `string` | optional, computed, provider-chosen, replaces on change |  | The scalable dimension. This string consists of the service namespace, resource type, and scaling property. |
| `ScalingTargetId` | scaling_target_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the ``AWS::ApplicationAutoScaling::ScalableTarget`` resource. |
| `ServiceNamespace` | service_namespace | `string` | optional, computed, provider-chosen, replaces on change |  | The namespace of the AWS service that provides the resource, or a ``custom-resource``. |
| `StepScalingPolicyConfiguration` | step_scaling_policy_configuration | `map` | optional, computed, provider-chosen |  | ``StepScalingPolicyConfiguration`` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a step scaling policy configuration for Application Auto Scaling. |
| `TargetTrackingScalingPolicyConfiguration` | target_tracking_scaling_policy_configuration | `map` | optional, computed, provider-chosen |  | ``TargetTrackingScalingPolicyConfiguration`` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a target tracking scaling policy configuration for Application Auto Scaling. Use a target tracking scaling policy to adjust the capacity of the specified scalable target in response to actual workloads, so that resource utilization remains at or near the target utilization value. |

Supports update: yes

Discovery: supported (parent resource required)
