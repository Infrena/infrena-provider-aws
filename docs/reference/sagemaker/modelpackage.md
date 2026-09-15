# aws.modelpackage

**CloudFormation type:** `AWS::SageMaker::ModelPackage`

Resource Type definition for AWS::SageMaker::ModelPackage

Region attribute: `region`

**Import ID:** `<region>/ModelPackageArn` (AWS::SageMaker::ModelPackage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalInferenceSpecifications` | additional_inference_specifications | `list` | optional, computed, provider-chosen |  | An array of additional Inference Specification objects. |
| `AdditionalInferenceSpecificationsToAdd` | additional_inference_specifications_to_add | `list` | optional, computed, provider-chosen, write-only |  | An array of additional Inference Specification objects. |
| `ApprovalDescription` | approval_description | `string` | optional, computed, provider-chosen |  | A description provided for the model approval. |
| `CertifyForMarketplace` | certify_for_marketplace | `boolean` | optional, computed, provider-chosen |  | Whether to certify the model package for listing on AWS Marketplace. |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A unique token that guarantees that the call to this API is idempotent. |
| `CreationTime` | creation_time | `string` | computed |  | The time at which the model package was created. |
| `CustomerMetadataProperties` | customer_metadata_properties | `map` | optional, computed, provider-chosen |  | The metadata properties associated with the model package versions. |
| `Domain` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The machine learning domain of the model package you specified. |
| `DriftCheckBaselines` | drift_check_baselines | `map` | optional, computed, provider-chosen, replaces on change |  | Represents the drift check baselines that can be used when the model monitor is set using the model package. |
| `InferenceSpecification` | inference_specification | `map` | optional, computed, provider-chosen, replaces on change |  | Details about inference jobs that can be run with models based on this model package. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the model package was last modified. |
| `MetadataProperties` | metadata_properties | `map` | optional, computed, provider-chosen, replaces on change |  | Metadata properties of the tracking entity, trial, or trial component. |
| `ModelApprovalStatus` | model_approval_status | `string` | optional, computed, provider-chosen |  | The approval status of the model package. |
| `ModelCard` | model_card | `map` | optional, computed, provider-chosen |  | The model card associated with the model package. |
| `ModelMetrics` | model_metrics | `map` | optional, computed, provider-chosen, replaces on change |  | A structure that contains model metrics reports. |
| `ModelPackageArn` | model_package_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the model package group. |
| `ModelPackageDescription` | model_package_description | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the model package. |
| `ModelPackageGroupName` | model_package_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the model package group. |
| `ModelPackageName` | model_package_name | `string` | optional, computed, provider-chosen |  | The name or arn of the model package. |
| `ModelPackageStatus` | model_package_status | `string` | computed |  | The current status of the model package. |
| `ModelPackageStatusDetails` | model_package_status_details | `map` | optional, computed, provider-chosen |  | Details about the current status of the model package. |
| `ModelPackageVersion` | model_package_version | `integer` | optional, computed, provider-chosen |  | The version of the model package. |
| `SamplePayloadUrl` | sample_payload_url | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Simple Storage Service (Amazon S3) path where the sample payload are stored pointing to single gzip compressed tar archive. |
| `SecurityConfig` | security_config | `map` | optional, computed, provider-chosen, replaces on change |  | An optional AWS Key Management Service key to encrypt, decrypt, and re-encrypt model package information for regulated workloads with highly sensitive data. |
| `SkipModelValidation` | skip_model_validation | `string` | optional, computed, provider-chosen |  | Indicates if you want to skip model validation. |
| `SourceAlgorithmSpecification` | source_algorithm_specification | `map` | optional, computed, provider-chosen, replaces on change |  | Details about the algorithm that was used to create the model package. |
| `SourceUri` | source_uri | `string` | optional, computed, provider-chosen |  | The URI of the source for the model package. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Task` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The machine learning task your model package accomplishes. |
| `ValidationSpecification` | validation_specification | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies configurations for one or more transform jobs that Amazon SageMaker runs to test the model package. |

Supports update: yes

Discovery: supported
