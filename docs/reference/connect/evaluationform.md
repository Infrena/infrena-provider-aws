# aws.evaluationform

**CloudFormation type:** `AWS::Connect::EvaluationForm`

Creates an evaluation form for the specified CON instance.

Region attribute: `region`

**Import ID:** `<region>/EvaluationFormArn` (AWS::Connect::EvaluationForm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoEvaluationConfiguration` | auto_evaluation_configuration | `map` | optional, computed, provider-chosen |  | Configuration information about automated evaluations. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the evaluation form. |
| `EvaluationFormArn` | evaluation_form_arn | `string` | computed |  |  |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Items` |  | `list` | required |  | Items that are part of the evaluation form. The total number of sections and questions must not exceed 100 each. Questions must be contained in a section. |
| `LanguageConfiguration` | language_configuration | `map` | optional, computed, provider-chosen |  | Language configuration for an evaluation form. |
| `ReviewConfiguration` | review_configuration | `map` | optional, computed, provider-chosen |  | Configuration settings for evaluation reviews. |
| `ScoringStrategy` | scoring_strategy | `map` | optional, computed, provider-chosen |  | A scoring strategy of the evaluation form. |
| `Status` |  | `string` | required |  | The status of the evaluation form. |
| `Tags` |  | `map` | tags map |  | The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }. |
| `TargetConfiguration` | target_configuration | `map` | optional, computed, provider-chosen |  | Configuration that specifies the target for an evaluation form. |
| `Title` |  | `string` | required |  | A title of the evaluation form. |

Supports update: yes

Discovery: supported
