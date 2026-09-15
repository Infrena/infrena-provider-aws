# aws.datazone.datasource

**CloudFormation type:** `AWS::DataZone::DataSource`

A data source is used to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssetFormsInput` | asset_forms_input | `list` | optional, computed, provider-chosen, write-only |  | The metadata forms that are to be attached to the assets that this data source works with. |
| `Configuration` |  | `string` | optional, computed, provider-chosen, write-only |  | Specifies the configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration or sageMakerRunConfiguration. |
| `ConnectionId` | connection_id | `string` | computed |  | The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run |
| `ConnectionIdentifier` | connection_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the data source was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the data source. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the Amazon DataZone domain where the data source is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The ID of the Amazon DataZone domain where the data source is created. |
| `EnableSetting` | enable_setting | `string` | optional, computed, provider-chosen |  | Specifies whether the data source is enabled. |
| `EnvironmentId` | environment_id | `string` | computed |  | The unique identifier of the Amazon DataZone environment to which the data source publishes assets. |
| `EnvironmentIdentifier` | environment_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The unique identifier of the Amazon DataZone environment to which the data source publishes assets. |
| `Id` |  | `string` | computed |  | The unique identifier of the data source. |
| `LastRunAssetCount` | last_run_asset_count | `float` | computed |  | The number of assets created by the data source during its last run. |
| `LastRunAt` | last_run_at | `string` | computed |  | The timestamp that specifies when the data source was last run. |
| `LastRunStatus` | last_run_status | `string` | computed |  | The status of the last run of this data source. |
| `Name` |  | `string` | required |  | The name of the data source. |
| `ProjectId` | project_id | `string` | computed |  | The ID of the Amazon DataZone project to which the data source is added. |
| `ProjectIdentifier` | project_identifier | `string` | required, replaces on change, write-only |  | The identifier of the Amazon DataZone project in which you want to add the data source. |
| `PublishOnImport` | publish_on_import | `boolean` | optional, computed, provider-chosen |  | Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog. |
| `Recommendation` |  | `map` | optional, computed, provider-chosen |  | The recommendation to be updated as part of the UpdateDataSource action. |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  | The schedule of the data source runs. |
| `Status` |  | `string` | computed |  | The status of the data source. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the data source. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp of when this data source was updated. |

Supports update: yes

Discovery: supported (parent resource required)
