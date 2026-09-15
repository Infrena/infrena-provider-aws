# aws.customergatewayassociation

**CloudFormation type:** `AWS::NetworkManager::CustomerGatewayAssociation`

The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|CustomerGatewayArn` (AWS::NetworkManager::CustomerGatewayAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CustomerGatewayArn` | customer_gateway_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the customer gateway. |
| `DeviceId` | device_id | `string` | required, replaces on change | aws.networkmanager.device.DeviceId | The ID of the device |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `LinkId` | link_id | `string` | optional, computed, provider-chosen, replaces on change | aws.networkmanager.link.LinkId | The ID of the link |

Supports update: no

Discovery: supported (parent resource required)
