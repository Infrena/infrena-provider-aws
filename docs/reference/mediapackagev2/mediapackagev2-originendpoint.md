# aws.mediapackagev2.originendpoint

**CloudFormation type:** `AWS::MediaPackageV2::OriginEndpoint`

<p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaPackageV2::OriginEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) associated with the resource.</p> |
| `ChannelGroupName` | channel_group_name | `string` | required, replaces on change |  |  |
| `ChannelName` | channel_name | `string` | required, replaces on change |  |  |
| `ContainerType` | container_type | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  | <p>The date and time the origin endpoint was created.</p> |
| `DashManifestUrls` | dash_manifest_urls | `list` | computed |  |  |
| `DashManifests` | dash_manifests | `list` | optional, computed, provider-chosen |  | <p>A DASH manifest configuration.</p> |
| `Description` |  | `string` | optional, computed, provider-chosen |  | <p>Enter any descriptive text that helps you to identify the origin endpoint.</p> |
| `ForceEndpointErrorConfiguration` | force_endpoint_error_configuration | `map` | optional, computed, provider-chosen |  | <p>The failover settings for the endpoint.</p> |
| `HlsManifestUrls` | hls_manifest_urls | `list` | computed |  |  |
| `HlsManifests` | hls_manifests | `list` | optional, computed, provider-chosen |  | <p>An HTTP live streaming (HLS) manifest configuration.</p> |
| `LowLatencyHlsManifestUrls` | low_latency_hls_manifest_urls | `list` | computed |  |  |
| `LowLatencyHlsManifests` | low_latency_hls_manifests | `list` | optional, computed, provider-chosen |  | <p>A low-latency HLS manifest configuration.</p> |
| `ModifiedAt` | modified_at | `string` | computed |  | <p>The date and time the origin endpoint was modified.</p> |
| `MssManifestUrls` | mss_manifest_urls | `list` | computed |  |  |
| `MssManifests` | mss_manifests | `list` | optional, computed, provider-chosen |  | <p>The Microsoft Smooth Streaming (MSS) manifest configurations associated with this origin endpoint.</p> |
| `OriginEndpointName` | origin_endpoint_name | `string` | required, replaces on change |  |  |
| `Segment` |  | `map` | optional, computed, provider-chosen |  | <p>The segment configuration, including the segment name, duration, and other configuration values.</p> |
| `StartoverWindowSeconds` | startover_window_seconds | `integer` | optional, computed, provider-chosen |  | <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p> |
| `StreamNameOutputMode` | stream_name_output_mode | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UriSeparator` | uri_separator | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
