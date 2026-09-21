# aws.transitgatewayroute

**CloudFormation type:** `AWS::EC2::TransitGatewayRoute`

Resource Type definition for AWS::EC2::TransitGatewayRoute

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayRouteTableId|DestinationCidrBlock` (AWS::EC2::TransitGatewayRoute)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Blackhole` |  | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether to drop traffic that matches this route. |
| `DestinationCidrBlock` | destination_cidr_block | `string` | required, replaces on change |  | The CIDR range used for destination matches. Routing decisions are based on the most specific match. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | optional, computed, provider-chosen, replaces on change | aws.transitgatewayattachment.Id | The ID of transit gateway attachment. |
| `TransitGatewayRouteTableId` | transit_gateway_route_table_id | `string` | required, replaces on change | aws.transitgatewayroutetable.TransitGatewayRouteTableId | The ID of transit gateway route table. |

Supports update: no

Discovery: supported (parent resource required)
