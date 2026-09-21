# aws.localgatewayroutetablevpcassociation

**CloudFormation type:** `AWS::EC2::LocalGatewayRouteTableVPCAssociation`

Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.

Region attribute: `region`

**Import ID:** `<region>/LocalGatewayRouteTableVpcAssociationId` (AWS::EC2::LocalGatewayRouteTableVPCAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocalGatewayId` | local_gateway_id | `string` | computed |  | The ID of the local gateway. |
| `LocalGatewayRouteTableId` | local_gateway_route_table_id | `string` | required, replaces on change | aws.localgatewayroutetable.LocalGatewayRouteTableId | The ID of the local gateway route table. |
| `LocalGatewayRouteTableVpcAssociationId` | local_gateway_route_table_vpc_association_id | `string` | computed |  | The ID of the association. |
| `State` |  | `string` | computed |  | The state of the association. |
| `Tags` |  | `map` | tags map |  | The tags for the association. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: yes

Discovery: supported
