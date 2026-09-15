# aws.personalize.datasetgroup

**CloudFormation type:** `AWS::Personalize::DatasetGroup`

Resource Schema for AWS::Personalize::DatasetGroup.

Region attribute: `region`

**Import ID:** `<region>/DatasetGroupArn` (AWS::Personalize::DatasetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetGroupArn` | dataset_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the dataset group. |
| `Domain` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The domain of a Domain dataset group. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets. |
| `Name` |  | `string` | required, replaces on change |  | The name for the new dataset group. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: no

Discovery: supported
