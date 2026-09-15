# aws.insightrule

**CloudFormation type:** `AWS::CloudWatch::InsightRule`

Resource Type definition for AWS::CloudWatch::InsightRule. Creates a Contributor Insights rule that analyzes log data to identify top contributors and usage patterns.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudWatch::InsightRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplyOnTransformedLogs` | apply_on_transformed_logs | `boolean` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `RuleBody` | rule_body | `string` | required |  |  |
| `RuleName` | rule_name | `string` | required, replaces on change |  |  |
| `RuleState` | rule_state | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
