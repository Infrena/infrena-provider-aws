# aws.policystatement

**CloudFormation type:** `AWS::EntityResolution::PolicyStatement`

Policy Statement defined in AWS Entity Resolution Service

Region attribute: `region`

**Import ID:** `<region>/Arn|StatementId` (AWS::EntityResolution::PolicyStatement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | required, replaces on change |  | Arn of the resource to which the policy statement is being attached. |
| `Condition` |  | `string` | optional, computed, provider-chosen |  |  |
| `Effect` |  | `string` | optional, computed, provider-chosen |  |  |
| `Principal` |  | `list` | optional, computed, provider-chosen |  |  |
| `StatementId` | statement_id | `string` | required, replaces on change |  | The Statement Id of the policy statement that is being attached. |

Supports update: yes

Discovery: supported (parent resource required)
