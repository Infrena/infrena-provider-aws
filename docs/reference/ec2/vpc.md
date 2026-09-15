# aws.vpc

**CloudFormation type:** `AWS::EC2::VPC`

Specifies a virtual private cloud (VPC).

Region attribute: `region`

**Import ID:** `<region>/VpcId` (AWS::EC2::VPC)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CidrBlock` | cidr, cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 network range for the VPC, in CIDR notation. For example, ``10.0.0.0/16``. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``. |
| `CidrBlockAssociations` | cidr_block_associations | `list` | computed |  |  |
| `DefaultNetworkAcl` | default_network_acl | `string` | computed |  |  |
| `DefaultSecurityGroup` | default_security_group | `string` | computed |  |  |
| `EnableDnsHostnames` | enable_dns_hostnames | `boolean` | optional, computed, provider-chosen |  | Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support). |
| `EnableDnsSupport` | enable_dns_support | `boolean` | optional, computed, provider-chosen |  | Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support). |
| `InstanceTenancy` | instance_tenancy | `string` | optional, computed, provider-chosen |  | The allowed tenancy of instances launched into the VPC. |
| `Ipv4IpamPoolId` | ipv4_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide*. |
| `Ipv4NetmaskLength` | ipv4_netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide*. |
| `Ipv6CidrBlocks` | ipv6_cidr_blocks | `list` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the VPC. |
| `VpcEncryptionControl` | vpc_encryption_control | `map` | optional, computed, provider-chosen, replaces on change |  | Describes the configuration and state of VPC encryption controls. |
| `VpcId` | vpc_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
