# aws.topicrule

**CloudFormation type:** `AWS::IoT::TopicRule`

Resource Type definition for AWS::IoT::TopicRule

Region attribute: `region`

**Import ID:** `<region>/RuleName` (AWS::IoT::TopicRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `RuleName` | rule_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TopicRulePayload` | topic_rule_payload | `map` | required |  |  |

Supports update: yes

Discovery: supported
