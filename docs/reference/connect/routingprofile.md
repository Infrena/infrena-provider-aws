# aws.routingprofile

**CloudFormation type:** `AWS::Connect::RoutingProfile`

Resource Type definition for AWS::Connect::RoutingProfile

Region attribute: `region`

**Import ID:** `<region>/RoutingProfileArn` (AWS::Connect::RoutingProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentAvailabilityTimer` | agent_availability_timer | `string` | optional, computed, provider-chosen |  | Whether agents with this routing profile will have their routing order calculated based on longest idle time or time since their last inbound contact. |
| `DefaultOutboundQueueArn` | default_outbound_queue_arn | `string` | required |  | The identifier of the default outbound queue for this routing profile. |
| `Description` |  | `string` | required |  | The description of the routing profile. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `ManualAssignmentQueueConfigs` | manual_assignment_queue_configs | `list` | optional, computed, provider-chosen |  | The manual assignment queues to associate with this routing profile. |
| `MediaConcurrencies` | media_concurrencies | `list` | required |  | The channels agents can handle in the Contact Control Panel (CCP) for this routing profile. |
| `Name` |  | `string` | required |  | The name of the routing profile. |
| `QueueConfigs` | queue_configs | `list` | optional, computed, provider-chosen |  | The queues to associate with this routing profile. |
| `RoutingProfileArn` | routing_profile_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the routing profile. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
