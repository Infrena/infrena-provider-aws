# aws.launch

**CloudFormation type:** `AWS::Evidently::Launch`

Resource Type definition for AWS::Evidently::Launch.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Evidently::Launch)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExecutionStatus` | execution_status | `map` | optional, computed, provider-chosen |  | Start or Stop Launch Launch. Default is not started. |
| `Groups` |  | `list` | required |  |  |
| `MetricMonitors` | metric_monitors | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Project` |  | `string` | required, replaces on change |  |  |
| `RandomizationSalt` | randomization_salt | `string` | optional, computed, provider-chosen |  |  |
| `ScheduledSplitsConfig` | scheduled_splits_config | `list` | required |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: not supported
