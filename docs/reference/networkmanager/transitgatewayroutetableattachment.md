# aws.transitgatewayroutetableattachment

**CloudFormation type:** `AWS::NetworkManager::TransitGatewayRouteTableAttachment`

AWS::NetworkManager::TransitGatewayRouteTableAttachment Resource Type definition.

Region attribute: `region`

**Import ID:** `<region>/AttachmentId` (AWS::NetworkManager::TransitGatewayRouteTableAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentId` | attachment_id | `string` | computed |  | The ID of the attachment. |
| `AttachmentPolicyRuleNumber` | attachment_policy_rule_number | `integer` | computed |  | The policy rule number associated with the attachment. |
| `AttachmentType` | attachment_type | `string` | computed |  | The type of attachment. |
| `CoreNetworkArn` | core_network_arn | `string` | computed |  | The ARN of a core network for the VPC attachment. |
| `CoreNetworkId` | core_network_id | `string` | computed |  | The ID of a core network where you're creating a site-to-site VPN attachment. |
| `CreatedAt` | created_at | `string` | computed |  | Creation time of the attachment. |
| `EdgeLocation` | edge_location | `string` | computed |  | The Region where the edge is located. |
| `LastModificationErrors` | last_modification_errors | `list` | computed |  | Errors from the last modification of the attachment. |
| `NetworkFunctionGroupName` | network_function_group_name | `string` | optional, computed, provider-chosen |  | The name of the network function group attachment. |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | Owner account of the attachment. |
| `PeeringId` | peering_id | `string` | required, replaces on change |  | The Id of peering between transit gateway and core network. |
| `ProposedNetworkFunctionGroupChange` | proposed_network_function_group_change | `map` | computed |  | The attachment to move from one network function group to another. |
| `ProposedSegmentChange` | proposed_segment_change | `map` | computed |  | The attachment to move from one segment to another. |
| `ResourceArn` | resource_arn | `string` | computed |  | The ARN of the Resource. |
| `RoutingPolicyLabel` | routing_policy_label | `string` | optional, computed, provider-chosen, write-only |  | Routing policy label |
| `SegmentName` | segment_name | `string` | computed |  | The name of the segment that attachment is in. |
| `State` |  | `string` | computed |  | The state of the attachment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TransitGatewayRouteTableArn` | transit_gateway_route_table_arn | `string` | required, replaces on change |  | The Arn of transit gateway route table. |
| `UpdatedAt` | updated_at | `string` | computed |  | Last update time of the attachment. |

Supports update: yes

Discovery: supported
