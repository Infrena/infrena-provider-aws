# aws.verifiedpermissions.policy

**CloudFormation type:** `AWS::VerifiedPermissions::Policy`

Definition of AWS::VerifiedPermissions::Policy Resource Type

Region attribute: `region`

**Import ID:** `<region>/PolicyId|PolicyStoreId` (AWS::VerifiedPermissions::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Definition` |  | `string` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `PolicyId` | policy_id | `string` | computed |  |  |
| `PolicyStoreId` | policy_store_id | `string` | required, replaces on change | aws.policystore.PolicyStoreId |  |
| `PolicyType` | policy_type | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
