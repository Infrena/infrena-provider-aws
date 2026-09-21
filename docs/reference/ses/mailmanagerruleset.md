# aws.mailmanagerruleset

**CloudFormation type:** `AWS::SES::MailManagerRuleSet`

Definition of AWS::SES::MailManagerRuleSet Resource Type

Region attribute: `region`

**Import ID:** `<region>/RuleSetId` (AWS::SES::MailManagerRuleSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RuleSetArn` | rule_set_arn | `string` | computed |  |  |
| `RuleSetId` | rule_set_id | `string` | computed |  |  |
| `RuleSetName` | rule_set_name | `string` | optional, computed, provider-chosen |  |  |
| `Rules` |  | `list` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
