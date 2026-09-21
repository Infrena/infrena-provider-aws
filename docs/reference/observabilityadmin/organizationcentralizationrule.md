# aws.organizationcentralizationrule

**CloudFormation type:** `AWS::ObservabilityAdmin::OrganizationCentralizationRule`

Resource schema for AWS:ObservabilityAdmin:OrganizationCentralizationRule

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::ObservabilityAdmin::OrganizationCentralizationRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Rule` |  | `map` | required |  |  |
| `RuleArn` | rule_arn | `string` | computed |  |  |
| `RuleName` | rule_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
