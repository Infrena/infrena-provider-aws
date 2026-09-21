# aws.iotfleetwise.fleet

**CloudFormation type:** `AWS::IoTFleetWise::Fleet`

Definition of AWS::IoTFleetWise::Fleet Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTFleetWise::Fleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | required, replaces on change |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `SignalCatalogArn` | signal_catalog_arn | `string` | required, replaces on change | aws.signalcatalog.Arn |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
