# aws.anomalydetector

**CloudFormation type:** `AWS::APS::AnomalyDetector`

AnomalyDetector schema for cloudformation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::APS::AnomalyDetector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | required, replaces on change |  | The AnomalyDetector alias. |
| `Arn` |  | `string` | computed |  | The AnomalyDetector ARN. |
| `Configuration` |  | `map` | required |  | Determines the anomaly detector's algorithm and its configuration. |
| `EvaluationIntervalInSeconds` | evaluation_interval_in_seconds | `integer` | optional, computed, provider-chosen |  | The AnomalyDetector period of detection and metric generation. |
| `Labels` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to provide meta-data. |
| `MissingDataAction` | missing_data_action | `map` | optional, computed, provider-chosen |  | The action to perform when running the expression returns no data. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Workspace` |  | `string` | required, replaces on change |  | Required to identify a specific APS Workspace associated with this Anomaly Detector. |

Supports update: yes

Discovery: supported (parent resource required)
