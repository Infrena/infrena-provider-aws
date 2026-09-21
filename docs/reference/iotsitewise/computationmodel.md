# aws.computationmodel

**CloudFormation type:** `AWS::IoTSiteWise::ComputationModel`

Resource schema for AWS::IoTSiteWise::ComputationModel.

Region attribute: `region`

**Import ID:** `<region>/ComputationModelId` (AWS::IoTSiteWise::ComputationModel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputationModelArn` | computation_model_arn | `string` | computed |  | The ARN of the computation model. |
| `ComputationModelConfiguration` | computation_model_configuration | `map` | required |  | The configuration for the computation model. |
| `ComputationModelDataBinding` | computation_model_data_binding | `map` | required |  | The data binding for the computation model. |
| `ComputationModelDescription` | computation_model_description | `string` | optional, computed, provider-chosen |  | A description about the computation model. |
| `ComputationModelId` | computation_model_id | `string` | computed |  | The ID of the computation model. |
| `ComputationModelName` | computation_model_name | `string` | required |  | The name of the computation model. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
