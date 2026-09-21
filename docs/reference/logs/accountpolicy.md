# aws.accountpolicy

**CloudFormation type:** `AWS::Logs::AccountPolicy`

The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.

Region attribute: `region`

**Import ID:** `<region>/AccountId|PolicyType|PolicyName` (AWS::Logs::AccountPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | User account id |
| `PolicyDocument` | policy_document | `string` | required |  | The body of the policy document you want to use for this topic. |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the account policy |
| `PolicyType` | policy_type | `string` | required, replaces on change |  | Type of the policy. |
| `Scope` |  | `string` | optional, computed, provider-chosen |  | Scope for policy application |
| `SelectionCriteria` | selection_criteria | `string` | optional, computed, provider-chosen |  | Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes |

Supports update: yes

Discovery: supported (parent resource required)
