# aws.ipampoolcidr

**CloudFormation type:** `AWS::EC2::IPAMPoolCidr`

Resource Schema of AWS::EC2::IPAMPoolCidr Type

Region attribute: `region`

**Import ID:** `<region>/IpamPoolId|IpamPoolCidrId` (AWS::EC2::IPAMPoolCidr)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Cidr` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Represents a single IPv4 or IPv6 CIDR |
| `IpamPoolCidrId` | ipam_pool_cidr_id | `string` | computed |  | Id of the IPAM Pool Cidr. |
| `IpamPoolId` | ipam_pool_id | `string` | required, replaces on change | aws.ipampool.IpamPoolId | Id of the IPAM Pool. |
| `NetmaskLength` | netmask_length | `integer` | optional, computed, provider-chosen, replaces on change |  | The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it. |
| `State` |  | `string` | computed |  | Provisioned state of the cidr. |

Supports update: no

Discovery: supported (parent resource required)
