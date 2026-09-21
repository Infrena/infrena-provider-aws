# aws.custommetric

**CloudFormation type:** `AWS::IoT::CustomMetric`

A custom metric published by your devices to Device Defender.

Region attribute: `region`

**Import ID:** `<region>/MetricName` (AWS::IoT::CustomMetric)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined. |
| `MetricArn` | metric_arn | `string` | computed |  | The Amazon Resource Number (ARN) of the custom metric. |
| `MetricName` | metric_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined. |
| `MetricType` | metric_type | `string` | required, replaces on change |  | The type of the custom metric. Types include string-list, ip-address-list, number-list, and number. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
