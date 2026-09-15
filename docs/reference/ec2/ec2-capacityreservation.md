# aws.ec2.capacityreservation

**CloudFormation type:** `AWS::EC2::CapacityReservation`

Resource Type definition for AWS::EC2::CapacityReservation

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::CapacityReservation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `AvailabilityZoneId` | availability_zone_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `AvailableInstanceCount` | available_instance_count | `integer` | computed |  |  |
| `CapacityAllocationSet` | capacity_allocation_set | `list` | computed |  |  |
| `CapacityReservationArn` | capacity_reservation_arn | `string` | computed |  |  |
| `CapacityReservationFleetId` | capacity_reservation_fleet_id | `string` | computed |  |  |
| `CommitmentInfo` | commitment_info | `map` | computed |  |  |
| `CreateDate` | create_date | `string` | computed |  |  |
| `DeliveryPreference` | delivery_preference | `string` | computed |  |  |
| `EbsOptimized` | ebs_optimized | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `EndDate` | end_date | `string` | optional, computed, provider-chosen |  |  |
| `EndDateType` | end_date_type | `string` | optional, computed, provider-chosen |  |  |
| `EphemeralStorage` | ephemeral_storage | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `InstanceCount` | instance_count | `integer` | required |  |  |
| `InstanceMatchCriteria` | instance_match_criteria | `string` | optional, computed, provider-chosen |  |  |
| `InstancePlatform` | instance_platform | `string` | required, replaces on change |  |  |
| `InstanceType` | instance_type | `string` | required, replaces on change |  |  |
| `OutPostArn` | out_post_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `OwnerId` | owner_id | `string` | computed |  |  |
| `PlacementGroupArn` | placement_group_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ReservationType` | reservation_type | `string` | computed |  |  |
| `StartDate` | start_date | `string` | computed |  |  |
| `State` |  | `string` | computed |  |  |
| `TagSpecifications` | tag_specifications | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tenancy` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `TotalInstanceCount` | total_instance_count | `integer` | computed |  |  |
| `UnusedReservationBillingOwnerId` | unused_reservation_billing_owner_id | `string` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
