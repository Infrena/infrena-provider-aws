# aws.events.rule

**CloudFormation type:** `AWS::Events::Rule`

Resource Type definition for AWS::Events::Rule

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Events::Rule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the rule, such as arn:aws:events:us-east-2:123456789012:rule/example. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the rule. |
| `EventBusName` | event_bus_name | `string` | optional, computed, provider-chosen |  | The name or ARN of the event bus associated with the rule. If you omit this, the default event bus is used. |
| `EventPattern` | event_pattern | `string` | optional, computed, provider-chosen |  | The event pattern of the rule. For more information, see Events and Event Patterns in the Amazon EventBridge User Guide. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the rule. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the role that is used for target invocation. |
| `RuleName` | rule_name | `string` | computed |  | The name of the rule, exposed as a read-only attribute for use with Fn::GetAtt. |
| `ScheduleExpression` | schedule_expression | `string` | optional, computed, provider-chosen |  | The scheduling expression. For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see Creating an Amazon EventBridge rule that runs on a schedule. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the rule. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the event rule. |
| `Targets` |  | `list` | optional, computed, provider-chosen |  | Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule. |

Supports update: yes

Discovery: supported
