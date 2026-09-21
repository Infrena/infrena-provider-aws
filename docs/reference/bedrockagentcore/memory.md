# aws.memory

**CloudFormation type:** `AWS::BedrockAgentCore::Memory`

Resource Type definition for AWS::BedrockAgentCore::Memory

Region attribute: `region`

**Import ID:** `<region>/MemoryArn` (AWS::BedrockAgentCore::Memory)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Memory resource |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | ARN format |
| `EventExpiryDuration` | event_expiry_duration | `integer` | required |  | Duration in days until memory events expire |
| `FailureReason` | failure_reason | `string` | computed |  |  |
| `IndexedKeys` | indexed_keys | `list` | optional, computed, provider-chosen |  | List of indexed keys for the memory |
| `MemoryArn` | memory_arn | `string` | computed |  | ARN of the Memory resource |
| `MemoryExecutionRoleArn` | memory_execution_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | ARN format |
| `MemoryId` | memory_id | `string` | computed |  | Unique identifier for the Memory resource |
| `MemoryStrategies` | memory_strategies | `list` | optional, computed, provider-chosen |  | List of memory strategies attached to this memory |
| `Name` |  | `string` | required, replaces on change |  | Name of the Memory resource |
| `Status` |  | `string` | computed |  | Status of the Memory resource |
| `StreamDeliveryResources` | stream_delivery_resources | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
