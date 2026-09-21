# aws.quicksight.datasource

**CloudFormation type:** `AWS::QuickSight::DataSource`

Definition of the AWS::QuickSight::DataSource Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|DataSourceId` (AWS::QuickSight::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlternateDataSourceParameters` | alternate_data_source_parameters | `list` | optional, computed, provider-chosen |  | <p>A set of alternate data source parameters that you want to share for the credentials |
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the data source.</p> |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that this data source was created.</p> |
| `Credentials` |  | `map` | optional, computed, provider-chosen, sensitive, write-only |  | <p>Data source credentials. This is a variant type structure. For this structure to be |
| `DataSourceId` | data_source_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.datasource.DataSourceId |  |
| `DataSourceParameters` | data_source_parameters | `map` | optional, computed, provider-chosen |  | <p>The parameters that Amazon QuickSight uses to connect to your underlying data source. |
| `ErrorInfo` | error_info | `map` | optional, computed, provider-chosen |  | <p>Error information for the data source creation or update.</p> |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, write-only | aws.folder.Arn |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The last time that this data source was updated.</p> |
| `Name` |  | `string` | required |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `SslProperties` | ssl_properties | `map` | optional, computed, provider-chosen |  | <p>Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |
| `VpcConnectionProperties` | vpc_connection_properties | `map` | optional, computed, provider-chosen |  | <p>VPC connection properties.</p> |

Supports update: yes

Discovery: supported
