# aws.trainingdataset

**CloudFormation type:** `AWS::CleanRoomsML::TrainingDataset`

Definition of AWS::CleanRoomsML::TrainingDataset Resource Type

Region attribute: `region`

**Import ID:** `<region>/TrainingDatasetArn` (AWS::CleanRoomsML::TrainingDataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset. |
| `TrainingData` | training_data | `list` | required, replaces on change |  |  |
| `TrainingDatasetArn` | training_dataset_arn | `string` | computed |  |  |

Supports update: yes

Discovery: supported
