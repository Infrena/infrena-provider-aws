# aws.observabilityconfiguration

**CloudFormation type:** `AWS::AppRunner::ObservabilityConfiguration`

The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration

Region attribute: `region`

**Import ID:** `<region>/ObservabilityConfigurationArn` (AWS::AppRunner::ObservabilityConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Latest` |  | `boolean` | computed |  | It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise. |
| `ObservabilityConfigurationArn` | observability_configuration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of this ObservabilityConfiguration |
| `ObservabilityConfigurationName` | observability_configuration_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. |
| `ObservabilityConfigurationRevision` | observability_configuration_revision | `integer` | computed |  | The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, write-only, tags map |  | A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair. |
| `TraceConfiguration` | trace_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Describes the configuration of the tracing feature within an AWS App Runner observability configuration. |

Supports update: no

Discovery: supported
