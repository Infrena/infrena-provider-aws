# aws.subnetcidrblock

**CloudFormation type:** `AWS::EC2::SubnetCidrBlock`

The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SubnetCidrBlock)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | Information about the IPv6 association. |
| `IpSource` | ip_source | `string` | computed |  | The IP Source of an IPv6 Subnet CIDR Block. |
| `Ipv6AddressAttribute` | ipv6_address_attribute | `string` | computed |  | The value denoting whether an IPv6 Subnet CIDR Block is public or private. |
| `Ipv6CidrBlock` | ipv6_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length |
| `Ipv6IpamPoolId` | ipv6_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR |
| `Ipv6NetmaskLength` | ipv6_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet |

Supports update: no

Discovery: supported
