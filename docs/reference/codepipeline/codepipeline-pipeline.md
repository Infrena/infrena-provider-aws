# aws.codepipeline.pipeline

**CloudFormation type:** `AWS::CodePipeline::Pipeline`

The AWS::CodePipeline::Pipeline resource creates a CodePipeline pipeline that describes how software changes go through a release process.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::CodePipeline::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the pipeline. |
| `ArtifactStore` | artifact_store | `map` | optional, computed, provider-chosen |  | The S3 bucket where artifacts for the pipeline are stored. |
| `ArtifactStores` | artifact_stores | `list` | optional, computed, provider-chosen |  | A mapping of artifactStore objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline. |
| `DisableInboundStageTransitions` | disable_inbound_stage_transitions | `list` | optional, computed, provider-chosen |  | Represents the input of a DisableStageTransition action. |
| `ExecutionMode` | execution_mode | `string` | optional, computed, provider-chosen |  | The method that the pipeline will use to handle multiple executions. The default mode is SUPERSEDED. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the pipeline. |
| `PipelineType` | pipeline_type | `string` | optional, computed, provider-chosen |  | CodePipeline provides the following pipeline types, which differ in characteristics and price, so that you can tailor your pipeline features and cost to the needs of your applications. |
| `RestartExecutionOnUpdate` | restart_execution_on_update | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates whether to rerun the CodePipeline pipeline after you update it. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn |
| `Stages` |  | `list` | required |  | Represents information about a stage and its definition. |
| `Tags` |  | `map` | tags map |  | Specifies the tags applied to the pipeline. |
| `Triggers` |  | `list` | optional, computed, provider-chosen |  | The trigger configuration specifying a type of event, such as Git tags, that starts the pipeline. |
| `Variables` |  | `list` | optional, computed, provider-chosen |  | A list that defines the pipeline variables for a pipeline resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9@\-_]+. |
| `Version` |  | `string` | computed |  | The version of the pipeline. |

Supports update: yes

Discovery: supported
