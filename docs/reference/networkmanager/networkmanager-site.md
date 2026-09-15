# aws.networkmanager.site

**CloudFormation type:** `AWS::NetworkManager::Site`

The AWS::NetworkManager::Site type describes a site.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|SiteId` (AWS::NetworkManager::Site)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the device was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the site. |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `Location` |  | `map` | optional, computed, provider-chosen |  | The location of the site |
| `SiteArn` | site_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the site. |
| `SiteId` | site_id | `string` | computed |  | The ID of the site. |
| `State` |  | `string` | computed |  | The state of the site. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the site. |

Supports update: yes

Discovery: supported (parent resource required)
