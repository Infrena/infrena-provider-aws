# aws.calculatedattributedefinition

**CloudFormation type:** `AWS::CustomerProfiles::CalculatedAttributeDefinition`

A calculated attribute definition for Customer Profiles

Region attribute: `region`

**Import ID:** `<region>/DomainName|CalculatedAttributeName` (AWS::CustomerProfiles::CalculatedAttributeDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttributeDetails` | attribute_details | `map` | required |  | Mathematical expression and a list of attribute items specified in that expression. |
| `CalculatedAttributeName` | calculated_attribute_name | `string` | required, replaces on change |  | The unique name of the calculated attribute. |
| `Conditions` |  | `map` | optional, computed, provider-chosen |  | The conditions including range, object count, and threshold for the calculated attribute. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the calculated attribute definition was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the calculated attribute. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The display name of the calculated attribute. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the calculated attribute definition was most recently edited. |
| `Readiness` |  | `map` | computed |  | The readiness status of the calculated attribute. |
| `Statistic` |  | `string` | required |  | The aggregation operation to perform for the calculated attribute. |
| `Status` |  | `string` | computed |  | The status of the calculated attribute definition. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UseHistoricalData` | use_historical_data | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether to use historical data for the calculated attribute. |

Supports update: yes

Discovery: supported (parent resource required)
