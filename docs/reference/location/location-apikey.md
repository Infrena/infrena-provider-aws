# aws.location.apikey

**CloudFormation type:** `AWS::Location::APIKey`

Definition of AWS::Location::APIKey Resource Type

Region attribute: `region`

**Import ID:** `<region>/KeyName` (AWS::Location::APIKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExpireTime` | expire_time | `string` | optional, computed, provider-chosen |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `ForceDelete` | force_delete | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `ForceUpdate` | force_update | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `KeyArn` | key_arn | `string` | computed |  |  |
| `KeyName` | key_name | `string` | required, replaces on change |  |  |
| `NoExpiry` | no_expiry | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `Restrictions` |  | `map` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
