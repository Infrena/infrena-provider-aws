# aws.gatewayrule

**CloudFormation type:** `AWS::BedrockAgentCore::GatewayRule`

Resource Type definition for AWS::BedrockAgentCore::GatewayRule

Region attribute: `region`

**Import ID:** `<region>/GatewayIdentifier|RuleId` (AWS::BedrockAgentCore::GatewayRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  |  |
| `Conditions` |  | `list` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `GatewayArn` | gateway_arn | `string` | computed |  |  |
| `GatewayIdentifier` | gateway_identifier | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Priority` |  | `float` | required |  |  |
| `RuleId` | rule_id | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `System` |  | `map` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
