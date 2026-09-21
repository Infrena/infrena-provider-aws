# aws.applicationinferenceprofile

**CloudFormation type:** `AWS::Bedrock::ApplicationInferenceProfile`

Definition of AWS::Bedrock::ApplicationInferenceProfile Resource Type

Region attribute: `region`

**Import ID:** `<region>/InferenceProfileIdentifier` (AWS::Bedrock::ApplicationInferenceProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the inference profile |
| `InferenceProfileArn` | inference_profile_arn | `string` | computed |  |  |
| `InferenceProfileId` | inference_profile_id | `string` | computed |  |  |
| `InferenceProfileIdentifier` | inference_profile_identifier | `string` | computed |  | Inference profile identifier. Supports both system-defined inference profile ids, and inference profile ARNs. |
| `InferenceProfileName` | inference_profile_name | `string` | required, replaces on change |  |  |
| `ModelSource` | model_source | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Various ways to encode a list of models in a CreateInferenceProfile request |
| `Models` |  | `list` | computed |  | List of model configuration |
| `Status` |  | `string` | computed |  | Status of the Inference Profile |
| `Tags` |  | `map` | tags map |  | List of Tags |
| `Type` | type_value | `string` | computed |  | Type of the Inference Profile |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp |

Supports update: yes

Discovery: supported
