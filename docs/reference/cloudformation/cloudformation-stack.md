# aws.cloudformation.stack

**CloudFormation type:** `AWS::CloudFormation::Stack`

The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.

Region attribute: `region`

**Import ID:** `<region>/StackId` (AWS::CloudFormation::Stack)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Capabilities` |  | `list` | optional, computed, provider-chosen |  |  |
| `ChangeSetId` | change_set_id | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisableRollback` | disable_rollback | `boolean` | optional, computed, provider-chosen |  |  |
| `EnableTerminationProtection` | enable_termination_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `LastUpdateTime` | last_update_time | `string` | computed |  |  |
| `NotificationARNs` | notification_ar_ns | `list` | optional, computed, provider-chosen |  |  |
| `Outputs` |  | `list` | computed |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  |  |
| `ParentId` | parent_id | `string` | computed |  |  |
| `RoleARN` | role_arn | `string` | optional, computed, provider-chosen |  |  |
| `RootId` | root_id | `string` | computed |  |  |
| `StackId` | stack_id | `string` | computed |  |  |
| `StackName` | stack_name | `string` | required, replaces on change |  |  |
| `StackPolicyBody` | stack_policy_body | `map` | optional, computed, provider-chosen |  |  |
| `StackPolicyURL` | stack_policy_url | `string` | optional, computed, provider-chosen, write-only |  |  |
| `StackStatus` | stack_status | `string` | computed |  |  |
| `StackStatusReason` | stack_status_reason | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen |  |  |
| `TemplateURL` | template_url | `string` | optional, computed, provider-chosen, write-only |  |  |
| `TimeoutInMinutes` | timeout_in_minutes | `integer` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
