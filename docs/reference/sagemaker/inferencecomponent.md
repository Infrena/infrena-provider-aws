# aws.inferencecomponent

**CloudFormation type:** `AWS::SageMaker::InferenceComponent`

Resource Type definition for AWS::SageMaker::InferenceComponent

Region attribute: `region`

**Import ID:** `<region>/InferenceComponentArn` (AWS::SageMaker::InferenceComponent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DeploymentConfig` | deployment_config | `map` | optional, computed, provider-chosen, write-only |  | The deployment config for the inference component |
| `EndpointArn` | endpoint_arn | `string` | optional, computed, provider-chosen | aws.sagemaker.endpoint.EndpointArn | The Amazon Resource Name (ARN) of the endpoint the inference component is associated with |
| `EndpointName` | endpoint_name | `string` | required |  | The name of the endpoint the inference component is associated with |
| `FailureReason` | failure_reason | `string` | computed |  | The failure reason if the inference component is in a failed state |
| `InferenceComponentArn` | inference_component_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the inference component |
| `InferenceComponentName` | inference_component_name | `string` | optional, computed, provider-chosen |  | The name of the inference component |
| `InferenceComponentStatus` | inference_component_status | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `RuntimeConfig` | runtime_config | `map` | optional, computed, provider-chosen |  | The runtime config for the inference component |
| `Specification` |  | `map` | optional, computed, provider-chosen |  | The specification for the inference component, for an endpoint with a single instance type. Specify exactly one of Specification or Specifications. InstanceType is not accepted here; use Specifications for per instance type configuration. |
| `Specifications` |  | `list` | optional, computed, provider-chosen |  | A list of specification objects for the inference component, one per instance type. The service requires at least two entries; use the singular Specification for a single instance type. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of tags to apply to the resource |
| `VariantName` | variant_name | `string` | optional, computed, provider-chosen |  | The name of the endpoint variant the inference component is associated with |

Supports update: yes

Discovery: supported
