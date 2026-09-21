# aws.metricattribution

**CloudFormation type:** `AWS::Personalize::MetricAttribution`

Creates a metric attribution for reporting on recommendation impact in Amazon Personalize.

Region attribute: `region`

**Import ID:** `<region>/MetricAttributionArn` (AWS::Personalize::MetricAttribution)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetGroupArn` | dataset_group_arn | `string` | required, replaces on change | aws.personalize.datasetgroup.DatasetGroupArn | The ARN of the destination dataset group. |
| `MetricAttributionArn` | metric_attribution_arn | `string` | computed |  | The ARN of the metric attribution. |
| `Metrics` |  | `list` | required, replaces on change |  | A list of metric attributes for the metric attribution. |
| `MetricsOutputConfig` | metrics_output_config | `map` | required |  | The output configuration details for the metric attribution. |
| `Name` |  | `string` | required, replaces on change |  | The name of the metric attribution. |
| `Status` |  | `string` | computed |  | The status of the metric attribution. |

Supports update: yes

Discovery: supported
