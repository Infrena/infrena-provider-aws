# aws.gamelift.alias

**CloudFormation type:** `AWS::GameLift::Alias`

The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.

Region attribute: `region`

**Import ID:** `<region>/AliasId` (AWS::GameLift::Alias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasArn` | alias_arn | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift Alias resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift Alias ARN, the resource ID matches the AliasId value. |
| `AliasId` | alias_id | `string` | computed |  | Unique alias ID |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A human-readable description of the alias. |
| `Name` |  | `string` | required |  | A descriptive label that is associated with an alias. Alias names do not need to be unique. |
| `RoutingStrategy` | routing_strategy | `map` | required |  | A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
