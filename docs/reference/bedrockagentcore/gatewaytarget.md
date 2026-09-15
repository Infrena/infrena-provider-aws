# aws.gatewaytarget

**CloudFormation type:** `AWS::BedrockAgentCore::GatewayTarget`

Definition of AWS::BedrockAgentCore::GatewayTarget Resource Type

Region attribute: `region`

**Import ID:** `<region>/GatewayIdentifier|TargetId` (AWS::BedrockAgentCore::GatewayTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthorizationData` | authorization_data | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CredentialProviderConfigurations` | credential_provider_configurations | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `GatewayArn` | gateway_arn | `string` | computed |  |  |
| `GatewayIdentifier` | gateway_identifier | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LastSynchronizedAt` | last_synchronized_at | `string` | computed |  |  |
| `MetadataConfiguration` | metadata_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `PrivateEndpoint` | private_endpoint | `string` | optional, computed, provider-chosen |  |  |
| `PrivateEndpointManagedResources` | private_endpoint_managed_resources | `list` | computed |  |  |
| `ProtocolType` | protocol_type | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusReasons` | status_reasons | `list` | computed |  |  |
| `TargetConfiguration` | target_configuration | `string` | required |  |  |
| `TargetId` | target_id | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
