# aws.rekognition.dataset

**CloudFormation type:** `AWS::Rekognition::Dataset`

The AWS::Rekognition::Dataset type creates an Amazon Rekognition Custom Labels dataset.

Region attribute: `region`

**Import ID:** `<region>/DatasetArn` (AWS::Rekognition::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetArn` | dataset_arn | `string` | computed |  | The ARN of the dataset. |
| `DatasetType` | dataset_type | `string` | required, replaces on change |  | The type of the dataset. Specify TRAIN to create a training dataset. Specify TEST to create a test dataset. |
| `ProjectArn` | project_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.rekognition.project.Arn | The ARN of the project to which the dataset belongs. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
