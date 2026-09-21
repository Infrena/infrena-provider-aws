# aws.deliverydestination

**CloudFormation type:** `AWS::Logs::DeliveryDestination`

This structure contains information about one delivery destination in your account.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Logs::DeliveryDestination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that uniquely identifies a resource. |
| `DeliveryDestinationPolicy` | delivery_destination_policy | `map` | optional, computed, provider-chosen, write-only |  | IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account. |
| `DeliveryDestinationType` | delivery_destination_type | `string` | optional, computed, provider-chosen, replaces on change |  | Displays whether this delivery destination is CloudWatch Logs, Amazon S3, Kinesis Data Firehose, or XRay. |
| `DestinationResourceArn` | destination_resource_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) that uniquely identifies a resource. |
| `Name` |  | `string` | required, replaces on change |  | The name of this delivery destination. |
| `OutputFormat` | output_format | `string` | optional, computed, provider-chosen, replaces on change |  | The format of the logs that are sent to this delivery destination. |
| `Tags` |  | `map` | tags map |  | The tags that have been assigned to this delivery destination. |

Supports update: yes

Discovery: supported
