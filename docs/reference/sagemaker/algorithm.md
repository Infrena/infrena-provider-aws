# aws.algorithm

**CloudFormation type:** `AWS::SageMaker::Algorithm`

Resource Type definition for AWS::SageMaker::Algorithm

Region attribute: `region`

**Import ID:** `<region>/AlgorithmArn` (AWS::SageMaker::Algorithm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlgorithmArn` | algorithm_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the algorithm. |
| `AlgorithmDescription` | algorithm_description | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the algorithm. |
| `AlgorithmName` | algorithm_name | `string` | required, replaces on change |  | The name of the algorithm. |
| `CertifyForMarketplace` | certify_for_marketplace | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether to certify the algorithm so that it can be listed in AWS Marketplace. |
| `CreationTime` | creation_time | `string` | computed |  | A timestamp specifying when the algorithm was created. |
| `InferenceSpecification` | inference_specification | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TrainingSpecification` | training_specification | `map` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
