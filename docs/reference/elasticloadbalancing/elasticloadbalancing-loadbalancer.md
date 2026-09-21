# aws.elasticloadbalancing.loadbalancer

**CloudFormation type:** `AWS::ElasticLoadBalancing::LoadBalancer`

Resource Type definition for AWS::ElasticLoadBalancing::LoadBalancer

Region attribute: `region`

**Import ID:** `<region>/LoadBalancerName` (AWS::ElasticLoadBalancing::LoadBalancer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessLoggingPolicy` | access_logging_policy | `map` | optional, computed, provider-chosen |  | Information about where and how access logs are stored for the load balancer. |
| `AppCookieStickinessPolicy` | app_cookie_stickiness_policy | `list` | optional, computed, provider-chosen |  | Information about a policy for application-controlled session stickiness. |
| `AvailabilityZones` | availability_zones | `list` | optional, computed, provider-chosen |  | The Availability Zones for a load balancer in a default VPC. For a load balancer in a nondefault VPC, specify Subnets instead. |
| `CanonicalHostedZoneName` | canonical_hosted_zone_name | `string` | computed |  | The name of the Route 53 hosted zone that is associated with the load balancer. Internal-facing load balancers. |
| `CanonicalHostedZoneNameID` | canonical_hosted_zone_name_id | `string` | computed |  | The ID of the Route 53 hosted zone name that is associated with the load balancer. |
| `ConnectionDrainingPolicy` | connection_draining_policy | `map` | optional, computed, provider-chosen |  | If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance. |
| `ConnectionSettings` | connection_settings | `map` | optional, computed, provider-chosen |  | If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration. |
| `CrossZone` | cross_zone | `boolean` | optional, computed, provider-chosen |  | If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones. |
| `DNSName` | dns_name | `string` | computed |  | The DNS name for the load balancer |
| `HealthCheck` | health_check | `map` | optional, computed, provider-chosen |  | The health check settings to use when evaluating the health of your EC2 instances. |
| `Instances` |  | `list` | optional, computed, provider-chosen |  | The IDs of the instances for the load balancer. |
| `LBCookieStickinessPolicy` | lb_cookie_stickiness_policy | `list` | optional, computed, provider-chosen |  | Information about a policy for duration-based session stickiness. |
| `Listeners` |  | `list` | required |  | The Listeners for the load balancer. You can specify at most one listener per port. |
| `LoadBalancerName` | load_balancer_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the load balancer. This name must be unique within your set of load balancers for the region. |
| `Policies` |  | `list` | optional, computed, provider-chosen |  | The policies defined for your Classic Load Balancer. Specify only back-end server policies. |
| `Scheme` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The type of load balancer. Valid only for load balancers in a VPC. |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen |  | The security groups for the load balancer. Valid only for load balancers in a VPC. |
| `SourceSecurityGroup` | source_security_group | `map` | computed |  |  |
| `Subnets` |  | `list` | optional, computed, provider-chosen |  | The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone. |
| `Tags` |  | `map` | tags map |  | The tags associated with a load balancer. |

Supports update: yes

Discovery: supported
