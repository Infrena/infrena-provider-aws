# aws.autoscalingconfiguration

**CloudFormation type:** `AWS::AppRunner::AutoScalingConfiguration`

Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.

Region attribute: `region`

**Import ID:** `<region>/AutoScalingConfigurationArn` (AWS::AppRunner::AutoScalingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingConfigurationArn` | auto_scaling_configuration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of this auto scaling configuration. |
| `AutoScalingConfigurationName` | auto_scaling_configuration_name | `string` | optional, computed, provider-chosen, replaces on change |  | The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration. |
| `AutoScalingConfigurationRevision` | auto_scaling_configuration_revision | `integer` | computed |  | The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName. |
| `Latest` |  | `boolean` | computed |  | It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code. |
| `MaxConcurrency` | max_concurrency | `integer` | optional, computed, provider-chosen, replaces on change |  | The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests. |
| `MaxSize` | max_size | `integer` | optional, computed, provider-chosen, replaces on change |  | The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service. |
| `MinSize` | min_size | `integer` | optional, computed, provider-chosen, replaces on change |  | The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset. |
| `Tags` |  | `map` | replaces on change, tags map |  | A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair. |

Supports update: no

Discovery: supported
