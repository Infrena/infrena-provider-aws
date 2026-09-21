# aws.personalize.dataset

**CloudFormation type:** `AWS::Personalize::Dataset`

Resource schema for AWS::Personalize::Dataset.

Region attribute: `region`

**Import ID:** `<region>/DatasetArn` (AWS::Personalize::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetArn` | dataset_arn | `string` | computed |  | The ARN of the dataset |
| `DatasetGroupArn` | dataset_group_arn | `string` | required, replaces on change | aws.personalize.datasetgroup.DatasetGroupArn | The Amazon Resource Name (ARN) of the dataset group to add the dataset to |
| `DatasetImportJob` | dataset_import_job | `map` | optional, computed, provider-chosen |  | Initial DatasetImportJob for the created dataset |
| `DatasetType` | dataset_type | `string` | required, replaces on change |  | The type of dataset |
| `Name` |  | `string` | required, replaces on change |  | The name for the dataset |
| `SchemaArn` | schema_arn | `string` | required, replaces on change | aws.personalize.schema.SchemaArn | The ARN of the schema to associate with the dataset. The schema defines the dataset fields. |
| `Tags` |  | `map` | replaces on change, tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: yes

Discovery: supported
