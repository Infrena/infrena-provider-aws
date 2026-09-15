# aws.corenetwork

**CloudFormation type:** `AWS::NetworkManager::CoreNetwork`

AWS::NetworkManager::CoreNetwork Resource Type Definition.

Region attribute: `region`

**Import ID:** `<region>/CoreNetworkId` (AWS::NetworkManager::CoreNetwork)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CoreNetworkArn` | core_network_arn | `string` | computed |  | The ARN (Amazon resource name) of core network |
| `CoreNetworkId` | core_network_id | `string` | computed |  | The Id of core network |
| `CreatedAt` | created_at | `string` | computed |  | The creation time of core network |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of core network |
| `Edges` |  | `list` | computed |  | The edges within a core network. |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network that your core network is a part of. |
| `NetworkFunctionGroups` | network_function_groups | `list` | computed |  | The network function groups within a core network. |
| `OwnerAccount` | owner_account | `string` | computed |  | Owner of the core network |
| `PolicyDocument` | policy_document | `map` | optional, computed, provider-chosen |  | Live policy document for the core network, you must provide PolicyDocument in Json Format |
| `Segments` |  | `list` | computed |  | The segments within a core network. |
| `State` |  | `string` | computed |  | The state of core network |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the global network. |

Supports update: yes

Discovery: supported
