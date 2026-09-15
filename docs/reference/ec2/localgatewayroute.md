# aws.localgatewayroute

**CloudFormation type:** `AWS::EC2::LocalGatewayRoute`

Resource Type definition for Local Gateway Route which describes a route for a local gateway route table.

Region attribute: `region`

**Import ID:** `<region>/DestinationCidrBlock|LocalGatewayRouteTableId` (AWS::EC2::LocalGatewayRoute)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DestinationCidrBlock` | destination_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The CIDR block used for destination matches. |
| `LocalGatewayRouteTableId` | local_gateway_route_table_id | `string` | optional, computed, provider-chosen, replaces on change | aws.localgatewayroutetable.LocalGatewayRouteTableId | The ID of the local gateway route table. |
| `LocalGatewayVirtualInterfaceGroupId` | local_gateway_virtual_interface_group_id | `string` | optional, computed, provider-chosen | aws.localgatewayvirtualinterfacegroup.LocalGatewayVirtualInterfaceGroupId | The ID of the virtual interface group. |
| `NetworkInterfaceId` | network_interface_id | `string` | optional, computed, provider-chosen | aws.networkinterface.Id | The ID of the network interface. |
| `State` |  | `string` | computed |  | The state of the route. |
| `Type` | type_value | `string` | computed |  | The route type. |

Supports update: yes

Discovery: supported
