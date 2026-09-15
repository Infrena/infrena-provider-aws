# aws.signalcatalog

**CloudFormation type:** `AWS::IoTFleetWise::SignalCatalog`

Definition of AWS::IoTFleetWise::SignalCatalog Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::SignalCatalog)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `NodeCounts` | node_counts | `map` | optional, computed, provider-chosen |  |  |
| `Nodes` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
