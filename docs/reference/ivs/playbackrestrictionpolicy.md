# aws.playbackrestrictionpolicy

**CloudFormation type:** `AWS::IVS::PlaybackRestrictionPolicy`

Resource Type definition for AWS::IVS::PlaybackRestrictionPolicy.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::PlaybackRestrictionPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedCountries` | allowed_countries | `list` | optional, computed, provider-chosen |  | A list of country codes that control geoblocking restriction. Allowed values are the officially assigned ISO 3166-1 alpha-2 codes. Default: All countries (an empty array). |
| `AllowedOrigins` | allowed_origins | `list` | optional, computed, provider-chosen |  | A list of origin sites that control CORS restriction. Allowed values are the same as valid values of the Origin header defined at https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin |
| `Arn` |  | `string` | computed |  | Playback-restriction-policy identifier. |
| `EnableStrictOriginEnforcement` | enable_strict_origin_enforcement | `boolean` | optional, computed, provider-chosen |  | Whether channel playback is constrained by origin site. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Playback-restriction-policy name. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
