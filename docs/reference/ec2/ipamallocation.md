# aws.ipamallocation

**CloudFormation type:** `AWS::EC2::IPAMAllocation`

Resource Schema of AWS::EC2::IPAMAllocation Type

Region attribute: `region`

**Import ID:** `<region>/IpamPoolId|IpamPoolAllocationId|Cidr` (AWS::EC2::IPAMAllocation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Cidr` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Represents an IPAM custom allocation of a single IPv4 or IPv6 CIDR |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `IpamPoolAllocationId` | ipam_pool_allocation_id | `string` | computed |  | Id of the allocation. |
| `IpamPoolId` | ipam_pool_id | `string` | required, replaces on change | aws.ipampool.IpamPoolId | Id of the IPAM Pool. |
| `NetmaskLength` | netmask_length | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it. |

Supports update: no

Discovery: supported (parent resource required)
