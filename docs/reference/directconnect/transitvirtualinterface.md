# aws.transitvirtualinterface

**CloudFormation type:** `AWS::DirectConnect::TransitVirtualInterface`

Resource Type definition for AWS::DirectConnect::TransitVirtualInterface

Region attribute: `region`

**Import ID:** `<region>/VirtualInterfaceArn` (AWS::DirectConnect::TransitVirtualInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocateTransitVirtualInterfaceRoleArn` | allocate_transit_virtual_interface_role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of the role to allocate the TransitVifAllocation. Needs directconnect:AllocateTransitVirtualInterface permissions and tag permissions if applicable. |
| `BgpPeers` | bgp_peers | `list` | required |  | The BGP peers configured on this virtual interface.. |
| `ConnectionId` | connection_id | `string` | required | aws.directconnect.connection.ConnectionId | The ID or ARN of the connection or LAG. |
| `DirectConnectGatewayId` | direct_connect_gateway_id | `string` | required, replaces on change | aws.directconnectgateway.DirectConnectGatewayId | The ID or ARN of the Direct Connect gateway. |
| `EnableSiteLink` | enable_site_link | `boolean` | optional, computed, provider-chosen |  | Indicates whether to enable or disable SiteLink. |
| `Mtu` |  | `integer` | optional, computed, provider-chosen |  | The maximum transmission unit (MTU), in bytes. The supported values are 1500 and 9001. The default value is 1500. |
| `RateLimit` | rate_limit | `string` | optional, computed, provider-chosen |  | The rate limit (bandwidth allocation) for the virtual interface. The value must be one of the supported bandwidth values (e.g., 50Mbps, 1Gbps, 10Gbps) and cannot exceed the bandwidth of the parent connection or LAG. |
| `Tags` |  | `map` | tags map |  | The tags associated with the private virtual interface. |
| `VirtualInterfaceArn` | virtual_interface_arn | `string` | computed |  | The ARN of the virtual interface. |
| `VirtualInterfaceId` | virtual_interface_id | `string` | computed |  | The ID of the virtual interface. |
| `VirtualInterfaceName` | virtual_interface_name | `string` | required |  | The name of the virtual interface assigned by the customer network. The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-). |
| `Vlan` |  | `integer` | required, replaces on change |  | The ID of the VLAN. |

Supports update: yes

Discovery: supported
