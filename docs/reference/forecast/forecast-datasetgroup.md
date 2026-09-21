# aws.forecast.datasetgroup

**CloudFormation type:** `AWS::Forecast::DatasetGroup`

Represents a dataset group that holds a collection of related datasets

Region attribute: `region`

**Import ID:** `<region>/DatasetGroupArn` (AWS::Forecast::DatasetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetArns` | dataset_arns | `list` | optional, computed, provider-chosen | aws.forecast.dataset.Arn | An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group. |
| `DatasetGroupArn` | dataset_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the dataset group to delete. |
| `DatasetGroupName` | dataset_group_name | `string` | required, replaces on change |  | A name for the dataset group. |
| `Domain` |  | `string` | required |  | The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match. |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | The tags of Application Insights application. |

Supports update: yes

Discovery: supported
