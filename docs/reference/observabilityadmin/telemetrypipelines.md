# aws.telemetrypipelines

**CloudFormation type:** `AWS::ObservabilityAdmin::TelemetryPipelines`

Resource Type definition for AWS::ObservabilityAdmin::TelemetryPipelines

Region attribute: `region`

**Import ID:** `<region>/PipelineIdentifier` (AWS::ObservabilityAdmin::TelemetryPipelines)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Configuration` |  | `map` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Pipeline` |  | `map` | computed |  |  |
| `PipelineIdentifier` | pipeline_identifier | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusReason` | status_reason | `map` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource |

Supports update: yes

Discovery: supported
