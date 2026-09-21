# aws.transitgatewaypeeringattachment

**CloudFormation type:** `AWS::EC2::TransitGatewayPeeringAttachment`

The AWS::EC2::TransitGatewayPeeringAttachment type

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayAttachmentId` (AWS::EC2::TransitGatewayPeeringAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time the transit gateway peering attachment was created. |
| `PeerAccountId` | peer_account_id | `string` | required, replaces on change |  | The ID of the peer account |
| `PeerRegion` | peer_region | `string` | required, replaces on change |  | Peer Region |
| `PeerTransitGatewayId` | peer_transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The ID of the peer transit gateway. |
| `State` |  | `string` | computed |  | The state of the transit gateway peering attachment. Note that the initiating state has been deprecated. |
| `Status` |  | `map` | computed |  | The status of the transit gateway peering attachment. |
| `Tags` |  | `map` | tags map |  | The tags for the transit gateway peering attachment. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | computed |  | The ID of the transit gateway peering attachment. |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The ID of the transit gateway. |

Supports update: yes

Discovery: supported
