# aws.iotsitewise.dataset

**CloudFormation type:** `AWS::IoTSiteWise::Dataset`

Resource schema for AWS::IoTSiteWise::Dataset.

Region attribute: `region`

**Import ID:** `<region>/DatasetId` (AWS::IoTSiteWise::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetArn` | dataset_arn | `string` | computed |  | The ARN of the dataset. |
| `DatasetDescription` | dataset_description | `string` | optional, computed, provider-chosen |  | A description about the dataset, and its functionality. |
| `DatasetId` | dataset_id | `string` | computed |  | The ID of the dataset. |
| `DatasetName` | dataset_name | `string` | required |  | The name of the dataset. |
| `DatasetSource` | dataset_source | `map` | required |  | The data source for the dataset. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
