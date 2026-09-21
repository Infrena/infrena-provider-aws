# aws.transitgatewayregistration

**CloudFormation type:** `AWS::NetworkManager::TransitGatewayRegistration`

The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|TransitGatewayArn` (AWS::NetworkManager::TransitGatewayRegistration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `TransitGatewayArn` | transit_gateway_arn | `string` | required, replaces on change | aws.transitgateway.TransitGatewayArn | The Amazon Resource Name (ARN) of the transit gateway. |

Supports update: no

Discovery: supported (parent resource required)
