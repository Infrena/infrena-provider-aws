# aws.verifiedpermissions.identitysource

**CloudFormation type:** `AWS::VerifiedPermissions::IdentitySource`

Definition of AWS::VerifiedPermissions::IdentitySource Resource Type

Region attribute: `region`

**Import ID:** `<region>/IdentitySourceId|PolicyStoreId` (AWS::VerifiedPermissions::IdentitySource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Configuration` |  | `string` | required |  |  |
| `Details` |  | `map` | computed |  |  |
| `IdentitySourceId` | identity_source_id | `string` | computed |  |  |
| `PolicyStoreId` | policy_store_id | `string` | required, replaces on change | aws.policystore.PolicyStoreId |  |
| `PrincipalEntityType` | principal_entity_type | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
