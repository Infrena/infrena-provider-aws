# aws.userpolicy

**CloudFormation type:** `AWS::IAM::UserPolicy`

Adds or updates an inline policy document that is embedded in the specified IAM user.

Global type (no region attribute)

**Import ID:** `global/PolicyName|UserName` (AWS::IAM::UserPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | optional, computed, provider-chosen |  | The policy document. |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the policy document. |
| `UserName` | user_name | `string` | required, replaces on change |  | The name of the user to associate the policy with. |

Supports update: yes

Discovery: not supported
