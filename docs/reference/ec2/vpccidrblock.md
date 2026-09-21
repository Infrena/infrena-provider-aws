# aws.vpccidrblock

**CloudFormation type:** `AWS::EC2::VPCCidrBlock`

Resource Type definition for AWS::EC2::VPCCidrBlock

Region attribute: `region`

**Import ID:** `<region>/Id|VpcId` (AWS::EC2::VPCCidrBlock)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonProvidedIpv6CidrBlock` | amazon_provided_ipv6_cidr_block | `boolean` | optional, computed, provider-chosen, replaces on change |  | Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IPv6 addresses, or the size of the CIDR block. |
| `CidrBlock` | cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | An IPv4 CIDR block to associate with the VPC. |
| `Id` |  | `string` | computed |  | The Id of the VPC associated CIDR Block. |
| `IpSource` | ip_source | `string` | computed |  | The IP Source of an IPv6 VPC CIDR Block. |
| `Ipv4IpamPoolId` | ipv4_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | The ID of the IPv4 IPAM pool to Associate a CIDR from to a VPC. |
| `Ipv4NetmaskLength` | ipv4_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. |
| `Ipv6AddressAttribute` | ipv6_address_attribute | `string` | computed |  | The value denoting whether an IPv6 VPC CIDR Block is public or private. |
| `Ipv6CidrBlock` | ipv6_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | An IPv6 CIDR block from the IPv6 address pool. |
| `Ipv6CidrBlockNetworkBorderGroup` | ipv6_cidr_block_network_border_group | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the location from which we advertise the IPV6 CIDR block. |
| `Ipv6IpamPoolId` | ipv6_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | The ID of the IPv6 IPAM pool to Associate a CIDR from to a VPC. |
| `Ipv6NetmaskLength` | ipv6_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. |
| `Ipv6Pool` | ipv6_pool | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: no

Discovery: supported (parent resource required)
