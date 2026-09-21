# aws.subscriptionfilter

**CloudFormation type:** `AWS::Logs::SubscriptionFilter`

The ``AWS::Logs::SubscriptionFilter`` resource specifies a subscription filter and associates it with the specified log group. Subscription filters allow you to subscribe to a real-time stream of log events and have them delivered to a specific destination. Currently, the supported destinations are:

Region attribute: `region`

**Import ID:** `<region>/FilterName|LogGroupName` (AWS::Logs::SubscriptionFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplyOnTransformedLogs` | apply_on_transformed_logs | `boolean` | optional, computed, provider-chosen |  | This parameter is valid only for log groups that have an active log transformer. For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html). |
| `DestinationArn` | destination_arn | `string` | required | aws.logs.destination.Arn | The Amazon Resource Name (ARN) of the destination. |
| `Distribution` |  | `string` | optional, computed, provider-chosen |  | The method used to distribute log data to the destination, which can be either random or grouped by log stream. |
| `EmitSystemFields` | emit_system_fields | `list` | optional, computed, provider-chosen |  | The list of system fields that are included in the log events sent to the subscription destination. Returns the ``emitSystemFields`` value if it was specified when the subscription filter was created. |
| `FieldSelectionCriteria` | field_selection_criteria | `string` | optional, computed, provider-chosen |  | The filter expression that specifies which log events are processed by this subscription filter based on system fields. Returns the ``fieldSelectionCriteria`` value if it was specified when the subscription filter was created. |
| `FilterName` | filter_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the subscription filter. |
| `FilterPattern` | filter_pattern | `string` | required |  | The filtering expressions that restrict what gets delivered to the destination AWS resource. For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html). |
| `LogGroupName` | log_group_name | `string` | required, replaces on change |  | The log group to associate with the subscription filter. All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery. |

Supports update: yes

Discovery: supported (parent resource required)
