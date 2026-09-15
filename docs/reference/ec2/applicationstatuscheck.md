# aws.applicationstatuscheck

**CloudFormation type:** `AWS::EC2::ApplicationStatusCheck`

An application status check monitors an HTTP or HTTPS endpoint on Amazon EC2 instances and reports application-layer health. Configure a health check with a Protocol, Port, and Path, and then associate it with EC2 instances via AWS::EC2::ApplicationStatusCheckInstanceAssociation or AWS::EC2::ApplicationStatusCheckTagAssociation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::EC2::ApplicationStatusCheck)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Aggregation` |  | `string` | optional, computed, provider-chosen |  | Whether this check is included in the rolled-up application status. |
| `ApplicationStatusCheckId` | application_status_check_id | `string` | computed |  | The unique identifier of the application status check. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the application status check. |
| `CreationTime` | creation_time | `string` | computed |  | When the application status check was created (ISO 8601). |
| `DeviceIndex` | device_index | `integer` | optional, computed, provider-chosen |  | The network interface device index used for the health check. |
| `FailureThreshold` | failure_threshold | `integer` | optional, computed, provider-chosen |  | The number of consecutive failed probes required to mark the instance unhealthy. |
| `HealthCheckPaths` | health_check_paths | `list` | optional, computed, provider-chosen |  | The source/destination network paths used for the health check. |
| `InitializationGracePeriodSeconds` | initialization_grace_period_seconds | `integer` | optional, computed, provider-chosen |  | Seconds to wait after instance launch before beginning health checks. |
| `Interval` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The interval, in seconds, between health check probes. |
| `IpScope` | ip_scope | `string` | optional, computed, provider-chosen, replaces on change |  | The IP scope used for the health check. |
| `IpVersion` | ip_version | `string` | optional, computed, provider-chosen |  | The IP version used for the health check. |
| `Path` |  | `string` | optional, computed, provider-chosen |  | The HTTP path used for the health check. |
| `Port` |  | `integer` | required |  | The port used for the health check. |
| `Protocol` |  | `string` | required |  | The network protocol used for the health check. |
| `StatusCodeMatcher` | status_code_matcher | `string` | optional, computed, provider-chosen |  | The HTTP status codes considered successful (e.g., "200-299"). |
| `SuccessThreshold` | success_threshold | `integer` | optional, computed, provider-chosen |  | The number of consecutive successful probes required to mark the instance healthy. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the application status check. |
| `Timeout` |  | `integer` | optional, computed, provider-chosen |  | The timeout, in seconds, for each health check probe. |

Supports update: yes

Discovery: supported
