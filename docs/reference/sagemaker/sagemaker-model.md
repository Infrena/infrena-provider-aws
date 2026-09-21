# aws.sagemaker.model

**CloudFormation type:** `AWS::SageMaker::Model`

Resource type definition for AWS::SageMaker::Model

Region attribute: `region`

**Import ID:** `<region>/ModelArn` (AWS::SageMaker::Model)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Containers` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Specifies the containers in the inference pipeline. |
| `EnableNetworkIsolation` | enable_network_isolation | `boolean` | optional, computed, provider-chosen, replaces on change |  | Isolates the model container. No inbound or outbound network calls can be made to or from the model container. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role that you specified for the model. |
| `InferenceExecutionConfig` | inference_execution_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies details about how containers in a multi-container endpoint are run. |
| `ModelArn` | model_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the model. |
| `ModelName` | model_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the new model. |
| `PrimaryContainer` | primary_container | `map` | optional, computed, provider-chosen, replaces on change |  | Describes the container, as part of model definition. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs. You can use tags to categorize your AWS resources in different ways, for example, by purpose, owner, or environment. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html). |

Supports update: yes

Discovery: supported
