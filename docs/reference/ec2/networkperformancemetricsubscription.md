# aws.networkperformancemetricsubscription

**CloudFormation type:** `AWS::EC2::NetworkPerformanceMetricSubscription`

Resource Type definition for AWS::EC2::NetworkPerformanceMetricSubscription

Region attribute: `region`

**Import ID:** `<region>/Source|Destination|Metric|Statistic` (AWS::EC2::NetworkPerformanceMetricSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Destination` |  | `string` | required, replaces on change |  | The target Region or Availability Zone for the metric to subscribe to. |
| `Metric` |  | `string` | required, replaces on change |  | The metric type to subscribe to. |
| `Source` |  | `string` | required, replaces on change |  | The starting Region or Availability Zone for metric to subscribe to. |
| `Statistic` |  | `string` | required, replaces on change |  | The statistic to subscribe to. |

Supports update: no

Discovery: supported
