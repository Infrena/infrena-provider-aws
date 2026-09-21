# aws.flowlog

**CloudFormation type:** `AWS::EC2::FlowLog`

Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::FlowLog)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeliverCrossAccountRole` | deliver_cross_account_role | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts. |
| `DeliverLogsPermissionArn` | deliver_logs_permission_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName. |
| `DestinationOptions` | destination_options | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  | The Flow Log ID |
| `LogDestination` | log_destination | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType. |
| `LogDestinationType` | log_destination_type | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3. |
| `LogFormat` | log_format | `string` | optional, computed, provider-chosen, replaces on change |  | The fields to include in the flow log record, in the order in which they should appear. |
| `LogGroupName` | log_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName. |
| `MaxAggregationInterval` | max_aggregation_interval | `integer` | optional, computed, provider-chosen, replaces on change |  | The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes). |
| `ResourceId` | resource_id | `string` | required, replaces on change |  | The ID of the subnet, network interface, or VPC for which you want to create a flow log. |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property. |
| `TagFieldSpecifications` | tag_field_specifications | `list` | optional, computed, provider-chosen, replaces on change |  | The resource types and associated tags for EC2 resources associated with the EC2 Tags feature for log enrichment. |
| `Tags` |  | `map` | tags map |  | The tags to apply to the flow logs. |
| `TrafficType` | traffic_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic. |

Supports update: yes

Discovery: supported
