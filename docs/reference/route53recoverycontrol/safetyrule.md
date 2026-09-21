# aws.safetyrule

**CloudFormation type:** `AWS::Route53RecoveryControl::SafetyRule`

Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/SafetyRuleArn` (AWS::Route53RecoveryControl::SafetyRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssertionRule` | assertion_rule | `map` | optional, computed, provider-chosen |  | An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted. |
| `ControlPanelArn` | control_panel_arn | `string` | optional, computed, provider-chosen | aws.controlpanel.ControlPanelArn | The Amazon Resource Name (ARN) of the control panel. |
| `GatingRule` | gating_rule | `map` | optional, computed, provider-chosen |  | A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name for the safety rule. |
| `RuleConfig` | rule_config | `map` | optional, computed, provider-chosen |  | The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes. |
| `SafetyRuleArn` | safety_rule_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the safety rule. |
| `Status` |  | `string` | computed |  | The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION. |
| `Tags` |  | `map` | tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported (parent resource required)
