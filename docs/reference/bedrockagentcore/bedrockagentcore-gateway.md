# aws.bedrockagentcore.gateway

**CloudFormation type:** `AWS::BedrockAgentCore::Gateway`

Definition of AWS::BedrockAgentCore::Gateway Resource Type

Region attribute: `region`

**Import ID:** `<region>/GatewayIdentifier` (AWS::BedrockAgentCore::Gateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthorizerConfiguration` | authorizer_configuration | `string` | optional, computed, provider-chosen |  |  |
| `AuthorizerType` | authorizer_type | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExceptionLevel` | exception_level | `string` | optional, computed, provider-chosen |  |  |
| `GatewayArn` | gateway_arn | `string` | computed |  |  |
| `GatewayIdentifier` | gateway_identifier | `string` | computed |  |  |
| `GatewayUrl` | gateway_url | `string` | computed |  |  |
| `InterceptorConfigurations` | interceptor_configurations | `list` | optional, computed, provider-chosen |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `PolicyEngineConfiguration` | policy_engine_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ProtocolConfiguration` | protocol_configuration | `string` | optional, computed, provider-chosen |  |  |
| `ProtocolType` | protocol_type | `string` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `StatusReasons` | status_reasons | `list` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |
| `WafConfiguration` | waf_configuration | `map` | optional, computed, provider-chosen |  |  |
| `WebAclArn` | web_acl_arn | `string` | computed |  |  |
| `WorkloadIdentityDetails` | workload_identity_details | `map` | computed |  |  |

Supports update: yes

Discovery: supported
