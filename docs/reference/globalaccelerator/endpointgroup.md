# aws.endpointgroup

**CloudFormation type:** `AWS::GlobalAccelerator::EndpointGroup`

Resource Type definition for AWS::GlobalAccelerator::EndpointGroup

Region attribute: `region`

**Import ID:** `<region>/EndpointGroupArn` (AWS::GlobalAccelerator::EndpointGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EndpointConfigurations` | endpoint_configurations | `list` | optional, computed, provider-chosen |  | The list of endpoint objects. |
| `EndpointGroupArn` | endpoint_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the endpoint group |
| `EndpointGroupRegion` | endpoint_group_region | `string` | required, replaces on change |  | The name of the AWS Region where the endpoint group is located |
| `HealthCheckIntervalSeconds` | health_check_interval_seconds | `integer` | optional, computed, provider-chosen |  | The time in seconds between each health check for an endpoint. Must be a value of 10 or 30 |
| `HealthCheckPath` | health_check_path | `string` | optional, computed, provider-chosen |  |  |
| `HealthCheckPort` | health_check_port | `integer` | optional, computed, provider-chosen |  | The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group. |
| `HealthCheckProtocol` | health_check_protocol | `string` | optional, computed, provider-chosen |  | The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group. |
| `ListenerArn` | listener_arn | `string` | required, replaces on change | aws.globalaccelerator.listener.ListenerArn | The Amazon Resource Name (ARN) of the listener |
| `PortOverrides` | port_overrides | `list` | optional, computed, provider-chosen |  |  |
| `ThresholdCount` | threshold_count | `integer` | optional, computed, provider-chosen |  | The number of consecutive health checks required to set the state of the endpoint to unhealthy. |
| `TrafficDialPercentage` | traffic_dial_percentage | `float` | optional, computed, provider-chosen |  | The percentage of traffic to sent to an AWS Region |

Supports update: yes

Discovery: supported
