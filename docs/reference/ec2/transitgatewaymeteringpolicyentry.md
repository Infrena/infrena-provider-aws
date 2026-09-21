# aws.transitgatewaymeteringpolicyentry

**CloudFormation type:** `AWS::EC2::TransitGatewayMeteringPolicyEntry`

AWS::EC2::TransitGatewayMeteringPolicyEntry Resource Definition

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayMeteringPolicyId|PolicyRuleNumber` (AWS::EC2::TransitGatewayMeteringPolicyEntry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DestinationCidrBlock` | destination_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The list of IP addresses of the instances receiving traffic from the transit gateway |
| `DestinationPortRange` | destination_port_range | `string` | optional, computed, provider-chosen, replaces on change |  | The list of ports on destination instances receiving traffic from the transit gateway |
| `DestinationTransitGatewayAttachmentId` | destination_transit_gateway_attachment_id | `string` | optional, computed, provider-chosen, replaces on change | aws.transitgatewayattachment.Id | The ID of the source attachment through which traffic leaves a transit gateway |
| `DestinationTransitGatewayAttachmentType` | destination_transit_gateway_attachment_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of the attachment through which traffic leaves a transit gateway |
| `MeteredAccount` | metered_account | `string` | required, replaces on change |  | The resource owner information responsible for paying default billable charges for the traffic flow |
| `PolicyRuleNumber` | policy_rule_number | `integer` | required, replaces on change |  | The rule number of the metering policy entry |
| `Protocol` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The protocol of the traffic |
| `SourceCidrBlock` | source_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The list of IP addresses of the instances sending traffic to the transit gateway for which the metering policy entry is applicable |
| `SourcePortRange` | source_port_range | `string` | optional, computed, provider-chosen, replaces on change |  | The list of ports on source instances sending traffic to the transit gateway |
| `SourceTransitGatewayAttachmentId` | source_transit_gateway_attachment_id | `string` | optional, computed, provider-chosen, replaces on change | aws.transitgatewayattachment.Id | The ID of the source attachment through which traffic enters a transit gateway |
| `SourceTransitGatewayAttachmentType` | source_transit_gateway_attachment_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of the attachment through which traffic enters a  transit gateway |
| `State` |  | `string` | computed |  | State of the transit gateway metering policy |
| `TransitGatewayMeteringPolicyId` | transit_gateway_metering_policy_id | `string` | required, replaces on change | aws.transitgatewaymeteringpolicy.TransitGatewayMeteringPolicyId | The ID of the transit gateway metering policy for which the entry is being created |
| `UpdateEffectiveAt` | update_effective_at | `string` | computed |  | The timestamp at which the latest action performed on the metering policy entry will become effective |

Supports update: no

Discovery: supported (parent resource required)
