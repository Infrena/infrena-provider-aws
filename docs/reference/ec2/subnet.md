# aws.subnet

**CloudFormation type:** `AWS::EC2::Subnet`

Specifies a subnet for the specified VPC.

Region attribute: `region`

**Import ID:** `<region>/SubnetId` (AWS::EC2::Subnet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssignIpv6AddressOnCreation` | assign_ipv6_address_on_creation | `boolean` | optional, computed, provider-chosen |  | Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``. |
| `AvailabilityZone` | az, availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone of the subnet. |
| `AvailabilityZoneId` | availability_zone_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AZ ID of the subnet. |
| `BlockPublicAccessStates` | block_public_access_states | `map` | computed |  |  |
| `CidrBlock` | cidr, cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 CIDR block assigned to the subnet. |
| `EnableDns64` | enable_dns64 | `boolean` | optional, computed, provider-chosen |  | Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. |
| `EnableLniAtDeviceIndex` | enable_lni_at_device_index | `integer` | optional, computed, provider-chosen, write-only |  | Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1). |
| `Ipv4IpamPoolId` | ipv4_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | An IPv4 IPAM pool ID for the subnet. |
| `Ipv4NetmaskLength` | ipv4_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | An IPv4 netmask length for the subnet. |
| `Ipv6CidrBlock` | ipv6_cidr_block | `string` | optional, computed, provider-chosen |  | The IPv6 CIDR block. |
| `Ipv6CidrBlocks` | ipv6_cidr_blocks | `list` | computed |  |  |
| `Ipv6IpamPoolId` | ipv6_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | An IPv6 IPAM pool ID for the subnet. |
| `Ipv6Native` | ipv6_native | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*. |
| `Ipv6NetmaskLength` | ipv6_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | An IPv6 netmask length for the subnet. |
| `MapPublicIpOnLaunch` | map_public_ip_on_launch | `boolean` | optional, computed, provider-chosen |  | Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``. |
| `NetworkAclAssociationId` | network_acl_association_id | `string` | computed |  |  |
| `OutpostArn` | outpost_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the Outpost. |
| `PrivateDnsNameOptionsOnLaunch` | private_dns_name_options_on_launch | `map` | optional, computed, provider-chosen |  | The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*. |
| `SubnetId` | subnet_id | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the subnet. |
| `VpcId` | vpc, vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC the subnet is in. |

Supports update: yes

Discovery: supported
