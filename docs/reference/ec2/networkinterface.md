# aws.networkinterface

**CloudFormation type:** `AWS::EC2::NetworkInterface`

The AWS::EC2::NetworkInterface resource creates network interface

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::NetworkInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionTrackingSpecification` | connection_tracking_specification | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the network interface. |
| `EnablePrimaryIpv6` | enable_primary_ipv6 | `boolean` | optional, computed, provider-chosen |  | If you have instances or ENIs that rely on the IPv6 address not changing, to avoid disrupting traffic to instances or ENIs, you can enable a primary IPv6 address. Enable this option to automatically assign an IPv6 associated with the ENI attached to your instance to be the primary IPv6 address. When you enable an IPv6 address to be a primary IPv6, you cannot disable it. Traffic will be routed to the primary IPv6 address until the instance is terminated or the ENI is detached. If you have multiple IPv6 addresses associated with an ENI and you enable a primary IPv6 address, the first IPv6 address associated with the ENI becomes the primary IPv6 address. |
| `GroupSet` | group_set | `list` | optional, computed, provider-chosen |  | A list of security group IDs associated with this network interface. |
| `Id` |  | `string` | computed |  | Network interface id. |
| `InterfaceType` | interface_type | `string` | optional, computed, provider-chosen, replaces on change |  | Indicates the type of network interface. |
| `Ipv4PrefixCount` | ipv4_prefix_count | `integer` | optional, computed, provider-chosen |  | The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses. |
| `Ipv4Prefixes` | ipv4_prefixes | `list` | optional, computed, provider-chosen |  | Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses. |
| `Ipv6AddressCount` | ipv6_address_count | `integer` | optional, computed, provider-chosen |  | The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property. |
| `Ipv6Addresses` | ipv6_addresses | `list` | optional, computed, provider-chosen |  | One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property. |
| `Ipv6PrefixCount` | ipv6_prefix_count | `integer` | optional, computed, provider-chosen |  | The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses. |
| `Ipv6Prefixes` | ipv6_prefixes | `list` | optional, computed, provider-chosen |  | Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses. |
| `PrimaryIpv6Address` | primary_ipv6_address | `string` | computed |  | The primary IPv6 address |
| `PrimaryPrivateIpAddress` | primary_private_ip_address | `string` | computed |  | Returns the primary private IP address of the network interface. |
| `PrivateIpAddress` | private_ip_address | `string` | optional, computed, provider-chosen, replaces on change |  | Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. |
| `PrivateIpAddresses` | private_ip_addresses | `list` | optional, computed, provider-chosen |  | Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property. |
| `PublicIpDnsHostnameTypeSpecification` | public_ip_dns_hostname_type_specification | `string` | optional, computed, provider-chosen, write-only |  | Public IP DNS hostname type |
| `PublicIpDnsNameOptions` | public_ip_dns_name_options | `map` | computed |  | Describes the public hostname type options, including public hostname type, IPv4-enabled public hostname, IPv6-enabled public hostname, and dual-stack public hostname. |
| `SecondaryPrivateIpAddressCount` | secondary_private_ip_address_count | `integer` | optional, computed, provider-chosen |  | The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses |
| `SecondaryPrivateIpAddresses` | secondary_private_ip_addresses | `list` | computed |  | Returns the secondary private IP addresses of the network interface. |
| `SourceDestCheck` | source_dest_check | `boolean` | optional, computed, provider-chosen |  | Indicates whether traffic to or from the instance is validated. |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet to associate with the network interface. |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this network interface. |
| `VpcId` | vpc_id | `string` | computed |  | The ID of the VPC |

Supports update: yes

Discovery: supported
