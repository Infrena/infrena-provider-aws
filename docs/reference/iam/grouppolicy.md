# aws.grouppolicy

**CloudFormation type:** `AWS::IAM::GroupPolicy`

Adds or updates an inline policy document that is embedded in the specified IAM group.

Global type (no region attribute)

**Import ID:** `global/PolicyName|GroupName` (AWS::IAM::GroupPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupName` | group_name | `string` | required, replaces on change |  | The name of the group to associate the policy with. |
| `PolicyDocument` | policy_document | `map` | optional, computed, provider-chosen |  | The policy document. |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the policy document. |

Supports update: yes

Discovery: not supported
