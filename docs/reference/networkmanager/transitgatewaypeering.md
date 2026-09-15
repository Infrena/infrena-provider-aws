# aws.transitgatewaypeering

**CloudFormation type:** `AWS::NetworkManager::TransitGatewayPeering`

AWS::NetworkManager::TransitGatewayPeering Resoruce Type.

Region attribute: `region`

**Import ID:** `<region>/PeeringId` (AWS::NetworkManager::TransitGatewayPeering)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CoreNetworkArn` | core_network_arn | `string` | computed |  | The ARN (Amazon Resource Name) of the core network that you want to peer a transit gateway to. |
| `CoreNetworkId` | core_network_id | `string` | required, replaces on change | aws.corenetwork.CoreNetworkId | The Id of the core network that you want to peer a transit gateway to. |
| `CreatedAt` | created_at | `string` | computed |  | The creation time of the transit gateway peering |
| `EdgeLocation` | edge_location | `string` | computed |  | The location of the transit gateway peering |
| `LastModificationErrors` | last_modification_errors | `list` | computed |  | Errors from the last modification of the transit gateway peering. |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | Peering owner account Id |
| `PeeringId` | peering_id | `string` | computed |  | The Id of the transit gateway peering |
| `PeeringType` | peering_type | `string` | computed |  | Peering type (TransitGatewayPeering) |
| `ResourceArn` | resource_arn | `string` | computed |  | The ARN (Amazon Resource Name) of the resource that you will peer to a core network |
| `State` |  | `string` | computed |  | The state of the transit gateway peering |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TransitGatewayArn` | transit_gateway_arn | `string` | required, replaces on change | aws.transitgateway.TransitGatewayArn | The ARN (Amazon Resource Name) of the transit gateway that you will peer to a core network |
| `TransitGatewayPeeringAttachmentId` | transit_gateway_peering_attachment_id | `string` | computed |  | The ID of the TransitGatewayPeeringAttachment |

Supports update: yes

Discovery: supported
