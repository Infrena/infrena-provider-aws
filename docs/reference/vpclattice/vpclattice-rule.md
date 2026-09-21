# aws.vpclattice.rule

**CloudFormation type:** `AWS::VpcLattice::Rule`

Creates a listener rule. Each listener has a default rule for checking connection requests, but you can define additional rules. Each rule consists of a priority, one or more actions, and one or more conditions.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::Rule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `map` | required |  |  |
| `Arn` |  | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `ListenerIdentifier` | listener_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Match` |  | `map` | required |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Priority` |  | `integer` | required |  |  |
| `ServiceIdentifier` | service_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
