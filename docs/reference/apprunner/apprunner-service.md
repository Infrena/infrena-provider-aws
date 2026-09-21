# aws.apprunner.service

**CloudFormation type:** `AWS::AppRunner::Service`

The AWS::AppRunner::Service resource specifies an AppRunner Service.

Region attribute: `region`

**Import ID:** `<region>/ServiceArn` (AWS::AppRunner::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingConfigurationArn` | auto_scaling_configuration_arn | `string` | optional, computed, provider-chosen, write-only | aws.autoscalingconfiguration.AutoScalingConfigurationArn | Autoscaling configuration ARN |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Encryption configuration (KMS key) |
| `HealthCheckConfiguration` | health_check_configuration | `map` | optional, computed, provider-chosen |  | Health check configuration |
| `InstanceConfiguration` | instance_configuration | `map` | optional, computed, provider-chosen |  | Instance Configuration |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  | Network configuration |
| `ObservabilityConfiguration` | observability_configuration | `map` | optional, computed, provider-chosen |  | Service observability configuration |
| `ServiceArn` | service_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the AppRunner Service. |
| `ServiceId` | service_id | `string` | computed |  | The AppRunner Service Id |
| `ServiceName` | service_name | `string` | optional, computed, provider-chosen, replaces on change |  | The AppRunner Service Name. |
| `ServiceUrl` | service_url | `string` | computed |  | The Service Url of the AppRunner Service. |
| `SourceConfiguration` | source_configuration | `map` | required |  | Source Code configuration |
| `Status` |  | `string` | computed |  | AppRunner Service status. |
| `Tags` |  | `map` | replaces on change, write-only, tags map |  |  |

Supports update: yes

Discovery: supported
