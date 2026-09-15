# aws.iot.stream

**CloudFormation type:** `AWS::IoT::Stream`

Resource Type definition for AWS IoT Stream. A stream is a publicly addressable resource that is an abstraction for a list of files that can be transferred to an IoT device.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IoT::Stream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the stream. |
| `CreatedAt` | created_at | `string` | computed |  | The date when the stream was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the stream. |
| `Files` |  | `list` | required |  | The files to stream. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The date when the stream was last updated. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | An IAM role that allows the IoT service principal to access your S3 files. |
| `StreamId` | stream_id | `string` | required, replaces on change | aws.iot.stream.StreamId | The stream ID. |
| `StreamVersion` | stream_version | `integer` | computed |  | The stream version. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata which can be used to manage streams. |

Supports update: yes

Discovery: supported
