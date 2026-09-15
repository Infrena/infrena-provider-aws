# aws.transitgatewaymeteringpolicy

**CloudFormation type:** `AWS::EC2::TransitGatewayMeteringPolicy`

AWS::EC2::TransitGatewayMeteringPolicy Resource Definition

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayMeteringPolicyId` (AWS::EC2::TransitGatewayMeteringPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `MiddleboxAttachmentIds` | middlebox_attachment_ids | `list` | optional, computed, provider-chosen |  | Middle box attachment Ids |
| `State` |  | `string` | computed |  | State of the transit gateway metering policy |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The Id of transit gateway |
| `TransitGatewayMeteringPolicyId` | transit_gateway_metering_policy_id | `string` | computed |  | The Id of the transit gateway metering policy |
| `UpdateEffectiveAt` | update_effective_at | `string` | computed |  | The timestamp at which the latest action performed on the metering policy will become effective |

Supports update: yes

Discovery: supported
