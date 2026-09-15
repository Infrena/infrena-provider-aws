# aws.transitgatewayroutetable

**CloudFormation type:** `AWS::EC2::TransitGatewayRouteTable`

Resource Type definition for AWS::EC2::TransitGatewayRouteTable

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayRouteTableId` (AWS::EC2::TransitGatewayRouteTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted. |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The ID of the transit gateway. |
| `TransitGatewayRouteTableId` | transit_gateway_route_table_id | `string` | computed |  | Transit Gateway Route Table primary identifier |

Supports update: yes

Discovery: supported
