# aws.registryrecord

**CloudFormation type:** `AWS::AgentRegistry::RegistryRecord`

Definition of AWS::AgentRegistry::RegistryRecord Resource Type

Region attribute: `region`

**Import ID:** `<region>/RecordArn` (AWS::AgentRegistry::RegistryRecord)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the registry record was created. |
| `CreatedBy` | created_by | `string` | computed |  | The identifier of the AWS account that created the registry record. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the registry record. |
| `Descriptors` |  | `map` | required |  | The typed set of descriptors for a registry record. Exactly one descriptor field is populated based on the record type. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The human-readable display name of the registry record. |
| `Name` |  | `string` | required |  | The name of the registry record. |
| `RecordArn` | record_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the registry record. |
| `RecordId` | record_id | `string` | computed |  | The unique identifier of the registry record. |
| `RecordType` | record_type | `string` | required, replaces on change |  | The type of the registry record. |
| `RecordVersion` | record_version | `string` | optional, computed, provider-chosen |  | The version of the registry record. |
| `RegistryArn` | registry_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the registry containing the record. |
| `RegistryId` | registry_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.agentregistry.registry.RegistryId | The identifier of the registry in which to create the record. You can specify either the registry ID or the registry Amazon Resource Name (ARN). Use the ARN form to reference a registry shared from another account via AWS Resource Access Manager (RAM). |
| `Status` |  | `string` | computed |  | The lifecycle status of the registry record. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the registry record. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the registry record was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
