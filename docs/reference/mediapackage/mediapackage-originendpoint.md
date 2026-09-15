# aws.mediapackage.originendpoint

**CloudFormation type:** `AWS::MediaPackage::OriginEndpoint`

Resource schema for AWS::MediaPackage::OriginEndpoint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaPackage::OriginEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) assigned to the OriginEndpoint. |
| `Authorization` |  | `map` | optional, computed, provider-chosen |  | CDN Authorization credentials |
| `ChannelId` | channel_id | `string` | required | aws.mediapackage.channel.Id | The ID of the Channel the OriginEndpoint is associated with. |
| `CmafPackage` | cmaf_package | `map` | optional, computed, provider-chosen |  | A Common Media Application Format (CMAF) packaging configuration. |
| `DashPackage` | dash_package | `map` | optional, computed, provider-chosen |  | A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A short text description of the OriginEndpoint. |
| `HlsPackage` | hls_package | `map` | optional, computed, provider-chosen |  | An HTTP Live Streaming (HLS) packaging configuration. |
| `Id` |  | `string` | required, replaces on change |  | The ID of the OriginEndpoint. |
| `ManifestName` | manifest_name | `string` | optional, computed, provider-chosen |  | A short string appended to the end of the OriginEndpoint URL. |
| `MssPackage` | mss_package | `map` | optional, computed, provider-chosen |  | A Microsoft Smooth Streaming (MSS) packaging configuration. |
| `Origination` |  | `string` | optional, computed, provider-chosen |  | Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination |
| `StartoverWindowSeconds` | startover_window_seconds | `integer` | optional, computed, provider-chosen |  | Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource |
| `TimeDelaySeconds` | time_delay_seconds | `integer` | optional, computed, provider-chosen |  | Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint. |
| `Url` |  | `string` | computed |  | The URL of the packaged OriginEndpoint for consumption. |
| `Whitelist` |  | `list` | optional, computed, provider-chosen |  | A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint. |

Supports update: yes

Discovery: supported
