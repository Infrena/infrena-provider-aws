# aws.vpnconnection

**CloudFormation type:** `AWS::EC2::VPNConnection`

Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.

Region attribute: `region`

**Import ID:** `<region>/VpnConnectionId` (AWS::EC2::VPNConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CustomerGatewayId` | customer_gateway_id | `string` | required | aws.customergateway.CustomerGatewayId | The ID of the customer gateway at your end of the VPN connection. |
| `EnableAcceleration` | enable_acceleration | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicate whether to enable acceleration for the VPN connection. |
| `LocalIpv4NetworkCidr` | local_ipv4_network_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection. |
| `LocalIpv6NetworkCidr` | local_ipv6_network_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection. |
| `OutsideIpAddressType` | outside_ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of IP address assigned to the outside interface of the customer gateway device. |
| `PreSharedKeyStorage` | pre_shared_key_storage | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the storage mode for the pre-shared key (PSK). Valid values are ``Standard`` (stored in the S2Slong service) or ``SecretsManager`` (stored in AWS Secrets Manager). |
| `RemoteIpv4NetworkCidr` | remote_ipv4_network_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 CIDR on the AWS side of the VPN connection. |
| `RemoteIpv6NetworkCidr` | remote_ipv6_network_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv6 CIDR on the AWS side of the VPN connection. |
| `StaticRoutesOnly` | static_routes_only | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the VPN connection. |
| `TransitGatewayId` | transit_gateway_id | `string` | optional, computed, provider-chosen | aws.transitgateway.Id | The ID of the transit gateway associated with the VPN connection. |
| `TransportTransitGatewayAttachmentId` | transport_transit_gateway_attachment_id | `string` | optional, computed, provider-chosen, replaces on change | aws.transitgatewayattachment.Id | The transit gateway attachment ID to use for the VPN tunnel. |
| `TunnelBandwidth` | tunnel_bandwidth | `string` | optional, computed, provider-chosen |  | The desired bandwidth specification for the VPN tunnel, used when creating or modifying VPN connection options to set the tunnel's throughput capacity. ``standard`` supports up to 1.25 Gbps per tunnel, while ``large`` supports up to 5 Gbps per tunnel. The default value is ``standard``. Existing VPN connections without a bandwidth setting will automatically default to ``standard``. |
| `TunnelInsideIpVersion` | tunnel_inside_ip_version | `string` | optional, computed, provider-chosen, replaces on change |  | Indicate whether the VPN tunnels process IPv4 or IPv6 traffic. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of VPN connection. |
| `VpnConcentratorId` | vpn_concentrator_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpnconcentrator.VpnConcentratorId | The ID of the VPN concentrator to associate with the VPN connection. |
| `VpnConnectionId` | vpn_connection_id | `string` | computed |  |  |
| `VpnGatewayId` | vpn_gateway_id | `string` | optional, computed, provider-chosen | aws.vpngateway.VPNGatewayId | The ID of the virtual private gateway at the AWS side of the VPN connection. |
| `VpnTunnelOptionsSpecifications` | vpn_tunnel_options_specifications | `list` | optional, computed, provider-chosen |  | The tunnel options for the VPN connection. |

Supports update: yes

Discovery: supported
