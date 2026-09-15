# aws.insight

**CloudFormation type:** `AWS::SecurityHub::Insight`

The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.

Region attribute: `region`

**Import ID:** `<region>/InsightArn` (AWS::SecurityHub::Insight)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Filters` |  | `map` | required |  | A collection of filters that are applied to all active findings aggregated by AWS Security Hub. |
| `GroupByAttribute` | group_by_attribute | `string` | required |  | Non-empty string definition. |
| `InsightArn` | insight_arn | `string` | computed |  | The ARN of a Security Hub insight |
| `Name` |  | `string` | required |  | The name of a Security Hub insight |

Supports update: yes

Discovery: supported
