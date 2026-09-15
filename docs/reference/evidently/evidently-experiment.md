# aws.evidently.experiment

**CloudFormation type:** `AWS::Evidently::Experiment`

Resource Type definition for AWS::Evidently::Experiment.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Evidently::Experiment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `MetricGoals` | metric_goals | `list` | required |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `OnlineAbConfig` | online_ab_config | `map` | required |  |  |
| `Project` |  | `string` | required, replaces on change |  |  |
| `RandomizationSalt` | randomization_salt | `string` | optional, computed, provider-chosen |  |  |
| `RemoveSegment` | remove_segment | `boolean` | optional, computed, provider-chosen |  |  |
| `RunningStatus` | running_status | `map` | optional, computed, provider-chosen |  | Start Experiment. Default is False |
| `SamplingRate` | sampling_rate | `integer` | optional, computed, provider-chosen |  |  |
| `Segment` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to this resource. |
| `Treatments` |  | `list` | required |  |  |

Supports update: yes

Discovery: not supported
