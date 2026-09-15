# aws.connectionalias

**CloudFormation type:** `AWS::WorkSpaces::ConnectionAlias`

Resource Type definition for AWS::WorkSpaces::ConnectionAlias

Region attribute: `region`

**Import ID:** `<region>/AliasId` (AWS::WorkSpaces::ConnectionAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasId` | alias_id | `string` | computed |  |  |
| `Associations` |  | `list` | computed |  |  |
| `ConnectionAliasState` | connection_alias_state | `string` | computed |  |  |
| `ConnectionString` | connection_string | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  |  |

Supports update: no

Discovery: supported
