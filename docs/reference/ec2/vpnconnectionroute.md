# aws.vpnconnectionroute

**CloudFormation type:** `AWS::EC2::VPNConnectionRoute`

Specifies a static route for a VPN connection between an existing virtual private gateway and a VPN customer gateway. The static route allows traffic to be routed from the virtual private gateway to the VPN customer gateway.

Region attribute: `region`

**Import ID:** `<region>/DestinationCidrBlock|VpnConnectionId` (AWS::EC2::VPNConnectionRoute)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DestinationCidrBlock` | destination_cidr_block | `string` | required, replaces on change |  | The CIDR block associated with the local subnet of the customer network. |
| `VpnConnectionId` | vpn_connection_id | `string` | required, replaces on change | aws.vpnconnection.VpnConnectionId | The ID of the VPN connection. |

Supports update: no

Discovery: supported
