# aws.directconnectgateway

**CloudFormation type:** `AWS::DirectConnect::DirectConnectGateway`

Resource Type definition for AWS::DirectConnect::DirectConnectGateway

Region attribute: `region`

**Import ID:** `<region>/DirectConnectGatewayArn` (AWS::DirectConnect::DirectConnectGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonSideAsn` | amazon_side_asn | `string` | optional, computed, provider-chosen, replaces on change |  | The autonomous system number (ASN) for the Amazon side of the connection. |
| `DirectConnectGatewayArn` | direct_connect_gateway_arn | `string` | computed |  | The ARN of the Direct Connect gateway. |
| `DirectConnectGatewayId` | direct_connect_gateway_id | `string` | computed |  | The ID of the Direct Connect gateway. |
| `DirectConnectGatewayName` | direct_connect_gateway_name | `string` | required |  | The name of the Direct Connect gateway. |
| `Tags` |  | `map` | tags map |  | The tags associated with the Direct Connect gateway. |

Supports update: yes

Discovery: supported
