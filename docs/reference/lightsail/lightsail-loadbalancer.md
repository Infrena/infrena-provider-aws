# aws.lightsail.loadbalancer

**CloudFormation type:** `AWS::Lightsail::LoadBalancer`

Resource Type definition for AWS::Lightsail::LoadBalancer

Region attribute: `region`

**Import ID:** `<region>/LoadBalancerName` (AWS::Lightsail::LoadBalancer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachedInstances` | attached_instances | `list` | optional, computed, provider-chosen |  | The names of the instances attached to the load balancer. |
| `HealthCheckPath` | health_check_path | `string` | optional, computed, provider-chosen |  | The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/"). |
| `InstancePort` | instance_port | `integer` | required, replaces on change |  | The instance port where you're creating your load balancer. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  | The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack. |
| `LoadBalancerArn` | load_balancer_arn | `string` | computed |  |  |
| `LoadBalancerName` | load_balancer_name | `string` | required, replaces on change |  | The name of your load balancer. |
| `SessionStickinessEnabled` | session_stickiness_enabled | `boolean` | optional, computed, provider-chosen |  | Configuration option to enable session stickiness. |
| `SessionStickinessLBCookieDurationSeconds` | session_stickiness_lb_cookie_duration_seconds | `string` | optional, computed, provider-chosen |  | Configuration option to adjust session stickiness cookie duration parameter. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TlsPolicyName` | tls_policy_name | `string` | optional, computed, provider-chosen |  | The name of the TLS policy to apply to the load balancer. |

Supports update: yes

Discovery: supported
