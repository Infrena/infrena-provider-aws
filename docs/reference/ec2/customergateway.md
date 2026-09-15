# aws.customergateway

**CloudFormation type:** `AWS::EC2::CustomerGateway`

Specifies a customer gateway.

Region attribute: `region`

**Import ID:** `<region>/CustomerGatewayId` (AWS::EC2::CustomerGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BgpAsn` | bgp_asn | `integer` | optional, computed, provider-chosen, replaces on change |  | For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``. |
| `BgpAsnExtended` | bgp_asn_extended | `float` | optional, computed, provider-chosen, replaces on change |  | For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``. |
| `CertificateArn` | certificate_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) for the customer gateway certificate. |
| `CustomerGatewayId` | customer_gateway_id | `string` | computed |  |  |
| `DeviceName` | device_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of customer gateway device. |
| `IpAddress` | ip_address | `string` | required, replaces on change |  | The IP address for the customer gateway device's outside interface. The address must be static. If ``OutsideIpAddressType`` in your VPN connection options is set to ``PrivateIpv4``, you can use an RFC6598 or RFC1918 private IPv4 address. If ``OutsideIpAddressType`` is set to ``Ipv6``, you can use an IPv6 address. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags for the customer gateway. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of VPN connection that this customer gateway supports (``ipsec.1``). |

Supports update: yes

Discovery: supported
