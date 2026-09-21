# aws.queuefleetassociation

**CloudFormation type:** `AWS::Deadline::QueueFleetAssociation`

Resource Type definition for AWS::Deadline::QueueFleetAssociation

Region attribute: `region`

**Import ID:** `<region>/FarmId|FleetId|QueueId` (AWS::Deadline::QueueFleetAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `FleetId` | fleet_id | `string` | required, replaces on change | aws.deadline.fleet.FleetId |  |
| `QueueId` | queue_id | `string` | required, replaces on change | aws.deadline.queue.QueueId |  |

Supports update: no

Discovery: supported (parent resource required)
