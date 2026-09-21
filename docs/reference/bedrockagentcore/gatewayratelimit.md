# aws.gatewayratelimit

**CloudFormation type:** `AWS::BedrockAgentCore::GatewayRateLimit`

Definition of AWS::BedrockAgentCore::GatewayRateLimit Resource Type

Region attribute: `region`

**Import ID:** `<region>/GatewayIdentifier|RateLimitId` (AWS::BedrockAgentCore::GatewayRateLimit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Optional human-readable description for this limit. |
| `DimensionKeys` | dimension_keys | `list` | required, replaces on change |  | Ordered list of dimension names defining the scope of this limit. |
| `Entries` |  | `list` | required |  | Rule entries mapping dimension values to rate configurations. |
| `GatewayIdentifier` | gateway_identifier | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RateLimitId` | rate_limit_id | `string` | optional, computed, provider-chosen, replaces on change |  | Limit identifier. Optional on Create (system-generates if not provided by customer). |
| `Status` |  | `string` | computed |  | Status of a gateway limit |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
