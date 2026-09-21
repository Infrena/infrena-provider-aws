# aws.modelcard

**CloudFormation type:** `AWS::SageMaker::ModelCard`

Resource Type definition for AWS::SageMaker::ModelCard.

Region attribute: `region`

**Import ID:** `<region>/ModelCardName` (AWS::SageMaker::ModelCard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Content` |  | `map` | required |  | The content of the model card. |
| `CreatedBy` | created_by | `map` | optional, computed, provider-chosen |  | Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card. |
| `CreationTime` | creation_time | `string` | computed |  | The date and time the model card was created. |
| `LastModifiedBy` | last_modified_by | `map` | optional, computed, provider-chosen |  | Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The date and time the model card was last modified. |
| `ModelCardArn` | model_card_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the successfully created model card. |
| `ModelCardName` | model_card_name | `string` | required, replaces on change |  | The unique name of the model card. |
| `ModelCardProcessingStatus` | model_card_processing_status | `string` | computed |  | The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps. |
| `ModelCardStatus` | model_card_status | `string` | required |  | The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval. |
| `ModelCardVersion` | model_card_version | `integer` | computed |  | A version of the model card. |
| `SecurityConfig` | security_config | `map` | optional, computed, provider-chosen, replaces on change |  | An optional Key Management Service key to encrypt, decrypt, and re-encrypt model card content for regulated workloads with highly sensitive data. |
| `Tags` |  | `map` | tags map |  | Key-value pairs used to manage metadata for model cards. |

Supports update: yes

Discovery: supported
