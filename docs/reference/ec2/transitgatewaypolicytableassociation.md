# aws.transitgatewaypolicytableassociation

**CloudFormation type:** `AWS::EC2::TransitGatewayPolicyTableAssociation`

AWS::EC2::TransitGatewayPolicyTableAssociation Resource Definition

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayPolicyTableId|TransitGatewayAttachmentId` (AWS::EC2::TransitGatewayPolicyTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `State` |  | `string` | computed |  | The state of the transit gateway policy table association. |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | required, replaces on change | aws.transitgatewayattachment.Id | The ID of transit gateway attachment. |
| `TransitGatewayPolicyTableId` | transit_gateway_policy_table_id | `string` | required, replaces on change | aws.transitgatewaypolicytable.TransitGatewayPolicyTableId | The ID of transit gateway policy table. |

Supports update: no

Discovery: supported (parent resource required)
