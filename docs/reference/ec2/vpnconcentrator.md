# aws.vpnconcentrator

**CloudFormation type:** `AWS::EC2::VPNConcentrator`

Describes a VPN concentrator.

Region attribute: `region`

**Import ID:** `<region>/VpnConcentratorId` (AWS::EC2::VPNConcentrator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the VPN concentrator. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | computed |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The ID of the transit gateway associated with the VPN concentrator. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of VPN concentrator. |
| `VpnConcentratorId` | vpn_concentrator_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
