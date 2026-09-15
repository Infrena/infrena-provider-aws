# aws.metricstream

**CloudFormation type:** `AWS::CloudWatch::MetricStream`

Resource Type definition for Metric Stream

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::CloudWatch::MetricStream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name of the metric stream. |
| `CreationDate` | creation_date | `string` | computed |  | The date of creation of the metric stream. |
| `ExcludeFilters` | exclude_filters | `list` | optional, computed, provider-chosen |  | Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null. |
| `FirehoseArn` | firehose_arn | `string` | optional, computed, provider-chosen |  | The ARN of the Kinesis Firehose where to stream the data. |
| `IncludeFilters` | include_filters | `list` | optional, computed, provider-chosen |  | Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null. |
| `IncludeLinkedAccountsMetrics` | include_linked_accounts_metrics | `boolean` | optional, computed, provider-chosen |  | If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. |
| `LastUpdateDate` | last_update_date | `string` | computed |  | The date of the last update of the metric stream. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the metric stream. |
| `OutputFormat` | output_format | `string` | optional, computed, provider-chosen |  | The output format of the data streamed to the Kinesis Firehose. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the role that provides access to the Kinesis Firehose. |
| `State` |  | `string` | computed |  | Displays the state of the Metric Stream. |
| `StatisticsConfigurations` | statistics_configurations | `list` | optional, computed, provider-chosen |  | By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A set of tags to assign to the delivery stream. |

Supports update: yes

Discovery: supported
