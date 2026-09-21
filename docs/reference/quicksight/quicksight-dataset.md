# aws.quicksight.dataset

**CloudFormation type:** `AWS::QuickSight::DataSet`

Definition of the AWS::QuickSight::DataSet Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|DataSetId` (AWS::QuickSight::DataSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the resource.</p> |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ColumnGroups` | column_groups | `list` | optional, computed, provider-chosen |  | <p>Groupings of columns that work together in certain Amazon QuickSight features. Currently, only geospatial hierarchy is supported.</p> |
| `ColumnLevelPermissionRules` | column_level_permission_rules | `list` | optional, computed, provider-chosen |  | <p>A set of one or more definitions of a <code> |
| `ConsumedSpiceCapacityInBytes` | consumed_spice_capacity_in_bytes | `float` | computed |  | <p>The amount of SPICE capacity used by this dataset. This is 0 if the dataset isn't |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that this dataset was created.</p> |
| `DataPrepConfiguration` | data_prep_configuration | `map` | optional, computed, provider-chosen |  |  |
| `DataSetId` | data_set_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.dataset.DataSetId |  |
| `DataSetRefreshProperties` | data_set_refresh_properties | `map` | optional, computed, provider-chosen |  | <p>The refresh properties of a dataset.</p> |
| `DataSetUsageConfiguration` | data_set_usage_configuration | `map` | optional, computed, provider-chosen |  | <p>The usage configuration to apply to child datasets that reference this dataset as a source.</p> |
| `DatasetParameters` | dataset_parameters | `list` | optional, computed, provider-chosen |  | <p>The parameter declarations of the dataset.</p> |
| `FieldFolders` | field_folders | `map` | optional, computed, provider-chosen, write-only |  |  |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, write-only | aws.folder.Arn | <p>When you create the dataset, Amazon QuickSight adds the dataset to these folders.</p> |
| `ImportMode` | import_mode | `string` | optional, computed, provider-chosen |  |  |
| `IngestionWaitPolicy` | ingestion_wait_policy | `map` | optional, computed, provider-chosen, write-only |  | <p>Wait policy to use when creating/updating dataset. Default is to wait for SPICE ingestion to finish with timeout of 36 hours.</p> |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The last time that this dataset was updated.</p> |
| `LogicalTableMap` | logical_table_map | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  | <p>The display name for the dataset.</p> |
| `OutputColumns` | output_columns | `list` | computed |  | <p>The list of columns after all transforms. These columns are available in templates, |
| `PerformanceConfiguration` | performance_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  | <p>A list of resource permissions on the dataset.</p> |
| `PhysicalTableMap` | physical_table_map | `map` | optional, computed, provider-chosen |  |  |
| `RowLevelPermissionDataSet` | row_level_permission_data_set | `map` | optional, computed, provider-chosen |  | <p>Information about a dataset that contains permissions for row-level security (RLS). |
| `RowLevelPermissionTagConfiguration` | row_level_permission_tag_configuration | `map` | optional, computed, provider-chosen |  | <p>The configuration of tags on a dataset to set row-level security. </p> |
| `SemanticModelConfiguration` | semantic_model_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p> |
| `UseAs` | use_as | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
