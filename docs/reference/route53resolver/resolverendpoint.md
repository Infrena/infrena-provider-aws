# aws.resolverendpoint

**CloudFormation type:** `AWS::Route53Resolver::ResolverEndpoint`

Resource type definition for AWS::Route53Resolver::ResolverEndpoint

Region attribute: `region`

**Import ID:** `<region>/ResolverEndpointId` (AWS::Route53Resolver::ResolverEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the resolver endpoint, such as arn:aws:route53resolver:us-east-1:123456789012:resolver-endpoint/resolver-endpoint-a1bzhi. |
| `Direction` |  | `string` | required, replaces on change |  | Indicates whether the Resolver endpoint allows inbound or outbound DNS queries: |
| `Dns64Enabled` | dns64_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether DNS64 is enabled for the Inbound Resolver Endpoint. When set to true, if a DNS AAAA query is made for a domain that has only an A (IPv4) record, the resolver automatically synthesizes an AAAA (IPv6) response by embedding the IPv4 address into the well-known prefix 64:ff9b::/96. Default is false. |
| `HostVPCId` | host_vpc_id | `string` | computed |  | The ID of the VPC that you want to create the resolver endpoint in. |
| `IpAddressCount` | ip_address_count | `string` | computed |  | The number of IP addresses that the resolver endpoint can use for DNS queries. |
| `IpAddresses` | ip_addresses | `list` | required |  | The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints). The subnet ID uniquely identifies a VPC. |
| `Ipv6InternetAccessEnabled` | ipv6_internet_access_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether IPv6 Internet Gateway access is enabled through the Outbound Resolver Endpoint. When set to true, this property allows your Endpoint ENIs to reach public IPv6 target nameservers through an internet gateway. Default is false. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console. |
| `OutpostArn` | outpost_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN (Amazon Resource Name) for the Outpost. |
| `PreferredInstanceType` | preferred_instance_type | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon EC2 instance type. |
| `Protocols` |  | `list` | optional, computed, provider-chosen |  | Protocols used for the endpoint. DoH-FIPS is applicable for inbound endpoints only. |
| `ResolverEndpointId` | resolver_endpoint_id | `string` | computed |  | The ID of the resolver endpoint. |
| `ResolverEndpointType` | resolver_endpoint_type | `string` | optional, computed, provider-chosen |  | The Resolver endpoint IP address type. |
| `RniEnhancedMetricsEnabled` | rni_enhanced_metrics_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether RNI enhanced metrics are enabled for the Resolver Endpoints. When set to true, one-minute granular metrics are published in CloudWatch for each RNI associated with this endpoint. When set to false, metrics are not published. Default is false. |
| `SecurityGroupIds` | security_group_ids | `list` | required, replaces on change | aws.securitygroup.Id | The ID of one or more security groups that control access to this VPC. The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetNameServerMetricsEnabled` | target_name_server_metrics_enabled | `boolean` | optional, computed, provider-chosen |  | Specifies whether target name server metrics are enabled for the Outbound Resolver Endpoint. When set to true, one-minute granular metrics are published in CloudWatch for each target name server associated with this endpoint. When set to false, metrics are not published. Default is false. |

Supports update: yes

Discovery: supported
