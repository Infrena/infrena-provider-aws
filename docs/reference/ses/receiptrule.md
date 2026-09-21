# aws.receiptrule

**CloudFormation type:** `AWS::SES::ReceiptRule`

Resource Type definition for AWS::SES::ReceiptRule

Region attribute: `region`

**Import ID:** `<region>/RuleName|RuleSetName` (AWS::SES::ReceiptRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `After` |  | `string` | optional, computed, provider-chosen, write-only |  | The name of an existing rule after which the new rule is placed. If this parameter is null, the new rule is inserted at the beginning of the rule list. |
| `Rule` |  | `map` | required |  | A data structure that contains the specified rule's name, actions, recipients, domains, enabled status, scan status, and TLS policy. |
| `RuleName` | rule_name | `string` | computed |  | The name of the rule |
| `RuleSetName` | rule_set_name | `string` | required, replaces on change |  | The name of the rule set where the receipt rule is added. |

Supports update: yes

Discovery: supported (parent resource required)
