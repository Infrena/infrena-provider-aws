# aws.metric

**CloudFormation type:** `AWS::Connect::Metric`

Resource Type definition for AWS::Connect::Metric, a custom metric configured for an Amazon Connect instance

Region attribute: `region`

**Import ID:** `<region>/MetricArn` (AWS::Connect::Metric)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Category` |  | `string` | computed |  | The category of the custom metric |
| `CreatedTime` | created_time | `float` | computed |  | The timestamp where the metric was created |
| `CreatedUser` | created_user | `map` | computed |  |  |
| `CreationMethod` | creation_method | `string` | computed |  | Whether the metric was built with the guided Service Level (SL) experience, or with the free-form metric builder |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the custom metric |
| `EffectiveTime` | effective_time | `float` | computed |  | Earliest time that can be queried for this metric |
| `Filters` |  | `list` | computed |  | List of filter types that may be used with this metric |
| `Groupings` |  | `list` | computed |  | List of groupings that may be used with this metric |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | The AWS region where the metric was last modified |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | The timestamp where the metric was last modified |
| `LastModifiedUser` | last_modified_user | `map` | computed |  |  |
| `MetricArn` | metric_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the custom metric. |
| `MetricCalculation` | metric_calculation | `map` | required |  | The calculation configuration for the metric |
| `Name` |  | `string` | required |  | The name of the custom metric |
| `PositiveTrendIndicator` | positive_trend_indicator | `string` | optional, computed, provider-chosen |  | Indicates how to classify a positive trend in metric data on the UI |
| `PrimaryEventSource` | primary_event_source | `string` | computed |  | Main provider of the document/row-level data for the metric; should match Data Lake table names |
| `PrimaryEventSourceEffectiveTimestampType` | primary_event_source_effective_timestamp_type | `string` | computed |  | Identifies the timestamp used to place the metrics on a time-series; should match public attribute name |
| `RefreshRate` | refresh_rate | `integer` | computed |  | Recommended time to wait between each refresh of data for the metric |
| `Status` |  | `string` | required, replaces on change |  | The status of the custom metric |
| `SupportedStats` | supported_stats | `list` | computed |  | List of stat aggregations available for the metric |
| `SupportsCustomCalculation` | supports_custom_calculation | `boolean` | computed |  | The metric may be used to compose other (custom) metrics |
| `SupportsPreaggregateCalculation` | supports_preaggregate_calculation | `boolean` | computed |  | The metric may be used to compose other (custom) metrics, meaning it can be used inside of aggregate stat functions |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `Type` | type_value | `string` | computed |  | Whether the metric is provided out-of-the-box or created by each customer |
| `Unit` |  | `string` | required |  | Display unit for the metric data |

Supports update: yes

Discovery: supported
