# aws.publicvirtualinterface

**CloudFormation type:** `AWS::DirectConnect::PublicVirtualInterface`

Resource Type definition for AWS::DirectConnect::PublicVirtualInterface

Region attribute: `region`

**Import ID:** `<region>/VirtualInterfaceArn` (AWS::DirectConnect::PublicVirtualInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatePublicVirtualInterfaceRoleArn` | allocate_public_virtual_interface_role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of the role to allocate the public virtual interface. Needs directconnect:AllocatePublicVirtualInterface permissions and tag permissions if applicable. |
| `BgpPeers` | bgp_peers | `list` | required |  | The BGP peers configured on this virtual interface. |
| `ConnectionId` | connection_id | `string` | required | aws.directconnect.connection.ConnectionId | The ID or ARN of the connection or LAG. |
| `RateLimit` | rate_limit | `string` | optional, computed, provider-chosen |  | The rate limit (bandwidth allocation) for the virtual interface. The value must be one of the supported bandwidth values (e.g., 50Mbps, 1Gbps, 10Gbps) and cannot exceed the bandwidth of the parent connection or LAG. |
| `RouteFilterPrefixes` | route_filter_prefixes | `list` | optional, computed, provider-chosen, replaces on change |  | The routes to be advertised to the AWS network in this region. |
| `Tags` |  | `map` | tags map |  | The tags associated with the public virtual interface. |
| `VirtualInterfaceArn` | virtual_interface_arn | `string` | computed |  | The ARN of the virtual interface. |
| `VirtualInterfaceId` | virtual_interface_id | `string` | computed |  | The ID of the virtual interface. |
| `VirtualInterfaceName` | virtual_interface_name | `string` | required |  | The name of the virtual interface assigned by the customer network. The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-). |
| `Vlan` |  | `integer` | required, replaces on change |  | The ID of the VLAN. |

Supports update: yes

Discovery: supported
