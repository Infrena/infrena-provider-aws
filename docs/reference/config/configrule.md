# aws.configrule

**CloudFormation type:** `AWS::Config::ConfigRule`

You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).

Region attribute: `region`

**Import ID:** `<region>/ConfigRuleName` (AWS::Config::ConfigRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Compliance` |  | `map` | optional, computed, provider-chosen |  | Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance. |
| `ConfigRuleId` | config_rule_id | `string` | computed |  |  |
| `ConfigRuleName` | config_rule_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description that you provide for the CC rule. |
| `EvaluationModes` | evaluation_modes | `list` | optional, computed, provider-chosen |  | The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only. |
| `InputParameters` | input_parameters | `string` | optional, computed, provider-chosen |  | A string, in JSON format, that is passed to the CC rule Lambda function. |
| `MaximumExecutionFrequency` | maximum_execution_frequency | `string` | optional, computed, provider-chosen |  | The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when: |
| `Scope` |  | `map` | optional, computed, provider-chosen |  | Defines which resources trigger an evaluation for an CC rule. The scope can include one or more resource types, a combination of a tag key and value, or a combination of one resource type and one resource ID. Specify a scope to constrain which resources trigger an evaluation for a rule. Otherwise, evaluations for the rule are triggered when any resource in your recording group changes in configuration. |
| `Source` |  | `map` | required |  | Provides the CustomPolicyDetails, the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the events that cause the evaluation of your AWS resources. |

Supports update: yes

Discovery: supported
