# aws.frauddetector.detector

**CloudFormation type:** `AWS::FraudDetector::Detector`

A resource schema for a Detector in Amazon Fraud Detector.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::FraudDetector::Detector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the detector. |
| `AssociatedModels` | associated_models | `list` | optional, computed, provider-chosen |  | The models to associate with this detector. |
| `CreatedTime` | created_time | `string` | computed |  | The time when the detector was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the detector. |
| `DetectorId` | detector_id | `string` | required, replaces on change | aws.frauddetector.detector.DetectorId | The ID of the detector |
| `DetectorVersionId` | detector_version_id | `string` | computed |  | The active version ID of the detector |
| `DetectorVersionStatus` | detector_version_status | `string` | optional, computed, provider-chosen |  | The desired detector version status for the detector |
| `EventType` | event_type | `map` | required |  | The event type to associate this detector with. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The time when the detector was last updated. |
| `RuleExecutionMode` | rule_execution_mode | `string` | optional, computed, provider-chosen |  |  |
| `Rules` |  | `list` | required |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Tags associated with this detector. |

Supports update: yes

Discovery: supported
