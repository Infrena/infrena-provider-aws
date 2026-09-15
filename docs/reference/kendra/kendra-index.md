# aws.kendra.index

**CloudFormation type:** `AWS::Kendra::Index`

A Kendra index

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Kendra::Index)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CapacityUnits` | capacity_units | `map` | optional, computed, provider-chosen |  | Capacity units |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the index |
| `DocumentMetadataConfigurations` | document_metadata_configurations | `list` | optional, computed, provider-chosen |  | Document metadata configurations |
| `Edition` |  | `string` | required, replaces on change |  | Edition of index |
| `Id` |  | `string` | computed |  | Unique ID of index |
| `Name` |  | `string` | required |  | Name of index |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role Arn |
| `ServerSideEncryptionConfiguration` | server_side_encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Server side encryption configuration |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of tags |
| `UserContextPolicy` | user_context_policy | `string` | optional, computed, provider-chosen |  |  |
| `UserTokenConfigurations` | user_token_configurations | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
