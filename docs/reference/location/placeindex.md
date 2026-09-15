# aws.placeindex

**CloudFormation type:** `AWS::Location::PlaceIndex`

Definition of AWS::Location::PlaceIndex Resource Type

Region attribute: `region`

**Import ID:** `<region>/IndexName` (AWS::Location::PlaceIndex)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `DataSource` | data_source | `string` | required, replaces on change |  |  |
| `DataSourceConfiguration` | data_source_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IndexArn` | index_arn | `string` | computed |  |  |
| `IndexName` | index_name | `string` | required, replaces on change |  |  |
| `PricingPlan` | pricing_plan | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
