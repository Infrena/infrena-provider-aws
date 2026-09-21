# aws.fleetmetric

**CloudFormation type:** `AWS::IoT::FleetMetric`

An aggregated metric of certain devices in your fleet

Region attribute: `region`

**Import ID:** `<region>/MetricName` (AWS::IoT::FleetMetric)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AggregationField` | aggregation_field | `string` | optional, computed, provider-chosen |  | The aggregation field to perform aggregation and metric emission |
| `AggregationType` | aggregation_type | `map` | optional, computed, provider-chosen |  | Aggregation types supported by Fleet Indexing |
| `CreationDate` | creation_date | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of a fleet metric |
| `IndexName` | index_name | `string` | optional, computed, provider-chosen |  | The index name of a fleet metric |
| `LastModifiedDate` | last_modified_date | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `MetricArn` | metric_arn | `string` | computed |  | The Amazon Resource Number (ARN) of a fleet metric metric |
| `MetricName` | metric_name | `string` | required, replaces on change |  | The name of the fleet metric |
| `Period` |  | `integer` | optional, computed, provider-chosen |  | The period of metric emission in seconds |
| `QueryString` | query_string | `string` | optional, computed, provider-chosen |  | The Fleet Indexing query used by a fleet metric |
| `QueryVersion` | query_version | `string` | optional, computed, provider-chosen |  | The version of a Fleet Indexing query used by a fleet metric |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource |
| `Unit` |  | `string` | optional, computed, provider-chosen |  | The unit of data points emitted by a fleet metric |
| `Version` |  | `float` | computed |  | The version of a fleet metric |

Supports update: yes

Discovery: supported
