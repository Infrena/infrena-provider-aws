# aws.transitgatewaymulticastdomain

**CloudFormation type:** `AWS::EC2::TransitGatewayMulticastDomain`

The AWS::EC2::TransitGatewayMulticastDomain type

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayMulticastDomainId` (AWS::EC2::TransitGatewayMulticastDomain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time the transit gateway multicast domain was created. |
| `Options` |  | `map` | optional, computed, provider-chosen |  | The options for the transit gateway multicast domain. |
| `State` |  | `string` | computed |  | The state of the transit gateway multicast domain. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the transit gateway multicast domain. |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The ID of the transit gateway. |
| `TransitGatewayMulticastDomainArn` | transit_gateway_multicast_domain_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the transit gateway multicast domain. |
| `TransitGatewayMulticastDomainId` | transit_gateway_multicast_domain_id | `string` | computed |  | The ID of the transit gateway multicast domain. |

Supports update: yes

Discovery: supported
