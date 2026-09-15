# aws.inferenceexperiment

**CloudFormation type:** `AWS::SageMaker::InferenceExperiment`

Resource Type definition for AWS::SageMaker::InferenceExperiment

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::SageMaker::InferenceExperiment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the inference experiment. |
| `CreationTime` | creation_time | `string` | computed |  | The timestamp at which you created the inference experiment. |
| `DataStorageConfig` | data_storage_config | `map` | optional, computed, provider-chosen |  | The Amazon S3 location and configuration for storing inference request and response data. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the inference experiment. |
| `DesiredState` | desired_state | `string` | optional, computed, provider-chosen, write-only |  | The desired state of the experiment after starting or stopping operation. |
| `EndpointMetadata` | endpoint_metadata | `map` | computed |  | The metadata of the endpoint on which the inference experiment ran. |
| `EndpointName` | endpoint_name | `string` | required, replaces on change |  | The name of the endpoint used to run the inference experiment. |
| `KmsKey` | kms_key | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The timestamp at which you last modified the inference experiment. |
| `ModelVariants` | model_variants | `list` | required |  | An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant. |
| `Name` |  | `string` | required, replaces on change |  | The name for the inference experiment. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment. |
| `Schedule` |  | `map` | optional, computed, provider-chosen, write-only |  | The duration for which you want the inference experiment to run. |
| `ShadowModeConfig` | shadow_mode_config | `map` | optional, computed, provider-chosen |  | The configuration of ShadowMode inference experiment type. Use this field to specify a production variant which takes all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant also specify the percentage of requests that Amazon SageMaker replicates. |
| `Status` |  | `string` | computed |  | The status of the inference experiment. |
| `StatusReason` | status_reason | `string` | optional, computed, provider-chosen, write-only |  | The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the inference experiment that you want to run. |

Supports update: yes

Discovery: supported
