# aws.configuredmodelalgorithm

**CloudFormation type:** `AWS::CleanRoomsML::ConfiguredModelAlgorithm`

Definition of AWS::CleanRoomsML::ConfiguredModelAlgorithm Resource Type

Region attribute: `region`

**Import ID:** `<region>/ConfiguredModelAlgorithmArn` (AWS::CleanRoomsML::ConfiguredModelAlgorithm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfiguredModelAlgorithmArn` | configured_model_algorithm_arn | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `InferenceContainerConfig` | inference_container_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm. |
| `TrainingContainerConfig` | training_container_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
