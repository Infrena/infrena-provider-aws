# aws.realtimelogconfig

**CloudFormation type:** `AWS::CloudFront::RealtimeLogConfig`

A real-time log configuration.

Global type (no region attribute)

**Import ID:** `global/Arn` (AWS::CloudFront::RealtimeLogConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `EndPoints` | end_points | `list` | required |  | Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration. |
| `Fields` |  | `list` | required |  | A list of fields that are included in each real-time log record. In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream. |
| `Name` |  | `string` | required, replaces on change |  | The unique name of this real-time log configuration. |
| `SamplingRate` | sampling_rate | `float` | required |  | The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive. |

Supports update: yes

Discovery: supported
