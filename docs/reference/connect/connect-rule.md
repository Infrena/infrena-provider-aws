# aws.connect.rule

**CloudFormation type:** `AWS::Connect::Rule`

Creates a rule for the specified CON instance.

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::Connect::Rule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `map` | required |  | A list of actions to be run when the rule is triggered. |
| `Function` |  | `string` | required |  | The conditions of the rule. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The Amazon Resource Name (ARN) of the instance. |
| `Name` |  | `string` | required |  | The name of the rule. |
| `PublishStatus` | publish_status | `string` | required |  | The publish status of the rule. |
| `RuleArn` | rule_arn | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }. |
| `TriggerEventSource` | trigger_event_source | `map` | required, replaces on change |  | The name of the event source. |

Supports update: yes

Discovery: supported (parent resource required)
