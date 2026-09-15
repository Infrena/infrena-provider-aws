# aws.dataintegration

**CloudFormation type:** `AWS::AppIntegrations::DataIntegration`

Resource Type definition for AWS::AppIntegrations::DataIntegration

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AppIntegrations::DataIntegration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataIntegrationArn` | data_integration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the data integration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The data integration description. |
| `FileConfiguration` | file_configuration | `map` | optional, computed, provider-chosen |  | The configuration for what files should be pulled from the source. |
| `Id` |  | `string` | computed |  | The unique identifer of the data integration. |
| `KmsKey` | kms_key | `string` | required, replaces on change |  | The KMS key of the data integration. |
| `Name` |  | `string` | required |  | The name of the data integration. |
| `ObjectConfiguration` | object_configuration | `map` | optional, computed, provider-chosen |  | The configuration for what data should be pulled from the source. |
| `ScheduleConfig` | schedule_config | `map` | optional, computed, provider-chosen, replaces on change |  | The name of the data and how often it should be pulled from the source. |
| `SourceURI` | source_uri | `string` | required, replaces on change |  | The URI of the data source. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the data integration. |

Supports update: yes

Discovery: supported
