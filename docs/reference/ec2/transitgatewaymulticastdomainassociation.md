# aws.transitgatewaymulticastdomainassociation

**CloudFormation type:** `AWS::EC2::TransitGatewayMulticastDomainAssociation`

The AWS::EC2::TransitGatewayMulticastDomainAssociation type

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayMulticastDomainId|TransitGatewayAttachmentId|SubnetId` (AWS::EC2::TransitGatewayMulticastDomainAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceId` | resource_id | `string` | computed |  | The ID of the resource. |
| `ResourceType` | resource_type | `string` | computed |  | The type of resource, for example a VPC attachment. |
| `State` |  | `string` | computed |  | The state of the subnet association. |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The IDs of the subnets to associate with the transit gateway multicast domain. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | required, replaces on change | aws.transitgatewayattachment.Id | The ID of the transit gateway attachment. |
| `TransitGatewayMulticastDomainId` | transit_gateway_multicast_domain_id | `string` | required, replaces on change | aws.transitgatewaymulticastdomain.TransitGatewayMulticastDomainId | The ID of the transit gateway multicast domain. |

Supports update: no

Discovery: supported (parent resource required)
