# aws.solution

**CloudFormation type:** `AWS::Personalize::Solution`

Resource schema for AWS::Personalize::Solution.

Region attribute: `region`

**Import ID:** `<region>/SolutionArn` (AWS::Personalize::Solution)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetGroupArn` | dataset_group_arn | `string` | required, replaces on change | aws.personalize.datasetgroup.DatasetGroupArn | The ARN of the dataset group that provides the training data. |
| `EventType` | event_type | `string` | optional, computed, provider-chosen, replaces on change |  | When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type. |
| `Name` |  | `string` | required, replaces on change |  | The name for the solution |
| `PerformAutoML` | perform_auto_ml | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn. |
| `PerformHPO` | perform_hpo | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false. |
| `RecipeArn` | recipe_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the recipe to use for model training. Only specified when performAutoML is false. |
| `SolutionArn` | solution_arn | `string` | computed |  | The ARN of the solution |
| `SolutionConfig` | solution_config | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration. |
| `Tags` |  | `map` | replaces on change, tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: no

Discovery: supported
