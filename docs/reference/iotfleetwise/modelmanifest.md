# aws.modelmanifest

**CloudFormation type:** `AWS::IoTFleetWise::ModelManifest`

Definition of AWS::IoTFleetWise::ModelManifest Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::ModelManifest)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Nodes` |  | `list` | optional, computed, provider-chosen |  |  |
| `SignalCatalogArn` | signal_catalog_arn | `string` | required | aws.signalcatalog.Arn |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
