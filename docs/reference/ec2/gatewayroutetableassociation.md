# aws.gatewayroutetableassociation

**CloudFormation type:** `AWS::EC2::GatewayRouteTableAssociation`

Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.

Region attribute: `region`

**Import ID:** `<region>/GatewayId` (AWS::EC2::GatewayRouteTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | The route table association ID. |
| `GatewayId` | gateway_id | `string` | required, replaces on change |  | The ID of the gateway. |
| `RouteTableId` | route_table_id | `string` | required | aws.routetable.RouteTableId | The ID of the route table. |

Supports update: yes

Discovery: not supported
