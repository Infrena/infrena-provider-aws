# aws.policytemplate

**CloudFormation type:** `AWS::VerifiedPermissions::PolicyTemplate`

Definition of AWS::VerifiedPermissions::PolicyTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/PolicyStoreId|PolicyTemplateId` (AWS::VerifiedPermissions::PolicyTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `PolicyStoreId` | policy_store_id | `string` | required, replaces on change | aws.policystore.PolicyStoreId |  |
| `PolicyTemplateId` | policy_template_id | `string` | computed |  |  |
| `Statement` |  | `string` | required |  |  |

Supports update: yes

Discovery: supported (parent resource required)
