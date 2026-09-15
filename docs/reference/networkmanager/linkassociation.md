# aws.linkassociation

**CloudFormation type:** `AWS::NetworkManager::LinkAssociation`

The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|DeviceId|LinkId` (AWS::NetworkManager::LinkAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeviceId` | device_id | `string` | required, replaces on change | aws.networkmanager.device.DeviceId | The ID of the device |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `LinkId` | link_id | `string` | required, replaces on change | aws.networkmanager.link.LinkId | The ID of the link |

Supports update: no

Discovery: supported (parent resource required)
