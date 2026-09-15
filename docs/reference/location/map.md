# aws.map

**CloudFormation type:** `AWS::Location::Map`

Definition of AWS::Location::Map Resource Type

Region attribute: `region`

**Import ID:** `<region>/MapName` (AWS::Location::Map)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Configuration` |  | `map` | required, replaces on change |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `MapArn` | map_arn | `string` | computed |  |  |
| `MapName` | map_name | `string` | required, replaces on change |  |  |
| `PricingPlan` | pricing_plan | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
