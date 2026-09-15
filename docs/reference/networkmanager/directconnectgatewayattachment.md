# aws.directconnectgatewayattachment

**CloudFormation type:** `AWS::NetworkManager::DirectConnectGatewayAttachment`

AWS::NetworkManager::DirectConnectGatewayAttachment Resource Type

Region attribute: `region`

**Import ID:** `<region>/AttachmentId` (AWS::NetworkManager::DirectConnectGatewayAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentId` | attachment_id | `string` | computed |  | Id of the attachment. |
| `AttachmentPolicyRuleNumber` | attachment_policy_rule_number | `integer` | computed |  | The policy rule number associated with the attachment. |
| `AttachmentType` | attachment_type | `string` | computed |  | Attachment type. |
| `CoreNetworkArn` | core_network_arn | `string` | computed |  | The ARN of a core network for the Direct Connect Gateway attachment. |
| `CoreNetworkId` | core_network_id | `string` | required, replaces on change | aws.corenetwork.CoreNetworkId | The ID of a core network for the Direct Connect Gateway attachment. |
| `CreatedAt` | created_at | `string` | computed |  | Creation time of the attachment. |
| `DirectConnectGatewayArn` | direct_connect_gateway_arn | `string` | required, replaces on change |  | The ARN of the Direct Connect Gateway. |
| `EdgeLocations` | edge_locations | `list` | required |  | The Regions where the edges are located. |
| `LastModificationErrors` | last_modification_errors | `list` | computed |  | Errors from the last modification of the attachment. |
| `NetworkFunctionGroupName` | network_function_group_name | `string` | computed |  | The name of the network function group attachment. |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | Owner account of the attachment. |
| `ProposedNetworkFunctionGroupChange` | proposed_network_function_group_change | `map` | computed |  | The attachment to move from one network function group to another. |
| `ProposedSegmentChange` | proposed_segment_change | `map` | computed |  | The attachment to move from one segment to another. |
| `ResourceArn` | resource_arn | `string` | computed |  | The ARN of the Resource. |
| `RoutingPolicyLabel` | routing_policy_label | `string` | optional, computed, provider-chosen, write-only |  | Routing policy label |
| `SegmentName` | segment_name | `string` | computed |  | The name of the segment attachment.. |
| `State` |  | `string` | computed |  | State of the attachment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the attachment. |
| `UpdatedAt` | updated_at | `string` | computed |  | Last update time of the attachment. |

Supports update: yes

Discovery: supported
