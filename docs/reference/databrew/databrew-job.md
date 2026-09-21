# aws.databrew.job

**CloudFormation type:** `AWS::DataBrew::Job`

Resource schema for AWS::DataBrew::Job.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Job)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataCatalogOutputs` | data_catalog_outputs | `list` | optional, computed, provider-chosen |  |  |
| `DatabaseOutputs` | database_outputs | `list` | optional, computed, provider-chosen |  |  |
| `DatasetName` | dataset_name | `string` | optional, computed, provider-chosen |  | Dataset name |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen |  | Encryption Key Arn |
| `EncryptionMode` | encryption_mode | `string` | optional, computed, provider-chosen |  | Encryption mode |
| `JobSample` | job_sample | `map` | optional, computed, provider-chosen |  | Job Sample |
| `LogSubscription` | log_subscription | `string` | optional, computed, provider-chosen |  | Log subscription |
| `MaxCapacity` | max_capacity | `integer` | optional, computed, provider-chosen |  | Max capacity |
| `MaxRetries` | max_retries | `integer` | optional, computed, provider-chosen |  | Max retries |
| `Name` |  | `string` | required, replaces on change |  | Job name |
| `OutputLocation` | output_location | `map` | optional, computed, provider-chosen |  | Output location |
| `Outputs` |  | `list` | optional, computed, provider-chosen |  |  |
| `ProfileConfiguration` | profile_configuration | `map` | optional, computed, provider-chosen |  | Profile Job configuration |
| `ProjectName` | project_name | `string` | optional, computed, provider-chosen |  | Project name |
| `Recipe` |  | `map` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role arn |
| `Tags` |  | `map` | tags map |  |  |
| `Timeout` |  | `integer` | optional, computed, provider-chosen |  | Timeout |
| `Type` | type_value | `string` | required, replaces on change |  | Job type |
| `ValidationConfigurations` | validation_configurations | `list` | optional, computed, provider-chosen |  | Data quality rules configuration |

Supports update: yes

Discovery: supported
