# aws.eip

**CloudFormation type:** `AWS::EC2::EIP`

Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.

Region attribute: `region`

**Import ID:** `<region>/PublicIp|AllocationId` (AWS::EC2::EIP)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Address` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | An Elastic IP address or a carrier IP address in a Wavelength Zone. |
| `AllocationId` | allocation_id | `string` | computed |  |  |
| `Domain` |  | `string` | optional, computed, provider-chosen |  | The network (``vpc``). |
| `InstanceId` | instance_id | `string` | optional, computed, provider-chosen | aws.ec2.instance.InstanceId | The ID of the instance. |
| `IpamPoolId` | ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipampool.IpamPoolId | The ID of an IPAM pool which has an Amazon-provided or BYOIP public IPv4 CIDR provisioned to it. For more information, see [Allocate sequential Elastic IP addresses from an IPAM pool](https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-eip-pool.html) in the *Amazon VPC IPAM User Guide*. |
| `NetworkBorderGroup` | network_border_group | `string` | optional, computed, provider-chosen, replaces on change |  | A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups. |
| `PublicIp` | public_ip | `string` | computed |  |  |
| `PublicIpv4Pool` | public_ipv4_pool | `string` | optional, computed, provider-chosen |  | The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool. |
| `Tags` |  | `map` | tags map |  | Any tags assigned to the Elastic IP address. |
| `TransferAddress` | transfer_address | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Elastic IP address you are accepting for transfer. You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide*. |

Supports update: yes

Discovery: supported
