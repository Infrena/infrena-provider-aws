# aws.transitgatewaypolicytableentry

**CloudFormation type:** `AWS::EC2::TransitGatewayPolicyTableEntry`

AWS::EC2::TransitGatewayPolicyTableEntry Resource Definition

Region attribute: `region`

**Import ID:** `<region>/TransitGatewayPolicyTableId|PolicyRuleNumber` (AWS::EC2::TransitGatewayPolicyTableEntry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyRule` | policy_rule | `map` | required |  | The policy rule associated with the entry. |
| `PolicyRuleNumber` | policy_rule_number | `string` | required, replaces on change |  | The rule number for the policy table entry. |
| `State` |  | `string` | computed |  | The state of the policy table entry. |
| `TargetRouteTableId` | target_route_table_id | `string` | required |  | The ID of the target route table. |
| `TransitGatewayPolicyTableId` | transit_gateway_policy_table_id | `string` | required, replaces on change | aws.transitgatewaypolicytable.TransitGatewayPolicyTableId | The ID of the transit gateway policy table. |

Supports update: yes

Discovery: supported (parent resource required)
