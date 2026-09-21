# aws.matchmakingruleset

**CloudFormation type:** `AWS::GameLift::MatchmakingRuleSet`

The AWS::GameLift::MatchmakingRuleSet resource creates an Amazon GameLift (GameLift) matchmaking rule set.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::GameLift::MatchmakingRuleSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking rule set resource and uniquely identifies it. |
| `CreationTime` | creation_time | `string` | computed |  | A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds. |
| `Name` |  | `string` | required, replaces on change |  | A unique identifier for the matchmaking rule set. |
| `RuleSetBody` | rule_set_body | `string` | required, replaces on change |  | A collection of matchmaking rules, formatted as a JSON string. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
