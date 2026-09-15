# aws.capacityreservationfleet

**CloudFormation type:** `AWS::EC2::CapacityReservationFleet`

Resource Type definition for AWS::EC2::CapacityReservationFleet

Region attribute: `region`

**Import ID:** `<region>/CapacityReservationFleetId` (AWS::EC2::CapacityReservationFleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocationStrategy` | allocation_strategy | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CapacityReservationFleetId` | capacity_reservation_fleet_id | `string` | computed |  |  |
| `EndDate` | end_date | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `InstanceMatchCriteria` | instance_match_criteria | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `InstanceTypeSpecifications` | instance_type_specifications | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `NoRemoveEndDate` | no_remove_end_date | `boolean` | optional, computed, provider-chosen |  |  |
| `RemoveEndDate` | remove_end_date | `boolean` | optional, computed, provider-chosen |  |  |
| `TagSpecifications` | tag_specifications | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tenancy` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `TotalTargetCapacity` | total_target_capacity | `integer` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
