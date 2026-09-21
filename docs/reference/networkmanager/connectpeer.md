# aws.connectpeer

**CloudFormation type:** `AWS::NetworkManager::ConnectPeer`

AWS::NetworkManager::ConnectPeer Resource Type Definition.

Region attribute: `region`

**Import ID:** `<region>/ConnectPeerId` (AWS::NetworkManager::ConnectPeer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BgpOptions` | bgp_options | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Bgp options |
| `Configuration` |  | `map` | computed |  | Configuration of the connect peer. |
| `ConnectAttachmentId` | connect_attachment_id | `string` | required, replaces on change | aws.connectattachment.AttachmentId | The ID of the attachment to connect. |
| `ConnectPeerId` | connect_peer_id | `string` | computed |  | The ID of the Connect peer. |
| `CoreNetworkAddress` | core_network_address | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The IP address of a core network. |
| `CoreNetworkId` | core_network_id | `string` | computed |  | The ID of the core network. |
| `CreatedAt` | created_at | `string` | computed |  | Connect peer creation time. |
| `EdgeLocation` | edge_location | `string` | computed |  | The Connect peer Regions where edges are located. |
| `InsideCidrBlocks` | inside_cidr_blocks | `list` | optional, computed, provider-chosen, replaces on change |  | The inside IP addresses used for a Connect peer configuration. |
| `LastModificationErrors` | last_modification_errors | `list` | computed |  | Errors from the last modification of the connect peer. |
| `PeerAddress` | peer_address | `string` | required, replaces on change |  | The IP address of the Connect peer. |
| `State` |  | `string` | computed |  | State of the connect peer. |
| `SubnetArn` | subnet_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The subnet ARN for the connect peer. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
