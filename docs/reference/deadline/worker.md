# aws.worker

**CloudFormation type:** `AWS::Deadline::Worker`

Definition of AWS::Deadline::Worker Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Worker)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the worker. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time the resource was created. |
| `CreatedBy` | created_by | `string` | computed |  | The user or system that created this resource. |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId | The farm ID. |
| `FleetId` | fleet_id | `string` | required, replaces on change | aws.deadline.fleet.FleetId | The fleet ID. |
| `HostProperties` | host_properties | `map` | optional, computed, provider-chosen, replaces on change |  | The host property details. |
| `Status` |  | `string` | computed |  | The status of the worker. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `WorkerId` | worker_id | `string` | computed |  | The worker ID. |

Supports update: yes

Discovery: supported (parent resource required)
