# aws.securityhub.automationrule

**CloudFormation type:** `AWS::SecurityHub::AutomationRule`

The ``AWS::SecurityHub::AutomationRule`` resource specifies an automation rule based on input parameters. For more information, see [Automation rules](https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html) in the *User Guide*.

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::SecurityHub::AutomationRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | One or more actions to update finding fields if a finding matches the conditions specified in ``Criteria``. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time, in UTC and ISO 8601 format. |
| `CreatedBy` | created_by | `string` | computed |  |  |
| `Criteria` |  | `map` | required |  | The criteria that determine which findings a rule applies to. |
| `Description` |  | `string` | required |  | A description of the rule. |
| `IsTerminal` | is_terminal | `boolean` | optional, computed, provider-chosen |  | Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria. This is useful when a finding matches the criteria for multiple rules, and each rule has different actions. If a rule is terminal, Security Hub CSPM applies the rule action to a finding that matches the rule criteria and doesn't evaluate other rules for the finding. By default, a rule isn't terminal. |
| `RuleArn` | rule_arn | `string` | computed |  |  |
| `RuleName` | rule_name | `string` | required |  | The name of the rule. |
| `RuleOrder` | rule_order | `integer` | required |  | An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings. Security Hub CSPM applies rules with lower values for this parameter first. |
| `RuleStatus` | rule_status | `string` | optional, computed, provider-chosen |  | Whether the rule is active after it is created. If this parameter is equal to ``ENABLED``, ASH applies the rule to findings and finding updates after the rule is created. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time, in UTC and ISO 8601 format. |

Supports update: yes

Discovery: supported
