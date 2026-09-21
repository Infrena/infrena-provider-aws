# aws.devopsagent.trigger

**CloudFormation type:** `AWS::DevOpsAgent::Trigger`

Resource Type definition for AWS::DevOpsAgent::Trigger. A trigger defines an automated action that fires on a schedule within an Agent Space.

Region attribute: `region`

**Import ID:** `<region>/AgentSpaceId|TriggerId` (AWS::DevOpsAgent::Trigger)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `map` | required, replaces on change |  | The action to perform when the trigger fires. A JSON object containing actionType and task. |
| `AgentSpaceId` | agent_space_id | `string` | required, replaces on change | aws.devopsagent.agentspace.AgentSpaceId | The unique identifier of the parent Agent Space. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the trigger. Nested under the parent Agent Space: arn:<partition>:aidevops:<region>:<account-id>:agentspace/<agentspace-id>/trigger/<trigger-id>. |
| `Condition` |  | `map` | required, replaces on change |  | The condition that causes the trigger to fire. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the trigger was created. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the trigger. Active triggers fire on schedule; Inactive triggers are paused. |
| `TriggerId` | trigger_id | `string` | computed |  | The unique identifier of the trigger (assigned by the service on Create). |
| `Type` | type_value | `string` | required, replaces on change |  | The type of trigger. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the trigger was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
