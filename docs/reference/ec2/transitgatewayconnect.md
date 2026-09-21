# aws.transitgatewayconnect

**CloudFormation type:** `AWS::EC2::TransitGatewayConnect`

The AWS::EC2::TransitGatewayConnect type

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayAttachmentId` (AWS::EC2::TransitGatewayConnect)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The creation time. |
| `Options` |  | `map` | required, replaces on change |  | The Connect attachment options. |
| `State` |  | `string` | computed |  | The state of the attachment. |
| `Tags` |  | `map` | tags map |  | The tags for the attachment. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | computed |  | The ID of the Connect attachment. |
| `TransitGatewayId` | transit_gateway_id | `string` | computed |  | The ID of the transit gateway. |
| `TransportTransitGatewayAttachmentId` | transport_transit_gateway_attachment_id | `string` | required, replaces on change | aws.transitgatewayattachment.Id | The ID of the attachment from which the Connect attachment was created. |

Supports update: yes

Discovery: supported
