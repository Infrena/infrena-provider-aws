# aws.cloudwatch.alarm

**CloudFormation type:** `AWS::CloudWatch::Alarm`

The ``AWS::CloudWatch::Alarm`` type specifies an alarm and associates it with the specified metric or metric math expression.

Region attribute: `region`

**Import ID:** `<region>/AlarmName` (AWS::CloudWatch::Alarm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionsEnabled` | actions_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE. |
| `AlarmActions` | alarm_actions | `list` | optional, computed, provider-chosen |  | The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*. |
| `AlarmDescription` | alarm_description | `string` | optional, computed, provider-chosen |  | The description of the alarm. |
| `AlarmName` | alarm_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name. |
| `Arn` |  | `string` | computed |  |  |
| `ComparisonOperator` | comparison_operator | `string` | optional, computed, provider-chosen |  | The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand. |
| `DatapointsToAlarm` | datapoints_to_alarm | `integer` | optional, computed, provider-chosen |  | The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*. |
| `Dimensions` |  | `list` | optional, computed, provider-chosen |  | The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``. |
| `EvaluateLowSampleCountPercentile` | evaluate_low_sample_count_percentile | `string` | optional, computed, provider-chosen |  | Used only for alarms based on percentiles. If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available. |
| `EvaluationCriteria` | evaluation_criteria | `map` | optional, computed, provider-chosen |  | The evaluation criteria for an alarm. This is a union type that currently supports ``PromQLCriteria``. |
| `EvaluationInterval` | evaluation_interval | `integer` | optional, computed, provider-chosen |  | The frequency, in seconds, at which the alarm is evaluated. |
| `EvaluationPeriods` | evaluation_periods | `integer` | optional, computed, provider-chosen |  | The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M. |
| `EvaluationWindow` | evaluation_window | `map` | optional, computed, provider-chosen |  | The evaluation window that an alarm uses to select the range of metric data that it evaluates each time it runs. This is a union type. Set exactly one of its members, ``SlidingWindow`` or ``WallClockWindow``. If you don't set ``EvaluationWindow``, the alarm uses a ``SlidingWindow`` by default. |
| `ExtendedStatistic` | extended_statistic | `string` | optional, computed, provider-chosen |  | The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100. |
| `InsufficientDataActions` | insufficient_data_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state. Each action is specified as an Amazon Resource Name (ARN). |
| `MetricName` | metric_name | `string` | optional, computed, provider-chosen |  | The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``. |
| `Metrics` |  | `list` | optional, computed, provider-chosen |  | An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression. |
| `Namespace` |  | `string` | optional, computed, provider-chosen |  | The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead. |
| `OKActions` | ok_actions | `list` | optional, computed, provider-chosen |  | The actions to execute when this alarm transitions to the ``OK`` state from any other state. Each action is specified as an Amazon Resource Name (ARN). |
| `Period` |  | `integer` | optional, computed, provider-chosen |  | The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 20, 30, 60, and any multiple of 60. |
| `Statistic` |  | `string` | optional, computed, provider-chosen |  | The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ``ExtendedStatistic``. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission. |
| `Threshold` |  | `float` | optional, computed, provider-chosen |  | The value to compare with the specified statistic. |
| `ThresholdMetricId` | threshold_metric_id | `string` | optional, computed, provider-chosen |  | In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm. |
| `TreatMissingData` | treat_missing_data | `string` | optional, computed, provider-chosen |  | Sets how this alarm is to handle missing data points. Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*. |
| `Unit` |  | `string` | optional, computed, provider-chosen |  | The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array. |
| `WarmUpConfiguration` | warm_up_configuration | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
