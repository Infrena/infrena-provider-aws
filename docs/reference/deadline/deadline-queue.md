# aws.deadline.queue

**CloudFormation type:** `AWS::Deadline::Queue`

Resource Type definition for AWS::Deadline::Queue

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Queue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedStorageProfileIds` | allowed_storage_profile_ids | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `DefaultBudgetAction` | default_budget_action | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `JobAttachmentSettings` | job_attachment_settings | `map` | optional, computed, provider-chosen |  |  |
| `JobRunAsUser` | job_run_as_user | `map` | optional, computed, provider-chosen |  |  |
| `QueueId` | queue_id | `string` | computed |  |  |
| `RequiredFileSystemLocationNames` | required_file_system_location_names | `list` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `SchedulingConfiguration` | scheduling_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
