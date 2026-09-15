# aws.localgatewayroutetablevirtualinterfacegroupassociation

**CloudFormation type:** `AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation`

Resource Type definition for Local Gateway Route Table Virtual Interface Group Association which describes a local gateway route table virtual interface group association for a local gateway.

Region attribute: `region`

**Import ID:** `<region>/LocalGatewayRouteTableVirtualInterfaceGroupAssociationId` (AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocalGatewayId` | local_gateway_id | `string` | computed |  | The ID of the local gateway. |
| `LocalGatewayRouteTableArn` | local_gateway_route_table_arn | `string` | computed |  | The ARN of the local gateway route table. |
| `LocalGatewayRouteTableId` | local_gateway_route_table_id | `string` | required, replaces on change | aws.localgatewayroutetable.LocalGatewayRouteTableId | The ID of the local gateway route table. |
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociationId` | local_gateway_route_table_virtual_interface_group_association_id | `string` | computed |  | The ID of the local gateway route table virtual interface group association. |
| `LocalGatewayVirtualInterfaceGroupId` | local_gateway_virtual_interface_group_id | `string` | required, replaces on change | aws.localgatewayvirtualinterfacegroup.LocalGatewayVirtualInterfaceGroupId | The ID of the local gateway route table virtual interface group. |
| `OwnerId` | owner_id | `string` | computed |  | The owner of the local gateway route table virtual interface group association. |
| `State` |  | `string` | computed |  | The state of the local gateway route table virtual interface group association. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the local gateway route table virtual interface group association. |

Supports update: yes

Discovery: supported
