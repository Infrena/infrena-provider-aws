# aws.system

**CloudFormation type:** `AWS::ResilienceHubV2::System`

Creates a system that represents a logical grouping of services.

Region attribute: `region`

**Import ID:** `<region>/SystemArn` (AWS::ResilienceHubV2::System)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the system was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the system. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The KMS key ID for encrypting system data. |
| `Name` |  | `string` | required, replaces on change |  | The name of the system. |
| `SharingEnabled` | sharing_enabled | `boolean` | optional, computed, provider-chosen |  | Whether the system is enabled to be shared with other members of the Organization. Only applicable if the system owner is a management account or delegated admin. |
| `SystemArn` | system_arn | `string` | computed |  | The ARN of the system. |
| `SystemId` | system_id | `string` | computed |  | The system ID. |
| `Tags` |  | `map` | tags map |  | Tags assigned to the system. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the system was last updated. |

Supports update: yes

Discovery: supported
