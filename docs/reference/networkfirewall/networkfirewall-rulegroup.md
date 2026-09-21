# aws.networkfirewall.rulegroup

**CloudFormation type:** `AWS::NetworkFirewall::RuleGroup`

Resource type definition for AWS::NetworkFirewall::RuleGroup

Region attribute: `region`

**Import ID:** `<region>/RuleGroupArn` (AWS::NetworkFirewall::RuleGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Capacity` |  | `integer` | required, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `RuleGroup` | rule_group | `map` | optional, computed, provider-chosen |  |  |
| `RuleGroupArn` | rule_group_arn | `string` | computed |  | A resource ARN. |
| `RuleGroupId` | rule_group_id | `string` | computed |  |  |
| `RuleGroupName` | rule_group_name | `string` | required, replaces on change |  |  |
| `SummaryConfiguration` | summary_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
