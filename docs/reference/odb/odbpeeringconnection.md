# aws.odbpeeringconnection

**CloudFormation type:** `AWS::ODB::OdbPeeringConnection`

Resource Type definition for AWS::ODB::OdbPeeringConnection.

Region attribute: `region`

**Import ID:** `<region>/OdbPeeringConnectionArn` (AWS::ODB::OdbPeeringConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalPeerNetworkCidrs` | additional_peer_network_cidrs | `list` | optional, computed, provider-chosen, write-only |  | The additional CIDR blocks for the ODB peering connection. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The name of the ODB peering connection. |
| `OdbNetworkArn` | odb_network_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the ODB network. |
| `OdbNetworkId` | odb_network_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.odbnetwork.OdbNetworkId | The unique identifier of the ODB network. |
| `OdbPeeringConnectionArn` | odb_peering_connection_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the ODB peering connection. |
| `OdbPeeringConnectionId` | odb_peering_connection_id | `string` | computed |  | The unique identifier of the ODB peering connection. |
| `PeerNetworkArn` | peer_network_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the peer network. |
| `PeerNetworkCidrs` | peer_network_cidrs | `list` | computed |  | The CIDR blocks for the ODB peering connection. |
| `PeerNetworkId` | peer_network_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The unique identifier of the peer network. |
| `PeerNetworkRouteTableIds` | peer_network_route_table_ids | `list` | optional, computed, provider-chosen, replaces on change, write-only | aws.routetable.RouteTableId | The unique identifier of the VPC route table for which a route to the ODB network is automatically created during peering connection establishment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Odb peering connection. |

Supports update: yes

Discovery: supported
