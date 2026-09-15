# aws.rolepolicy

**CloudFormation type:** `AWS::IAM::RolePolicy`

Adds or updates an inline policy document that is embedded in the specified IAM role.

Global type (no region attribute)

**Import ID:** `global/PolicyName|RoleName` (AWS::IAM::RolePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | optional, computed, provider-chosen |  | The policy document. |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the policy document. |
| `RoleName` | role_name | `string` | required, replaces on change |  | The name of the role to associate the policy with. |

Supports update: yes

Discovery: not supported
