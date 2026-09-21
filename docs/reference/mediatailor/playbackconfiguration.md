# aws.playbackconfiguration

**CloudFormation type:** `AWS::MediaTailor::PlaybackConfiguration`

Resource schema for AWS::MediaTailor::PlaybackConfiguration

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::MediaTailor::PlaybackConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdConditioningConfiguration` | ad_conditioning_configuration | `map` | optional, computed, provider-chosen |  | <p>The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns.</p> |
| `AdDecisionServerConfiguration` | ad_decision_server_configuration | `map` | optional, computed, provider-chosen |  | The configuration for the request to the specified Ad Decision Server URL. |
| `AdDecisionServerUrl` | ad_decision_server_url | `string` | required |  | The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters. |
| `AdsPersonalizationConcurrency` | ads_personalization_concurrency | `map` | optional, computed, provider-chosen |  | The settings that control how many concurrent requests MediaTailor makes to the ad decision server (ADS). |
| `AdsPersonalizationTimeouts` | ads_personalization_timeouts | `map` | optional, computed, provider-chosen |  | The ad decision server (ADS) request timeouts and personalization time budgets for live, VOD, and prefetch workflows. |
| `AvailSuppression` | avail_suppression | `map` | optional, computed, provider-chosen |  | The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html). |
| `Bumper` |  | `map` | optional, computed, provider-chosen |  | The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html). |
| `CdnConfiguration` | cdn_configuration | `map` | optional, computed, provider-chosen |  | The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management. |
| `ConfigurationAliases` | configuration_aliases | `map` | optional, computed, provider-chosen |  | The predefined aliases for dynamic variables. |
| `DashConfiguration` | dash_configuration | `map` | optional, computed, provider-chosen |  | The configuration for DASH PUT operations. |
| `FunctionMapping` | function_mapping | `map` | optional, computed, provider-chosen |  | A map of event names to function identifiers for custom processing during session lifecycle events. |
| `HlsConfiguration` | hls_configuration | `map` | optional, computed, provider-chosen |  | The configuration for HLS content. |
| `InsertionMode` | insertion_mode | `string` | optional, computed, provider-chosen |  | The setting that controls whether players can use stitched or guided ad insertion. The default, STITCHED_ONLY, forces all player sessions to use stitched (server-side) ad insertion. Choosing PLAYER_SELECT allows players to select either stitched or guided ad insertion at session-initialization time. The default for players that do not specify an insertion mode is stitched. |
| `LivePreRollConfiguration` | live_pre_roll_configuration | `map` | optional, computed, provider-chosen |  | The configuration for pre-roll ad insertion. |
| `LogConfiguration` | log_configuration | `map` | optional, computed, provider-chosen |  | The configuration that defines where AWS Elemental MediaTailor sends logs for the playback configuration. |
| `ManifestProcessingRules` | manifest_processing_rules | `map` | optional, computed, provider-chosen |  | The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor. |
| `Name` |  | `string` | required, replaces on change |  | The identifier for the playback configuration. |
| `PersonalizationThresholdSeconds` | personalization_threshold_seconds | `integer` | optional, computed, provider-chosen |  | Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html). |
| `PlaybackConfigurationArn` | playback_configuration_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the playback configuration. |
| `PlaybackEndpointPrefix` | playback_endpoint_prefix | `string` | computed |  | The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting. |
| `SessionInitializationEndpointPrefix` | session_initialization_endpoint_prefix | `string` | computed |  | The URL that the player uses to initialize a session that uses client-side reporting. |
| `SlateAdUrl` | slate_ad_url | `string` | optional, computed, provider-chosen |  | The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video. |
| `Tags` |  | `map` | tags map |  | The tags to assign to the playback configuration. |
| `TranscodeProfileName` | transcode_profile_name | `string` | optional, computed, provider-chosen |  | The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support. |
| `VideoContentSourceUrl` | video_content_source_url | `string` | required |  | The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters. |

Supports update: yes

Discovery: supported
