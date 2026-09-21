# aws.connectattachment

**CloudFormation type:** `AWS::NetworkManager::ConnectAttachment`

AWS::NetworkManager::ConnectAttachment Resource Type Definition

Region attribute: `region`

**Import ID:** `<region>/AttachmentId` (AWS::NetworkManager::ConnectAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentId` | attachment_id | `string` | computed |  | The ID of the attachment. |
| `AttachmentPolicyRuleNumber` | attachment_policy_rule_number | `integer` | computed |  | The policy rule number associated with the attachment. |
| `AttachmentType` | attachment_type | `string` | computed |  | The type of attachment. |
| `CoreNetworkArn` | core_network_arn | `string` | computed |  | The ARN of a core network. |
| `CoreNetworkId` | core_network_id | `string` | required, replaces on change | aws.corenetwork.CoreNetworkId | ID of the CoreNetwork that the attachment will be attached to. |
| `CreatedAt` | created_at | `string` | computed |  | Creation time of the attachment. |
| `EdgeLocation` | edge_location | `string` | required, replaces on change |  | Edge location of the attachment. |
| `LastModificationErrors` | last_modification_errors | `list` | computed |  | Errors from the last modification of the attachment. |
| `NetworkFunctionGroupName` | network_function_group_name | `string` | optional, computed, provider-chosen |  | The name of the network function group attachment. |
| `Options` |  | `map` | required, replaces on change |  | Connect attachment options for protocol |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | The ID of the attachment account owner. |
| `ProposedNetworkFunctionGroupChange` | proposed_network_function_group_change | `map` | computed |  | The attachment to move from one network function group to another. |
| `ProposedSegmentChange` | proposed_segment_change | `map` | computed |  | The attachment to move from one segment to another. |
| `ResourceArn` | resource_arn | `string` | computed |  | The attachment resource ARN. |
| `RoutingPolicyLabel` | routing_policy_label | `string` | optional, computed, provider-chosen, write-only |  | Routing policy label |
| `SegmentName` | segment_name | `string` | computed |  | The name of the segment attachment. |
| `State` |  | `string` | computed |  | State of the attachment. |
| `Tags` |  | `map` | tags map |  | Tags for the attachment. |
| `TransportAttachmentId` | transport_attachment_id | `string` | required, replaces on change |  | Id of transport attachment |
| `UpdatedAt` | updated_at | `string` | computed |  | Last update time of the attachment. |

Supports update: yes

Discovery: supported
