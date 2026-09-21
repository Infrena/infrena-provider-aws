# aws.transitgatewaymulticastgroupsource

**CloudFormation type:** `AWS::EC2::TransitGatewayMulticastGroupSource`

The AWS::EC2::TransitGatewayMulticastGroupSource registers and deregisters members and sources (network interfaces) with the transit gateway multicast group

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayMulticastDomainId|GroupIpAddress|NetworkInterfaceId` (AWS::EC2::TransitGatewayMulticastGroupSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupIpAddress` | group_ip_address | `string` | required, replaces on change |  | The IP address assigned to the transit gateway multicast group. |
| `GroupMember` | group_member | `boolean` | computed |  | Indicates that the resource is a transit gateway multicast group member. |
| `GroupSource` | group_source | `boolean` | computed |  | Indicates that the resource is a transit gateway multicast group member. |
| `NetworkInterfaceId` | network_interface_id | `string` | required, replaces on change | aws.networkinterface.Id | The ID of the transit gateway attachment. |
| `ResourceId` | resource_id | `string` | computed |  | The ID of the resource. |
| `ResourceType` | resource_type | `string` | computed |  | The type of resource, for example a VPC attachment. |
| `SourceType` | source_type | `string` | computed |  | The source type. |
| `SubnetId` | subnet_id | `string` | computed |  | The ID of the subnet. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | computed |  | The ID of the transit gateway attachment. |
| `TransitGatewayMulticastDomainId` | transit_gateway_multicast_domain_id | `string` | required, replaces on change | aws.transitgatewaymulticastdomain.TransitGatewayMulticastDomainId | The ID of the transit gateway multicast domain. |

Supports update: no

Discovery: supported (parent resource required)
