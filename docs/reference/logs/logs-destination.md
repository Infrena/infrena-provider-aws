# aws.logs.destination

**CloudFormation type:** `AWS::Logs::Destination`

The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.

Region attribute: `region`

**Import ID:** `<region>/DestinationName` (AWS::Logs::Destination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DestinationName` | destination_name | `string` | required, replaces on change |  | The name of the destination resource |
| `DestinationPolicy` | destination_policy | `string` | optional, computed, provider-chosen |  | An IAM policy document that governs which AWS accounts can create subscription filters against this destination. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetArn` | target_arn | `string` | required |  | The ARN of the physical target where the log events are delivered (for example, a Kinesis stream) |

Supports update: yes

Discovery: supported
