# aws.agentalias

**CloudFormation type:** `AWS::Bedrock::AgentAlias`

Definition of AWS::Bedrock::AgentAlias Resource Type

Region attribute: `region`

**Import ID:** `<region>/AgentId|AgentAliasId` (AWS::Bedrock::AgentAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentAliasArn` | agent_alias_arn | `string` | computed |  | Arn representation of the Agent Alias. |
| `AgentAliasHistoryEvents` | agent_alias_history_events | `list` | computed |  | The list of history events for an alias for an Agent. |
| `AgentAliasId` | agent_alias_id | `string` | computed |  | Id for an Agent Alias generated at the server side. |
| `AgentAliasName` | agent_alias_name | `string` | required |  | Name for a resource. |
| `AgentAliasStatus` | agent_alias_status | `string` | computed |  | The statuses an Agent Alias can be in. |
| `AgentId` | agent_id | `string` | required, replaces on change | aws.bedrock.agent.AgentId | Identifier for a resource. |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Resource. |
| `RoutingConfiguration` | routing_configuration | `list` | optional, computed, provider-chosen |  | Routing configuration for an Agent alias. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |

Supports update: yes

Discovery: supported (parent resource required)
