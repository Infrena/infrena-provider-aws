# aws.corenetworkprefixlistassociation

**CloudFormation type:** `AWS::NetworkManager::CoreNetworkPrefixListAssociation`

Resource Type definition for AWS::NetworkManager::CoreNetworkPrefixListAssociation which associates a prefix list with a core network.

Region attribute: `region`

**Import ID:** `<region>/CoreNetworkId|PrefixListArn` (AWS::NetworkManager::CoreNetworkPrefixListAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CoreNetworkId` | core_network_id | `string` | required, replaces on change | aws.corenetwork.CoreNetworkId | The ID of the core network. |
| `PrefixListAlias` | prefix_list_alias | `string` | required, replaces on change |  | The alias of the prefix list |
| `PrefixListArn` | prefix_list_arn | `string` | required, replaces on change | aws.prefixlist.Arn | The Amazon Resource Name (ARN) of the prefix list. |

Supports update: no

Discovery: supported (parent resource required)
