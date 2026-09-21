# aws.agentstatus

**CloudFormation type:** `AWS::Connect::AgentStatus`

Resource Type definition for AWS::Connect::AgentStatus

Region attribute: `region`

**Import ID:** `<region>/AgentStatusArn` (AWS::Connect::AgentStatus)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentStatusArn` | agent_status_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the agent status. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the status. |
| `DisplayOrder` | display_order | `integer` | optional, computed, provider-chosen |  | The display order of the status. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | Last modified region. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | Last modified time. |
| `Name` |  | `string` | required |  | The name of the status. |
| `ResetOrderNumber` | reset_order_number | `boolean` | optional, computed, provider-chosen |  | A number indicating the reset order of the agent status. |
| `State` |  | `string` | required |  | The state of the status. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | The type of agent status. |

Supports update: yes

Discovery: supported (parent resource required)
