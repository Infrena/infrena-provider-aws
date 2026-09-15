# aws.networkmanager.link

**CloudFormation type:** `AWS::NetworkManager::Link`

The AWS::NetworkManager::Link type describes a link.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|LinkId` (AWS::NetworkManager::Link)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Bandwidth` |  | `map` | required |  | The bandwidth for the link. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the device was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the link. |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `LinkArn` | link_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the link. |
| `LinkId` | link_id | `string` | computed |  | The ID of the link. |
| `Provider` | provider_value | `string` | optional, computed, provider-chosen |  | The provider of the link. |
| `SiteId` | site_id | `string` | required, replaces on change | aws.networkmanager.site.SiteId | The ID of the site |
| `State` |  | `string` | computed |  | The state of the link. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the link. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | The type of the link. |

Supports update: yes

Discovery: supported (parent resource required)
