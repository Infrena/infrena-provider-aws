# aws.devopsagent.service

**CloudFormation type:** `AWS::DevOpsAgent::Service`

The AWS::DevOpsAgent::Service resource registers external services (like Dynatrace, MCP servers, GitLab) for integration with DevOpsAgent.

Region attribute: `region`

**Import ID:** `<region>/ServiceId` (AWS::DevOpsAgent::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessibleResources` | accessible_resources | `list` | computed |  | List of accessible resources for this service |
| `AdditionalServiceDetails` | additional_service_details | `map` | computed |  | Additional details specific to the service type returned after registration |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Service. |
| `ExchangeUrlPrivateConnectionName` | exchange_url_private_connection_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the private connection to use for OAuth token exchange requests only. Cannot be specified when PrivateConnectionName is provided. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the KMS key to use for encryption. |
| `PrivateConnectionName` | private_connection_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the private connection to use for VPC connectivity. |
| `ServiceDetails` | service_details | `map` | optional, computed, provider-chosen, write-only |  | Service-specific configuration details - MCPServerSigV4 supports in-place updates; GitLab (TokenValue), MCPServer (ApiKey or BearerToken credential rotation), MCPServerNewRelic (ApiKey rotation), and MCPServerGrafana (BearerToken rotation) support in-place credential rotation; all other service types and fields require replacement when modified |
| `ServiceId` | service_id | `string` | computed |  | The unique identifier of the service |
| `ServiceType` | service_type | `string` | required, replaces on change |  | The type of service being registered |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetUrlPrivateConnectionName` | target_url_private_connection_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the private connection to use for API calls (target URL) only. Cannot be specified when PrivateConnectionName is provided. |

Supports update: yes

Discovery: supported
