# aws.iotsitewise.pipeline

**CloudFormation type:** `AWS::IoTSiteWise::Pipeline`

Resource schema for AWS::IoTSiteWise::Pipeline. A pipeline defines a directed acyclic graph (DAG) of compute nodes, where each node references a task definition and can declare dependencies on other nodes.

Region attribute: `region`

**Import ID:** `<region>/PipelineArn` (AWS::IoTSiteWise::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Computations` |  | `list` | required |  | The list of compute nodes that form the pipeline DAG. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the pipeline. |
| `EnvironmentVariables` | environment_variables | `map` | optional, computed, provider-chosen |  | A map of environment variable key-value pairs. |
| `PipelineArn` | pipeline_arn | `string` | computed |  | The ARN of the pipeline. |
| `PipelineName` | pipeline_name | `string` | required, replaces on change |  | The name of the pipeline. Must be unique within the workspace. |
| `Status` |  | `string` | computed |  | The current lifecycle status of the pipeline. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `WorkspaceName` | workspace_name | `string` | required, replaces on change |  | The name of the workspace. |

Supports update: yes

Discovery: supported (parent resource required)
