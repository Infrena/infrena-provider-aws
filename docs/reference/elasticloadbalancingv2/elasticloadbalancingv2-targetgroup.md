# aws.elasticloadbalancingv2.targetgroup

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::TargetGroup`

Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup

Region attribute: `region`

**Import ID:** `<region>/TargetGroupArn` (AWS::ElasticLoadBalancingV2::TargetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HealthCheckEnabled` | health_check_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled. |
| `HealthCheckIntervalSeconds` | health_check_interval_seconds | `integer` | optional, computed, provider-chosen |  | The approximate amount of time, in seconds, between health checks of an individual target. |
| `HealthCheckPath` | health_check_path | `string` | optional, computed, provider-chosen |  | [HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck. |
| `HealthCheckPort` | health_check_port | `string` | optional, computed, provider-chosen |  | The port the load balancer uses when performing health checks on targets. |
| `HealthCheckProtocol` | health_check_protocol | `string` | optional, computed, provider-chosen |  | The protocol the load balancer uses when performing health checks on targets. |
| `HealthCheckTimeoutSeconds` | health_check_timeout_seconds | `integer` | optional, computed, provider-chosen |  | The amount of time, in seconds, during which no response from a target means a failed health check. |
| `HealthyThresholdCount` | healthy_threshold_count | `integer` | optional, computed, provider-chosen |  | The number of consecutive health checks successes required before considering an unhealthy target healthy. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of IP address used for this target group. The possible values are ipv4 and ipv6. |
| `LoadBalancerArns` | load_balancer_arns | `list` | computed |  | The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group. |
| `Matcher` |  | `map` | optional, computed, provider-chosen |  | [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the target group. |
| `Port` |  | `integer` | optional, computed, provider-chosen, replaces on change |  | The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081. |
| `Protocol` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The protocol to use for routing traffic to the targets. |
| `ProtocolVersion` | protocol_version | `string` | optional, computed, provider-chosen, replaces on change |  | [HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2. |
| `Tags` |  | `map` | tags map |  | The tags. |
| `TargetControlPort` | target_control_port | `integer` | optional, computed, provider-chosen |  | The port that the target control agent uses to communicate the available capacity of targets to the load balancer. |
| `TargetGroupArn` | target_group_arn | `string` | computed |  | The ARN of the Target Group |
| `TargetGroupAttributes` | target_group_attributes | `list` | optional, computed, provider-chosen |  | The attributes. |
| `TargetGroupFullName` | target_group_full_name | `string` | computed |  | The full name of the target group. |
| `TargetGroupName` | target_group_name | `string` | computed |  | The name of the target group. |
| `TargetType` | target_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of target that you must specify when registering targets with this target group. You can't specify targets for a target group using more than one target type. |
| `Targets` |  | `list` | optional, computed, provider-chosen |  | The targets. |
| `UnhealthyThresholdCount` | unhealthy_threshold_count | `integer` | optional, computed, provider-chosen |  | The number of consecutive health check failures required before considering a target unhealthy. |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId | The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply. |

Supports update: yes

Discovery: supported
