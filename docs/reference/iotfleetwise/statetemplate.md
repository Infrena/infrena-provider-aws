# aws.statetemplate

**CloudFormation type:** `AWS::IoTFleetWise::StateTemplate`

Definition of AWS::IoTFleetWise::StateTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::StateTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DataExtraDimensions` | data_extra_dimensions | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `MetadataExtraDimensions` | metadata_extra_dimensions | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `SignalCatalogArn` | signal_catalog_arn | `string` | required, replaces on change | aws.signalcatalog.Arn |  |
| `StateTemplateProperties` | state_template_properties | `list` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
