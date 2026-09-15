# aws.localgatewayvirtualinterface

**CloudFormation type:** `AWS::EC2::LocalGatewayVirtualInterface`

Resource Type definition for Local Gateway Virtual Interface which describes a virtual interface for AWS Outposts local gateways.

Region attribute: `region`

**Import ID:** `<region>/LocalGatewayVirtualInterfaceId` (AWS::EC2::LocalGatewayVirtualInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigurationState` | configuration_state | `string` | computed |  | The current state of the local gateway virtual interface |
| `LocalAddress` | local_address | `string` | required, replaces on change |  | The local address. |
| `LocalBgpAsn` | local_bgp_asn | `integer` | computed |  | The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP) |
| `LocalGatewayId` | local_gateway_id | `string` | computed |  | The ID of the local gateway |
| `LocalGatewayVirtualInterfaceGroupId` | local_gateway_virtual_interface_group_id | `string` | required, replaces on change | aws.localgatewayvirtualinterfacegroup.LocalGatewayVirtualInterfaceGroupId | The ID of the virtual interface group |
| `LocalGatewayVirtualInterfaceId` | local_gateway_virtual_interface_id | `string` | computed |  | The ID of the virtual interface |
| `OutpostLagId` | outpost_lag_id | `string` | required, replaces on change |  | The Outpost LAG ID. |
| `OwnerId` | owner_id | `string` | computed |  | The ID of the Amazon Web Services account that owns the local gateway virtual interface group |
| `PeerAddress` | peer_address | `string` | required, replaces on change |  | The peer address. |
| `PeerBgpAsn` | peer_bgp_asn | `integer` | optional, computed, provider-chosen, replaces on change |  | The peer BGP ASN. |
| `PeerBgpAsnExtended` | peer_bgp_asn_extended | `integer` | optional, computed, provider-chosen, replaces on change |  | The extended 32-bit ASN of the BGP peer for use with larger ASN values. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Vlan` |  | `integer` | required, replaces on change |  | The ID of the VLAN. |

Supports update: yes

Discovery: supported
