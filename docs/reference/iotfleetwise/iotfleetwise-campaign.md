# aws.iotfleetwise.campaign

**CloudFormation type:** `AWS::IoTFleetWise::Campaign`

Definition of AWS::IoTFleetWise::Campaign Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::Campaign)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CollectionScheme` | collection_scheme | `string` | required, replaces on change |  |  |
| `Compression` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DataDestinationConfigs` | data_destination_configs | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `DataExtraDimensions` | data_extra_dimensions | `list` | optional, computed, provider-chosen |  |  |
| `DataPartitions` | data_partitions | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DiagnosticsMode` | diagnostics_mode | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ExpiryTime` | expiry_time | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `PostTriggerCollectionDuration` | post_trigger_collection_duration | `float` | optional, computed, provider-chosen, replaces on change |  |  |
| `Priority` |  | `integer` | optional, computed, provider-chosen, replaces on change |  |  |
| `SignalCatalogArn` | signal_catalog_arn | `string` | required, replaces on change | aws.signalcatalog.Arn |  |
| `SignalsToCollect` | signals_to_collect | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `SignalsToFetch` | signals_to_fetch | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `SpoolingMode` | spooling_mode | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StartTime` | start_time | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TargetArn` | target_arn | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
