# aws.endpointconfig

**CloudFormation type:** `AWS::SageMaker::EndpointConfig`

Resource Type definition for AWS::SageMaker::EndpointConfig

Region attribute: `region`

**Import ID:** `<region>/EndpointConfigArn` (AWS::SageMaker::EndpointConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AsyncInferenceConfig` | async_inference_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies configuration for how an endpoint performs asynchronous inference. |
| `DataCaptureConfig` | data_capture_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies how to capture endpoint data for model monitor. The data capture configuration applies to all production variants hosted at the endpoint. |
| `EnableNetworkIsolation` | enable_network_isolation | `boolean` | optional, computed, provider-chosen, replaces on change |  | Sets whether all model containers deployed to the endpoint are isolated. If they are, no inbound or outbound network calls can be made to or from the model containers. |
| `EndpointConfigArn` | endpoint_config_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the endpoint configuration. |
| `EndpointConfigName` | endpoint_config_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the endpoint configuration. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform actions on your behalf. |
| `ExplainerConfig` | explainer_config | `map` | optional, computed, provider-chosen, replaces on change |  | A parameter to activate explainers. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint. |
| `MetricsConfig` | metrics_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the metrics that the endpoint publishes to Amazon CloudWatch, the frequency of publication, and whether to enable enhanced or detailed observability metrics. |
| `ProductionVariants` | production_variants | `list` | required, replaces on change |  | A list of ProductionVariant objects, one for each model that you want to host at this endpoint. |
| `ShadowProductionVariants` | shadow_production_variants | `list` | optional, computed, provider-chosen, replaces on change |  | Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to apply to this resource. |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. |

Supports update: yes

Discovery: supported
