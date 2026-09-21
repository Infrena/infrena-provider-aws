# aws.modelbiasjobdefinition

**CloudFormation type:** `AWS::SageMaker::ModelBiasJobDefinition`

Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition

Region attribute: `region`

**Import ID:** `<region>/JobDefinitionArn` (AWS::SageMaker::ModelBiasJobDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time at which the job definition was created. |
| `EndpointName` | endpoint_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the endpoint used to run the monitoring job. |
| `JobDefinitionArn` | job_definition_arn | `string` | computed |  | The Amazon Resource Name (ARN) of job definition. |
| `JobDefinitionName` | job_definition_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the job definition. |
| `JobResources` | job_resources | `map` | required, replaces on change |  | Identifies the resources to deploy for a monitoring job. |
| `ModelBiasAppSpecification` | model_bias_app_specification | `map` | required, replaces on change |  | Container image configuration object for the monitoring job. |
| `ModelBiasBaselineConfig` | model_bias_baseline_config | `map` | optional, computed, provider-chosen, replaces on change |  | Baseline configuration used to validate that the data conforms to the specified constraints and statistics. |
| `ModelBiasJobInput` | model_bias_job_input | `map` | required, replaces on change |  | The inputs for a monitoring job. |
| `ModelBiasJobOutputConfig` | model_bias_job_output_config | `map` | required, replaces on change |  | The output configuration for monitoring jobs. |
| `NetworkConfig` | network_config | `map` | optional, computed, provider-chosen, replaces on change |  | Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf. |
| `StoppingCondition` | stopping_condition | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies a time limit for how long the monitoring job is allowed to run. |
| `Tags` |  | `map` | replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: no

Discovery: supported
