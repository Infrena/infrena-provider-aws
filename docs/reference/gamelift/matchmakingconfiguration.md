# aws.matchmakingconfiguration

**CloudFormation type:** `AWS::GameLift::MatchmakingConfiguration`

The AWS::GameLift::MatchmakingConfiguration resource creates an Amazon GameLift (GameLift) matchmaking configuration.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::GameLift::MatchmakingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptanceRequired` | acceptance_required | `boolean` | required |  | A flag that indicates whether a match that was created with this configuration must be accepted by the matched players |
| `AcceptanceTimeoutSeconds` | acceptance_timeout_seconds | `integer` | optional, computed, provider-chosen |  | The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required. |
| `AdditionalPlayerCount` | additional_player_count | `integer` | optional, computed, provider-chosen |  | The number of player slots in a match to keep open for future players. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking configuration resource and uniquely identifies it. |
| `BackfillMode` | backfill_mode | `string` | optional, computed, provider-chosen |  | The method used to backfill game sessions created with this matchmaking configuration. |
| `CreationTime` | creation_time | `string` | optional, computed, provider-chosen |  | A time stamp indicating when this data object was created. |
| `CustomEventData` | custom_event_data | `string` | optional, computed, provider-chosen |  | Information to attach to all events related to the matchmaking configuration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A descriptive label that is associated with matchmaking configuration. |
| `FlexMatchMode` | flex_match_mode | `string` | optional, computed, provider-chosen |  | Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution. |
| `GameProperties` | game_properties | `list` | optional, computed, provider-chosen |  | A set of custom properties for a game session, formatted as key:value pairs. |
| `GameSessionData` | game_session_data | `string` | optional, computed, provider-chosen |  | A set of custom game session properties, formatted as a single string value. |
| `GameSessionQueueArns` | game_session_queue_arns | `list` | optional, computed, provider-chosen | aws.gamesessionqueue.Arn | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it. |
| `Name` |  | `string` | required, replaces on change |  | A unique identifier for the matchmaking configuration. |
| `NotificationTarget` | notification_target | `string` | optional, computed, provider-chosen |  | An SNS topic ARN that is set up to receive matchmaking notifications. |
| `RequestTimeoutSeconds` | request_timeout_seconds | `integer` | required |  | The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out. |
| `RuleSetArn` | rule_set_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses. |
| `RuleSetName` | rule_set_name | `string` | required |  | A unique identifier for the matchmaking rule set to use with this configuration. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
