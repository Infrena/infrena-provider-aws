# aws.agentregistry.registry

**CloudFormation type:** `AWS::AgentRegistry::Registry`

Definition of AWS::AgentRegistry::Registry Resource Type

Region attribute: `region`

**Import ID:** `<region>/RegistryArn` (AWS::AgentRegistry::Registry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApprovalConfiguration` | approval_configuration | `map` | optional, computed, provider-chosen |  | Configuration for the registry's record approval workflow. |
| `AuthorizerType` | authorizer_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of authorizer that controls how consumers access the registry's search and MCP invoke operations. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the registry was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the registry. |
| `DiscoveryConfiguration` | discovery_configuration | `map` | optional, computed, provider-chosen |  | Discovery configuration for the registry. Controls how consumers are authorized to search the registry and invoke its MCP endpoint. |
| `Name` |  | `string` | required |  | The name of the registry. |
| `RegistryArn` | registry_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the registry. |
| `RegistryId` | registry_id | `string` | computed |  | The unique identifier of the registry. |
| `Status` |  | `string` | computed |  | The status of the registry. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the registry. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the registry was last updated. |

Supports update: yes

Discovery: supported
