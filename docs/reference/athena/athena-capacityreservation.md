# aws.athena.capacityreservation

**CloudFormation type:** `AWS::Athena::CapacityReservation`

Resource schema for AWS::Athena::CapacityReservation

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Athena::CapacityReservation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedDpus` | allocated_dpus | `integer` | computed |  | The number of DPUs Athena has provisioned and allocated for the reservation |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified capacity reservation |
| `CapacityAssignmentConfiguration` | capacity_assignment_configuration | `map` | optional, computed, provider-chosen |  | Assignment configuration to assign workgroups to a reservation |
| `CreationTime` | creation_time | `string` | computed |  | The date and time the reservation was created. |
| `LastSuccessfulAllocationTime` | last_successful_allocation_time | `string` | computed |  | The timestamp when the last successful allocated was made |
| `Name` |  | `string` | required, replaces on change |  | The reservation name. |
| `Status` |  | `string` | computed |  | The status of the reservation. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetDpus` | target_dpus | `integer` | required |  | The number of DPUs to request to be allocated to the reservation. |

Supports update: yes

Discovery: supported
