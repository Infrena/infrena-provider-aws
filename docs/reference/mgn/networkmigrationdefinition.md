# aws.networkmigrationdefinition

**CloudFormation type:** `AWS::MGN::NetworkMigrationDefinition`

Resource schema for AWS::MGN::NetworkMigrationDefinition

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MGN::NetworkMigrationDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the network migration definition. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the network migration definition was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the network migration definition. |
| `Name` |  | `string` | required |  | The name of the network migration definition. |
| `NetworkMigrationDefinitionID` | network_migration_definition_id | `string` | computed |  | The unique identifier of the network migration definition. |
| `ScopeTags` | scope_tags | `map` | optional, computed, provider-chosen |  | Scope tags map for the network migration definition. |
| `SourceConfigurations` | source_configurations | `list` | required |  | A list of source configurations for the network migration. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the network migration definition. |
| `TargetDeployment` | target_deployment | `string` | optional, computed, provider-chosen |  | The target deployment configuration for the migrated network. |
| `TargetNetwork` | target_network | `map` | required |  | Configuration for the target network topology and addressing. |
| `TargetS3Configuration` | target_s3_configuration | `map` | required |  | S3 configuration for storing target network artifacts. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the network migration definition was last updated. |

Supports update: yes

Discovery: supported
