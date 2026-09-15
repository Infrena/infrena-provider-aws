# aws.sagemaker.endpoint

**CloudFormation type:** `AWS::SageMaker::Endpoint`

Resource Type definition for AWS::SageMaker::Endpoint

Region attribute: `region`

**Import ID:** `<region>/EndpointArn` (AWS::SageMaker::Endpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeploymentConfig` | deployment_config | `map` | optional, computed, provider-chosen |  | Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies. |
| `EndpointArn` | endpoint_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the endpoint. |
| `EndpointConfigName` | endpoint_config_name | `string` | required |  | The name of the endpoint configuration for the SageMaker endpoint. This is a required property. |
| `EndpointName` | endpoint_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the SageMaker endpoint. This name must be unique within an AWS Region. |
| `ExcludeRetainedVariantProperties` | exclude_retained_variant_properties | `list` | optional, computed, provider-chosen, write-only |  | Specifies a list of variant properties that you want to exclude when updating an endpoint. |
| `RetainAllVariantProperties` | retain_all_variant_properties | `boolean` | optional, computed, provider-chosen, write-only |  | When set to true, retains all variant properties for an endpoint when it is updated. |
| `RetainDeploymentConfig` | retain_deployment_config | `boolean` | optional, computed, provider-chosen, write-only |  | When set to true, retains the deployment configuration during endpoint updates. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
