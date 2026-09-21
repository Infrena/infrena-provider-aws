# aws.trialcomponent

**CloudFormation type:** `AWS::SageMaker::TrialComponent`

Resource type definition for AWS::SageMaker::TrialComponent. A trial component is a stage of a machine learning trial, such as a preprocessing job, training job, or batch transform job.

Region attribute: `region`

**Import ID:** `<region>/TrialComponentArn` (AWS::SageMaker::TrialComponent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | When the trial component was created. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The name of the component as displayed. If DisplayName isn't specified, TrialComponentName is displayed. |
| `InputArtifacts` | input_artifacts | `map` | optional, computed, provider-chosen |  | The input artifacts for the component. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | When the trial component was last modified. |
| `LineageGroupArn` | lineage_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the lineage group. |
| `MetadataProperties` | metadata_properties | `map` | optional, computed, provider-chosen, replaces on change |  | Metadata properties of the tracking entity, trial, or trial component. |
| `OutputArtifacts` | output_artifacts | `map` | optional, computed, provider-chosen |  | The output artifacts for the component. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | The hyperparameters for the component. |
| `Status` |  | `map` | optional, computed, provider-chosen |  | The status of the trial component. |
| `Tags` |  | `map` | tags map |  | A list of tags to associate with the trial component. |
| `TrialComponentArn` | trial_component_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the trial component. |
| `TrialComponentName` | trial_component_name | `string` | required, replaces on change |  | The name of the trial component. Must be unique in your AWS account and is not case-sensitive. |

Supports update: yes

Discovery: supported
