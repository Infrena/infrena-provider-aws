# aws.samplingrule

**CloudFormation type:** `AWS::XRay::SamplingRule`

This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.

Region attribute: `region`

**Import ID:** `<region>/RuleARN` (AWS::XRay::SamplingRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RuleARN` | rule_arn | `string` | computed |  | The ARN of the sampling rule. Specify a rule by either name or ARN, but not both. |
| `RuleName` | rule_name | `string` | optional, computed, provider-chosen |  | The ARN of the sampling rule. Specify a rule by either name or ARN, but not both. |
| `SamplingRule` | sampling_rule | `map` | optional, computed, provider-chosen |  |  |
| `SamplingRuleRecord` | sampling_rule_record | `map` | optional, computed, provider-chosen |  |  |
| `SamplingRuleUpdate` | sampling_rule_update | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
