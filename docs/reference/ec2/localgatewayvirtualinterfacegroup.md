# aws.localgatewayvirtualinterfacegroup

**CloudFormation type:** `AWS::EC2::LocalGatewayVirtualInterfaceGroup`

Resource Type definition for LocalGatewayVirtualInterfaceGroup which describes a group of LocalGateway VirtualInterfaces

Region attribute: `region`

**Import ID:** `<region>/LocalGatewayVirtualInterfaceGroupId` (AWS::EC2::LocalGatewayVirtualInterfaceGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigurationState` | configuration_state | `string` | computed |  | The current state of the local gateway virtual interface group |
| `LocalBgpAsn` | local_bgp_asn | `integer` | optional, computed, provider-chosen, replaces on change |  | The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP) |
| `LocalBgpAsnExtended` | local_bgp_asn_extended | `integer` | optional, computed, provider-chosen, replaces on change |  | The extended 32-bit ASN for the local BGP configuration |
| `LocalGatewayId` | local_gateway_id | `string` | required, replaces on change |  | The ID of the local gateway |
| `LocalGatewayVirtualInterfaceGroupArn` | local_gateway_virtual_interface_group_arn | `string` | computed |  | The Amazon Resource Number (ARN) of the local gateway virtual interface group |
| `LocalGatewayVirtualInterfaceGroupId` | local_gateway_virtual_interface_group_id | `string` | computed |  | The ID of the virtual interface group |
| `LocalGatewayVirtualInterfaceIds` | local_gateway_virtual_interface_ids | `list` | computed |  | The IDs of the virtual interfaces |
| `OwnerId` | owner_id | `string` | computed |  | The ID of the Amazon Web Services account that owns the local gateway virtual interface group |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the virtual interface group |

Supports update: yes

Discovery: supported
