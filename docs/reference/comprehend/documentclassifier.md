# aws.documentclassifier

**CloudFormation type:** `AWS::Comprehend::DocumentClassifier`

Document Classifier enables training document classifier models.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Comprehend::DocumentClassifier)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DataAccessRoleArn` | data_access_role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `DocumentClassifierName` | document_classifier_name | `string` | required, replaces on change |  |  |
| `InputDataConfig` | input_data_config | `map` | required, replaces on change |  |  |
| `LanguageCode` | language_code | `string` | required, replaces on change |  |  |
| `Mode` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ModelKmsKeyId` | model_kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ModelPolicy` | model_policy | `string` | optional, computed, provider-chosen |  |  |
| `OutputDataConfig` | output_data_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VersionName` | version_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `VolumeKmsKeyId` | volume_kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
