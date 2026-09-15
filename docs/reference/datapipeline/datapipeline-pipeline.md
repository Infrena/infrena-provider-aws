# aws.datapipeline.pipeline

**CloudFormation type:** `AWS::DataPipeline::Pipeline`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/PipelineId` (AWS::DataPipeline::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Activate` |  | `boolean` | optional, computed, provider-chosen |  | Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the pipeline. |
| `Name` |  | `string` | required, replaces on change |  | The name of the pipeline. |
| `ParameterObjects` | parameter_objects | `list` | optional, computed, provider-chosen |  | The parameter objects used with the pipeline. |
| `ParameterValues` | parameter_values | `list` | optional, computed, provider-chosen |  | The parameter values used with the pipeline. |
| `PipelineId` | pipeline_id | `string` | computed |  |  |
| `PipelineObjects` | pipeline_objects | `list` | optional, computed, provider-chosen |  | The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide. |
| `PipelineTags` | pipeline_tags | `map` | optional, computed, provider-chosen, tags map |  | A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide. |

Supports update: yes

Discovery: supported
