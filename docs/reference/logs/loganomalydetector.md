# aws.loganomalydetector

**CloudFormation type:** `AWS::Logs::LogAnomalyDetector`

The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.

Region attribute: `region`

**Import ID:** `<region>/AnomalyDetectorArn` (AWS::Logs::LogAnomalyDetector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | optional, computed, provider-chosen, write-only |  | Account ID for owner of detector |
| `AnomalyDetectorArn` | anomaly_detector_arn | `string` | computed |  | ARN of LogAnomalyDetector |
| `AnomalyDetectorStatus` | anomaly_detector_status | `string` | computed |  | Current status of detector. |
| `AnomalyVisibilityTime` | anomaly_visibility_time | `float` | optional, computed, provider-chosen |  |  |
| `CreationTimeStamp` | creation_time_stamp | `float` | computed |  | When detector was created. |
| `DetectorName` | detector_name | `string` | optional, computed, provider-chosen |  | Name of detector |
| `EvaluationFrequency` | evaluation_frequency | `string` | optional, computed, provider-chosen |  | How often log group is evaluated |
| `FilterPattern` | filter_pattern | `string` | optional, computed, provider-chosen |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the CMK to use when encrypting log data. |
| `LastModifiedTimeStamp` | last_modified_time_stamp | `float` | computed |  | When detector was lsat modified. |
| `LogGroupArnList` | log_group_arn_list | `list` | optional, computed, provider-chosen |  | List of Arns for the given log group |

Supports update: yes

Discovery: supported
