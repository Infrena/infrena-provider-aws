# aws.automationrulev2

**CloudFormation type:** `AWS::SecurityHub::AutomationRuleV2`

Resource schema for AWS::SecurityHub::AutomationRuleV2

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::SecurityHub::AutomationRuleV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | A list of actions to be performed when the rule criteria is met |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `Criteria` |  | `map` | required |  | Defines the parameters and conditions used to evaluate and filter security findings |
| `Description` |  | `string` | required |  | A description of the automation rule |
| `RuleArn` | rule_arn | `string` | computed |  | The ARN of the automation rule |
| `RuleId` | rule_id | `string` | computed |  | The ID of the automation rule |
| `RuleName` | rule_name | `string` | required |  | The name of the automation rule |
| `RuleOrder` | rule_order | `float` | required |  | The value for the rule priority |
| `RuleStatus` | rule_status | `string` | optional, computed, provider-chosen |  | The status of the automation rule |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp formatted in ISO8601 |

Supports update: yes

Discovery: supported
