# aws.scalabletarget

**CloudFormation type:** `AWS::ApplicationAutoScaling::ScalableTarget`

The ``AWS::ApplicationAutoScaling::ScalableTarget`` resource specifies a resource that Application Auto Scaling can scale, such as an AWS::DynamoDB::Table or AWS::ECS::Service resource.

Region attribute: `region`

**Import ID:** `<region>/ResourceId|ScalableDimension|ServiceNamespace` (AWS::ApplicationAutoScaling::ScalableTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `MaxCapacity` | max_capacity | `integer` | required |  | The maximum value that you plan to scale out to. When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand. |
| `MinCapacity` | min_capacity | `integer` | required |  | The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand. |
| `ResourceId` | resource_id | `string` | required, replaces on change |  | The identifier of the resource associated with the scalable target. This string consists of the resource type and unique identifier. |
| `RoleARN` | role_arn | `string` | optional, computed, provider-chosen, write-only |  | Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide*. |
| `ScalableDimension` | scalable_dimension | `string` | required, replaces on change |  | The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property. |
| `ScheduledActions` | scheduled_actions | `list` | optional, computed, provider-chosen |  | The scheduled actions for the scalable target. Duplicates aren't allowed. |
| `ServiceNamespace` | service_namespace | `string` | required, replaces on change |  | The namespace of the AWS service that provides the resource, or a ``custom-resource``. |
| `SuspendedState` | suspended_state | `map` | optional, computed, provider-chosen |  | ``SuspendedState`` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies whether the scaling activities for a scalable target are in a suspended state. |

Supports update: yes

Discovery: supported
