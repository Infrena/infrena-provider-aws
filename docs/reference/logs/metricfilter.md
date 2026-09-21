# aws.metricfilter

**CloudFormation type:** `AWS::Logs::MetricFilter`

The ``AWS::Logs::MetricFilter`` resource specifies a metric filter that describes how CWL extracts information from logs and transforms it into Amazon CloudWatch metrics. If you have multiple metric filters that are associated with a log group, all the filters are applied to the log streams in that group.

Region attribute: `region`

**Import ID:** `<region>/LogGroupName|FilterName` (AWS::Logs::MetricFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplyOnTransformedLogs` | apply_on_transformed_logs | `boolean` | optional, computed, provider-chosen |  | This parameter is valid only for log groups that have an active log transformer. For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html). |
| `EmitSystemFieldDimensions` | emit_system_field_dimensions | `list` | optional, computed, provider-chosen |  | The list of system fields that are emitted as additional dimensions in the generated metrics. Returns the ``emitSystemFieldDimensions`` value if it was specified when the metric filter was created. |
| `FieldSelectionCriteria` | field_selection_criteria | `string` | optional, computed, provider-chosen |  | The filter expression that specifies which log events are processed by this metric filter based on system fields. Returns the ``fieldSelectionCriteria`` value if it was specified when the metric filter was created. |
| `FilterName` | filter_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the metric filter. |
| `FilterPattern` | filter_pattern | `string` | required |  | A filter pattern for extracting metric data out of ingested log events. For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html). |
| `LogGroupName` | log_group_name | `string` | required, replaces on change |  | The name of an existing log group that you want to associate with this metric filter. |
| `MetricTransformations` | metric_transformations | `list` | required |  | The metric transformations. |

Supports update: yes

Discovery: supported
