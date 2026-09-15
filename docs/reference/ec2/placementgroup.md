# aws.placementgroup

**CloudFormation type:** `AWS::EC2::PlacementGroup`

Resource Type definition for AWS::EC2::PlacementGroup

Region attribute: `region`

**Import ID:** `<region>/GroupName` (AWS::EC2::PlacementGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupId` | group_id | `string` | computed |  | The ID of the placement group. |
| `GroupName` | group_name | `string` | computed |  | The Group Name of Placement Group. |
| `ParentGroupId` | parent_group_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of a parent placement group. Valid for strategies that support parent group linking. |
| `PartitionCount` | partition_count | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of partitions. Valid only when **Strategy** is set to `partition` |
| `SpreadLevel` | spread_level | `string` | optional, computed, provider-chosen, replaces on change |  | The Spread Level of Placement Group is an enum where it accepts either host or rack when strategy is spread |
| `Strategy` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The placement strategy. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: no

Discovery: supported
