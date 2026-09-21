# aws.policystore

**CloudFormation type:** `AWS::VerifiedPermissions::PolicyStore`

Represents a policy store that you can place schema, policies, and policy templates in to validate authorization requests

Region attribute: `region`

**Import ID:** `<region>/PolicyStoreId` (AWS::VerifiedPermissions::PolicyStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DeletionProtection` | deletion_protection | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EncryptionSettings` | encryption_settings | `map` | optional, computed, provider-chosen, write-only |  |  |
| `EncryptionState` | encryption_state | `map` | computed |  |  |
| `PolicyStoreId` | policy_store_id | `string` | computed |  |  |
| `Schema` |  | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | The tags to add to the policy store |
| `ValidationSettings` | validation_settings | `map` | required |  |  |

Supports update: yes

Discovery: supported
