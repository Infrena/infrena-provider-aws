# aws.queuelimitassociation

**CloudFormation type:** `AWS::Deadline::QueueLimitAssociation`

Resource Type definition for AWS::Deadline::QueueLimitAssociation

Region attribute: `region`

**Import ID:** `<region>/FarmId|LimitId|QueueId` (AWS::Deadline::QueueLimitAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `LimitId` | limit_id | `string` | required, replaces on change | aws.limit.LimitId |  |
| `QueueId` | queue_id | `string` | required, replaces on change | aws.deadline.queue.QueueId |  |

Supports update: no

Discovery: supported (parent resource required)
