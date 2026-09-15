# aws.transitgatewaypolicytable

**CloudFormation type:** `AWS::EC2::TransitGatewayPolicyTable`

AWS::EC2::TransitGatewayPolicyTable Resource Definition

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayPolicyTableId` (AWS::EC2::TransitGatewayPolicyTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Creation time of the transit gateway policy table |
| `State` |  | `string` | computed |  | State of the transit gateway policy table |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id | The Id of transit gateway |
| `TransitGatewayPolicyTableId` | transit_gateway_policy_table_id | `string` | computed |  | The Id of transit gateway policy table. |

Supports update: yes

Discovery: supported
