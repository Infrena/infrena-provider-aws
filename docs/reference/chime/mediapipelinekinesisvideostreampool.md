# aws.mediapipelinekinesisvideostreampool

**CloudFormation type:** `AWS::Chime::MediaPipelineKinesisVideoStreamPool`

Resource Type definition for an Amazon Chime SDK Media Pipeline Kinesis Video Stream Pool

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Chime::MediaPipelineKinesisVideoStreamPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Kinesis Video Stream Pool. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The time at which the Kinesis Video Stream Pool was created. |
| `PoolId` | pool_id | `string` | computed |  | The unique identifier of the Kinesis Video Stream Pool. |
| `PoolName` | pool_name | `string` | required, replaces on change |  | The name of the Kinesis Video Stream Pool. |
| `PoolStatus` | pool_status | `string` | computed |  | The status of the Kinesis Video Stream Pool. |
| `StreamConfiguration` | stream_configuration | `map` | required |  | The configuration settings for the Kinesis video stream. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the Kinesis Video Stream Pool. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  | The time at which the Kinesis Video Stream Pool was last updated. |

Supports update: yes

Discovery: supported
