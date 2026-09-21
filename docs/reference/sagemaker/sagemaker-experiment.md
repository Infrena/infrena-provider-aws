# aws.sagemaker.experiment

**CloudFormation type:** `AWS::SageMaker::Experiment`

Resource type definition for AWS::SageMaker::Experiment. An experiment is a collection of related trials used to organize and track machine learning workflows.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::Experiment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the experiment. |
| `CreationTime` | creation_time | `string` | computed |  | When the experiment was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the experiment. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The name of the experiment as displayed. The name does not need to be unique. |
| `ExperimentName` | experiment_name | `string` | required, replaces on change |  | The name of the experiment. Must be unique in your AWS account and is not case-sensitive. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | When the experiment was last modified. |
| `Tags` |  | `map` | tags map |  | A list of tags to associate with the experiment. |

Supports update: yes

Discovery: supported
