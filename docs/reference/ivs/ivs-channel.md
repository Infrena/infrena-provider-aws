# aws.ivs.channel

**CloudFormation type:** `AWS::IVS::Channel`

Resource Type definition for AWS::IVS::Channel

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Channel ARN is automatically generated on creation and assigned as the unique identifier. |
| `Authorized` |  | `boolean` | optional, computed, provider-chosen |  | Whether the channel is authorized. |
| `ContainerFormat` | container_format | `string` | optional, computed, provider-chosen |  | Indicates which content-packaging format is used (MPEG-TS or fMP4). If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS. |
| `IngestEndpoint` | ingest_endpoint | `string` | computed |  | Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software. |
| `InsecureIngest` | insecure_ingest | `boolean` | optional, computed, provider-chosen |  | Whether the channel allows insecure ingest. |
| `LatencyMode` | latency_mode | `string` | optional, computed, provider-chosen |  | Channel latency mode. |
| `MultitrackInputConfiguration` | multitrack_input_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Channel |
| `PlaybackUrl` | playback_url | `string` | computed |  | Channel Playback URL. |
| `Preset` |  | `string` | optional, computed, provider-chosen |  | Optional transcode preset for the channel. This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string (""). |
| `RecordingConfigurationArn` | recording_configuration_arn | `string` | optional, computed, provider-chosen | aws.recordingconfiguration.Arn | Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled). |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately. |

Supports update: yes

Discovery: supported
