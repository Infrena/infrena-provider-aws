# aws.transitgatewayconnectpeer

**CloudFormation type:** `AWS::EC2::TransitGatewayConnectPeer`

Resource Type definition for AWS::EC2::TransitGatewayConnectPeer

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayConnectPeerId` (AWS::EC2::TransitGatewayConnectPeer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectPeerConfiguration` | connect_peer_configuration | `map` | required, replaces on change |  | The Connect peer details. |
| `CreationTime` | creation_time | `string` | computed |  | The creation time. |
| `State` |  | `string` | computed |  | The state of the Connect peer. |
| `Tags` |  | `map` | tags map |  | The tags for the Connect Peer. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | required, replaces on change | aws.transitgatewayattachment.Id | The ID of the Connect attachment. |
| `TransitGatewayConnectPeerId` | transit_gateway_connect_peer_id | `string` | computed |  | The ID of the Connect peer. |

Supports update: yes

Discovery: supported
