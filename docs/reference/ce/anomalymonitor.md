# aws.anomalymonitor

**CloudFormation type:** `AWS::CE::AnomalyMonitor`

AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.

Region attribute: `region`

**Import ID:** `<region>/MonitorArn` (AWS::CE::AnomalyMonitor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationDate` | creation_date | `string` | computed |  | The date when the monitor was created. |
| `DimensionalValueCount` | dimensional_value_count | `integer` | computed |  | The value for evaluated dimensions. |
| `LastEvaluatedDate` | last_evaluated_date | `string` | computed |  | The date when the monitor last evaluated for anomalies. |
| `LastUpdatedDate` | last_updated_date | `string` | computed |  | The date when the monitor was last updated. |
| `MonitorArn` | monitor_arn | `string` | computed |  | Monitor ARN |
| `MonitorDimension` | monitor_dimension | `string` | optional, computed, provider-chosen, replaces on change |  | The dimensions to evaluate |
| `MonitorName` | monitor_name | `string` | required |  | The name of the monitor. |
| `MonitorSpecification` | monitor_specification | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `MonitorType` | monitor_type | `string` | required, replaces on change |  |  |
| `ResourceTags` | resource_tags | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | Tags to assign to monitor. |

Supports update: yes

Discovery: supported
