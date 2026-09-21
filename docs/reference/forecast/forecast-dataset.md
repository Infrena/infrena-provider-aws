# aws.forecast.dataset

**CloudFormation type:** `AWS::Forecast::Dataset`

Resource Type Definition for AWS::Forecast::Dataset

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Forecast::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DataFrequency` | data_frequency | `string` | optional, computed, provider-chosen, replaces on change |  | Frequency of data collection. This parameter is required for RELATED_TIME_SERIES |
| `DatasetName` | dataset_name | `string` | required, replaces on change |  | A name for the dataset |
| `DatasetType` | dataset_type | `string` | required, replaces on change |  | The dataset type |
| `Domain` |  | `string` | required, replaces on change |  | The domain associated with the dataset |
| `EncryptionConfig` | encryption_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Schema` |  | `map` | required, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported
