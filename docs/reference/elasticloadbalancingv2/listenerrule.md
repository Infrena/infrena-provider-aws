# aws.listenerrule

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::ListenerRule`

Specifies a listener rule. The listener must be associated with an Application Load Balancer. Each rule consists of a priority, one or more actions, and one or more conditions.

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::ElasticLoadBalancingV2::ListenerRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | The actions. |
| `Conditions` |  | `list` | required |  | The conditions. |
| `IsDefault` | is_default | `boolean` | computed |  |  |
| `ListenerArn` | listener_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.elasticloadbalancingv2.listener.ListenerArn | The Amazon Resource Name (ARN) of the listener. |
| `Priority` |  | `integer` | required |  | The rule priority. A listener can't have multiple rules with the same priority. |
| `RuleArn` | rule_arn | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Transforms` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
