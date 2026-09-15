# aws.mediatailor.channel

**CloudFormation type:** `AWS::MediaTailor::Channel`

Definition of AWS::MediaTailor::Channel Resource Type

Region attribute: `region`

**Import ID:** `<region>/ChannelName` (AWS::MediaTailor::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The ARN of the channel.</p> |
| `Audiences` |  | `list` | optional, computed, provider-chosen |  | <p>The list of audiences defined in channel.</p> |
| `ChannelName` | channel_name | `string` | required, replaces on change |  |  |
| `FillerSlate` | filler_slate | `map` | optional, computed, provider-chosen |  | <p>Slate VOD source configuration.</p> |
| `LogConfiguration` | log_configuration | `map` | optional, computed, provider-chosen |  | <p>The log configuration for the channel.</p> |
| `Outputs` |  | `list` | required |  | <p>The channel's output properties.</p> |
| `PlaybackMode` | playback_mode | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to assign to the channel. |
| `Tier` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `TimeShiftConfiguration` | time_shift_configuration | `map` | optional, computed, provider-chosen |  | <p>The configuration for time-shifted viewing.</p> |

Supports update: yes

Discovery: supported
