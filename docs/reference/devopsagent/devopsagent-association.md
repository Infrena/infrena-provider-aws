# aws.devopsagent.association

**CloudFormation type:** `AWS::DevOpsAgent::Association`

Resource Type definition for AWS::DevOpsAgent::Association defining how the AgentSpace interacts with external services like GitHub, Slack, AWS accounts, and others.

Region attribute: `region`

**Import ID:** `<region>/AgentSpaceId|AssociationId` (AWS::DevOpsAgent::Association)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentSpaceId` | agent_space_id | `string` | required, replaces on change | aws.devopsagent.agentspace.AgentSpaceId | The unique identifier of the AgentSpace |
| `AssociationId` | association_id | `string` | computed |  | The unique identifier of the association |
| `Configuration` |  | `map` | required |  | The configuration that directs how AgentSpace interacts with the given service |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the association was created |
| `LinkedAssociationIds` | linked_association_ids | `list` | optional, computed, provider-chosen, write-only | aws.devopsagent.association.AssociationId | Set of linked association IDs for parent-child relationships |
| `ServiceId` | service_id | `string` | required | aws.devopsagent.service.ServiceId | The identifier for the associated service. For SourceAws and Aws configurations, this must be 'aws'. For all other service types, this is a UUID generated from the RegisterService command |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the association was last updated |

Supports update: yes

Discovery: supported (parent resource required)
