# aws.recordingconfiguration

**CloudFormation type:** `AWS::IVS::RecordingConfiguration`

Resource Type definition for AWS::IVS::RecordingConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::RecordingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier. |
| `DestinationConfiguration` | destination_configuration | `map` | required, replaces on change |  | Recording Destination Configuration. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Recording Configuration Name. |
| `RecordingReconnectWindowSeconds` | recording_reconnect_window_seconds | `integer` | optional, computed, provider-chosen, replaces on change |  | Recording Reconnect Window Seconds. (0 means disabled) |
| `RenditionConfiguration` | rendition_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Rendition Configuration describes which renditions should be recorded for a stream. |
| `State` |  | `string` | computed |  | Recording Configuration State. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the asset model. |
| `ThumbnailConfiguration` | thumbnail_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Recording Thumbnail Configuration. |

Supports update: yes

Discovery: supported
