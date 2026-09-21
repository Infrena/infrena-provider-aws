# aws.deliverysource

**CloudFormation type:** `AWS::Logs::DeliverySource`

A delivery source is an AWS resource that sends logs to an AWS destination. The destination can be CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Logs::DeliverySource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that uniquely identifies this delivery source. |
| `DeliverySourceConfiguration` | delivery_source_configuration | `map` | optional, computed, provider-chosen |  | A map of key-value pairs to configure the delivery source. Both keys and values must be between 1 and 255 characters in length. |
| `LogType` | log_type | `string` | optional, computed, provider-chosen |  | The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options. |
| `Name` |  | `string` | required, replaces on change |  | The unique name of the Log source. |
| `ResourceArn` | resource_arn | `string` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Name (ARN) that uniquely identifies this delivery source. |
| `ResourceArns` | resource_arns | `list` | computed |  | This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array. |
| `Service` |  | `string` | computed |  | The AWS service that is sending logs. |
| `Status` |  | `string` | computed |  | The status of this delivery source. The value can be ACTIVE or INACTIVE. |
| `StatusReason` | status_reason | `string` | computed |  | The reason for the status of this delivery source, such as RESOURCE_DELETED. |
| `Tags` |  | `map` | tags map |  | The tags that have been assigned to this delivery source. |

Supports update: yes

Discovery: supported
