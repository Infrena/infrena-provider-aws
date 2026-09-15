# aws.vpngateway

**CloudFormation type:** `AWS::EC2::VPNGateway`

Specifies a virtual private gateway. A virtual private gateway is the endpoint on the VPC side of your VPN connection. You can create a virtual private gateway before creating the VPC itself.

Region attribute: `region`

**Import ID:** `<region>/VPNGatewayId` (AWS::EC2::VPNGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonSideAsn` | amazon_side_asn | `integer` | optional, computed, provider-chosen, replaces on change |  | The private Autonomous System Number (ASN) for the Amazon side of a BGP session. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the virtual private gateway. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of VPN connection the virtual private gateway supports. |
| `VPNGatewayId` | vpn_gateway_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
