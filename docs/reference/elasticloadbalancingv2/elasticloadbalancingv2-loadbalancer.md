# aws.elasticloadbalancingv2.loadbalancer

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::LoadBalancer`

Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.

Region attribute: `region`

**Import ID:** `<region>/LoadBalancerArn` (AWS::ElasticLoadBalancingV2::LoadBalancer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CanonicalHostedZoneID` | canonical_hosted_zone_id | `string` | computed |  |  |
| `DNSName` | dns_name | `string` | computed |  |  |
| `EnableCapacityReservationProvisionStabilize` | enable_capacity_reservation_provision_stabilize | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates whether to enable stabilization when creating or updating an LCU reservation. This ensures that the final stack status reflects the status of the LCU reservation. The default is ``false``. |
| `EnablePrefixForIpv6SourceNat` | enable_prefix_for_ipv6_source_nat | `string` | optional, computed, provider-chosen |  | [Network Load Balancers with UDP listeners] Indicates whether to use an IPv6 prefix from each subnet for source NAT. The IP address type must be ``dualstack``. The default value is ``off``. |
| `EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic` | enforce_security_group_inbound_rules_on_private_link_traffic | `string` | optional, computed, provider-chosen |  | Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink. The default is ``on``. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | The IP address type. Internal load balancers must use ``ipv4``. |
| `Ipv4IpamPoolId` | ipv4_ipam_pool_id | `string` | optional, computed, provider-chosen | aws.ipampool.IpamPoolId | The ID of the IPv4 IPAM pool. |
| `LoadBalancerArn` | load_balancer_arn | `string` | computed |  |  |
| `LoadBalancerAttributes` | load_balancer_attributes | `list` | optional, computed, provider-chosen |  | The load balancer attributes. Attributes that you do not modify retain their current values. |
| `LoadBalancerFullName` | load_balancer_full_name | `string` | computed |  |  |
| `LoadBalancerName` | load_balancer_name | `string` | computed |  |  |
| `MinimumLoadBalancerCapacity` | minimum_load_balancer_capacity | `map` | optional, computed, provider-chosen |  | The minimum capacity for a load balancer. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-". |
| `Scheme` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet. |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen |  | [Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer. |
| `SubnetMappings` | subnet_mappings | `list` | optional, computed, provider-chosen |  | The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. |
| `Subnets` |  | `list` | optional, computed, provider-chosen |  | The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets. |
| `Tags` |  | `map` | tags map |  | The tags to assign to the load balancer. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The type of load balancer. The default is ``application``. |

Supports update: yes

Discovery: supported
