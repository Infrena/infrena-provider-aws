# aws.sagemaker.pipeline

**CloudFormation type:** `AWS::SageMaker::Pipeline`

Resource Type definition for AWS::SageMaker::Pipeline

Region attribute: `region`

**Import ID:** `<region>/PipelineName` (AWS::SageMaker::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ParallelismConfiguration` | parallelism_configuration | `map` | optional, computed, provider-chosen |  |  |
| `PipelineDefinition` | pipeline_definition | `map` | required |  |  |
| `PipelineDescription` | pipeline_description | `string` | optional, computed, provider-chosen |  | The description of the Pipeline. |
| `PipelineDisplayName` | pipeline_display_name | `string` | optional, computed, provider-chosen |  | The display name of the Pipeline. |
| `PipelineName` | pipeline_name | `string` | required, replaces on change |  | The name of the Pipeline. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role Arn |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
