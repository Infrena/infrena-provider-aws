# aws.kinesisvideo.stream

**CloudFormation type:** `AWS::KinesisVideo::Stream`

Resource Type Definition for AWS::KinesisVideo::Stream

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::KinesisVideo::Stream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Kinesis Video stream. |
| `DataRetentionInHours` | data_retention_in_hours | `integer` | optional, computed, provider-chosen |  | The number of hours till which Kinesis Video will retain the data in the stream |
| `DeviceName` | device_name | `string` | optional, computed, provider-chosen |  | The name of the device that is writing to the stream. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data. |
| `MediaType` | media_type | `string` | optional, computed, provider-chosen |  | The media type of the stream. Consumers of the stream can use this information when processing the stream. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Kinesis Video stream. |
| `StreamStorageConfiguration` | stream_storage_configuration | `map` | optional, computed, provider-chosen |  | Configuration for the storage tier of the Kinesis Video Stream. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs associated with the Kinesis Video Stream. |

Supports update: yes

Discovery: supported
