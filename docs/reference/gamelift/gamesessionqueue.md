# aws.gamesessionqueue

**CloudFormation type:** `AWS::GameLift::GameSessionQueue`

The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::GameLift::GameSessionQueue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it. |
| `CustomEventData` | custom_event_data | `string` | optional, computed, provider-chosen |  | Information that is added to all events that are related to this game session queue. |
| `Destinations` |  | `list` | optional, computed, provider-chosen |  | A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue. |
| `FilterConfiguration` | filter_configuration | `map` | optional, computed, provider-chosen |  | A list of locations where a queue is allowed to place new game sessions. |
| `Name` |  | `string` | required, replaces on change |  | A descriptive label that is associated with game session queue. Queue names must be unique within each Region. |
| `NotificationTarget` | notification_target | `string` | optional, computed, provider-chosen |  | An SNS topic ARN that is set up to receive game session placement notifications. |
| `PlayerLatencyPolicies` | player_latency_policies | `list` | optional, computed, provider-chosen |  | A set of policies that act as a sliding cap on player latency. |
| `PriorityConfiguration` | priority_configuration | `map` | optional, computed, provider-chosen |  | Custom settings to use when prioritizing destinations and locations for game session placements. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TimeoutInSeconds` | timeout_in_seconds | `integer` | optional, computed, provider-chosen |  | The maximum time, in seconds, that a new game session placement request remains in the queue. |

Supports update: yes

Discovery: supported
