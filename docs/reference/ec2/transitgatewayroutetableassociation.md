# aws.transitgatewayroutetableassociation

**CloudFormation type:** `AWS::EC2::TransitGatewayRouteTableAssociation`

Resource Type definition for AWS::EC2::TransitGatewayRouteTableAssociation

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayRouteTableId|TransitGatewayAttachmentId` (AWS::EC2::TransitGatewayRouteTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | required, replaces on change | aws.transitgatewayattachment.Id | The ID of transit gateway attachment. |
| `TransitGatewayRouteTableId` | transit_gateway_route_table_id | `string` | required, replaces on change | aws.transitgatewayroutetable.TransitGatewayRouteTableId | The ID of transit gateway route table. |

Supports update: no

Discovery: supported (parent resource required)
