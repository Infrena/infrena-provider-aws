# aws.cloudwatchalarmtemplate

**CloudFormation type:** `AWS::MediaLive::CloudWatchAlarmTemplate`

Definition of AWS::MediaLive::CloudWatchAlarmTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/Identifier` (AWS::MediaLive::CloudWatchAlarmTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | A cloudwatch alarm template's ARN (Amazon Resource Name) |
| `ComparisonOperator` | comparison_operator | `string` | required |  | The comparison operator used to compare the specified statistic and the threshold. |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DatapointsToAlarm` | datapoints_to_alarm | `float` | optional, computed, provider-chosen |  | The number of datapoints within the evaluation period that must be breaching to trigger the alarm. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A resource's optional description. |
| `EvaluationPeriods` | evaluation_periods | `float` | required |  | The number of periods over which data is compared to the specified threshold. |
| `GroupId` | group_id | `string` | computed |  | A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-` |
| `GroupIdentifier` | group_identifier | `string` | optional, computed, provider-chosen, write-only |  | A cloudwatch alarm template group's identifier. Can be either be its id or current name. |
| `Id` |  | `string` | computed |  | A cloudwatch alarm template's id. AWS provided templates have ids that start with `aws-` |
| `Identifier` |  | `string` | computed |  |  |
| `MetricName` | metric_name | `string` | required |  | The name of the metric associated with the alarm. Must be compatible with targetResourceType. |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  | A resource's name. Names must be unique within the scope of a resource type in a specific region. |
| `Period` |  | `float` | required |  | The period, in seconds, over which the specified statistic is applied. |
| `Statistic` |  | `string` | required |  | The statistic to apply to the alarm's metric data. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Represents the tags associated with a resource. |
| `TargetResourceType` | target_resource_type | `string` | required |  | The resource type this template should dynamically generate cloudwatch metric alarms for. |
| `Threshold` |  | `float` | required |  | The threshold value to compare with the specified statistic. |
| `TreatMissingData` | treat_missing_data | `string` | required |  | Specifies how missing data points are treated when evaluating the alarm's condition. |

Supports update: yes

Discovery: supported
