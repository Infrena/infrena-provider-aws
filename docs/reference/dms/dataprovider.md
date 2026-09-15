# aws.dataprovider

**CloudFormation type:** `AWS::DMS::DataProvider`

Resource schema for AWS::DMS::DataProvider

Region attribute: `region`

**Import ID:** `<region>/DataProviderArn` (AWS::DMS::DataProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataProviderArn` | data_provider_arn | `string` | computed |  | The data provider ARN. |
| `DataProviderCreationTime` | data_provider_creation_time | `string` | computed |  | The data provider creation time. |
| `DataProviderIdentifier` | data_provider_identifier | `string` | optional, computed, provider-chosen, write-only |  | The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn |
| `DataProviderName` | data_provider_name | `string` | optional, computed, provider-chosen |  | The property describes a name to identify the data provider. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The optional description of the data provider. |
| `Engine` |  | `string` | required |  | The property describes a data engine for the data provider. |
| `ExactSettings` | exact_settings | `boolean` | optional, computed, provider-chosen, write-only |  | The property describes the exact settings which can be modified |
| `Settings` |  | `map` | optional, computed, provider-chosen |  | The property identifies the exact type of settings for the data provider. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
