# aws.events.endpoint

**CloudFormation type:** `AWS::Events::Endpoint`

Resource Type definition for AWS::Events::Endpoint.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Events::Endpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EndpointId` | endpoint_id | `string` | computed |  |  |
| `EndpointUrl` | endpoint_url | `string` | computed |  |  |
| `EventBuses` | event_buses | `list` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ReplicationConfig` | replication_config | `map` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `RoutingConfig` | routing_config | `map` | required |  |  |
| `State` |  | `string` | computed |  |  |
| `StateReason` | state_reason | `string` | computed |  |  |

Supports update: yes

Discovery: supported
