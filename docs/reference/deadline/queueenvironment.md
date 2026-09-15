# aws.queueenvironment

**CloudFormation type:** `AWS::Deadline::QueueEnvironment`

Resource Type definition for AWS::Deadline::QueueEnvironment

Region attribute: `region`

**Import ID:** `<region>/FarmId|QueueId|QueueEnvironmentId` (AWS::Deadline::QueueEnvironment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `Name` |  | `string` | computed |  |  |
| `Priority` |  | `integer` | required |  |  |
| `QueueEnvironmentId` | queue_environment_id | `string` | computed |  |  |
| `QueueId` | queue_id | `string` | required, replaces on change | aws.deadline.queue.QueueId |  |
| `Template` |  | `string` | required |  |  |
| `TemplateType` | template_type | `string` | required |  |  |

Supports update: yes

Discovery: supported (parent resource required)
